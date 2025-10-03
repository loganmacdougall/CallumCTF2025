import sys, os
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

from world import World

class World5(World):
  sign_message = [
    "The ground is riddled with spikes ahead.",
    "The area ahead is a 7x7 square of hidden traps.",
    "A spike will be active is determined by this function:\n(hash((x, y, t)) % 10) < 3",
    "x and y are the coordinates of the spike\nTop Left: (0,0), Bottom Right: (6,6)",
    "t is the number of code actions taken since entering the world\n(up, down, left, right, interact)",
    "Make it across safely",
  ]
  
  def __init__(self):
    super().__init__("world5")
    self.init_args = []

  def sign_message_func(self):
    return self.sign_message
  
  def control_spikes(self):
    first_spike = [1,1]

    t = self.tick

    for y in range(7):
      sy = y + first_spike[1]
      for x in range(7):
        sx = x + first_spike[0]

        rise = (hash((x, y, t)) % 10) < 3

        self.set_tile(sx, sy, 8 if rise else 1)

  def setup(self):
    sign_pos = [4, 9]
    self.run_on_interact[self.get_tile_index(*sign_pos)] = self.sign_message_func
    self.run_on_tick.append(self.control_spikes)
    