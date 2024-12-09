input = open("2024/09/sample.txt").readline()

def get_disk_map(input):
  # take input and read as alternating sets of 'file size (and id index) and file space
  is_file = True
  disk = []
  file_num = 0
  for c in input:
    if is_file:
      v = int(c)
      for i in range(0, v):
        disk.append(str(file_num))
      file_num += 1
    else:
      v = int(c)
      # add v blocks of empty space.
      for i in range(0, v):
        disk.append('.')
    is_file = not is_file
  return disk

def get_first_empty_idx(disk):
  for i in range(0, len(disk)):
    if disk[i] == '.':
      return i
  return -1 # no empty spaces in entire array

def get_last_full_idx(disk):
  for i in range(len(disk)-1, -1, -1):
    if disk[i] != '.':
      return i
  return -1 # no non empty spaces in entire array


def compact(disk):
  # moves file blocks from the end of the disk to the leftmost free space block until there are no gaps.
  first_empty = get_first_empty_idx(disk)
  last_full = get_last_full_idx(disk)
  # if first_empty is *after* last_full, then it's already compact. return disk.
  while not first_empty > last_full:
    print(f"First empty is at {first_empty}, last full is at {last_full}")
    v = disk[last_full]
    disk[last_full] = disk[first_empty]
    disk[first_empty] = v
    first_empty = get_first_empty_idx(disk)
    last_full = get_last_full_idx(disk)
  return disk


# this currently tries to fill each empty spot. 
# bad requirement: only evaluating each file once.
def compact_no_fragmentation(disk):
  print(f"Evaluating {''.join(disk)}")
  first_empty = get_first_empty_idx(disk)
  if first_empty == -1: return disk # no more empty sections left.
  # get of first_empty
  empty_count = 0
  for i in range(first_empty, len(disk)):
    if disk[i] == '.':
      empty_count += 1
    else:
      break
  (fs, fe) = find_file_less_than_or_equal_to_size(disk, empty_count)
  if fs == -1 and fe == -1:
    # this section cannot be replaced. recurse evaluate compact rest of disk, everything is good up to this point.
    return disk[:first_empty + empty_count] + compact_no_fragmentation(disk[first_empty + empty_count:])
  # swap disk[first_empty, last_empty_idx+1] with disk[fs, fe+1]
  f = disk[fs: fe+1]
  print(f"Choosing {''.join(f)} to swap into spot.")
  # put file in position
  empty_replaced = disk[first_empty:first_empty + len(f)]
  disk = disk[:first_empty] + f + disk[first_empty + len(f):] 
  disk = disk[:fs] + empty_replaced + disk[fe+1:]
  # if there is no free space to fit the file, the file does not move. chop off the end.

  return compact_no_fragmentation(disk[:fe+1]) + disk[fe+1:] 
  



# returns start and end index of 
def find_file_less_than_or_equal_to_size(disk, size):
  # check to see if the last one has 
  i = len(disk)
  while i >= 1:
    i -= 1
    if disk[i] == '.':
      continue
    # found a non-empty space. check if this file is less than or equal to 'empty' space.
    start_idx = i
    c = disk[i]
    for j in range(i, -1, -1):
      if disk[j] == c:
        start_idx = j
      else:
        break
    # disk[i:j] = the section.
    file_size = i - start_idx + 1
    if file_size <= size:
      print(f"Selected {disk[start_idx:i+1]} as fit for {size}")
      return (start_idx, i)
    i = start_idx # if this one didn't match, move to next one.
    
  return -1, -1

def calculate_checksum(disk):
  total = 0
  for i in range(0, len(disk)):
    if disk[i] == '.':
      continue
    total += (i * int(disk[i]))
  return total

def part1(input):
  disk = get_disk_map(input)
  compacted = compact(disk)
  return calculate_checksum(compacted)
  # calculate checksum

def part2(input):
  disk = get_disk_map(input)
  compacted = compact_no_fragmentation(disk)
  print(''.join(compacted))
  return calculate_checksum(compacted)


print(part1(input))
print(part2(input))