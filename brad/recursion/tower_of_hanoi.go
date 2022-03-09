package main

import "fmt"

func move_disk(from_pole, to_pole string) string {
	return fmt.Sprintf("move disk from %v to %v", from_pole, to_pole)
}

func move_tower(height int, from_pole, to_pole, with_pole string) {
	if height >= 1 {
		move_tower(height - 1, from_pole, with_pole, to_pole)
		fmt.Println(move_disk(from_pole, to_pole))
		move_tower(height - 1, with_pole, to_pole, from_pole)
	}
}

func main() {
	move_tower(3, "A", "B", "C")
}