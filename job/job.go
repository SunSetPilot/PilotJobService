package job

import (
	"time"

	"github.com/gorhill/cronexpr"
	"github.com/spf13/viper"

	"PilotJobService/svc"
	"PilotJobService/utils/log"
)

var jobs []IJob

type IJob interface {
	GetName() string
	Do(ctx *svc.ServiceContext)

	shouldExecute(job IJob) bool
}

type AbstractJob struct {
	lastExecuteTime time.Time
}

func (j *AbstractJob) shouldExecute(job IJob) bool {
	jobName := job.GetName()
	enable := viper.GetBool("Job." + jobName + ".Enabled")
	if !enable {
		return false
	}
	cronStr := viper.GetString("Job." + jobName + ".Cron")
	expr, err := cronexpr.Parse(cronStr)
	if err != nil {
		log.Errorf("failed to parse cron expression: %s, error: %v", cronStr, err)
		return false
	}
	now := time.Now()
	if j.lastExecuteTime.IsZero() {
		j.lastExecuteTime = now
	}
	nextTime := expr.Next(j.lastExecuteTime)
	log.Debugf("job %s next execute time: %s", jobName, nextTime.Format("2006-01-02 15:04:05"))
	if now.Equal(nextTime) || now.After(nextTime) {
		j.lastExecuteTime = now
		return true
	}
	return false
}

type Scheduler struct {
	Ctx *svc.ServiceContext
}

func (s *Scheduler) Start() {
	for _, job := range jobs {
		log.Infof("register job %s", job.GetName())
		go func(job IJob) {
			for {
				time.Sleep(1 * time.Second)
				if job.shouldExecute(job) {
					log.Infof("job %s start", job.GetName())
					job.Do(s.Ctx)
					log.Infof("job %s finish", job.GetName())
				}
			}
		}(job)
	}
}

func NewScheduler(ctx *svc.ServiceContext) *Scheduler {
	return &Scheduler{Ctx: ctx}
}
