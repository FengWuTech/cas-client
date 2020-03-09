package models

import "time"

type Query map[string]interface{}
type RequestBody map[string]interface{}
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Model struct {
	ID         *int       `json:"id"`
	CreateTime *time.Time `json:"create_time"`
	UpdateTime *time.Time `json:"update_time"`
}

type App struct {
	Model
	Code   *string `json:"code,omitempty"`    // 标识
	Name   *string `json:"name,omitempty"`    // 名称
	ApiKey *string `json:"api_key,omitempty"` // APIKEY
}

type Action struct {
	Model
	AppID *int    `json:"app_id,omitempty"` // 应用id
	Code  *string `json:"code,omitempty"`   // 操作代码
	Name  *string `json:"name,omitempty"`   // 操作名称
}

type Resource struct {
	Model
	AppID *int    `json:"app_id,omitempty"` // 应用id
	Type  *int    `json:"type,omitempty"`   // 资源类型：1菜单
	Code  *string `json:"code,omitempty"`   // 资源代码
	Name  *string `json:"name,omitempty"`   // 资源名称
	Pid   *int    `json:"pid,omitempty"`    // 父id
	Data  *string `json:"data,omitempty"`   // 资源携带的数据
}

type ResourceTree struct {
	Resource Resource       `json:"resource"`
	Action   []Action       `json:"action"`
	Children []ResourceTree `json:"children"`
}

type ResourceAction struct {
	Model
	AppID      *int `json:"app_id,omitempty"`      // 应用ID
	ResourceID *int `json:"resource_id,omitempty"` // 资源ID
	ActionID   *int `json:"action_id,omitempty"`   // 操作ID
}

type Role struct {
	Model
	AppID       *int    `json:"app_id,omitempty"`      // 角色对应的APPID
	Code        *string `json:"code,omitempty"`        // 角色描述字符串
	Name        *string `json:"name,omitempty"`        // 角色名称
	Description *string `json:"description,omitempty"` // 其他描述信息
}

type RoleAction struct {
	Model
	AppID    *int `json:"app_id,omitempty"`    // 应用ID
	RoleID   *int `json:"role_id,omitempty"`   // 角色ID
	ActionID *int `json:"action_id,omitempty"` // 操作ID
}

type RoleResource struct {
	Model
	AppID      *int `json:"app_id,omitempty"`      // 应用ID
	RoleID     *int `json:"role_id,omitempty"`     // 角色ID
	ResourceID *int `json:"resource_id,omitempty"` // 资源ID
}

type UserRole struct {
	Model
	AppID  *int `json:"app_id,omitempty"`  // 应用ID
	UserID *int `json:"user_id,omitempty"` // 用户在CAS系统UID
	RoleID *int `json:"role_id,omitempty"` // 角色ID
}
