package glresp

import (
	"net/http"
	"strconv"
	"strings"
)

// Pagination represents pagination info from header
type Pagination struct {
	Total      int    // the total number of items
	TotalPages int    // the total number of pages
	PerPage    int    // the number of items per page
	Page       int    // the index of current page, started from 1
	NextPage   int    // the index of next page
	PrevPage   int    // the index of previous page
	LinkHeader string // raw Link header
}

// ParsePagination extracts pagination info from http headers
func ParsePagination(h http.Header) *Pagination {
	ret := &Pagination{}

	fill := func(dest *int, h http.Header, k string) {
		if i, err := strconv.Atoi(h.Get(k)); err == nil {
			*dest = i
		}
	}

	fill(&(ret.Total), h, "X-Total")
	fill(&(ret.TotalPages), h, "X-Total-Pages")
	fill(&(ret.PerPage), h, "X-Per-Page")
	fill(&(ret.Page), h, "X-Page")
	fill(&(ret.NextPage), h, "X-Next-Page")
	fill(&(ret.PrevPage), h, "X-Prev-Page")
	if l := h.Get("Link"); l != "" {
		ret.LinkHeader = l
	}
	return ret
}

// Links parses LinkHeader and returns a map[rel]url
//
// This method does not test against comma and colon, since they are not expected
// to exist in reply of gitlab. The malformed input is also ignored.
// You should be aware of this.
func (p *Pagination) Links() map[string]string {
	ret := make(map[string]string)

	// should not to met , and ; in url and rel, so just skip related tests
	for _, l := range strings.Split(p.LinkHeader, ",") {
		data := strings.Split(l, ";")
		if len(data) != 2 {
			// malform data, skip it
			continue
		}

		uri := strings.TrimSpace(data[0])
		if !strings.HasPrefix(uri, "<") || !strings.HasSuffix(uri, ">") {
			// malform data, skip it
			continue
		}
		// trim beginning < and tailing >
		uri = uri[1 : len(uri)-1]

		rel := strings.TrimSpace(data[1])
		if !strings.HasPrefix(rel, `rel="`) || !strings.HasSuffix(rel, `"`) || len(rel) < 7 {
			// malform data, skip it
			continue
		}
		rel = rel[5 : len(rel)-1]

		ret[rel] = uri
	}

	return ret
}
