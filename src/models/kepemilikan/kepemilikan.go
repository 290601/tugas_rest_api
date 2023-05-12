package kepemilikan

import (
	"tugas-rest-api/models/games"
	"tugas-rest-api/models/users"
)

type Kepemilikan struct {
	IdUser         int64       `gorm:"unsigned;not null" json:"id_user"`
	IdGame         int64       `gorm:"unsigned;not null" json:"id_game"`
	WaktuPembelian string      `gorm:"type:timestamp;not null" json:"waktu_pembelian"`
	Game           games.Games `gorm:"foreignKey:IdGame;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User           users.Users `gorm:"foreignKey:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Kepemilikan) TableName() string {
	return "kepemilikan"
}
