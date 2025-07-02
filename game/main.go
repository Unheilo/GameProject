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

// declare items
var Keys = Item{Name: "ключи", NeededBackpack: true}
var Phone = Item{Name: "телефон", NeededBackpack: true}
var Notes = Item{Name: "конспекты", NeededBackpack: true}

// declare dress
var Backpack = Dress{Name: "рюкзак", Backpack: true}

// declare player
var Stas = Player{Name: "стас"}

// declare locations
var Room = Location{Name: "комната"}
var Hallway = Location{Name: "коридор", Object: []Object{door}}
var Kitchen = Location{Name: "кухня", Object: []Object{table, chair}}
var Street = Location{Name: "улица"}

// declare objects
var vardrobe = Object{Name: "шкаф"} //засуну его на улицу, шкаф в кустах
var table = Object{Name: "стол"}
var chair = Object{Name: "стул"}
var door = Object{Name: "дверь", Item: []Item{Keys}, Condition: true}

// declare portals
var FromRoomToHallway = Portal{Name: "от комнаты к коридору", LocationSource: Room, LocationDestination: Hallway}
var FromHallwayToRoom = Portal{Name: "от коридора к комнате", LocationSource: Hallway, LocationDestination: Room}
var FromHallwayToKitchen = Portal{Name: "от коридора к кухне", LocationSource: Hallway, LocationDestination: Kitchen}
var FromKitchenToHallway = Portal{Name: "от кухни к коридору", LocationSource: Kitchen, LocationDestination: Hallway}
var FromHallwayToStreet = Portal{Name: "от коридора к улице", Object: door, LocationSource: Hallway, LocationDestination: Street}
var FromStreetToHallway = Portal{Name: "от улицы к коридору", Object: door, LocationSource: Street, LocationDestination: Hallway}

func main() {
	InitItems()
}
