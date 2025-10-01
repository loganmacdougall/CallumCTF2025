extends Area2D


func _on_input_event(viewport: Node, event: InputEvent, shape_idx: int) -> void:
	print("Detected Input!!!!!!!!!!!")
	if event is InputEventMouseButton and event.pressed:
		print("Clicked!!!!!!!!!!!!!!!")
