package entity_test

import (
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/extend"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUsersEntity_NewUsersEntity(t *testing.T) {
	Convey("New Instance", t, func() {
		usersEntity := entity.Users{}
		testObj := entity.NewUsersEntity()

		So(testObj, ShouldEqual, usersEntity)
	})
}

func TestUsersEntity_SetHashPassword(t *testing.T) {
	Convey("密碼雜湊", t, func() {
		str := "test"

		hashStr := extend.Helper.Hash(str)
		testStr := entity.NewUsersEntity().SetHashPassword(str)

		So(testStr, ShouldEqual, hashStr)
	})
}
