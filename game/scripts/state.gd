extends Node

@export_storage var light_code = "char state(char f) {\n\treturn f;\n}"
@export_storage var castle_code = "def action_up():\n  up()\n\ndef action_down():\n  down()\n\ndef action_left():\n  left()\n\ndef action_right():\n  right()\n\ndef action_interact():\n  text = interact()\n  if text:\n    display(text)"

@export_storage var lights_fixed = false
@export_storage var challenges_completed = [false, false, false]

signal action_signal

func do_action(action: String):
	action_signal.emit(action)
	
func update_completion(completed):
	for i in range(len(completed)):
		var c = completed[i]
		
		if c == "1":
			challenges_completed[i] = true
		else:
			challenges_completed[i] = false
