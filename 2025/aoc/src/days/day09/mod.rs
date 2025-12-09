use std::collections::HashSet;

use crate::Day;

pub struct Day09 {}

impl Day for Day09 {
    fn solve_1(&self, input: String) -> i64 {
        let points: Vec<Point> = input.split("\n").map(|p: &str| point_from_line(p)).collect();
        let mut max_area = 0;
        for i in 0..points.len()-1 {
            let p = &points[i];
            for j in i..points.len() {
                let other = &points[j];
                if p.area_rectangle(other) > max_area {
                    max_area = p.area_rectangle(other)
                }
            }
        }
        max_area
    }

    fn solve_2(&self, input: String) -> i64 {
        let red: Vec<Point> = input.split("\n").map(|p: &str| point_from_line(p)).collect();
        // make it larger on purpose. 
        let exp_red: Vec<Point> = red.iter().map(|p: &Point| Point{x: p.x*2, y: p.y*2}).collect();
        let edge = edge_tiles(&exp_red);
        let exp_grid = &mut get_grid(&exp_red);

        let mut visited: HashSet<Point> = HashSet::new();
        for i in 0..exp_grid.len() {
            for j in 0..exp_grid[i].len() {
                let np = Point{y: i as i32, x: j as i32};
                if !visited.contains(&np) {
                    visited.insert(np.clone());
                    fill_grid(&np, exp_grid, &mut visited, &edge);
                }
            }
            println!("Finished filling row of grid.");
        }
        
        println!("Finished filling grid");
        print_grid(exp_grid);

        println!();
        let grid = compress_grid(exp_grid.clone());
        print_comp_grid(&grid);

        // print_grid(&grid);

        // all tiles on the edge are red/green.
        // all tiles *within* the loop are red/green.

        // rectangle still must have red tiles in opposite corners. But all tiles within it must be red/green.
        
        // what if i make a grid. and then figure out all colored tiles with "can escape grid" with a visited set. i treat the lines like walls. 
        // to handle ability to squeeze out, i will multiply every point by 2. 

        println!("Grid calculations completed.");
        let mut max_area: i64 = 0;
        for i in 0..red.len()-1 {
            let p = &red[i];
            for j in i..red.len() {
                let other: &Point = &red[j];
                if p.area_rectangle(other) > max_area && p.rectangle_all_enclosed(other, &grid) {
                    println!("New max area: {},{} to {},{}", p.x, p.y, other.x, other.y);
                    max_area = p.area_rectangle(other)
                }
            }
            if i % 1000 == 0 {
                println!("Progress being made. Start corner rows complete: {} out of {}", i+1, red.len()-1)
            }
        }
        max_area
    }
}

fn compress_grid(grid: Vec<Vec<Option<bool>>>) -> Vec<Vec<bool>> {
    // for simplicity, earlier, i multipled all p.x and p.y by 2. i now need to undo that. 
    let mut comp: Vec<Vec<bool>> = vec![];
    for y in 0..grid.len() {
        if y % 2 == 1 {
            continue
        }
        let mut row: Vec<bool> = vec![];
        for x in 0..grid[y].len() {
            if x % 2 == 1 {
                continue
            }
            row.push(grid[y][x].unwrap());
        }
        if y % 1000 == 0 {
            println!("Compression progress being made. rows complete: {} out of {}", y+1, grid.len()-1)
        }
        comp.push(row);
    }
    comp
}

// something is missing with 'visited'. 
// return whether or not start can escape.
fn fill_grid(start: &Point, grid: &mut Vec<Vec<Option<bool>>>, visited: &mut HashSet<Point>, lines: &Vec<Line>) -> bool {
    if visited.len() % 100000 == 0 {
        println!("Making progress! {} out of {} visited so far.", visited.len(), grid.len() as i64 * grid[0].len() as i64)
    }
    if grid[start.y as usize][start.x as usize].is_some() {
        return grid[start.y as usize][start.x as usize].unwrap()
    }
    
    if point_on_a_line(lines, start) {
        grid[start.y as usize][start.x as usize] = Some(false);
        return grid[start.y as usize][start.x as usize].unwrap()
    }

    let y_turns = vec![-1, 0, 1];
    let x_turns = vec![-1, 0, 1];
    for x_turn in &x_turns {
        for y_turn in &y_turns {
            if x_turn == &0 && y_turn == &0 || (x_turn != &0 && y_turn != &0) {
                continue // either self or diagonal.
            }
            let np = Point{x: start.x+x_turn, y: start.y+y_turn};
            if visited.contains(&np) {
                // already visited this point. it should be handled by the recursion. continue.
                continue
            }
            visited.insert(np.clone());
            // if any neighbors can escape, that means i can escape. 
            if fill_grid(&np, grid, visited, lines) {
                grid[start.y as usize][start.x as usize] = Some(true);
                return true
            }
        }
    }
    // no neighbors could escape. that means i cannot escape.
    grid[start.y as usize][start.x as usize] = Some(false);
    false
}

fn point_on_a_line(lines: &Vec<Line>, point: &Point) -> bool {
    for line in lines {
        if line.point_is_on_line(point) {
            return true
        }
    }
    return false
}

fn get_grid(points: &Vec<Point>) -> Vec<Vec<Option<bool>>>{
    let max_x = points.iter().map(|p| p.x).max().unwrap();
    let max_y = points.iter().map(|p| p.y).max().unwrap();
    let w = max_x  as usize + 2;
    let h = max_y as usize + 2;

    // Prebuild the top/bottom row and middle-row template
    let top_or_bottom = vec![Some(true); w];

    let mut middle = vec![None; w];
    middle[0] = Some(true);
    middle[w - 1] = Some(true);

    let mut grid = Vec::with_capacity(h);

    grid.push(top_or_bottom.clone());
    for y in 1..h-1 {
        grid.push(middle.clone());
        if y % 1000 == 0 {
            println!("making grid row {} of len {} done", y, h);
        }
    }
    grid.push(top_or_bottom.clone());

    grid
}

#[derive(Hash, Eq, PartialEq, Clone)]
struct Point {
    x: i32,
    y: i32
}

impl Point {
    fn area_rectangle(&self, other: &Point) -> i64 {
        // xdiff * ydiff inclusive.
        ((self.x - other.x + 1).abs() as i64 * (self.y - other.y  + 1).abs() as i64) as i64
    }

    fn rectangle_all_enclosed(&self, other: &Point, colored_grid: &Vec<Vec<bool>>) -> bool {
        let mut small_y = self.y;
        let mut big_y = other.y;
        if small_y > big_y {
            small_y = other.y;
            big_y = self.y;
        }

        let mut small_x = self.x;
        let mut big_x = other.x;
        if small_x > big_x {
            small_x = other.x;
            big_x = self.x;
        }


        for y in small_y..big_y+1 {
            for x in small_x..big_x+1 {  
                if point_can_escape(&Point{x: x, y: y}, &colored_grid) {
                    return false // can escape, that means not closed in
                }
            }
        }
        return true
    }
}

fn point_can_escape(p: &Point, colored_grid: &Vec<Vec<bool>>) -> bool {
    return colored_grid[p.y as usize][p.x as usize]
}

fn edge_tiles(points: &Vec<Point>) -> Vec<Line> {
    let mut lines = vec![];
    for i in 0..points.len()-1 {
        lines.push(Line{
            start: points[i].clone(),
            end: points[i+1].clone()
        })
    }
    lines.push(Line{
        start: points[points.len()-1].clone(),
        end: points[0].clone()
    });
    lines
}

fn point_from_line(input: &str) -> Point {
    let spl: Vec<i32> = input.split(",").map(|c| c.parse::<i32>().unwrap()).collect();
    Point{x: spl[0], y: spl[1]}
}

struct Line {
    start: Point,
    end: Point
}

impl Line {
    fn point_is_on_line(&self, p: &Point) -> bool {
        // either all x's have to be equal or all y's have to be equal.
        if p.y == self.start.y && p.y == self.end.y {
            return (p.x >= self.start.x && p.y <= self.end.x) || (p.x >= self.end.x && p.x <= self.start.x)
        }

        if p.x == self.start.x && p.x == self.end.x {
            // p.y must be between start.y and end.y but i dont know hwich is bigger.
            return (p.y >= self.start.y && p.y <= self.end.y) || (p.y >= self.end.y && p.y <= self.start.y)
        }
        return false
    }
}

fn print_grid(grid: &Vec<Vec<Option<bool>>>) {
    for line in grid {
        let formatted = line.iter().map(|x| opt_bool_to_char(*x).to_string()).collect::<Vec<_>>().join("");
        println!("{}", formatted);
    }
}

fn print_comp_grid(grid: &Vec<Vec<bool>>) {
    for line in grid {
        let formatted = line.iter().map(|x| bool_to_char(*x).to_string()).collect::<Vec<_>>().join("");
        println!("{}", formatted);
    }
}

fn bool_to_char(b: bool) -> char {
    if b {
        '.'
    } else {
        'X'
    }
}

fn opt_bool_to_char(b: Option<bool>) -> char {
    if b.is_some() {
        return bool_to_char(b.unwrap());
    }
    'N'
}