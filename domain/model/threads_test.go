package model_test

import (
	"onion-architecrure-go/domain/model"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestThreadsModel_NewThreadsModel(t *testing.T) {
	Convey("New Instance", t, func() {
		threadsModel := model.Threads{}
		testObj := model.NewThreadsModel()

		So(testObj, ShouldEqual, threadsModel)
	})
}
