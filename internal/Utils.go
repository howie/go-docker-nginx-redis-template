package internal

import (
	"context"
	"google.golang.org/appengine"
	"net/http"
)

func getContext(req *http.Request) context.Context {

	//For Go1.9
	ctx := appengine.NewContext(req)
	//For Go1.11
	//ctx:=req.Context()
	return ctx
}
