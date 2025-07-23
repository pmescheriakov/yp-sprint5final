package personaldata

import "fmt"

type Personal struct {
	Name           string
	Weight, Height float64
}

// Print - just info about user
func (p Personal) Print() {
	fmt.Printf("Имя: %v\nВес: %.2f кг.\nРост: %.2f м.\n", p.Name, p.Weight, p.Height)
}
