from abc import ABC, abstractmethod

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
  def __init__(self):
    self.val = 0
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
    if (registers['A']) == 0:
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
  pointer = InstructionPointer()
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

def find_lowest_copy(registers, program):
  output = []
  a_init = -1
  e = 0
  while output != program:
    if a_init % 10000 == 0:
      print(f"{a_init} A start value did not result in output, {e} instructions executed")
      e = 0
    a_init += 1
    registers['A'] = a_init
    registers['B'] = 0
    registers['C'] = 0
    output, executed = run(registers, program, True)
    e += executed
  return a_init


output, _ = run(registers, program, False)
print(','.join(map(str, output)))
print(find_lowest_copy(registers, program))