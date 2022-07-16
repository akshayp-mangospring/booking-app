package booking

// Public Exports

// All fields inside an exported struct needs to be exported
type Conf struct {
	Name    string
	Tickets uint8
	names   []string
}

// Attach a method to a struct
// These are also called as Receiver functions
func (c *Conf) GetTicketBookers() []string {
	return c.names
}
