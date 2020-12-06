package leeter

import (
	"fmt"
	"testing"
)

type Test struct {
	Type string
	Args string
	Want string
}

func TestDefaultLeeter(t *testing.T) {
	tests := []Test{
		{Type: "success", Args: "hey", Want: "h3y"},
		{Type: "success", Args: "This is me", Want: "7h15 15 m3"},
		{Type: "success", Args: "Leeeeeet word is fun!", Want: "13333337 w0rd 15 fun!"},
	}

	for _, tt := range tests {
		t.Run(tt.Type, func(t *testing.T) {
			got, err := DefaultLeeter(tt.Args)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.Want {
				t.Errorf("Test failed : expected->%s, autual->%s", tt.Want, got)
			}

			fmt.Printf("Test Success: expected->%s, autual->%s\n", tt.Want, got)
		})
	}
}

func TestLowercaseLeeter(t *testing.T) {
	tests := []Test{
		{Type: "success", Args: "hey", Want: "h3y"},
		{Type: "success", Args: "This is me", Want: "Th15 15 m3"},
		{Type: "success", Args: "Leeeeeet word IS fun!", Want: "L3333337 w0rd IS fun!"},
	}

	for _, tt := range tests {
		t.Run(tt.Type, func(t *testing.T) {
			got, err := LowerLeeter(tt.Args)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.Want {
				t.Errorf("Test failed : expected->%s, autual->%s", tt.Want, got)
			}

			fmt.Printf("Test Success: expected->%s, autual->%s\n", tt.Want, got)
		})
	}
}

func TestUppercaseLeeter(t *testing.T) {
	tests := []Test{
		{Type: "success", Args: "hey", Want: "hey"},
		{Type: "success", Args: "This IS me", Want: "7his 15 me"},
		{Type: "success", Args: "Leeeeeet word IS fun!", Want: "1eeeeeet word 15 fun!"},
	}

	for _, tt := range tests {
		t.Run(tt.Type, func(t *testing.T) {
			got, err := UpperLeeter(tt.Args)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.Want {
				t.Errorf("Test failed : expected->%s, autual->%s", tt.Want, got)
			}

			fmt.Printf("Test Success: expected->%s, autual->%s\n", tt.Want, got)
		})
	}
}
