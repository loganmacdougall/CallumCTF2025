extends Node

var cursor = load("res://assets/sprites/cursor.png")
var hand = load("res://assets/sprites/hand.png")
var ibeam = load("res://assets/sprites/ibeam.png")

func _ready():
	Input.set_custom_mouse_cursor(cursor)
	Input.set_custom_mouse_cursor(hand, Input.CURSOR_POINTING_HAND)
	Input.set_custom_mouse_cursor(ibeam, Input.CURSOR_IBEAM)
