records = []
with open("input.txt") as file:
  for line in file:
    records.append(list(map(int, line.split())))



  def is_safe(record):
    record_increasing = record[0] < record[1]
    for i in range(0, len(record)-1):
      first = record[i]
      second = record[i+1]
      diff = abs(first-second)
      if diff < 1 or diff > 3:
        return False # too large of a step
      is_increasing = second > first
      if (is_increasing != record_increasing):
        return False
    return True

  def is_safe_with_dampener(record):
    if (is_safe(record)):
      return True
    # try removing each level. 
    for i in range(0, len(record)):
      bad_level = record.pop(i)
      if (is_safe(record)):
        return True
      # removing this level didn't save anything. put it back.
      record.insert(i, bad_level)
    return False

  def part1():
    count = 0
    for record in records:
      if is_safe(record):
        count += 1
    return count

  def part2():
    count = 0
    for record in records:
      if is_safe_with_dampener(record):
        count += 1
    return count


print(part1())
print(part2())