package tun2Pipe

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type ByPassIPs struct {
	Masks map[string]net.IPMask
	IP    map[string]struct{}
	sync.RWMutex
	global bool
}

var _instance *ByPassIPs
var once sync.Once

func ByPassInst() *ByPassIPs {
	once.Do(func() {
		_instance = &ByPassIPs{
			Masks: make(map[string]net.IPMask),
			IP:    make(map[string]struct{}),
		}
	})
	return _instance
}

func (bp *ByPassIPs) Load(IPS string) {
	bp.IP = make(map[string]struct{})
	bp.Masks = make(map[string]net.IPMask)
	array := strings.Split(IPS, "\n")
	for _, cidr := range array {
		ip, subNet, _ := net.ParseCIDR(cidr)
		bp.IP[ip.String()] = struct{}{}
		bp.Masks[subNet.Mask.String()] = subNet.Mask
	}
	log.Println(fmt.Sprintf("====By Pass===>Total bypass ips:%d groups:%d \n", len(bp.IP), len(bp.Masks)))
}

func (bp *ByPassIPs) Hit(ip net.IP) bool {

	bp.RLock()
	defer bp.RUnlock()

	if bp.global {
		return false
	}

	for _, mask := range bp.Masks {
		maskIP := ip.Mask(mask)
		if _, ok := bp.IP[maskIP.String()]; ok {
			log.Println(fmt.Sprintf("\nHit success ip:%s->ip mask:%s", ip, maskIP))
			return true
		}
	}

	//TODO:: pac domain list

	return false
}

func (bp *ByPassIPs) SetGlobal(g bool) {
	bp.Lock()
	defer bp.Unlock()
	bp.global = g
}

func (bp *ByPassIPs) IsGlobal() bool {
	bp.RLock()
	defer bp.RUnlock()
	return bp.global
}
