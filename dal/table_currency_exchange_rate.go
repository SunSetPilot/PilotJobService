package dal

import (
	"context"
	"strings"
	"time"

	"gorm.io/gorm/clause"

	"PilotJobService/model/table"
)

var TableCurrencyExchangeRate _TableCurrencyExchangeRate

type _TableCurrencyExchangeRate struct{}

func (*_TableCurrencyExchangeRate) GetCurrencyExchangeRateOne(ctx context.Context, source, name, orderBy, sort string, startTime, endTime time.Time) (*table.CurrencyExchangeRateModel, error) {
	currencyExchangeRate := new(table.CurrencyExchangeRateModel)
	var desc bool
	if strings.ToLower(sort) == "desc" {
		desc = true
	} else {
		desc = false
	}
	result := DB.NewRequest(ctx).Table(currencyExchangeRate.TableName()).
		Where("source = ? and name = ?", source, name).
		Where("timestamp between ? and ?", startTime, endTime).
		Order(clause.OrderByColumn{Column: clause.Column{Name: orderBy}, Desc: desc}).
		Limit(1).
		Find(currencyExchangeRate)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return currencyExchangeRate, nil
}

func (*_TableCurrencyExchangeRate) CreateCurrencyExchangeRateOne(ctx context.Context, data *table.CurrencyExchangeRateModel) error {
	err := DB.NewRequest(ctx).Table(data.TableName()).Create(data).Error
	return err
}

func (*_TableCurrencyExchangeRate) DeleteCurrencyExchangeRateByTime(ctx context.Context, before time.Time) error {
	data := new(table.CurrencyExchangeRateModel)
	err := DB.NewRequest(ctx).Table(data.TableName()).Where("timestamp < ?", before).Delete(data).Error
	return err
}
