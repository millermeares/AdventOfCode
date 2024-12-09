grid_in = []
with open("input.txt") as file:
  for line in file:
    grid_in.append(list(line.rstrip()))

def make_antennas_map(grid):
  antennas = {}
  for i in range(0, len(grid)):
    for j in range(0, len(grid[i])):
      if grid[i][j] == '.':
        continue
      if grid[i][j] not in antennas:
        antennas[grid[i][j]] = []
      # (x, y)
      antennas[grid[i][j]].append((j, i))
  return antennas


def is_on_grid(grid, row, col):
  return row >= 0 and col >= 0 and row < len(grid) and col < len(grid[row])

def get_antinodes(grid, do_full_line):
  antinodes = []
  antennas = make_antennas_map(grid)
  for antenna in antennas.keys():
    ant_antinodes = get_antinodes_for_antenna(grid, antennas[antenna], do_full_line)
    antinodes += ant_antinodes
  return list(set(antinodes))


def get_antinodes_for_antenna(grid, antenna_nodes, do_full_line):
  antinodes = []
  for i in range(0, len(antenna_nodes)):
    for j in range(0, len(antenna_nodes)):
      if i == j:
        continue # self
      (fx, fy) = antenna_nodes[i]
      (sx, sy) = antenna_nodes[j]
      if do_full_line:
        x_diff = fx - sx
        y_diff = fy - sy
        m = (sy - fy) / (sx - fx)
        b = sy - (m * sx)
        for ay in range(0, len(grid)):
          ax = round((ay - b) / m)
          if not is_on_grid(grid, ay, ax):
            continue # not on grid: 
          if ((ay - fy) % y_diff == 0 and (ax - fx) % x_diff == 0):
            antinodes.append((ax, ay))
      else: 
        # calculate where antinode would be if 'second' is the midpoint.
        (anti_x, anti_y) = (int(sx + sx - fx), int(sy + sy - fy))
        if is_on_grid(grid, anti_y, anti_x):
          antinodes.append((anti_x, anti_y))
  return antinodes



def print_grid_with_antinodes(grid, antinodes):
  for i in range(0, len(grid)):
    l = ""
    for j in range(0, len(grid[i])):
      p = (j, i)
      if p not in antinodes:
        l += grid[i][j]
      elif grid[i][j] != '.':
        l += grid[i][j]
      else:
        l += "#"
    print(l)


# takes [XO, XO] and turns it onto [XX, OO]
def flip90(data):
  flipped = []
  for i in range(0, len(data[0])):
    flipped_line = ""
    for line in data:
      flipped_line += line[i]
    flipped.append(flipped_line)
  return flipped

print(len(get_antinodes(grid_in, False))) # part1
print(len(get_antinodes(grid_in, True))) # part2
