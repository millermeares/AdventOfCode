use aoc::Day;
use std::env;

use crate::days::day01::Day01;


pub mod days;

fn main() {
    println!("Hello, world!");
    let args: Vec<String> = env::args().collect();
    println!("{:?}", args);

    let day_arg = &args[1];
    println!("Day: {day_arg}");

    let day = choose_day(day_arg.parse::<i32>().unwrap());

    let s_answer = day.solve_sample_1();
    println!("Part 1 Sample: {s_answer}");
    let real_answer = day.solve_real_1();
    println!("Part 1 Real: {real_answer}");

    let s_answer_2 = day.solve_sample_2();
    println!("Part 2 Sample: {s_answer_2}");
    let real_answer_2 = day.solve_real_2();
    println!("Part 2 Real: {real_answer_2}")

}

fn choose_day(d: i32) -> Box<dyn Day> {
    match d {
        1 => Box::new(Day01 {}),
        _ => panic!("Could not find day for {d}"),
    }
}
