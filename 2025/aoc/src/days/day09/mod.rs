use crate::Day;

pub struct Day09 {}

impl Day for Day09 {
    fn solve_1(&self, input: String) -> i64 {
        let points: Vec<Point> = input.split("\n").map(|p| point_from_line(p)).collect();
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
        todo!()
    }
}

#[derive(Eq, PartialEq)]
struct Point {
    x: i32,
    y: i32
}

impl Point {
    fn area_rectangle(&self, other: &Point) -> i64 {
        // xdiff + ydiff 
        ((self.x - other.x + 1).abs() as i64 * (self.y - other.y  + 1).abs() as i64) as i64
    }
}

fn point_from_line(input: &str) -> Point {
    let spl: Vec<i32> = input.split(",").map(|c| c.parse::<i32>().unwrap()).collect();
    Point{x: spl[0], y: spl[1]}
}