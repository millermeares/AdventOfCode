use std::{collections::{HashMap, HashSet}, path};

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

        let path_ct = find_paths(&device_map, &"you".to_string(), &"out".to_string(), &mut HashSet::new(), &mut HashMap::new(), None);
        // find every path from "you" to "out". 
        path_ct
    }

    fn solve_2(&self, input: String) -> i64 {
        let devices: Vec<Device> = input.split("\n").map(|d| device_from_line(d.to_string())).collect();
        let mut graph: HashMap<String, &Device> = HashMap::new();
        for d in &devices {
            graph.insert(d.name.clone(), d);
        }
        let out_device = &Device { name: "out".to_string(), outputs: vec![] };
        graph.insert("out".to_string(), out_device); // out has no paths away.

        let memo: &mut HashMap<String, i64> = &mut HashMap::new();
        // ok so i need:
        
        // d_o = dac to out (no fft or svr in path)
        let d_o = find_paths_without_prohibited(&graph, &"dac", &"out", &vec!["fft", "svr"], memo);
        println!("Calculated {} paths from {} to {} without {} or {}", d_o, "dac", "out", "fft", "svr");

        // f_d = fft to dac (no svr or out in path)
        let f_d = find_paths_without_prohibited(&graph, &"fft", &"dac", &vec!["svr", "out"], memo);
        println!("Calculated {} paths from {} to {} without {} or {}", f_d, "fft", "dac", "svr", "out");

        // d_f = dac to fft (no svr or out in path)
        let d_f = find_paths_without_prohibited(&graph, &"dac", &"fft", &vec!["svr", "out"], memo);
        println!("Calculated {} paths from {} to {} without {} or {}", d_f, "dac", "fft", "svr", "out");

        // s_f = svr to fft (no dac or out in path)
        let s_f = find_paths_without_prohibited(&graph, &"svr", &"fft", &vec!["dac", "out"], memo);
        println!("Calculated {} paths from {} to {} without {} or {}", s_f, "svr", "fft", "dac", "out");

        // s_d = svr to dac (no fft or out in path)
        let s_d = find_paths_without_prohibited(&graph, &"svr", &"dac", &vec!["fft", "out"], memo);
        println!("Calculated {} paths from {} to {} without {} or {}", s_d, "svr", "dac", "fft", "out");

        // f_o = fft to out (no dac or svr in path)
        let f_o = find_paths_without_prohibited(&graph, &"fft", &"out", &vec!["dac", "svr"], memo);
        println!("Calculated {} paths from {} to {} without {} or {}", f_o, "fft", "out", "dac", "svr");


        // really 2 different ways to do this. 
        // svr -> fft -> dac -> out
        // svr -> dac -> fft -> out

        // s_f * f_d * d_o
        // + 
        // s_d * d_f * f_o
        ((s_f * f_d * d_o) + (s_d * d_f * f_o)) as i64
    }
}

fn find_paths_without_prohibited(graph: &HashMap<String, &Device>, start: &str, end: &str, prohibited: &Vec<&str>, memo: &mut HashMap<String, i64>) -> i64 {
    // i can 'prohibit visiting' by manipulating the visited map 
    let visited: &mut HashSet<String> = &mut HashSet::new();
    for node in prohibited {
        visited.insert(node.to_string().clone());
    }
    let can_reach_map: HashMap<String, bool>= build_can_reach_map(graph, &end.to_string(), &prohibited);
    println!("Can reach map size: {}, graph size: {}", can_reach_map.len(), graph.iter().count());
    return find_paths(graph, &start.to_string(), &end.to_string(), visited, memo, Some(&can_reach_map));
}

fn build_can_reach_map(graph: &HashMap<String, &Device>, end: &String, prohibited: &Vec<&str>) -> HashMap<String, bool> {
    let can_reach_map: &mut HashMap<String, bool> = &mut HashMap::new();
    let visited = &mut HashSet::new();
    for node in prohibited {
        visited.insert(node.to_string());
    }
    for d_name in graph.keys() {
        can_reach_map.insert(d_name.clone(), can_reach(graph, d_name, end, visited));
    }
    can_reach_map.clone()
}

fn can_reach(graph: &HashMap<String, &Device>, start: &String, end: &String, visited: &mut HashSet<String>) -> bool {
    if start == end {
        return true
    }
    visited.insert(start.clone());
    let cur = graph.get(start).unwrap_or_else(|| {
        panic!("Key `{start}` not found in map");
    });
    for way in &cur.outputs {
        if can_reach(graph, way, end, visited) {
            visited.remove(start);
            return true
        }
    }
    visited.remove(start);
    return false
}

// directed graph that has possible cycles. 
// basically get all of the paths from node to another node.
fn find_paths(graph: &HashMap<String, &Device>, start: &String, end: &String, visited: &mut HashSet<String>, memo: &mut HashMap<String, i64>, can_reach_map: Option<&HashMap<String, bool>>) -> i64 {
    visited.insert(start.clone());
    if start == end {
        visited.remove(start);
        return 1 // we are at destination, we can stop.
    }
    let memo_key = get_memo_key(start, end, visited);
    if memo.contains_key(&memo_key) {
        return *memo.get(&memo_key).unwrap();
    }

    if can_reach_map.is_some() {
        if !can_reach_map.unwrap().get(start).unwrap() {
            println!("Cannot reach {} from {}, terminating early", end, start);
            return 0 // cannot reach destination from here.
        }
    }
    let cur = graph.get(start).unwrap_or_else(|| {
        panic!("Key `{start}` not found in map");
    });
    let mut total_paths: i64 = 0;
    for way in &cur.outputs {
        if visited.contains(way) {
            continue // already visited this node, re-entering would be a cycle.
        }
        total_paths += find_paths(&graph, &way, &end, visited, memo, can_reach_map);
    }
    visited.remove(start);
    memo.insert(memo_key, total_paths);
    return total_paths;
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

fn get_memo_key(start: &String, end: &String, visited: &mut HashSet<String>) -> String {
    let mut v: Vec<_> = visited.iter().collect();
    v.sort(); // sorts by &String
    let visited_keys = v.iter().map(|s| s.as_str()).collect::<Vec<_>>().join(",");
    format!("{}-{}-{}", start, end, visited_keys)
}