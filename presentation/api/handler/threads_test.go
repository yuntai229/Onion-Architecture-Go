package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"onion-architecrure-go/cmd"
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"
	mock_ports "onion-architecrure-go/mocks"
	"onion-architecrure-go/presentation/api/handler"
	"onion-architecrure-go/presentation/api/middleware"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/mitchellh/mapstructure"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func TestThreadsHandler_CreatePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mockThreadApp := mock_ports.NewMockThreadApp(ctrl)
	mockLogger := zap.NewNop()
	config := cmd.InitAppEnv()
	handler := handler.NewThreadHandler(mockThreadApp, mockLogger)
	defer ctrl.Finish()

	var authClaims entity.UserAuthClaims
	monkey.PatchInstanceMethod(reflect.TypeOf(&authClaims), "VerifyJwt", func(e *entity.UserAuthClaims, key, tokenString string) (bool, *entity.UserAuthClaims) {
		e.UserID = 1
		return true, e
	})
	defer monkey.UnpatchAll()

	router := gin.New()
	jwtToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDUxMzQ4NzcsInN1YiI6IlVzZXIiLCJVc2VySUQiOjF9.tPfCaNFUG-X8lu5ABtNot3sy_7FEV90PNeTtToA0adOkH4PU_hAXCbiP7BRzTpAWL-gPNaD67DrkrVdaCnFahw"
	authHeader := fmt.Sprintf("Bearer %v", jwtToken)
	jwtMiddelware := middleware.NewJwtMiddleware(config, mockLogger)
	testUrl := "/threads/post"
	router.POST(testUrl, jwtMiddelware.Auth(), handler.CreatePost)

	var userId uint = 1

	Convey("請求參數錯誤", t, func() {
		signupRequest := dto.PostRequest{}
		jsons, _ := json.Marshal(signupRequest)
		req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
		req.Header.Set("Authorization", authHeader)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		respBody, _ := io.ReadAll(resp.Body)
		var resultBody entity.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, entity.MissingFieldErr.HttpCode)
		So(resultBody.Code, ShouldEqual, entity.MissingFieldErr.Code)
		So(resultBody.Message, ShouldEqual, entity.MissingFieldErr.Message)
	})

	Convey("ctx get userid error", t, func() {
		router.POST("/threads/post/ctxGetError", handler.CreatePost)
		postRequest := dto.PostRequest{
			Content: "test",
		}

		jsons, _ := json.Marshal(postRequest)
		req := httptest.NewRequest(http.MethodPost, "/threads/post/ctxGetError", bytes.NewBuffer(jsons))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		respBody, _ := io.ReadAll(resp.Body)
		var resultBody entity.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, entity.TokenInvalidErr.HttpCode)
		So(resultBody.Code, ShouldEqual, entity.TokenInvalidErr.Code)
		So(resultBody.Message, ShouldEqual, entity.TokenInvalidErr.Message)
	})

	Convey("新增貼文成功", t, func() {
		postRequest := dto.PostRequest{
			Content: "test",
		}

		gomock.InOrder(
			mockThreadApp.EXPECT().CreatePost(gomock.Any(), postRequest, userId).Return(nil),
		)

		jsons, _ := json.Marshal(postRequest)
		req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
		req.Header.Set("Authorization", authHeader)
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

	Convey("Db Connect Error", t, func() {
		postRequest := dto.PostRequest{
			Content: "test",
		}

		gomock.InOrder(
			mockThreadApp.EXPECT().CreatePost(gomock.Any(), postRequest, userId).Return(&entity.DbConnectErr),
		)

		jsons, _ := json.Marshal(postRequest)
		req := httptest.NewRequest(http.MethodPost, testUrl, bytes.NewBuffer(jsons))
		req.Header.Set("Authorization", authHeader)
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

func TestThreadsHandler_GetPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mockThreadApp := mock_ports.NewMockThreadApp(ctrl)
	mockLogger := zap.NewNop()
	config := cmd.InitAppEnv()
	handler := handler.NewThreadHandler(mockThreadApp, mockLogger)
	defer ctrl.Finish()

	var authClaims entity.UserAuthClaims
	monkey.PatchInstanceMethod(reflect.TypeOf(&authClaims), "VerifyJwt", func(e *entity.UserAuthClaims, key, tokenString string) (bool, *entity.UserAuthClaims) {
		e.UserID = 1
		return true, e
	})
	defer monkey.UnpatchAll()

	router := gin.New()
	jwtToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDUxMzQ4NzcsInN1YiI6IlVzZXIiLCJVc2VySUQiOjF9.tPfCaNFUG-X8lu5ABtNot3sy_7FEV90PNeTtToA0adOkH4PU_hAXCbiP7BRzTpAWL-gPNaD67DrkrVdaCnFahw"
	authHeader := fmt.Sprintf("Bearer %v", jwtToken)
	jwtMiddelware := middleware.NewJwtMiddleware(config, mockLogger)

	router.GET("/threads/post", jwtMiddelware.Auth(), handler.GetPost)

	dateTime, _ := time.Parse("2006-01-02 15:04:05", "2023-01-01 12:30:30")
	getPostContent := []entity.Threads{{
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

	Convey("ctx get userid error", t, func() {
		router.GET("/threads/post/ctxGetError", handler.GetPost)

		req := httptest.NewRequest(http.MethodGet, "/threads/post/ctxGetError", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		respBody, _ := io.ReadAll(resp.Body)
		var resultBody entity.ResFail
		_ = json.Unmarshal(respBody, &resultBody)

		So(resp.Code, ShouldEqual, entity.TokenInvalidErr.HttpCode)
		So(resultBody.Code, ShouldEqual, entity.TokenInvalidErr.Code)
		So(resultBody.Message, ShouldEqual, entity.TokenInvalidErr.Message)
	})

	Convey("取得貼文成功", t, func() {
		Convey("Url 夾帶 parameters 參數", func() {
			var userId uint = 10
			var page int = 1
			var pageSize int = 20
			var orderBy string = "id"
			var sort string = "asc"
			requestParams := fmt.Sprintf("userId=%v&page=%v&pageSize=%v&orderBy=%v&sort=%v", userId, page, pageSize, orderBy, sort)
			testUrl := fmt.Sprintf("%v?%v", "/threads/post", requestParams)
			pageParams := entity.Pagination{
				Page:     page,
				PageSize: pageSize,
				OrderBy:  orderBy,
				Sort:     sort,
			}
			getPostRequest := dto.GetPostRequest{
				UserId:     userId,
				Pagination: pageParams,
			}

			gomock.InOrder(
				mockThreadApp.EXPECT().GetPost(gomock.Any(), &pageParams, getPostRequest).Return(getPostContent, nil),
			)

			req := httptest.NewRequest(http.MethodGet, testUrl, nil)
			req.Header.Set("Authorization", authHeader)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			respBody, _ := io.ReadAll(resp.Body)
			var resultBody entity.ResSucc
			var data entity.PageResponse

			_ = json.Unmarshal(respBody, &resultBody)
			mapstructure.Decode(resultBody.Data, &data)

			var content []entity.Threads
			var contentElement entity.Threads
			for _, item := range data.Collection.([]any) {
				createdAt, _ := time.Parse("2006-01-02 15:04:05", item.(map[string]any)["createdAt"].(string))
				updatedAt, _ := time.Parse("2006-01-02 15:04:05", item.(map[string]any)["updatedAt"].(string))
				item.(map[string]any)["createdAt"] = createdAt
				item.(map[string]any)["updatedAt"] = updatedAt
				item.(map[string]any)["id"] = uint(item.(map[string]any)["id"].(float64))
				mapstructure.Decode(item, &contentElement)
				content = append(content, contentElement)
			}

			So(resp.Code, ShouldEqual, http.StatusOK)
			So(resultBody.Code, ShouldEqual, "0000")
			So(resultBody.Message, ShouldEqual, "Succ")
			So(data.Meta, ShouldEqual, pageParams)
			So(content, ShouldEqual, getPostContent)
		})
		Convey("Url 不夾帶 parameters 參數", func() {
			var userId uint = 1
			var page int = 1
			var pageSize int = 20
			var orderBy string = "id"
			var sort string = "desc"
			testUrl := "/threads/post"
			pageParams := entity.Pagination{
				Page:     page,
				PageSize: pageSize,
				OrderBy:  orderBy,
				Sort:     sort,
			}
			getPostRequest := dto.GetPostRequest{
				UserId:     userId,
				Pagination: pageParams,
			}

			gomock.InOrder(
				mockThreadApp.EXPECT().GetPost(gomock.Any(), &pageParams, getPostRequest).Return(getPostContent, nil),
			)

			req := httptest.NewRequest(http.MethodGet, testUrl, nil)
			req.Header.Set("Authorization", authHeader)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			respBody, _ := io.ReadAll(resp.Body)
			var resultBody entity.ResSucc
			var data entity.PageResponse

			_ = json.Unmarshal(respBody, &resultBody)
			mapstructure.Decode(resultBody.Data, &data)

			var content []entity.Threads
			var contentElement entity.Threads
			for _, item := range data.Collection.([]any) {
				createdAt, _ := time.Parse("2006-01-02 15:04:05", item.(map[string]any)["createdAt"].(string))
				updatedAt, _ := time.Parse("2006-01-02 15:04:05", item.(map[string]any)["updatedAt"].(string))
				item.(map[string]any)["createdAt"] = createdAt
				item.(map[string]any)["updatedAt"] = updatedAt
				item.(map[string]any)["id"] = uint(item.(map[string]any)["id"].(float64))
				mapstructure.Decode(item, &contentElement)
				content = append(content, contentElement)
			}

			So(resp.Code, ShouldEqual, http.StatusOK)
			So(resultBody.Code, ShouldEqual, "0000")
			So(resultBody.Message, ShouldEqual, "Succ")
			So(data.Meta, ShouldEqual, pageParams)
			So(content, ShouldEqual, getPostContent)
		})
	})

	Convey("Db Connect Error", t, func() {
		var userId uint = 1
		var page int = 1
		var pageSize int = 20
		var orderBy string = "id"
		var sort string = "asc"
		requestParams := fmt.Sprintf("userId=%v&page=%v&pageSize=%v&orderBy=%v&sort=%v", userId, page, pageSize, orderBy, sort)
		testUrl := fmt.Sprintf("%v?%v", "/threads/post", requestParams)
		pageParams := entity.Pagination{
			Page:     page,
			PageSize: pageSize,
			OrderBy:  orderBy,
			Sort:     sort,
		}
		getPostRequest := dto.GetPostRequest{
			UserId:     userId,
			Pagination: pageParams,
		}

		gomock.InOrder(
			mockThreadApp.EXPECT().GetPost(gomock.Any(), &pageParams, getPostRequest).Return(getPostContent, &entity.DbConnectErr),
		)

		req := httptest.NewRequest(http.MethodGet, testUrl, nil)
		req.Header.Set("Authorization", authHeader)
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
