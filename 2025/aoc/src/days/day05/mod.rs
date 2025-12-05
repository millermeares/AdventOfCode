use crate::Day;

pub struct Day05 {}

struct Range {
    min: i64,
    max: i64
}

impl Range {
    fn in_range(&self, v: i64) -> bool {
        return v >= self.min && v <= self.max
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
        let (ranges, _) = parse_input(input);
        let (min, max) = get_min_max_of_ranges(&ranges);
        println!("{min}-{max}");
        let mut fresh = 0;

        // ok here is my idea. sort by minimum. merge all ranges that are overlapping. 
        // if there are no overlapping ranges, then just add up all of the numbers.
        for ingredient in min..max+1 {
            if ingredient % 1000000 == 0 {
                println!("million: {ingredient}");
            }
            if ingredient % 1000000000 == 0 {
                println!("billion: {ingredient}");
            }
            if ingredient % 1000000000000 == 0 {
                println!("trillion: {ingredient}");
            }
            if ingredient_in_a_range(ingredient, &ranges) {
                fresh += 1
            }   
        }
        fresh as i64
    }
}

fn get_min_max_of_ranges(ranges: &Vec<Range>) -> (i64, i64) {
    let mins = ranges.iter().map(|f| f.min);
    let maxs = ranges.iter().map(|f| f.max);
    return (mins.min().unwrap(), maxs.max().unwrap())

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