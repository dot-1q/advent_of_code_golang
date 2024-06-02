package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Event struct {
	date  time.Time // Time of the event
	event string    // Description of what happened.
}

type Minute struct {
	minute int         // Minute value. Ex: 1,2,3,...,54,55,56...
	sleep  map[int]int // Sleep occurrences on this minute from a given guard.
}

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	// Create the events and sort them
	events := createEvents(lines)
	slices.SortFunc(events, func(a, b Event) int {
		if a.date.Before(b.date) {
			return -1
		} else {
			return 1
		}
	})
	minutes := simulate(events)
	guard := sleepiestGuard(minutes)
	minute := sleepiestMinute(minutes, guard)
	fmt.Printf("Guard %d sleeps the most, and sleeps the most on minute %d\n", guard, minute)
	fmt.Println("Part 1 |  answer: ", guard*minute)
	guard, minute = asleepSameMinute(minutes)
	fmt.Println("Part 2 |  answer: ", guard*minute)

}

func simulate(events []Event) []Minute {
	minutes := createMinutes()
	guard := 0
	lastFallSleep := 0
	for _, event := range events {
		sep := strings.Fields(event.event)
		switch sep[0] {
		case "falls":
			minutes[event.date.Minute()].sleep[guard]++
			// Save the time we went to sleep
			lastFallSleep = event.date.Minute()
		case "wakes":
			now := event.date.Minute()
			for i := lastFallSleep + 1; i < now; i++ {
				// Fill the times from last fall asleep until he wakes up
				// as sleep time
				minutes[i].sleep[guard]++
			}
		case "Guard":
			guard, _ = strconv.Atoi(strings.TrimPrefix(sep[1], "#"))
		}
	}
	return minutes
}

func asleepSameMinute(minutes []Minute) (int, int) {
	maximum := 0
	guard := 0
	m := 0
	for _, minute := range minutes {
		for key, value := range minute.sleep {
			if value > maximum {
				maximum = value
				guard = key
				m = minute.minute
			}
		}
	}
	fmt.Printf("Guard %d sleeps the most on minute %d\n", guard, m)
	return guard, m
}

func sleepiestMinute(minutes []Minute, guard int) int {
	maximum := 0
	m := 0
	for _, minute := range minutes {
		for key, value := range minute.sleep {
			if key == guard {
				if value > maximum {
					maximum = value
					m = minute.minute
				}
			}
		}
	}
	return m
}

func sleepiestGuard(minutes []Minute) int {
	// For every guard, save the times he slept
	occurrences := map[int]int{}

	for _, minute := range minutes {
		for key, value := range minute.sleep {
			occurrences[key] += value
		}
	}
	maximum := 0
	sleepiest := 0
	for key, value := range occurrences {
		if value > maximum {
			maximum = value
			sleepiest = key
		}
	}
	return sleepiest
}

func createMinutes() []Minute {
	m := make([]Minute, 60)

	for i := range 60 {
		m[i] = Minute{i, map[int]int{}}
	}
	return m
}
func createEvents(lines []string) []Event {
	const shortForm = "2006-01-02 15:04"
	events := []Event{}

	for _, line := range lines {
		before, after, _ := strings.Cut(line, "]")
		t, _ := time.Parse(shortForm, strings.TrimPrefix(before, "["))
		// We don't know yet, since we haven't sorted the events
		event := Event{t, after}
		events = append(events, event)
	}
	return events
}
