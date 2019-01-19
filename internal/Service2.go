package internal

import (
	"fmt"
	"google.golang.org/appengine/log"
	"html"
	"net/http"
)

func Echo(resp http.ResponseWriter, req *http.Request) {

	msg := html.UnescapeString(req.URL.Query().Get("msg"))

	ctx := getContext(req)
	//For Go1.9
	log.Infof(ctx, "%v", msg)
	_, _ = fmt.Fprint(resp, msg)

}
