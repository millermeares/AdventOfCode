use crate::Day;

pub struct Day01 {}

struct Turn {
    l: bool,
    dist: i32,
    raw: String
}

struct Dial {
    pos: i32,
}

impl Dial {
    fn apply(&mut self, t: &Turn) {
        let prev = self.pos;
        let mut m: i32 = 1;
        if t.l {
            m *= -1;
        }
        self.pos = self.pos + (m * t.dist);
    }
}

impl Day for Day01 {
    fn solve_1(&self, input: String) -> i32 {
        let mut at_zero = 0;
        let turns = parse_turns(input);
        let mut dial = Dial { pos: 50 };
        for turn in turns {
            dial.apply(&turn);
            if dial.pos % 100 == 0 {
                at_zero += 1;
            }
        }
        at_zero
    }

    fn solve_2(&self, input: String) -> i32 {
        let mut crossed_zero = 0;
        let turns = parse_turns(input);
        let mut dial = Dial { pos: 50 };
        for turn in turns {
            let prev = dial.pos;
            dial.apply(&turn);
            let new = dial.pos;
            // if turn was left, the smaller number is 'new'. 
            let amt: i32 = if turn.l {
                // multiply by negative one due to the `multiples_of_100_between_two_numbers` only being inclusive on the larger side. 
                // this works around a deficiency on the `multiples_of_100_between_two_numbers` not counting correctly when turning left.
                multiples_of_100_between_two_numbers(prev * -1, new * -1)
            } else {
                 multiples_of_100_between_two_numbers(prev, new)
            };
            crossed_zero += amt;
        }
        crossed_zero
    }
}

/*
 * Counts multiples of 100. Only inclusive on the 'l' side.
 */
fn multiples_of_100_between_two_numbers(s: i32, l: i32) -> i32 {
    if s >= l {
        panic!("l should always be bigger than s.");
    }
    let mut d = s;
    let mut i = 0;
    while l - d > 100 {
        d += 100;
        i += 1;
    }

    
    // rust division truncates.
    let s_div = d.div_euclid(100);
    let l_div = l.div_euclid(100);
    if s_div != l_div {
        i += 1;
    }
    i
}

fn parse_turns(input: String) -> Vec<Turn> {
    let mut turns: Vec<Turn> = vec![];

    let spl = input.split("\n");
    spl.for_each(|t| {
        // println!("{}", t[1..]);
        turns.push(Turn {
            l: (t.starts_with("L")),
            dist: t[1..].parse::<i32>().unwrap(),
            raw: t.to_string()
        })
    });
    turns
}

mod test {
    use crate::days::day01::multiples_of_100_between_two_numbers;

    #[test]
    fn test_multiples_between_two_numbers() {
        let n = multiples_of_100_between_two_numbers(1, 2);
        assert_eq!(n, 0);

        let one = multiples_of_100_between_two_numbers(-1, 1);
        assert_eq!(one, 1);

        let two = multiples_of_100_between_two_numbers(-97, 199);
        assert_eq!(two, 2);

        let even = multiples_of_100_between_two_numbers(0, 100);
        assert_eq!(even, 1); 

        // if i start at zero, i don't count that.
        let start_zero = multiples_of_100_between_two_numbers(0, 2);
        assert_eq!(0, start_zero);

        // if i end at zero, i count that. 
        let end_zero = multiples_of_100_between_two_numbers(98, 100);
        assert_eq!(1, end_zero);


        let one_neg = multiples_of_100_between_two_numbers(-1, 0);
        assert_eq!(1, one_neg);

        let one_pos = multiples_of_100_between_two_numbers(0, 1);
        assert_eq!(0, one_pos);

        let ten_ex = multiples_of_100_between_two_numbers(50, 1000);
        assert_eq!(10, ten_ex);
    }
}
