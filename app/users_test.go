package app_test

import (
	"errors"
	"onion-architecrure-go/app"
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"
	"onion-architecrure-go/extend"
	mock_ports "onion-architecrure-go/mocks"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"

	gomock "github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	"gorm.io/gorm"
)

func TestUsersApp_Signup(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockUserRepo := mock_ports.NewMockUserRepo(ctrl)
	app := app.NewUserApp(mockUserRepo)
	defer ctrl.Finish()
	userData := entity.Users{
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
				mockUserRepo.EXPECT().Create(userData).Return(nil),
			)

			err := app.Signup(requestData)
			So(err, ShouldBeNil)
		})

		Convey("失敗 - 帳號已存在", func() {
			gomock.InOrder(
				mockUserRepo.EXPECT().Create(userData).Return(&entity.UserExistErr),
			)
			err := app.Signup(requestData)
			errData := &entity.UserExistErr
			So(err, ShouldEqual, errData)
		})
	})

	Convey("Db Connect Error", t, func() {
		gomock.InOrder(
			mockUserRepo.EXPECT().Create(userData).Return(&entity.DbConnectErr),
		)
		err := app.Signup(requestData)
		errData := &entity.DbConnectErr
		So(err, ShouldEqual, errData)
	})
}

func TestUsersApp_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := mock_ports.NewMockUserRepo(ctrl)
	app := app.NewUserApp(mockUserRepo)
	jwtToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDUxMzQ4NzcsInN1YiI6IlVzZXIiLCJVc2VySUQiOjF9.tPfCaNFUG-X8lu5ABtNot3sy_7FEV90PNeTtToA0adOkH4PU_hAXCbiP7BRzTpAWL-gPNaD67DrkrVdaCnFahw"
	dateTime, _ := time.Parse("2006-01-02 15:04:05", "2023-01-01 12:30:30")
	correctPwd := "fiiewl"
	faultPwd := "fiiew]l"
	defer ctrl.Finish()

	user := entity.Users{
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

		var authClaims = entity.UserAuthClaims{
			UserID: 1,
		}

		monkey.PatchInstanceMethod(reflect.TypeOf(&authClaims), "GenJwt", func(_ *entity.UserAuthClaims) (string, error) {
			return jwtToken, nil
		})

		defer monkey.UnpatchAll()
		Convey("密碼正確", func() {
			requestBody := dto.LoginRequest{
				Email:    "1",
				Password: correctPwd,
			}
			gomock.InOrder(
				mockUserRepo.EXPECT().GetByMail("1").Return(user, nil),
			)
			token, err := app.Login(requestBody)
			So(err, ShouldBeNil)
			So(token, ShouldEqual, jwtToken)
		})

		Convey("密碼錯誤", func() {
			requestBody := dto.LoginRequest{
				Email:    "1",
				Password: faultPwd,
			}
			user := entity.Users{
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
				mockUserRepo.EXPECT().GetByMail("1").Return(user, nil),
			)
			token, err := app.Login(requestBody)
			So(token, ShouldEqual, "")
			So(err, ShouldEqual, &entity.PasswordIncorrectErr)
		})

	})

	Convey("使用者不存在", t, func() {
		requestBody := dto.LoginRequest{
			Email:    "1",
			Password: correctPwd,
		}
		var user entity.Users
		gomock.InOrder(
			mockUserRepo.EXPECT().GetByMail("1").Return(user, &entity.UserNotFoundErr),
		)
		token, err := app.Login(requestBody)
		So(token, ShouldEqual, "")
		So(err, ShouldEqual, &entity.UserNotFoundErr)
	})

	Convey("Jwt gen error", t, func() {
		requestBody := dto.LoginRequest{
			Email:    "1",
			Password: correctPwd,
		}
		gomock.InOrder(
			mockUserRepo.EXPECT().GetByMail("1").Return(user, nil),
		)
		defer monkey.UnpatchAll()
		var authClaims = entity.UserAuthClaims{
			UserID: 1,
		}

		monkey.PatchInstanceMethod(reflect.TypeOf(&authClaims), "GenJwt", func(_ *entity.UserAuthClaims) (string, error) {
			return "", errors.New("error")
		})
		token, err := app.Login(requestBody)
		So(token, ShouldEqual, "")
		So(err, ShouldEqual, &entity.TokenGenFail)
	})

	Convey("Db Connect Error", t, func() {
		requestBody := dto.LoginRequest{
			Email:    "1",
			Password: correctPwd,
		}
		user = entity.Users{}
		gomock.InOrder(
			mockUserRepo.EXPECT().GetByMail("1").Return(user, &entity.DbConnectErr),
		)
		errData := &entity.DbConnectErr
		token, err := app.Login(requestBody)
		So(token, ShouldEqual, "")
		So(err, ShouldEqual, errData)
	})

}
