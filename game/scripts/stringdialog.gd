class_name StringDialog
extends BaseDialog

@export var messages: Array = []

func _init(p_messages: Array = []):
	self.messages = p_messages.duplicate()

func get_messages():
	return self.messages
