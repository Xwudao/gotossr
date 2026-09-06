package main

import (
	"log"
	"net/http"

	gossr "github.com/Xwudao/gotossr"
)

type PageProps struct {
	Message string `json:"message"`
}

func main() {
	engine, err := gossr.New(gossr.Config{
		AppEnv:            "production",
		AssetRoute:        "/assets",
		FrontendDir:       "./frontend",
		ClientAppPath:     "App.tsx",
		SPAHydrationMode:  "tanstack",
		JSRuntimePoolSize: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write(engine.RenderRoute(gossr.RenderConfig{
			File:        "App.tsx",
			Title:       "gotossr TanStack Router",
			RequestPath: r.URL.Path,
			Props:       PageProps{Message: "Rendered by Go on the server"},
		}))
	})
	log.Println("Open http://localhost:8080 or /about")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
