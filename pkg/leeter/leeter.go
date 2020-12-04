package leeter

import (
	"fmt"
)

var leet []rune

func DefaultLeeter(args []string) {

	str := args[0]

	for _, c := range str {
		switch c {
		case 'a', 'A':
			leet = append(leet, c-(c-'@'))
		case 'e', 'E':
			leet = append(leet, c-(c-'3'))
		case 'l', 'L':
			leet = append(leet, c-(c-'1'))
		case 'i', 'I':
			leet = append(leet, c-(c-'1'))
		case 's', 'S':
			leet = append(leet, c-(c-'5'))
		case 't', 'T':
			leet = append(leet, c-(c-'7'))
		default:
			leet = append(leet, c)
		}
	}

	fmt.Printf("[ %s ] --LEET(default)--> : ", str)
	fmt.Println(string(leet))
}
