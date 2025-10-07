extends AnimatedSprite2D

@onready var dialog1: PreparedDialog = load("res://resources/dialog/001_troll.tres")
@onready var dialog2: PreparedDialog = load("res://resources/dialog/002_troll.tres")
@onready var textbox: Textbox = $"../Textbox"

func _ready():
	play("Standing")
	
func _on_generic_button_pressed() -> void:
	play("Talking")
	if GlobalState.challenges_completed[2]:
		dialog2.run_dialog(textbox)
		await dialog2.finished_dialog
	else:
		dialog1.run_dialog(textbox)
		await dialog1.finished_dialog
	play("Standing")
	


func _on_check_completed_request_completed(_result: int, _response_code: int, _headers: PackedStringArray, _body: PackedByteArray) -> void:
	pass
