package main

import (
	"kerberos/pkg/core"
	"kerberos/pkg/sdk/http/client"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetPrincipal() (clientId string, clientHash []byte, serverId string) {
	return "001", core.GetHash32("001"), "002"
}

func main() {
	ctx := client.NewSecurityContext(GetPrincipal())
	cli := client.NewKrbClient(ctx, core.NewEAS())

	for i := 0; i < 100; i++ {
		request, _ := http.NewRequest(http.MethodPost, "http://localhost:9999/server/echo", strings.NewReader(""))
		log.Println(cli.HttpPost(request))
		time.Sleep(time.Millisecond * 505)
	}
}
