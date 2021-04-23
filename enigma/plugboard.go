package enigma

import "fmt"

type PlugBoard struct {
	Configuration map[rune]rune
}

func checkValidConf(cfg map[rune]rune) error {
	// valid plugboard checking
	for k, v := range cfg {
		if vv, ok := cfg[v]; ok && vv != k {
			return fmt.Errorf("invalid plugboard configuration %c and %c are not ordered properly (have %c and %c)", k, v, vv, v)
		}
	}

	return nil
}

func NewPlugBoard() *PlugBoard {
	return &PlugBoard{}
}

func (pb *PlugBoard) SetConfig(cfg map[rune]rune) error {
	if err := checkValidConf(cfg); err != nil {
		return err
	}

	ncfg := make(map[rune]rune, len(cfg)*2)
	for k, v := range cfg {
		ncfg[k] = v
		ncfg[v] = k
	}

	pb.Configuration = ncfg
	return nil
}

func (pb *PlugBoard) Reverse(c rune) rune {
	if _, ok := pb.Configuration[c]; ok {
		return pb.Configuration[c]
	}

	return c
}
