package users

type Users struct {
	Id       int64  `gorm:"primaryKey;autoIncrement;unsigned" json:"id"`
	Nama     string `gorm:"type:varchar(255);not null" json:"nama"`
	Email    string `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

func (Users) TableName() string {
	return "user"
}
