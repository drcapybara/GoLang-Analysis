package main

import "fmt"

var (
	//Go retains the usual types of the C based languages:
	byte_example   uint8  = 255
	sixteen_bits   uint16 = 65535
	sixtyfour_bits uint64 = 0xFFFFFFFFFFFFFFFF
	a_boolean      bool   = false
	a_byte         byte   = 255
	zero_value     int
	rune1          rune = '♬'
)

var no_type = "no problem"

func main() {
	fmt.Printf("Type:  %T  Value:  %v\n", byte_example, byte_example)
	fmt.Printf("Type:  %T  Value:  %v\n", sixteen_bits, sixteen_bits)
	fmt.Printf("Type:  %T  Value:  %v\n", a_boolean, a_boolean)
	fmt.Printf("Type:  %T  Value:  %v\n", sixtyfour_bits, sixtyfour_bits)
	fmt.Printf("bytes are equal to uint8: %v\n", a_byte == byte_example)
	fmt.Printf("no value? initialize to zero: %v\n", zero_value)
	fmt.Printf("Runes display unique unicode values: %v\n", rune1)
	fmt.Printf("Go can also infer types: %T\n", no_type)
}
