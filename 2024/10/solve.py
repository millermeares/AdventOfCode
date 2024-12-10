grid = [list(map(int, line.strip())) for line in open("2024/10/input.txt")] 

trailheads = [
  (j, i) for j, row in enumerate(grid) for i, height in enumerate(row) if height == 0
]
# returns a list of paths reached. duplicates will happen, if they are reached by different paths.
def paths_to_peaks(grid, i, j):
  cur = grid[i][j]
  if cur == 9:
    return [(i, j)]
  
  peaks = []
  for (ci, cj) in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
    new_i = i + ci
    new_j = j + cj
    if new_i < 0 or new_j < 0 or new_i >= len(grid) or new_j >= len(grid[i]):
      continue # this would take off of grid
    if grid[new_i][new_j] == cur + 1:
      peaks += paths_to_peaks(grid, new_i, new_j)
  return peaks


print(sum(len(set(paths_to_peaks(grid, i, j))) for (i, j) in trailheads))
print(sum(len(paths_to_peaks(grid, i, j)) for (i, j) in trailheads))