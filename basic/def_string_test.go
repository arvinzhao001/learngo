package basic

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

//A string is a slice of bytes in Go, are UTF-8 Encoded
func TestPrintBytes(t *testing.T) {
	message := "Hello World"
	printBytes(message)
}

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // for hexadecimal
		fmt.Printf("%c ", s[i]) // for charactor
	}
}

//In UTF-8 encoding a code point can occupy more than 1 byte
func TestPrintCharactor(t *testing.T) {
	message := "Hello World : 你好世界"
	printChars(message)
	charsAndBytePosition(message)
}

//iterate charactor method 1
func printChars(s string) {
	fmt.Printf("Characters: ")
	/*
		string is converted to a slice of runes.
		A rune is a builtin type in GO and it's the alias  of int32
		It's doesn't matter how many bytes the code points occupies,
		It can be represented by a rune
	*/
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

//iterate charactor method 2
func charsAndBytePosition(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func TestCreateString(t *testing.T) {
	//from a slices of bytes hex values
	byteSlice1 := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str1 := string(byteSlice1)
	fmt.Println(str1)
	//from a slices of bytes decimal values
	byteSlice2 := []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str2 := string(byteSlice2)
	fmt.Println(str2)
	//from a slices of rune
	runeSlice3 := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str3 := string(runeSlice3)
	fmt.Println(str3)
}

func TestStrLen(t *testing.T) {
	word1 := "Señor"
	fmt.Printf("String: %s\n", word1)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1)) // the number of charactor
	fmt.Printf("Number of bytes: %d\n", len(word1))           // the number of byte
}

func TestModifyString(t *testing.T) {
	str1 := "hello"
	//any valid Unicode character within a single quote is a rune
	//string is immutable, this command fails to compile with errors
	//str[1] = '0'

	//string are converted to a slice of runes, a new string is created
	str2 := []rune(str1)
	str2[1] = 'E'
	str3 := string(str2)

	fmt.Println(str1) //hello
	fmt.Println(str3) //hEllo
}
