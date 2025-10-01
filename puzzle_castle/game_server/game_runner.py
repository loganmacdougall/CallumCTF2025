import multiprocessing
import traceback
from typing import List
from world import GameoverException, World, WorldCompleteException
from ast_validator import compile_code

class GameRunner:
  def __init__(self, worlds: List[World], ast: any, tick_timeout = 2000000):
    self.worlds = [w.recreate() for w in worlds]
    self.ast = ast
    self.tick_timeout = tick_timeout
    self.console = []
    self.text = []

    for world in self.worlds:
      world.action_tick_timeout = tick_timeout

    try:
      self.actions = compile_code(ast, self.worlds[0], {
        "print": self.console_print,
        "display": self.text_display,
      })
    except Exception as e:
      raise e
    
    self.current_world = 0

  def console_print(self, msgs: str) -> None:
    if len(self.console) >= 50:
      return
    
    if isinstance(msgs, list):
      for msg in msgs:
        self.console.append(str(msg))  
        if len(self.console) >= 50:
          break
      return
    self.console.append(str(msgs))

  def text_display(self, msgs: str) -> None:
    if len(self.text) >= 50:
      return
    
    if isinstance(msgs, list):
      for msg in msgs:
        self.text.append(str(msg))  
        if len(self.text) >= 50:
          break
      return
    self.text.append(str(msgs))

  def run_action(self, action: str) -> None:
    world = self.worlds[self.current_world]
    world.reset_action_ticks()
    
    if world.game_over:
      return
        
    if action not in {
        "action_up", "action_down", "action_left", "action_right", "action_interact",
        "action_1", "action_2", "action_3", "action_4", "action_5",
        "action_6", "action_7", "action_8", "action_9", "action_0"
      }:
      raise ValueError(f"Invalid action: {action}")
      
    if action not in self.actions:
      raise ValueError(f"No function defined for action: {action}")

    try: 
      self.actions[action]()
    except GameoverException as e:
      world.game_over = True 
      return 
    except WorldCompleteException as e:
      pass

    if world.is_complete():
      self.current_world += 1
      if self.current_world >= len(self.worlds):
        self.game_over = True
        return
      else:
        self.actions = compile_code(self.ast, self.worlds[self.current_world], {
          "print": self.console_print,
          "display": self.text_display,
        })
