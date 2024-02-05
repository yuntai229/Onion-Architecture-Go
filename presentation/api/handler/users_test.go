package handler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"
	mock_ports "onion-architecrure-go/mocks"
	"onion-architecrure-go/presentation/api/handler"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUsersHandler_Signup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mockUserApp := mock_ports.NewMockUserApp(ctrl)
	handler := handler.NewUserHandler(mockUserApp)
	defer ctrl.Finish()

	router := gin.New()
	testUrl := "/users/signup"
	router.POST(testUrl, handler.Signup)

	Convey("請求參數錯誤", t, func() {
		signupRequest := dto.SignupRequest{
			Name:  "test",
			Email: "test",
		}
		jsons, _ := json.Marshal(signupRequest)
		req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		respBody, _ := io.ReadAll(resp.Body)
		var resultBody entity.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, entity.MissingFieldErr.HttpCode)
		So(resultBody.Code, ShouldEqual, entity.MissingFieldErr.Code)
		So(resultBody.Message, ShouldEqual, entity.MissingFieldErr.Message)
	})

	Convey("註冊", t, func() {
		signupRequest := dto.SignupRequest{
			Name:     "test",
			Email:    "test",
			Password: "ooo",
		}
		jsons, _ := json.Marshal(signupRequest)
		Convey("成功", func() {
			gomock.InOrder(
				mockUserApp.EXPECT().Signup(signupRequest).Return(nil),
			)
			req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			respBody, _ := io.ReadAll(resp.Body)
			var resultBody entity.ResSucc
			_ = json.Unmarshal(respBody, &resultBody)

			So(resp.Code, ShouldEqual, http.StatusOK)
			So(resultBody.Code, ShouldEqual, "0000")
			So(resultBody.Message, ShouldEqual, "Succ")
			So(resultBody.Data, ShouldEqual, nil)
		})

		Convey("失敗 - 帳號已存在", func() {
			gomock.InOrder(
				mockUserApp.EXPECT().Signup(signupRequest).Return(&entity.UserExistErr),
			)
			req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			respBody, _ := io.ReadAll(resp.Body)
			var resultBody entity.ResFail
			_ = json.Unmarshal(respBody, &resultBody)

			So(resp.Code, ShouldEqual, entity.UserExistErr.HttpCode)
			So(resultBody.Code, ShouldEqual, entity.UserExistErr.Code)
			So(resultBody.Message, ShouldEqual, entity.UserExistErr.Message)
		})
	})

	Convey("Db Connect Error", t, func() {
		signupRequest := dto.SignupRequest{
			Name:     "test",
			Email:    "test",
			Password: "ooo",
		}
		gomock.InOrder(
			mockUserApp.EXPECT().Signup(signupRequest).Return(&entity.DbConnectErr),
		)

		jsons, _ := json.Marshal(signupRequest)
		req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		respBody, _ := io.ReadAll(resp.Body)
		var resultBody entity.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, entity.DbConnectErr.HttpCode)
		So(resultBody.Code, ShouldEqual, entity.DbConnectErr.Code)
		So(resultBody.Message, ShouldEqual, entity.DbConnectErr.Message)
	})
}

func TestUsersHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mockUserApp := mock_ports.NewMockUserApp(ctrl)
	handler := handler.NewUserHandler(mockUserApp)
	defer ctrl.Finish()

	router := gin.New()
	testUrl := "/users/login"
	router.POST(testUrl, handler.Login)
	jwtToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDUxMzQ4NzcsInN1YiI6IlVzZXIiLCJVc2VySUQiOjF9.tPfCaNFUG-X8lu5ABtNot3sy_7FEV90PNeTtToA0adOkH4PU_hAXCbiP7BRzTpAWL-gPNaD67DrkrVdaCnFahw"

	Convey("請求參數錯誤", t, func() {
		loginRequest := dto.LoginRequest{
			Email: "test",
		}

		jsons, _ := json.Marshal(loginRequest)
		req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		respBody, _ := io.ReadAll(resp.Body)
		var resultBody entity.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, entity.MissingFieldErr.HttpCode)
		So(resultBody.Code, ShouldEqual, entity.MissingFieldErr.Code)
		So(resultBody.Message, ShouldEqual, entity.MissingFieldErr.Message)
	})

	Convey("使用者不存在", t, func() {
		loginRequest := dto.LoginRequest{
			Email:    "test",
			Password: "test",
		}
		gomock.InOrder(
			mockUserApp.EXPECT().Login(loginRequest).Return("", &entity.UserNotFoundErr),
		)

		jsons, _ := json.Marshal(loginRequest)
		req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		respBody, _ := io.ReadAll(resp.Body)
		var resultBody entity.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, entity.UserNotFoundErr.HttpCode)
		So(resultBody.Code, ShouldEqual, entity.UserNotFoundErr.Code)
		So(resultBody.Message, ShouldEqual, entity.UserNotFoundErr.Message)
	})

	Convey("登入密碼", t, func() {
		loginRequest := dto.LoginRequest{
			Email:    "test",
			Password: "test",
		}
		resData := map[string]any{
			"token": jwtToken,
		}
		Convey("密碼正確", func() {
			gomock.InOrder(
				mockUserApp.EXPECT().Login(loginRequest).Return(jwtToken, nil),
			)

			jsons, _ := json.Marshal(loginRequest)
			req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			respBody, _ := io.ReadAll(resp.Body)
			var resultBody entity.ResSucc
			_ = json.Unmarshal(respBody, &resultBody)

			So(resp.Code, ShouldEqual, http.StatusOK)
			So(resultBody.Code, ShouldEqual, "0000")
			So(resultBody.Message, ShouldEqual, "Succ")
			So(resultBody.Data, ShouldEqual, resData)
		})

		Convey("密碼錯誤", func() {
			gomock.InOrder(
				mockUserApp.EXPECT().Login(loginRequest).Return("", &entity.PasswordIncorrectErr),
			)

			jsons, _ := json.Marshal(loginRequest)
			req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			respBody, _ := io.ReadAll(resp.Body)
			var resultBody entity.ResFail
			_ = json.Unmarshal(respBody, &resultBody)

			So(resp.Code, ShouldEqual, entity.PasswordIncorrectErr.HttpCode)
			So(resultBody.Code, ShouldEqual, entity.PasswordIncorrectErr.Code)
			So(resultBody.Message, ShouldEqual, entity.PasswordIncorrectErr.Message)
		})
	})

	Convey("Jwt gen error", t, func() {
		loginRequest := dto.LoginRequest{
			Email:    "test",
			Password: "test",
		}
		gomock.InOrder(
			mockUserApp.EXPECT().Login(loginRequest).Return("", &entity.TokenGenFail),
		)

		jsons, _ := json.Marshal(loginRequest)
		req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		respBody, _ := io.ReadAll(resp.Body)
		var resultBody entity.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, entity.TokenGenFail.HttpCode)
		So(resultBody.Code, ShouldEqual, entity.TokenGenFail.Code)
		So(resultBody.Message, ShouldEqual, entity.TokenGenFail.Message)
	})

	Convey("Db Connect Error", t, func() {
		loginRequest := dto.LoginRequest{
			Email:    "test",
			Password: "test",
		}
		gomock.InOrder(
			mockUserApp.EXPECT().Login(loginRequest).Return("", &entity.DbConnectErr),
		)

		jsons, _ := json.Marshal(loginRequest)
		req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		respBody, _ := io.ReadAll(resp.Body)
		var resultBody entity.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, entity.DbConnectErr.HttpCode)
		So(resultBody.Code, ShouldEqual, entity.DbConnectErr.Code)
		So(resultBody.Message, ShouldEqual, entity.DbConnectErr.Message)
	})
}
