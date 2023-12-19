package nineteen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_Example(t *testing.T) {
	input := []string{
		"px{a<2006:qkq,m>2090:A,rfg}",
		"pv{a>1716:R,A}",
		"lnx{m>1548:A,A}",
		"rfg{s<537:gd,x>2440:R,A}",
		"qs{s>3448:A,lnx}",
		"qkq{x<1416:A,crn}",
		"crn{x>2662:A,R}",
		"in{s<1351:px,qqz}",
		"qqz{s>2770:qs,m<1801:hdj,R}",
		"gd{a>3333:R,R}",
		"hdj{m>838:A,pv}",
		"",
		"{x=787,m=2655,a=1222,s=2876}",
		"{x=1679,m=44,a=2067,s=496}",
		"{x=2036,m=264,a=79,s=2244}",
		"{x=2461,m=1339,a=466,s=291}",
		"{x=2127,m=1623,a=2188,s=1013}",
	}
	assert.Equal(t, 19114, Part1(input))
}

func TestPart1_Single(t *testing.T) {
	input := []string{
		"px{a<2006:qkq,m>2090:A,rfg}",
		"pv{a>1716:R,A}",
		"lnx{m>1548:A,A}",
		"rfg{s<537:gd,x>2440:R,A}",
		"qs{s>3448:A,lnx}",
		"qkq{x<1416:A,crn}",
		"crn{x>2662:A,R}",
		"in{s<1351:px,qqz}",
		"qqz{s>2770:qs,m<1801:hdj,R}",
		"gd{a>3333:R,R}",
		"hdj{m>838:A,pv}",
		"",
		"{x=787,m=2655,a=1222,s=2876}",
	}
	assert.Equal(t, 7540, Part1(input))
}

func TestSplitRange(t *testing.T) {
	input := []string{
		"in{s<1351:px,qqz}",
		"",
		"{x=787,m=2655,a=1222,s=2876}",
	}
	workflows, _ := parseInput(input)
	inWorkflow := workflows["in"]
	rule := inWorkflow.rules[0]
	starterRanges := starterRanges(4000)

	matched, nonMatched := rule.splitRanges(starterRanges)
	assert.Equal(t, 1, matched.rs['s'].min)
	assert.Equal(t, 1351, matched.rs['s'].max)
	assert.Equal(t, 0, matched.getCount()%1350)
	assert.Equal(t, 1351, nonMatched.rs['s'].min)
	assert.Equal(t, 4001, nonMatched.rs['s'].max)
}

func TestPart2_Example(t *testing.T) {
	input := []string{
		"px{a<2006:qkq,m>2090:A,rfg}",
		"pv{a>1716:R,A}",
		"lnx{m>1548:A,A}",
		"rfg{s<537:gd,x>2440:R,A}",
		"qs{s>3448:A,lnx}",
		"qkq{x<1416:A,crn}",
		"crn{x>2662:A,R}",
		"in{s<1351:px,qqz}",
		"qqz{s>2770:qs,m<1801:hdj,R}",
		"gd{a>3333:R,R}",
		"hdj{m>838:A,pv}",
		"",
	}
	assert.Equal(t, 167409079868000, Part2(input))
}

func TestRanges_GetCount(t *testing.T) {
	ranges := starterRanges(10)
	assert.Equal(t, 10000, ranges.getCount())
}

func TestPart2_Single(t *testing.T) {
	input := []string{
		"in{s<1351:hdj,A}",
		"hdj{R}",
		"",
	}
	assert.Equal(t, 2650*4000*4000*4000, Part2(input))
}

func TestPart2_Single2(t *testing.T) {
	input := []string{
		"in{s>1350:px,qqz}",
		"px{R}",
		"qqz{s>2770:abc,m<1801:px,A}",
		"",
	}
	assert.Equal(t, 1350*4000*4000*2200, Part2(input))
}

func TestPart2_ComparePart1(t *testing.T) {
	input := []string{
		"px{a<2006:qkq,m>2090:A,rfg}",
		"pv{a>1716:R,A}",
		"lnx{m>1548:A,A}",
		"rfg{s<537:gd,x>2440:R,A}",
		"qs{s>3448:A,lnx}",
		"qkq{x<1416:A,crn}",
		"crn{x>2662:A,R}",
		"in{s<1351:px,qqz}",
		"qqz{s>2770:qs,m<1801:hdj,R}",
		"gd{a>3333:R,R}",
		"hdj{m>838:A,pv}",
		"",
		"{x=787,m=2655,a=1222,s=2876}",
		"{x=1679,m=44,a=2067,s=496}", // wrong
		"{x=2036,m=264,a=79,s=2244}",
		"{x=2461,m=1339,a=466,s=291}", // wrong
		"{x=2127,m=1623,a=2188,s=1013}",
	}

	workflows, prs := parseInput(input)
	for _, pr := range prs {
		toComp := 0
		if evaluatePartRating(workflows, pr) {
			toComp = 1
		}
		ans := parameterizedPart2(input, pr.getRanges())
		if toComp != ans {
			fmt.Println("Invalid answer for pr", pr)
		}
		assert.Equal(t, toComp, ans)
	}
}

func TestPart2_Single3(t *testing.T) {
	input := []string{
		"in{m<1:R,A}",
	}
	assert.Equal(t, 4000*4000*4000*4000, Part2(input))
}

func TestPart2_Single4(t *testing.T) {
	input := []string{
		"px{a<2006:qkq,m>2090:A,rfg}",
		"pv{a>1716:R,A}",
		"lnx{m>1548:A,A}",
		"rfg{s<537:gd,x>2440:R,A}",
		"qs{s>3448:A,lnx}",
		"qkq{x<1416:A,crn}",
		"crn{x>2662:A,R}",
		"in{s<1351:px,qqz}",
		"qqz{s>2770:qs,m<1801:hdj,R}",
		"gd{a>3333:R,R}",
		"hdj{m>838:A,pv}",
		"",
		"{x=1679,m=44,a=2067,s=496}", // wrong
	}
	// ok.. for some reason, when it gets to px, it is matching qkq, which is not correct.
	p1Ans := Part1(input)
	if p1Ans != 0 {
		p1Ans /= p1Ans
	}
	_, prs := parseInput(input)
	for _, pr := range prs {
		ans := parameterizedPart2(input, pr.getRanges())
		if ans != p1Ans {
			fmt.Println("Invalid answer for pr", pr)
		}
		assert.Equal(t, p1Ans, ans)
	}
}
