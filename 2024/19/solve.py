from functools import cache

lines = list(map(lambda l: l.rstrip(), open("2024/19/input.txt").readlines()))
towels = list(map(lambda l: l.rstrip().lstrip(), lines[0].split(",")))
patterns = []
for i in range(2, len(lines)):
  patterns.append(lines[i])



@cache
def possible_ways(pattern):
  if len(pattern) == 0:
    return 1
  ways = 0
  for t in towels:
    if len(t) > len(pattern):
      continue # towel too big
    match_so_far = t == pattern[:len(t)]
    if not match_so_far:
      continue
    
    ways += possible_ways(pattern[len(t):])
  return ways


total_ways = 0
s = 0
for i, pattern in enumerate(patterns):
  print(f"Working on {i} pattern {pattern}")
  ways = possible_ways(pattern)
  total_ways += ways
  possible = ways > 0
  if possible:
    s += 1

print(f"Part 1: {s}")
print(f"Part 2: {total_ways}")