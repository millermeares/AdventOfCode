package twenty

import (
	"days"
	"fmt"
	"strings"
)

func GetDay() days.Day {
	return days.MakeDay(Part1, Part2, "20")
}

func Part1(input []string) int {
	modules := parseInput(input)
	lowPulses := 0
	highPulses := 0
	for i := 0; i < 1000; i++ {
		low, high, _ := pushButton(modules)
		fmt.Println(low, high)
		lowPulses += low
		highPulses += high
	}
	return lowPulses * highPulses
}

func pushButton(modules map[string]Module) (int, int, int) {
	lowPulseCount := 0
	highPulseCount := 0

	lowPulseToRx := 0
	buttonPulse := Pulse{isHigh: false, destination: "broadcaster", source: "button"}
	queue := []Pulse{buttonPulse}
	for len(queue) > 0 {
		pulse := queue[0]
		queue = queue[1:]
		if pulse.isHigh {
			highPulseCount++
		} else {
			lowPulseCount++
			if pulse.destination == "rx" {
				lowPulseToRx++
			}
		}
		resultingPulses := modules[pulse.destination].processPulse(pulse)
		queue = append(queue, resultingPulses...)
	}
	return lowPulseCount, highPulseCount, lowPulseToRx
}

func Part2(input []string) int {
	modules := parseInput(input)
	ql := modules["ql"].(*ConjunctionModule)
	var iterationsUntilHigh []int
	for d := range ql.receivedPulseStrength {
		fmt.Println("Evaluating", d)
		iterationsUntilHigh = append(iterationsUntilHigh, firstInstanceOfQlHavingHigh(input, d))
	}
	fmt.Println("Iterations until high for", ql.receivedPulseStrength, ":", iterationsUntilHigh)
	return LCM(iterationsUntilHigh[0], iterationsUntilHigh[1], iterationsUntilHigh[2:])
}

func firstInstanceOfQlHavingHigh(input []string, dest string) int {
	modules := parseInput(input)
	count := 0
	qlHasHigh := false
	for !qlHasHigh {
		count++
		pushButton(modules)
		// fmt.Println(ql.receivedPulseStrength, "while waiting for", dest)
		if pushButtonWaitForStrength(modules, dest) {
			qlHasHigh = true
		}
	}
	return count
}

func pushButtonWaitForStrength(modules map[string]Module, dest string) bool {
	buttonPulse := Pulse{isHigh: false, destination: "broadcaster", source: "button"}
	queue := []Pulse{buttonPulse}
	for len(queue) > 0 {
		pulse := queue[0]
		queue = queue[1:]
		if pulse.source == "ql" {
			qlMod := modules["ql"].(*ConjunctionModule)
			hasReceived := qlMod.receivedPulseStrength[dest]
			if hasReceived {
				return true
			}
		}
		resultingPulses := modules[pulse.destination].processPulse(pulse)
		queue = append(queue, resultingPulses...)
	}
	return false
}

type Module interface {
	getId() string
	getDestinations() []string
	processPulse(p Pulse) []Pulse
	setupSources() bool
}

type BroadcastModule struct {
	id           string
	destinations []string
}

func (bm *BroadcastModule) getId() string {
	return bm.id
}

func (bm *BroadcastModule) getDestinations() []string {
	return bm.destinations
}

func (bm *BroadcastModule) setupSources() bool {
	return false
}

func (bm *BroadcastModule) processPulse(p Pulse) []Pulse {
	return makePulses(bm.id, bm.destinations, p.isHigh)
}

// broadcast module is only used once and it receives a low pulse. make it a FlipFlop that starts with "on".
type FlipFlopModule struct {
	id           string
	destinations []string
	on           bool
}

func (fm *FlipFlopModule) getId() string {
	return fm.id
}

func (fm *FlipFlopModule) setupSources() bool {
	return false
}

func (fm *FlipFlopModule) getDestinations() []string {
	return fm.destinations
}

func (fm *FlipFlopModule) processPulse(p Pulse) []Pulse {
	if p.isHigh {
		return []Pulse{}
	}
	if fm.on {
		fm.on = false
		return makePulses(fm.id, fm.destinations, false)
	} else {
		fm.on = true
		return makePulses(fm.id, fm.destinations, true)
	}
}

type ConjunctionModule struct {
	id                    string
	destinations          []string
	receivedPulseStrength map[string]bool
}

func (cm *ConjunctionModule) getId() string {
	return cm.id
}

func (cm *ConjunctionModule) setupSources() bool {
	return true
}

func (cm *ConjunctionModule) getDestinations() []string {
	return cm.destinations
}

func (cm *ConjunctionModule) processPulse(p Pulse) []Pulse {
	cm.receivedPulseStrength[p.source] = p.isHigh
	if areAllValuesHigh(cm.receivedPulseStrength) {
		return makePulses(cm.id, cm.destinations, false)
	} else {
		return makePulses(cm.id, cm.destinations, true)
	}
}

type UntypedModule struct {
	id string
}

func (um *UntypedModule) processPulse(p Pulse) []Pulse {
	return []Pulse{}
}

func (um *UntypedModule) getId() string {
	return um.id
}

func (um *UntypedModule) setupSources() bool {
	return false
}

func (um *UntypedModule) getDestinations() []string {
	return []string{}
}

type Pulse struct {
	isHigh      bool
	source      string
	destination string
}

func makePulses(source string, destinations []string, isHigh bool) []Pulse {
	var pulses []Pulse
	for _, d := range destinations {
		pulses = append(pulses, Pulse{source: source, destination: d, isHigh: isHigh})
	}
	return pulses
}

func areAllValuesHigh(m map[string]bool) bool {
	for _, v := range m {
		if !v {
			return false
		}
	}
	return true
}

func parseInput(input []string) map[string]Module {
	modules := map[string]Module{}
	for _, line := range input {
		module := parseModule(line)
		modules[module.getId()] = module
	}
	for _, v := range modules {
		for _, d := range v.getDestinations() {
			_, e := modules[d]
			if !e {
				modules[d] = &UntypedModule{id: d}
			}
			if !modules[d].setupSources() {
				continue // module does not need to know who its source is.
			}
			// send a low pulse to initialize (really here for conjunction modules)
			modules[d].processPulse(Pulse{isHigh: false, source: v.getId(), destination: d})
		}
	}

	return modules
}

func parseModule(input string) Module {
	input = strings.Replace(input, ",", "", -1)
	fields := strings.Fields(input)
	destinations := fields[2:]
	id := fields[0]
	if id[0] == '%' || id[0] == '&' {
		id = id[1:]
	}

	if input[0] == '%' {
		return &FlipFlopModule{
			destinations: destinations,
			id:           id,
			on:           false, // initially off.
		}
	}
	if input[0] == '&' {
		return &ConjunctionModule{
			destinations:          destinations,
			id:                    id,
			receivedPulseStrength: map[string]bool{},
		}
	}
	return &BroadcastModule{
		id:           id,
		destinations: destinations,
	}
}

// copy-pasted from https://go.dev/play/p/SmzvkDjYlb and amended
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers []int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i], integers[i+1:])
	}
	return result
}
