input = open("2024/09/input.txt").readline()

def s(arr):
  return ''.join(arr)

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
    v = disk[last_full]
    disk[last_full] = disk[first_empty]
    disk[first_empty] = v
    first_empty = get_first_empty_idx(disk)
    last_full = get_last_full_idx(disk)
  return disk


# what is recursion exit case?? i think it will happen automatically
# we only want to move a file once. not sure if we are doing that right at the moment.
def compact_no_fragmentation(disk):
  (fs, fe) = get_next_file(disk)
  if fs == -1 and fe == -1:
    return disk # no more files to find - already compact
  
  f = disk[fs:fe+1]
  print(f"Evaluating moving file {f}")
  # find the left-most free space that would fit the file. space must be before
  (es, ee) = find_leftmost_free_space(disk[:fs], fe - fs + 1)
  if es == -1 and ee == -1:
    # this file cannot fit in any space- recurse and just append rest.
    return compact_no_fragmentation(disk[:fs]) + disk[fs:]
  # swap disk[fs:fe+1], disk[es, ee+1]
  e = disk[es:es + len(f)]
  disk = disk[:es] + f + disk[es + len(f):]
  disk = disk[:fs] + e + disk[fe+1:]
  # everything after fs has already been considered
  return compact_no_fragmentation(disk[:fs]) + disk[fs:] 


def get_next_file(disk):
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
    return (start_idx, i)
  return -1, -1 # no more files to move?


def find_leftmost_free_space(disk, size):
  i = -1
  while i < len(disk) - 1:
    i += 1
    if disk[i] != '.':
      continue 
    # we are at left-most 
    end_idx = i
    for j in range(i, len(disk)):
      if disk[j] == '.':
        end_idx = j
      else:
        break
    empty_size = end_idx - i + 1
    if empty_size >= size:
      return (i, end_idx)
    i = end_idx
  return -1, -1

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


#print(part1(input))
print(part2(input))