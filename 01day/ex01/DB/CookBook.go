package DB

import "encoding/xml"

type Ingredients struct {
	Name string  `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit string  `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
  }

  type Cake struct {
	Name string               `json:"name" xml:"name"`
	Time string               `json:"time" xml:"stovetime"`
	Ingredients []Ingredients `json:"ingredients" xml:"ingredients>item"`
  }

  type CookBook struct {
	XMLName xml.Name          `json:"-" xml:"recipes"`
	Cakes []Cake              `json:"cake" xml:"cake"`
  }
