extends Node2D

@onready var output_label = %RichTextLabel
@onready var codebox = %CodeEdit
@onready var http_req = %HTTPRequest
@onready var code_check = $CheckCompleted

const headers = ["Content-Type: application/json"]
const url = "http://localhost:8081/submit_code"
const default_code = "char state(char f) {\n\treturn f;\n}"

var sending = false
var body : Dictionary

func _ready() -> void:
	codebox.text = GlobalState.light_code
	
func _on_reset_button_pressed() -> void:
	codebox.text = default_code
	GlobalState.light_code = default_code

func _on_send_button_pressed() -> void:
	if sending: return
	
	sending = true
	body["Code"] = codebox.text
	http_req.request(url, headers, HTTPClient.METHOD_POST, JSON.stringify(body))
	output_label.bbcode_enabled = true;
	output_label.text = "\n[wave amp=50.0 freq=5.0 connected=1]loading...[/wave]"

func _on_http_request_request_completed(result: int, response_code: int, headers: PackedStringArray, body: PackedByteArray) -> void:
	output_label.bbcode_enabled = false;
	sending = false
	
	if result == HTTPRequest.RESULT_TIMEOUT:
		output_label.text = "Error: Server timed out"
	elif result != HTTPRequest.RESULT_SUCCESS:
		output_label.text = "Error: Something went wrong"
	else:
		var json = JSON.parse_string(body.get_string_from_utf8())
		output_label.text = json["Output"]
		code_check.check()


func _on_code_edit_text_changed() -> void:
	GlobalState.light_code = codebox.text
