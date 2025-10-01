import sys, os
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

import random

from world import GameoverException, World

class World1(World):
  starting_sign_message = "Interact with the corresponding lever:"
  
  def __init__(self):
    super().__init__("world1")
    self.init_args = []

  def sign_callable(self, num: int) -> tuple[callable, bool]:
    msgs = [self.starting_sign_message]
    left_true = False

    if num > 10:
      left_true = bool(random.randint(0,1))
      msgs.append(f"{'left' if left_true else 'right'}")
    else:
      left = bool(random.randint(0,1))
      nots = random.randint(0, 3)

      left_true = left and (nots % 2 == 0) or (not left and (nots % 2 == 1))
      msg = " ".join(["not" for _ in range(nots)] + ["left" if left else "right"])
      msgs.append(msg)

    def func():
      return msgs
    
    return func, left_true
  
  def good_lever_callable(self, pos: list[int]) -> callable:
    def func():
      tile = self.get_tile(*pos)
      if tile != 3:
        return
      
      self.correct_counter += 1
      self.set_tile(*pos, 7)

      if self.correct_counter == 15:
        self.set_tile(5, 1, 6)
    return func

  def bad_lever_callable(self, pos: list[int]) -> callable:
    def func():
      tile = self.get_tile(*pos)
      if tile != 3:
        return
      
      self.game_over = True
      self.set_tile(*pos, 7)

      raise GameoverException()
    
    return func
  
  def setup(self):
    pos = [5, 3]

    self.correct_counter = 0

    for i in range(15):
      func, left_true = self.sign_callable(i)
      self.run_on_interact[self.get_tile_index(*pos)] = func

      left_idx = self.get_tile_index(pos[0] - 1, pos[1])
      right_idx = self.get_tile_index(pos[0] + 1, pos[1])

      if left_true:
        self.run_on_interact[left_idx] = self.good_lever_callable([pos[0] - 1, pos[1]])
        self.run_on_interact[right_idx] = self.bad_lever_callable([pos[0] + 1, pos[1]])
      else:
        self.run_on_interact[left_idx] = self.bad_lever_callable([pos[0] - 1, pos[1]])
        self.run_on_interact[right_idx] = self.good_lever_callable([pos[0] + 1, pos[1]])

      pos[1] += 3