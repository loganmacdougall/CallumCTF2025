extends CodeEdit

func _ready() -> void:
	text = GlobalState.castle_code
	
func _on_text_changed() -> void:
	GlobalState.castle_code = text
