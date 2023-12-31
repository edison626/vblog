package token

import (
	"encoding/json"
	"time"

	"github.com/edison626/vblog/apps/user"
	"github.com/edison626/vblog/exception"
	"github.com/rs/xid"
)

func NewToken() *Token {
	return &Token{
		// 生产一个UUID的字符串
		AccessToken:           xid.New().String(),
		AccessTokenExpiredAt:  7200,
		RefreshToken:          xid.New().String(),
		RefreshTokenExpiredAt: 3600 * 24 * 7, //3600 秒 * 24 小时 * 7天
		CreatedAt:             time.Now().Unix(),
	}
}

type Token struct {
	// 该Token是颁发
	UserId int64 `json:"user_id"`
	// 人的名称，- 名显示是 user_name ， 需要gorm 标签和数据一样 username
	UserName string `json:"username" gorm:"column:username"`
	// 办法给用户的访问令牌(用户需要携带Token来访问接口)
	AccessToken string `json:"access_token"`
	// 过期时间(2h), 单位是秒
	AccessTokenExpiredAt int `json:"access_token_expired_at"`
	// 刷新Token
	RefreshToken string `json:"refresh_token"`
	// 刷新Token过期时间(7d)
	RefreshTokenExpiredAt int `json:"refresh_token_expired_at"`

	// 创建时间
	CreatedAt int64 `json:"created_at"`
	// 更新实现
	UpdatedAt int64 `json:"updated_at"`
	// 额外补充信息, gorm忽略处理
	Role user.Role `gorm:"-"`
}

func (t *Token) TableName() string {
	return "tokens"
}

func (t *Token) IsExpired() error {
	//duration 是个减法
	duration := time.Since(t.ExpiredTime())
	expiredSeconds := duration.Seconds()
	if expiredSeconds > 0 {
		return exception.NewTokenExpired("token %s 过期了 %f秒",
			t.AccessToken, expiredSeconds)
	}
	return nil
}

// 计算Token的过期时间
func (t *Token) ExpiredTime() time.Time {
	return time.Unix(t.CreatedAt, 0).
		Add(time.Duration(t.AccessTokenExpiredAt) * time.Second)
}

func (u *Token) String() string {
	dj, _ := json.Marshal(u)
	return string(dj)
}
