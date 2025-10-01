import sys, os
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

from world import World

class WorldF(World):
  sign_message = [
    "You've surpassed all my challenges!",
    "You must now take this to the realm of the web and enter the following code:",
    "'CASTLE_CRASHER_9457'"
  ]
  
  def __init__(self):
    super().__init__("worldF")
    self.init_args = []

  def sign_message_func(self):
    return self.sign_message
  
  def setup(self):
    sign_pos = [3, 2]
    self.run_on_interact[self.get_tile_index(*sign_pos)] = self.sign_message_func
    