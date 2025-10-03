extends Node2D

@onready var http: HTTPRequest = $HTTPRequest
@onready var text: LineEdit = $CanvasLayer/LineEdit
@onready var code_label: Label = $CanvasLayer/code_label

const send_url = "http://localhost:8080/enter_code"
const headers = ["Content-Type: text/plain"]

func _send_button_pressed() -> void:
	http.request(send_url, headers, HTTPClient.METHOD_POST, text.text)
	
func _on_send_completed(result, response_code, headers, body):
	code_label.text = body.get_string_from_utf8()
