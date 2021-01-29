package models

import (
	"fmt"
	u "p3/utils"

	"github.com/jinzhu/gorm"
)

type Vector2 struct {
	X float32
	Y float32
}

type Building struct {
	gorm.Model
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"description"`
	Domain   string `json:"domain"`

	Pos     Vector2 `json:"posxy"`
	PosU    string  `json:"posxyu"`
	PosZ    float32 `json:"posz"`
	PosZU   string  `json:"poszu"`
	Size    float32 `json:"size"`
	SizeU   string  `json:"sizeu"`
	Height  float32 `json:"height"`
	HeightU string  `json:"heightu"`
	Room    []Room  `gorm:"foreignKey:Room"`
}

func (bldg *Building) Validate() (map[string]interface{}, bool) {
	if bldg.Name == "" {
		return u.Message(false, "Building Name should be on payload"), false
	}

	if bldg.Category == "" {
		return u.Message(false, "Category should be on the payload"), false
	}

	if bldg.Desc == "" {
		return u.Message(false, "Description should be on the paylad"), false
	}

	if bldg.Domain != "" {
		return u.Message(false, "Domain should NULL!"), false
	}

	if bldg.Pos.X < 0.0 || bldg.Pos.Y < 0.0 {
		return u.Message(false, "Invalid XYcoordinates on payload"), false
	}

	if bldg.PosU == "" {
		return u.Message(false, "PositionXY string should be on the payload"), false
	}

	if bldg.PosZ < 0.0 {
		return u.Message(false, "Invalid Z coordinates on payload"), false
	}

	if bldg.PosZU == "" {
		return u.Message(false, "PositionZ string should be on the payload"), false
	}

	if bldg.Size <= 0.0 {
		return u.Message(false, "Invalid building size on the payload"), false
	}

	if bldg.SizeU == "" {
		return u.Message(false, "Building size string should be on the payload"), false
	}

	if bldg.Height <= 0.0 {
		return u.Message(false, "Invalid Height on payload"), false
	}

	if bldg.HeightU == "" {
		return u.Message(false, "Building Height string should be on the payload"), false
	}

	//Successfully validated bldg
	return u.Message(true, "success"), true
}

func (bldg *Building) Create() map[string]interface{} {
	if resp, ok := bldg.Validate(); !ok {
		return resp
	}

	GetDB().Create(bldg)

	resp := u.Message(true, "success")
	resp["building"] = bldg
	return resp
}

//Obtain the first building
//Given the site
func GetBuilding(site *Site) *Building {
	bldg := &Building{}
	err := GetDB().Table("buildings").Where("foreignkey = ?", site.ID).First(bldg).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return bldg
}

//Obtain all buildings of a site
func GetBuildings(site *Site) []*Building {
	bldgs := make([]*Building, 0)

	err := GetDB().Table("buildings").Where("foreignkey = ?", site.ID).Find(&bldgs).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return bldgs
}

//More methods should be made to
//Meet CRUD capabilities
//Need Update and Delete
//These would be a bit more complicated
//So leave them out for now
