package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	logV1 "github.com/ZQCard/kbk-log/api/log/v1"
	"github.com/ZQCard/kbk-log/internal/biz"
	"github.com/ZQCard/kbk-log/internal/domain"
)

type LogEntity struct {
	Id        int64     `gorm:"primarykey;type:int;comment:主键id"`
	Name      string    `gorm:"type:varchar(255);comment:名称"`
	Domain    string    `gorm:"type:varchar(255);comment:域"`
	TraceId   string    `gorm:"type:varchar(255);comment:trace id"`
	Component string    `gorm:"type:varchar(255);comment:请求http/tpc"`
	UserId    string    `gorm:"type:varchar(255);comment:用户Id"`
	Username  string    `gorm:"type:varchar(255);comment:用户名"`
	Role      string    `gorm:"type:varchar(255);comment:角色"`
	Method    string    `gorm:"type:varchar(255);comment:请求方式"`
	Operation string    `gorm:"type:varchar(255);comment:kratos操作"`
	Path      string    `gorm:"type:varchar(255);comment:请求path"`
	Request   string    `gorm:"type:text;comment:请求参数"`
	Code      string    `gorm:"type:varchar(255);comment:返回code"`
	Reason    string    `gorm:"type:varchar(255);comment:返回reason"`
	IP        string    `gorm:"type:varchar(255);comment:IP"`
	Latency   string    `gorm:"type:varchar(255);comment:响应时长"`
	CreatedAt time.Time `gorm:"column:created_at;not null;comment:创建时间"`
}

func (LogEntity) TableName() string {
	return "log"
}

type LogRepo struct {
	data *Data
	log  *log.Helper
}

func NewLogRepo(data *Data, logger log.Logger) biz.LogRepo {
	r := &LogRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/log")),
	}
	return r
}

// searchParam 搜索条件
func (r LogRepo) searchParam(ctx context.Context, params map[string]interface{}) *gorm.DB {
	conn := r.data.db.Model(&LogEntity{})
	if Id, ok := params["id"]; ok && Id.(int64) != 0 {
		conn = conn.Where("id = ?", Id)
	}
	if value, ok := params["domain"]; ok && value.(string) != "" {
		conn = conn.Where("domain = ?", value)
	}
	if value, ok := params["name"]; ok && value.(string) != "" {
		conn = conn.Where("name = ?", value)
	}
	if value, ok := params["user_id"]; ok && value.(string) != "" {
		conn = conn.Where("user_id = ?", value)
	}
	if value, ok := params["trace_id"]; ok && value.(string) != "" {
		conn = conn.Where("trace_id = ?", value)
	}
	if value, ok := params["username"]; ok && value.(string) != "" {
		conn = conn.Where("username = ?", value)
	}
	if value, ok := params["role"]; ok && value.(string) != "" {
		conn = conn.Where("role = ?", value)
	}
	if value, ok := params["operation"]; ok && value.(string) != "" {
		conn = conn.Where("operation = ?", value)
	}
	if value, ok := params["ip"]; ok && value.(string) != "" {
		conn = conn.Where("ip = ?", value)
	}
	// 开始时间
	if start, ok := params["created_at_start"]; ok && start.(string) != "" {
		conn = conn.Where("created_at >= ?", start.(string)+" 00:00:00")
	}
	// 结束时间
	if end, ok := params["created_at_end"]; ok && end.(string) != "" {
		conn = conn.Where("created_at <= ?", end.(string)+" 23:59:59")
	}
	conn = getDbWithDomain(ctx, conn)
	return conn
}

func (r LogRepo) CreateLog(ctx context.Context, domain *domain.Log) error {
	entity := &LogEntity{}
	entity.Name = domain.Name
	entity.Domain = getDomain(ctx)
	entity.TraceId = domain.TraceId
	entity.Component = domain.Component
	entity.Operation = domain.Operation
	entity.UserId = domain.UserId
	entity.Username = domain.Username
	entity.Role = domain.Role
	entity.Method = domain.Method
	entity.Path = domain.Path
	entity.Request = domain.Request
	entity.Code = domain.Code
	entity.Reason = domain.Reason
	entity.IP = domain.IP
	entity.Latency = domain.Latency
	entity.CreatedAt = time.Now()
	if err := r.data.db.Model(entity).Create(entity).Error; err != nil {
		return logV1.ErrorSystemError("CreateLog Create %s", err)
	}
	return nil
}

func (r LogRepo) ListLog(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*domain.Log, error) {
	list := []*LogEntity{}
	conn := r.searchParam(ctx, params)
	err := conn.Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&list).Error
	if err != nil {
		return nil, logV1.ErrorSystemError("ListLog Find %s", err)
	}

	rv := make([]*domain.Log, 0, len(list))
	for _, record := range list {
		log := toDomainLog(record)
		rv = append(rv, log)
	}
	return rv, nil
}

func (r LogRepo) GetLogCount(ctx context.Context, params map[string]interface{}) (count int64, err error) {
	if len(params) == 0 {
		return 0, logV1.ErrorBadRequest("参数不得为空")
	}
	r.searchParam(ctx, params).Count(&count)
	return count, nil
}

func toDomainLog(log *LogEntity) *domain.Log {
	if log == nil {
		return &domain.Log{}
	}
	return &domain.Log{
		Id:        log.Id,
		Name:      log.Name,
		TraceId:   log.TraceId,
		Component: log.Component,
		Operation: log.Operation,
		UserId:    log.UserId,
		Username:  log.Username,
		Role:      log.Role,
		Method:    log.Method,
		Path:      log.Path,
		Request:   log.Request,
		Code:      log.Code,
		Reason:    log.Reason,
		IP:        log.IP,
		Latency:   log.Latency,
		CreatedAt: log.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
