package enigma

type Reflector struct {
	Configuration []rune
}

func NewReflector(cfg map[rune]rune) *Reflector {
	ncfg := make([]rune, len(cfg)*2)
	for k, v := range cfg {
		ncfg[int(k-'a')] = v
		ncfg[int(v-'a')] = k
	}
	return &Reflector{
		Configuration: ncfg,
	}
}

func (r *Reflector) Reflect(c rune) rune {
	return r.Configuration[int(c-'a')]
}
