from abc import ABC, abstractmethod
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
    print(f"Literal combo operand {operand}")
    return operand
  if operand == 4:
    print(f"Combo operand from register A")
    return registers['A']
  if operand == 5:
    print(f"Combo operand from register B")
    return registers['B']
  if operand == 6:
    print(f"Combo operand from register C")
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
    val = int(registers['A'] / pow(2, combo_operand(operand, registers)))
    print(f"Adv changing register A from {registers['A']} to {val}")
    registers['A'] = val
    pointer.add(2)

class Bxl(Instruction):
  def do(self, operand, registers, pointer):
    val = registers['B'] ^ operand
    print(f"Bxl changing register B from {registers['B']} to {val}")
    registers['B'] = val
    pointer.add(2)

class Bst(Instruction):
  def do(self, operand, registers, pointer):
    val = combo_operand(operand, registers) % 8
    print(f"Bst changing register B from {registers['B']} to {val}")
    registers['B'] = val
    pointer.add(2)

class Jnz(Instruction):
  def do(self, operand, registers, pointer):
    if registers['A'] == 0:
      print(f"Jnz not setting pointer as Register A has 0.")
      pointer.add(2)
      return
    print(f"Jnz setting pointer to {operand} with {registers}")
    pointer.set(operand)
    
    
class Bxc(Instruction):
  def do(self, operand, registers, pointer):
    val = registers['B'] ^ registers['C']
    print(f"Bxc changing register B from {registers['B']} to {val}")
    registers['B'] = val
    pointer.add(2)

class Out(Instruction):
  def __init__(self,output):
    self.output = output

  def do(self, operand, registers, pointer):
    val = combo_operand(operand, registers) % 8
    print(f"Out outputting {val} from pointer {pointer.get()}")
    print()
    self.output.append(val)
    pointer.add(2)


class Bdv(Instruction):
  def do(self, operand, registers, pointer):
    val = int(registers['A'] / pow(2, combo_operand(operand, registers)))
    print(f"Bdv changing register B from {registers['B']} to {val}")
    registers['B'] = val
    pointer.add(2)


class Cdv(Instruction):
  def do(self, operand, registers, pointer):
    val = int(registers['A'] / pow(2, combo_operand(operand, registers)))
    print(f"Bdv changing register B from {registers['C']} to {val}")
    registers['C'] = val
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
    print(f"Pointer: {pointer.get()}")
    opcode = program[pointer.get()]
    operand = program[pointer.get()+1]
    instruction = get_instruction(opcode, output)
    instruction.do(operand, registers, pointer)
    # print()
    if cancel_if_not_equal and not out_equal_so_far(output, program):
      break
    # print(f"Did instruction: {type(instruction).__name__} resulting in registers {registers}")
  return output

# copied from pignataj github as I've had a very hard time with this problem.
def search(program):
    target = list(reversed(program))
    candidates = []

    # i have no idea how this works!
    def step(A):
        # this is the part that i have to figure out for myself.
        B = A & 7
        B ^= 1
        C = A // (2**B)
        A //= 2**3
        B ^= 4
        B ^= C

        return B & 7

    def find(A, column=0):
        if step(A) == target[column]:
            if column == len(target) - 1:
                yield A
            else:
                for i in range(8):
                    yield from find(A * 8 + i, column + 1)

    for A in range(8):
        candidates.extend(list(find(A)))

    return min(candidates)

# output = run(registers, program, False)
# print(','.join(map(str, output)))
# print(search(program))


registers['A'] = 0
registers['B'] = 0
registers['C'] = 0

run(registers, program, False)