package migration

type Sat struct {
	Id   uint `gorm:"primaryKey"`
	Name string
	Juh  string
	Slug string
	Gok  string
}

func (sat *Sat) TableName() string {
	return "sat"
}
