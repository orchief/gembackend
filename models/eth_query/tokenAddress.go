package eth_query

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type TokenAddress struct {
	Id              int64
	WalletId        string    `orm:"index"`
	Addr            string    `orm:"index"`
	Created         time.Time `orm:"auto_now_add;type(datetime)"`
	ContractAddr    string    `orm:"index"`
	Amount          string
	UnconfirmAmount string
	Updated         time.Time `orm:"auto_now;type(datetime)"`
	Decimal         int64     `orm:"default(18)"`
	TokenName       string
	Added           int
}

func (u *TokenAddress) TableUnique() [][]string {
	return [][]string{
		{"Addr", "ContractAddr"},
	}
}

func (address *TokenAddress) Update(s string) *TokenAddress {
	o := orm.NewOrm()
	o.Using(databases)
	qs := o.QueryTable(address)
	p := orm.Params{
		"amount":           address.Amount,
		"unconfirm_amount": address.UnconfirmAmount,
		"updated":          time.Now(),
		"decimal":          address.Decimal}

	qs.Filter("addr", s).Filter("contract_addr", address.ContractAddr).Update(p)
	return address
}

func (address *TokenAddress) UpdateAmount(s string) *TokenAddress {
	o := orm.NewOrm()
	o.Using(databases)
	qs := o.QueryTable(address)
	p := orm.Params{
		"amount":           address.Amount,
		"unconfirm_amount": address.UnconfirmAmount,
		"updated":          time.Now(),
	}

	qs.Filter("addr", s).Filter("contract_addr", address.ContractAddr).Update(p)
	return address
}

func (Self *TokenAddress) InsertOneRaw(data *TokenAddress) *TokenAddress {
	o := orm.NewOrm()
	o.Using(databases)
	data.Id = 0
	data.Created = time.Now()
	id, err := o.Insert(data)
	if err != nil {
		log.Errorf("TokenAddress insert error %s", err)
	}
	log.Debugf("TokenAddress insert id %d ", id)
	return Self
}
//func (u *TokenAddress) TableEngine() string {
//	return "MYISAM"
//}
