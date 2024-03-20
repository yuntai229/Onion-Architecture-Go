package model_test

import (
	"onion-architecrure-go/domain/model"
	"onion-architecrure-go/extend"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUsersModel_NewUsersModel(t *testing.T) {
	Convey("New Instance", t, func() {
		usersModel := model.Users{}
		testObj := model.NewUsersModel()

		So(testObj, ShouldEqual, usersModel)
	})
}

func TestUsersModel_SetHashPassword(t *testing.T) {
	Convey("密碼雜湊", t, func() {
		str := "test"

		hashStr := extend.Helper.Hash(str)
		testStr := model.NewUsersModel().SetHashPassword(str)

		So(testStr, ShouldEqual, hashStr)
	})
}
