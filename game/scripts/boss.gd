extends AnimatedSprite2D

@onready var button = $"./GenericButton"
@onready var textbox = %Textbox

var d_explaing: PreparedDialog = load("res://resources/dialog/001_boss.tres")
var d_solved: PreparedDialog = load("res://resources/dialog/002_boss.tres")

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
			

func _ready() -> void:
	if GlobalState.challenges_completed[0]:
		mad = false
		

func _on_button_pressed() -> void:
	talking = true
	
	var d = d_explaing if mad else d_solved
	d.run_dialog(textbox)
	await d.finished_dialog
	
	talking = false

func _on_check_completed_code_check_finished() -> void:
	if GlobalState.challenges_completed[0]:
		mad = false
