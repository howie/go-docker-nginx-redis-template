package internal

import (
	"fmt"
	"google.golang.org/appengine/log"
	"net/http"
)

func Index(resp http.ResponseWriter, req *http.Request) {
	msg := "Welcome to GAE"
	ctx := getContext(req)
	//For Go1.9
	log.Infof(ctx, "%v", msg)
	_, _ = fmt.Fprint(resp, msg)
}
