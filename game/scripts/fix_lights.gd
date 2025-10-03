extends Node2D

@onready var http: HTTPRequest = $HTTPRequest
@onready var boss = $PotatoBoss

const get_url = "http://localhost:8080/code_check?c=1"

func _ready():
	http.request(get_url, [], HTTPClient.METHOD_GET)
	

func _on_http_request_request_completed(result: int, response_code: int, headers: PackedStringArray, body: PackedByteArray) -> void:
	if result != HTTPRequest.RESULT_SUCCESS:
		return
		
	var res = body.get_string_from_utf8()
	if res == "true":
		GlobalState.lights_fixed = true
		boss.mad = false
	
