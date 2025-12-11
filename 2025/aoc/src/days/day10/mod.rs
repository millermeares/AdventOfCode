use std::collections::HashMap;

use crate::Day;

pub struct Day10 {}

impl Day for Day10 {
    fn solve_1(&self, input: String) -> i64 {
        let mut machines: Vec<Machine> = input.split("\n").map(|p: &str| parse_machine(p.to_string())).collect();
        let mut min_presses = 0;
        for i in 0..machines.len() {
            let m = machines.get_mut(i).unwrap();
            let machine_min = m.min_buttons_to_reach_lights(&mut m.get_buttons(), &mut HashMap::new());
            println!("Calculated min {} for machine {} out of {}", machine_min, i+1, machines.len());
            min_presses += machine_min
        }
        min_presses
    }

    fn solve_2(&self, input: String) -> i64 {
        let mut machines: Vec<Machine> = input.split("\n").map(|p: &str| parse_machine(p.to_string())).collect();
        let mut min_presses = 0;
        for i in 0..machines.len() {
            let m = machines.get_mut(i).unwrap();
            let machine_min = m.min_buttons_to_reach_joltage(&mut m.get_buttons(), &mut HashMap::new());
            println!("Calculated min {} for machine {} out of {}", machine_min, i+1, machines.len());
            min_presses += machine_min
        }
        min_presses
    }
}

struct Machine {
    desired_lights: Vec<bool>,
    buttons: Vec<Button>,
    desired_joltage: Vec<i32>,
    lights: Vec<bool>,
    joltage: Vec<i32>
}



// it is never beneficial to press a button twice in a row. it is never beneficial to enter a cycle. because we often would just end up where we started.
impl Machine {
    fn get_desired_joltages(&self) -> Vec<Joltage> {
        let mut jolts = vec![];
        for i in 0..self.desired_joltage.len() {
            jolts.push(Joltage{
                idx: i,
                desired_value: self.desired_joltage[i]
            })
        }
        jolts
    }
    fn memo_key_lights(&self, buttons: &Vec<Button>) -> String {
        let b_char_c: Vec<String> = buttons.iter().map(|b| bool_to_char(b.pushed).to_string()).collect();
        let b_key = b_char_c.join("");
        let cur_lights: Vec<String> = self.lights.iter().map(|l| bool_to_char(*l).to_string()).collect();
        let l_key = cur_lights.join("");
        format!("{}-{}", b_key, l_key)
    }

    fn memo_key_jolts(&self) -> String {
        let cur_joltage: Vec<String> = self.joltage.iter().map(|j| j.to_string()).collect();
        let j_key = cur_joltage.join(",");
        format!("{}", j_key)
    }

    fn can_reach_joltage(&self) -> bool {
        for i in 0..self.desired_joltage.len() {
            let desired = self.desired_joltage[i];
            let actual = self.joltage[i];
            if actual > desired {
                return false;
            }
        }
        return true;

    }

    fn get_buttons(&self) -> Vec<Button> {
        return self.buttons.clone()
    }

    fn flip_lights(&mut self, flipping: &Vec<usize>) {
        for light_idx in flipping {
            self.lights[*light_idx] = !self.lights[*light_idx];
        }
    }

    fn increment_jolts(&mut self, incrementing: &Vec<usize>) {
        for jolt_idx in incrementing {
            self.joltage[*jolt_idx] = self.joltage[*jolt_idx] + 1;
        }
    }

    fn decrement_jolts(&mut self, decrementing: &Vec<usize>) {
        for jolt_idx in decrementing {
            self.joltage[*jolt_idx] = self.joltage[*jolt_idx] - 1;
        }
    }

    fn min_buttons_to_reach_lights(&mut self, buttons: &mut Vec<Button>, memo: &mut HashMap<String, i64>) -> i64 {
        if self.desired_lights == self.lights {
            return 0;
        }
        let memo_key = self.memo_key_lights(buttons);
        if memo.contains_key(&memo_key) {
            return *memo.get(&memo_key).unwrap();
        }

        let mut min_cost = i64::MAX;
        for i in 0..buttons.len() {
            if buttons[i].pushed {
                continue // already pushed. no value in double-pushing (for part 1).
            }
            push_button(self, &mut buttons[i], true);
            let b_min = &self.min_buttons_to_reach_lights(buttons, memo);
            if b_min != &i64::MAX { // it's not possible to reach from this point.
                if b_min+1 < min_cost {
                    min_cost = b_min+1
                }
            } // else, it's impossible if we push this button first. 
            push_button(self, &mut buttons[i], false);
        }
        memo.insert(memo_key, min_cost);
        min_cost
    }

    fn min_buttons_to_reach_joltage(&mut self, buttons: &mut Vec<Button>, memo: &mut HashMap<String, i64>) -> i64 {
        if self.joltage == self.desired_joltage {
            return 0;
        }
        let memo_key = self.memo_key_jolts();
        if memo.contains_key(&memo_key) {
            return *memo.get(&memo_key).unwrap();
        }

        let mut min_cost = i64::MAX;
        if !self.can_reach_joltage() {
            return min_cost;
        }
        for i in 0..buttons.len() {
            push_button(self, &mut buttons[i], true);
            // println!("Trying to push button {i}");
            let b_min = &self.min_buttons_to_reach_joltage(buttons, memo);
            if b_min != &i64::MAX { // it's not possible to reach from this point.
                // println!("Calculated min {} steps to reach when previous min cost was {}.", b_min+1, min_cost);
                if b_min+1 < min_cost {
                    min_cost = b_min+1
                }
            } // else, it's impossible if we push this button first. 
            push_button(self, &mut buttons[i], false);
        }
        memo.insert(memo_key, min_cost); // i am not sure this is right anymore.
        min_cost
    }
}

fn push_button(m: &mut Machine, b: &mut Button, increment_jolts: bool) {
    m.flip_lights(&b.lights_to_flip);
    b.pushed = !b.pushed;
    if increment_jolts {
        m.increment_jolts(&b.lights_to_flip);
    } else {
        m.decrement_jolts(&b.lights_to_flip);
    }
}

#[derive(PartialEq, Eq, Clone)]
struct Button {
    lights_to_flip: Vec<usize>,
    pushed: bool
}

fn parse_machine(input: String) -> Machine {
    let mut desired_lights: Vec<bool> = vec![];
    let light_end = index_of(&input, ']');
    let raw_lights = &input[1..light_end];
    let mut actual_lights: Vec<bool> = vec![];
    for c in raw_lights.chars() {
        actual_lights.push(false); // all lights start turned off.
        if c == '.' {
            desired_lights.push(false);
        } else if c == '#' {
            desired_lights.push(true);
        } else {
            panic!("unexpected light character");
        }
    }

    let jolt_beginning = index_of(&input, '{');

    let buttons: Vec<Button> = parse_buttons(input[light_end+2..jolt_beginning].trim().to_string());
    let raw_joltage = &input[jolt_beginning+1..input.len()-1];
    let joltages: Vec<i32> = raw_joltage.split(",").map(|j| j.parse::<i32>().unwrap()).collect();
    let amt_joltages = joltages.len();
    Machine{
        lights: actual_lights,
        desired_lights,
        buttons: buttons,
        desired_joltage: joltages,
        joltage: vec![0; amt_joltages]
    }
}

fn index_of(input: &String, f: char) -> usize {
    return input.chars().position(|c| c == f).unwrap()
}

fn parse_buttons(input: String) -> Vec<Button> {
    let mut buttons: Vec<Button> = vec![];
    let spl = input.split(" ");
    for rb in spl.into_iter() {
        // trim front and back parentheses.
        let actual = &rb.to_string()[1..rb.len()-1];
        let button_lights: Vec<usize> = actual.split(",").map(|l| l.parse::<usize>().unwrap()).collect();
        buttons.push(Button{
            lights_to_flip: button_lights,
            pushed: false
        })
    }
    buttons
}

fn bool_to_char(b: bool) -> char {
    if b {
        return 'T'
    } else {
        return 'F'
    }
}

fn buttons_with_index(idx: usize, buttons: &Vec<Button>) -> Vec<Button> {
    buttons
        .iter()
        .filter(|b| b.lights_to_flip.contains(&idx))
        .cloned()
        .collect()
}

struct Joltage {
    idx: usize,
    desired_value: i32
}

fn min_button_to_reach_joltage_smart(m: &mut Machine) -> i64 {
    let mut joltages = m.get_desired_joltages();
    let mut buttons_for_each_index: Vec<Vec<Button>> = vec![];
    for i in 0..m.desired_joltage.len() {
        buttons_for_each_index.push(buttons_with_index(i, &m.get_buttons()))
    }

    // this is like calculus with a bunch of dimensions. there are many options, i want min. 


    // we know each button can only be pressed a max of <joltage> times. 

    // ok now it's about the combination of possible button pushes that would yield the desired value in each joltage field. 

    panic!("Cannot reach joltage");
}

// so what is the effiicent solutoin here? some type of lcm i think

mod tests {
    use std::collections::HashMap;

    use crate::days::day10::{min_button_to_reach_joltage_smart, parse_machine};

    #[test]
    fn test_sample_2() {
        let t = "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}".to_string();
        let mut m = parse_machine(t);
        let min = m.min_buttons_to_reach_joltage(&mut m.get_buttons(), &mut HashMap::new());
        assert_eq!(10, min);
    }

    #[test]
    fn test_sample_2_smarter() {
        let t = "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}".to_string();
        let mut m: crate::days::day10::Machine = parse_machine(t);
        let min = min_button_to_reach_joltage_smart(&mut m);

        // system of equations but calculus?
        assert_eq!(10, min);
    }
}