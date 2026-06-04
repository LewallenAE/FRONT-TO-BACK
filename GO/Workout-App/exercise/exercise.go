package exercise

import "fmt"

// create player struct name and inventory
// inventory is a slice of item
// create item struct name and type
// Methods
// PickUpItem: add item to player inventory
// DropItem: Item is deleted from player inventory
// UseItem: Depending on the item - it will print different things.

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name      string
	Inventory map[string]Item
}

func (p *Player) PickUpItem(newItem Item) {
	p.Inventory[newItem.Name] = newItem
	fmt.Printf("%s picked up %s!\n", p.Name, newItem.Name)
}

func (p *Player) DropItem(itemName string) {
	_, exists := p.Inventory[itemName]

	if exists {
		delete(p.Inventory, itemName)
		fmt.Printf("%s threw %s item away.\n", p.Name, itemName)
		return
	} else {
		fmt.Printf("%s, does not have %s to throw away!\n", p.Name, itemName)
		return
	}
}

func (p *Player) UseItem(itemName string) {
	fmt.Printf("%s uses %s\n", p.Name, itemName)
}
