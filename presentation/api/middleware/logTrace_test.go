package middleware_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/presentation/api/middleware"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
)

func TestJwtAuthMiddleware_InjectRequestId(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	mockLogger := zap.NewNop()

	logTraceMiddleware := middleware.NewLogTraceMiddleware(mockLogger)
	var ctxRequestId string

	Convey("request id 注入", t, func() {
		router.GET("/requestId", logTraceMiddleware.InjectRequestId(), func(ctx *gin.Context) {
			res := entity.NewResEntity().ResWithSucc(nil)
			requestId := fmt.Sprintf("%v", ctx.Value("requestId"))

			ctxRequestId = requestId
			ctx.JSON(http.StatusOK, res)
		})

		req := httptest.NewRequest(http.MethodGet, "/requestId", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		resBody, _ := io.ReadAll(resp.Body)
		var resultResBody entity.ResSucc
		_ = json.Unmarshal(resBody, &resultResBody)

		So(reflect.TypeOf(resp.Header()["X-Request-Id"][0]).Kind(), ShouldEqual, reflect.String)
		So(ctxRequestId, ShouldNotBeNil)
		So(resp.Code, ShouldEqual, http.StatusOK)
		So(resultResBody.Code, ShouldEqual, "0000")
		So(resultResBody.Message, ShouldEqual, "Succ")
		So(resultResBody.Data, ShouldEqual, nil)
	})
}
