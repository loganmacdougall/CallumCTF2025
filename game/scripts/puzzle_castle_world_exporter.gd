@tool
extends Node

@export var worldname: String = "world1"

@export_tool_button("Export Current World")
var export_world_action = export_world_filename
@export_tool_button("Import Current World")
var import_world_action = import_world_filename

const worldpath = "res://resources/world_exports/%s.dat"
@onready var tiles: TileMapLayer = %Ground

func export_world_bytes():
	var world_rect = tiles.get_used_rect()
	var world_pos = world_rect.position
	var world_dims = world_rect.size
	
	var buf = PackedByteArray()
	var offset = 9
	
	buf.resize(offset + (world_dims.x * world_dims.y))
	buf.encode_s32(0, world_dims.x)
	buf.encode_s32(4, world_dims.y)
	buf.encode_s8(8, 0) # Alive
	
	for y in range(world_dims.y):
		var ay = y + world_pos.y
		for x in range(world_dims.x):
			var ax = x + world_pos.x
			
			var tile_cords = tiles.get_cell_atlas_coords(Vector2i(ax, ay))
			if tile_cords.x == -1:
				buf.encode_s8(offset, 0)
			else:
				var tile_idx = tile_cords.y * 4 + tile_cords.x + 1
				buf.encode_s8(offset, tile_idx)
			offset += 1
	
	return buf

func export_world_filename():
	var file = FileAccess.open(worldpath % worldname, FileAccess.WRITE)
	
	var buf = export_world_bytes()
	file.store_buffer(buf)

func import_world_bytes(buf: PackedByteArray):
	var dim_x = buf.decode_s32(0)
	var dim_y = buf.decode_s32(4)
	var gameover = bool(buf.decode_s8(8))
	var offset = 9
	
	tiles.clear()
	
	for y in range(dim_y):
		for x in range(dim_x):
			var idx = buf.decode_s8(offset) - 1
			offset += 1
			
			if idx == -1:
				continue
				
			var idx_x = idx % 4
			var idx_y = idx / 4
			
			tiles.set_cell(Vector2i(x, y), 2, Vector2i(idx_x, idx_y))
	
	return gameover
	
func import_world_filename():
	var file = FileAccess.open(worldpath % worldname, FileAccess.READ)
	file.big_endian = true;
	
	import_world_bytes(file.get_buffer(file.get_length()))
