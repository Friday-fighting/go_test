package users

import (
	"context"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/golang-jwt/jwt/v5"
	"star/internal/consts"
	"star/internal/dao"
	"star/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

type jwtClaims struct {
	ID       uint
	Username string
	jwt.RegisteredClaims
}

func (u *Users) Login(ctx context.Context, username, password string) (tokenString string, err error) {
	var user entity.Users
	err = dao.Users.Ctx(ctx).Where("username", username).Scan(&user)
	if err != nil {
		return "", gerror.New("用户名或密码错误")
	}
	if user.Id == 0 {
		return "", gerror.New("用户不存在")
	}
	if user.Password != u.encryptPassword(password) {
		return "", gerror.New("用户名或密码错误")
	}
	uc := &jwtClaims{
		ID:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)), // 设置过期时间为6小时
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	return token.SignedString([]byte(consts.JwtKey))
}

func (u *Users) Info(ctx context.Context) (user *entity.Users, err error) {
	tokenString := g.RequestFromCtx(ctx).Request.Header.Get("Authorization")
	tokenClaims, _ := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return consts.JwtKey, nil
	})
	if claims, ok := tokenClaims.Claims.(*jwtClaims); ok && tokenClaims.Valid {
		err = dao.Users.Ctx(ctx).Where("id", claims.Id).Scan(&user)
	}
	return
}
