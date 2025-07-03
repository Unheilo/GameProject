package main

type Location struct {
	Name   string
	Object []Object
}

type Portal struct {
	Name                string
	Object              Object
	LocationSource      Location
	LocationDestination Location
}

type Object struct {
	Name      string
	Item      []Item
	Condition bool
}

type Item struct {
	Name           string
	NeededBackpack bool
}

type Dress struct {
	Name     string
	Backpack bool
}

type Player struct {
	Name      string
	Inventory []Item
}

func NewItem(name string, NeededBackpack bool) *Item {
	return &Item{
		Name:           name,
		NeededBackpack: NeededBackpack,
	}
}

func NewDress(name string, Backpack bool) *Dress {
	return &Dress{
		Name:     name,
		Backpack: Backpack,
	}
}

func NewPlayer(name string, Inventory []Item) *Player {
	return &Player{
		Name:      name,
		Inventory: Inventory,
	}
}

func NewLocation(name string, Object []Object) *Location {
	return &Location{
		Name:   name,
		Object: Object,
	}
}

func NewObject(name string, Item []Item, Condition bool) *Object {
	return &Object{
		Name:      name,
		Item:      Item,
		Condition: Condition,
	}
}

func NewPortal(name string, Object Object, LocationSource Location, LocationDestination Location) *Portal {
	return &Portal{
		Name:                name,
		Object:              Object,
		LocationSource:      LocationSource,
		LocationDestination: LocationDestination,
	}
}

func main() {

	Keys := NewItem("ключи", true)
	Phone := NewItem("телефон", true)
	Notes := NewItem("конспекты", true)

	Vardrobe := NewObject("шкаф", []Item{}, false)
	Table := NewObject("стол", []Item{}, false)
	Chair := NewObject("стул", []Item{}, false)
	Door := NewObject("дверь", []Item{*Keys}, true)

	Backpack := NewDress("рюкзак", true)

	Stas := NewPlayer("стас", []Item{})

	Room := NewLocation("омната", []Object{})
	Hallway := NewLocation("коридор", []Object{})
	Kitchen := NewLocation("кухня", []Object{*Table, *Chair})
	Street := NewLocation("улица", []Object{*Vardrobe})

	FromRoomToHallway := NewPortal("от комнаты к коридору", Object{}, *Room, *Hallway)
	FromHallwayToRoom := NewPortal("от коридора к комнате", Object{}, *Hallway, *Room)
	FromHallwayToKitchen := NewPortal("от коридора к кухне", Object{}, *Hallway, *Kitchen)
	FromKitchenToHallway := NewPortal("от кухни к коридору", Object{}, *Kitchen, *Hallway)
	FromHallwayToStreet := NewPortal("от коридора к улице", *Door, *Hallway, *Street)
	FromStreetToHallway := NewPortal("от улицы к коридору", *Door, *Street, *Hallway)

}
