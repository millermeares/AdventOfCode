garden = []
with open("2024/12/input.txt") as file:
  for line in file:
    garden.append(list(line.rstrip()))


# returns the area and diff (or non existent) perimeter for the plots traversed.
def get_area_perim_for_region(grid, x, y, visited, full_side):
  cur = grid[y][x]
  visited.add((x, y))

  plot_count = 0
  plot_perim = 0
  # gets area and perimeter of region.
  steps = [(-1, 0), (1, 0), (0, 1), (0, -1)]
  for (cx, cy) in steps:
    (nx, ny) = (cx+x, cy+y)
    if (nx < 0 or ny < 0 or ny >= len(grid) or nx >= len(grid[ny])):
      # on the edge of board. need to increment perimeter
      plot_perim += 1
      continue
    if grid[ny][nx] != cur:
      plot_perim += 1
      continue 

    if (nx, ny) not in visited:
      nx_plot_count, nx_perim = get_area_perim_for_region(grid, nx, ny, visited, full_side)
      plot_count += nx_plot_count
      plot_perim += nx_perim

  return plot_count+1, plot_perim
      

def total_price(garden, full_side):
  price = 0
  starts = [
    (x, y) for y, row in enumerate(garden) for x, _ in enumerate(row)
  ]
  visited = set()
  for (x, y) in starts:
    if (x, y) in visited:
      continue # region already considered
    area, perimeter = get_area_perim_for_region(garden, x, y, visited, full_side)
    print(f"Region {garden[y][x]} has {area} area, {perimeter} perimeter")
    price += (area * perimeter)
  return price




print(total_price(garden, False))
print(total_price(garden, True))