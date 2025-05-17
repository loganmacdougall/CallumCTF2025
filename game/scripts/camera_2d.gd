extends Camera2D

@export var max_zoom = 4
@export var min_zoom = 1
@export var scroll_amount = 0.1

var moving = false
var previous_mouse_position = Vector2(0,0)

func _process(delta: float) -> void:
	var current_mouse_position = get_viewport().get_mouse_position()
		
	if moving:
		var movement = previous_mouse_position - current_mouse_position
		position += movement / zoom.x
		
	previous_mouse_position = current_mouse_position
	
func _input(event: InputEvent) -> void:
	if event is InputEventMouseButton and event.is_pressed():
		match event.button_index:
			MOUSE_BUTTON_LEFT:
				moving = true
			MOUSE_BUTTON_WHEEL_UP:
				zoom.x -= scroll_amount * zoom.x
				zoom.x = clamp(zoom.x, min_zoom, max_zoom)
				zoom.y = zoom.x
			MOUSE_BUTTON_WHEEL_DOWN:
				zoom.x += scroll_amount * zoom.x
				zoom.x = clamp(zoom.x, min_zoom, max_zoom)
				zoom.y = zoom.x
				
	print(event is InputEventMouseButton and event.is_released() and event.button_index == MOUSE_BUTTON_LEFT)
	if event is InputEventMouseButton and event.is_released() and event.button_index == MOUSE_BUTTON_LEFT:
		moving = false
