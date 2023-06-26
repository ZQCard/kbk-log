package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/ZQCard/kbk-log/internal/domain"
)

type LogRepo interface {
	ListLog(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*domain.Log, error)
	GetLogCount(ctx context.Context, params map[string]interface{}) (int64, error)
	CreateLog(ctx context.Context, log *domain.Log) error
}

type LogUsecase struct {
	repo   LogRepo
	logger *log.Helper
}

func NewLogUsecase(repo LogRepo, logger log.Logger) *LogUsecase {
	return &LogUsecase{repo: repo, logger: log.NewHelper(log.With(logger, "module", "usecase/log"))}
}

func (suc *LogUsecase) ListLog(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*domain.Log, int64, error) {
	list, err1 := suc.repo.ListLog(ctx, page, pageSize, params)
	if err1 != nil {
		return nil, 0, err1
	}
	count, err2 := suc.repo.GetLogCount(ctx, params)
	if err2 != nil {
		return nil, 0, err2
	}
	return list, count, nil
}

func (suc *LogUsecase) CreateLog(ctx context.Context, log *domain.Log) error {
	return suc.repo.CreateLog(ctx, log)
}
