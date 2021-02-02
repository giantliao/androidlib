package androidlib

import (
	"encoding/json"
	"errors"
	"github.com/giantliao/beatles-client-lib/clientwallet"
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



func GetWalletInfo() (string,error) {
	if !IsWalletOpen(){
		return "",errors.New("wallet not opened")
	}

	wi,err:=clientwallet.GetWalletInfo()

	if err!=nil{
		return "", err
	}

	j,_:=json.Marshal(*wi)

	return string(j),nil
}



func Balance() (string,error) {
	if !IsWalletOpen(){
		return "",errors.New("wallet not opened")
	}

	wb,err:=clientwallet.GetBalance()
	if err!=nil{
		return "", err
	}

	j,_:=json.Marshal(*wb)

	return string(j),nil
}




