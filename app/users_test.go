package app_test

import (
	"errors"
	"onion-architecrure-go/app"
	"onion-architecrure-go/cmd"
	"onion-architecrure-go/domain/model"
	"onion-architecrure-go/dto"
	"onion-architecrure-go/extend"
	mock_ports "onion-architecrure-go/mocks"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	"go.uber.org/zap"

	gomock "github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	"gorm.io/gorm"
)

func TestUsersApp_Signup(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := mock_ports.NewMockUserRepo(ctrl)
	mockLogger := zap.NewNop()
	config := cmd.InitAppEnv()
	app := app.NewUserApp(config, mockUserRepo, mockLogger)
	defer ctrl.Finish()

	requestId := "test-request-id"
	userData := model.Users{
		Name:         "test",
		Email:        "test",
		HashPassword: extend.Helper.Hash("test"),
	}
	requestData := dto.SignupRequest{
		Name:     "test",
		Email:    "test",
		Password: "test",
	}
	Convey("註冊", t, func() {
		Convey("成功", func() {
			gomock.InOrder(
				mockUserRepo.EXPECT().Create(gomock.Any(), userData).Return(nil),
			)

			err := app.Signup(requestId, requestData)
			So(err, ShouldBeNil)
		})

		Convey("失敗 - 帳號已存在", func() {
			gomock.InOrder(
				mockUserRepo.EXPECT().Create(gomock.Any(), userData).Return(&model.UserExistErr),
			)
			err := app.Signup(requestId, requestData)
			errData := &model.UserExistErr
			So(err, ShouldEqual, errData)
		})
	})

	Convey("Db Connect Error", t, func() {
		gomock.InOrder(
			mockUserRepo.EXPECT().Create(gomock.Any(), userData).Return(&model.DbConnectErr),
		)
		err := app.Signup(requestId, requestData)
		errData := &model.DbConnectErr
		So(err, ShouldEqual, errData)
	})
}

func TestUsersApp_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := mock_ports.NewMockUserRepo(ctrl)
	mockLogger := zap.NewNop()
	config := cmd.InitAppEnv()
	app := app.NewUserApp(config, mockUserRepo, mockLogger)
	requestId := "test-request-id"
	jwtToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDUxMzQ4NzcsInN1YiI6IlVzZXIiLCJVc2VySUQiOjF9.tPfCaNFUG-X8lu5ABtNot3sy_7FEV90PNeTtToA0adOkH4PU_hAXCbiP7BRzTpAWL-gPNaD67DrkrVdaCnFahw"
	dateTime, _ := time.Parse("2006-01-02 15:04:05", "2023-01-01 12:30:30")
	correctPwd := "fiiewl"
	faultPwd := "fiiew]l"
	defer ctrl.Finish()

	user := model.Users{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: dateTime,
			UpdatedAt: dateTime,
		},
		Name:         "test",
		Email:        "8",
		HashPassword: "1e1543e12484ce0db6275fe1d2948ef2",
	}

	Convey("登入密碼", t, func() {

		var authClaims = model.UserAuthClaims{
			UserID: 1,
		}

		monkey.PatchInstanceMethod(reflect.TypeOf(&authClaims), "GenJwt", func(_ *model.UserAuthClaims, key string) (string, error) {
			return jwtToken, nil
		})

		defer monkey.UnpatchAll()
		Convey("密碼正確", func() {
			requestBody := dto.LoginRequest{
				Email:    "1",
				Password: correctPwd,
			}
			gomock.InOrder(
				mockUserRepo.EXPECT().GetByMail(gomock.Any(), "1").Return(user, nil),
			)
			token, err := app.Login(requestId, requestBody)
			So(err, ShouldBeNil)
			So(token, ShouldEqual, jwtToken)
		})

		Convey("密碼錯誤", func() {
			requestBody := dto.LoginRequest{
				Email:    "1",
				Password: faultPwd,
			}
			user := model.Users{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: dateTime,
					UpdatedAt: dateTime,
				},
				Name:         "test",
				Email:        "8",
				HashPassword: "1e1543e12484ce0db6275fe1d2948ef2",
			}
			gomock.InOrder(
				mockUserRepo.EXPECT().GetByMail(gomock.Any(), "1").Return(user, nil),
			)
			token, err := app.Login(requestId, requestBody)
			So(token, ShouldEqual, "")
			So(err, ShouldEqual, &model.PasswordIncorrectErr)
		})

	})

	Convey("使用者不存在", t, func() {
		requestBody := dto.LoginRequest{
			Email:    "1",
			Password: correctPwd,
		}
		var user model.Users
		gomock.InOrder(
			mockUserRepo.EXPECT().GetByMail(gomock.Any(), "1").Return(user, &model.UserNotFoundErr),
		)
		token, err := app.Login(requestId, requestBody)
		So(token, ShouldEqual, "")
		So(err, ShouldEqual, &model.UserNotFoundErr)
	})

	Convey("Jwt gen error", t, func() {
		requestBody := dto.LoginRequest{
			Email:    "1",
			Password: correctPwd,
		}
		gomock.InOrder(
			mockUserRepo.EXPECT().GetByMail(gomock.Any(), "1").Return(user, nil),
		)
		defer monkey.UnpatchAll()
		var authClaims = model.UserAuthClaims{
			UserID: 1,
		}

		monkey.PatchInstanceMethod(reflect.TypeOf(&authClaims), "GenJwt", func(_ *model.UserAuthClaims, key string) (string, error) {
			return "", errors.New("error")
		})
		token, err := app.Login(requestId, requestBody)
		So(token, ShouldEqual, "")
		So(err, ShouldEqual, &model.TokenGenFail)
	})

	Convey("Db Connect Error", t, func() {
		requestBody := dto.LoginRequest{
			Email:    "1",
			Password: correctPwd,
		}
		user = model.Users{}
		gomock.InOrder(
			mockUserRepo.EXPECT().GetByMail(gomock.Any(), "1").Return(user, &model.DbConnectErr),
		)
		errData := &model.DbConnectErr
		token, err := app.Login(requestId, requestBody)
		So(token, ShouldEqual, "")
		So(err, ShouldEqual, errData)
	})

}
