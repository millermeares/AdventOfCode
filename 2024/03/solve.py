import re

with open('input.txt', 'r') as file:
    data = file.read().rstrip()


def multiply_match(match):
  # given mul(12,34) formatted, return 12 * 34
  nums = match[4:len(match)-1].split(',')
  return int(nums[0]) * int(nums[1])

def part1():
  matcher = 'mul\(\d{1,3},\d{1,3}\)'
  matches = re.findall(matcher, data)
  total = 0
  for match in matches:
    total += multiply_match(match)
  return total


def part2():
  matcher  = "do\(\)|mul\(\d{1,3},\d{1,3}\)|don\'t\(\)"
  matches = re.findall(matcher, data)
  enabled = True # starts enabled
  total = 0
  for match in matches:
    if (match[0:4]) == 'do()':
      enabled = True
    if (match[0:3]) == 'don':
      enabled = False
    if (match[0:3] == 'mul'):
      if (enabled):
        total += multiply_match(match)
  return total

print(part1())
print(part2())