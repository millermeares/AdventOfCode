use crate::Day;

pub struct Day03 {}

struct Bank {
    batteries: Vec<i32>
}

impl Bank {
    fn max_joltage(&self) -> i64 {
        // find two largest numbers in my batteries array. 
        let mut max_num: i32 = 0;
        for l in 0..self.batteries.len()-1 {
            for r in l+1..self.batteries.len() {
                let lnum = self.batteries.get(l).unwrap();
                let rnum = self.batteries.get(r).unwrap();
                let num = format!("{lnum}{rnum}").parse::<i32>().unwrap();
                if num > max_num {
                    max_num = num;
                }
            }
        }
        max_num as i64
    }
}

impl Day for Day03 {
    fn solve_1(&self, input: String) -> i64 {
        let banks = parse_banks(input);
        let mut max_joltage = 0;
        banks.iter().for_each(|b| max_joltage += b.max_joltage());
        max_joltage
    }

    fn solve_2(&self, input: String) -> i64 {
        // needs to be some search algorithm which is like 'take biggest' starting with this. DFS basically i think with recursion.
        todo!()
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
    }
}