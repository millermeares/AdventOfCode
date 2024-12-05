from functools import cmp_to_key
import random

input_rules = []
input_updates = []
with open("input.txt") as file:
  reading_rules = True
  for line in file:
    stripped = line.rstrip()
    if (stripped == ""):
      reading_rules = False
      continue
    if (reading_rules):
      input_rules.append(list(map(int, stripped.split('|'))))
    else:
      input_updates.append(list(map(int, stripped.split(','))))

def update_is_valid(update, rules_set):
  for i in range(0, len(update)):
    page = update[i]
    if page not in rules_set:
      continue # no rules for this character
    # page must be before all of rules_set['before']
    page_must_precede = rules_set[page]['before']
    # check preceding section to see if any came before that should not have
    union_preceding = list(set(update[0:i]) & set(page_must_precede))
    if (len(union_preceding) > 0):
      return False
    
    # page must be after all of rules_set['after']
    page_must_succeed = rules_set[page]['after']
    union_succeeding = list(set(update[i:]) & set(page_must_succeed))
    if (len(union_succeeding) > 0):
      return False
    
  return True


def aggregate_rules(rules):
  rules_set = {}
  for rule in rules:
    if rule[0] not in rules_set:
      rules_set[rule[0]] = {
        'before': [],
        'after': []
      }
    if rule[1] not in rules_set:
      rules_set[rule[1]] = {
        'before': [],
        'after': []
      }
    rules_set[rule[0]]['before'].append(rule[1])
    rules_set[rule[1]]['after'].append(rule[0])
  return rules_set
    

def fix_update(update, rules_set):
  def comp(first, second):
    if second not in rules_set[first]['before'] and second not in rules_set[first]['after']:
      return 0 # not related to each other, don't change.
    # negative value when first should be before second.
    if second in rules_set[first]['before']:
      return -1
    return 1
    

  return sorted(update, key=cmp_to_key(comp))


def get_sorted_pages(rules_set):
  ordered = []
  pages = list(rules_set.keys())
  random.shuffle(pages)
  for page in pages:
    must_succeed = rules_set[page]['after']
    # print(f"{page} must FOLLOW {must_succeed}")
    # insert into ordered after index of last must_succeed value.
    idx = -1
    for succeeder in must_succeed:
      if succeeder not in ordered:
        continue # not inserted yet.
      succ_idx = ordered.index(succeeder)
      if succ_idx >= idx:
        idx = succ_idx
    # print(f"Inserting {page} into {ordered} at idx {idx}")
    ordered.insert(idx+1, page)
  return ordered


def part1(updates, rules):
  rules_set = aggregate_rules(rules)
  total = 0
  for update in updates:
    if (update_is_valid(update, rules_set)):
      total += update[int((len(update)-1)/2)]
  return total

def part2(updates, rules):
  rules_set = aggregate_rules(rules)
  fixed_updates = []
  for update in updates:
    if not update_is_valid(update, rules_set):
        fixed_updates.append(fix_update(update, rules_set))
  return part1(fixed_updates, rules)


print(part1(input_updates, input_rules))
print(part2(input_updates, input_rules))