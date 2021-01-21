package androidlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/giantliao/androidlib/tun2Pipe"
	"github.com/giantliao/beatles-client-lib/app/cmd"
	"github.com/giantliao/beatles-client-lib/bootstrap"
	"github.com/giantliao/beatles-client-lib/config"
	"github.com/giantliao/beatles-client-lib/resource/pacserver"
	"github.com/giantliao/beatles-client-lib/streamserver"
	"io"
	"log"
	"sync"
)

type Beetle struct {
	CurrentMiner string
	VpnMode int
	BasDir string
}

var ProtectFD func(fd int32) bool
var ListenSock io.Writer
var TunInst *tun2Pipe.Tun2Pipe = nil


var beetleInstance *Beetle
var beetleInstLock sync.Mutex

func beetleIsInit() bool {
	if beetleInstance == nil{
		return false
	}
	return true
}

func InitBeetle(basdir string,bypassIPs string) bool {
	if beetleInstance != nil{
		return true
	}

	beetleInstLock.Lock()
	defer beetleInstLock.Unlock()

	if beetleInstance != nil{
		return true
	}

	beetleInstance = &Beetle{
		CurrentMiner: "",
		VpnMode: 0,
		BasDir: basdir,
	}

	config.SetHomeDir(basdir)

	cmd.InitCfg()
	cfg:=config.GetCBtlc()
	cfg.Save()

	tun2Pipe.ByPassInst().Load(bypassIPs)

	return true
}

func LoadBypassIPs(bypassIPs string)  {
	tun2Pipe.ByPassInst().Load(bypassIPs)
}

func StartBeetle() error {
	if !beetleIsInit(){
		return beetleInitialErr
	}
	if !IsWalletOpen(){
		return errors.New("wallet is not open")
	}

	cfg:=config.GetCBtlc()
	if len(cfg.Miners) == 0{
		err := bootstrap.UpdateBootstrap()
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}
	if len(cfg.Miners) == 0 {
		return errors.New("no miner to start vpn")
	}

	go pacserver.StartWebDaemon()

	log.Println("start Beetle success")

	return nil

}

func StopBeetle() error {
	if !beetleIsInit(){
		return beetleInitialErr
	}

	if !IsWalletOpen(){
		return errors.New("wallet is not open")
	}

	cfg:=config.GetCBtlc()

	cfg.Save()

	pacserver.StopWebDaemon()

	//setting.ClearProxy()

	return nil
}

func StartVpn(minerId string) error  {
	if !beetleIsInit(){
		return beetleInitialErr
	}

	if !IsWalletOpen(){
		return errors.New("wallet is not open")
	}

	if !pacserver.WebDaemonIsStarted(){
		return errors.New("beetle is not started")
	}

	if streamserver.StreamServerIsStart(){
		return errors.New("stream server is started")
	}

	tun2Pipe.VpnInstance = ListenSock

	cfg := config.GetCBtlc()

	var idx = -1

	for i:= 0;i<len(cfg.Miners);i++{
		if cfg.Miners[i].MinerId.String() == minerId{
			idx = i
			break
		}
	}

	if idx == -1{
		return errors.New("not find miners")
	}

	srvAddr:=fmt.Sprintf("%s:%d",cfg.Miners[idx].Ipv4Addr,cfg.Miners[idx].Port)

	t2s,err:=tun2Pipe.New(srvAddr, func(fd uintptr) {
		ProtectFD(int32(fd))
	})

	if err!=nil{
		return err
	}

	cfg.CurrentMiner = cfg.Miners[idx].MinerId

	cfg.Save()

	go streamserver.StartStreamServer(idx,ProtectFD,t2s.GetTarget,t2s.ProxyClose)

	go t2s.Proxying(nil)

	TunInst = t2s

	return nil
}

func StopVpn() error {
	if !beetleIsInit(){
		return beetleInitialErr
	}

	if !IsWalletOpen(){
		return errors.New("wallet is not open")
	}

	if !pacserver.WebDaemonIsStarted(){
		return errors.New("beetle is not started")
	}

	if !streamserver.StreamServerIsStart(){
		return errors.New("stream server is started")
	}

	streamserver.StopStreamserver()

	if TunInst != nil{
		TunInst.Finish()
		TunInst = nil
	}

	return nil
}

func ListAllMiner() ([]string,error) {
	if !beetleIsInit(){
		return nil,beetleInitialErr
	}

	cfg:=config.GetCBtlc()

	var ms []string

	for i:=0;i<len(cfg.Miners);i++{
		j,_:=json.Marshal(cfg.Miners[i])
		ms = append(ms,string(j))
	}

	return ms,nil
}

func VpnIsStarted() bool  {
	if !beetleIsInit() || !IsWalletOpen() || !pacserver.WebDaemonIsStarted() ||	!streamserver.StreamServerIsStart(){
		return false
	}

	return true
}

func InputPacket(data []byte)  error{
	if TunInst == nil {
		return fmt.Errorf("Tun2Proxy has stopped")
	}
	TunInst.InputPacket(data)
	return nil
}

func SetGlobalModel(g bool) {
	tun2Pipe.ByPassInst().SetGlobal(g)
}

func IsGlobalMode() bool {
	return tun2Pipe.ByPassInst().IsGlobal()
}