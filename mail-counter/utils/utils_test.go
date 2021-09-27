package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestDetectSpam(t *testing.T) {
	file, err := os.Open("../TestCases.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	type Tests struct {
		name    string
		message string
		want    bool
	}
	type TestCase struct {
		Cases []Tests
	}
	tests := TestCase{}
	successCase := true
	for scanner.Scan() {
		message := scanner.Text()
		testCase := Tests{name: "", message: scanner.Text(), want: true}
		if message == "" {
			successCase = false
		}
		if !successCase {
			testCase = Tests{name: "", message: scanner.Text(), want: false}
		}
		tests.Cases = append(tests.Cases, testCase)
	}
	for _, tt := range tests.Cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectSpam(tt.message); got != tt.want {
				fmt.Printf("\n%s", tt.message)
				t.Errorf("DetectSpam() = %v, want %v", got, tt.want)
			}
		})
	}
}
