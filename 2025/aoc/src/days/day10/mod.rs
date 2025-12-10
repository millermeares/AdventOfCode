use std::collections::HashMap;

use crate::Day;

pub struct Day10 {}

impl Day for Day10 {
    fn solve_1(&self, input: String) -> i64 {
        let mut machines: Vec<Machine> = input.split("\n").map(|p: &str| parse_machine(p.to_string())).collect();
        let mut min_presses = 0;
        for i in 0..machines.len() {
            let m = machines.get_mut(i).unwrap();
            let machine_min = m.min_buttons_to_reach_desired(&mut m.get_buttons(), &mut HashMap::new());
            println!("Calculated min {} for machine {} out of {}", machine_min, i+1, machines.len());
            min_presses += machine_min
        }
        min_presses
    }

    fn solve_2(&self, input: String) -> i64 {
        todo!()
    }
}

struct Machine {
    desired_lights: Vec<bool>,
    buttons: Vec<Button>,
    // i bet joltages will just be a dynamic cost to press each button.
    required_joltage: Vec<i32>,
    lights: Vec<bool>,
}



// it is never beneficial to press a button twice in a row. it is never beneficial to enter a cycle. because we often would just end up where we started.
impl Machine {
    fn memo_key(&self, buttons: &Vec<Button>) -> String {
        let b_char_c: Vec<String> = buttons.iter().map(|b| bool_to_char(b.pushed).to_string()).collect();
        let b_key = b_char_c.join("");
        let cur_lights: Vec<String> = self.lights.iter().map(|l| bool_to_char(*l).to_string()).collect();
        let l_key = cur_lights.join("");
        format!("{}-{}", b_key, l_key)
    }
    fn get_buttons(&self) -> Vec<Button> {
        return self.buttons.clone()
    }
    fn flip_lights(&mut self, flipping: &Vec<usize>) {
        for light_idx in flipping {
            self.lights[*light_idx] = !self.lights[*light_idx];
        }
    }

    // can i memoize this? 
    fn min_buttons_to_reach_desired(&mut self, buttons: &mut Vec<Button>, memo: &mut HashMap<String, i64>) -> i64 {
        if self.desired_lights == self.lights {
            return 0;
        }
        let memo_key = self.memo_key(buttons);
        if memo.contains_key(&memo_key) {
            return *memo.get(&memo_key).unwrap();
        }

        let mut min_cost = i64::MAX;
        for i in 0..buttons.len() {
            if buttons[i].pushed {
                continue // already pushed. no value in double-pushing.
            }
            // println!("Pushing button {}", i);
            push_button(self, &mut buttons[i]);
            // if pushing a button makes it impossible, i need to handle that case. 
            let b_min = &self.min_buttons_to_reach_desired(buttons, memo);
            if b_min != &i64::MAX { // it's not possible to reach from this point.
                if b_min+1 < min_cost {
                    min_cost = b_min+1
                }
            } // else, it's impossible if we push this button first. 
            push_button(self, &mut buttons[i]);
        }
        memo.insert(memo_key, min_cost);
        min_cost
    }
}

fn push_button(m: &mut Machine, b: &mut Button) {
    m.flip_lights(&b.lights_to_flip);
    b.pushed = !b.pushed
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

    Machine{
        lights: actual_lights,
        desired_lights,
        buttons: buttons,
        required_joltage: joltages
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