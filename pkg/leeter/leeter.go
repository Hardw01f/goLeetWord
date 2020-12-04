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
		case 'o', 'O':
			leet = append(leet, c-(c-'0'))
		default:
			leet = append(leet, c)
		}
	}

	fmt.Printf("[ %s ] --LEET(default)--> : ", str)
	fmt.Println(string(leet))
}

func LowerLeeter(args []string) {

	str := args[0]

	for _, c := range str {
		switch c {
		case 'a':
			leet = append(leet, c-(c-'@'))
		case 'e':
			leet = append(leet, c-(c-'3'))
		case 'l':
			leet = append(leet, c-(c-'1'))
		case 'i':
			leet = append(leet, c-(c-'1'))
		case 's':
			leet = append(leet, c-(c-'5'))
		case 't':
			leet = append(leet, c-(c-'7'))
		case 'o':
			leet = append(leet, c-(c-'0'))
		default:
			leet = append(leet, c)
		}
	}

	fmt.Printf("[ %s ] --LEET(Only Lowercase)--> : ", str)
	fmt.Println(string(leet))
}

func UpperLeeter(args []string) {

	str := args[0]

	for _, c := range str {
		switch c {
		case 'A':
			leet = append(leet, c-(c-'@'))
		case 'E':
			leet = append(leet, c-(c-'3'))
		case 'L':
			leet = append(leet, c-(c-'1'))
		case 'I':
			leet = append(leet, c-(c-'1'))
		case 'S':
			leet = append(leet, c-(c-'5'))
		case 'T':
			leet = append(leet, c-(c-'7'))
		case 'O':
			leet = append(leet, c-(c-'0'))
		default:
			leet = append(leet, c)
		}
	}

	fmt.Printf("[ %s ] --LEET(Only Uppercase)--> : ", str)
	fmt.Println(string(leet))
}
