use aoc::{Day, timed};
use std::env;

use crate::days::day01::Day01;
use crate::days::day02::Day02;
use crate::days::day03::Day03;
use crate::days::day04::Day04;
use crate::days::day05::Day05;
use crate::days::day06::Day06;
use crate::days::day07::Day07;
use crate::days::day08::Day08;


pub mod days;


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
        3 => Box::new(Day03 {}),
        4 => Box::new(Day04 {}),
        5 => Box::new(Day05 {}),
        6 => Box::new(Day06 {}),
        7 => Box::new(Day07 {}),
        8 => Box::new(Day08 {}),
        _ => panic!("Could not find day for {d}"),
    }
}

