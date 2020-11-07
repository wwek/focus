// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"focus/app/model/internal"
	"github.com/gogf/gf/os/gtime"
)

const (
	TopicListDefaultSize = 10
	TopicListMaxSize     = 50
	TopicSortDefault     = 0
	TopicSortActive      = 1
	TopicSortHot         = 2
)

// Topic is the golang structure for table gf_topic.
type Topic internal.Topic

// 主要用于列表展示
type TopicListItem struct {
	Id         uint        `json:"id"`          // 自增ID
	CategoryId uint        `json:"category_id"` // 栏目ID
	UserId     uint        `json:"user_id"`     // 用户ID
	Title      string      `json:"title"`       // 标题
	Content    string      `json:"content"`     // 内容
	Sort       uint        `json:"sort"`        // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	Brief      string      `json:"brief"`       // 摘要
	Thumb      string      `json:"thumb"`       // 缩略图
	Tags       string      `json:"tags"`        // 标签名称列表，以JSON存储
	Referer    string      `json:"referer"`     // 内容来源，例如github/gitee
	Status     uint        `json:"status"`      // 状态 0: 正常, 1: 禁用
	ViewCount  uint        `json:"view_count"`  // 浏览数量
	ZanCount   uint        `json:"zan_count"`   // 赞
	CaiCount   uint        `json:"cai_count"`   // 踩
	CreatedAt  *gtime.Time `json:"created_at"`  // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at"`  // 修改时间
}

// 绑定到Topic列表中的用户信息
type TopicUserItem struct {
	Id       uint   `json:"id"`       // UID
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像地址
}

// Topic详情
type TopicDetail struct {
	Topic Topic
	User  User
}

// API查看主题详情
type TopicApiDetailReq struct {
	Id uint `v:"min:1#请选择查看的主题"`
}

// API创建/修改主题基类
type TopicApiCreateUpdateBase struct {
	TopicServiceCreateUpdateBase
	CategoryId uint   `v:"min:1#请输入栏目ID"`    // 栏目ID
	Title      string `v:"required#请输入主题标题"` // 标题
	Content    string `v:"required#请输入主题内容"` // 内容
}

// API创建主题
type TopicApiCreateReq struct {
	TopicApiCreateUpdateBase
}

// API修改主题
type TopicApiUpdateReq struct {
	TopicApiCreateUpdateBase
	Id uint `v:"min:1#请选择需要修改的主题"` // 修改时ID不能为空
}

// API删除主题
type TopicApiDeleteReq struct {
	Id uint `v:"min:1#请选择需要删除的主题"` // 删除时ID不能为空
}

// Service查询列表
type TopicServiceGetListReq struct {
	Cate uint // 栏目ID
	Page int  // 分页号码
	Size int  // 分页数量，最大50
	Sort int  // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// Service查询列表结果
type TopicServiceGetListRes struct {
	List  []*TopicServiceGetListResItem `json:"list"`  // 列表
	Page  int                           `json:"page"`  // 分页码
	Size  int                           `json:"size"`  // 分页数量
	Total int                           `json:"total"` // 数据总数
}

type TopicServiceGetListResItem struct {
	Topic *TopicListItem `json:"topic"`
	User  *TopicUserItem `json:"user"`
}

// Service查询详情结果
type TopicServiceGetDetailRes struct {
	Topic *Topic         `json:"topic"`
	User  *TopicUserItem `json:"user"`
}

// Service创建/修改主题基类
type TopicServiceCreateUpdateBase struct {
	CategoryId uint     // 栏目ID
	Title      string   // 标题
	Content    string   // 内容
	Brief      string   // 摘要
	Thumb      string   // 缩略图
	Tags       []string // 标签名称列表，以JSON存储
	Referer    string   // 内容来源，例如github/gitee
}

// Service创建主题
type TopicServiceCreateReq struct {
	TopicServiceCreateUpdateBase
	UserId uint
}

// Service修改主题
type TopicServiceUpdateReq struct {
	TopicServiceCreateUpdateBase
	Id uint
}
