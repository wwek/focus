// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package define

// API用户注册
type UserApiRegisterReq struct {
	Passport string `v:"required#请输入账号"`  // 账号
	Password string `v:"required#请输入密码"`  // 密码
	Nickname string `v:"required#请输入昵称"`  // 昵称
	Captcha  string `v:"required#请输入验证码"` // 验证码
}

// API修改用户密码
type UserApiPasswordReq struct {
	OldPassword string `v:"required#请输入原始密码"` // 原密码
	NewPassword string `v:"required#请输入新密码"`  // 新密码
}

// API修改用户
type UserApiUpdateProfileReq struct {
	Id       uint   // 用户ID
	Nickname string `v:"required#请输入昵称信息"` // 昵称
	Avatar   string // 头像地址
	Gender   int    // 性别 0: 未设置 1: 男 2: 女
}

// Service用户信息
type UserProfileRes struct {
	Id       uint           // 用户ID
	Nickname string         `v:"required#请输入昵称信息"` // 昵称
	Avatar   string         // 头像地址
	Gender   int            // 性别 0: 未设置 1: 男 2: 女
	Stats    map[string]int // 发布内容数量
}

// Service修改用户头像
type UserServiceUpdateAvatarReq struct {
	Avatar string // 头像地址
}

// Service修改用户
type UserServiceUpdateProfileReq struct {
	Nickname string // 昵称
	Gender   int    // 性别 0: 未设置 1: 男 2: 女
}

// API禁用用户
type UserApiDisableReq struct {
	Id *uint `v:"required#请选择需要禁用的用户"` // 删除时ID不能为空
}

// Api用户登录
type UserApiLoginReq struct {
	Passport string `v:"required#请输入账号"`  // 账号
	Password string `v:"required#请输入密码"`  // 密码(明文)
	Captcha  string `v:"required#请输入验证码"` // 验证码
}

// Service用户登录
type UserServiceLoginReq struct {
	Passport string `v:"required#请输入账号"` // 账号
	Password string `v:"required#请输入密码"` // 密码(明文)
}

// Service创建用户
type UserServiceRegisterReq struct {
	Passport string // 账号
	Password string // 密码(明文)
	Nickname string // 昵称
}

// Service查询用户列表请求
type UserServiceGetListReq struct {
	ContentServiceGetListReq
	Id uint
}

// Service查询用户详情结果
type UserServiceGetListRes struct {
	Content *ContentServiceGetListRes `json:"content"` // 查询用户
	User    *UserProfileRes           `json:"user"`    // 查询用户
	Stats   map[string]int            // 发布内容数量
}

// Service查询用户列表查询请求
type UserServiceGetMessageListReq struct {
	Page       int    `json:"page"`        // 分页码
	Size       int    `json:"size"`        // 分页数量
	TargetType string `json:"target_type"` // 数据类型
	TargetId   int    `json:"target_id"`   // 数据ID
	UserId     uint   `json:"user_id"`     // 用户ID
}

// Service查询用户列表查询结果
type UserServiceGetMessageListRes struct {
	List  []*ReplyServiceGetListResItem `json:"list"`  // 列表
	Page  int                           `json:"page"`  // 分页码
	Size  int                           `json:"size"`  // 分页数量
	Total int                           `json:"total"` // 数据总数
	Stats map[string]int                // 发布内容数量
}
