import sys
import math

def calculate_tokens_to_reach_brute(ax, ay, bx, by, px, py, limit):
  if px % math.gcd(ax, bx) != 0 or py % math.gcd(ay, by) != 0:
    return sys.maxsize # not possible!
  
  tokens = sys.maxsize
  # figure out combinations that make it reachable. 
  max_a = min(limit, math.floor(px / ax), math.floor(py / ay))+1 # add one to make sure that pressing button 0 times is considered
  for a in range(0, max_a):
    if (px - (a *ax)) % bx != 0 or round((px - (a *ax)) / bx) != round((py - (a *ay)) / by) :
      continue # not feasbible
    b = round((px - (a *ax)) / bx)
    iter_tokens = (a * 3) + (b * 1)
    if iter_tokens < tokens:
      tokens = iter_tokens
  return tokens

def calculate_tokens_to_reach(ax, ay, bx, by, px, py, limit):
  if px % math.gcd(ax, bx) != 0 or py % math.gcd(ay, by) != 0:
    return sys.maxsize, 0, 0 # not possible!
  # solve system of equations by elimination.
  mult = -1 * (by / bx)
  px_mult = mult * px
  ax_mult = mult * ax
  a_coef = ax_mult + ay
  p_comb = px_mult + py
  a = p_comb / a_coef
  if not is_almost_integer(a):
    return sys.maxsize, 0, 0
  b = round((px - (a * ax)) / bx)
  if a > limit or b > limit:
    return sys.maxsize, 0, 0 # max size exceeded, this isn't valid
  return a * 3 + b * 1, a, b


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

def is_almost_integer(value, tol=1e-9):
    return math.isclose(value, round(value), abs_tol=tol)

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
  games.append(parse_game(game_in)) # read in the last
  return games

games = parse_games()
def calculate(limit, add_prize_dist):
  tokens = 0
  b_t = 0
  for i, game in enumerate(games):
    (ax, ay) = game[0]
    (bx, by) = game[1]
    (px, py) = game[2]
    a_slope = ay / ax
    b_slope = by / bx
    # print(f"Calculating game {i}")
    tokens_to_win_brute = calculate_tokens_to_reach_brute(ax, ay, bx, by, px + add_prize_dist, py + add_prize_dist, limit)
    tokens_to_win, a, b = calculate_tokens_to_reach(ax, ay, bx, by, px + add_prize_dist, py + add_prize_dist, limit)
    if tokens_to_win_brute != sys.maxsize:
      b_t += tokens_to_win
    if tokens_to_win != sys.maxsize:
      tokens += tokens_to_win
    if not is_almost_integer(tokens_to_win):
      print(f"number not whole: {tokens_to_win}")
    if tokens_to_win_brute != round(tokens_to_win):
      print(f"different answers: {tokens_to_win}, {tokens_to_win_brute}. {a} a presses, {b} b presses")
  print(b_t)
  return tokens

print(calculate(100, 0))
# print(calculate(sys.maxsize, 10000000000000))

