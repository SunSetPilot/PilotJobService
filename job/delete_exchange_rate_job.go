package job

import (
	"context"
	"time"

	"github.com/spf13/viper"

	"PilotJobService/dal"
	"PilotJobService/model"
	"PilotJobService/svc"
	"PilotJobService/utils/log"
)

func init() {
	jobs = append(jobs, &DeleteExchangeRateJob{})
}

type DeleteExchangeRateJob struct {
	AbstractJob
}

func (j *DeleteExchangeRateJob) GetName() string {
	return "DeleteExchangeRateJob"
}

func (j *DeleteExchangeRateJob) Do(ctx *svc.ServiceContext) {
	expireDays := viper.GetInt(model.ExchangeRateExpireDaysConfigKey)
	log.Debugf("delete exchange rate data before %d days", expireDays)
	err := dal.TableCurrencyExchangeRate.DeleteCurrencyExchangeRateByTime(context.Background(), time.Now().AddDate(0, 0, -1*expireDays))
	if err != nil {
		log.Errorf("failed to delete exchange rate: %v", err)
	}
}
