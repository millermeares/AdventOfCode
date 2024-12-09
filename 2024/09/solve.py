import time
import math

input = open("2024/09/input.txt").readline()

def s(arr):
  return ''.join(arr)

def record_time(fun):
    def inner_func(*args, **kwargs):
        start_time = time.perf_counter_ns()
        res = fun(*args, **kwargs)
        end_time = time.perf_counter_ns()
        elapsed = get_elapsed_time(start_time, end_time)
        fun_name = f"{fun.__name__}"  # Get the function's name directly
        print(f"Elapsed time for {fun_name}: {elapsed}")
        return res
    return inner_func

def get_elapsed_time(start_ns, end_ns):
    total_nanoseconds = end_ns - start_ns
    total_seconds = total_nanoseconds / 1_000_000_000  # Convert ns to seconds
    hours = int(total_seconds // 3600)
    minutes = int((total_seconds % 3600) // 60)
    seconds = int(total_seconds % 60)
    milliseconds = int((total_nanoseconds % 1_000_000_000) // 1_000_000)
    microseconds = int((total_nanoseconds % 1_000_000) // 1_000)
    nanoseconds = int(total_nanoseconds % 1_000)
    return f"{hours:02}:{minutes:02}:{seconds:02}.{milliseconds:03}{microseconds:03}{nanoseconds:03}"



def get_disk_map(input):
  # take input and read as alternating sets of 'file size (and id index) and file space
  is_file = True
  disk = []
  file_num = 0
  for c in input:
    if is_file:
      for _ in range(0, int(c)):
        disk.append(str(file_num))
      file_num += 1
    else:
      # add v blocks of empty space.
      for _ in range(0, int(c)):
        disk.append('.')
    is_file = not is_file
  return disk

# @record_time
def get_last_full_idx(disk, last_idx_to_search):
  for i in range(last_idx_to_search-1, -1, -1):
    if disk[i] != '.':
      return i
  return -1 # no non empty spaces in entire array


def compact(disk):
  # moves file blocks from the end of the disk to the leftmost free space block until there are no gaps.
  (first_empty, _) = find_leftmost_free_space(disk, 1, 0)
  last_full = get_last_full_idx(disk, len(disk))
  # if first_empty is *after* last_full, then it's already compact. return disk.
  while not first_empty > last_full:
    v = disk[last_full]
    disk[last_full] = disk[first_empty]
    disk[first_empty] = v
    (first_empty, l) = find_leftmost_free_space(disk, 1, first_empty)
    last_full = get_last_full_idx(disk, last_full)
  return disk


# what is recursion exit case?? i think it will happen automatically
# we only want to move a file once. not sure if we are doing that right at the moment.
def compact_no_fragmentation(disk):
  (fs, fe) = get_next_file(disk, len(disk))
  while fs != -1 and fe != -1:
    f = disk[fs:fe+1]
    # print(f"Evaluating moving file {f}")
    # find the left-most free space that would fit the file. space must be before
    (es, ee) = find_leftmost_free_space(disk[:fs], fe - fs + 1, 0)
    if es == -1 and ee == -1:
      # this file cannot fit in any space - move to next file.
      (fs, fe) = get_next_file(disk, fs)
      continue
    # swap disk[fs:fe+1], disk[es, ee+1]
    e = disk[es:es + len(f)]
    disk = disk[:es] + f + disk[es + len(f):]
    disk = disk[:fs] + e + disk[fe+1:]
    # everything after fs has already been considered, move to next file.
    (fs, fe) = get_next_file(disk, fs)
  return disk


def get_next_file(disk, last_index):
  i = last_index
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

# @record_time
def find_leftmost_free_space(disk, size, start_search):
  i = start_search - 1
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

def calculate_checksum(disk):
  total = 0
  for i in range(0, len(disk)):
    if disk[i] == '.':
      continue
    total += (i * int(disk[i]))
  return total

@record_time
def part1(input):
  disk = get_disk_map(input)
  compacted = compact(disk)
  return calculate_checksum(compacted)
  # calculate checksum

@record_time
def part2(input):
  disk = get_disk_map(input)
  compacted = compact_no_fragmentation(disk)
  return calculate_checksum(compacted)


print(part1(input))
print(part2(input))


