package main

import (
	"code.bigogo.com/gopkg/ginlog"
	"code.bigogo.com/gopkg/logs"
	"flag"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func main() {
	defer logs.Flush()

	var (
		flagAddress = flag.String("a", ":9000", "http address")
	)
	flag.Parse()

	ln, err := net.Listen("tcp", *flagAddress)

	if err != nil {
		logs.Fatalf("create listener failed: %s", *flagAddress)
		logs.Flush()
		return
	}

	logs.Infof("listening on %s", ln.Addr())
	r := gin.Default()
	CreateRouter(r)

	logs.Errorf("server running with error: %v", http.Serve(ln, r))
	ginlog.CloseLog()
}
