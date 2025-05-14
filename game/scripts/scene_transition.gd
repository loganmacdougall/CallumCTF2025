extends CanvasLayer

func change_scene(path: String) -> void:
	$AnimationPlayer.play("dissolve")
	await $AnimationPlayer.animation_finished
	get_tree().change_scene_to_file(path)
	$AnimationPlayer.play_backwards("dissolve")
