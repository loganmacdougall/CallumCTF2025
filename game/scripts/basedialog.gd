class_name BaseDialog
extends Resource

signal finished_dialog

func get_messages():
	return []

func run_dialog(textbox: Textbox):
	textbox.show_textbox(true)
	
	var messages = get_messages()
	for msg in messages:
		textbox.run_message(msg)
		await textbox.finished_presenting_dialog
	
	textbox.show_textbox(false)
	finished_dialog.emit()
