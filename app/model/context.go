// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import "github.com/gogf/gf/net/ghttp"

const (
	ContextUserKey          = "ContextUserKey"
	ContextMessageTypeInfo  = "info"
	ContextMessageTypeWarn  = "warn"
	ContextMessageTypeError = "error"
)

// 请求上下文结构
type Context struct {
	User    *ContextUser    // 上下文用户信息
	Session *ghttp.Session  // 当前Session管理对象
	Message *ContextMessage // 上下文消息提示
}

// 请求上下文中的用户信息
type ContextUser struct {
	Id       uint   // 用户ID
	Passport string // 用户账号
	Nickname string // 用户名称
}

// 请求上下文提示信息，请求执行结束后清空
type ContextMessage struct {
	Type    string // 消息类型: info, warn, error
	Content string // 消息内容
}
