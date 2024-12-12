garden = []
with open("2024/12/input.txt") as file:
  for line in file:
    garden.append(list(line.rstrip()))


# returns the area and diff (or non existent) perimeter for the plots traversed.
def get_area_perim_for_region(grid, x, y, visited):
  cur = grid[y][x]
  visited.add((x, y))
  plot_count = 0
  plot_perim = 0
  corners = count_corners(grid, x, y)
  # gets area and perimeter of region.
  steps = [(1, 0), (-1, 0), (0, 1), (0, -1)]
  for (cx, cy) in steps:
    (nx, ny) = (cx+x, cy+y)
    if not on_grid(grid, nx, ny):
      plot_perim += 1
      # on the edge of board. need to increment perimeter
      # trying to step off the board. 
      continue
    if grid[ny][nx] != cur:
      plot_perim += 1
      continue 
    if (nx, ny) not in visited:
      nx_plot_count, nx_perim, nx_corners = get_area_perim_for_region(grid, nx, ny, visited)
      plot_count += nx_plot_count
      plot_perim += nx_perim
      corners += nx_corners
  return plot_count+1, plot_perim, corners

def count_corners(grid, x, y):
  cur = grid[y][x]
  e_corners = 0
  # count external corners
  if not is_matching(grid, cur, x, y+1):
    # above is not matching
    if not is_matching(grid, cur, x+1, y):
      e_corners += 1
    if not is_matching(grid, cur, x-1, y):
      e_corners +=1

  if not is_matching(grid, cur, x, y-1):
    # below is matching.
    if not is_matching(grid, cur, x+1, y):
      e_corners += 1
    if not is_matching(grid, cur, x-1, y):
      e_corners +=1
  # count internal corners. these can be shared.
  i_corners = 0
  diag_steps = [(1, 1), (-1, -1), (1, -1), (-1, 1)]
  for (sx, sy) in diag_steps:
    (dx, dy) = (x + sx, y + sy)
    if not is_matching(grid, cur, dx, dy):
      continue # diagonal doesn't match, so this can't be internal corner
    # check the other two coordinates bordering both (x, y) and (dx, dy). if exactly 1 matches, then increment i_corners
    matches = 0
    (fx, fy) = (x + sx, y)
    if is_matching(grid, cur, fx, fy):
      matches += 1
    (tx, ty) = (x, y + sy)
    if is_matching(grid, cur, tx, ty):
      matches += 1
    # an internal corner is where diagonal matches and only 1 of others matches.
    if matches == 1:
      i_corners += 1


  # internal corners are shared with others, so we only add half.
  return e_corners + (i_corners / 2)
    

def on_grid(grid, x, y):
  return x >= 0 and y >= 0 and y < len(grid) and x < len(grid[y])



def is_matching(grid, cur, nx, ny):
  if not on_grid(grid, nx, ny):
    return False # can't match if you don't exist
  return grid[ny][nx] == cur

def total_price(garden, full_side):
  price = 0
  starts = [
    (x, y) for y, row in enumerate(garden) for x, _ in enumerate(row)
  ]
  visited = set()
  for (x, y) in starts:
    if (x, y) in visited:
      continue # region already considered
    area, perimeter, corners = get_area_perim_for_region(garden, x, y, visited)
    print(f"Region {garden[y][x]} has {area} area, {perimeter} perimeter, {corners} corners")
    if full_side:
      sides = corners
      price += (area * sides)
    else:
      price += (area * perimeter)
  return price




print(total_price(garden, False))
print(total_price(garden, True))