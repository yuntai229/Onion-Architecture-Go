package handler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"onion-architecrure-go/domain/constant"
	"onion-architecrure-go/domain/model"
	"onion-architecrure-go/dto"
	mock_ports "onion-architecrure-go/mocks"
	"onion-architecrure-go/presentation/api/handler"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
)

func TestUsersHandler_Signup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mockUserApp := mock_ports.NewMockUserApp(ctrl)
	mockLogger := zap.NewNop()
	handler := handler.NewUserHandler(mockUserApp, mockLogger)
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
		var resultBody model.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, constant.MissingFieldErr.HttpCode)
		So(resultBody.Code, ShouldEqual, constant.MissingFieldErr.Code)
		So(resultBody.Message, ShouldEqual, constant.MissingFieldErr.Message)
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
				mockUserApp.EXPECT().Signup(gomock.Any(), signupRequest).Return(nil),
			)
			req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			respBody, _ := io.ReadAll(resp.Body)
			var resultBody model.ResSucc
			_ = json.Unmarshal(respBody, &resultBody)

			So(resp.Code, ShouldEqual, http.StatusOK)
			So(resultBody.Code, ShouldEqual, "0000")
			So(resultBody.Message, ShouldEqual, "Succ")
			So(resultBody.Data, ShouldEqual, nil)
		})

		Convey("失敗 - 帳號已存在", func() {
			gomock.InOrder(
				mockUserApp.EXPECT().Signup(gomock.Any(), signupRequest).Return(&constant.UserExistErr),
			)
			req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			respBody, _ := io.ReadAll(resp.Body)
			var resultBody model.ResFail
			_ = json.Unmarshal(respBody, &resultBody)

			So(resp.Code, ShouldEqual, constant.UserExistErr.HttpCode)
			So(resultBody.Code, ShouldEqual, constant.UserExistErr.Code)
			So(resultBody.Message, ShouldEqual, constant.UserExistErr.Message)
		})
	})

	Convey("Db Connect Error", t, func() {
		signupRequest := dto.SignupRequest{
			Name:     "test",
			Email:    "test",
			Password: "ooo",
		}
		gomock.InOrder(
			mockUserApp.EXPECT().Signup(gomock.Any(), signupRequest).Return(&constant.DbConnectErr),
		)

		jsons, _ := json.Marshal(signupRequest)
		req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		respBody, _ := io.ReadAll(resp.Body)
		var resultBody model.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, constant.DbConnectErr.HttpCode)
		So(resultBody.Code, ShouldEqual, constant.DbConnectErr.Code)
		So(resultBody.Message, ShouldEqual, constant.DbConnectErr.Message)
	})
}

func TestUsersHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mockUserApp := mock_ports.NewMockUserApp(ctrl)
	mockLogger := zap.NewNop()
	handler := handler.NewUserHandler(mockUserApp, mockLogger)
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
		var resultBody model.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, constant.MissingFieldErr.HttpCode)
		So(resultBody.Code, ShouldEqual, constant.MissingFieldErr.Code)
		So(resultBody.Message, ShouldEqual, constant.MissingFieldErr.Message)
	})

	Convey("使用者不存在", t, func() {
		loginRequest := dto.LoginRequest{
			Email:    "test",
			Password: "test",
		}
		gomock.InOrder(
			mockUserApp.EXPECT().Login(gomock.Any(), loginRequest).Return("", &constant.UserNotFoundErr),
		)

		jsons, _ := json.Marshal(loginRequest)
		req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		respBody, _ := io.ReadAll(resp.Body)
		var resultBody model.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, constant.UserNotFoundErr.HttpCode)
		So(resultBody.Code, ShouldEqual, constant.UserNotFoundErr.Code)
		So(resultBody.Message, ShouldEqual, constant.UserNotFoundErr.Message)
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
				mockUserApp.EXPECT().Login(gomock.Any(), loginRequest).Return(jwtToken, nil),
			)

			jsons, _ := json.Marshal(loginRequest)
			req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			respBody, _ := io.ReadAll(resp.Body)
			var resultBody model.ResSucc
			_ = json.Unmarshal(respBody, &resultBody)

			So(resp.Code, ShouldEqual, http.StatusOK)
			So(resultBody.Code, ShouldEqual, "0000")
			So(resultBody.Message, ShouldEqual, "Succ")
			So(resultBody.Data, ShouldEqual, resData)
		})

		Convey("密碼錯誤", func() {
			gomock.InOrder(
				mockUserApp.EXPECT().Login(gomock.Any(), loginRequest).Return("", &constant.PasswordIncorrectErr),
			)

			jsons, _ := json.Marshal(loginRequest)
			req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			respBody, _ := io.ReadAll(resp.Body)
			var resultBody model.ResFail
			_ = json.Unmarshal(respBody, &resultBody)

			So(resp.Code, ShouldEqual, constant.PasswordIncorrectErr.HttpCode)
			So(resultBody.Code, ShouldEqual, constant.PasswordIncorrectErr.Code)
			So(resultBody.Message, ShouldEqual, constant.PasswordIncorrectErr.Message)
		})
	})

	Convey("Jwt gen error", t, func() {
		loginRequest := dto.LoginRequest{
			Email:    "test",
			Password: "test",
		}
		gomock.InOrder(
			mockUserApp.EXPECT().Login(gomock.Any(), loginRequest).Return("", &constant.TokenGenFail),
		)

		jsons, _ := json.Marshal(loginRequest)
		req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		respBody, _ := io.ReadAll(resp.Body)
		var resultBody model.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, constant.TokenGenFail.HttpCode)
		So(resultBody.Code, ShouldEqual, constant.TokenGenFail.Code)
		So(resultBody.Message, ShouldEqual, constant.TokenGenFail.Message)
	})

	Convey("Db Connect Error", t, func() {
		loginRequest := dto.LoginRequest{
			Email:    "test",
			Password: "test",
		}
		gomock.InOrder(
			mockUserApp.EXPECT().Login(gomock.Any(), loginRequest).Return("", &constant.DbConnectErr),
		)

		jsons, _ := json.Marshal(loginRequest)
		req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		respBody, _ := io.ReadAll(resp.Body)
		var resultBody model.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, constant.DbConnectErr.HttpCode)
		So(resultBody.Code, ShouldEqual, constant.DbConnectErr.Code)
		So(resultBody.Message, ShouldEqual, constant.DbConnectErr.Message)
	})
}
