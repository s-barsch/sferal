package sitemaps

import (
	"net/http"
	"st/pkg/server"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/sitemaps.xml":
		Index(s, w, r)
	case "/sitemaps/core.xml":
		Core(s, w, r)
	case "/sitemaps/holds.xml":
		Holds(s, w, r)
	case "/sitemaps/graph-els.xml":
		GraphEls(s, w, r)
	default:
		http.NotFound(w, r)
	}
}

/*
	r.HandleFunc("/sitemaps.xml", sitemapIndex)
	r.HandleFunc("/sitemaps/core.xml", sitemapCore)
	r.HandleFunc("/sitemaps/holds.xml", sitemapHolds)
	r.HandleFunc("/sitemaps/graph-els.xml", sitemapGraphEls)
	r.HandleFunc("/sitemaps/index-els.xml", sitemapIndexEls)
*/
