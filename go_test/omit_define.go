package go_test

import "time"

type UserInfo struct {
	ID         uint64    `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`              // 自增主键id
	UID        int64     `gorm:"unique;column:uid;type:bigint(20);not null;default:0"`                // 用户 id
	IsOpenGame uint8     `gorm:"column:is_open_game;type:tinyint(4) unsigned;not null;default:0"`     // 是否开启游戏
	Coin       uint32    `gorm:"column:coin;type:int(10) unsigned;not null;default:0"`                // 用户金币
	Assets     uint32    `gorm:"column:assets;type:int(10) unsigned;not null;default:0"`              // 用户资产
	GarageLen  uint32    `gorm:"column:garage_len;type:int(10) unsigned;not null;default:0"`          // 用户已有车辆数
	ParkLen    uint32    `gorm:"column:park_len;type:int(10) unsigned;not null;default:4"`            // 用户停车场解锁车位数
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP"` // 更新时间
}

func OmitTime(u UserInfo) UserInfo {
	return UserInfo{
		ID:         u.ID,
		UID:        u.UID,
		IsOpenGame: u.IsOpenGame,
		Coin:       u.Coin,
		Assets:     u.Assets,
		GarageLen:  u.GarageLen,
		ParkLen:    u.ParkLen,
	}

}
