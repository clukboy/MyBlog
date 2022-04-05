// Code generated by goctl. DO NOT EDIT.
package types

type TagIdBody struct {
	Id string `json:"id,optional" path:"id,optional"` // 标签id
}

type TagInfo struct {
	Id          string `json:"id,optional"`          // 标签id
	Name        string `json:"name"`                 // 标签名称
	AliasName   string `json:"alias_name"`           // 标签别名
	Description string `json:"description"`          // 标签描述
	CreateTime  string `json:"create_time,optional"` // 创建时间
	UpdateTime  string `json:"update_time,optional"` // 修改时间
}

type TagPageReq struct {
	Page     int64  `json:"page"`             // 页码
	PageSize int64  `json:"pageSize"`         // 条数
	KeyWord  string `json:"keyWord,optional"` // 关键字
}
