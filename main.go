package main

import (
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/braintree/manners"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// @title stats api for the project of links123.com
// @version 1.0
// @description stats api for the project of links123.com
// @host 127.0.0.1:8080/v1
func main() {
	var host, port string
	cmd := &cobra.Command{
		Use:   "stats",
		Short: "stats api for the project of links123.com",
		Run: func(cmd *cobra.Command, args []string) {
			go start(host, port)
			// 阻塞退出 捕获信号
			signalChan := make(chan os.Signal)
			signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
			logrus.Infof("caught signal %+v, begin garbage collection", <-signalChan)
			signal.Reset()
			stop()
		},
	}

	cmd.PersistentFlags().StringVarP(&host, "host", "o", "127.0.0.1", "server hostname")
	cmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "server port")

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

// start start the http server
func start(host, port string) {
	//go RunGitStats()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(c.StatsDir)))

	logrus.Infof("start http server on %s:%s", host, port)
	err := manners.ListenAndServe(strings.Join([]string{host, port}, ":"), mux)
	if err != nil {
		panic(err)
	}
}

// Stop stop the http service graceful
func stop() {
	if manners.Close() {
		logrus.Info("http server stopped")
	}
}
