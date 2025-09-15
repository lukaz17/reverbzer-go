// Copyright (C) Nguyen Nhat Tung
//
// Reverbzer is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package server

import (
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

// RequestInfo describes basic information about an HTTP request.
type RequestInfo struct {
	Method string `json:"method"`

	Host  string `json:"host"`
	Path  string `json:"path"`
	Query string `json:"queryString"`
	Frag  string `json:"fragment"`

	Headers map[string]string `json:"headers"`

	Length  int64  `json:"contentLength"`
	Body    []byte `json:"body"`
	Content string `json:"content"`

	Origin string    `json:"origin"`
	Time   time.Time `json:"time"`
}

// Return new RequestInfo from http.Request.
func NewRequestInfo(hr *http.Request) *RequestInfo {
	r := &RequestInfo{
		Method: hr.Method,

		Host:  hr.Host,
		Path:  hr.URL.Path,
		Query: hr.URL.RawQuery,
		Frag:  hr.URL.Fragment,

		Origin: hr.RemoteAddr,
		Time:   time.Now(),
	}

	var keys []string
	for key := range hr.Header {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	headers := make(map[string]string)
	for _, key := range keys {
		values := hr.Header[key]
		headers[key] = strings.Join(values, "\n")
	}
	r.Headers = headers

	body, err := io.ReadAll(hr.Body)
	if err == nil && len(body) > 0 {
		r.Length = hr.ContentLength
		r.Body = body
		r.Content = string(body)
	}

	return r
}
