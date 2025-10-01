extends Node

var conn : StreamPeerTCP
var initialized := false

@onready var runner = $".."

func _ready():
	conn = StreamPeerTCP.new()
	var err = conn.connect_to_host("127.0.0.1", 8083)
	if err != OK:
		push_error("connect_to_host returned error: %s" % err)

func _process(_delta: float) -> void:
	conn.poll()
	
	if not connected():
		return
		
	# wait until the TCP socket reaches connected status
	if not initialized and conn.get_status() == StreamPeerTCP.STATUS_CONNECTED:
		var b64_code = Marshalls.utf8_to_base64(GlobalState.castle_code)
		send('{"code": "%s"}' % b64_code)
		
		initialized = true

	# receive packets if any (only after initialized)
	if initialized and conn.get_status() == StreamPeerTCP.STATUS_CONNECTED:
		if bytes_available():
			var msg = recv()
			runner.process_packet(msg)
			
func connected() -> bool:
	return conn.get_status() == StreamPeerTCP.STATUS_CONNECTED

func bytes_available() -> bool:
	return conn.get_available_bytes() > 0

func send(msg: String) -> void:
	if not connected():
		print("Not initialized yet; dropping send:", msg)
		return
	conn.put_string(msg)

func recv() -> String:
	if conn.get_available_bytes() == 0:
		return ""
		
	var length = conn.get_32()
	return conn.get_utf8_string(length)

func _exit_tree() -> void:
	if conn:
		conn.disconnect_from_host()
