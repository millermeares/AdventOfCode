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
        let edges = edge_tiles(&red);

        let mut max_area: i64 = 0;
        for i in 0..red.len()-1 {
            let p = &red[i];
            println!("Evaluating {},{}", p.x, p.y);
            for j in i..red.len() {
                let other: &Point = &red[j];
                println!("Evaluating {},{} to {},{}", p.x, p.y, other.x, other.y);
                if p.area_rectangle(other) > max_area && p.rectangle_all_enclosed(other, &edges) {
                    max_area = p.area_rectangle(other);
                    println!("New max area: {} from {},{} to {},{}", max_area, p.x, p.y, other.x, other.y);
                }
            }
            println!("{} of {} complete.", i+1, red.len())
        }
        max_area
    }
}

fn point_on_a_line(lines: &Vec<Line>, point: &Point) -> bool {
    for line in lines {
        if line.point_is_on_line(point) {
            return true
        }
    }
    return false
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

    fn rectangle_all_enclosed(&self, other: &Point, edges: &Vec<Line>) -> bool {
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

        for x in small_x..big_x+1 {
            if point_can_escape(&Point{x: x, y: small_y}, &edges) {
                return false;
            }
            if point_can_escape(&Point{x: x, y: big_y}, &edges) {
                return false;
            }
        }

        for y in small_y..big_y+1 {
            if point_can_escape(&Point{x: small_x, y: y}, &edges) {
                return false;
            }
            if point_can_escape(&Point{x: big_x, y: y}, &edges) {
                return false;
            }
        }

        
        return true
    }
}

fn point_can_escape(p: &Point, lines: &Vec<Line>) -> bool {
    if point_on_a_line(lines, p) {  
        return false;
    }
    let edge_point = Point{x: 0, y: p.y};
    let line_to_edge = Line{start: edge_point, end: p.clone()};
    let can_escape = line_to_edge.count_intersecting_lines(lines) % 2 == 0;
    can_escape
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
    fn ordered_by_x(&self) -> (Point, Point) {
        if self.start.x < self.end.x {
            return (self.start.clone(), self.end.clone());
        }
        (self.end.clone(), self.start.clone())
    }

    fn ordered_by_y(&self) -> (Point, Point) {
        if self.start.y < self.end.y {
            return (self.start.clone(), self.end.clone());
        }
        (self.end.clone(), self.start.clone())
    }


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

    fn is_vertical(&self) -> bool {
        return self.start.x == self.end.x;
    }

    fn is_horizontal(&self) -> bool {
        self.start.y == self.end.y
    }

    fn intersects(&self, line: &Line) -> bool {
        if self.is_vertical() == line.is_vertical() {
            return false // either both vertical or both horizontal.
        }
        
        let horiz_y = self.start.y; // end.y and start.y are equal.
        // self.y must be between line.lower.y and line.higher.y
        let (sy, ey) = line.ordered_by_y();
        if !(horiz_y >= sy.y && horiz_y <= ey.y) {
            return false;
        }

        let vert_x = line.start.x;
        let (sx, ex) = self.ordered_by_x();
        if !(vert_x >= sx.x && vert_x <= ex.x) {
            return false;
        }
        return true;
    }

    fn count_intersecting_lines(&self, lines: &Vec<Line>) -> i32 {
        let mut intersections = 0;
        if self.is_vertical() {
            panic!("This is only written for horizontal lines");
        }
        for line in lines {
            if self.intersects(&line) {
                intersections += 1;
            }
        }
        intersections
    }
}
