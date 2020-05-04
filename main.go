package main

import (
	"fmt"
	"net"
	"os"
	"errors"

	"github.com/tsthght/s3folder/s3common/s3castleclient"
	dsync "github.com/pingcap/tidb-binlog/drainer/sync"
)

func newCastleClient( ) (castleManager *s3castleclient.CastleClientManager, err error) {
	var ip, host string
	err, ip = GetLocalIP()
	if err != nil {
		return nil, err
	}
	err, host = GetLocalHost()
	if err != nil {
		return nil, err
	}
	if castleManager, err = s3castleclient.NewCastleClientManager(host, ip,
		"", "", ""); err != nil {
		return nil, err
	}
	return
}

func main() {
	newCastleClient()
}

func GetLocalIP() (error, string) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return err, ""
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return nil, ipnet.IP.String()
			}
		}
	}
	return errors.New("can not get local ip"), ""
}

func GetLocalHost() (error, string) {
	host, err := os.Hostname()
	if err != nil {
		return err, ""
	}
	return nil, host
}
