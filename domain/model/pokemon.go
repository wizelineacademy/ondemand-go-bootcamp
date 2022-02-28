package model

//import "time"

type Pokemon struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	//CreatedAt time.Time `json:"created_at"`
}

func (Pokemon) TableName() string {
	return "pokemon"
}
