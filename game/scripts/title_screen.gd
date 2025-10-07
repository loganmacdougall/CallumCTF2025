extends Node2D

@onready var http: HTTPRequest = $HTTPRequest
@onready var code_check: HTTPRequest = $CheckCompleted
@onready var text: LineEdit = $CanvasLayer/LineEdit
@onready var code_label: Label = $CanvasLayer/code_label

const send_url = "http://localhost:8080/enter_code"
const headers = ["Content-Type: text/plain"]

var sending_code = false

func _send_button_pressed() -> void:
	sending_code = true
	http.request(send_url, headers, HTTPClient.METHOD_POST, text.text)
	
func _on_send_completed(_result, _response_code, _headers, body):
	if sending_code:
		sending_code = false
		var msg = body.get_string_from_utf8()
		code_label.text = msg
		
		if msg.contains("Accepted"):
			code_check.check()
			
	
