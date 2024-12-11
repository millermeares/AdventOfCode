from ..helpers.perf import record_time
from functools import cache

stones = open("2024/11/input.txt").read().rstrip().split(' ')

def stone_blink(stone):
  new_stones = []
  if stone == '0':
    new_stones.append('1')
  elif len(stone) % 2 == 0:
    first = stone[:int(len(stone) / 2)]
    second = stone[int(len(stone) / 2):]
    new_stones.append(first)
    new_stones.append(str(int(second)))
  else:
    new_stones.append(str(int(stone) * 2024))
  return new_stones

@cache
def count_stones_after_blinks(stone, blinks):
  if blinks == 0:
    return 1
  next_stones = stone_blink(stone)
  return sum(count_stones_after_blinks(s, blinks-1) for s in next_stones)
  

def do_blinks(stones, blinks):
  return sum(count_stones_after_blinks(s, blinks) for s in stones)

# treat each stone independently, as they do not interact with each other once the simulation starts. 
print(do_blinks(stones, 25))
print(do_blinks(stones, 75))
