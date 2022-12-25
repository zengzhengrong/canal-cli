package canal

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/withlin/canal-go/client"
	"github.com/zengzhengrong/canal-cli/config"
)

var canalClient *client.SimpleCanalConnector
var watchTable []string

func init() {
	host := config.Conf.Canal.Host
	port := cast.ToInt(config.Conf.Canal.Port)
	destination := config.Conf.Canal.Destination
	reg := config.Conf.Canal.ListenReg
	// New a connector
	connector := client.NewSimpleCanalConnector(host, port, "", "", destination, 60000, 60*60*1000)
	if err := connector.Connect(); err != nil {
		logrus.Fatal(err)
	}
	canalClient = connector
	// subscribe
	if reg == "" {
		// listen all
		reg = ".\\..*"
	}
	if err := canalClient.Subscribe(reg); err != nil {
		logrus.Fatal(err)
	}
	watchTable = strings.Split(config.Conf.Canal.ListenReg, ",")
	// std info
	logrus.Info("Connect Canal Success:", fmt.Sprintf("Canal Server %s:%d (%s) (%s)", host, port, destination, reg))
}
