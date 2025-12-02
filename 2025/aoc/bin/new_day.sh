#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <day_number>"
    exit 1
fi

DAY=$(printf "%02d" $1)
DAY_DIR="src/days/day$DAY"

# Create day directory and files
mkdir -p "$DAY_DIR"
touch "$DAY_DIR/input.txt"
touch "$DAY_DIR/sample.txt"

# Create mod.rs
cat > "$DAY_DIR/mod.rs" << EOF
use crate::Day;

pub struct Day$DAY {}

impl Day for Day$DAY {
    fn solve_1(&self, input: String) -> i64 {
        todo!()
    }

    fn solve_2(&self, input: String) -> i64 {
        todo!()
    }
}
EOF

# Add to days/mod.rs on new line
echo "pub mod day$DAY;" >> src/days/mod.rs

# Find last Day import line number and insert after it
LAST_IMPORT=$(grep -n "use crate::days::day[0-9]*::Day[0-9]*;" src/main.rs | tail -1 | cut -d: -f1)
sed -i '' "${LAST_IMPORT}a\\
use crate::days::day$DAY::Day$DAY;
" src/main.rs

# Find last match case line number and insert after it
LAST_CASE=$(grep -n "^        [0-9]* => Box::new(Day[0-9]* {})," src/main.rs | tail -1 | cut -d: -f1)
sed -i '' "${LAST_CASE}a\\
        $1 => Box::new(Day$DAY {}),
" src/main.rs

echo "Created boilerplate for Day $1"
