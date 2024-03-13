package entity_test

import (
	"onion-architecrure-go/domain/entity"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResponseEntity_NewResEntity(t *testing.T) {
	Convey("New Instance", t, func() {
		resEntity := entity.ResponseEntity{}
		testObj := entity.NewResEntity()

		So(testObj, ShouldEqual, resEntity)
	})
}

func TestResponseEntity_ResWithSucc(t *testing.T) {
	Convey("響應成功", t, func() {

		Convey("data payload - nil", func() {
			res := entity.ResSucc{
				Code:    "0000",
				Message: "Succ",
				Data:    nil,
			}
			testRes := entity.NewResEntity().ResWithSucc(nil)

			So(testRes, ShouldEqual, res)
		})
		Convey("data payload - struct", func() {
			type testStruct struct {
				val int
			}
			testData := testStruct{3}
			res := entity.ResSucc{
				Code:    "0000",
				Message: "Succ",
				Data:    testData,
			}

			testRes := entity.NewResEntity().ResWithSucc(testData)

			So(testRes, ShouldEqual, res)
		})
		Convey("data payload - nest struct", func() {
			type subTestStruct struct {
				val int
			}
			type testStruct struct {
				val           int
				subTestStruct subTestStruct
			}

			testData := testStruct{
				val: 3,
				subTestStruct: subTestStruct{
					val: 4,
				},
			}
			res := entity.ResSucc{
				Code:    "0000",
				Message: "Succ",
				Data:    testData,
			}
			testRes := entity.NewResEntity().ResWithSucc(testData)
			So(testRes, ShouldEqual, res)
		})
		Convey("data payload - array", func() {
			type testStruct struct {
				val int
			}
			var testData []testStruct

			for i := 0; i < 10; i++ {
				testData = append(testData, testStruct{i})
			}
			res := entity.ResSucc{
				Code:    "0000",
				Message: "Succ",
				Data:    testData,
			}
			testRes := entity.NewResEntity().ResWithSucc(testData)
			So(testRes, ShouldEqual, res)
		})

	})
}

func TestResponseEntity_ResWithFail(t *testing.T) {
	Convey("響應失敗", t, func() {
		err := entity.NotFoundErr
		res := entity.ResFail{
			Code:    err.Code,
			Message: err.Message,
		}

		testRes := entity.NewResEntity().ResWithFail(err)

		So(testRes, ShouldEqual, res)
	})
}
