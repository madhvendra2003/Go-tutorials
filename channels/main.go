package main

import (
	"fmt"
	"testing"
	"time"
)

func TestProcessMessages(t *testing.T) {
	tests := []struct {
		messages []string
		expect   []string
	}{
		{
			messages: []string{"Sunlit", "Man"},
			expect:   []string{"Man-processed", "Sunlit-processed"},
		},
		{
			messages: []string{"Nomad do you copy?"},
			expect:   []string{"Nomad do you copy?-processed"},
		},
		{
			messages: []string{"Scadriel", "Roshar", "Sel", "Nalthis", "Taldain"},
			expect:   []string{"Taldain-processed", "Roshar-processed", "Sel-processed", "Nalthis-processed", "Scadriel-processed"},
		},
	}

	if withSubmit {
		tests = append(tests,
			struct {
				messages []string
				expect   []string
			}{
				messages: []string{},
				expect:   []string{},
			},
			struct {
				messages []string
				expect   []string
			}{
				messages: []string{"Scadriel"},
				expect:   []string{"Scadriel-processed"},
			},
		)
	}

	passCount := 0
	failCount := 0

	for _, tc := range tests {
		fail := false
		result := processMessages(tc.messages)

		if len(result) != len(tc.expect) {
			fail = true
		}

		counts := make(map[string]int)
		for _, res := range result {
			counts[res]++
		}
		for _, exp := range tc.expect {
			counts[exp]--
			if counts[exp] < 0 {
				fail = true
			}
		}

		if fail {
			failCount++
			t.Errorf(`
---------------------------------
Test Failed:
  inputs:   %v
  expected: %v
  actual:   %v
  `, tc.messages, tc.expect, result)
		} else {
			passCount++
			fmt.Printf(`
---------------------------------
Test Passed:
  inputs:   %v
  expected: %v
  actual:   %v
`, tc.messages, tc.expect, result)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true

func main() {
	fmt.Println("Running tests...")
	TestProcessMessages(&testing.T{})
}

func processMessages(messages []string) []string {
	// Result slice to collect processed messages
	ans := make([]string, 0, len(messages))
	// Channel to collect processed messages
	ch := make(chan string, len(messages))

	// Start goroutines for each message
	for _, value := range messages {
		go func(value string) {
			ch <- process(value) // send the processed message to the channel
		}(value)
	}

	time.Sleep(3 * time.Second)

	// Collect the results from the channel
	for i := 0; i < len(messages); i++ {
		ans = append(ans, <-ch)
	}

	return ans

}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
