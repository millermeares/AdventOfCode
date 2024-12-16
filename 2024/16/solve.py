import sys
import random

# python default recursion limit is only 1000. 
sys.setrecursionlimit(10**6)

maze = []
with open("2024/16/input.txt") as file:
  for line in file:
    maze.append(list(line.rstrip()))

# this works but is too expensive.
def cheapest_path(maze, x, y, prevX, prevY, visited):
  print(f"Path from ({x}, {y})")
  if maze[y][x] == 'E':
    return 0 # we're here! 
  visited.add((x, y))
  changes = [(0, 1), (0, -1), (-1, 0), (1, 0)]
  random.shuffle(changes)
  cost = sys.maxsize
  for (cx, cy) in changes:
    (nx, ny) = (x + cx, y + cy)
    if (nx, ny) in visited or maze[ny][nx] == '#':
      continue # already visited
    # no need to account for backwards moves because 'visited' does that.
    move_cost = 1001 if is_90_turn(x, y, prevX, prevY, nx, ny) else 1
    cost = min(cost, move_cost + cheapest_path(maze, nx, ny, x, y, visited))
  visited.remove((x, y))
  return cost

def is_90_turn(x, y, prevX, prevY, nx, ny):
  return not ((x == prevX and prevX == nx) or (y == prevY and prevY == ny))
  

def find_location(maze, location):
  for y, line in enumerate(maze):
    for x, c in enumerate(line):
      if c == location: return (x, y)
  raise Exception("could not find start")

def print_maze(maze, visited):
  for y, line in enumerate(maze):
    l = ''
    for x, c in enumerate(line):
      if (x, y) in visited:
        l += '&'
      else:
        l += c
    print(l)

def get_cost_grid(maze):
  grid = []
  for line in maze:
    gl = []
    for c in line:
      # it's *possible* that i need to add 4 directions rather than 2.
      gl.append({
        'vertical': sys.maxsize,
        'horizontal': sys.maxsize
      })
    grid.append(gl)
  return grid


def fill(maze, costs):
  queue = []
  (ex, ey) = find_location(maze, 'E')
  costs[ey][ex]['vertical'] = 0
  costs[ey][ex]['horizontal'] = 0
  queue.append((ex, ey))
  while len(queue) > 0:
    (x, y) = queue.pop(0)
    print(x, y)
    neighbors = get_neighbors(maze, x, y)
    for (nx, ny) in neighbors:
      should_update = False
      is_horizontal = ny == y
      horiz_add = 1 if is_horizontal else 1001
      vert_add = 1001 if is_horizontal else 1
      cost_to_neighbor = costs[y][x]['horizontal'] if is_horizontal else costs[y][x]['vertical']
      horiz_cost = cost_to_neighbor + horiz_add
      if costs[ny][nx]['horizontal'] > horiz_cost:
        costs[ny][nx]['horizontal'] = horiz_cost
        should_update = True
      vert_cost = cost_to_neighbor + vert_add
      if costs[ny][nx]['vertical'] > vert_cost:
        costs[ny][nx]['vertical'] = vert_cost
        should_update = True
      if should_update:
        queue.append((nx, ny))

  

def get_neighbors(maze, x, y):
  neighbors = []
  changes = [(0, 1), (0, -1), (1, 0), (-1, 0)]
  for (cx, cy) in changes:
    (nx, ny) = (x+cx,y+cy)
    if maze[ny][nx] != '#':
      neighbors.append((nx, ny))
  return neighbors



def part1(maze):
  costs = get_cost_grid(maze)
  # can i do with dp?
  (sx, sy) = find_location(maze, 'S')
  fill(maze, costs)
  return costs[sy][sx]
  # return cheapest_path(maze, sx, sy, sx-1, sy, set())



print(part1(maze))