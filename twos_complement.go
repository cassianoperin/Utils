package main

import "fmt"

func GetTwoComplement(num int8) string {
        var ret = []byte("00000000")

        for i := range ret {
                if ((1 << uint(i)) & uint(num)) > 0 {
                        ret[i] = '1'
                }
        }
        beg := 0
        end := len(ret) - 1
        for beg <= end {
                ret[beg], ret[end] = ret[end], ret[beg]
                beg++
                end--
        }
        return string(ret)
}

func DecodeTwoComplement(num byte) int8 {

	var sum int8 = 0

	for i := 0 ; i < 8 ; i++ {
		// Sum each bit and sum the value of the bit power of i (<<i)
		sum += (int8(num) >> i & 0x01) << i
	}
	fmt.Printf("\n")

	return sum
}

func main() {

var (
	value	int
)
	   value = 0xAA
	   fmt.Printf("Hex: %d \t\tTwo's complement %s\n", value, GetTwoComplement(int8(value)))
	   fmt.Printf("Hex: %d \t\tDecode Two's complement %d\n", value, DecodeTwoComplement(byte(value)))


}
