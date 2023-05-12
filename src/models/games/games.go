package games

type Games struct {
	Id         int64  `gorm:"primaryKey;autoIncrement;unsigned" json:"id"`
	Nama       string `gorm:"type:varchar(255);not null" json:"nama"`
	Genre      string `gorm:"type:varchar(255);not null" json:"genre"`
	Publisher  string `gorm:"type:varchar(255);not null" json:"publisher"`
	TahunRilis string `gorm:"type:varchar(255);not null" json:"tahun_rilis"`
	Platform   string `gorm:"type:varchar(255);not null" json:"platform"`
	Harga      int64  `gorm:"type:int;unsigned;not null" json:"harga"`
}

func (Games) TableName() string {
	return "games"
}
