package exchange

import (
	"time"
	"github.com/astaxie/beego/orm"
	"os"
)

type EthToken struct {
	Id            int64
	TokenName     string    `orm:"unique"`
	TokenFullName string
	ContractAddr  string    `orm:"unique"`
	TokenDecimal  int64
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	ContractAbi   string    `orm:"type(text)"`
	GasLimit      string
	Cny           string
	Usdt          string
	Updated       time.Time `orm:"auto_now;type(datetime)"`
}

func (e *EthToken) UpdateCnyAndUsdt() *EthToken {
	o := orm.NewOrm()
	os.Unsetenv(databases)
	qs := o.QueryTable(e)
	p := orm.Params{
		"cny":     e.Cny,
		"usdt":    e.Usdt,
		"updated": time.Now(),
	}
	qs.Filter("token_name", e.TokenName).Update(p)
	return e
}
