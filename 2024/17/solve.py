from abc import ABC, abstractmethod
import math
import sys

sys.setrecursionlimit(10**6)


lines = open("2024/17/input.txt").readlines()
def read_register_in(line):
  return int(line.split(":")[1].lstrip().rstrip())

registers = {
  'A': read_register_in(lines[0]),
  'B': read_register_in(lines[1]),
  'C': read_register_in(lines[2])
}

program = list(map(int, lines[-1].split(":")[1].rstrip().split(",")))

def combo_operand(operand, registers):
  if operand < 4:
    return operand
  if operand == 4:
    return registers['A']
  if operand == 5:
    return registers['B']
  if operand == 6:
    return registers['C']
  raise Exception(f"Invalid operand: {operand}")

class InstructionPointer():
  def __init__(self, v):
    self.val = v
    self.i = 0

  def get(self):
    return self.val
  
  def add(self, to_add):
    self.val += to_add
    self.i += 1

  def set(self, to_set):
    self.val = to_set
    self.i += 1
  
  def executed(self):
    return self.i

class Instruction(ABC):
  @abstractmethod
  def do(self, operand, registers, pointer):
    pass


class Adv(Instruction):
  def do(self, operand, registers, pointer):
    registers['A'] = int(registers['A'] / pow(2, combo_operand(operand, registers)))
    pointer.add(2)

class Bxl(Instruction):
  def do(self, operand, registers, pointer):
    registers['B'] = registers['B'] ^ operand
    pointer.add(2)

class Bst(Instruction):
  def do(self, operand, registers, pointer):
    registers['B'] = combo_operand(operand, registers) % 8
    pointer.add(2)

class Jnz(Instruction):
  def do(self, operand, registers, pointer):
    if registers['A'] == 0:
      pointer.add(2)
      return
    pointer.set(operand)
    
    
class Bxc(Instruction):
  def do(self, operand, registers, pointer):
    registers['B'] = registers['B'] ^ registers['C']
    pointer.add(2)

class Out(Instruction):
  def __init__(self,output):
    self.output = output

  def do(self, operand, registers, pointer):
    self.output.append(combo_operand(operand, registers) % 8)
    pointer.add(2)


class Bdv(Instruction):
  def do(self, operand, registers, pointer):
    registers['B'] = int(registers['A'] / pow(2, combo_operand(operand, registers)))
    pointer.add(2)


class Cdv(Instruction):
  def do(self, operand, registers, pointer):
    registers['C'] = int(registers['A'] / pow(2, combo_operand(operand, registers)))
    pointer.add(2)

def get_instruction(opcode, output):
  if opcode == 0:
    return Adv()
  if opcode == 1:
    return Bxl()
  if opcode == 2:
    return Bst()
  if opcode == 3:
    return Jnz()
  if opcode == 4:
    return Bxc()
  if opcode == 5:
    return Out(output)
  if opcode == 6:
    return Bdv()
  if opcode == 7:
    return Cdv()
  raise Exception("Unknown opcode")

def out_equal_so_far(output, program):
  if len(output) > len(program):
    return False
  for i in range(0, len(output)):
    if output[i] != program[i]:
      return False
  return True

def run(registers, program, cancel_if_not_equal):
  pointer = InstructionPointer(0)
  output = []
  while pointer.get() < len(program):
    opcode = program[pointer.get()]
    operand = program[pointer.get()+1]
    instruction = get_instruction(opcode, output)
    instruction.do(operand, registers, pointer)
    if cancel_if_not_equal and not out_equal_so_far(output, program):
      break
    # print(f"Did instruction: {type(instruction).__name__} resulting in registers {registers}")
  return output, pointer.executed()

def find_lowest_copy(registers, program, start, up_by, max_tries):
  base_a = start
  output = []
  best = []
  prod_best = []
  tries = 0
  while output != program: # and max_tries > tries:
    if tries % 100000 == 0:
      print(f"{base_a}: {best} so far produced by {prod_best}")
    tries += 1
    base_a += up_by
    a_init = base_a
    registers['A'] = a_init
    registers['B'] = 0
    registers['C'] = 0
    output, _ = run(registers, program, True)
    same = len(best) > 0 and len(best[0]) == len(output)
    if same:
      best.append(output)
      prod_best.append(a_init)
    better = len(best) == 0 or len(output) > len(best[0])
    if better:
      prod_best = [a_init]
      best = [output]

  return base_a, best, prod_best

output, _ = run(registers, program, False)
print(','.join(map(str, output)))
lowest, _, _ = find_lowest_copy(registers, program, 10**12, 1, sys.maxsize)
print(lowest)

# seed_out, seed_in = find_lowest_copy(registers, program, -1, 1, 10 ** 8)
# print(f"Exploring further with seeds, {seed_in} which produced {seed_out}")


# finders = []
# for seed in seed_in:
#   start = seed * int((10 ** 11) / seed)
#   end = 10 ** 14
#   explore_out, explore_in = find_lowest_copy(registers, program, start, seed, 10 ** 8) # just a million tries for now.
#   print(f"First million tries for {seed} produced {explore_out} from {explore_in}")
#   for i in range(0, len(explore_in)):
#     o = explore_out[i]
#     ein = explore_in[i]
#     if o == program:
#       print(f"Produced the appropriate values from {seed} seed and {ein} A register")
#       finders.append(ein)


# if len(finders) > 0:
#   print(f"Found: {min(finders)}")
# else:
#   print("Did not find any producers")
