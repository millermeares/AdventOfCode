import sys
import random

# python default recursion limit is only 1000. 


maze = []
with open("2024/16/input.txt") as file:
  for line in file:
    maze.append(list(line.rstrip()))

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
        l += 'O'
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

def get_members_of_pathways(x, y, prevX, prevY, costs):
  members = [(x, y)]
  if costs[y][x]['horizontal'] == 0:
    return members # we are at the end!
  neighbors = get_neighbors(maze, x, y)
  min_cost = sys.maxsize
  # figure out which neighbor(s) can be reached through the min cost.
  # bug - i'm preferring to turn for some reason. 
  for (nx, ny) in neighbors:
    is_horizontal = ny == y
    cost = costs[ny][nx]['horizontal'] if is_horizontal else costs[ny][nx]['vertical']
    if is_90_turn(x, y, prevX, prevY, nx, ny):
      cost += 1000 # going to this square would be changing direction.
    if min_cost > cost:
      min_cost = cost
  for (nx, ny) in neighbors:
    is_horizontal = ny == y
    cost = costs[ny][nx]['horizontal'] if is_horizontal else costs[ny][nx]['vertical']
    if is_90_turn(x, y, prevX, prevY, nx, ny):
      cost += 1000 # going to this square would be changing direction.
    if cost == min_cost:
      members += get_members_of_pathways(nx, ny, x, y, costs)
  return members



costs = get_cost_grid(maze)
(sx, sy) = find_location(maze, 'S')
fill(maze, costs)
print(f"Part 1: {costs[sy][sx]['horizontal']}")
path_members = get_members_of_pathways(sx, sy, sx-1, sy, costs)
print(f"Part 2: {len(list(set(path_members)))}")
