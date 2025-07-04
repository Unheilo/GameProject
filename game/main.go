package main

import (
	"strings"
)

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
	Name          string
	Item          []Item
	Dress         []Dress
	Condition     bool
	ConditionItem Item
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
	Name     string
	Item     []Item
	Dress    []Dress
	Location Location
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

func NewPlayer(name string, Item []Item, Dress []Dress, Location Location) *Player {
	return &Player{
		Name:     name,
		Item:     Item,
		Dress:    Dress,
		Location: Location,
	}
}

func NewLocation(name string, Object []Object) *Location {
	return &Location{
		Name:   name,
		Object: Object,
		Portal: make([]*Portal, 0),
	}
}

func NewObject(name string, Item []Item, Dress []Dress, Condition bool, ConditionItem Item) *Object {
	return &Object{
		Name:          name,
		Item:          Item,
		Dress:         Dress,
		Condition:     Condition,
		ConditionItem: ConditionItem,
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

// Move activity
func Move(Player *Player, Location *Location) string {

	moveCheckResult := CheckMove(Player, Location)

	switch moveCheckResult {
	case "можно переместиться":
		Player.Location = *Location
		return "Ты переместился в " + Location.Name
	case "дверь закрыта":
		return "Дверь закрыта. Не можешь пройти в " + Location.Name
	case "нет пути":
		return "Нет пути в " + Location.Name
	default:
		return "Нет пути в " + Location.Name
	}
}

func CheckMove(Player *Player, Location *Location) string {

	for _, portal := range Player.Location.Portal {
		if &portal.LocationDestination == Location {
			if portal.Object.Condition {
				return "closed" // Дверь закрыта
			}
			return "open" // Дверь открыта, перемещение возможно
		}
	}
	return "noway" // Нет портала в направлении

}

// end move activity

// PutOnDress activity

func PutOnDress(Player *Player, Dress *Dress) string {
	for i, object := range Player.Location.Object {
		for j, valueDress := range object.Dress {
			if valueDress == *Dress {

				Player.Dress = append(Player.Dress, *Dress)
				object.Dress = append(object.Dress[:j], object.Dress[j+1:]...)

				Player.Location.Object[i] = object

				return "вы надели: " + Dress.Name
			}
		}
	}
	return "не удалось надеть " + Dress.Name
}

// end PutOnDress activity

// TakeBreakfast activity

func TakeBreakfast() string {
	return "неизвестная команда"
}

// end TakeBreakfast activity

//  TakeItem activity

func TakeItem(Player *Player, Item *Item) string {
	for i, object := range Player.Location.Object {
		for j, valueItem := range object.Item {
			if valueItem == *Item {

				if len(Player.Dress) > 0 {
					Player.Item = append(Player.Item, *Item)
					object.Item = append(object.Item[:j], object.Item[j+1:]...)

					Player.Location.Object[i] = object
					return "предмет добавлен в инвентарь: " + Item.Name
				} else {
					return "некуда класть"
				}
			}
		}
	}
	return "нет такого"
}

// end TakeItem activity

// Use activity

func UseItem(Player *Player, Item *Item, Object *Object) string {
	for _, playerItem := range Player.Item {
		if playerItem == *Item {
			if Object.ConditionItem == *Item {
				Object.Condition = false
				return Object.Name + " открыта"
			} else {
				return "не к чему применить"
			}
		}
	}

	return "нет предмета в инвентаре: " + Item.Name
}

// end use activity

func InitGame() {

	Keys := NewItem("ключи", true)
	//Phone := NewItem("телефон", true)
	Notes := NewItem("конспекты", true)
	Tea := NewItem("чай, надо собрать рюкзак и идти в универ", true)

	Backpack := NewDress("рюкзак", true)

	Wardrobe := NewObject("шкаф", []Item{}, []Dress{}, true, Item{})
	TableRoom := NewObject("на столе:", []Item{*Keys, *Notes}, []Dress{*Backpack}, false, Item{})
	TableKitchen := NewObject("на столе:", []Item{*Tea}, []Dress{}, false, Item{})
	Chair := NewObject("на стуле:", []Item{}, []Dress{}, false, Item{})
	Door := NewObject("дверь", []Item{*Keys}, []Dress{}, true, *Keys)

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

	Player := NewPlayer("студент", []Item{}, []Dress{}, *Kitchen)

}

func handleCommand(command string) string {
	SplittedString := strings.Split(command, " ")
	switch SplittedString[0] {
	case "осмотреться":
		return Look(Player)
	case "завтракать":
		return TakeBreakfast()
	case "надеть":
		return PutOnDress(Player, SplittedString[1])
	case "взять":
		return TakeItem(Player, SplittedString[1])
	case "применить":
		return UseItem(Player, SplittedString[1], SplittedString[2])
	case "идти":
		return Move(Player, SplittedString[1])
	default:
		return "неизвестная команда"

	}
}
