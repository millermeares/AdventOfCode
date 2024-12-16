grid = []
moves = []
with open("2024/15/input.txt") as file:
  for line in file:
    if "#" in line:
      grid.append(list(line.rstrip()))
    elif line.rstrip() != '':
      moves +=  list(line.rstrip())

def calculate_score(grid):
  score = 0
  for y, line in enumerate(grid):
    for x, c in enumerate(line):
      if c != 'O' and c != '[':
        continue # not a box.
      score += (100 * y) + x
  return score


def get_robot_idx(grid):
  for y, line in enumerate(grid):
    for x, c in enumerate(line):
      if c == '@':
        return (x, y)
  raise Exception("couldn't find robot")

def get_dir(move):
  if move == '>': return (1, 0)
  elif move == '<': return (-1, 0)
  elif move == 'v': return (0, 1)
  elif move == '^': return (0, -1)
  print("move: " + move)
  raise Exception("invalid move")

def do_move(grid, rx, ry, move):
  (cx, cy) = get_dir(move)
  (nx, ny) = (cx+rx, cy+ry)
  next = grid[ny][nx]
  if next == '#':
    return (rx, ry) # tried to move into wall, did not move.
  if next == '.':
    # empty space, move into that space.
    grid[ry][rx] = '.'
    grid[ny][nx] = '@'
    return (nx, ny)
  
  # next is a box. need to do box move.
  # figure out if boxes can be pushed
  (dx, dy) = (nx, ny)
  while grid[dy][dx] != '.' and grid[dy][dx] != '#':
    (dx, dy) = (dx+cx, dy+cy)

  can_push = grid[dy][dx] == '.'
  if not can_push:
    return (rx, ry) # tried to push boxes, but there are no gaps before a wall.
  
  grid[dy][dx] = 'O' # move a box into the gap.
  grid[ry][rx] = '.' 
  grid[ny][nx] = '@' # robot moves one square

  return (nx, ny)


def part1(grid, moves):
  (rx, ry) = get_robot_idx(grid)
  for move in moves:
    (rx, ry) = do_move(grid, rx, ry, move)
  return calculate_score(grid)


def widen_map(grid):
  wider = []
  for l in grid:
    wide_l = []
    for c in l:
      if c == '#':
        wide_l += ['#','#']
      elif c == '@':
        wide_l += ['@','.']
      elif c == '.':
        wide_l += ['.','.']
      elif c == 'O':
        wide_l += ['[',']']
    wider.append(wide_l)
  return wider


def do_wide_move(grid, rx, ry, move):
  (cx, cy) = get_dir(move)
  (nx, ny) = (cx+rx, cy+ry)
  next = grid[ny][nx]
  if next == '#':
    return (rx, ry) # tried to move into wall, did not move.
  if next == '.':
    # empty space, move into that space.
    grid[ry][rx] = '.'
    grid[ny][nx] = '@'
    return (nx, ny)
  
  # next is a box, need to do box move.
  # if side to side, it's simpler as effects are contained to this column.
  if cy == 0:
    (dx, dy) = (nx, ny)
    while grid[dy][dx] != '.' and grid[dy][dx] != '#':
      (dx, dy) = (dx+cx, dy+cy)

    can_push = grid[dy][dx] == '.'
    if not can_push:
      return (rx, ry) # tried to push boxes, but there are no gaps before a wall.
    # shift boxes one square
    do_left = cx == -1
    while dy != ny or dx != nx:
      if do_left:
        grid[dy][dx] = '['
      else:
        grid[dy][dx] = ']'
      do_left = not do_left
      (dx, dy) = (dx - cx, dy - cy)  
    grid[ry][rx] = '.' 
    grid[ny][nx] = '@' # robot moves one square
    return (nx, ny)
  
  impacted_boxes = list(set(get_all_impacted_boxes(grid, nx, ny, cy, [])))
  # can only move if *all* boxes are not blocked.
  if not can_move_all_boxes(grid, impacted_boxes, cy):
    return (rx, ry) # boxes cannot be moved, don't move.
  
  move_boxes_vertical(grid, nx, ny, cy)
  grid[ny][nx] = '@'
  grid[ry][rx] = '.'
  return (nx, ny)

def move_boxes_vertical(grid, bx, by, cy):
  if grid[by][bx] == '.':
    # this coordinate already moved. do nothing.
    return

  (nx, ny) = (bx, by+cy)
  if grid[ny][nx] == '#':
    raise Exception("Cannot move box into a wall, this is invalid")
  
  if grid[ny][nx] == '[' or grid[ny][nx] == ']':
    move_boxes_vertical(grid, nx, ny, cy) # move other impacted boxes first.

  # move this one.
  if grid[ny][nx] != '.':
    raise Exception("trying to move box into non-empty space.")
  
  (sx, sy) = second_half_coords(grid, bx, by)


  grid[ny][nx] = grid[by][bx]
  grid[by][bx] = '.'
  
  # move second half of box.
  move_boxes_vertical(grid, sx, sy, cy)
  return grid

def can_move_all_boxes(grid, box_coords, cy):
  for (bx, by) in box_coords:
    (nx, ny) = (bx, by+cy)
    v = grid[ny][nx]
    # coordinates in this direction are *either* another impacted box *or* empty space. if wall, 
    if v not in ['.', '[', ']']:
      return False
  return True

# returns array of coordinates of impacted boxes. I *think* this could return duplicates.
def get_all_impacted_boxes(grid, bx, by, cy, checked):

  if (bx, by) in checked or grid[by][bx] not in ['[', ']']:
    return [] # already checked this one, no need to recurse here.
  checked.append((bx, by))
  impacted = [(bx, by)] # i already know that this coordinate is impacted.
  (sx, sy) = second_half_coords(grid, bx, by) # this is second half of my box. 
  impacted += get_all_impacted_boxes(grid, sx, sy, cy, checked)

  (nx, ny) = (bx, by + cy) # directly above/below
  if grid[ny][nx] == '.' or grid[ny][nx] == '#':
    return impacted # reached wall or empty space, no more boxes add *for this side*.

  # next is a box. get those too.
  impacted += get_all_impacted_boxes(grid, nx, ny, cy, checked)
  return impacted
  

def second_half_coords(grid, bx, by):
  val = grid[by][bx]
  if val == ']':
    return (bx-1, by)
  return (bx+1, by)

def part2(wider_grid, moves):
  (rx, ry) = get_robot_idx(wider_grid)
  for move in moves:
    (rx, ry) = do_wide_move(wider_grid, rx, ry, move)
  return calculate_score(wider_grid)

def print_grid(grid):
  for line in grid:
    print(arr_to_str(line))

def arr_to_str(line):
  l = ''
  for c in line:
    l += c
  return l

wider = widen_map(grid)
print(part1(grid, moves))
print_grid(wider)
print(part2(wider, moves))
print_grid(wider)