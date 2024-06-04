package bank

import (
	"fmt"
	"strconv"

	"PilotJobService/utils"
	"PilotJobService/utils/log"
)

type CmbApi struct {
	ExchangeRateUrl string
}

func (c *CmbApi) GetExchangeRateList() ([]*CurrencyExchangeRateCmb, error) {
	var (
		resp            map[string]interface{}
		err             error
		params, headers map[string]string
		result          []*CurrencyExchangeRateCmb
	)
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	headers = map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36",
		"Accept":     "application/json, text/plain, */*",
	}
	resp, err = utils.HttpRequest("GET", c.ExchangeRateUrl, "", params, headers, true)
	log.Debugf("cmb response: %+v", resp)
	if err != nil {
		return nil, fmt.Errorf("send request error: %v", err)
	}
	returnCode, ok := resp["returnCode"]
	if !ok || returnCode.(string) != "SUC0000" {
		return nil, fmt.Errorf("server return error: %+v", resp)
	}
	body, ok := resp["body"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("server return error: %+v", resp)
	}
	for _, data := range body["data"].([]interface{}) {
		ZCcyNbr, _ := data.(map[string]interface{})["ccyNbr"].(string)
		ZRtbBid, _ := strconv.ParseFloat(data.(map[string]interface{})["rtbBid"].(string), 64)
		ZRthOfr, _ := strconv.ParseFloat(data.(map[string]interface{})["rthOfr"].(string), 64)
		ZRtcOfr, _ := strconv.ParseFloat(data.(map[string]interface{})["rtcOfr"].(string), 64)
		ZRthBid, _ := strconv.ParseFloat(data.(map[string]interface{})["rthBid"].(string), 64)
		ZRtcBid, _ := strconv.ParseFloat(data.(map[string]interface{})["rtcBid"].(string), 64)
		ZRatTim, _ := data.(map[string]interface{})["ratTim"].(string)
		ZRatDat, _ := data.(map[string]interface{})["ratDat"].(string)
		result = append(result, &CurrencyExchangeRateCmb{
			ZCcyNbr: ZCcyNbr,
			ZRtbBid: ZRtbBid,
			ZRthOfr: ZRthOfr,
			ZRtcOfr: ZRtcOfr,
			ZRthBid: ZRthBid,
			ZRtcBid: ZRtcBid,
			ZRatTim: ZRatTim,
			ZRatDat: ZRatDat,
		})
	}
	return result, nil
}
