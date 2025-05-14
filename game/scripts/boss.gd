extends AnimatedSprite2D

@onready var button = $"./GenericButton"
@onready var textbox = %Textbox

var d_explaing: Dialog = load("res://resources/dialog/001_boss.tres")

var _mad = true
@export var mad: bool:
	get:
		return _mad
	set(val):
		if val != _mad:
			if val == true:
				self.animation = "mad"
			else:
				self.animation = "happy"
			_mad = val
			
var _talking = false
@export var talking: bool:
	get:
		return _talking
	set(val):
		if val != _talking:
			if val == true:
				self.frame = 1
				self.play()
			else:
				self.pause()
				self.frame = 0
			_talking = val
			

func _on_button_pressed() -> void:
	talking = true
	d_explaing.run_dialog(textbox)
	await d_explaing.finished_dialog
	talking = false
