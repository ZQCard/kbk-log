package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	logV1 "github.com/ZQCard/kbk-log/api/log/v1"
	"github.com/ZQCard/kbk-log/internal/biz"
	"github.com/ZQCard/kbk-log/internal/domain"
)

type LogService struct {
	logV1.UnimplementedLogServiceServer
	logUsecase *biz.LogUsecase
	log        *log.Helper
}

func NewLogService(logUsecase *biz.LogUsecase, logger log.Logger) *LogService {
	return &LogService{
		log:        log.NewHelper(log.With(logger, "module", "service/LogService")),
		logUsecase: logUsecase,
	}
}

func (s *LogService) GetLogList(ctx context.Context, req *logV1.GetLogListReq) (*logV1.GetLogListPageRes, error) {
	params := make(map[string]interface{})
	params["name"] = req.Name
	params["user_id"] = req.UserId
	params["username"] = req.Username
	params["role"] = req.Role
	params["operation"] = req.Operation
	params["ip"] = req.Ip
	params["trace_id"] = req.TraceId

	list, count, err := s.logUsecase.ListLog(ctx, req.Page, req.PageSize, params)
	if err != nil {
		return nil, err
	}
	res := &logV1.GetLogListPageRes{}
	res.Total = int64(count)
	for _, v := range list {
		res.List = append(res.List, v.Pb())
	}
	return res, nil
}

func (s *LogService) CreateLog(ctx context.Context, req *logV1.CreateLogReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.logUsecase.CreateLog(ctx, domain.ToDomainLog(req))
}
