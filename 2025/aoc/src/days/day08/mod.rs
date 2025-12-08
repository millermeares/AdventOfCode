use std::{collections::HashMap, iter::Map};

use crate::Day;

pub struct Day08 {}

impl Day for Day08 {
    fn solve_1(&self, input: String) -> i64 {
        let points = parse(input);
        let mut unconnected_distances: HashMap<&Point3d, SortedDistanceList<'_>> = get_distances(&points);
        let mut amount = amount_to_connect(&points);
        let mut circuits = init_circuits(&points);
        while amount > 0 {
            // first, figure out 'smallest' new connection.
            let smallest = point_with_shortest_distance(&unconnected_distances);
            let p1 = smallest;
            let p2 = &pop_change(&mut unconnected_distances, p1);
            pop_change(&mut unconnected_distances, p2);

            // second, identify which circuits these points are in.
            let c1 = get_index_of_circuit(p1, &circuits);
            let c2 = get_index_of_circuit(p2, &circuits);

            // third, combine the circuits (if they are different).
            if c1 != c2 {
                let old_circuit = circuits.remove(c2);
                let adjusted_c1 = if c1 > c2 { c1 - 1 } else { c1 };
                let updated = &mut circuits[adjusted_c1];
                updated.combine(old_circuit);
            }
            amount-= 1;
        }
        
        circuits.sort_by(|a, b| b.points.len().cmp(&a.points.len()));
        println!("{}", circuits[0].points.len());
        return (circuits[0].points.len() * circuits[1].points.len() * circuits[2].points.len()) as i64
    }

    fn solve_2(&self, input: String) -> i64 {
        let points = parse(input);
        let mut unconnected_distances: HashMap<&Point3d, SortedDistanceList<'_>> = get_distances(&points);
        let mut circuits = init_circuits(&points);
        let mut last_x_1 = -1;
        let mut last_x_2 = -1;
        while circuits.len() != 1 {
            // first, figure out 'smallest' new connection.
            let smallest = point_with_shortest_distance(&unconnected_distances);
            let p1 = smallest;
            let p2 = &pop_change(&mut unconnected_distances, p1);
            pop_change(&mut unconnected_distances, p2);
            last_x_1 = p1.x;
            last_x_2 = p2.x;

            // second, identify which circuits these points are in.
            let c1 = get_index_of_circuit(p1, &circuits);
            let c2 = get_index_of_circuit(p2, &circuits);

            // third, combine the circuits (if they are different).
            if c1 != c2 {
                let old_circuit = circuits.remove(c2);
                let adjusted_c1 = if c1 > c2 { c1 - 1 } else { c1 };
                let updated = &mut circuits[adjusted_c1];
                updated.combine(old_circuit);
            }
            
        }
        (last_x_1 * last_x_2) as i64
    }
}

fn get_index_of_circuit(point: &Point3d, circuits: &Vec<Circuit>) -> usize {
    for i in 0..circuits.len() {
        for p in &circuits[i].points {
            if p == point {
                return i;
            }
        }
    }
    panic!("Could not find!!");
}

fn pop_change<'a>(unconnected_distances: &mut HashMap<&Point3d, SortedDistanceList<'_>>, p: &Point3d) -> Point3d {
    unconnected_distances.get_mut(p).unwrap().pop().clone()
}


fn init_circuits (points: &Vec<Point3d>) -> Vec<Circuit> {
    points.iter().map(|p| Circuit{points: vec![p.clone()]}).collect()
}
// handles both sample and real input lengths.
fn amount_to_connect(points: &Vec<Point3d>) -> i32 {
    if points.len() > 100 {
        1000
    } else {
        10
    }
}

fn parse(input: String) -> Vec<Point3d> {
    let mut points = vec![];
    for line in input.split("\n") {
        points.push(point_from_line(line));
    }
    points
}

fn point_from_line(input: &str) -> Point3d {
    let coords: Vec<i32> = input.split(",").map(|c| c.parse::<i32>().unwrap()).collect();
    return Point3d { x: coords[0], y: coords[1], z: coords[2] }
}

fn get_distances<'a>(points: &'a Vec<Point3d>) -> HashMap<&'a Point3d, SortedDistanceList<'a>> { 
    let mut distances: HashMap<&Point3d, SortedDistanceList> = HashMap::new();
    for point in points {
        let mut this_distances= SortedDistanceList{points: vec![]};
        for other in points {
            if point == other {
                continue
            }
            this_distances.insert(point, other);
        }
        distances.insert(&point, this_distances);
    }
    distances
}

#[derive(PartialEq, Eq, Hash, Clone)]
struct Point3d {
    x: i32,
    y: i32,
    z: i32,
}

impl Point3d {
    fn straight_line_distance(&self, other: &Point3d) -> f64 {
        // square root of the sum of the squares of each dimension
        let x_diff = (self.x - other.x) as i64;
        let z_diff = (self.z - other.z) as i64;
        let y_diff = (self.y - other.y) as i64;
        let total = (x_diff * x_diff) + (y_diff * y_diff) + (z_diff * z_diff);
        return (total as f64).sqrt();
    }

    fn as_str(&self) -> String {
        format!("{},{},{}" ,self.x, self.y, self.z)
    }
}

struct SortedDistanceList<'a> {
    points: Vec<&'a Point3d>
}

impl <'a> SortedDistanceList<'a> {
    fn peek(&self) -> &Point3d {
        if self.points.len() == 0 {
            panic!("Should not have empty point list");
        }
        self.points.first().unwrap()
    }

    fn insert(&mut self, source: &'a Point3d, other: &'a Point3d) {
        let to_insert = self.get_index_to_insert(source, other, 0, self.points.len());
        self.points.insert(to_insert, other);
    }

    fn pop(&mut self) -> &Point3d {
        self.points.remove(0)
    }

    // just linear. could binary if u want
    fn get_index_to_insert(&self, source: &'a Point3d, other: &'a Point3d, min: usize, max: usize) -> usize {
        let dist = source.straight_line_distance(other);
        if min == max {
            return min; // both 0 for example.
        }

        let mid = (min + max) / 2;
        let mid_dist = source.straight_line_distance(self.points[mid]);
        if dist > mid_dist {
            // greater. use mid_dist+1 as the min
            return self.get_index_to_insert(source, other, mid+1, max)
        } else {
            return self.get_index_to_insert(source, other, min, mid)
        }   
    }
}

fn point_with_shortest_distance<'a>(unconnected_distances: &HashMap<&'a Point3d, SortedDistanceList<'a>>) -> &'a Point3d {
    let mut smallest_point = &Point3d{x: -1, y: -1, z: -1};
    let mut smallest = f64::MAX;
    for p in unconnected_distances.keys() {
        let sorted_list = unconnected_distances.get(p).unwrap();
        let smallest_for_this_point = sorted_list.peek();
        let dist = p.straight_line_distance(smallest_for_this_point);
        if dist < smallest {
            smallest_point = p;
            smallest = dist;
        }
    }
    smallest_point
}


struct Circuit {
    points: Vec<Point3d>
}

impl<'a> Circuit {
    fn combine(&mut self, other: Circuit) {
        for p in other.points {
            self.points.push(p);
        }
    }
}