extends GenericButton

@export_file("*.tscn") var scene

func _on_pressed() -> void:
	SceneTransition.change_scene(scene)
