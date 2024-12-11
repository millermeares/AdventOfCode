from ..helpers.perf import record_time
from functools import cache

stones = open("2024/11/input.txt").read().rstrip().split(' ')

# can i leverage memoization?
@record_time
def blink(stones):
  new_stones = []
  for i in range(0, len(stones)):
    new_stones += single_blink(stones[i])
  return new_stones

@cache
def single_blink(stone):
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


def do_blinks(stones, blinks):
  while blinks > 0:
    print(f"{blinks} blinks remaining for {len(stones)} stones.")
    blinks -= 1
    stones = blink(stones)
  return stones

# i'm not sure how i can make this faster
# each stone can be treated independently - they do not interact with each other at all once process starts.
print(len(do_blinks(stones, 25)))
print(len(do_blinks(stones, 75)))
