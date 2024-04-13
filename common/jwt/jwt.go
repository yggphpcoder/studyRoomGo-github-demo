package jwt

import (
	"context"
	"errors"
	config2 "studyRoomGo/common/config"
	"studyRoomGo/common/jwt/gen"
	"time"

	aj "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v4"
	jwtV4 "github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	OpenId               string `json:"openId"`     //wexin openid
	MemberId             int64  `json:"memberId"`   //wexin openid
	MemberType           int8   `json:"memberType"` //wexin openid
	RandToken            string `json:"randToken"`  //wexin openid
	jwt.RegisteredClaims        // 注意!这是jwt-go的v4版本新增的，原先是jwt.StandardClaims
}

var MySecret = []byte("yuyue2023_dsafvgytukjzc")       // 定义secret，后面会用到
var MySecretDev = []byte("yuyue2023_zxcvbfghdfghrtey") // 定义secret，后面会用到

func CreateToken(OpenId string, MemberId int64, MemberType int8, randToken string) (tokenString string) {
	claim := MyClaims{
		OpenId:     OpenId,
		MemberId:   MemberId,
		MemberType: MemberType,
		RandToken:  randToken,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 365 * 10 * time.Hour)), // 过期时间10年
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, _ = token.SignedString(MySecret)
	if config2.AppConf.Environment != "pro" {
		tokenString, _ = token.SignedString(MySecretDev)

	}
	return tokenString
}

func CheckIsAdmin(ctx context.Context) (ok bool, err error) {
	jwtCtx, _ := aj.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
	memberMgr := gen.SrMemberMgr(config2.AppConf.Data.DB)
	mem, err := memberMgr.GetFromID(memberId)
	if err != nil {
		return false, err

	}
	if mem.Type != 99 {
		return false, errors.New("no admin")
	}
	return true, nil
}

func CheckIsTeacher(ctx context.Context) (ok bool, err error) {
	jwtCtx, _ := aj.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
	memberMgr := gen.SrMemberMgr(config2.AppConf.Data.DB)
	mem, err := memberMgr.GetFromID(memberId)
	if err != nil {
		return false, err

	}
	if mem.Type != 99 && mem.Type != 98 {
		return false, errors.New("no teacher or admin")
	}
	return true, nil
}
