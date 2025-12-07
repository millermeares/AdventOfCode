use crate::Day;

struct Operation {
    nums: Vec<i32>,
    op: String
}

impl Operation {
    fn get_value(&self) -> i64 {
        if self.op == "+" {
            return add(&self.nums);
        }
        if self.op == "*" {
            return multiply(&self.nums)
        }
        panic!("Illegal operation {}", self.op)
    }

    #[allow(dead_code)]
    fn print_op(&self) {
        println!("{}", &self.nums.iter().map(|x| x.to_string()).collect::<Vec<_>>()
            .join(&self.op.chars().nth(0).unwrap().to_string()));
    }
}

fn multiply(nums: &Vec<i32>) -> i64 {
    let mut v: i64 = 1;
    for n in nums {
        v *= *n as i64
    }
    v
}

fn add(nums: &Vec<i32>) -> i64 {
    let mut v: i64 = 0;
    for n in nums {
        v += *n as i64
    }
    v
}

pub struct Day06 {}

impl Day for Day06 {
    fn solve_1(&self, input: String) -> i64 {
        let operations = parse_input(input);
        let mut total = 0;
        for op in operations {
            let v = op.get_value();
            total += v;
        }
        total
    }

    fn solve_2(&self, input: String) -> i64 {
        let operations = parse_input_2(input);
        let mut total = 0;
        for op in operations {
            let v = op.get_value();
            total += v;
        }
        total
    }
}

fn parse_input(input: String) -> Vec<Operation> {
    let mut operations: Vec<Operation> = vec![];
    let mut all: Vec<Vec<i32>> = vec![];
    let mut ops_line: Vec<&str> = vec![];
    let spl: Vec<&str> = input.split("\n").collect();
    for line in spl {
        if line.chars().nth(0).unwrap() == '+' ||  line.chars().nth(0).unwrap() == '*' {
            ops_line = to_string_vec(line);
            continue
        }
        all.push(string_to_int_vec(line))
    }
    
    for line in 0..all[0].len() {
        let mut line_nums: Vec<i32> = vec![];
        for a in &all {
            line_nums.push(a[line]);
        }
        operations.push(Operation { nums: line_nums, op: ops_line[line].to_string() })
    }
    operations
}

fn parse_input_2(input: String) -> Vec<Operation> {
    let mut operations: Vec<Operation> = vec![];
    let mut number_lines: Vec<&str> = vec![];
    let mut ops_line: Vec<&str> = vec![];
    let spl: Vec<&str> = input.split("\n").collect();
    for line in spl {
        if line.chars().nth(0).unwrap() == '+' ||  line.chars().nth(0).unwrap() == '*' {
            ops_line = to_string_vec(line);
            continue
        }
        number_lines.push(line)
    }
    
    let mut current_nums: Vec<i32> = vec![];
    let mut i = (number_lines[0].len()-1) as i32;
    while i >= -1 {
        if column_fully_blank(i, &number_lines) {
            let ops_index = ops_line.len() - 1 -  operations.len();
            let this_op = ops_line[ops_index];
            operations.push(Operation { nums: current_nums.clone(), op: this_op.to_string() });
            current_nums = vec![];
            i -= 1;
            continue
        }
        current_nums.push(number_from_column(i as usize, &number_lines));
        i-= 1;
    }

    operations
}

fn number_from_column(i: usize, number_lines: &Vec<&str>) -> i32 {
    let mut s = "".to_string();
    for line in number_lines {
        let c = line.chars().nth(i).unwrap();
        if c != ' ' {
            s.push(c);
        }
    }
    s.parse::<i32>().unwrap()
}

fn column_fully_blank(i: i32, number_lines: &Vec<&str>) -> bool {
    if i == -1 {
        return true // out of bounds => it's blank.
    }
    for line in number_lines {
        if line.chars().nth(i as usize).unwrap() != ' ' {
            return false
        }
    }
    return true
}

fn to_string_vec(input: &str) -> Vec<&str> {
    let mut strs: Vec<&str> = vec![];
    for n in input.split(" ") {
        if n != "" {
            strs.push(n);
        }
    }
    strs
}

fn string_to_int_vec(input: &str) -> Vec<i32> {
    let mut nums = vec![];
    for n in input.split(" ") {
        if n == "" {
            continue
        }
        nums.push(n.parse::<i32>().unwrap());
    }
    nums
}

mod tests {
    #[allow(unused_imports)]
    use crate::days::day06::{Operation, parse_input_2};

    #[test] 
    fn test_sample_1() {
        let op = Operation{
            nums: vec![64, 23, 314],
            op: "+".to_string()
        };
        assert_eq!(401, op.get_value());
    }

    #[test]
    fn test_sample_2() {
        let s = "123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + ";
        parse_input_2(s.to_string());
    }
}