package enigma

import "fmt"

type RotorController struct {
	Rotors []*Rotor
}

type Rotor struct {
	Position               int
	Configuration          []int
	ConfigurationReflected []int
}

func NewRotorController(r []*Rotor) *RotorController {
	return &RotorController{
		Rotors: r,
	}
}

func (rc *RotorController) Turn() {
	turnRotor := true
	for _, rotor := range rc.Rotors {
		if turnRotor {
			turnRotor = rotor.Turn()
		}
	}
}

func (rc *RotorController) Process(c rune) rune {
	res := c

	for i := int(0); i < len(rc.Rotors); i++ {
		res = rc.Rotors[i].Process(res)
	}

	return res
}

func (rc *RotorController) ProcessReflected(c rune) rune {
	res := c
	for i := len(rc.Rotors) - 1; i >= 0; i-- {
		res = rc.Rotors[i].ProcessReflected(res)
	}

	return res
}

func NewRotor(pos int, cfg []rune) *Rotor {
	ncfg := make([]int, len(cfg))
	for i := int(0); i < len(cfg); i++ {
		ncfg[i] = int(cfg[i] - 'a')
	}
	cfgReflected := make([]int, len(cfg))
	for i := int(0); i < len(cfg); i++ {
		cfgReflected[ncfg[i]] = i
	}

	return &Rotor{
		Position:               pos,
		Configuration:          ncfg,
		ConfigurationReflected: cfgReflected,
	}
}

func (r *Rotor) SetPosition(pos int) error {
	if pos >= len(r.Configuration) {
		return fmt.Errorf("Invalid position")
	}
	r.Position = pos
	return nil
}

func (r *Rotor) ProcessReflected(c rune) rune {
	val := int(c - 'a')
	return 'a' + rune((len(r.ConfigurationReflected)+r.ConfigurationReflected[(len(r.ConfigurationReflected)+val-r.Position)%len(r.ConfigurationReflected)]-r.Position)%len(r.ConfigurationReflected))
}

func (r *Rotor) Process(c rune) rune {
	val := int(c - 'a')
	return 'a' + rune((r.Configuration[(val+r.Position)%len(r.Configuration)]+r.Position)%len(r.Configuration))
}

func (r *Rotor) Turn() bool {
	prev := r.Position
	r.Position = (r.Position + 1) % len(r.Configuration)
	// will return whether the rotor made a full turn or not
	return r.Position < prev
}
