package app_test

import (
	"onion-architecrure-go/app"
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"
	mock_ports "onion-architecrure-go/mocks"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func TestThreadApp_CreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockThreadRepo := mock_ports.NewMockThreadRepo(ctrl)
	mockLogger := zap.NewNop()

	requestId := "test-request-id"
	var userId uint = 1
	threadData := entity.Threads{
		UserId:  userId,
		Content: "test",
	}
	app := app.NewThreadApp(mockThreadRepo, mockLogger)

	defer ctrl.Finish()
	Convey("新增貼文成功", t, func() {
		gomock.InOrder(
			mockThreadRepo.EXPECT().Create(gomock.Any(), threadData).Return(nil),
		)
		requestData := dto.PostRequest{
			Content: "test",
		}

		err := app.CreatePost(requestId, requestData, userId)
		So(err, ShouldBeNil)
	})

	Convey("Db Connect Error", t, func() {
		gomock.InOrder(
			mockThreadRepo.EXPECT().Create(gomock.Any(), threadData).Return(&entity.DbConnectErr),
		)
		requestData := dto.PostRequest{
			Content: "test",
		}
		errData := &entity.DbConnectErr

		err := app.CreatePost(requestId, requestData, userId)
		So(err, ShouldEqual, errData)
	})
}

func TestThreadApp_GetPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockThreadRepo := mock_ports.NewMockThreadRepo(ctrl)
	mockLogger := zap.NewNop()
	app := app.NewThreadApp(mockThreadRepo, mockLogger)

	requestId := "test-request-id"
	dateTime, _ := time.Parse("2006-01-02 15:04:05", "2023-01-01 12:30:30")

	res := []entity.Threads{{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: dateTime,
			UpdatedAt: dateTime,
		},
		UserId:  1,
		Content: "tt",
	}, {
		Model: gorm.Model{
			ID:        2,
			CreatedAt: dateTime,
			UpdatedAt: dateTime,
		},
		UserId:  1,
		Content: "tts",
	}}
	defer ctrl.Finish()

	Convey("取得貼文成功", t, func() {
		var userId uint = 1
		pageParams := entity.Pagination{
			Page:     1,
			PageSize: 20,
			OrderBy:  "id",
			Sort:     "asc",
		}
		requestData := dto.GetPostRequest{
			UserId:     userId,
			Pagination: pageParams,
		}

		gomock.InOrder(
			mockThreadRepo.EXPECT().GetByUserId(gomock.Any(), &pageParams, userId).Return(res, nil),
		)
		data, err := app.GetPost(requestId, &pageParams, requestData)
		So(err, ShouldBeNil)
		So(data, ShouldEqual, res)
	})

	Convey("Db Connect Error", t, func() {
		var userId uint = 1
		pageParams := entity.Pagination{
			Page:     1,
			PageSize: 20,
			OrderBy:  "id",
			Sort:     "asc",
		}
		requestData := dto.GetPostRequest{
			UserId:     userId,
			Pagination: pageParams,
		}

		gomock.InOrder(
			mockThreadRepo.EXPECT().GetByUserId(gomock.Any(), &pageParams, userId).Return(res, &entity.DbConnectErr),
		)
		_, err := app.GetPost(requestId, &pageParams, requestData)
		errData := &entity.DbConnectErr
		So(err, ShouldEqual, errData)
	})

}
