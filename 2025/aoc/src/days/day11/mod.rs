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
        let mut graph: HashMap<String, &Device> = HashMap::new();
        for d in &devices {
            graph.insert(d.name.clone(), d);
        }

        // ok so i need:
        // f_o = fft to out (no dac or svr in path)
        let f_o = find_paths_without_prohibited(&graph, &"fft", &"out", &vec!["dac", "svr"]);
        println!("Calculated {} paths from {} to {} without {} or {}", f_o.len(), "fft", "out", "dac", "svr");
        // d_o = dac to out (no fft or svr in path)
        let d_o = find_paths_without_prohibited(&graph, &"dac", &"out", &vec!["fft", "svr"]);
        println!("Calculated {} paths from {} to {} without {} or {}", d_o.len(), "dac", "out", "fft", "svr");

        // f_d = fft to dac (no svr or out in path)
        let f_d: Vec<Vec<String>> = find_paths_without_prohibited(&graph, &"fft", &"dac", &vec!["svr", "out"]);
        println!("Calculated {} paths from {} to {} without {} or {}", f_d.len(), "fft", "dac", "svr", "out");

        // d_f = dac to fft (no svr or out in path)
        let d_f = find_paths_without_prohibited(&graph, &"dac", &"fft", &vec!["svr", "out"]);
        println!("Calculated {} paths from {} to {} without {} or {}", d_f.len(), "dac", "fft", "svr", "out");

        // s_f = svr to fft (no dac or out in path)
        let s_f = find_paths_without_prohibited(&graph, &"svr", &"fft", &vec!["dac", "out"]);
        println!("Calculated {} paths from {} to {} without {} or {}", s_f.len(), "svr", "fft", "dac", "out");

        // s_d = svr to dac (no fft or out in path)
        let s_d = find_paths_without_prohibited(&graph, &"svr", &"dac", &vec!["fft", "out"]);
        println!("Calculated {} paths from {} to {} without {} or {}", s_d.len(), "svr", "dac", "fft", "out");


        // really 2 different ways to do this. 
        // svr -> fft -> dac -> out
        // svr -> dac -> fft -> out

        // s_f * f_d * d_o
        // + 
        // s_d * d_f * f_o
        ((s_f.len() * f_d.len() * d_o.len()) + (s_d.len() * d_f.len() * f_o.len())) as i64
    }
}

fn find_paths_without_prohibited(graph: &HashMap<String, &Device>, start: &str, end: &str, prohibited: &Vec<&str>) -> Vec<Vec<String>> {
    // i can 'prohibit visiting' by manipulating the visited map 
    let mut results: &mut Vec<Vec<String>> = &mut vec![];
    let mut path_so_far: &mut Vec<String> = &mut vec![];
    let mut visited: &mut HashSet<String> = &mut HashSet::new();
    for node in prohibited {
        visited.insert(node.to_string().clone());
    }
    find_paths(graph, &start.to_string(), &end.to_string(), path_so_far, visited, results);
    results.to_vec()
}

// directed graph that has possible cycles. 
// basically get all of the paths from node to another node.
fn find_paths(graph: &HashMap<String, &Device>, start: &String, end: &String, path: &mut Vec<String>, visited: &mut HashSet<String>, results: &mut Vec<Vec<String>>) {
    visited.insert(start.clone());
    path.push(start.clone());

    if start == end {
        let completed_path = path.clone();
        results.push(completed_path);
        visited.remove(start);
        path.remove(path.len()-1);
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

fn get_memo() {

}