package logic


import (
	"context"
	"github.com/hibiken/asynq"
	"order/cmd/job/internal/svc"
	"order/cmd/jobtype"
)
type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register register job
func (l *CronJob) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	mux.Handle(jobtype.DeferCloseOrderJob, NewCloseOrderHandler(l.svcCtx))

	return mux
}