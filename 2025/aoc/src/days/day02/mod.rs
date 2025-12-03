use std::collections::HashSet;

use crate::Day;
use crate::timed;

pub struct Day02 {}

// ranges are inclusive on both sides. 
struct Range {
    begin: String,
    end: String
}

impl Range {
    fn get_invalid(&self, max_copies: i32) -> HashSet<String> {
        let mut search_count = 0;
        let mut invalids: HashSet<String> = HashSet::new();
        let first_possible_start = self.first_possible(max_copies);
        let last_possible_start = first_half(&self.end, false);
        let first_possible_start_num = first_possible_start.parse::<i64>().unwrap();
        let last_possible_start_num = last_possible_start.parse::<i64>().unwrap();
        let mut cur = first_possible_start_num;
        // the perf bottleneck here is 'cur'. cur is basically brute force searching all possibilites. 
        // could maybe be smarter by looking at 'end' earlier.
        while cur <= last_possible_start_num {
            search_count += 1;
            let mut num = format!("{cur}");
            let mut copies = 1;
            while &num.len() <= &self.end.len() && copies < max_copies {
                num = format!("{num}{cur}");
                copies += 1;
                if !invalids.contains(&num) && self.str_in_range(&num) {
                    invalids.insert(num.to_string());
                }
            }
            cur += 1;
        }
        let b = self.begin.to_string();
        let e = self.end.to_string();
        println!("Search count: {search_count} for {b},{e}");
        invalids
    }

    fn num_in_range(&self, i: i64) -> bool {
        let snum = self.begin.parse::<i64>().unwrap();
        let endnum = self.end.parse::<i64>().unwrap();
        return i >= snum && i <= endnum
    }

    fn str_in_range(&self, s: &String) -> bool {
        let i = s.parse::<i64>().unwrap();
        return self.num_in_range(i);
    }

    fn first_possible(&self, max_copies: i32) -> String {
        // x * max_copies <= self.begin.len()
        // x = self.begin.len() / max_copies
        if self.begin.len() == 1 {
            return "0".to_string();
        }
        let desired_len: usize = self.begin.len() / (max_copies as usize);
        let s = &self.begin;
        let h = s[..desired_len].to_string();
        if h.is_empty() {
            return "0".to_string();
        }
        h
    }
}




fn first_half(s: &String, trunc_left: bool) -> String {
    if s.len() == 1 {
        return "0".to_string()
    }
    let mut d = 0;
    if !trunc_left {
        d = 1;
    }
    s[..(s.len() / 2) + d].to_string()
}


impl Day for Day02 {
    fn solve_1(&self, input: String) -> i64 {
        let mut sum = 0;
        let ranges = parse_ranges(input.trim().to_string());
        for range in ranges {
            let invalids = range.get_invalid(2);
            let nums = invalids.iter().map(|s| {
                s.parse::<i64>().unwrap()
            });
            sum += nums.sum::<i64>();
        }
        sum
    }

    fn solve_2(&self, input: String) -> i64 {
        let mut sum = 0;
        let ranges = parse_ranges(input.trim().to_string());
        for range in ranges {
            let invalids = timed!("Range time", range.get_invalid(range.end.len().try_into().unwrap()));
            let nums = invalids.iter().map(|s: &String| {
                s.parse::<i64>().unwrap()
            });
            sum += nums.sum::<i64>();
        }
        sum
    }
}

fn parse_ranges(input: String) -> Vec<Range> {
    let mut ranges: Vec<Range> = vec![];
    let spl = input.split(",");
    spl.for_each(|r| {
        let mut be = r.split("-");
        let start = be.next().unwrap();
        let end = be.next().unwrap();
        ranges.push(Range { begin: start.to_string(), end: end.to_string() })
    });
    ranges
}

mod test {
    use crate::{Day, days::day02::{Day02, first_half}};

    #[test]
    fn test_example_1() {
        let d = Day02{};
        assert_eq!(33, d.solve_1("11-22".to_owned()));
        assert_eq!(99, d.solve_1("95-115".to_owned()));
        assert_eq!(1010, d.solve_1("998-1012".to_owned()));
        assert_eq!(1188511885, d.solve_1("1188511880-1188511890".to_owned()));
        assert_eq!(222222, d.solve_1("222220-222224".to_owned()));
        assert_eq!(0, d.solve_1("1698522-1698528".to_owned()));
        assert_eq!(446446, d.solve_1("446443-446449".to_owned()));
        assert_eq!(38593859, d.solve_1("38593856-38593862".to_owned()));
    }

    #[test]
    fn test_example_2() {
        let d = Day02{};
        assert_eq!(33, d.solve_2("11-22".to_owned()));
        assert_eq!(210, d.solve_2("95-115".to_owned()));
        assert_eq!(2009, d.solve_2("998-1012".to_owned()));
        assert_eq!(1188511885, d.solve_2("1188511880-1188511890".to_owned()));
        assert_eq!(222222, d.solve_2("222220-222224".to_owned()));
        assert_eq!(0, d.solve_2("1698522-1698528".to_owned()));
        assert_eq!(446446, d.solve_2("446443-446449".to_owned()));
        assert_eq!(38593859, d.solve_2("38593856-38593862".to_owned()));
        assert_eq!(565656, d.solve_2("565653-565659".to_owned()));
        assert_eq!(824824824, d.solve_2("824824821-824824827".to_owned()));
        assert_eq!(2121212121, d.solve_2("2121212118-2121212124".to_owned()));
    }



    #[test]
    fn first_half_test() {
        let f = first_half(&"abcd".to_string(), true);
        assert_eq!("ab", f);
        let b = first_half(&"abcde".to_string(), true);
        assert_eq!("ab", b);
        let c = first_half(&"abcde".to_string(), false);
        assert_eq!("abc", c);
    }
}