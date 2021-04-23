package enigma

import (
	"bytes"
	"strings"
)

type Enigma struct {
	RotorSet  *RotorController
	Reflector *Reflector
	PlugBoard *PlugBoard
}

func NewEnigmaBoard(rc *RotorController, rf *Reflector, pb *PlugBoard) *Enigma {
	return &Enigma{
		RotorSet:  rc,
		Reflector: rf,
		PlugBoard: pb,
	}
}

func (e *Enigma) cipherLetter(letter rune) (rune, error) {
	// turn rotor
	res := e.PlugBoard.Reverse(letter)
	e.RotorSet.Turn()
	res = e.RotorSet.Process(res)
	res = e.Reflector.Reflect(res)
	res = e.RotorSet.ProcessReflected(res)
	res = e.PlugBoard.Reverse(res)

	return res, nil
}

func (e *Enigma) CipherText(text string) (string, error) {
	var b bytes.Buffer

	for _, c := range strings.ToLower(text) {
		newC, err := e.cipherLetter(c)
		if err != nil {
			return "", err
		}
		if _, err := b.WriteRune(newC); err != nil {
			return "", err
		}
	}

	return b.String(), nil
}
