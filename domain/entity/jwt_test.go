package entity_test

import (
	"encoding/json"
	"errors"
	"onion-architecrure-go/cmd"
	"onion-architecrure-go/domain/entity"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/golang-jwt/jwt"
	. "github.com/smartystreets/goconvey/convey"
)

func TestJwtEntity_VerifyJwt(t *testing.T) {
	config := cmd.InitAppEnv()
	var authClaims entity.UserAuthClaims
	jwtToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDUxMzQ4NzcsInN1YiI6IlVzZXIiLCJVc2VySUQiOjF9.tPfCaNFUG-X8lu5ABtNot3sy_7FEV90PNeTtToA0adOkH4PU_hAXCbiP7BRzTpAWL-gPNaD67DrkrVdaCnFahw"
	Convey("token 驗證成功", t, func() {
		var Parser jwt.Parser
		monkey.PatchInstanceMethod(reflect.TypeOf(&Parser), "ParseWithClaims", func(p *jwt.Parser, tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc) (*jwt.Token, error) {
			token := jwt.Token{
				Valid: true,
			}
			return &token, nil
		})
		defer monkey.UnpatchAll()

		isValid, _ := authClaims.VerifyJwt(config.JwtConfig.Key, jwtToken)
		So(isValid, ShouldBeTrue)
	})
	Convey("token 驗證失敗", t, func() {
		testErrString := "err str"
		Convey("token 失效", func() {
			isValid, _ := authClaims.VerifyJwt(config.JwtConfig.Key, testErrString)
			So(isValid, ShouldBeFalse)
		})
		Convey("token 無效 (verify 錯誤)", func() {
			isValid, _ := authClaims.VerifyJwt(config.JwtConfig.Key, testErrString)
			So(isValid, ShouldBeFalse)
		})
	})
}

func TestJwtEntity_GenJwt(t *testing.T) {
	config := cmd.InitAppEnv()
	Convey("token 生成成功", t, func() {
		var authClaims entity.UserAuthClaims
		token, _ := authClaims.GenJwt(config.JwtConfig.Key)
		isValid, _ := authClaims.VerifyJwt(config.JwtConfig.Key, token)
		So(isValid, ShouldBeTrue)
	})
	Convey("token 生成失敗", t, func() {
		var authClaims entity.UserAuthClaims
		errorString := "error testing"
		monkey.Patch(json.Marshal, func(v any) ([]byte, error) {
			return []byte{}, errors.New(errorString)
		})
		defer monkey.UnpatchAll()
		_, err := authClaims.GenJwt(config.JwtConfig.Key)
		So(err, ShouldEqual, errorString)
	})
}
