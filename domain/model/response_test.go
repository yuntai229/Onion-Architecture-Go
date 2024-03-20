package model_test

import (
	"onion-architecrure-go/domain/model"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResponseModel_NewResModel(t *testing.T) {
	Convey("New Instance", t, func() {
		resModel := model.ResponseModel{}
		testObj := model.NewResModel()

		So(testObj, ShouldEqual, resModel)
	})
}

func TestResponseModel_ResWithSucc(t *testing.T) {
	Convey("響應成功", t, func() {

		Convey("data payload - nil", func() {
			res := model.ResSucc{
				Code:    "0000",
				Message: "Succ",
				Data:    nil,
			}
			testRes := model.NewResModel().ResWithSucc(nil)

			So(testRes, ShouldEqual, res)
		})
		Convey("data payload - struct", func() {
			type testStruct struct {
				val int
			}
			testData := testStruct{3}
			res := model.ResSucc{
				Code:    "0000",
				Message: "Succ",
				Data:    testData,
			}

			testRes := model.NewResModel().ResWithSucc(testData)

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
			res := model.ResSucc{
				Code:    "0000",
				Message: "Succ",
				Data:    testData,
			}
			testRes := model.NewResModel().ResWithSucc(testData)
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
			res := model.ResSucc{
				Code:    "0000",
				Message: "Succ",
				Data:    testData,
			}
			testRes := model.NewResModel().ResWithSucc(testData)
			So(testRes, ShouldEqual, res)
		})

	})
}

func TestResponseModel_ResWithFail(t *testing.T) {
	Convey("響應失敗", t, func() {
		err := model.NotFoundErr
		res := model.ResFail{
			Code:    err.Code,
			Message: err.Message,
		}

		testRes := model.NewResModel().ResWithFail(err)

		So(testRes, ShouldEqual, res)
	})
}
