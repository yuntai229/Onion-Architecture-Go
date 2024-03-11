package middleware_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"onion-architecrure-go/cmd"
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/presentation/api/middleware"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
)

func TestJwtAuthMiddleware_Auth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	mockLogger := zap.NewNop()
	config := cmd.InitAppEnv()

	jwtMiddelware := middleware.NewJwtMiddleware(config, mockLogger)
	logTraceMiddleware := middleware.NewLogTraceMiddleware(mockLogger)
	router.Use(logTraceMiddleware.InjectRequestId())

	jwtToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDUxMzQ4NzcsInN1YiI6IlVzZXIiLCJVc2VySUQiOjF9.tPfCaNFUG-X8lu5ABtNot3sy_7FEV90PNeTtToA0adOkH4PU_hAXCbiP7BRzTpAWL-gPNaD67DrkrVdaCnFahw"
	authHeader := fmt.Sprintf("Bearer %v", jwtToken)
	Convey("驗證成功", t, func() {
		var ctxUserId any
		router.GET("/jwt/success", jwtMiddelware.Auth(), func(ctx *gin.Context) {
			res := entity.Response.ResWithSucc(nil)
			userId, _ := ctx.Get("UserId")
			ctxUserId = userId
			ctx.JSON(http.StatusOK, res)
		})
		var authClaims entity.UserAuthClaims
		monkey.PatchInstanceMethod(reflect.TypeOf(&authClaims), "VerifyJwt", func(e *entity.UserAuthClaims, key, tokenString string) (bool, *entity.UserAuthClaims) {
			e.UserID = 999
			return true, e
		})
		defer monkey.UnpatchAll()

		req := httptest.NewRequest(http.MethodGet, "/jwt/success", nil)
		req.Header.Set("Authorization", authHeader)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		resBody, _ := io.ReadAll(resp.Body)
		var resultResBody entity.ResSucc
		_ = json.Unmarshal(resBody, &resultResBody)

		So(ctxUserId.(uint), ShouldEqual, 999)
		So(resp.Code, ShouldEqual, http.StatusOK)
		So(resultResBody.Code, ShouldEqual, "0000")
		So(resultResBody.Message, ShouldEqual, "Succ")
		So(resultResBody.Data, ShouldEqual, nil)
	})

	Convey("驗證失敗", t, func() {
		Convey("沒有 token", func() {
			router.GET("/jwt/missingToken", jwtMiddelware.Auth(), func(ctx *gin.Context) {
				newErr := entity.MissingTokenErr
				res := entity.Response.ResWithFail(newErr)
				ctx.JSON(newErr.HttpCode, res)
			})

			req := httptest.NewRequest(http.MethodGet, "/jwt/missingToken", nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			resBody, _ := io.ReadAll(resp.Body)
			var resultResBody entity.ResFail
			_ = json.Unmarshal(resBody, &resultResBody)

			So(resp.Code, ShouldEqual, entity.MissingTokenErr.HttpCode)
			So(resultResBody.Code, ShouldEqual, entity.MissingTokenErr.Code)
			So(resultResBody.Message, ShouldEqual, entity.MissingTokenErr.Message)
		})
		Convey("無效 token", func() {
			router.GET("/jwt/invalidToken", jwtMiddelware.Auth(), func(ctx *gin.Context) {
				newErr := entity.TokenInvalidErr
				res := entity.Response.ResWithFail(newErr)
				ctx.JSON(newErr.HttpCode, res)
			})

			var authClaims entity.UserAuthClaims
			monkey.PatchInstanceMethod(reflect.TypeOf(&authClaims), "VerifyJwt", func(e *entity.UserAuthClaims, key, tokenString string) (bool, *entity.UserAuthClaims) {
				return false, nil
			})
			defer monkey.UnpatchAll()

			req := httptest.NewRequest(http.MethodGet, "/jwt/invalidToken", nil)
			req.Header.Set("Authorization", authHeader)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			respBody, _ := io.ReadAll(resp.Body)
			var resultResBody entity.ResFail
			_ = json.Unmarshal(respBody, &resultResBody)

			So(resp.Code, ShouldEqual, entity.TokenInvalidErr.HttpCode)
			So(resultResBody.Code, ShouldEqual, entity.TokenInvalidErr.Code)
			So(resultResBody.Message, ShouldEqual, entity.TokenInvalidErr.Message)
		})

	})
}
