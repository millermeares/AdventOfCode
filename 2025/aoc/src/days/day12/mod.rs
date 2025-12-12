use std::any;

use crate::Day;

pub struct Day12 {}

impl Day for Day12 {
    fn solve_1(&self, input: String) -> i64 {
        let (blueprints, mut spaces, mut required) = parse_input(input);
        let mut can_fit_ct = 0;
        for i in 0..spaces.len() {
            let can_fit = spaces[i].can_place_presents_to_satisfy_required(&blueprints, &mut required[i]);
            println!("Space {} can fit its presents? {}", i, can_fit);
            if can_fit {
                can_fit_ct += 1;
            }
        }
        can_fit_ct as i64
    }

    fn solve_2(&self, input: String) -> i64 {
        todo!()
    }
}

struct Blueprint {
    grid: Vec<Vec<char>>,
    idx: usize
}

impl Blueprint {
    // returns 4 shapes. 
    // in theory, this could be optimized to only return 1 if the shape is entirely symmetrical. none of my input would benefit though.
    fn possible_present_shapes(&self) -> Vec<Present> {
        // shape rotated 4 ways. 
        let original = self.grid.clone();
        let one = flip90(&original);
        let two = flip90(&one);
        let three = flip90(&two);
        vec![Present{ grid: original, blueprint_idx: self.idx}, Present{ grid: one, blueprint_idx: self.idx }, Present { grid: two, blueprint_idx: self.idx }, Present { grid: three, blueprint_idx: self.idx }]
    }
}

struct Present {
    grid: Vec<Vec<char>>,
    blueprint_idx: usize
}

impl Present {
    fn height(&self) -> usize {
        self.grid.len()
    }
    fn width(&self) -> usize {
        self.grid[0].len()
    }

    fn get_populated_points(&self) -> Vec<(usize, usize)> {
        let  mut points: Vec<(usize, usize)> = vec![];
        for y in 0..self.grid.len() {
            for x in 0..self.grid[y].len() {
                if self.grid[y][x] == '#' {
                    points.push((x, y));
                }
            }
        }
        points
    }
}

struct Space {
    grid: Vec<Vec<char>>,
}

impl Space {
    fn can_place_presents_to_satisfy_required(&mut self, blueprints: &Vec<Blueprint>, required: &mut Vec<i32>) -> bool {
        print_grid(&self.grid);

        let any_unsatisfied: bool = required.iter().any(|r| *r > 0);
        if !any_unsatisfied {
            return true // all placed, yay!
        }

        // print_grid(&self.grid);
        let possible_presents = get_all_possible_presents(blueprints, required);
        for present in possible_presents {
            let i = present.blueprint_idx;
            let (placed, (px, py)) = self.greedy_place_present(&present);
            if !placed {
                // not place-able, try to place the next one.
                continue
            }
            // it was successfully placed. decrement required and see if it works.
            decrement_required(i, required);
            if self.can_place_presents_to_satisfy_required(blueprints, required) {
                return true
            }
            // if we were not able to fit everyting with this placement, undo the placement and continue on. 
            increment_required(i, required);
            self.remove_present((px, py), &present);
        }
        false
    }

    fn choose_next_present(&self, blueprint: &Vec<Blueprint>, required: Vec<i32>) {

    }

    // places present at first possible location.
    fn greedy_place_present(&mut self, p: &Present) -> (bool, (usize, usize)) {
        for y in 0..self.grid.len() {
            for x in 0..self.grid[y].len() {
                if self.place_present((x, y), p) {
                    return (true, (x, y)) // nice.
                }
            }
        }
        (false, (0, 0))
    }


    fn place_present(&mut self, (p_x, p_y): (usize, usize), p: &Present) -> bool {
        // if would place present out of bounds, return false early.
        if p_y + p.height()-1 >= self.grid.len() {
            return false
        }
        if p_x + p.width()-1 >= self.grid[0].len() {
            return false
        }
        let populated_points = p.get_populated_points();
        // check if populated.
        for (x, y) in &populated_points {
            if self.grid[y+p_y][x+p_x] == '#' {
                return false // already populated.
            }
        }
        // populate
        for (x, y) in &populated_points {
            self.grid[y+p_y][x+p_x] = '#'
        }
        return true
    }

    fn remove_present(&mut self, (p_x, p_y): (usize, usize), p: &Present) {
        for (x, y) in &p.get_populated_points() {
            if self.grid[y+p_y][x+p_x] != '#' {
                panic!("Trying to remove something that does not exist");
            }
            self.grid[y+p_y][x+p_x] = '.'
        }
    }

    // returns amount of squares in the grid that are '.' and have '#' around them. for now, just starts with single holes. in theory, could also have two. 
    fn count_holes(&self) -> i32 {
        let mut holes = 0;
        for y in 0..self.grid.len() {
            for x in 0..self.grid[y].len() {
                if self.grid[y][x] != '.' {
                    continue
                }
                let left_filled = x == 0 || self.grid[y][x-1] == '#';
                let right_filled = x == self.grid[y].len()-1 || self.grid[y][x+1] == '#';
                let above_filled = y == 0 || self.grid[y-1][x] == '#';
                let below_filled = y == self.grid.len()-1 || self.grid[y+1][x] == '#';
                if left_filled && right_filled && above_filled && below_filled {
                    holes += 1;
                }
            }
        }
        holes
    }
}


fn get_all_possible_presents(blueprints: &Vec<Blueprint>, required: &Vec<i32>) -> Vec<Present> {
    let mut presents: Vec<Present> = vec![];
    for r in 0..required.len() {
        if required[r] == 0 {
            continue
        }
        for present in blueprints[r].possible_present_shapes() {
            presents.push(present);
        }
    }
    presents
}

fn decrement_required(i: usize, required: &mut Vec<i32>) {
    required[i] -= 1;
}

fn increment_required(i: usize, required: &mut Vec<i32>) {
    required[i] += 1;
}

fn parse_input(input: String) -> (Vec<Blueprint>, Vec<Space>, Vec<Vec<i32>>) {
    let double_new_line_split: Vec<&str> = input.split("\n\n").collect();
    let mut blueprints: Vec<Blueprint> = vec![];
    for i in 0..double_new_line_split.len()-1 {
        blueprints.push(blueprint_from_input(i, double_new_line_split[i].split("\n").collect()));
    }

    let last = double_new_line_split[double_new_line_split.len()-1];
    let (spaces, required): (Vec<Space>, Vec<Vec<i32>>) = last.split("\n").map(|s| space_from_input(s.to_string())).collect();
    (blueprints, spaces, required)
}

fn blueprint_from_input(idx: usize, input: Vec<&str>) -> Blueprint {
    let mut lines: Vec<Vec<char>> = vec![];
    for i in 1..input.len() {
        let mut c_line: Vec<char> = vec![];
        for c in input[i].chars() {
            c_line.push(c);
        }
        lines.push(c_line);
    }
    Blueprint { grid: lines, idx: idx }
}

fn space_from_input(input: String) -> (Space, Vec<i32>) {
    let spl: Vec<&str> = input.split(": ").collect();

    let dimensions: Vec<i32> = spl[0].split("x").map(|d| d.parse::<i32>().unwrap()).collect();
    // create empty grid.
    let grid: Vec<Vec<char>> = vec![vec!['.'; dimensions[0] as usize]; dimensions[1] as usize];
    let required: Vec<i32> = spl[1].split(" ").map(|n| n.parse::<i32>().unwrap()).collect();
    return (Space{grid: grid}, required)
}

// used this same concept in 2024 day 8 but for debugging purposes.
fn flip90(data: &[Vec<char>]) -> Vec<Vec<char>> {
    let rows = data.len();
    let cols = data[0].len();

    let mut rotated = vec![vec![' '; rows]; cols];

    for r in 0..rows {
        for c in 0..cols {
            rotated[c][rows - 1 - r] = data[r][c];
        }
    }

    rotated
}

fn print_grid(grid: &Vec<Vec<char>>) {
    println!();
    for row in grid {
        let line: String = row.iter().collect();
        println!("{}", line);
    }
    println!();
}

mod tests {
    use crate::days::day12::{Blueprint, print_grid};

    #[test]
    fn test_rotation() {
        let b = Blueprint{
            grid: vec![
                vec!['#','#','#'],
                vec!['#','.','.'],
                vec!['#','#','#']
            ],
            idx: 0
        };
        for p in b.possible_present_shapes() {
            print_grid(&p.grid);
        }
    }
}