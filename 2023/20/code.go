package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/20/input-user.txt", false)
}

const BroadcasterLabel = "broadcaster"

type ModuleBase struct {
	label        string
	destinations []string
}

type Module interface {
	Send(event Event) *Pulse
	GetDestinations() []string
}

type Broadcaster struct {
	ModuleBase
}

func (b *Broadcaster) GetDestinations() []string {
	return b.destinations
}

func (b *Broadcaster) Send(event Event) *Pulse {
	return &event.pulse
}

type Flipflop struct {
	ModuleBase
	on bool
}

func (f *Flipflop) GetDestinations() []string {
	return f.destinations
}

func (f *Flipflop) Send(event Event) *Pulse {
	if event.pulse {
		return nil
	}
	if f.on {
		f.on = false
		return &Low
	} else {
		f.on = true
		return &High
	}
}

type Conjunction struct {
	ModuleBase
	memory map[string]Pulse
}

func (c *Conjunction) AddInput(label string) {
	c.memory[label] = false
}

func (c *Conjunction) GetDestinations() []string {
	return c.destinations
}

func (c *Conjunction) Send(event Event) *Pulse {
	c.memory[event.sender] = event.pulse

	allHigh := true
	for _, pulse := range c.memory {
		if !pulse {
			allHigh = false
		}
	}
	outPulse := Pulse(!allHigh)
	return &outPulse
}

type Pulse bool

var (
	Low  Pulse = false
	High Pulse = true
)

func (p Pulse) String() string {
	if p {
		return "high"
	} else {
		return "low"
	}
}

type Event struct {
	pulse               Pulse
	sender, destination string
}

func (e Event) String() string {
	return fmt.Sprintf("%s -%s-> %s", e.sender, e.pulse, e.destination)
}

func pushButton(modules map[string]Module) (int, int) {
	eventBus := []Event{{Low, "button", BroadcasterLabel}}
	lowPulses, highPulses := 0, 0
	for len(eventBus) > 0 {
		event := eventBus[0]
		//fmt.Println(event)
		var module Module
		var ok bool
		var outPulse *Pulse
		if event.pulse {
			highPulses++
		} else {
			lowPulses++
		}
		if module, ok = modules[event.destination]; ok {
			outPulse = module.Send(event)
		}
		if outPulse != nil {
			for _, dest := range module.GetDestinations() {
				eventBus = append(eventBus, Event{*outPulse, event.destination, dest})
			}
		}
		eventBus = eventBus[1:]
	}
	return lowPulses, highPulses
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	modules := map[string]Module{}
	for _, line := range util.Lines(input) {
		lineParts := strings.Split(line, " -> ")
		moduleLabel := lineParts[0]
		destinations := strings.Split(lineParts[1], ", ")

		firstChar := moduleLabel[0]
		if firstChar == '%' || firstChar == '&' {
			moduleLabel = moduleLabel[1:]
		}
		base := ModuleBase{moduleLabel, destinations}

		conjunctions := map[string]*Conjunction{}
		if firstChar == '%' {
			modules[moduleLabel] = &Flipflop{base, false}
		} else if firstChar == '&' {
			c := &Conjunction{base, map[string]Pulse{}}
			conjunctions[moduleLabel] = c
			modules[moduleLabel] = c
		} else if moduleLabel == BroadcasterLabel {
			modules[moduleLabel] = &Broadcaster{base}
		} else {
			panic("unexpected module!")
		}
	}

	for label, module := range modules {
		for _, dest := range module.GetDestinations() {
			switch modules[dest].(type) {
			case *Conjunction:
				modules[dest].(*Conjunction).AddInput(label)
			}
		}
	}

	lowPulses, highPulses := 0, 0
	for i := 0; i < 1000; i++ {
		low, high := pushButton(modules)
		//fmt.Println(low, high)
		lowPulses += low
		highPulses += high
	}
	//fmt.Println(pushButton(modules))
	// TODO: printing "4251 2749", but it should be "4250 2750"!
	fmt.Println(lowPulses, highPulses)
	return lowPulses * highPulses
}
