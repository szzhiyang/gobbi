package generate

type Interfaces []*Interface

func (ii Interfaces) init(ns *Namespace) {
	for _, i := range ii {
		i.init(ns, "Interface")
	}
}

func (ii Interfaces) generate(f *file) {
	for _, i := range ii {
		i.generate(f)
	}
}

func (ii Interfaces) byName(name string) (*Interface, bool) {
	for _, iface := range ii {
		if iface.Name == name {
			return iface, true
		}
	}

	return nil, false
}