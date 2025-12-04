use std::collections::VecDeque;

use crate::Day;

pub struct Day04 {}

fn count_adjaciencies(grid: &Vec<Vec<char>>, px: usize, py: usize) -> i32 {
    box_neighbors(grid, px, py).len() as i32
}

fn is_accessible(grid: &Vec<Vec<char>>, px: usize, py: usize) -> bool {
    return count_adjaciencies(grid, px, py) < 4
} 

fn box_neighbors(grid: &Vec<Vec<char>>, px: usize, py: usize) -> Vec<(usize, usize)> {
    let mut neighbors: Vec<(usize, usize)> = vec![];
    let xmoves: Vec<i32> = vec![-1, 0, 1];
    let ymoves: Vec<i32> = vec![-1, 0, 1];
    for my in ymoves {
        for mx in &xmoves {
            // is self
            if my == 0 && *mx == 0 {
                continue
            }

            let x: i32 = px as i32 + mx;
            let y: i32 = py as i32 + my;
            // is out of grid.
            if x < 0 || y < 0 || y as usize >= grid.len() || x as usize >= grid.get(y as usize).unwrap().len() {
                continue
            }
            // is not box
            if grid[y as usize][x as usize] == '@' {
                neighbors.push((x as usize, y as usize));
            }
        }
    }
    neighbors
}

fn get_accessibles(grid: &Vec<Vec<char>>) -> Vec<(usize, usize)> {
    let mut accessibles= vec![];
    for y in 0..grid.len() {
        let row = &grid.get(y).unwrap();
        for x in 0..row.len() {
            let c = row[x];
            if c == '@' {
                if is_accessible(grid, x, y) {
                    accessibles.push((x, y));
                }
            }
        }
    }
    accessibles
}

impl Day for Day04 {
    fn solve_1(&self, input: String) -> i64 {
        let grid: Vec<Vec<char>> = input.split("\n").map(|v| v.to_string().chars().collect()).collect();
        return get_accessibles(&grid).len() as i64
    }


    fn solve_2(&self, input: String) -> i64 {
        let mut accessible_removed = 0;
        let grid: &mut Vec<Vec<char>> = &mut input.split("\n").map(|v: &str| v.to_string().chars().collect()).collect();
        
        let mut queue: VecDeque<(usize, usize)> = VecDeque::new();
        let accessibles = get_accessibles(&grid);
        for accessible in accessibles {
            queue.push_back(accessible);
        }
        while queue.len() > 0 {
            let (nx, ny) = queue.pop_front().unwrap();
            if grid[ny][nx] == '@' && is_accessible(grid, nx, ny) {
                // remove box and increment count.
                accessible_removed += 1;
                grid[ny][nx] = '.';
                // enqueue neighbors who have had their box numbers change.
                let neighbors = box_neighbors(grid, nx, ny);
                for neigh in neighbors {
                    queue.push_back(neigh);
                }
            }
        }
        accessible_removed as i64
    }
}