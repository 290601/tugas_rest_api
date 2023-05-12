package games

type Games struct {
	Id         int64  `gorm:"primaryKey;autoIncrement;unsigned" json:"id"`
	Nama       string `gorm:"type:varchar(255)" json:"nama"`
	Genre      string `gorm:"type:varchar(255)" json:"genre"`
	Publisher  string `gorm:"type:varchar(255)" json:"publisher"`
	TahunRilis string `gorm:"type:varchar(255)" json:"tahun_rilis"`
	Platform   string `gorm:"type:varchar(255)" json:"platform"`
}

func (Games) TableName() string {
	return "games"
}
