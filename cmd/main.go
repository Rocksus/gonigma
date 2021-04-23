package main

import (
	"fmt"
	"os"

	"github.com/Rocksus/gonigma/enigma"
)

func main() {
	pos1 := 7
	pos2 := 6
	pos3 := 13
	pos4 := 14
	pos5 := 3
	rotor1 := enigma.NewRotor(pos1, []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'})
	rotor2 := enigma.NewRotor(pos2, []rune{'r', 'n', 'e', 'l', 'u', 'b', 's', 't', 'm', 'y', 'f', 'o', 'g', 'k', 'd', 'w', 'j', 'v', 'q', 'p', 'i', 'c', 'x', 'z', 'h', 'a'})
	rotor3 := enigma.NewRotor(pos3, []rune{'z', 'y', 'x', 'w', 'v', 'u', 't', 's', 'r', 'q', 'p', 'o', 'n', 'm', 'l', 'k', 'j', 'i', 'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a'})
	rotor4 := enigma.NewRotor(pos4, []rune{'r', 'n', 'e', 'l', 'u', 'b', 's', 't', 'm', 'y', 'f', 'o', 'g', 'k', 'd', 'w', 'j', 'v', 'q', 'p', 'i', 'c', 'x', 'z', 'h', 'a'})
	rotor5 := enigma.NewRotor(pos5, []rune{'r', 'n', 'e', 'l', 'u', 'b', 's', 't', 'm', 'y', 'f', 'o', 'g', 'k', 'd', 'w', 'j', 'v', 'q', 'p', 'i', 'c', 'x', 'z', 'h', 'a'})
	rotorSet := enigma.NewRotorController([]*enigma.Rotor{
		rotor1,
		rotor2,
		rotor3,
		rotor4,
		rotor5,
	})

	reflector := enigma.NewReflector(map[rune]rune{
		'a': 'n',
		'b': 'o',
		'c': 'p',
		'd': 'q',
		'e': 'r',
		'f': 's',
		'g': 't',
		'h': 'u',
		'i': 'v',
		'j': 'w',
		'k': 'x',
		'l': 'y',
		'm': 'z',
	})

	plugboard := enigma.NewPlugBoard()
	if err := plugboard.SetConfig(map[rune]rune{
		'a': 'b',
		'd': 'f',
		'z': 'x',
	}); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	board := enigma.NewEnigmaBoard(rotorSet, reflector, plugboard)

	encrypted, err := board.CipherText("thisisanexamplestatement")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("Encrypted: ", encrypted)

	// Reset position
	rotor1.SetPosition(pos1)
	rotor2.SetPosition(pos2)
	rotor3.SetPosition(pos3)
	rotor4.SetPosition(pos4)
	rotor5.SetPosition(pos5)
	decrypted, err := board.CipherText(encrypted)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("Decrypted: ", decrypted)
}
