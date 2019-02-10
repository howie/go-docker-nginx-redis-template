# Running  Go 1.11x program on GAE 2.0 environment

本範例是示範如何在如何在 GAE 2.0 上執行 Go 1.11x 程式(with Go module)，並且繼續使用 GAE 1.0 的 API。

## 緣起

Go1.11 版本以後開始支援 go module ，大大簡化了開發路徑與依賴管理的麻煩，所以一開始學go 就可以直接使用go module 是幸福的。
不過因為 1.11 版本還很新，所以在許多地方可能會遇到麻煩，例如 GAE 的環境。

GAE 2.0 版本以後才能執行 Go 1.11以後的版本，但是有許多GAE的API(如 Search API)都還沒支援GAE 2.0的環境，因此需要一些 workaround.

## 初始化

整個專案的Layout 還是依照 go 1.11 的開發環境，只是有些需要初始化的地方仍必須依照Go1.9版本以前的範例程式撰寫，請看main.go
主要有兩點差異：

### 需要初始化 app engine
```
func main() {

	registerHandlers()
	// For Go1.9 runtime
	appengine.Main()
}
```
### 處理 url 的方法

```
    // For Go1.9 runtime
	// [START request_logging]
	// Delegate all of the HTTP routing and serving to the gorilla/mux router.
	// Log all requests using the standard Apache format.
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
	// [END request_logging]

	// For Go1.11 runtime
	//  port := os.Getenv("PORT")
	//  if port == "" {
	//	port = "8080"
	//	log.Printf("Defaulting to port %s", port)
	//  }
	//  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))

```

### 取得 context 的方法

根據我的實驗，如果要使用GAE1.0 的API 必須使用 app engine api 取得的 context，而不能直接從 http.Request 取得。 

```
func getContext(req *http.Request) context.Context {

	//For Go1.9
	ctx := appengine.NewContext(req)
	//For Go1.11
	//ctx:=req.Context()
	return ctx
}
```

### 使用log 的方法

```
//For Go1.9
log.Infof(ctx, "%v", msg)
	
```

只要這幾點都有遵守，就可以開心的在GAE 2.0 的環境，執行和使用GAE1.0 的API