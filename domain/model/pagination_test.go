package model_test

import (
	"onion-architecrure-go/domain/model"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPaginationModel_NewDbPaginationScope(t *testing.T) {
	Convey("設定搜索範圍", t, func() {
		pagination := model.Pagination{
			Page:     2,
			PageSize: 10,
		}
		scope := pagination.NewDbPaginationScope()
		So(reflect.TypeOf(scope).Kind(), ShouldEqual, reflect.Func)
	})
}

func TestPaginationModel_ComposeOrderQuery(t *testing.T) {
	Convey("設定資料排序方式", t, func() {
		pagination := model.Pagination{
			OrderBy: "id",
			Sort:    "asc",
		}

		testQuery := "id asc"
		query := pagination.ComposeOrderQuery()

		So(query, ShouldEqual, testQuery)
	})
}

func TestPaginationModel_CalculatePage(t *testing.T) {
	Convey("計算搜索數量", t, func() {
		pagination := model.Pagination{
			PageSize: 10,
		}

		var count int64 = 101
		pagination.CalculatePage(count)

		So(pagination.LastPage, ShouldEqual, 11)
		So(pagination.Total, ShouldEqual, count)
	})
}

func TestPaginationModel_Format(t *testing.T) {
	Convey("結構化搜索結果", t, func() {
		type testStruct struct {
			val int
		}

		testData := testStruct{3}
		pagination := model.Pagination{
			Page:     1,
			PageSize: 1,
			LastPage: 1,
			OrderBy:  "id",
			Sort:     "asc",
		}

		pageRes := model.PageResponse{
			Meta: model.Pagination{
				Page:     pagination.Page,
				PageSize: pagination.PageSize,
				Total:    pagination.Total,
				LastPage: pagination.LastPage,
				OrderBy:  pagination.OrderBy,
				Sort:     pagination.Sort,
			},
			Collection: testData,
		}
		testRes := pagination.Format(testData)

		So(testRes, ShouldEqual, pageRes)
	})
}
