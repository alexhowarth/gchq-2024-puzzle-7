package main

import "fmt"

var charMap = map[rune]rune{
	'A': 'B',
	'B': 'I',
	'C': 'M',
	'D': 'N',
	'E': 'D',
	'F': 'O',
	'G': 'H',
	'H': '?',
	'I': 'P',
	'J': '?',
	'K': 'R',
	'L': 'C',
	'M': 'S',
	'N': 'T',
	'O': 'U',
	'P': 'V',
	'Q': 'L',
	'R': '?',
	'S': 'E',
	'T': '?',
	'U': 'G',
	'V': '?',
	'W': 'A',
	'X': 'Y',
	'Y': '?',
	'Z': '?',
	' ': ' ',
}

func main() {
	fmt.Println(string(decipher("ISKGWIM KSWEBDU BN DFN LBIGSKSE BM IKBCWKBQX W QSWEBDU BDEBLWNBFD NF XFO MFQPBDU NGBM WECBNNSEQX ESPBFOM BDBNBWQQX ODKSWEWAQS CSMMWUS")))
}

func decipher(in string) (deciphered []rune) {
	for _, char := range in {
		deciphered = append(deciphered, charMap[char])
	}
	return
}
