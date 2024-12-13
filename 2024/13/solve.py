import sys
import math

def calculate_tokens_to_reach(ax, ay, bx, by, px, py, limit):
  # unsure if this is appropriate here. it changes output at highest number which is odd.
  if px % math.gcd(ax, bx) != 0 or py % math.gcd(ay, by) != 0:
    return sys.maxsize # not possible!
  # solve system of equations by elimination.
  mult = -1 * (by / bx)
  px_mult = mult * px
  ax_mult = mult * ax
  a_coef = ax_mult + ay
  p_comb = px_mult + py
  a = p_comb / a_coef
  b = (px - (a * ax)) / bx
  if not is_almost_integer(a) or not is_almost_integer(b) or a < 0 or b < 0:
    return sys.maxsize
  a = round(a)
  b = round(b)
  if a > limit or b > limit:
    return sys.maxsize # max size exceeded, this isn't valid
  return a * 3 + b * 1


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

# python highkey annoying with this.
def is_almost_integer(value, tol=1e-3):
    return abs(value - round(value)) < tol

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
  for i, game in enumerate(games):
    (ax, ay) = game[0]
    (bx, by) = game[1]
    (px, py) = game[2]
    tokens_to_win = calculate_tokens_to_reach(ax, ay, bx, by, px + add_prize_dist, py + add_prize_dist, limit)
    if tokens_to_win != sys.maxsize:
      tokens += tokens_to_win
  return tokens

print(calculate(100, 0))
print(calculate(sys.maxsize, 10000000000000))
