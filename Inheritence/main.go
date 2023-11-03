package main

import "fmt"

type Hero struct {
	PlayerController interface{}
	AttributeCharacter
}

// func (h *Hero) SpecialMovePos(x, y int) {
// 	h.posX -= x
// 	h.posY -= y
// }

type EnemyBot struct {
	BotController interface{}
	AttributeCharacter
}

type AttributeCharacter struct {
	Position
	Att    int
	Health int
	Speed  int
}

type Position struct {
	posX int
	posY int
}

func (p *Position) ToZeroPos() {
	p.posX = 0
	p.posY = 0
}

func (p *Position) MovePos(x, y int) {
	if x == 0 && y == 0 {
		fmt.Println("hero not moving")
		return
	}

	p.posX += x
	p.posY += y
}

func (p *Position) SpecialMovePos(x, y int) *Position {
	if x == 0 && y == 0 {
		fmt.Println("hero not moving")
		return p
	}

	p.posX += x * x
	p.posY += y * y
	return p
}

func main() {
	// init hero
	hero_1 := Hero{}
	fmt.Println("Posisi ke - 1 : ", hero_1.Position)
	hero_1.Position.MovePos(4, 4)
	fmt.Println("Posisi ke - 2 : ", hero_1.Position)
	// hero_1.ToZeroPos()
	// fmt.Println("Posisi ke - 3 : ", hero_1.Position)
	hero_1.AttributeCharacter.Position.
		SpecialMovePos(2, 2). // => 4
		SpecialMovePos(1, 1). // 1
		SpecialMovePos(1, 1)  // 1
	fmt.Println("Posisi ke - 3 : ", hero_1.Position)
}
