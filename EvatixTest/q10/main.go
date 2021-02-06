package main

import (
	"bufio"
	"fmt"
	"os"
)

func cipher(text string, direction int, k int) string {

	shift := rune(k)
	offset := rune(26)
	rn := []rune(text)
	for i, ch := range rn {

		switch direction {
		// encoding
		case +1:
			if ch >= 'A' && ch <= 'Z'-shift || ch >= 'a' && ch <= 'z'-shift {
				ch = ch + shift
			} else if ch > 'Z'-shift && ch <= 'Z' ||
				ch > 'z'-shift && ch <= 'z' {
				ch = ch + shift - offset
			}
		// decoding
		case -1:
			if ch >= 'A'+shift && ch <= 'Z' || ch >= 'a'+shift && ch <= 'z' {
				ch = ch - shift
			} else if ch >= 'A' && ch < 'A'+shift || ch >= 'a' && ch < 'a'+shift {
				ch = ch - shift + offset
			}
		}
		rn[i] = ch
	}

	return string(rn)
}

func encode(text string, k int) string {
	return cipher(text, +1, k)
}
func decode(text string, k int) string {
	return cipher(text, -1, k)
}

func main() {
	var k int
	fmt.Println("Enter K: (Shift value)")
	_, err := fmt.Scanf("%d", &k)
	if err != nil {
		fmt.Println("input number error: ", err)
	}

	fmt.Println("Enter String to encode and decode: ")
	// fmt.Scanln(&s)
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("there was an error")
	}
	encoded := encode(line, k)
	println("  encoded: " + encoded)
	decoded := decode(encoded, k)
	println("  decoded: " + decoded)
}
