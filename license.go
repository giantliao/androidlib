package androidlib

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/giantliao/beatles-client-lib/config"
	"github.com/giantliao/beatles-client-lib/db"
	"github.com/giantliao/beatles-client-lib/licenses"
	prolic "github.com/giantliao/beatles-protocol/licenses"
	"github.com/kprc/libeth/account"
)

func QueryPrice(month int, payType int, addr string) (string,error) {
	if !IsWalletOpen(){
		return "",errors.New("wallet not opened")
	}

	if month <= 0{
		return "",errors.New("month must large lan 1")
	}
	if payType != prolic.PayTypETH && payType != prolic.PayTypBTLC {
		return "",errors.New("pay type error")
	}

	if addr != "" {
		if !account.BeatleAddress(addr).IsValid() {
			return "",errors.New("not a correct receiver address")
		}
	}

	var cp *licenses.CurrentPrice
	var err error
	cp, err = licenses.NewCurrentPrice(int64(month), payType, account.BeatleAddress(addr))
	if err != nil {
		return "",err
	}

	np := cp.Get()
	if np == nil {
		return "",errors.New("get price failed")
	}

	config.GetCBtlc().MemPrice = np

	var j []byte
	j, err = json.Marshal(*np)
	if err != nil {
		return "",err
	}

	return string(j),nil
}

func Buy(name, email, cell string) (string,error)  {
	if !beetleIsInit(){
		return "",beetleInitialErr
	}

	cfg := config.GetCBtlc()
	if cfg.MemPrice == nil {
		return "",errors.New("please get price first")
	}

	lr := licenses.NewClientLicenseRenew(cfg.MemPrice, name, email, cell)

	err := lr.Buy()
	if err != nil {
		return "",err
	}

	tdb := db.GetClientTransactionDb()
	if v := tdb.Find(*lr.Transaction); v != nil {
		j,_:=json.Marshal(*lr)
		return string(j),nil
	} else {
		return "",errors.New("buy license info not in db")
	}

}

func LicenseRenew(tx string) (string,error) {
	if !beetleIsInit(){
		return "",beetleInitialErr
	}
	tdb := db.GetClientTransactionDb()
	var (
		cti *db.ClientTranstionItem
		err error
	)

	if tx == "" {
		cti, err = tdb.FindLatest()
		if err != nil {
			return "",err
		}
	} else {
		cti = tdb.Find(common.HexToHash(tx))
		if cti == nil {
			return "",errors.New("not found transaction")
		}
	}
	if cti.Used {
		return "",errors.New("transaction is used")
	}

	clr := licenses.NewClientLicenseRenew(cti.Price, cti.Name, cti.Email, cti.Cell)
	clr.Transaction = &cti.Tx

	l := clr.GetLicense()
	if l == nil {
		return "",errors.New("something wrong, get license failed")
	}

	j, _ := json.Marshal(*l)

	return string(j),nil
}


type TransactionItem struct {
	TxStr string  			`json:"tx_str"`
	*db.ClientTranstionItem
}

func LicenseLog() ([]string,error) {

	if !beetleIsInit(){
		return nil,beetleInitialErr
	}

	tdb := db.GetClientTransactionDb()
	cursor:=tdb.Iterator()

	var logs []string

	for {
		k, v, e := tdb.Next(cursor)
		if k == nil || e != nil {
			break
		}

		ti:=&TransactionItem{}
		ti.TxStr = k.String()
		ti.ClientTranstionItem = v

		j,_:=json.Marshal(*ti)
		logs = append(logs,string(j))
	}

	return logs,nil
}

func ShowLicense() (string,error) {
	if !beetleIsInit(){
		return "",beetleInitialErr
	}
	ldb:=db.GetClientLicenseDb()
	cli:=ldb.FindNewestLicense()
	if cli == nil{
		return "",errors.New("not found")
	}

	j,_:=json.Marshal(*cli.License)

	return string(j),nil

}

func RefreshLicense() error  {
	cfl := licenses.ClientFreshLicense{}

	if err := cfl.FreshLicense(); err != nil {
		return err
	}

	return nil

}