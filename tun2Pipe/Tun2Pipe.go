package tun2Pipe

import (
	"errors"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"io"
	"log"
	"math"
	"net"
	"strconv"
	"sync"
	"time"
)


/*
    rc!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\lv.                   \[!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!x~
     j!                                                                  !x`                   ii                                                                                                                                                                                                        [i.
     j!                            application                           !f`                   ii                                                                                   vpn program                                                                                                          [i.
     j!                .   ....   ..     .. .     .......                !e`                   ii                                                                                                     ...   ...  .                                                                                       [i.
     c!               .._~-!__:.~~_!,|;.:i~!_:|.;~;\!i!..               .!x`                   ii . .                                                                           ..   '~!~`  ...  .   .:"..;;\\"`!;._...                                                                               . .[i.
     c\...          ....!r/|;t;.ri;e:]!."Lis\~l`!|-\j\_..            ....!e`                   ii . ..                                                                        .......ur"]o\Irj;xfi"}v?;u!ixvT[L|qbtz,  ..                                                                            ....[i.
     c\.................\;/!i~,.r\"x_I\./u;_\!;.\_`~[:_..................!e`                   it...................................................................................`e[ifllb?u|Tf\!];u~~Lu:|n!:_u!f6:....................................................................................[i.
     c\................`_,-,___-_.-!:|_..~-,_:_,:..""....................\e`                   it.....................................................................................:_,.!v_-.:_,.'.,..:-..-...,..:.....................................................................................[i.
     c\..................................................................\x`                   it........................................................................................................................................................................................................]i.
     j\..................................................................\x`                   ir........................................................................................................................................................................................................]i.
     j|..................................................................|e`                   ir........................................................................................................................................................................................................]i.
     r[iiiiiiiiiiri!`..!iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiljt|iiiiiiiiiiiii[j.                   \ziiiiiiiiiiiiti!`.`!iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiicjtiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii|it[jjslliiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii||iiiiiiiiiiiiiiiiiiiiiiiiiiiy|.
      .'''''''''']!.  .i_,-'''''''''''''''''''''''''-,~xTYLr_--'--''''''''..                    .`-''''''''''j\.  .!;---''''''''''''''''''''''''---,"zkCni:-,,''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''`'-|Z8T25n;---'''''''''''''''''''''----,;:.."?~--`'---'''''''''''''''-''.
                 c/   ,e-                          .~[oun6UZ$j:                                              r!   -f;                            .;syoL6SZYt`                                                                        /Z8Luu5:                          .~i~.   .!i_
                 j!   ,x,                        ."tljzonbUZO@Zz:                                            t\   `]".                          ;tll[oLdS%O@Sv`                                                                      /Z8Luu5:                        `!i".       ,||_
                 "_   ,x:                       .ifj[y5onbUZFCUUo"                                           __   `["                         .\}c[InoLdSZqCUgI_.                                                                    /Z8Luu5:                       ,l|..`.    ;?||sx;
                      ,x:                         ...!LuLbUP|`....                                                .[~.                          ...;LoLdSC;.....                                                                     /Z8Luu5:                            's.   il....
    . .    .~`  ..,'  .~:.  .;'-.                    ~LonkSP|.                                               ..   .;_.                             _LoLdSC_                                                                          !ZFLuun_     .:-:                   .}.   \t.
   tu|e}}zc|bntonZkb_,f]ui}oixjCtJ5f.                "LunkUPi.                                           .   !|    .,.    .                        _LoLdSC;                                                            ,[|;:;;/!_  .;rmPFY6Zzxi_!"\L!T"                  .}.   ._.
   !n]yIxyz|auj55Cau: shxLix:z|aiuu?.                ;LunkUPt.                                  _~`_~:-~~e\;:[]vi..!~;,;""ti[":~_                  _LoLdSC;                                                            -Y[oy8bgnL.  rZWDNDU@6UxLPwn5,wi                  .}.   .:.
                 j!   ,x'                            ;nunb8Pr.                                 `oC]woqToikPaTZ6oy`:$UT6aL!uJCrUUn-                 _LoLdSC_                                                             _;_/~__:~iti_s%ZkT5T/_":~""I!J:                  .?.   !i
                 c/   ,x-                            _nunb$Zv.                                 ._"'_":,":__;:s[:;.-[j_;.:.:-__:":                  _LoLdSC;.                                                                         !ZgTuu5:                             .    iv
                 \;   :UL'                         't!2nZ@K&L,                                             .![utiJ%dI,                            :aOZphS%p"                                                                        .\Z8b6wn:                            ..  ,~ej.
              .jUTYFYd$Qm,                         "ml5PQZZNDf                                             _ZN6oI6Q%n_                           `g%ZUNDON@!                                                                       ;k%Q@DNNN|                           ?SZT58LZm,
               ;Tz,`..'LU,                         "NSED@q8%@o.                                             .|o5yLNs".                           .gYToLPDQQ!                                                                       o8t%WZ55W2.                         ,PoJ|mr rW!
                .c"    :;.                         ._|TTLw8Zn_                                               :i."]]_:.                            _jwoL6SZf,        .                                                              \TjZZ65LNv                          .o5J-~- !S_    ..:,:
                 j/    _'                         ;. :onL68Zf`        ';                                     :j.   -_.                `"::~"';~]s":tFT%8PPr;"~-_/_:!n|                                                              .!Z$L5un_                 _[";,;_~/:.:x;~_"|!/~-;inyx|L:
                 c/   ,x'               .?u|fcocj/Unrnm%EDP%Z]t2uisLloLp'                                    ,j.   tt                 \kuTJYLTvnZnxm@Z@SP@u[ZTjuZgT5Sr                                                               /Z8Luu5:                 _CtLoZouuu..[yzl%SLL%LhuPgz;L;
                 c/   ,x-                !ojIlol[!oaln%PFOP%Z!!5ftnT[oIa.                                    -[.   tt.                -;:,"_.;_::_,|C5qpPS~._;'_"_,;"`                                                               /Z8Luu5:                  _,:;;::-"\|t"'.;}j;"::;;\j|[.
                 t"   ,x:                            .ynnkgZy-                                               .s.   is.                             _LoLdSC;                                                                          /Z8Luu5:                            ..    is
                 ..   ,x:                            .annwgZy-                                                .    il.                             _LoLdSC;                                                                          /Z8Luu5:                            .~.   ic.
            .'---.    `t:  ``                        .In2TgZu,                                           .,-''.    ~|.  `.                         _LoLdSC;                                                                          /Z8L5u5:                            .}.   il.
           .iov!!_     .  .yc.                       .xnnTgZu,                                          .tot!/-      ..;yt.                        _LoLdSC;                                                                     .    !ZgL5un:                            .}.   _".
             .~i!.      :tl;.                         xnnTFZ5,                                            '\i,      .!v|:                          _LoLdSC;                                                                     !6XX%Z%FTuoux[zJ"                        .f.   .:.
               ."i\`  .|l~,                           fLnLgZn,                                              'i|:   "v\,.                           _LoLdSC;.                             ..',`".   ... '|"i'                     -cZNmZ8L5Ijcc|_.                        .z.   _~.
     ,"""""""""""~vv;;~\!~~"""""""""""""""""""""""""""ukTCS%L~"""""""""""""""""""""""""""""""""""""""""""""""~i]"_;/!~"""""""""""""""""""""""""""""i6L6gZ$!""""""""""""":      `. ~}]itxlsLvzoS.  !j\[cfL:Li                       -sCZ8L2ofi_.                           _    iv
     lr;;;;;;;;;;;;;__;_;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;"/!!!~;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;_____;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;~!/!~;;;;;;;;;;;;;\f.   `sUt.ioxLT6o|njuca   lj~6SuL-nt  ......................._lwwur;,..................................rl.........................
     c!                                                                                                                                                                _x: .iC@Dj......:'.;[c:,|i!:...`_v/v, .z\!!!!!!!!!!!!!!!!!!!!!!!|t!~!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!|iiiri!!!!!!!!!!!!!!!!!!!!!\\\x~
     c!                                                                                                                                                                _x:_L%XOXL[llllllleZMRylllllllslsccsll[u.                                                                                         ji.
     c!                                                                                                                                                                _uIh8UUUUSSSSSSSSmQNNQ%SUSSSSSSSSSSUUUPb:.                                                                                        ji.
     c!                                                                tcp/ip stack                                                                                    _TLL2nn2n2nnnnnnZMDOKMN6nnnnnnnnnnnn2nLL_.                       tun dev                                                          [i.
     c!                                                                               ..            .                                                                  _x"taf[[uu[[[[[[yLLL%Wha][[[[[[[[[[[[[}o_.                                            .                                           [i.
     c!                                                                               .....         .......                                                            ;x:.;rciaj`       ..-;....         ...']`.                                            ...    ..              . ..                 [i.
     c!.........................................................::_`.,'.'_,...",..::..~|`|",..-;._;":.,i-"|\;..........................................................;x,   _iol.                           .f`.. ..... ...-:::``....``....,":::'..'..':'..:"i_;_.-|!~~!\.:"::-_/...`r-,-  .............[i.
     c\........................................................:?bz~}ti!uxyv.|xs[ihc5;if/sti-."!:~!j:.~L\i[!|..........................................................;x,    .__.                           .f`............"aCii5-_[;Lb;o,`zr\da~,fn",L}ot_j|]tr|._xsl~|v.~rsr\[?.',_o";/-..............[i.
     c\........................................................._L!_y:-;Tos";y;lJtPcr`\]i||j/.~j;iu!.`|u";oj"..........................................................;x,                                   .f`.............tT-!L_~o;T]uF:r[..Li.[noL"Yac;;s|?i\;.~}!f|li'~vi||[?...-x!i;...............[i.
     c\.........................................................-i_.|si;|;.-f!."!;t..."i;:|;.._t/~"i/.:i:;/it..........................................................;x,                           ;t;.    .e'.............;t..|vv!`t;\z!e:..i;,i_,i/t_..;""~/""._t_~!~!;;/...;i...'i'.................[i.
     c\.....................................................................`..........................................................................................;x,              ."|\:         ."i/.  .f-.........................................................................................[i.
     c\................................................................................................................................................................;I,   :~~//~~///_[5u6Na!!~~~~/;. ."t!``x-.........................................................................................[i.
     c\................................................................................................................................................................;z.   .,:::,,,,,.   !Na_:,,,,,`    ."_-x-.........................................................................................]i.
     c|................................................................................................................................................................"[.               ,[Cc..             .:v:.........................................................................................]i.
     j|................................................................................................................................................................"u||"'..._!|||||in@m}r|,/|||||t:   `\t~[:.........................................................................................zi.
     j|................................................................................................................................................................"J_...  ........:juuuu[-......[/ :|v!-,a;:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::_:xi.
     j|..................................................``.........................................`.`.`.............................................................,/I,      .`.             ..   _"ir~,...|iiizxIiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiit;.
     "itiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiittor|!:-irttttiiiiiiiiiiiiiiiiiiiiiiiiiiiiiititttvodkIritiititiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiittitk%w|!z!l|t}%|  'virL~\z|t|n[.  !ne|lljliztLt6\L"
                                                  .[_   `".                                     `!ILTCFy".                                                           .[yuunZyL"uuZ|  :uLc6tLZLL_ux.  ~gYjojTcnniksF!L/
                                                   [_   "}.                                   `!jIJuLCPE%u".                                                          ....',.' .-:"|i!:`.'.-,.` .:/ii~.```.'.``.''t|j.
                                                   ;`   "}.                                 :|rrv[IuLCPEN@Nn"
                                                        "f`                                .!rtrvnn5LqZ$IexIv.
                                                   -.   "e'                                      ]L5LqZT:
                                                   j_   _r-                                      ]LuLCPT:
                                                 .'[_./sv|:                                      [L5LCZSr.
                                                 oTf"LZIo%5.                                   "eUE%mN@ME-
                                                 LEL[Zt ;FZ_                                   !m@PUZENQm,
                                                 |2L5a; "Zu.                                    ;CTuLCZZf.
                                                        "x`                                      ]LuLCZT:
                                                   .    "e-                                      ]LuLCZT_
                                                  .l_   _r-                                      [LuLCZT_
                                              .sr\|t`    .. ,[],                                 [LuLCZT_
                                               .i|,       ."st_.                                 [LuLCZT_
                                                 :i|:    "st;.                                   [LuLCZT_
                                                  ._i!..:/_`.                                    ]LuLCZT_
                                 ;j!!!!!!!!!!!!!!!!!!\~~//\\!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!v[}xJoz!!!!!!!!!!!!!!!!!!!!!j;
                                 !j                                                                ......                    [i.
                                 !j                                                                                          jt.
                                 /j                                real network device                                       jt.
                                 /j                                                                                          ji.
                                 /j                           ..   .. . ..     . .   ... ..  ..                              [i.
                                 /j .......................... .-!;~-...._:_:,_.:_'``:"...:~................................ [i.
                                 /j.............................io]llss`:e~?v~i.rvr\!te'.-/["~/..............................[i.
                                 /j............................-"a?!r|t.:y~~r"".rsj|scx`.."[\t;..............................[i.
                                 /[..............................i!;~s"`~~_;!""-|/...~j..._\.................................[i.
                                 /[..........................................................................................[i.
                                 /[..........................................................................................[i.
                                 ![..........................................................................................[i.
                                 ![..........................................................................................[i.
                                 ![..........................................................................................[i.
                                 ![````````````````.......``````````````````````````````````````...``````````````````````````[i.
                                 -|||iiiiiiiiiiiiiii`  `t5iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiivuphyriiiiiiiiiiiiiiiiiiiiit;.
                                                   [_`..!e.                                     -|aLLCU5!.
                                                   ]rnrvIa'                                   -\jeJuLCP%%L!.
                                                   [!ttzox-                                 _is?s}uqP%%END@T!
                                               .  .}~!|vj|` .::.                           .~iittnb%%EmF[jjjt.
                                              .!~..;_?sv~.  ;ye-                                 IqFZEmT:
                                               .~i~.,rtv!.,il!`                                  ]L5LCZT_
                                                  _i|-  .|r;`                                    ]LuLCZT_
                                                    _~.`,:.                                      [6TdSZL_
                                                       ..                                        .;""""_.



*/



const (
	SysDialTimeOut    = time.Second * 2
	UDPSessionTimeOut = time.Second * 80
	MTU               = math.MaxInt16
	InnerPivotPort    = 51414 //TODO:: this port should be refactored
)

type Tun2Pipe struct {
	sync.RWMutex
	innerTcpPivot *net.TCPListener
	SessionCache  map[int]*Session
	udpProxy      *UdpProxy
	youPipeProxy  int
	tunIP         net.IP
}

type VpnDelegate interface {
	io.Writer
}

var VpnInstance VpnDelegate = nil

type ConnProtect func(fd uintptr)

var Protector ConnProtect

func New(proxyAddr string, p func(fd uintptr)) (*Tun2Pipe, error) {

	l, e := net.ListenTCP("tcp", &net.TCPAddr{
		Port: InnerPivotPort,
	})
	if e != nil {
		return nil, e
	}

	ip, port, _ := net.SplitHostPort(proxyAddr)
	intPort, _ := strconv.Atoi(port)

	tsc := &Tun2Pipe{
		innerTcpPivot: l,
		SessionCache:  make(map[int]*Session),
		udpProxy:      NewUdpProxy(),
		tunIP:         net.ParseIP(ip),
		youPipeProxy:  intPort,
	}
	Protector = p
	return tsc, nil
}

//TODO::
func (t2s *Tun2Pipe) Proxying(done chan struct{}) {

	for {
		conn, err := t2s.innerTcpPivot.Accept()
		if err != nil {
			log.Println(fmt.Sprintln("Tun2proxy inner pivot accept:", err))
			return
		}

		go t2s.simpleForward(conn)
	}
}

func (t2s *Tun2Pipe) ProxyClose(conn net.Conn) {
	rAddr := conn.RemoteAddr().String()
	_, port, _ := net.SplitHostPort(rAddr)
	keyPort, _ := strconv.Atoi(port)
	t2s.removeSession(keyPort)
}
func (t2s *Tun2Pipe) GetTarget(conn net.Conn) (string,error) {
	keyPort := conn.RemoteAddr().(*net.TCPAddr).Port
	s := t2s.getSession(keyPort)
	if s == nil {
		return "",errors.New("not found session")
	}

	if len(s.HostName) != 0 {
		return fmt.Sprintf("%s:%d", s.HostName, s.RemotePort),nil
	}

	addr := &net.TCPAddr{
		IP:   s.RemoteIP,
		Port: s.RemotePort,
	}

	return addr.String(),nil
}

func (t2s *Tun2Pipe) InputPacket(buf []byte) {

	var ip4 *layers.IPv4 = nil
	packet := gopacket.NewPacket(buf, layers.LayerTypeIPv4, gopacket.Default)

	//VpnInstance.VpnLog(packet.Dump())

	if ip4Layer := packet.Layer(layers.LayerTypeIPv4); ip4Layer != nil {
		ip4 = ip4Layer.(*layers.IPv4)
	} else {
		log.Println("Unsupported network layer :")
		return
	}

	var tcp *layers.TCP = nil
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		tcp = tcpLayer.(*layers.TCP)
		t2s.ProcessTcpPacket(ip4, tcp)
		return
	}

	var udp *layers.UDP = nil
	if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
		udp = udpLayer.(*layers.UDP)
		t2s.udpProxy.ReceivePacket(ip4, udp)
		return
	}

	log.Println(fmt.Sprintf("Unsupported transport layer :%s", ip4.Protocol.String()))
}

func (t2s *Tun2Pipe) Finish() {
	t2s.Lock()
	defer t2s.Unlock()

	if t2s.innerTcpPivot != nil {
		t2s.innerTcpPivot.Close()
		t2s.innerTcpPivot = nil
	}

	if t2s.udpProxy.Done != nil {
		t2s.udpProxy.Done <- fmt.Errorf("finished by outer controller")
		t2s.udpProxy.Done = nil
	}
}

func (t2s *Tun2Pipe) tun2Proxy(ip4 *layers.IPv4, tcp *layers.TCP) {

	//PrintFlow("-=->tun2Proxy", ip4, tcp)
	srcPort := int(tcp.SrcPort)
	s := t2s.getSession(srcPort)

	if s == nil {
		var srvPort = InnerPivotPort
		bypass := ByPassInst().Hit(ip4.DstIP)
		if !bypass {
			srvPort = t2s.youPipeProxy
			log.Println(fmt.Sprintln("This session will be proxy:", ip4.DstIP, tcp.DstPort, srcPort))
		}

		s = newSession(ip4, tcp, srvPort, bypass)
		t2s.addSession(srcPort, s)
	}

	tcpLen := len(tcp.Payload)

	s.UPTime = time.Now()
	s.packetSent++
	if s.packetSent == 2 && tcpLen == 0 {
		log.Println(fmt.Sprintf("Discard the ack:%t\n syn:%t psh:%t", tcp.ACK, tcp.SYN, tcp.PSH))
		return
	}

	if s.byteSent == 0 && tcpLen > 10 {
		host := ParseHost(tcp.Payload)
		if len(host) > 0 {
			log.Println(fmt.Sprintln("Session host success:", host))
			s.HostName = host
		}
	}

	ip4.SrcIP = ip4.DstIP
	ip4.DstIP = t2s.tunIP
	tcp.DstPort = layers.TCPPort(s.ServerPort)

	data := ChangePacket(ip4, tcp)
	//PrintFlow("-=->tun2Proxy", ip4, tcp)

	if _, err := VpnInstance.Write(data); err != nil {
		log.Println(fmt.Sprintln("-=->tun2Proxy write to tun err:", err))
		return
	}
	s.byteSent += tcpLen
}

func (t2s *Tun2Pipe) proxy2Tun(ip4 *layers.IPv4, tcp *layers.TCP, rPort int) {
	//PrintFlow("<-=-proxy2Tun", ip4, tcp)

	ip4.SrcIP = ip4.DstIP
	ip4.DstIP = t2s.tunIP
	tcp.SrcPort = layers.TCPPort(rPort)
	data := ChangePacket(ip4, tcp)

	//PrintFlow("<-=-proxy2Tun", ip4, tcp)
	if _, err := VpnInstance.Write(data); err != nil {
		log.Println(fmt.Sprintln("<-=-proxy2Tun write to tun err:", err))
		return
	}
}

func (t2s *Tun2Pipe) ProcessTcpPacket(ip4 *layers.IPv4, tcp *layers.TCP) {
	srcPort := int(tcp.SrcPort)
	if srcPort == InnerPivotPort ||
		srcPort == t2s.youPipeProxy {
		dstPort := int(tcp.DstPort)
		if s := t2s.getSession(dstPort); s != nil {
			t2s.proxy2Tun(ip4, tcp, s.RemotePort)
		}
		return
	}

	t2s.tun2Proxy(ip4, tcp)
}

func (t2s *Tun2Pipe) getSession(key int) *Session {
	t2s.RLock()
	defer t2s.RUnlock()
	return t2s.SessionCache[key]
}
func (t2s *Tun2Pipe) addSession(portKey int, s *Session) {
	t2s.Lock()
	defer t2s.Unlock()
	t2s.SessionCache[portKey] = s
}

func (t2s *Tun2Pipe) removeSession(key int) {
	t2s.Lock()
	defer t2s.Unlock()
	delete(t2s.SessionCache, key)
}
