package hex

import "fmt"

/*
Dump takes a byte slice and prints out the hex codes
it can be used in conjunction with ioutil.ReadFile
rom is
*/
func Hexdump(rom []byte, numBytes int) {
	for i, element := range rom {

		if (i % numBytes) == 0 {
			// hex with 8 digits. used to split into lines
			if i == 0 {
				fmt.Printf("%08x ", i)
			} else {
				fmt.Printf("\n%08x ", i)
			}
		}
		// hex with 2 digits.  used for opcodes
		fmt.Printf("%02x", element)
		if (i % numBytes) != numBytes-1 {
			// space between op codes
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}
