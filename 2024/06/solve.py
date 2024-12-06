maze_input = []

with open("sample.txt") as file:
  for line in file:
    maze_input.append(list(line))


guard_chars = ['v','<','>','^']

def take_guard_step(maze):
  for i in range(0, len(maze)):
    for j in range(0, len(maze[i])):
      if (maze[i][j] not in guard_chars):
        continue # not guard
      # ok we are guard.
      guard = maze[i][j]
      maze[i][j] = 'X' # guard is walking away from this one.
      if guard == '^':
        take_step_up(maze, i, j)
      elif guard == '<':
        take_step_left(maze, i, j)
      elif guard == '>':
        take_step_right(maze, i, j)
      else: # guard = 'v'
        take_step_down(maze, i, j)
        

def take_step_up(maze, guard_row, guard_column):
  if guard_row == 0:
    return # guard at top of maze, steps out.
  next_step = maze[guard_row-1][guard_column]
  if next_step == '#': # can't step up, turn 90 degrees.
    maze[guard_row][guard_column] = '>'
    take_guard_step(maze)
    return
  maze[guard_row-1][guard_column] = '^'

def take_step_down(maze, guard_row, guard_column):
  if guard_row+1 == len(maze): # guard at bottom of maze, steps out
    return
  next_step = maze[guard_row+1][guard_column]
  if next_step == '#':
    maze[guard_row][guard_column] = '<'
    take_guard_step(maze)
    return
  maze[guard_row+1][guard_column] = 'v'

def take_step_right(maze, guard_row, guard_column):
  if guard_column+1 == len(maze[guard_row]): # guard at right of maze, steps out.
    return
  next_step = maze[guard_row][guard_column+1]
  if next_step == '#':
    maze[guard_row][guard_column] = 'v'
    take_guard_step(maze)
    return
  maze[guard_row][guard_column+1] = '>'

def take_step_left(maze, guard_row, guard_column):
  if guard_column == 0: # guard at left of maze, steps out
    return
  next_step = maze[guard_row][guard_column-1]
  if next_step == '#':
    maze[guard_row][guard_column] = '^'
    take_guard_step(maze)
    return
  maze[guard_row][guard_column-1] = '<'

def is_guard_in_maze(maze):
  return any(list(map(lambda l: any(c in guard_chars for c in l), maze)))


def replace(str, c, idx):
  return str[:idx] + c + str[idx+1:]

def part1(maze):
  while is_guard_in_maze(maze):
    take_guard_step(maze)
  # count 'X'
  total = 0
  for l in maze:
    total += l.count('X')
  return total


print(part1(maze_input))
