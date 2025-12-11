use std::collections::{HashMap, HashSet};

use crate::Day;

pub struct Day11 {}

impl Day for Day11 {
    fn solve_sample_2(&self) -> i64 {
        let sample_input = self.get_input(self.get_day(), "sample_2");
        return self.solve_2(sample_input);
    }

    fn solve_1(&self, input: String) -> i64 {
        let devices: Vec<Device> = input.split("\n").map(|d| device_from_line(d.to_string())).collect();
        let mut device_map: HashMap<String, &Device> = HashMap::new();
        for d in &devices {
            device_map.insert(d.name.clone(), d);
        }

        let mut results: Vec<Vec<String>> = vec![];
        let mut path_so_far: Vec<String> = vec![];
        find_paths(&device_map, &"you".to_string(), &"out".to_string(), &mut path_so_far, &mut HashSet::new(), &mut results);
        // find every path from "you" to "out". 
        results.len() as i64
    }

    fn solve_2(&self, input: String) -> i64 {
        let devices: Vec<Device> = input.split("\n").map(|d| device_from_line(d.to_string())).collect();
        let mut device_map: HashMap<String, &Device> = HashMap::new();
        for d in &devices {
            device_map.insert(d.name.clone(), d);
        }

        let mut results: Vec<Vec<String>> = vec![];
        let mut path_so_far: Vec<String> = vec![];
        find_paths(&device_map, &"svr".to_string(), &"out".to_string(), &mut path_so_far, &mut HashSet::new(), &mut results);
        
        let mut visited_both = 0;
        for i in 0..results.len() {
            let path = &results[i];
            println!("Evaluating fft and dac containment for {} out of {} with path len of {} ", i+1, results.len(), path.len());
            if path.contains(&"dac".to_string()) && path.contains(&"fft".to_string()) {
                visited_both += 1;
            }
        }
        visited_both as i64
    }
}

// fn find_paths_containing(graph: &HashMap<String, &Device>, start: &String, end: &String, containing: String) -> Vec<Vec<String>> {

// }

// directed graph that has possible cycles. 
// basically get all of the paths from node to another node.
fn find_paths(graph: &HashMap<String, &Device>, start: &String, end: &String, path: &mut Vec<String>, visited: &mut HashSet<String>, results: &mut Vec<Vec<String>>) {
    visited.insert(start.clone());
    path.push(start.clone());

    if start == end {
        let completed_path = path.clone();
        let qualifies = path_contains_required_nodes(&completed_path);
        results.push(completed_path);
        visited.remove(start);
        path.remove(path.len()-1);
        if results.len() % 1000000 == 0 {
            println!("New path found with qualification {}. Total paths so far: {}.", qualifies, results.len());
        }
        return // we are at destination, we can stop.
    }
    let cur = graph.get(start).unwrap_or_else(|| {
        panic!("Key `{start}` not found in map");
    });
    for way in &cur.outputs {
        if visited.contains(way) {
            continue // already visited this node, re-entering would be a cycle.
        }
        find_paths(&graph, &way, &end, path, visited, results);
    }
    visited.remove(start);
    path.remove(path.len()-1);
}

fn path_contains_required_nodes(path: &Vec<String>) -> bool {
    path.contains(&"dac".to_string()) && path.contains(&"fft".to_string())
}

fn device_from_line(input: String) -> Device {
    let spl: Vec<&str> = input.split(" ").collect();
    let r_name = spl[0];
    let name = &r_name[0..r_name.len()-1];
    let outputs = &spl[1..spl.len()];
    return Device{
        name: name.to_string(),
        outputs: outputs.iter().map(|o| o.to_string()).collect()
    }
}

struct Device {
    name: String,
    outputs: Vec<String>
}