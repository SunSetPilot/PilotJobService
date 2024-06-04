package table

import (
	"fmt"
	"time"
)

type CurrencyExchangeRateModel struct {
	Id        int       `gorm:"id" json:"id"`               // 主键
	Name      string    `gorm:"name" json:"name"`           // 货币名称
	ZRtbBid   float64   `gorm:"z_rtb_bid" json:"z_rtb_bid"` // 中行折算价
	ZRthOfr   float64   `gorm:"z_rth_ofr" json:"z_rth_ofr"` // 现汇卖出价
	ZRtcOfr   float64   `gorm:"z_rtc_ofr" json:"z_rtc_ofr"` // 现钞卖出价
	ZRthBid   float64   `gorm:"z_rth_bid" json:"z_rth_bid"` // 现汇买入价
	ZRtcBid   float64   `gorm:"z_rtc_bid" json:"z_rtc_bid"` // 现钞买入价
	Source    string    `gorm:"source" json:"source"`       // 数据来源
	Hash      string    `gorm:"hash" json:"hash"`           // 数据哈希
	Timestamp time.Time `gorm:"timestamp" json:"timestamp"` // 时间戳
}

func (c *CurrencyExchangeRateModel) TableName() string {
	return "currency_exchange_rate"
}

func (c *CurrencyExchangeRateModel) ToString() string {
	return fmt.Sprintf(
		"货币名称: %s\n"+
			"中行折算价: %.2f\n"+
			"现汇卖出价: %.2f\n"+
			"现钞卖出价: %.2f\n"+
			"现汇买入价: %.2f\n"+
			"现钞买入价: %.2f\n"+
			"更新时间: %s",
		c.Name, c.ZRtbBid, c.ZRthOfr, c.ZRtcOfr, c.ZRthBid, c.ZRtcBid, c.Timestamp.Format("2006-01-02 15:04:05"))
}
