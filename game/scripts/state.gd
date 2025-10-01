extends Node

@export_storage var lights_fixed = false
@export_storage var light_code = "char state(char f) {\n\treturn f;\n}"

@export_storage var castle_code = "def action_up():\n  up()\n\ndef action_down():\n  down()\n\ndef action_left():\n  left()\n\ndef action_right():\n  right()\n\ndef action_interact():\n  text = interact()\n  if text:\n    display(text)"

signal action_signal

func do_action(action: String):
	action_signal.emit(action)
