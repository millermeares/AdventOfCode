one = []
two = []
with open("sample.txt") as file:
  for line in file:
    one_two = line.split(' ')
    one.append(int(one_two[0]))
    two.append(int(one_two[3].strip()))

one.sort()
two.sort()


def part1(first, second):
  total_diff = 0
  for i in range(0, len(first)):
    diff = abs(first[i]-second[i])
    total_diff += diff
  return total_diff


def freq(l):
  m = {}
  for v in l:
    if v not in m:
      m[v] = 0
    m[v] = m[v] + 1
  return m

def part2(left, right):
  total_diff = 0
  right_frequency = freq(right)
  for i in range(0, len(left)):
    occurrences = get_occurrences(left[i], right_frequency)
    total_diff += (left[i] * occurrences)
  return total_diff

def get_occurrences(k, frequency):
  if k not in frequency:
    return 0
  return frequency[k]

print(part1(one, two))
print(part2(one, two))

