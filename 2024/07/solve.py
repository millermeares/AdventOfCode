eq_in = []
with open("input.txt") as file:
  for line in file:
    split = line.rstrip().split(':')
    eq_in.append({
      'test_value': int(split[0]),
      'numbers': list(map(int, split[1].lstrip().split(' ')))
    })    

def can_equation_be_true(eq, allow_concat):
  nums = eq['numbers']
  if len(nums) == 1:
    return eq['test_value'] == nums[0]
  # take the first two elements, do an operation, and combine them into one.
  add_val = nums[0] + nums[1]
  add_possible = can_equation_be_true({
    'test_value': eq['test_value'],
    'numbers': [add_val] + nums[2:]
  }, allow_concat)
  mult_val = nums[0] * nums[1]
  mult_possible = can_equation_be_true({
    'test_value': eq['test_value'],
    'numbers': [mult_val] + nums[2:]
  }, allow_concat)
  concat_value = int(str(nums[0]) + str(nums[1]))
  concat_possible = can_equation_be_true({
    'test_value': eq['test_value'],
    'numbers': [concat_value] + nums[2:]
  }, allow_concat)

  return add_possible or mult_possible or (allow_concat and concat_possible)

def part1(equations):
  return sum(eq['test_value'] for eq in equations if can_equation_be_true(eq, False))

def part2(equations):
  return sum(eq['test_value'] for eq in equations if can_equation_be_true(eq, True))


print(part1(eq_in))
print(part2(eq_in))