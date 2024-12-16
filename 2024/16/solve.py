import sys
import random

# python default recursion limit is only 1000. 
sys.setrecursionlimit(10**6)

maze = []
with open("2024/16/sample2.txt") as file:
  for line in file:
    maze.append(list(line.rstrip()))


def cheapest_path(maze, x, y, prevX, prevY, visited, memo):
  mk = memo_key(x, y, prevX, prevY, visited)
  if mk in memo:
    return memo[mk]
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
    cost = min(cost, move_cost + cheapest_path(maze, nx, ny, x, y, visited, memo))
  visited.remove((x, y))
  memo[mk] = cost
  return cost

def is_90_turn(x, y, prevX, prevY, nx, ny):
  return not ((x == prevX and prevX == nx) or (y == prevY and prevY == ny))
  

def find_start(maze):
  for y, line in enumerate(maze):
    for x, c in enumerate(line):
      if c == 'S': return (x, y)
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

def memo_key(x, y, prevX, prevY, visited):
  # l_v = list(visited)
  # l_v.sort()
  # v_str = ''
  # for (vx, vy) in l_v:
  #   v_str += f"({vx},{vy})"
  return f"({x},{y}),({prevX},{prevY})" #-{v_str}"


def part1(maze):
  (sx, sy) = find_start(maze)
  return cheapest_path(maze, sx, sy, sx-1, sy, set(), {})



print(part1(maze))