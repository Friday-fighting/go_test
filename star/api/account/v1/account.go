package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type InfoReq struct {
	g.Meta `path:"account/info" method:"get" sm:"获取信息" tags:"用户"`
}

type InfoRes struct {
	Username string `json:"username" dc:"用户名"`
	Email    string `json:"email" dc:"邮箱"`
	CreaetAt string `json:"create_at" dc:"创建时间"`
	UpdateAt string `json:"update_at" dc:"更新时间"`
}
