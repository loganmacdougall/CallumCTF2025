extends HTTPRequest

const get_url = "http://localhost:8080/code_check"

@export var check_on_ready = true

signal code_check_finished

func _ready():
	if check_on_ready:
		check()

func check() -> void:
	self.request(get_url, [], HTTPClient.METHOD_GET)
	
func _on_send_completed(_result, _response_code, _headers, body):
	var msg = body.get_string_from_utf8()
	GlobalState.update_completion(msg)
	code_check_finished.emit()
	
