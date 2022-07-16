package booking

// Public Exports

// All fields inside an exported struct needs to be exported
type Conf struct {
	Name    string
	Tickets uint8
}

func Init(n string, c uint8) Conf {
	// Usage of Structs for maintaining data
	return Conf{
		Name:    n,
		Tickets: c,
	}
}
