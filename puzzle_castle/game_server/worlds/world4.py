import sys, os
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

from world import World
import random

class RoomTree():
  def __init__(self, depth = 4, parent = None):
    self.depth = depth
    self.parent = parent
    self.lever = False
    if depth > 0:
      self.left = RoomTree(depth - 1, self)
      self.right = RoomTree(depth - 1, self)
    else:
      self.left = None
      self.right = None
    self.left_unlocked = False
    self.right_unlocked = False

  def root(self) -> "RoomTree":
    if self.parent:
      return self.parent.root()
    return self

  def unlock_door_in_tree(self) -> bool:
    unlock_left = bool(random.randint(0, 1))
    if unlock_left:
      if not self.left_unlocked:
        self.left_unlocked = True
        return True
      elif self.left:
        return self.left.unlock_door_in_tree()
    else:
      if not self.right_unlocked:
        self.right_unlocked = True
        return True
      elif self.right:
        return self.right.unlock_door_in_tree()
    return False


class World4(World):
  sign_message = [
    "All doors begin locked, each lever opening a new unknown door.",
    "At the greatest depth of the maze, the exit lies."
  ]

  sign_pos = [8, 44]

  start_left_up = [[1, 28],[8, 24]]
  start_right_up = [[15, 28],[8, 24]]
  start_left_down = [[8, 25], [1, 29]]
  start_right_down = [[8, 25], [15, 29]]
  main_left_up = [[1, 12], [8, 24]]
  main_right_up = [[15, 12], [8, 24]]
  main_left_down = [[8, 25], [1, 13]]
  main_right_down = [[8, 25], [15, 13]]
  end_left_up = [[1, 12], [8, 8]]
  end_right_up = [[15, 12], [8, 8]]
  end_left_down = [[8, 9], [1, 13]]
  end_right_down = [[8, 9], [15, 13]]

  starter_lever = [8, 34]
  main_lever = [8, 18]
  end_lever = [8, 2]

  starter_left_door = [6, 34]
  starter_right_door = [10, 34]
  main_left_door = [6, 18]
  main_right_door = [10, 18]
  end_left_door = [6, 2]
  end_right_door = [10, 2]
  
  def __init__(self):
    super().__init__("world4")
    self.init_args = []

  def sign_message_func(self):
    return self.sign_message
  
  def teleport_callable(self, pos, dir: int):
    def func():
      self.player_x = pos[0]
      self.player_y = pos[1]

      if dir == 0:
        self.room = self.room.left
      elif dir == 1:
        self.room = self.room.right
      else:
        self.room = self.room.parent

      self.update_current_room()
    return func
  
  def lever_on_callable(self, pos):
    def func():
      tile = self.get_tile(*pos)
      if tile == 7:
        return
      
      self.tree.unlock_door_in_tree()

      self.room.lever = True

      self.update_current_room()

    return func

  def update_current_room(self):
    if not self.room.parent:
      return self.update_starter_room()
    elif not self.room.left:
      return self.update_end_room()
    else:
      return self.update_main_room()
    
  def update_starter_room(self):
    self.set_tile(*self.starter_left_door, 6 if self.room.left_unlocked else 5)
    self.set_tile(*self.starter_right_door, 6 if self.room.right_unlocked else 5)
    self.set_tile(*self.starter_lever, 7 if self.room.lever else 3)

  def update_main_room(self):
    self.set_tile(*self.main_left_door, 6 if self.room.left_unlocked else 5)
    self.set_tile(*self.main_right_door, 6 if self.room.right_unlocked else 5)
    self.set_tile(*self.main_lever, 7 if self.room.lever else 3)
  
    close_to_start = self.room.parent.parent == None
    close_to_end = self.room.depth == 1

    from_left = self.room.parent.left == self.room

    if close_to_end:
      self.run_on_stand[self.get_tile_index(*self.end_left_up[0])] = self.teleport_callable(self.end_left_up[1], 0)
      self.run_on_stand[self.get_tile_index(*self.end_right_up[0])] = self.teleport_callable(self.end_right_up[1], 1)
    else:
      self.run_on_stand[self.get_tile_index(*self.main_left_up[0])] = self.teleport_callable(self.main_left_up[1], 0)
      self.run_on_stand[self.get_tile_index(*self.main_right_up[0])] = self.teleport_callable(self.main_right_up[1], 1)

    if close_to_start:
      if from_left:
        self.run_on_stand[self.get_tile_index(*self.start_left_down[0])] = self.teleport_callable(self.start_left_down[1], 2)
      else:
        self.run_on_stand[self.get_tile_index(*self.start_right_down[0])] = self.teleport_callable(self.start_right_down[1], 2)
    else:
      if from_left:
        self.run_on_stand[self.get_tile_index(*self.main_left_down[0])] = self.teleport_callable(self.main_left_down[1], 2)
      else:
        self.run_on_stand[self.get_tile_index(*self.main_right_down[0])] = self.teleport_callable(self.main_right_down[1], 2)

  def update_end_room(self):
    self.set_tile(*self.end_left_door, 6 if self.room.left_unlocked else 5)
    self.set_tile(*self.end_right_door, 6 if self.room.right_unlocked else 5)
    self.set_tile(*self.end_lever, 7 if self.room.lever else 3)
  
    from_left = self.room.parent.left == self.room

    if from_left:
      self.run_on_stand[self.get_tile_index(*self.end_left_down[0])] = self.teleport_callable(self.end_left_down[1], 2)
    else:
      self.run_on_stand[self.get_tile_index(*self.end_right_down[0])] = self.teleport_callable(self.end_right_down[1], 2)

  def setup(self):
    self.run_on_interact[self.get_tile_index(*self.sign_pos)] = self.sign_message_func

    self.tree = RoomTree(4, None)
    self.room = self.tree

    self.run_on_stand[self.get_tile_index(*self.start_left_up[0])] = self.teleport_callable(self.start_left_up[1], 0)
    self.run_on_stand[self.get_tile_index(*self.start_right_up[0])] = self.teleport_callable(self.start_right_up[1], 1)

    self.run_on_interact[self.get_tile_index(*self.starter_lever)] = self.lever_on_callable(self.starter_lever)
    self.run_on_interact[self.get_tile_index(*self.main_lever)] = self.lever_on_callable(self.main_lever)
    self.run_on_interact[self.get_tile_index(*self.end_lever)] = self.lever_on_callable(self.end_lever)

    self.update_current_room()