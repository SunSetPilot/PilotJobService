package job

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"

	"PilotJobService/dal"
	"PilotJobService/model"
	"PilotJobService/model/table"
	"PilotJobService/svc"
	"PilotJobService/thirdparty/bank"
	"PilotJobService/utils"
	"PilotJobService/utils/log"
)

func init() {
	jobs = append(jobs, &FetchExchangeRateJob{})
}

type FetchExchangeRateJob struct {
	AbstractJob
}

func (j *FetchExchangeRateJob) GetName() string {
	return "FetchExchangeRateJob"
}

func (j *FetchExchangeRateJob) Do(ctx *svc.ServiceContext) {
	exchangeRateUrl := viper.GetString(model.CmbExchangeRateUrlConfigKey)
	cmb := &bank.CmbApi{ExchangeRateUrl: exchangeRateUrl}
	exchangeRateList, err := utils.Retry(3, 30*time.Second, cmb.GetExchangeRateList)
	if err != nil {
		log.Errorf("FetchExchangeRateJob error: request reach the max retry count, error: %v", err)
		return
	}
	for _, exchangeRate := range exchangeRateList {
		err = SaveExchangeRate(exchangeRate)
		if err != nil {
			log.Errorf("FetchExchangeRateJob error: %v", err)
		}
	}

}

func SaveExchangeRate(exchangeRate *bank.CurrencyExchangeRateCmb) error {
	var (
		err              error
		lastExchangeRate *table.CurrencyExchangeRateModel
	)

	currencyExchangeRate := &table.CurrencyExchangeRateModel{
		Name:      exchangeRate.ZCcyNbr,
		ZRtbBid:   exchangeRate.ZRtbBid,
		ZRthOfr:   exchangeRate.ZRthOfr,
		ZRtcOfr:   exchangeRate.ZRtcOfr,
		ZRthBid:   exchangeRate.ZRthBid,
		ZRtcBid:   exchangeRate.ZRtcBid,
		Source:    "cmb",
		Timestamp: time.Now(),
	}
	currencyExchangeRate.Hash = utils.GetHash(fmt.Sprintf("%f-%f-%f-%f-%f-%s", currencyExchangeRate.ZRtbBid, currencyExchangeRate.ZRthOfr, currencyExchangeRate.ZRtcOfr, currencyExchangeRate.ZRthBid, currencyExchangeRate.ZRtcBid, currencyExchangeRate.Source))

	ctx := context.Background()
	lastExchangeRate, err = dal.TableCurrencyExchangeRate.GetCurrencyExchangeRateOne(ctx, currencyExchangeRate.Source, currencyExchangeRate.Name, "timestamp", "desc", time.Time{}, time.Now())
	if err != nil {
		return err
	} else {
		if lastExchangeRate != nil && lastExchangeRate.Hash == currencyExchangeRate.Hash {
			log.Debugf("%s exchange rate not change, skip", currencyExchangeRate.Name)
			return nil
		}
	}

	log.Debugf("%s exchange rate change, insert", currencyExchangeRate.Name)
	log.Debugf("%s", currencyExchangeRate.ToString())
	err = dal.TableCurrencyExchangeRate.CreateCurrencyExchangeRateOne(ctx, currencyExchangeRate)
	return err
}
