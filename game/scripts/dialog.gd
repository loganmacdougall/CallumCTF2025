class_name Dialog
extends Resource

@export var string_indices: PackedInt32Array

signal finished_dialog

func run_dialog(textbox: Textbox):
	textbox.show_textbox(true)
	for index in string_indices:
		textbox.run_dialog(index)
		await textbox.finished_presenting_dialog
	textbox.show_textbox(false)
	finished_dialog.emit()
