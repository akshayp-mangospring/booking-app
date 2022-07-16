package booking

// Public Exports

// All fields inside an exported struct needs to be exported
type Conf struct {
	Tickets uint8
	Name    string
}

func Init(c uint8, n string) Conf {
	// Usage of Structs for maintaining data
	return Conf{
		Tickets: c,
		Name:    n,
	}
}
