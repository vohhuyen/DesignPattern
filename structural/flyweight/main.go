package main

import (
	"fmt"
)

// Step 1: Chia các trường thành intrinsic state và extrinsic state
// Intrinsic state: `color` (không thay đổi, được chia sẻ giữa các đối tượng)
// Extrinsic state: `lat`, `long` (khác biệt giữa các đối tượng)

// dress.go: Flyweight interface
// Flyweight interface
type Dress interface {
	getColor() string
}

// terroristDress.go: Concrete flyweight object
// TerroristDress: Concrete Flyweight Object
type TerroristDress struct {
	color string
}

func (t *TerroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}

// counterTerroristDress.go: Concrete flyweight object
// CounterTerroristDress: Concrete Flyweight Object
type CounterTerroristDress struct {
	color string
}

func (c *CounterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "green"}
}

// Step 4: Factory Class để quản lý pool các Flyweight Objects
// dressFactory.go: Flyweight factory
const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var dressFactorySingleInstance = &DressFactory{
	dressMap: make(map[string]Dress),
}

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	switch dressType {
	case TerroristDressType:
		d.dressMap[dressType] = newTerroristDress()
	case CounterTerroristDressType:
		d.dressMap[dressType] = newCounterTerroristDress()
	default:
		return nil, fmt.Errorf("wrong dress type passed")
	}

	return d.dressMap[dressType], nil
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

// Step 5: Context lưu trữ Extrinsic State
// player.go: Context
type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) setLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

// game.go: Client code
// Game: Client Code
type Game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *Game {
	return &Game{
		terrorists:        make([]*Player, 0),
		counterTerrorists: make([]*Player, 0),
	}
}

func (g *Game) addTerrorist() {
	player := newPlayer("T", TerroristDressType)
	g.terrorists = append(g.terrorists, player)
}

func (g *Game) addCounterTerrorist() {
	player := newPlayer("CT", CounterTerroristDressType)
	g.counterTerrorists = append(g.counterTerrorists, player)
}

// Step 3: Client sử dụng Flyweight thông qua Factory
// main.go: Client code
func main() {
	game := newGame()

	// Thêm Terrorist
	game.addTerrorist()
	game.addTerrorist()
	game.addTerrorist()

	// Thêm Counter-Terrorist
	game.addCounterTerrorist()
	game.addCounterTerrorist()

	// In thông tin các dress objects
	dressFactoryInstance := getDressFactorySingleInstance()
	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressType: %s, Color: %s\n", dressType, dress.getColor())
	}
}