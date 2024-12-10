grid = []
with open("2024/10/input.txt") as file:
  for line in file:
    l = []
    for c in line.rstrip():
      l.append(int(c))
    grid.append(l)

# i am currently counting routes
def count_peaks(grid, i, j):
  cur = grid[i][j]
  print(f"Counting trails at {i}, {j} with {cur} val")
  if cur == 9:
    return [(i, j)]
  
  peaks = []
  changes = [(0, 1), (0, -1), (1, 0), (-1, 0)]
  for (ci, cj) in changes:
    new_i = i + ci
    new_j = j + cj
    if new_i < 0 or new_j < 0 or new_i >= len(grid) or new_j >= len(grid[i]):
      print(f"{new_i}, {new_j} would take off of grid")
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

print(part1(grid))