package accesscontrol

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	config2 "studyRoomGo/common/config"
)

var receiveAddr = net.IPv4(0, 0, 0, 0)
var sendAddr *net.UDPAddr

type Option func(o *Options)

type Server struct {
	Option     Options
	initServer *net.UDPConn
}

type Options struct {
	Key string
}

func WithKey(key string) Option {
	return func(o *Options) {
		o.Key = key
	}
}
func NewServer(opts ...Option) *Server {
	options := Options{
		Key: "accesscontrol",
	}
	for _, o := range opts {
		o(&options)
	}
	return &Server{
		Option: options,
	}
}

func (s *Server) GetKey() string {
	return s.Option.Key
}
func (s *Server) Setup() {

	// 监听当前的tcp连接
	server, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   receiveAddr,
		Port: 61005,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	s.initServer = server
	go func() {
		for {
			var data [1024]byte
			_, addr, err := s.initServer.ReadFromUDP(data[:]) // 接收数据
			if err != nil {
				fmt.Println("read udp failed, err:", err)
				continue
			}
			//hexStr := hex.EncodeToString(data[:n])
			sendAddr = addr
			//fmt.Printf("data:%v addr:%v count:%v\n", hexStr, sendAddr, n)
		}
	}()
}

func (s *Server) Close() {
	err := s.initServer.Close()
	if err != nil {
		fmt.Println(err.Error())

		return
	}
	fmt.Println("accesscontrol exit!")
}

func (s *Server) Automate(mac string, code string, status int) (ok bool, err error) {
	str := "17400000%s%s00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	hexStr := fmt.Sprintf(str, mac, code)
	decoded, _ := hex.DecodeString(hexStr)
	fmt.Println(hexStr)
	if config2.AppConf.Environment == "pro" {
		fmt.Println("send open door...")

		_, err = s.initServer.WriteToUDP(decoded, sendAddr) // 发送数据
		if err != nil {
			fmt.Println(err)
			return false, errors.New("open door error")
		}
	} else {
		fmt.Println("not send open door...")
	}

	return true, nil
}
func (s *Server) AutomateMain(mac string, code []string, status int) (ok bool, err error) {
	return
}
