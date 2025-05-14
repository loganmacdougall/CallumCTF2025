class_name GenericButton
extends Button

var _invisable
@export var invisable : bool = false : 
	get() : 
		return _invisable
	set(val) : 
		if _invisable == val: return
		_invisable = val
		visibility_layer = 0 if val else 1
		light_mask = 0 if val else 1

func _ready():
	mouse_default_cursor_shape = CURSOR_POINTING_HAND
