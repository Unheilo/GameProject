package main

import (
	"fmt"
	"strings"
)

var RealPlayer *Player
var AllLocations []*Location
var InitialLocation *Location

type Location struct {
	Name         string
	MoveLocation string
	LookLocation string
	Object       []*Object
	Portal       []*Portal
}

type Portal struct {
	Name                string
	Obj                 *Object
	LocationSource      *Location
	LocationDestination *Location
}

type Object struct {
	Name          string
	Item          []*Item
	Dress         []*Dress
	Condition     bool
	ConditionItem *Item
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
	Item     []*Item
	Dress    []*Dress
	Location *Location
}

func NewItem(name string, neededBackpack bool) *Item {
	return &Item{
		Name:           name,
		NeededBackpack: neededBackpack,
	}
}

func NewDress(name string, backpack bool) *Dress {
	return &Dress{
		Name:     name,
		Backpack: backpack,
	}
}

func NewPlayer(name string, startLocation *Location) *Player {
	return &Player{
		Name:     name,
		Item:     []*Item{},
		Dress:    []*Dress{},
		Location: startLocation,
	}
}

func NewLocation(name string, looklocation string, movelocation string) *Location {
	return &Location{
		Name:         name,
		LookLocation: looklocation,
		MoveLocation: movelocation,
		Object:       []*Object{},
		Portal:       []*Portal{},
	}
}

func NewObject(name string, condition bool, conditionItem *Item) *Object {
	return &Object{
		Name:          name,
		Item:          []*Item{},
		Dress:         []*Dress{},
		Condition:     condition,
		ConditionItem: conditionItem,
	}
}

func NewPortal(name string, obj *Object, src, dst *Location) *Portal {
	return &Portal{
		Name:                name,
		Obj:                 obj,
		LocationSource:      src,
		LocationDestination: dst,
	}
}

// Look activity
func Look(player *Player, selector bool) string {
	env := LookEnvironment(player.Location, selector)
	portals := LookPortals(player.Location)

	result := env
	if env != "" && portals != "" {
		result += ". "
	}
	result += portals

	return result
}

func LookEnvironment(loc *Location, selector bool) string {

	var descriptions []string

	if !selector {
		descriptions = append(descriptions, loc.LookLocation)
		return strings.Join(descriptions, ", ")
	}
	if loc == InitialLocation {
		descriptions = append(descriptions, loc.LookLocation)
		//InitialLocation = nil
	}

	for _, obj := range loc.Object {

		if !obj.Condition {
			items := make([]string, len(obj.Item))

			for i, item := range obj.Item {
				items[i] = item.Name
			}

			dresses := make([]string, len(obj.Dress))
			for i, dress := range obj.Dress {
				dresses[i] = dress.Name
			}

			allThings := append(items, dresses...)

			if len(allThings) > 0 {
				descriptions = append(descriptions, fmt.Sprintf("%s: %s", obj.Name, strings.Join(allThings, ", ")))
			}
		}
	}

	return strings.Join(descriptions, ", ")
}

func LookPortals(loc *Location) string {
	if len(loc.Portal) == 0 {
		return ""
	}

	portalNames := make([]string, len(loc.Portal))
	for i, portal := range loc.Portal {
		portalNames[i] = portal.LocationDestination.Name
	}

	return "можно пройти - " + strings.Join(portalNames, ", ")
}

// Move activity
func Move(player *Player, locationName string) string {
	for _, portal := range player.Location.Portal {
		if portal.LocationDestination.Name == locationName {

			if portal.Obj != nil && portal.Obj.Condition {
				return "дверь закрыта"
			}

			player.Location = portal.LocationDestination
			return Look(player, false)
		}
	}
	return "нет пути в " + locationName
}

// PutOnDress activity
func PutOnDress(player *Player, dressName string) string {
	for _, obj := range player.Location.Object {
		for i, dress := range obj.Dress {
			if dress.Name == dressName {

				player.Dress = append(player.Dress, dress)
				obj.Dress = append(obj.Dress[:i], obj.Dress[i+1:]...)
				KitchenIntent(InitialLocation)
				return "вы надели: " + dressName
			}
		}
	}
	return "не удалось надеть " + dressName
}

func KitchenIntent(Location *Location) bool {
	for _, obj := range Location.Object {
		for i, item := range obj.Item {
			if item.Name == "надо собрать рюкзак и идти в универ" {
				obj.Item = append(obj.Item[:i], obj.Item[i+1:]...)
				obj.Item = append(obj.Item, NewItem("надо идти в универ", false))
				return true
			}
		}
	}
	return false
}

// TakeItem activity
func TakeItem(player *Player, itemName string) string {

	hasBackpack := false
	for _, dress := range player.Dress {
		if dress.Backpack {
			hasBackpack = true
			break
		}
	}

	if !hasBackpack {
		return "некуда класть"
	}

	for _, obj := range player.Location.Object {
		for i, item := range obj.Item {
			if item.Name == itemName {

				player.Item = append(player.Item, item)
				obj.Item = append(obj.Item[:i], obj.Item[i+1:]...)
				return "предмет добавлен в инвентарь: " + itemName
			}
		}
	}
	return "нет такого"
}

// Use activity
func UseItem(player *Player, itemName, objectName string) string {
	// Ищем предмет у игрока

	var targetItem *Item
	for _, item := range player.Item {
		if item.Name == itemName {
			targetItem = item
			break
		}
	}

	if targetItem == nil {
		return "нет предмета в инвентаре - " + itemName
	}

	// Ищем объект в текущей локации
	var targetObject *Object
	for _, obj := range player.Location.Object {
		if obj.Name == objectName {
			targetObject = obj
			break
		}
	}
	if targetObject == nil {
		return "не к чему применить"
	}

	// Проверяем условие использования
	if targetObject.ConditionItem != nil && targetObject.ConditionItem.Name == targetItem.Name {
		targetObject.Condition = false
		return objectName + " открыта"
	}
	return "не к чему применить"
}

// TakeBreakfast activity
func TakeBreakfast() string {
	return "неизвестная команда"
}

// Инициализация игры с правильными связями
func initGame() {

	keys := NewItem("ключи", true)
	notes := NewItem("конспекты", true)
	tea := NewItem("чай", true)

	NoBackpackIntent := NewItem("надо собрать рюкзак и идти в универ", false)

	backpack := NewDress("рюкзак", true)

	wardrobe := NewObject("шкаф", true, nil)
	tableRoom := NewObject("на столе", false, nil)
	tableKitchen := NewObject("на столе", false, nil)
	chair := NewObject("на стуле", false, nil)
	door := NewObject("дверь", true, keys)

	tableRoom.Item = append(tableRoom.Item, keys, notes)
	chair.Dress = append(chair.Dress, backpack)
	tableKitchen.Item = append(tableKitchen.Item, tea, NoBackpackIntent)

	room := NewLocation("комната", "ты в своей комнате", "")
	hallway := NewLocation("коридор", "ничего интересного", "ничего интересного")
	kitchen := NewLocation("кухня", "ты находишься на кухне", "кухня, ничего интересного")
	street := NewLocation("улица", "на улице весна. можно пройти - домой", "жарковато")

	room.Object = append(room.Object, tableRoom, chair)
	hallway.Object = append(hallway.Object, door, wardrobe)
	kitchen.Object = append(kitchen.Object, tableKitchen)

	roomToHallway := NewPortal("дверь", nil, room, hallway)
	hallwayToRoom := NewPortal("дверь", nil, hallway, room)
	hallwayToKitchen := NewPortal("проход", nil, hallway, kitchen)
	kitchenToHallway := NewPortal("проход", nil, kitchen, hallway)
	hallwayToStreet := NewPortal("дверь", door, hallway, street)
	//streetToHallway := NewPortal("дверь", door, street, hallway)

	room.Portal = append(room.Portal, roomToHallway)
	hallway.Portal = append(hallway.Portal, hallwayToKitchen, hallwayToRoom, hallwayToStreet)
	kitchen.Portal = append(kitchen.Portal, kitchenToHallway)
	//street.Portal = append(street.Portal, streetToHallway)

	AllLocations = []*Location{room, hallway, kitchen, street}
	InitialLocation = kitchen
	RealPlayer = NewPlayer("студент", kitchen)

}

func handleCommand(command string) string {
	//parts := strings.Fields(command)
	parts := strings.Split(command, " ")
	if len(parts) == 0 {
		return "неизвестная команда"
	}

	switch parts[0] {
	case "осмотреться":
		return Look(RealPlayer, true)
	case "завтракать":
		return TakeBreakfast()
	case "надеть":
		if len(parts) < 2 {
			return "укажите, что надеть"
		}
		return PutOnDress(RealPlayer, parts[1])
	case "взять":
		if len(parts) < 2 {
			return "укажите, что взять"
		}
		return TakeItem(RealPlayer, parts[1])
	case "применить":
		if len(parts) < 3 {
			return "не хватает аргументов"
		}
		return UseItem(RealPlayer, parts[1], parts[2])
	case "идти":
		if len(parts) < 2 {
			return "куда идти?"
		}
		return Move(RealPlayer, parts[1])
	default:
		return "неизвестная команда"
	}
}

func main() {

	initGame()

	command1 := "идти коридор"
	command2 := "идти комната"
	command3 := "надеть рюкзак"

	fmt.Println(handleCommand(command1))
	fmt.Println(handleCommand(command2))
	fmt.Println(handleCommand(command3))
}
