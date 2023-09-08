package domain

type Greeter struct {
	Greeting string
}

func (g *Greeter) Greet() string {
	return g.Greeting + ", world!"
}
