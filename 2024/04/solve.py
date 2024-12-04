import re

def count_horizontal_xmas(data):
  count = 0
  matcher = 'XMAS'
  for line in data:
    matches = re.findall(matcher, line)
    count += len(matches)
  for line in data:
    matches = re.findall(matcher, line[::-1])
    count += len(matches)
  return count

def reverse(data):
  # return [XOOX, XXXO] into [XOOX, OXXX]
  print(data)
  rev = []
  for line in data:
    rev.append(line[::-1])
  return rev


def count_horizontal_vertical_xmas(data):
  count = count_horizontal_xmas(data)
  flipped = flip90(data)
  return count + count_horizontal_xmas(flipped)


# takes [XO, XO] and turns it onto [XX, OO]
def flip90(data):
  flipped = []
  for i in range(0, len(data[0])):
    flipped_line = ""
    for line in data:
      flipped_line += line[i]
    flipped.append(flipped_line)
  return flipped

def get_all_diagonals(data):
  return get_top_left_to_bottom_right_diagonals(data) + get_top_left_to_bottom_right_diagonals(reverse(data))

def get_top_left_to_bottom_right_diagonals(data):
  diagonals = []
  for row in range(0, len(data)):
    for i in range(0, len(data[row])):
      if row != 0 and i != 0: # all top, then only first.
        continue
      word = ""
      for line_row in range(row, len(data)):
        line = data[line_row]
        idx = i + len(word)
        if (idx >= len(line)):
          continue
        word += line[idx]
      diagonals.append(word)
  return diagonals


def part1(data):
  count = count_horizontal_vertical_xmas(data)
  diagonals = get_all_diagonals(data)
  return count + count_horizontal_xmas(diagonals)





lines = []
with open("input.txt") as file:
  for line in file:
    lines.append(line.rstrip())


print(part1(lines))