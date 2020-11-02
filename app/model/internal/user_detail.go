// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// UserDetail is the golang structure for table gf_user_detail.
type UserDetail struct {
    UserId    uint        `orm:"user_id,primary" json:"user_id"`    // 用户ID                                        
    TrueName  string      `orm:"true_name"       json:"true_name"`  // 真实姓名                                      
    IdNumber  string      `orm:"id_number"       json:"id_number"`  // 身份证号                                      
    City      string      `orm:"city"            json:"city"`       // 城市                                          
    Birth     *gtime.Time `orm:"birth"           json:"birth"`      // 出生日期 (字符串，例如：1986-10-07 00:00:00)  
    Phone     string      `orm:"phone"           json:"phone"`      // 手机号码                                      
    Qq        string      `orm:"qq"              json:"qq"`         // QQ                                            
    Email     string      `orm:"email"           json:"email"`      // 邮件                                          
    From      string      `orm:"from"            json:"from"`       // 用户来源                                      
    Brief     string      `orm:"brief"           json:"brief"`      // 用户说明                                      
    Remark    string      `orm:"remark"          json:"remark"`     // github 返回                                   
    CreatedAt *gtime.Time `orm:"created_at"      json:"created_at"` //                                               
    UpdatedAt *gtime.Time `orm:"updated_at"      json:"updated_at"` //                                               
}