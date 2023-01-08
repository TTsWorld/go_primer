package go_orm

import (
	"fmt"
	"go_primer/go_test"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var uid1, uid2, uid3, uid4, uid5, uid6 int64 = 1508688644892925952, 1496731025781432320, 885876700272762881, 1481465403430543360, 100000020, 30013
var db *gorm.DB

func init() {
	dsn := "root:1qaz1QAZ2wsx2WSX@tcp(43.134.6.226:3306)/mico_parking?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

}

func TestUser(t *testing.T) {
	userCases := []go_test.UserInfo{
		{ID: 1, UID: uid1, IsOpenGame: 1, Coin: 10000000, Assets: 1494, GarageLen: 2, ParkLen: 2},
		{ID: 2, UID: uid2, IsOpenGame: 1, Coin: 10000000, Assets: 1494, GarageLen: 2, ParkLen: 2},
		{ID: 3, UID: uid3, IsOpenGame: 1, Coin: 10000000, Assets: 1494, GarageLen: 2, ParkLen: 2},
		{ID: 4, UID: uid4, IsOpenGame: 1, Coin: 10000000, Assets: 1494, GarageLen: 2, ParkLen: 2},
		{ID: 5, UID: uid5, IsOpenGame: 1, Coin: 10000000, Assets: 1494, GarageLen: 2, ParkLen: 2},
		{ID: 6, UID: uid6, IsOpenGame: 1, Coin: 10000000, Assets: 1494, GarageLen: 2, ParkLen: 2},
	}
	userList := []go_test.UserInfo{}
	uids := []int64{uid1, uid2, uid3, uid4, uid5, uid6}
	if err := db.Table("user_info").Model(&go_test.UserInfo{}).Where("uid IN ?", uids).Find(&userList).Error; err != nil {
		assert.Error(t, err)
	}
	assert.True(t, cmp.Equal(userCases[0], userList[0], cmp.Transformer("omitUser", go_test.OmitTime)),
		fmt.Sprintf("%v != %v", userCases[0], userList[0]))

}

func TestFind(t *testing.T) {
	rows, err := db.Debug().Table("user_info").Select("uid").Rows()
	fmt.Sprintf("%+v, %+v", rows, err)
	for rows.Next() {

	}
}
