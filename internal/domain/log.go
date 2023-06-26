package domain

import (
	logV1 "github.com/ZQCard/kbk-log/api/log/v1"
	"github.com/jinzhu/copier"
)

type Log struct {
	// 主键id
	Id int64
	// 名称
	Name string
	// kratos操作路径
	Operation string
	// trace id
	TraceId string
	// http/ grpc
	Component string
	// 相关表
	Table string
	// 记录主键
	PrimaryKey string
	// user_id
	UserId string
	// 用户名
	Username string
	// 角色
	Role string
	// 请求方法
	Method string
	// 请求路径
	Path string
	// 请求参数
	Request string
	// 返回参数
	Code string
	// 返回信息
	Reason string
	// 请求ip
	IP string
	// 响应时长
	Latency string
	// 创建时间
	CreatedAt string
}

// Pb 将domain结构体转换为pb结构体
func (log *Log) Pb() *logV1.Log {
	pb := &logV1.Log{}
	copier.Copy(pb, log)
	return pb
}

// ToDomainExample 将pb结构体转换为domain包下的Example结构体
func ToDomainLog(data interface{}) *Log {
	Log := &Log{}
	copier.Copy(Log, data)
	return Log
}
