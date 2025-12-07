use crate::Day;

pub struct Day07 {}

impl Day for Day07 {
    fn solve_1(&self, input: String) -> i64 {
        let (split, grid) = fill_grid(input);
        print_grid(&grid);
        split
    }

    fn solve_2(&self, input: String) -> i64 {
        let (_, grid) = fill_grid(input);
        let n_grid = &mut get_zeros_sized_grid(&grid);
        // fill n_grid with fastest pathways. 
        // fill any bottom | with a path of 1. 
        for i in 0..grid[grid.len()-1].len() {
            if grid[grid.len()-1][i] == '|' {
                n_grid[grid.len()-1][i] = 1
            }
        }

        let mut row_i: i32 = grid.len() as i32 -2;
        while row_i >= 0 {
            let row = row_i as usize;
            // fill grid. 
            for col in 0..grid[row].len() {
                if grid[row][col] == '|' || grid[row][col] == 'S' {
                    // just take the one directly below.
                    n_grid[row][col] = n_grid[row+1][col]
                }
                if grid[row][col] == '^' {
                    let mut ways = 0;
                    if col != 0 {
                        ways += n_grid[row+1][col-1]
                    }
                    if col != grid[row].len()-1 {
                        ways += n_grid[row+1][col+1]
                    }
                    n_grid[row][col] = ways
                }
            }
            row_i-=1;
        }

        let mut s_idx: i32 = -1;
        // return n_grid index of s. 
        for i in 0..grid[0].len() {
            if grid[0][i] == 'S' {
                s_idx = i as i32
            }
        }
        print_num_grid(&n_grid);
        n_grid[0][s_idx as usize]
    }
}

fn fill_grid(input: String) -> (i64, Vec<Vec<char>>) {
    let mut split = 0;
        let grid: &mut Vec<Vec<char>> = &mut input.split("\n").map(|v: &str| v.to_string().chars().collect()).collect();
        for row in 0..grid.len()-1 {
            for col in 0..grid[row].len() {
                let c = grid[row][col];
                if c != '|' && c != 'S' {
                    continue
                }
                // okay now let's look below. if below is . just change it. 
                let below = grid[row+1][col];
                if below == '.' {
                    grid[row+1][col] = '|'
                }
                if below == '^' {
                    split += 1;
                    if col != 0 {
                        grid[row+1][col-1] = '|'

                    }
                    if col != grid[row].len()-1 {
                        grid[row+1][col+1] = '|'
                    }
                }
            }
        }
        (split as i64, grid.clone())
}

fn get_zeros_sized_grid(grid: &Vec<Vec<char>>) -> Vec<Vec<i64>> {
    #[allow(unused_variables)]
    let n_grid: Vec<Vec<i64>> = grid.iter().map(|l| l.iter().map(|x| 0).collect()).collect();
    n_grid
}

fn print_grid(grid: &Vec<Vec<char>>) {
    for line in grid {
        let formatted = line.iter().map(|x| x.to_string()).collect::<Vec<_>>().join("");
        println!("{}", formatted);
    }
}

fn print_num_grid(grid: &Vec<Vec<i64>>) {
    for line in grid {
        let formatted = line.iter().map(|x| x.to_string()).collect::<Vec<_>>().join("");
        println!("{}", formatted);
    }
}