class_name PreparedDialog
extends BaseDialog

@export var string_indices: PackedInt32Array

func get_messages():
	var messages = Array()
	
	for index in string_indices:
		var message = DialogManager.get_dialog(index)
		messages.append(message)
		
	return messages
