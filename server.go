package socks5proxy

import (
	"log"
	"net"
	"sync"
)

func handleClientRequest(conn *net.TCPConn, auth socks5Auth) {
	if conn == nil {
		return
	}
	defer conn.Close()

	// 初始化一个字符串buff
	buff := make([]byte, 255)

	// 认证协商
	var proto ProtocolVersion
	n, err := auth.DecodeRead(conn, buff) //解密
	resp, err := proto.HandleHandshake(buff[0:n])
	auth.EncodeWrite(conn, resp) //加密
	if err != nil {
		log.Print(conn.RemoteAddr(), err)
		return
	}

	// 获取客户端代理的请求
	var request Socks5Resolution
	n, err = auth.DecodeRead(conn, buff)
	resp, err = request.LSTRequest(buff[0:n])
	auth.EncodeWrite(conn, resp)
	if err != nil {
		log.Print(conn.RemoteAddr(), err)
		return
	}

	log.Println(conn.RemoteAddr(), request.DSTDOMAIN, request.DSTADDR, request.DSTPORT)

	// 连接真正的远程服务
	dstServer, err := net.DialTCP("tcp", nil, request.RAWADDR)
	if err != nil {
		log.Print(conn.RemoteAddr(), err)
		return
	}
	defer dstServer.Close()

	wg := new(sync.WaitGroup)
	wg.Add(2)

	// 本地的内容copy到远程端
	go func() {
		defer wg.Done()
		SecureCopy(conn, dstServer, auth.Decrypt)
	}()

	// 远程得到的内容copy到源地址
	go func() {
		defer wg.Done()
		SecureCopy(dstServer, conn, auth.Encrypt)
	}()
	wg.Wait()
}

func Server(listenPort string, encrypt string, passwd string) {
	// 所有客户服务端的流都加密,
	auth, err := CreateAuth(encrypt, passwd)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("服务器||验证密码是:%s", passwd)

	// 监听客户端
	listenAddr, err := net.ResolveTCPAddr("tcp", listenPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("服务器||监听的端口: %s ", listenPort)

	listener, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go handleClientRequest(conn, auth)
	}
}
