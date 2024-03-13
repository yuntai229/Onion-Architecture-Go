package entity_test

import (
	"onion-architecrure-go/domain/entity"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestThreadsEntity_NewThreadsEntity(t *testing.T) {
	Convey("New Instance", t, func() {
		threadsEntity := entity.Threads{}
		testObj := entity.NewThreadsEntity()

		So(testObj, ShouldEqual, threadsEntity)
	})
}
