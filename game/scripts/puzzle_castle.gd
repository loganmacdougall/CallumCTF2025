extends Node2D

@onready var textbox = $Textbox
@onready var timer = $UI/Timer
@onready var console = $UI/TextEdit
@onready var tile_exporter = $PuzzleCastleWorldExporter
@onready var connection = $PuzzleCastleConnection
@onready var player = $World/Player
@onready var center_text = $UI/CenterLabel

var displaying_text = false
var gameover = false
var gameover_handled = false
var has_connected = false
var allow_input = false

func _process(_delta: float) -> void:	
	allow_input = false
	
	if gameover:
		textbox.visible = false
		if not gameover_handled:
			player.play("dead")
			center_text.text = "GAME OVER"
			center_text.visible = true
			gameover_handled = true
		return
			
	if not connection.connected():
		center_text.text = "CONNECTING"
		center_text.visible = true
	else:
		center_text.visible = false
	
	if displaying_text:
		if Input.is_action_just_pressed("Interact"):
			textbox.click()
		return
		
	allow_input = true

func _input(event: InputEvent) -> void:
	if not allow_input:
		return
		
	if event.is_action_pressed("Up"):
		connection.send('{"action":"action_up"}')
	elif event.is_action_pressed("Down"):
		connection.send('{"action":"action_down"}')
	elif event.is_action_pressed("Left"):
		connection.send('{"action":"action_left"}')
	elif event.is_action_pressed("Right"):
		connection.send('{"action":"action_right"}')
	elif event.is_action_pressed("Interact"):
		connection.send('{"action":"action_interact"}')
	elif event is InputEventKey and event.pressed:
		var keycode = event.keycode
		if keycode >= KEY_0 and keycode <= KEY_9:
			connection.send('{"action":"action_%d"}' % (keycode - KEY_0))	

func reset():
	get_tree().reload_current_scene()

func process_packet(msg: String):
	var json_result = JSON.parse_string(msg)
	if json_result == null:
		print("Recieved packet was not valid JSON")
		return
	
	for key in json_result.keys():
		if key == "state":
			update_board(json_result[key])
		elif key == "time":
			update_timer(json_result[key])
		elif key == "console":
			append_console(json_result[key])
		elif key == "text":
			display_text(json_result[key])
		
func update_board(msg: String):
	var bytes = Marshalls.base64_to_raw(msg)
	gameover = tile_exporter.import_world_bytes(bytes)
	
func update_timer(time: float):
	timer.text = "%0.2f" % time
	
func append_console(msgs: Array):
	var v_scroll = console.get_v_scroll_bar()
	var force_bottom = v_scroll.max_value == v_scroll.value
	for msg in msgs:
		console.text += msg + "\n"
		
	var count = len(console.text)
	if count > 3000:
		console.text = console.text.substr(count - 3000)
		
	if force_bottom:
		v_scroll.value = v_scroll.max_value

func display_text(msgs: Array):
	var dialog = StringDialog.new(msgs)
	displaying_text = true
	dialog.run_dialog(textbox)
	await dialog.finished_dialog
	displaying_text = false
	

func _on_reset_button_pressed() -> void:
	reset()
