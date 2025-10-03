extends AnimatedSprite2D

@onready var dialog: PreparedDialog = load("res://resources/dialog/001_troll.tres")
@onready var textbox: Textbox = $"../Textbox"

func _ready():
	play("Standing")
	

func _on_generic_button_pressed() -> void:
	play("Talking")
	dialog.run_dialog(textbox)
	await dialog.finished_dialog
	play("Standing")
	
