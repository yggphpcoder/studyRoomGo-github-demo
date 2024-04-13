package mqtt

// 与后端mqtt服务交互

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	config2 "studyRoomGo/common/config"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var ClientId = "" //配置mqtt 服务器信息
var Username = ""
var Password = ""
var EMQServerAddress = ""

// var initClient mqtt.Client

// 创建全局mqtt publish消息处理 handler
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("Push Message:")
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

//// 创建全局mqtt sub消息处理 handler
//var messageSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
//	fmt.Println("收到订阅消息:")
//	fmt.Printf("Sub Client Topic : %s \n", msg.Topic())
//	fmt.Printf("Sub Client msg : %s \n", msg.Payload())
//}

// 连接的回掉函数
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("mqtt新的连接!" + " Connected")
}

// 丢失连接的回掉函数
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect loss: %v\n", err)
}

func init() {
	// 配置错误提示
	//mqtt.DEBUG = log.New(os.Stdout, "		[mqttDEBUG]", 0)
	mqtt.ERROR = log.New(os.Stdout, " 	[mqttERROR]", 0)
}

type Option func(o *Options)

type Options struct {
	Key string
}

func WithKey(key string) Option {
	return func(o *Options) {
		o.Key = key
	}
}

type Server struct {
	Option     Options
	initClient mqtt.Client
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
	if s.initClient == nil {
		if config2.AppConf.Environment != "pro" {
			ClientId += "_" + config2.AppConf.Environment
		}
		opts := mqtt.NewClientOptions().
			AddBroker("mqtt://" + EMQServerAddress + ":1883").
			SetClientID(ClientId).
			SetUsername(Username).
			SetPassword(Password)
		opts.SetKeepAlive(3600 * time.Second)
		// Message callback handler，在没有任何订阅时，发布端调用此函数
		opts.SetDefaultPublishHandler(messagePubHandler)
		//opts.SetPingTimeout(1 * time.Second)
		opts.OnConnect = connectHandler
		opts.OnConnectionLost = connectLostHandler
		s.initClient = mqtt.NewClient(opts)
		if token := s.initClient.Connect(); token.Wait() && token.Error() != nil {
			fmt.Println("Connect")
			fmt.Println(token.Error())
		}
		//if config2.AppConf.Environment == "pro" {
		//	go Subscription("C8C9A35FFFB3", "C8C9A35FFFB3_event", 0, true, brokerLoadHandler)
		//	go Subscription("AC75252AA3FD", "AC75252AA3FD_event", 0, true, brokerLoadHandler)
		//	go Subscription("AC752529B2C5", "AC752529B2C5_event", 0, true, brokerLoadHandler)
		//	go Subscription("58BF25D6C8D5", "58BF25D6C8D5_event", 0, true, brokerLoadHandler)
		//	go Subscription("58BF25D72676", "58BF25D72676_event", 0, true, brokerLoadHandler)
		//}
		//Subscription("home_event", "home_event", 0, true, brokerLoadHandler)

	}
}

func (s *Server) Close() {
	s.initClient.Disconnect(250)
	fmt.Println("mqtt exit!")
}

func (s *Server) Automate(mac string, code string, status int) (ok bool, err error) {
	var automateStr = make(map[string]int)

	automateStr[code], err = func() (automateStr int, error error) {
		op := fmt.Sprintf("1%d0000", status)
		s, _ := strconv.Atoi(op)
		return s, nil
	}()

	mqttJson, err := json.Marshal(automateStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(mqttJson))
	if ok, _ := s.push(fmt.Sprintf("%v_control", mac), 1, false, string(mqttJson)); !ok {
		fmt.Printf("mqtt推送失败,%v_control,json:%v", mac, string(mqttJson))

		return false, errors.New("推送失败，请找管理员")
	}
	return true, nil
}
func (s *Server) AutomateMain(mac string, code []string, status int) (ok bool, err error) {
	var automateStr = make(map[string]int)
	for _, c := range code {
		automateStr[c], err = func() (automateStr int, error error) {
			op := fmt.Sprintf("1%d0000", status)
			s, _ := strconv.Atoi(op)
			return s, nil
		}()
	}
	mqttJson, err := json.Marshal(automateStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(mqttJson))
	if ok, _ := s.push(fmt.Sprintf("%v_control", mac), 1, false, string(mqttJson)); !ok {
		fmt.Printf("mqtt推送失败,%v_control,json:%v", mac, string(mqttJson))

		return false, errors.New("推送失败，请找管理员")
	}
	return true, nil
}

/**
 * @Description: 发布订阅
 * @param clientID
 * @param addr
 * @param topic
 * @param payload
 */
func (s *Server) push(topic string, qos byte, retain bool, payload string) (ok bool, error error) {
	// opts ClientOptions 用于设置 broker，端口，客户端 id ，用户名密码等选项
	if !s.initClient.IsConnected() {
		if token := s.initClient.Connect(); token.Wait() && token.Error() != nil {
			//fmt.Println("Push Connect")

			//fmt.Println(token.Error())
			return false, token.Error()
		}

	}
	//发布消息
	// qos是服务质量: ==1: 一次, >=1: 至少一次, <=1:最多一次
	// retained: 表示mqtt服务器要保留这次推送的信息，如果有新的订阅者出现，就会把这消息推送给它（持久化推送）
	if config2.AppConf.Environment == "pro" {
		fmt.Println("send open light...")

		token := s.initClient.Publish(topic, qos, retain, payload)
		token.Wait()
		if token.Error() != nil {
			return false, token.Error()
		}
	} else {
		fmt.Println("not send open light...")
	}
	fmt.Println("Push Data : "+topic, "Data Size is "+strconv.Itoa(len(payload)))
	//fmt.Println("Disconnect with broker")
	return true, nil
}

// Subscription /**
func (s *Server) subscription(mac string, topic string, qos byte, isSub bool, handleFun func(string, []byte)) {
	if token := s.initClient.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Subscription Connect error")
		fmt.Println(token.Error())
	}

	if isSub {
		// 订阅消息
		if token := s.initClient.Subscribe(topic, qos, func(initClient mqtt.Client, msg mqtt.Message) {
			fmt.Printf("Receive Subscribe Message :")
			fmt.Printf("Sub Client Topic : %s, Data size is  %d \n", msg.Topic(), len(msg.Payload()))
			if len(msg.Payload()) > 0 {
				handleFun(mac, msg.Payload())

			}
		}); token.Wait() && token.Error() != nil {
			fmt.Println("Subscription")
			fmt.Println(token.Error())
			//os.Exit(1)
		}
	} else {
		// 取消订阅
		if token := s.initClient.Unsubscribe(topic); token.Wait() && token.Error() != nil {
			fmt.Println("cancel Subscription")
			fmt.Println(token.Error())
			//os.Exit(1)
		}
	}
}

func (s *Server) brokerLoadHandler(mac string, bytes []byte) {
	//fmt.Printf("BrokerLoadHandler         ")
	//fmt.Printf("%s\n", bytes)
	var callback = make(map[string]int)
	if err := json.Unmarshal(bytes, &callback); err != nil {
		fmt.Printf("json Unmarshal error\n")
	}
	fmt.Printf("mac:%s  json:%v\n", mac, callback)

}
