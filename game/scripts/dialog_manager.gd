extends Node

var dialog : PackedStringArray
var filepath = "res://resources/dialog.txt"

func _ready():
	var dialog_list = []
	
	var f = FileAccess.open(filepath, FileAccess.READ)
	
	while not f.eof_reached(): # iterate through all lines until the end of file is reached
		dialog_list.append(f.get_line())
		
	f.close()
	
	dialog = PackedStringArray(dialog_list)
	
func get_dialog(i):
	return dialog[i-1]
