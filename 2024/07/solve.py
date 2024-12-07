eq_in = []
with open("sample.txt") as file:
  for line in file:
    split = line.rstrip().split(':')
    eq_in.append({
      'test_value': int(split[0]),
      'numbers': list(map(int, split[1].lstrip().split(' ')))
    })    

def can_equation_be_true(eq, supported_operations):
  nums = eq['numbers']
  if len(nums) == 1:
    return eq['test_value'] == nums[0]

  any_possible = False
  for op in supported_operations:
    val = do_operation(nums[0], nums[1], op)
    eq['numbers'] = [val] + nums[2:]
    possible = can_equation_be_true(eq, supported_operations)
    eq['numbers'] = nums
    if possible:
      any_possible = True
  return any_possible

def do_operation(num1, num2, op):
  if op == '*':
    return num1 * num2
  elif op == '+':
    return num1 + num2
  else: # concat
    return int(str(num1) + str(num2))


def part1(equations):
  return sum(eq['test_value'] for eq in equations if can_equation_be_true(eq, ['*', '+']))

def part2(equations):
  return sum(eq['test_value'] for eq in equations if can_equation_be_true(eq, ['*', '+', '|']))


print(part1(eq_in))
print(part2(eq_in))