package entity_test

import (
	"onion-architecrure-go/domain/entity"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPaginationEntity_NewDbPaginationScope(t *testing.T) {
	Convey("設定搜索範圍", t, func() {
		pagination := entity.Pagination{
			Page:     2,
			PageSize: 10,
		}
		scope := pagination.NewDbPaginationScope()
		So(reflect.TypeOf(scope).Kind(), ShouldEqual, reflect.Func)
	})
}

func TestPaginationEntity_ComposeOrderQuery(t *testing.T) {
	Convey("設定資料排序方式", t, func() {
		pagination := entity.Pagination{
			OrderBy: "id",
			Sort:    "asc",
		}

		testQuery := "id asc"
		query := pagination.ComposeOrderQuery()

		So(query, ShouldEqual, testQuery)
	})
}

func TestPaginationEntity_CalculatePage(t *testing.T) {
	Convey("計算搜索數量", t, func() {
		pagination := entity.Pagination{
			PageSize: 10,
		}

		var count int64 = 101
		pagination.CalculatePage(count)

		So(pagination.LastPage, ShouldEqual, 11)
		So(pagination.Total, ShouldEqual, count)
	})
}

func TestPaginationEntity_Format(t *testing.T) {
	Convey("結構化搜索結果", t, func() {
		type testStruct struct {
			val int
		}

		testData := testStruct{3}
		pagination := entity.Pagination{
			Page:     1,
			PageSize: 1,
			LastPage: 1,
			OrderBy:  "id",
			Sort:     "asc",
		}

		pageRes := entity.PageResponse{
			Meta: entity.Pagination{
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
