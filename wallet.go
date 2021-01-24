package androidlib

import (
	"encoding/json"
	"errors"
	"github.com/giantliao/beatles-client-lib/clientwallet"
	"github.com/giantliao/beatles-client-lib/coin"
	"github.com/giantliao/beatles-client-lib/config"
	"github.com/kprc/libeth/wallet"
	"math/big"
)

var beetleInitialErr = errors.New("beetle is not initialized")

func IsWalletCreate() bool {
	return clientwallet.IsWalletCreate()
}

func IsWalletOpen() bool {
	if _,err:=clientwallet.GetWallet();err!=nil{
		return false
	}

	return true
}



func OpenWallet(auth string) error {

	if !beetleIsInit(){
		return beetleInitialErr
	}

	if err:=clientwallet.LoadWallet(auth);err!=nil{
		return err
	}

	return nil
}

type WalletInfo struct {
	EthAddr string
	BeetleAddr string
	TrxAddr string
}

func GetWalletInfo() (string,error) {
	if !IsWalletOpen(){
		return "",errors.New("wallet not opened")
	}

	w,_:=clientwallet.GetWallet()

	wi:=&WalletInfo{
		EthAddr: w.AccountString(),
		BeetleAddr: w.BtlAddress().String(),
		TrxAddr: "",
	}

	j,_:=json.Marshal(*wi)

	return string(j),nil
}

type BeetleBalance struct {
	Eth float64
	BtlcGas float64
	Btlc float64
}

func Balance() (string,error) {
	if !IsWalletOpen(){
		return "",errors.New("wallet not opened")
	}

	w,_:=clientwallet.GetWallet()

	b:=&BeetleBalance{}
	var err error
	b.Eth,err = w.BalanceOf(true)
	if err!=nil{
		return "",err
	}
	b.BtlcGas, err=w.BalanceOfGas(config.GetCBtlc().BTLCAccessPoint)
	if err!=nil{
		return "",err
	}

	var btlc *big.Int

	btlc, err = coin.GetBTLCoinToken().BtlCoinBalance(w.Address())
	if err!=nil{
		return "",err
	}
	b.Btlc = wallet.BalanceHuman(btlc)

	j,_:=json.Marshal(*b)

	return string(j),nil
}




