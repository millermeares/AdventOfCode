import sys

bs = []
sample = False
file = 'input'
size = 70
bs_to_consider = 1024
if sample:
  file = 'sample'
  size = 6
  bs_to_consider = 12
with open(f"2024/18/{file}.txt") as file:
  for line in file:
    s = line.rstrip().split(",")
    bs.append((int(s[0]), int(s[1])))


def get_grid(v):
  grid = []
  for i in range(0, size+1):
    line = []
    for j in range(0, size+1):
      line.append(v)
    grid.append(line)
  return grid

def corrupt_spaces(grid, bs, amount):
  for (bx, by) in bs[:amount]:
    grid[by][bx] = '#'

def get_neighbors(grid, x, y):
  neighbors = []
  changes = [(0, 1), (0, -1), (1, 0), (-1, 0)]
  for (cx, cy) in changes:
    (nx, ny) = (x+cx,y+cy)
    if nx < 0 or ny < 0 or ny == len(grid) or nx == len(grid[ny]):
      continue # out of grid
    if grid[ny][nx] != '#':
      neighbors.append((nx, ny))
  return neighbors

def quickest_path(grid):
  costs = get_grid(sys.maxsize)
  costs[size][size] = 0
  queue = []
  queue.append((size, size))
  while len(queue) > 0:
    (x, y) = queue.pop(0)
    p_cost = costs[y][x]
    neighbors = get_neighbors(grid, x, y)
    for (nx, ny) in neighbors:
      neighbor_cost = costs[ny][nx]
      if p_cost+1 < neighbor_cost:
        # neighbor cost is cheaper through this node. 
        costs[ny][nx] = p_cost+1
        queue.append((nx, ny))
  # print_grid(grid)
  # print_grid(costs, True)
  return costs[0][0]


def print_grid(grid, do_comma = False):
  for line in grid:
    s = map(str, list(line))
    if do_comma:
      s = ','.join(s)
    else:
      s = ''.join(s)
    print(s)


def part1(bs):
  grid = get_grid('.')
  corrupt_spaces(grid, bs, bs_to_consider)
  return quickest_path(grid)

def part2(bs):
  p = 0
  consider = 0
  while p != sys.maxsize:
    consider += 1
    print(f"Trying {consider}")
    grid = get_grid('.')
    corrupt_spaces(grid, bs, consider)
    p = quickest_path(grid)
  print(consider)
  return bs[consider-1] 
    


print(part1(bs))
print(part2(bs))

