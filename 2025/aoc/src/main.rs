use aoc::Day;
use std::env;

use crate::days::day01::Day01;
use crate::days::day02::Day02;


pub mod days;

macro_rules! timed {
    ($label:literal, $expr:expr) => {{
        let t = std::time::Instant::now();
        let result = $expr;
        let dt = t.elapsed();
        println!();
        println!("{}: {}  (took {:?})", $label, result, dt);
        result
    }};
}


fn main() {
    let args: Vec<String> = env::args().collect();

    let day_arg = &args[1];
    println!("Day: {day_arg}");

    let day = choose_day(day_arg.parse::<i32>().unwrap());
    timed!("Part 1 Sample", day.solve_sample_1());
    timed!("Part 1 Real", day.solve_real_1());
    
    timed!("Part 2 Sample", day.solve_sample_2());
    timed!("Part 2 Real", day.solve_real_2());
}

fn choose_day(d: i32) -> Box<dyn Day> {
    match d {
        1 => Box::new(Day01 {}),
        2 => Box::new(Day02 {}),
        _ => panic!("Could not find day for {d}"),
    }
}

