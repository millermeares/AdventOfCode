import sys
import math

  

def extended_gcd(a, b):
  if b == 0:
    return a, 1, 0
  g, x1, y1 = extended_gcd(b, a % b)
  x = y1
  y = x1 - (a // b) * y1
  return g, x, y

def calculate_tokens_to_reach(ax, ay, bx, by, px, py, limit):
  if px % math.gcd(ax, bx) != 0 or py % math.gcd(ay, by) != 0:
    return sys.maxsize # not possible!
  
  tokens = sys.maxsize
  # figure out combinations that make it reachable. 
  max_a = min(limit, math.floor(px / ax), math.floor(py / ay))+1 # add one to make sure that pressing button 0 times is considered
  for a in range(0, max_a):
    if a % 1000000 == 0:
      print(f"{a}/{max_a}")
    if (px - (a *ax)) % bx != 0 or round((px - (a *ax)) / bx) != round((py - (a *ay)) / by) :
      continue # not feasbible
    b = round((px - (a *ax)) / bx)
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
  return [a_diff, b_diff, prize_coords]

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
  tokens = 0
  for i, game in enumerate(games):
    (ax, ay) = game[0]
    (bx, by) = game[1]
    (px, py) = game[2]
    print(f"Calculating game {i}")
    tokens_to_win = calculate_tokens_to_reach(ax, ay, bx, by, px + add_prize_dist, py + add_prize_dist, limit)
    if tokens_to_win != sys.maxsize:
      tokens += tokens_to_win
  return tokens

print(calculate(100, 0))
print(calculate(sys.maxsize, 10000000000000))

