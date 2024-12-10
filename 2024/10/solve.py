grid = []
with open("2024/10/sample.txt") as file:
  for line in file:
    l = []
    for c in line.rstrip():
      l.append(int(c))
    grid.append(l)

def count_trails(grid, i, j):
  cur = grid[i][j]
  if cur == 9:
    return 1
  
  count = 0
  changes = [(0, 1), (0, -1), (1, 0), (-1, 0)]
  for (ci, cj) in changes:
    new_i = i + ci
    new_j = j + cj
    if new_i < 0 or new_j < 0 or new_i >= len(grid) or new_j >= len(grid[i]):
      continue # this would take off of grid
    if grid[new_i][new_j] == cur + 1:
      count += count_trails(grid, new_i, new_j)
  return count

def count_peaks(grid, i, j):
  cur = grid[i][j]
  if cur == 9:
    return [(i, j)]
  
  peaks = []
  changes = [(0, 1), (0, -1), (1, 0), (-1, 0)]
  for (ci, cj) in changes:
    new_i = i + ci
    new_j = j + cj
    if new_i < 0 or new_j < 0 or new_i >= len(grid) or new_j >= len(grid[i]):
      continue # this would take off of grid
    if grid[new_i][new_j] == cur + 1:
      peaks += count_peaks(grid, new_i, new_j)
  return peaks


def part1(grid):
  total = 0
  for i in range(0, len(grid)):
    for j in range(0, len(grid[i])):
      if grid[i][j] == 0:
        total += len(set(count_peaks(grid, i, j)))
  return total

def part2(grid):
  total = 0
  for i in range(0, len(grid)):
    for j in range(0, len(grid[i])):
      if grid[i][j] == 0:
        total += count_trails(grid, i, j)
  return total

print(part1(grid))
print(part2(grid))