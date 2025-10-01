import sys, os
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

from world import World
import random

class World3(World):
  sign_message = [
    "One combination of levers will open the door.",
    "It's up to you to figure out which one it is.",
  ]

  correct_count = 0
  
  def __init__(self):
    super().__init__("world3")
    self.init_args = []

  def sign_message_func(self):
    return self.sign_message
  
  def lever_on_callable(self, pos):
    def func():
      tile = self.get_tile(*pos)
      flipped = 3 if tile == 7 else 7
      self.set_tile(*pos, flipped)

      if flipped == 7:
        self.correct_count += 1
      else:
        self.correct_count -= 1
      
      exit_tile = 6 if self.correct_count == 9 else 5
      self.set_tile(6, 2, exit_tile)
      
    return func
  
  def lever_off_callable(self, pos):
    def func():
      tile = self.get_tile(*pos)
      flipped = 3 if tile == 7 else 7
      self.set_tile(*pos, flipped)

      if flipped == 3:
        self.correct_count += 1
      else:
        self.correct_count -= 1
      
      exit_tile = 6 if self.correct_count == 9 else 5
      self.set_tile(6, 2, exit_tile)
    
    return func

  def setup(self):
    sign_pos = [6, 5]
    self.run_on_interact[self.get_tile_index(*sign_pos)] = self.sign_message_func

    start_lever_pos = [2, 3]
    for i in range(9):
      lever_pos = [start_lever_pos[0] + i, start_lever_pos[1]]
      lever_idx = self.get_tile_index(*lever_pos)

      should_be_on = bool(random.randint(0,1))

      if not should_be_on:
        self.correct_count += 1

      if should_be_on:
        self.run_on_interact[lever_idx] = self.lever_on_callable(lever_pos)
      else:
        self.run_on_interact[lever_idx] = self.lever_off_callable(lever_pos)