package models

import (
	"fmt"
	u "p3/utils"

	"github.com/jinzhu/gorm"
)

type Room struct {
	gorm.Model
	Name        string          `json:"name"`
	Category    string          `json:"category"`
	Desc        string          `json:"description"`
	Domain      string          `json:"domain"`
	Color       string          `json:"color"`
	Orientation ECardinalOrient `json:"eorientation"`
	Rack        []Rack          `gorm:"foreignKey:Rack"`
}

func (room *Room) Validate() (map[string]interface{}, bool) {
	if room.Name == "" {
		return u.Message(false, "Room Name should be on payload"), false
	}

	if room.Category == "" {
		return u.Message(false, "Category should be on the payload"), false
	}

	if room.Desc == "" {
		return u.Message(false, "Description should be on the paylad"), false
	}

	if room.Domain != "" {
		return u.Message(false, "Domain should NULL!"), false
	}

	if room.Color == "" {
		return u.Message(false, "Color should be on the payload"), false
	}

	switch room.Orientation {
	case "NE", "NW", "SE", "SW":
	case "":
		return u.Message(false, "Orientation should be on the payload"), false

	default:
		return u.Message(false, "Orientation is invalid!"), false
	}

	//Successfully validated Room
	return u.Message(true, "success"), true
}

func (room *Room) Create() map[string]interface{} {
	if resp, ok := room.Validate(); !ok {
		return resp
	}

	GetDB().Create(room)

	resp := u.Message(true, "success")
	resp["room"] = room
	return resp
}

//Get the first room given the room
func GetRoom(bldg *Building) *Room {
	room := &Room{}
	err := GetDB().Table("rooms").Where("foreignkey = ?", bldg.ID).First(room).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return room
}

//Obtain all rooms of a site
func GetRooms(bldg *Building) []*Room {
	rooms := make([]*Room, 0)

	err := GetDB().Table("rooms").Where("foreignkey = ?", bldg.ID).Find(&rooms).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return rooms
}

//More methods should be made to
//Meet CRUD capabilities
//Need Update and Delete
//These would be a bit more complicated
//So leave them out for now
