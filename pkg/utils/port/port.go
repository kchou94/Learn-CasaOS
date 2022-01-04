package port

import (
	"fmt"
	"net"
)

// 获取可用端口
func GetAvailablePort(t string) (int,error) {
	address:=fmt.Sprintf("%s:0", "0.0.0.0")
	if t == "udp" {
		udpAddr,err := net.ResolveUDPAddr(t, address)
		if err != nil {
			return 0,err
		}

		udpConn,err:=net.ListenUDP(t, udpAddr)
		if err != nil {
			return 0,err
		}

		defer udpConn.Close()
		return udpConn.LocalAddr().(*net.UDPAddr).Port,nil
	} else {
		tcpAddr,err:=net.ResolveTCPAddr(t, address)
		if err != nil {
			return 0,err
		}

		tcpConn,err:=net.ListenTCP(t, tcpAddr)
		if err != nil {
			return 0,err
		}

		defer tcpConn.Close()
		return tcpConn.Addr().(*net.TCPAddr).Port,nil
	}
}

// 判断端口是否可以（未被占用）
// param t tcp/udp
func IsPortAvailable(port int, t string) bool {
	address := fmt.Sprintf("%s:%d", "0.0.0.0", port)
	if t == "udp" {
		udpAddr, err := net.ResolveUDPAddr(t address)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}

		udpConn, err := net.ListenUDP(t, udpAddr)
		if err != nil {
			fmt.Println(err.Error())
			return false
		} else {
			defer udpConn.Close()
			return true
		}
	} else {
		listener, err := net.Listen(t, address)

		if err != nil {
			return false
		}

		defer listener.Close()
		return true
	}
}
