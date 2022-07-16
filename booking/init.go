package booking

// Public Exports

func Init(n string, c uint8) Conf {
	// Usage of Structs for maintaining data
	return Conf{
		Name:    n,
		Tickets: c,
	}
}
