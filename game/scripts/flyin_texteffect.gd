@tool
class_name FlyRichTextEffect
extends RichTextEffect


const speed = 35
var bbcode = "fly"

func _process_custom_fx(char_fx):
	var time = char_fx.elapsed_time
	var index = char_fx.range.x
	
	var alpha = 1 if speed * time > index else 0 
	char_fx.color.a = alpha
	return true
