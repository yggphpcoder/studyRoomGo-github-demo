package checkloginEquipment

import (
	"context"
	"fmt"

	config2 "studyRoomGo/common/config"

	"github.com/go-kratos/kratos/v2/errors"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
)

// Server is a server middleware.
// 检查登录设备，如果是新设备，则逼退旧设备
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			jwtCtx, _ := jwt.FromContext(ctx)
			memberId := jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)
			rand := jwtCtx.(jwtV4.MapClaims)["randToken"]
			if rand == nil {
				return nil, errors.Unauthorized("UNAUTHORIZED", "请重新登录")
			}
			eqCode, _ := GetLastLoginEquipment(ctx, int64(memberId))
			fmt.Println("eqCode", eqCode)

			// if eqCode != rand.(string) {
			// 	return nil, errors.Unauthorized("OTHER_EQUIPMENT_LOGIN", "您已在其他设备登录，请重新登录")
			// }
			return handler(ctx, req)
		}
	}
}

func SetLastLoginEquipment(ctx context.Context, memberId int64, eqCode string) (ok bool, err error) {
	// 创建redis客户端
	err = config2.AppConf.Data.Redis.Set(ctx, fmt.Sprintf("LastLoginEquipment_%d", memberId), eqCode, -1).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetLastLoginEquipment(ctx context.Context, memberId int64) (eqCode string, err error) {
	// 创建redis客户端
	str, err := config2.AppConf.Data.Redis.Get(ctx, fmt.Sprintf("LastLoginEquipment_%d", memberId)).Result()
	if err != nil {
		return "", err
	}
	return str, nil
}
