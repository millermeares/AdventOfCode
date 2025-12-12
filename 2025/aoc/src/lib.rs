use std::{any::type_name, env, fs::File, io::Read};

mod days;

#[macro_export]
macro_rules! timed {
    ($label:literal, $expr:expr) => {{
        let t = std::time::Instant::now();
        let result = $expr;
        let dt = t.elapsed();
        println!();
        println!("{}: {:?}  (took {:?})", $label, result, dt);
        result
    }};
}


pub trait Day {
    fn solve_1(&self, input: String) -> i64;
    fn solve_2(&self, input: String) -> i64;

    fn solve_sample_1(&self) -> i64 {
        let sample_input = self.get_input(self.get_day(), "sample");
        return self.solve_1(sample_input);
    }

    fn solve_real_1(&self) -> i64 {
        let real_input: String = self.get_input(self.get_day(), "input");
        return self.solve_1(real_input);
    }

    fn solve_sample_2(&self) -> i64 {
        let sample_input = self.get_input(self.get_day(), "sample");
        return self.solve_2(sample_input);
    }

    fn solve_real_2(&self) -> i64 {
        let real_input: String = self.get_input(self.get_day(), "input");
        return self.solve_2(real_input);
    }

    fn get_day(&self) -> i32 {
        let t = type_name::<Self>();
        let s: String = t
            .chars()
            .rev()
            .take(2)
            .collect::<Vec<_>>()
            .into_iter()
            .rev()
            .collect();
        s.parse::<i32>().unwrap()
    }
    fn get_input(&self, d: i32, t: &str) -> String {
        let path = get_path(d, t);
        let mut file = match File::open(&path) {
            Err(why) => panic!("could not open {}; {}", path, why),
            Ok(file) => file,
        };
        let mut s = String::new();
        file.read_to_string(&mut s).unwrap();
        s
    }

}

fn get_path(d: i32, t: &str) -> String {
    // this is directory from which cargo run was run.
    let current_dir: std::path::PathBuf = env::current_dir().unwrap();
    let formatted_day = format!("day{:0>2}", d);
    let p =
        current_dir.to_str().unwrap().to_owned() + &(format!("/src/days/{formatted_day}/{t}.txt"));
    p
}

pub fn index_of(input: &String, f: char) -> usize {
    return input.chars().position(|c| c == f).unwrap()
}
