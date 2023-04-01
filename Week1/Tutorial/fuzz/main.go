package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)

	// Below is go blog explaining slice
	// https://go.dev/blog/slices
	fmt.Print("--------------------------\n")

	// Create a couple of starter slices.
	slice := []int{1, 2, 3}
	slice2 := []int{55, 66, 77}
	fmt.Println("Start slice: ", slice)
	fmt.Println("Start slice2:", slice2)

	// Add an item to a slice.
	slice = append(slice, 4)
	fmt.Println("Add one item:", slice)

	// Add one slice to another.
	slice = append(slice, slice2...)
	fmt.Println("Add one slice:", slice)

	// Make a copy of a slice (of int).
	slice3 := append([]int(nil), slice...)
	fmt.Println("Copy a slice:", slice3)

	// Copy a slice to the end of itself.
	fmt.Println("Before append to self:", slice)
	slice = append(slice, slice...)
	fmt.Println("After append to self:", slice)

	// Below is go blog explaining strings
	// https://go.dev/blog/strings
	fmt.Print("--------------------------\n")

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)

	fmt.Printf("%x\n", sample)
	fmt.Printf("%+q\n", sample)

	const nihongo = "日本語"
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}
