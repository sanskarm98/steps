package stair

// Stair represents a stair in the Endless Stairs game.
type Stair struct {
	Direction string // "left" or "right"
}

// NewStair creates a new Stair with the given direction.
func NewStair(direction string) *Stair {
	return &Stair{Direction: direction}
}

// LeftRender returns the stair ASCII art at the left offset.
func (s *Stair) LeftRender() []string {
	return []string{
		("       O  "),
		("      /|\\ "),
		("[___] / \\ "),
		("     [___]"),
	}
}

// RightRender returns the stair ASCII art at the right offset.
func (s *Stair) RightRender() []string {

	return []string{
		("      O        "),
		("     /|\\       "),
		("     / \\ [___]"),
		("    [___]      "),
	}
} 