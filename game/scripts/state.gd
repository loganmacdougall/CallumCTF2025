extends Node

@export_storage var lights_fixed = false
@export_storage var light_code = "char state(char f) {\n\treturn f;\n}"

signal action_signal

func do_action(action: String):
	action_signal.emit(action)
