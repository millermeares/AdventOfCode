import sys
import math

  
def calculate_tokens_to_reach(game, limit, add_prize_dist = 0):
  print(f"Beginning game calculations for {game['prize_coords']}")
  (ax, ay) = game['a_diff']
  (bx, by) = game['b_diff']
  (px, py) = game['prize_coords']
  px += add_prize_dist
  py += add_prize_dist
  print(f"Beginning game calculations for {(px, py)}")
  tokens = sys.maxsize
  # figure out combinations that make it reachable. 
  max_a = min(limit, math.floor(px / ax), math.floor(py / ay))+1 # add one to make sure that pressing button 0 times is considered
  max_b = min(limit, math.floor(px / bx), math.floor(py / by))+1
  for a in range(0, max_a):
    for b in range(0, max_b):
      x_reach = (ax * a) + (bx * b) == px
      y_reach = (ay * a) + (by * b) == py
      if not x_reach or not y_reach:
        continue
      iter_tokens = (a * 3) + (b * 1)
      if iter_tokens < tokens:
        tokens = iter_tokens
  return tokens


def get_button_diff(line):
  suff = line.split(":")[1].lstrip().split(",")
  return (int(suff[0].split("+")[1]), int(suff[1].split("+")[1]))

def parse_game(game_in):
  # parse game from game_in
  a_diff = get_button_diff(game_in[0])
  b_diff = get_button_diff(game_in[1])
  prize_suff = game_in[2].split(":")[1].lstrip().split(",")
  prize_coords = (int(prize_suff[0].split("=")[1]), int(prize_suff[1].split("=")[1]))
  return {
    'a_diff': a_diff,
    'b_diff': b_diff,
    'prize_coords': prize_coords
  }

def parse_games():
  games = []
  with open("2024/13/input.txt") as file:
    game_in = []
    for line in file:
      if line.rstrip() == '':
        games.append(parse_game(game_in))
        game_in = [] # reset current game inputs
      else:
        game_in.append(line.rstrip())
  games.append(parse_game(game_in)) # read in  the las 
  return games

games = parse_games()
def calculate(limit, add_prize_dist):
  return sum(tokens_to_win for game in games if (tokens_to_win := calculate_tokens_to_reach(game, limit, add_prize_dist)) != sys.maxsize)

print(calculate(100, 0))
print(calculate(sys.maxsize, 10000000000000))

