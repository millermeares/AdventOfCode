def read_input(file_name):
  bots = []
  with open(f"2024/14/{file_name}.txt") as file:
    for line in file:
      spl = line.rstrip().split(" ")
      pos_s = spl[0].split('=')[1].split(",")
      pos = (int(pos_s[0]), int(pos_s[1]))
      v_s = spl[1].split('=')[1].split(",")
      v = (int(v_s[0]), int(v_s[1]))
      bots.append([pos, v])
    return bots

def position_after_n(bot, w, h, n):
  (posX, posY) = bot[0]
  (vx, vy) = bot[1]
  # wrap around
  new_x = (posX + (vx * n)) % w
  new_y = (posY + (vy * n)) % h
  return (new_x, new_y)


# given position and width + height of space, return quadrant.
def get_quadrant(pos, w, h):
  (x, y) = pos
  mid_w = w // 2
  mid_h = h // 2
  if x == mid_w or y == mid_h:
    return -1 # in middle, no quadrant matched
  elif x > mid_w and y > mid_h:
    return 1
  elif x < mid_w and y < mid_h:
    return 2
  elif x > mid_w and y < mid_h:
    return 3
  else:
    return 4


def get_grid(positions, w, h, print_middle_lines = True):
  mid_w = w // 2
  mid_h = h // 2
  total = []
  for y in range(0, h):
    l = ""
    for x in range(0, w):
      if not print_middle_lines:
        if y == mid_h or x == mid_w:
          l += ' '
        continue
      if (x, y) not in positions:
        l += "."
        continue
      else:
        l += str(positions[(x, y)])
    total.append(l)
  return total

def print_grid(grid):
  for line in grid:
    print(line)

def print_positions(positions, w, h, print_middle_lines = True):
  grid = get_grid(positions, w, h, print_middle_lines)
  print_grid(grid)

def get_positions(bots, w, h, n):
  positions = {}
  for bot in bots:
    pos = position_after_n(bot, w, h, n)
    if pos not in positions:
      positions[pos] = 0
    positions[pos] += 1
  return positions

def part1(bots, w, h):
  positions = get_positions(bots, w, h, 100)
  quads = {1: 0, 2: 0, 3: 0, 4: 0}
  for pos in positions.keys():
    quad = get_quadrant(pos, w, h)
    if quad == -1:
      continue # who cares
    quads[quad] += positions[pos]
  answer = 1
  for quad in quads.keys():
    answer *= quads[quad]
  return answer

sample_bots = read_input('sample')
print(part1(sample_bots, 11, 7))

bots = read_input('input')
print(part1(bots, 101, 103))

def is_grid_like_christmas_tree(grid, l):
  has_line = has_line_of_length(grid, 20)
  return has_line

def get_v(c, l):
  v = ''
  for i in range(0, l):
    v += c
  return v

def has_line_of_length(grid, l):
  v = get_v('1', l)
  for line in grid:
    if v in line:
      return True
  return False

def is_middle_populated(grid, top_bottom_allowance):
  for i in range(top_bottom_allowance, len(grid[0])-top_bottom_allowance):
    mid_w = len(grid[i]) // 2
    if grid[i][mid_w] == '.':
      return False
  return True

def print_iterations(bots, w, h, n):
  for i in range(0, n):
    positions = get_positions(bots, w, h, i)
    grid = get_grid(positions, w, h)
    if is_grid_like_christmas_tree(grid, 40):
      print(i)
      print_grid(grid)
      print()
    else: 
      if i % 10000 == 0:
        print(f"{i} no")

# m is max
def find_christmas_tree(bots, w, h, m):
  for i in range(0, m):
    positions = get_positions(bots, w, h, i)
    grid = get_grid(positions, w, h)
    if is_grid_like_christmas_tree(grid, 40):
      print(i)
      print_grid(grid)
      print()
      return i
    else: 
      if i % 10000 == 0:
        print(f"{i} no")

print(find_christmas_tree(bots, 101, 103, 10000))