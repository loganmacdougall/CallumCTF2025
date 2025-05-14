class_name Textbox
extends CanvasLayer

@onready var text = %RichTextLabel
@onready var arrow = %TextboxArrow
@onready var arrow_animation = %ArrowAnimation
var update_time = 0
var fade_finish_time = 0
var current_dialog = ""

signal finished_presenting_dialog

func _ready() -> void:
	fade_finish_time = 1 + text.get_total_character_count() / FlyRichTextEffect.speed

func show_textbox(val):
	visible = val
	
func run_dialog(i):
	print("Loading ",i)
	current_dialog = DialogManager.get_dialog(i)
	text.text = "[fly]" + current_dialog + "[/fly]"
	update_time = 0
	fade_finish_time = 1 + text.get_total_character_count() / FlyRichTextEffect.speed

func _process(delta: float) -> void:
	update_time += delta
	if not arrow_animation.is_playing() && update_time > fade_finish_time:
		arrow.visible = true
		arrow_animation.play("blink")
	elif arrow_animation.is_playing() && update_time <= fade_finish_time:
		arrow.visible = false
		arrow_animation.stop()

func _on_generic_button_pressed() -> void:
	if update_time <= fade_finish_time:
		text.text = current_dialog
		fade_finish_time = 0
		return
		
	finished_presenting_dialog.emit()
