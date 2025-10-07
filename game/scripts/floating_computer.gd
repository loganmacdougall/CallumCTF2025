extends AnimatedSprite2D

@onready var dialog: PreparedDialog = load("res://resources/dialog/001_computer.tres")
@onready var textbox: Textbox = $"../Textbox"


func _on_button_pressed() -> void:
	dialog.run_dialog(textbox)
	await dialog.finished_dialog
