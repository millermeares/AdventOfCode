use crate::Day;

pub struct Day04 {}

fn count_adjaciencies(grid: &Vec<Vec<char>>, px: usize, py: usize) -> i32 {
    let xmoves: Vec<i32> = vec![-1, 0, 1];
    let ymoves: Vec<i32> = vec![-1, 0, 1];

    let mut adj = 0;
    for my in ymoves {
        for mx in &xmoves {
            if my == 0 && *mx == 0 {
                continue
            }

            let x: i32 = px as i32 + mx;
            let y: i32 = py as i32 + my;
            if x < 0 || y < 0 || y as usize >= grid.len() || x as usize >= grid.get(y as usize).unwrap().len() {
                continue
            }
            let c = grid.get(y as usize).unwrap()[x as usize];
            if c == '@' {
                adj += 1;
            }
        }
    }
    adj
}

fn count_of_accessible(grid: &Vec<Vec<char>>) -> i64 {
    let mut accessible: i32 = 0;
    for y in 0..grid.len() {
        let row = &grid.get(y).unwrap();
        for x in 0..row.len() {
            let c = row[x];
            if c == '@' {
                let adj = count_adjaciencies(&grid, x, y);
                if adj < 4 {
                    accessible += 1;
                }
            }
        }
    }
    accessible as i64
}

fn remove_accessibles(grid: &mut Vec<Vec<char>>) -> i64 {
    let mut removed = 0;
    for y in 0..grid.len() {
        for x in 0..grid[y].len() {
            if grid[y][x] == '@' {
                let adj = count_adjaciencies(&grid, x, y);
                if adj < 4 {
                    removed += 1;
                    grid[y][x] = '.';
                }
            }
        }
    }
    removed
}

impl Day for Day04 {
    fn solve_1(&self, input: String) -> i64 {
        let grid: Vec<Vec<char>> = input.split("\n").map(|v| v.to_string().chars().collect()).collect();
        return count_of_accessible(&grid)
    }


    // this is brute force. the optimal solution is some sort of queue. i'll do that next.
    fn solve_2(&self, input: String) -> i64 {
        let mut accessible_removed = 0;
        let grid: &mut Vec<Vec<char>> = &mut input.split("\n").map(|v| v.to_string().chars().collect()).collect();
        while count_of_accessible(&grid) > 0 {
            accessible_removed += remove_accessibles(grid);
        }
        accessible_removed
    }
}