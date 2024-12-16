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

def part2(grid, moves):
  return calculate_score(grid)

print(part1(grid, moves))