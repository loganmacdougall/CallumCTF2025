def action_up():
  up()

def action_down():
  down()

def action_left():
  left()

def action_right():
  right()

def action_interact():
  text = interact()
  if text:
    print(text)
    display(text)
    
def action_1():
  while True:
    for _ in range(3):
      up()
    
    text = interact()[1]
    broken_up = text.split(" ")
    is_left = broken_up[-1] == "left"
    negate = len(broken_up) % 2 == 0
    
    ans_left = is_left ^ negate
    
    if ans_left:
      left()
      interact()
      right()
    else:
      right()
      interact()
      left()

def action_2():
  up()
  up()
  text = interact()[1]
  
  dir_map = {
    "U": up,
    "R": right,
    "D": down,
    "L": left
  }
  
  r_dir_map = {
    "U": right,
    "R": down,
    "D": left,
    "L": up
  }
  
  for dir in text:
    dir_map[dir]()
    
  for dir in text:
    r_dir_map[dir]()
    
def action_3():
  for _ in range(5):
    up()
    
  tracked_x = 5
  
  def move_to(num):
    nonlocal tracked_x
    dist = num - tracked_x
    if dist > 0:
      for _ in range(dist):
        left()
    else:
      for _ in range(abs(dist)):
        right()
    tracked_x = num
  
  def flip(num):
    move_to(num)
    interact()
    return tile_on() == 7
    
  def check_door_open():
    move_to(5)
    return tile_up() == 6
  
  for _ in range(1024):
    i = 1
    while flip(i):
      i += 1
    if check_door_open():
      break
      
  up()
  up()
  up()
    
def action_4():
  for _ in range(13):
    up()
  interact()
    
  LEFT = -1
  RIGHT = 1
  BACK = 0
      
  def check(dir) -> bool:
    if dir == BACK:
      return True
      
    left() if dir == LEFT else right()
    res = tile_left() if dir == LEFT else tile_right()
    right() if dir == LEFT else left()
    return res == 6
    
  def go_back():
    for _ in range(12):
      down()
    can_go_left = tile_left() == 1
    for _ in range(7):
      left() if can_go_left else right()
    return False
      
  def go_room(dir):
    if dir == BACK:
      return go_back()
      
    for _ in range(7):
      left() if dir == LEFT else right()
    for _ in range(12):
      up()
    res = tile_on() == 3
    interact()
    return res
    
  def next_dir(dir):
    if dir == LEFT:
      return RIGHT
    if dir == RIGHT:
      return BACK
    return LEFT
      
  stack = [LEFT]
  
  for _ in range(10000):
    if len(stack) == 1 and stack[0] == BACK:
      stack = [LEFT]
      
    elem = stack.pop()
    
    if elem == BACK:
      go_back()
      stack[-1] = next_dir(stack[-1])
    elif check(elem):
      go_room(elem)
      stack.extend([elem, LEFT])
    else:
      stack.append(next_dir(elem))
      
def action_5():
  for _ in range(4):
    up()
    
  t = 4
  
  def rise(x, y, t):
    return (hash((x, y, t)) % 10) < 3
    
  starting_pos = [3, 6]
  
  for _ in range(1000):
    found = True
    for i in range(7):
      tx = t + i + 1
      xx = starting_pos[0]
      yx = starting_pos[1] - i
      
      if rise(xx, yx, tx):
        found = False
        break
    
    if found:
      break
    else:
      interact()
      t += 1
      
  
  for _ in range(10):
    up()