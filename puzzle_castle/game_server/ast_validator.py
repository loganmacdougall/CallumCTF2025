import ast
import math, itertools, functools, string, re
from typing import Dict, Any

from world import World

# -------------------- Policy --------------------
FORBIDDEN_NAMES = {
    "open", "exec", "eval", "compile", "globals", "locals", "vars",
    "__import__", "__builtins__", "input", "help", "quit", "exit",
    "os", "sys", "subprocess", "shutil", "socket", "requests", "importlib",
}

FORBIDDEN_ATTRS = {
    "__globals__", "__class__", "__mro__",
    "__subclasses__", "__dict__", "__getattribute__",
}

ALLOWED_NODES = {
    ast.Module, ast.FunctionDef, ast.AsyncFunctionDef, ast.Assign, ast.AugAssign,
    ast.Expr, ast.Return, ast.Call, ast.Name, ast.Load, ast.Store,
    ast.Constant, ast.BinOp, ast.UnaryOp, ast.Compare, ast.If, ast.For,
    ast.While, ast.Break, ast.Continue, ast.List, ast.Tuple, ast.Dict,
    ast.Set, ast.ListComp, ast.DictComp, ast.SetComp, ast.GeneratorExp,
    ast.Attribute, ast.Subscript, ast.Slice, ast.IfExp, ast.Pass,
    ast.Raise, ast.Try, ast.ExceptHandler, ast.With, ast.Assert,
    ast.Lambda, ast.BoolOp,
    ast.Add, ast.Sub, ast.Mult, ast.Div, ast.Mod, ast.Pow, ast.UAdd, ast.USub,
    ast.MatMult, ast.LShift, ast.RShift, 
    ast.Eq, ast.NotEq, ast.Lt, ast.LtE, ast.Gt, ast.GtE, ast.Is, ast.IsNot,
    ast.arguments, ast.BitOr, ast.BitAnd, ast.BitXor, ast.FloorDiv, ast.Invert,
    ast.In, ast.NotIn, ast.Not, ast.And, ast.Or, ast.Yield, ast.YieldFrom, ast.Await,
    ast.Slice, ast.Starred, ast.arg, ast.Nonlocal, ast.comprehension,
    ast.JoinedStr, ast.FormattedValue
}

class UnsafeCodeError(Exception):
    pass

class SafeASTVisitor(ast.NodeVisitor):
    def generic_visit(self, node):
        if type(node) not in ALLOWED_NODES:
            raise UnsafeCodeError(f"Disallowed AST node: {type(node).__name__}")
        super().generic_visit(node)

    def visit_Import(self, node):
        raise UnsafeCodeError("Import statements are not allowed")

    def visit_ImportFrom(self, node):
        raise UnsafeCodeError("Import-from statements are not allowed")

    def visit_Name(self, node):
        if node.id in FORBIDDEN_NAMES:
            raise UnsafeCodeError(f"Use of name '{node.id}' is forbidden")
        if node.id.startswith("__"):
            raise UnsafeCodeError(f"Use of dunder name '{node.id}' is forbidden")
        self.generic_visit(node)

    def visit_Attribute(self, node):
        if node.attr in FORBIDDEN_ATTRS or node.attr.startswith("__"):
            raise UnsafeCodeError(f"Access to attribute '{node.attr}' is forbidden")
        self.generic_visit(node)

    def visit_Call(self, node):
        if isinstance(node.func, ast.Name) and node.func.id in FORBIDDEN_NAMES:
            raise UnsafeCodeError(f"Calling '{node.func.id}' is forbidden")
        self.generic_visit(node)

    def visit_Global(self, node):
        raise UnsafeCodeError("Global statement is forbidden. If you're attempting to create a closure, try using 'nonlocal' instead")


# -------------------- Runtime environment --------------------
SAFE_BUILTINS: Dict[str, Any] = {
    "abs": abs, "min": min, "max": max, "sum": sum, "round": round,
    "int": int, "float": float, "str": str, "bool": bool, "complex": complex,
    "len": len, "range": range, "enumerate": enumerate, "zip": zip,
    "map": map, "filter": filter, "sorted": sorted, "reversed": reversed,
    "all": all, "any": any, "next": next, "iter": iter, "repr": repr,
    "divmod": divmod, "pow": pow, "hash": hash
}

SAFE_MODULES = {
    "math": math,
    "itertools": itertools,
    "functools": functools,
    "string": string,
    "re": re,
}

# -------------------- API --------------------
def analyze_code(code_str: str) -> ast.Module:
    """Parse and validate code. Raises UnsafeCodeError if rejected."""
    try:
        tree = ast.parse(code_str)
    except SyntaxError as e:
        raise UnsafeCodeError(f"Syntax error: {e}")
    SafeASTVisitor().visit(tree)
    return tree


def compile_code(tree: ast.Module, world: World, other_builtins: Dict[str, callable]) -> dict:
    """Compile a validated AST and return an execution namespace."""
    builtins = SAFE_BUILTINS | world.export_as_builtins() | other_builtins
    exec_globals = {"__builtins__": builtins}
    exec_globals.update(SAFE_MODULES)
    exec_locals = {}
    try:
        exec(compile(tree, filename="<pocket_dimension>", mode="exec"), exec_globals, exec_locals)
    except Exception as e:
        raise UnsafeCodeError(f"Error during first the pass of code: {e}")
    exec_globals.update(exec_locals)
    return exec_globals