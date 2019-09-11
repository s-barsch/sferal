package index 

import (
	"log"
	"net/http"
	"st/pkg/el"
	"st/pkg/head"
	"st/pkg/paths"
	"st/pkg/server"
)

type indexMain struct {
	Head    *head.Head
	Hold    *el.Hold
	Recents el.Els
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	head := &head.Head{
		Title:   "Index",
		Section: "index",
		Path:    path,
		Host:    r.Host,
		El:      s.Trees["index"],
	}
	err = head.Make()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "index-hold", &indexMain{
		Head:    head,
		Hold:    s.Trees["index"],
		Recents: s.Recents["index"],
	})
	if err != nil {
		log.Println(err)
	}
}
