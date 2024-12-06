import copy

maze_input = []

with open("input.txt") as file:
  for line in file:
    maze_input.append(list(line.rstrip()))


guard_chars = ['v','<','>','^']

# returns whether or not we are in a cycle?
# visited is an object { i,j: { }}
# returns step of guard
def take_guard_step(maze, guard_row, guard_column):
  i = guard_row
  j = guard_column
  
  guard = maze[i][j]
  next_row, next_column = get_next_step_index(guard, i, j)
  maze[i][j] = 'X' # guard is walking away from this one.
  if not index_in_maze(maze, next_row, next_column):
    return next_row, next_column # exited the maze!
  if maze[next_row][next_column] == '#':
    maze[i][j] = get_right_turn_guard_char(guard)
    return i, j
  maze[next_row][next_column] = guard
  return next_row, next_column
        
def get_next_step_index(guard, guard_row, guard_column):
  if guard == '^':
    return guard_row-1, guard_column
  elif guard =='<':
    return guard_row, guard_column-1
  elif guard == '>':
    return guard_row, guard_column+1
  else: # guard = 'v'
    return guard_row+1, guard_column

def index_in_maze(maze, guard_row, guard_column):
  return guard_row >= 0 and guard_row < len(maze) and guard_column >= 0 and guard_column < len(maze[guard_row])

def is_guard_in_maze(maze):
  return any(list(map(lambda l: any(c in guard_chars for c in l), maze)))


def get_guard_location(maze):
  for i in range(0, len(maze)):
    for j in range(0, len(maze[i])):
      if maze[i][j] in guard_chars:
        return i, j
  return -1, -1

def part1(maze):
  guard_row, guard_column = get_guard_location(maze)
  while index_in_maze(maze, guard_row, guard_column):
    guard_row, guard_column = take_guard_step(maze, guard_row, guard_column)
  # count 'X'
  total = 0
  for l in maze:
    total += l.count('X')
  return total

def get_right_turn_guard_char(c):
  if c == '^':
    return '>'
  elif c == '<':
    return '^'
  elif c == '>':
    return 'v'
  else: # guard = 'v'
    return '<'

print(part1(maze_input))