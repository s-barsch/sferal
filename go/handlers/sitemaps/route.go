package sitemaps

import (
	"net/http"
	"sacer/go/server"
	"sacer/go/server/users"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, a *users.Auth) {
	switch r.URL.Path {
	case "/sitemaps.xml":
		Index(s, w, r, a)
	case "/sitemaps/core.xml":
		Core(s, w, r, a)
	case "/sitemaps/trees.xml":
		Trees(s, w, r, a)
	case "/sitemaps/kines.xml":
		Kines(s, w, r, a)
	case "/sitemaps/graph-entries.xml":
		GraphEntries(s, w, r, a)
	default:
		http.NotFound(w, r)
	}
}

/*
	r.HandleFunc("/sitemaps.xml", sitemapIndex)
	r.HandleFunc("/sitemaps/core.xml", sitemapCore)
	r.HandleFunc("/sitemaps/holds.xml", sitemapHolds)
	r.HandleFunc("/sitemaps/graph-els.xml", sitemapGraphEls)
	r.HandleFunc("/sitemaps/indecs-els.xml", sitemapIndexEls)
*/
