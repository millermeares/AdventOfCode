use std::cmp::Ordering;

use crate::Day;

pub struct Day05 {}

struct Range {
    min: i64,
    max: i64
}

impl Eq for Range {

}

impl PartialEq for Range {
    fn eq(&self, other: &Self) -> bool {
        self.min == other.min && self.max == other.max
    }
}

impl PartialOrd for Range {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        return Some(self.cmp(other))
    }
}

impl Ord for Range {
    fn cmp(&self, other: &Self) -> Ordering {
        if self.min == other.min {
            return self.max.cmp(&other.max);
        }
        return self.min.cmp(&other.min)
    }
}

impl Range {
    fn in_range(&self, v: i64) -> bool {
        return v >= self.min && v <= self.max
    }

    fn overlaps(&self, other: &Range) -> bool {
        self.max >= other.min
    }

    fn combine(&self, other: &Range) -> Range {
        if !self.overlaps(other) {
            panic!("cannot combine ranges which do not overlap")
        }
        // take the min and the max
        let mins = vec![self.min, other.min];
        let min = mins.iter().min().unwrap();
        let maxs = vec![self.max, other.max];
        let max = maxs.iter().max().unwrap();
        return Range { min: *min, max: *max }
    }
}

fn ingredient_in_a_range(ingredient: i64, ranges: &Vec<Range>) -> bool {
    for range in ranges {
        if range.in_range(ingredient) {
            return true
        }
    }
    return false
}

// returns false if no overlapping ranges were found.
fn combine_first_overlapping(sorted_ranges: &mut Vec<Range>) -> bool {
    for i in 0..sorted_ranges.len() - 1 {
        if sorted_ranges[i].overlaps(&sorted_ranges[i+1]) {
            let earlier = &sorted_ranges[i];
            let later = &sorted_ranges[i+1];
            let combined = earlier.combine(&later);
            sorted_ranges.remove(i+1);
            sorted_ranges.remove(i);
            sorted_ranges.insert(i, combined);
            return true
        }
    }
    false
}

impl Day for Day05 {
    fn solve_1(&self, input: String) -> i64 {
        let mut fresh = 0;
        let (ranges, ingredients) = parse_input(input);
        for ingredient in ingredients {
            if ingredient_in_a_range(ingredient, &ranges) {
                fresh += 1
            }
        }
        fresh as i64
    }

    fn solve_2(&self, input: String) -> i64 {
        let (mut ranges, _) = parse_input(input);
        ranges.sort();
        println!("before combining range count: {}", ranges.len());
        while combine_first_overlapping(&mut ranges) {
            // some combination happened which maybe could have messed up the sorting.
            ranges.sort();
        }
        println!("after combining range count: {}", ranges.len());

        let mut fresh = 0;
        for range in ranges {
            fresh += range.max+1-range.min
        }
        fresh
    }
}

fn parse_input(input: String) -> (Vec<Range>, Vec<i64>){
    let mut ranges: Vec<Range> = vec![];
    let mut ingredients: Vec<i64> = vec![];
    for line in input.split("\n") {
        if line == "" {
            continue // delimiter between ranges and ingredients
        }
        if line.contains("-") {
            let range: Vec<&str> = line.split("-").collect();
            let min = range[0].parse::<i64>().unwrap();
            let max = range[1].parse::<i64>().unwrap();
            ranges.push(Range { min, max })
        } else {
            ingredients.push(line.parse::<i64>().unwrap());
        }
    }
    (ranges, ingredients)
}