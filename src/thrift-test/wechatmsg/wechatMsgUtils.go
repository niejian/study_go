package wechatmsg

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"net"
	"os"
	"time"
)

const (
	IP = "localhost"
	PORT = "9991"
	RPCSERVICENAME = "wechatAlarmMsgService"
)

// 获取客户端信息
func getWechatMsgClient() (*WechatAlarmMsgServiceClient, *thrift.TSocket){
	thrift.NewTBufferedTransportFactory(8192)
	//protocolFactory := thrift.NewTCompactProtocolFactory()
	var transport thrift.TTransport


	//transport, err := thrift.NewTSocketTimeout(net.JoinHostPort("127.0.0.1", "8800"), 30 * time.Second)
	socket, err := thrift.NewTSocket(net.JoinHostPort(IP, PORT))
	socket.SetTimeout(30 * time.Second)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	//thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	// 传输方式
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	transport, err = transportFactory.GetTransport(socket)

	if err != nil {
		fmt.Fprintln(os.Stderr, "get transport error: ", err)
		os.Exit(1)
	}
	// 传输协议
	protocol := thrift.NewTBinaryProtocolTransport(transport)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "open transport error: ", err)
		os.Exit(1)
	}
	//defer socket.Close()
	// 客户端 (serviceName要与服务提供者注册上的一致)
	iprot := thrift.NewTMultiplexedProtocol(protocol, RPCSERVICENAME)
	oprot := thrift.NewTMultiplexedProtocol(protocol, RPCSERVICENAME)
	tStandardClient := thrift.NewTStandardClient(iprot, oprot)

	client := NewWechatAlarmMsgServiceClient(tStandardClient)

	return client, socket
}

func SendMsg(touser, content string)  {
	client, socket := getWechatMsgClient()
	defer socket.Close()
	ctx := context.Background()
	alarmMsgRequest := &AlarmMsgRequest{
		Touser:  touser,
		Content: content,
	}

	response, _ := client.SendAlarmMsg(ctx, alarmMsgRequest)
	fmt.Printf("发送消息，参数：%v，返回结果：%v \n", alarmMsgRequest, response)
}
