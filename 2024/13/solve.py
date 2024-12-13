import sys
class Game:
  def __init__(self, a_diff, b_diff, prize_coords):
    self.a_diff = a_diff
    self.b_diff = b_diff
    self.prize_coords = prize_coords
    self.curr_coords = (0, 0)
    self.presses = {
      'a': 0,
      'b': 0
    }
    self.tokens_spent = 0
    self.token_cost = {
      'a': 3,
      'b': 1
    }
  
  def memo_key_from_state(self):
    # return a memo key for presses and 
    return f"{self.presses['a']},{self.presses['b']}"

  def print_game_state(self):
    print(f"Spent {self.tokens_spent} tokens")
    print(f"Presses: {self.presses}")
    print(f"c: {self.curr_coords}, p: {self.prize_coords}")
    print()

  def press(self, button):
    (xDiff, yDiff) = self.a_diff if button == 'a' else self.b_diff
    (curX, curY) = self.curr_coords
    self.curr_coords = (curX + xDiff, curY + yDiff)
    self.presses[button] += 1
    self.tokens_spent += self.token_cost[button]
  
  def unpress(self, button):
    (xDiff, yDiff) = self.a_diff if button == 'a' else self.b_diff
    (curX, curY) = self.curr_coords
    self.curr_coords = (curX - xDiff, curY - yDiff)
    self.presses[button] -= 1
    self.tokens_spent -= self.token_cost[button]


  def at_prize(self):
    return self.curr_coords == self.prize_coords

  def already_passed(self):
    (curX, curY) = self.curr_coords
    (prizeX, prizeY) = self.prize_coords
    return curX > prizeX or curY > prizeY

  def min_tokens_to_reach(self, memo):
    key = self.memo_key_from_state()
    if key in memo:
      return memo[key]
    # self.print_game_state()
    if self.at_prize():
      return self.tokens_spent # made it!
    if self.already_passed():
      return sys.maxsize
    buttons = self.token_cost.keys()
    min_tokens = sys.maxsize
    for button in buttons:
      if self.presses[button] >= 100:
        continue # already pressed this button 100 times
      self.press(button)
      min_tokens = min(min_tokens, self.min_tokens_to_reach(memo))
      self.unpress(button) # undo
    memo[key] = min_tokens
    return memo[key]
  
  def calculate_tokens_to_reach(self, limit):
    (ax, ay) = self.a_diff
    (bx, by) = self.b_diff
    (px, py) = self.prize_coords
    tokens = sys.maxsize
    # figure out combinations that make it reachable. 
    for a in range(0, limit+1):
      for b in range(0, limit+1):
        x_reach = (ax * a) + (bx * b) == px
        y_reach = (ay * a) + (by * b) == py
        if not x_reach or not y_reach:
          continue
        iter_tokens = (a * 3) + (b * 1)
        if iter_tokens < tokens:
          tokens = iter_tokens
          print(f"New cheaper way - a {a} times, b {b} times")
    return tokens





def get_button_diff(line):
  suff = line.split(":")[1].lstrip().split(",")
  return (int(suff[0].split("+")[1]), int(suff[1].split("+")[1]))

games = []
with open("2024/13/input.txt") as file:
  game_in = []
  for line in file:
    if line.rstrip() == '':
      # parse game from game_in
      a_diff = get_button_diff(game_in[0])
      b_diff = get_button_diff(game_in[1])
      prize_suff = game_in[2].split(":")[1].lstrip().split(",")
      prize_coords = (int(prize_suff[0].split("=")[1]), int(prize_suff[1].split("=")[1]))
      # next game?
      games.append(Game(a_diff, b_diff, prize_coords))
      game_in = [] # reset current game inputs
    else:
      game_in.append(line.rstrip())


def part1(games):
  total_tokens = 0
  for i, game in enumerate(games):
    print(f"Calculating for game {i} of {len(games)}")
    tokens_to_win = game.calculate_tokens_to_reach(100)
    # brute = game.min_tokens_to_reach({})
    if tokens_to_win != sys.maxsize:
      total_tokens += tokens_to_win
      print(f"Calculated tokens {tokens_to_win} for game {i} of {len(games)}")
    else:
      print(f"Could not reach end for game {i}")

  return total_tokens


print(part1(games))

