package bank

type CurrencyExchangeRateCmb struct {
	// 货币名称
	ZCcyNbr string
	// 中行折算价
	ZRtbBid float64
	// 现汇卖出价
	ZRthOfr float64
	// 现钞卖出价
	ZRtcOfr float64
	// 现汇买入价
	ZRthBid float64
	// 现钞买入价
	ZRtcBid float64
	// 更新时间
	ZRatTim string
	// 更新日期
	ZRatDat string
}
