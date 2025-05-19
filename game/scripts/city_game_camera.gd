extends Camera2D

@onready var viewport: SubViewportContainer = %CitySubViewportContainer

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
		correct_camera_position()
		
	previous_mouse_position = current_mouse_position

func correct_camera_position() -> void:
		var x_camera_halfsize = viewport.size.x / (2 * zoom.x)
		var y_camera_halfsize = viewport.size.y / (2 * zoom.y)
		var left_clamp = limit_left + x_camera_halfsize
		var right_clamp = limit_right - x_camera_halfsize
		var top_clamp = limit_top + y_camera_halfsize
		var bottom_clamp = limit_bottom - y_camera_halfsize
		
		position.x = clamp(position.x, left_clamp, right_clamp)
		position.y = clamp(position.y, top_clamp, bottom_clamp)

func _input(event: InputEvent) -> void:
	var current_mouse_position = get_viewport().get_mouse_position()
	
	if event is InputEventMouseButton and event.is_pressed():
		match event.button_index:
			MOUSE_BUTTON_LEFT:
				moving = true
			MOUSE_BUTTON_WHEEL_UP:
				if zoom.x > min_zoom: 
					zoom.x -= scroll_amount * zoom.x
					zoom.x = clamp(zoom.x, min_zoom, max_zoom)
					zoom.y = zoom.x
			MOUSE_BUTTON_WHEEL_DOWN:
				if zoom.x < max_zoom:
					zoom.x += scroll_amount * zoom.x
					zoom.x = clamp(zoom.x, min_zoom, max_zoom)
					zoom.y = zoom.x
				
	if event is InputEventMouseButton and event.is_released() and event.button_index == MOUSE_BUTTON_LEFT:
		moving = false
