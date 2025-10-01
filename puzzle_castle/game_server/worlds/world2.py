import sys, os
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

from world import GameoverException, World

class World2(World):
  sign1_message = "After this sign, follow these directions exactly as they are given:"
  sign2_message = "Now do the instructions from the previous sign rotated 90 degrees clockwise."

  move_list = "UURULLLLLDLLUUURDRUUUURRUUUULLLLDLULLDLDDLLUURUUR"
  
  def __init__(self):
    super().__init__("world2")
    self.init_args = []
  
  def sign1_callable(self):
    return [self.sign1_message, self.move_list]
  
  def sign2_callable(self):
    return [self.sign2_message]
  
  def drop_spike_on_player(self):
    self.set_tile(self.player_x, self.player_y, 8)

  def move_order_callable(self):
    phase = 0
    index = 0

    prev_pos = [self.player_x, self.player_y]

    def func():
      nonlocal phase, index, prev_pos

      if self.player_x == prev_pos[0] and self.player_y == prev_pos[1]:
        return

      if phase == 0:
        if self.player_x == 16 and self.player_y == 28:
          phase = 1
      elif phase == 1:
        dir = self.move_list[index]
        if dir == 'U':
          if prev_pos[1] - 1 != self.player_y:
            self.drop_spike_on_player()
        elif dir == 'D':
          if prev_pos[1] + 1 != self.player_y:
            self.drop_spike_on_player()
        elif dir == 'L':
          if prev_pos[0] - 1 != self.player_x:
            self.drop_spike_on_player()
        elif dir == 'R':
          if prev_pos[0] + 1 != self.player_x:
            self.drop_spike_on_player()
        index += 1

        if index >= len(self.move_list):
          phase = 2
          index = 0
      else:
        dir = self.move_list[index]
        if dir == 'L':
          if prev_pos[1] - 1 != self.player_y:
            self.drop_spike_on_player()
        elif dir == 'R':
          if prev_pos[1] + 1 != self.player_y:
            self.drop_spike_on_player()
        elif dir == 'D':
          if prev_pos[0] - 1 != self.player_x:
            self.drop_spike_on_player()
        elif dir == 'U':
          if prev_pos[0] + 1 != self.player_x:
            self.drop_spike_on_player()
        index += 1
        
      prev_pos = [self.player_x, self.player_y]
    return func

  def setup(self):
    sign1_pos = [16, 28]
    sign2_pos = [6, 15]

    self.run_on_interact[self.get_tile_index(*sign1_pos)] = self.sign1_callable
    self.run_on_interact[self.get_tile_index(*sign2_pos)] = self.sign2_callable
    self.run_on_tick.append(self.move_order_callable())
    
