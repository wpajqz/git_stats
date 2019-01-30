package http

import (
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"git.links123.net/links123.com/stats/config"

	"github.com/braintree/manners"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// RunCommand cobra subcommand http
func RunCommand() *cobra.Command {
	var host, port string
	cmd := &cobra.Command{
		Use:   "http",
		Short: "Run the http service",
		Run: func(cmd *cobra.Command, args []string) {
			go Start(host, port)
			// 阻塞退出 捕获信号
			signalChan := make(chan os.Signal)
			signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
			logrus.Infof("caught signal %+v, begin garbage collection", <-signalChan)
			signal.Reset()
			Stop()
		},
	}
	cmd.PersistentFlags().StringVarP(&host, "host", "o", "127.0.0.1", "server hostname")
	cmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "server port")

	return cmd
}

// Start start the http server
func Start(host, port string) {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(config.C.StatsDir)))

	logrus.Infof("start http server on %s:%s", host, port)
	err := manners.ListenAndServe(strings.Join([]string{host, port}, ":"), mux)
	if err != nil {
		panic(err)
	}
}

// Stop stop the http service graceful
func Stop() {
	if manners.Close() {
		logrus.Info("http server stopped")
	}
}
