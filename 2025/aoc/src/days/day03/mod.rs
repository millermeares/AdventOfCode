use crate::{Day, timed};

pub struct Day03 {}

struct Bank {
    batteries: Vec<i32>
}

impl Bank {
    fn max_joltage(&self, start_idx: usize, nums_to_choose: usize) -> i64 {
        if nums_to_choose == 1 {
            return self.batteries[start_idx..].iter().max().unwrap().to_owned() as i64
        }

        let mut first_idx_of_max = start_idx;
        for i in start_idx..self.batteries.len() - nums_to_choose+1 {
            if self.batteries.get(i).unwrap() > self.batteries.get(first_idx_of_max).unwrap() {
                first_idx_of_max = i;
            }
        }
        let rest = self.max_joltage(first_idx_of_max+1, nums_to_choose-1).to_string();
        let s = self.batteries.get(first_idx_of_max).unwrap();
        return format!("{s}{rest}").parse::<i64>().unwrap();
    }
}

impl Day for Day03 {
    fn solve_1(&self, input: String) -> i64 {
        let banks = parse_banks(input);
        let mut max_joltage = 0;
        banks.iter().for_each(|b| max_joltage += b.max_joltage(0, 2));
        max_joltage
    }

    fn solve_2(&self, input: String) -> i64 {
        let banks = parse_banks(input);
        let mut max_joltage = 0;
        banks.iter().for_each(|b| max_joltage += timed!("Max joltage calculation", b.max_joltage(0, 12)));
        max_joltage
    }
}

fn parse_banks(input: String) -> Vec<Bank> {
    let mut banks: Vec<Bank> = vec![];
    for bank in input.split("\n") {
        let batteries: Vec<i32> = bank.chars().map(|x| x.to_digit(10).unwrap() as i32).collect();
        banks.push(Bank{
            batteries: batteries
        })
    }
    banks
}

mod tests {
    #[allow(unused_imports)]
    use crate::{Day, days::day03::Day03};

    #[test]
    fn example_1() {
        let d = Day03{};
        assert_eq!(98, d.solve_1("987654321111111".to_string()));
    }

    #[test]
    fn example_2() {
        let d = Day03{};
        assert_eq!(987654321111, d.solve_2("987654321111111".to_string()));
        assert_eq!(811111111119, d.solve_2("811111111111119".to_string()));
        assert_eq!(434234234278, d.solve_2("234234234234278".to_string()));
        assert_eq!(888911112111, d.solve_2("818181911112111".to_string()));
    }
}