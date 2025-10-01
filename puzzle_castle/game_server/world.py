import struct

class GameoverException(Exception):
  pass

class WorldCompleteException(Exception):
  pass


class World:
  WORLD_PATH = "puzzle_castle/world_data/{}.dat"

  def __init__(self, filename: str):
    self.init_args = [filename]

    self.world_view: int = 9
    with open(self.WORLD_PATH.format((filename)), 'rb') as f:
      self.game_over = False
      self.width = struct.unpack('<I', f.read(4))[0]
      self.height = struct.unpack('<I', f.read(4))[0]
      self.game_over = bool(struct.unpack('<b', f.read(1))[0])
      self.tiles = f.read(self.width * self.height)
      if len(self.tiles) != self.width * self.height:
        raise ValueError("Invalid world file: insufficient tile data")
      
      spawn_index = self.tiles.index(16)
      self.spawn_x = spawn_index % self.width
      self.spawn_y = spawn_index // self.width

      self.player_y = self.spawn_y
      self.player_x = self.spawn_x

      self.run_on_tick = []
      self.run_on_interact = {}
      self.run_on_stand = {}

      self.tick = 0
      self.action_tick_count = 0
      self.action_tick_timeout = -1

      self.current_tiles = [t for t in self.tiles]

      self.setup()

  def recreate(self):
    return self.__class__(*self.init_args)
  
  def reset_action_ticks(self):
    self.action_tick_count = 0

  def export_view(self) -> bytes:
    view_size = self.world_view
    half_view = view_size // 2
    offset = 9
    view_data = bytearray(offset + view_size * view_size)
    
    struct.pack_into('<I', view_data, 0, self.world_view)
    struct.pack_into('<I', view_data, 4, self.world_view)
    struct.pack_into('<b', view_data, 8, int(self.game_over))

    sy, sx = self.player_y - half_view, self.player_x - half_view
    for dy in range(view_size):
      ty = sy + dy
      for dx in range(view_size):
        tx = sx + dx
        view_data[offset] = self.get_tile(tx, ty)
        offset += 1

    return bytes(view_data)

  def reset(self):
    self.current_tiles = [t for t in self.tiles]

    self.player_x = self.spawn_x
    self.player_y = self.spawn_y
    self.game_over = False

    self.run_on_tick = []
    self.run_on_interact = {}
    self.run_on_stand = {}

    self.tick = 0
    self.action_tick_count = 0

    self.setup()

  def setup(self):
    pass

  def get_player_position(self) -> tuple[int, int]:
    return (self.player_x, self.player_y)

  def get_tile(self, x: int, y: int) -> int:
    if x < 0 or x >= self.width or y < 0 or y >= self.height:
      return 0
    return self.current_tiles[y * self.width + x]
  
  def set_tile(self, x: int, y: int, tile: int) -> None:
    if x < 0 or x >= self.width or y < 0 or y >= self.height:
      return
    
    if tile < 0 or tile > 16:
      return
    
    self.current_tiles[y * self.width + x] = tile
  
  def get_player_position_index(self) -> int:
    return self.player_y * self.width + self.player_x
  
  def get_tile_index(self, x: int, y: int) -> int:
    return y * self.width + x
  
  def get_tile_up(self) -> int:
    return self.get_tile(self.player_x, self.player_y - 1)
  
  def get_tile_down(self) -> int:
    return self.get_tile(self.player_x, self.player_y + 1)

  def get_tile_left(self) -> int:
    return self.get_tile(self.player_x - 1, self.player_y)
  
  def get_tile_right(self) -> int:
    return self.get_tile(self.player_x + 1, self.player_y)
  
  def get_player_tile(self) -> int:
    return self.get_tile(self.player_x, self.player_y)

  def is_complete(self) -> bool:
    tile = self.get_player_tile()
    return tile == 15

  def tile_is_walkable(self, x: int, y: int) -> bool:
    tile = self.get_tile(x, y)
    return tile not in (0, 4, 5)  # EMPTY, WALL, DOOR LOCKED
  
  def tile_is_safe(self, x: int, y: int) -> bool:
    tile = self.get_tile(x, y)
    return tile != 8  # Spike
  
  def up(self) -> bool:
    res = False
    if self.tile_is_walkable(self.player_x, self.player_y - 1):
      self.player_y -= 1
      res = True
    
    self.world_tick()
    return res
  
  def down(self) -> bool:
    res = False
    if self.tile_is_walkable(self.player_x, self.player_y + 1):
      self.player_y += 1
      res = True
    
    self.world_tick()
    return res
  
  def left(self) -> bool:
    res = False
    if self.tile_is_walkable(self.player_x - 1, self.player_y):
      self.player_x -= 1
      res = True
    
    self.world_tick()
    return res
  
  def right(self) -> bool: 
    res = False
    if self.tile_is_walkable(self.player_x + 1, self.player_y):
      self.player_x += 1
      res = True

    self.world_tick()
    return res

  def world_tick(self) -> None:
    self.tick += 1
    
    player_tile_idx = self.get_player_position_index()

    if player_tile_idx in self.run_on_stand:
      self.run_on_stand[player_tile_idx]()
    
    for func in self.run_on_tick:
      func()

    if not self.tile_is_safe(self.player_x, self.player_y):
      self.game_over = True
      raise GameoverException()
    
    if self.is_complete():
      raise WorldCompleteException()
    
    if self.hit_timeout():
      raise TimeoutError(f"Action ended after running {self.action_tick_timeout} ticks")
    
  def hit_timeout(self) -> bool:
    return self.action_tick_timeout > 0 and self.action_tick_count >= self.action_tick_timeout

  def interact(self) -> None:
    res = None
    idx = self.get_player_position_index()
    if idx in self.run_on_interact:
      res = self.run_on_interact[idx]()

    self.world_tick()
    return res

  def export_as_builtins(self) -> dict:
    return {
      "up": self.up,
      "down": self.down,
      "left": self.left,
      "right": self.right,
      "tile_on": self.get_player_tile,
      "tile_up": self.get_tile_up,
      "tile_down": self.get_tile_down,
      "tile_left": self.get_tile_left,
      "tile_right": self.get_tile_right,
      "interact": self.interact,
      "get_position": self.get_player_position,
    }
      

if __name__ == "__main__":
  world = World("world1")
  print(f"World loaded: {world.width}x{world.height} tiles")

  for y in range(world.height):
    row = []
    for x in range(world.width):
      tile = world.tiles[y * world.width + x]
      row.append(f"{tile:02}")
    print(" ".join(row))

  print(world.export_view())
    