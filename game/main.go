package main

//structs

type Location struct {
	Name   string
	Object []Object
	Portal []*Portal
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
	Dress     []Dress
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
	Dress     []Dress
	Location  Location
}

//entity constructors

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

func NewPlayer(name string, Inventory []Item, Dress []Dress, Location Location) *Player {
	return &Player{
		Name:      name,
		Inventory: Inventory,
		Dress:     Dress,
		Location:  Location,
	}
}

func NewLocation(name string, Object []Object) *Location {
	return &Location{
		Name:   name,
		Object: Object,
		Portal: make([]*Portal, 0),
	}
}

func NewObject(name string, Item []Item, Dress []Dress, Condition bool) *Object {
	return &Object{
		Name:      name,
		Item:      Item,
		Dress:     Dress,
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
	InitGame()
}

// Look activity
func Look(Player *Player) string {
	var InitialString string

	EnvironmentString := LookEnvironment(&Player.Location)

	PortalsString := LookPortals(&Player.Location)

	InitialString += EnvironmentString + PortalsString

	return InitialString
}

func LookEnvironment(Location *Location) string {
	var InitialString string

	for _, value := range Location.Object {
		if !value.Condition {
			InitialString += value.Name + " "
			for _, value2 := range value.Item {
				InitialString += value2.Name + ", "
			}

			for _, value3 := range value.Dress {
				InitialString += value3.Name + ", "
			}
		}
	}

	return InitialString
}

func LookPortals(Location *Location) string {
	var InitialString string
	InitialString += "можно пройти - "

	for _, value := range Location.Portal {
		InitialString += value.LocationDestination.Name + ", "
	}

	return InitialString
}

// Look end activity

func InitGame() {

	Keys := NewItem("ключи", true)
	//Phone := NewItem("телефон", true)
	Notes := NewItem("конспекты", true)
	Tea := NewItem("чай, надо собрать рюкзак и идти в универ", true)

	Backpack := NewDress("рюкзак", true)

	Wardrobe := NewObject("шкаф", []Item{}, []Dress{}, true)
	TableRoom := NewObject("на столе:", []Item{*Keys, *Notes}, []Dress{*Backpack}, false)
	TableKitchen := NewObject("на столе:", []Item{*Tea}, []Dress{}, false)
	Chair := NewObject("на стуле:", []Item{}, []Dress{}, false)
	Door := NewObject("дверь", []Item{*Keys}, []Dress{}, true)

	Room := NewLocation("ты в своей комнате", []Object{*TableRoom})
	Hallway := NewLocation("ничего интересного", []Object{*Door, *Wardrobe})
	Kitchen := NewLocation("ты находишься на кухне", []Object{*TableKitchen, *Chair})
	Street := NewLocation("на улице весна", []Object{})

	FromRoomToHallway := NewPortal("от комнаты к коридору", Object{}, *Room, *Hallway)
	FromHallwayToRoom := NewPortal("от коридора к комнате", Object{}, *Hallway, *Room)
	FromHallwayToKitchen := NewPortal("от коридора к кухне", Object{}, *Hallway, *Kitchen)
	FromKitchenToHallway := NewPortal("от кухни к коридору", Object{}, *Kitchen, *Hallway)
	FromHallwayToStreet := NewPortal("от коридора к улице", *Door, *Hallway, *Street)
	FromStreetToHallway := NewPortal("от улицы к коридору", *Door, *Street, *Hallway)

	Room.Portal = append(Room.Portal, FromRoomToHallway)
	Hallway.Portal = append(Hallway.Portal, FromHallwayToRoom, FromHallwayToKitchen, FromHallwayToStreet)
	Kitchen.Portal = append(Kitchen.Portal, FromKitchenToHallway)
	Street.Portal = append(Street.Portal, FromStreetToHallway)

	student := NewPlayer("студент", []Item{}, []Dress{}, *Kitchen)

}
