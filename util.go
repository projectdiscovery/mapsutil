package mapsutil

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/miekg/dns"
)

// MergeMaps into a new one
func MergeMaps(m1, m2 map[string]interface{}) (m map[string]interface{}) {
	m = make(map[string]interface{})

	for k, v := range m1 {
		m[k] = v
	}

	for k, v := range m2 {
		m[k] = v
	}

	return
}

// MergeMapsWithStrings into a new string one
func MergeMapsWithStrings(m1, m2 map[string]string) (m map[string]string) {
	m = make(map[string]string)
	for k, v := range m1 {
		m[k] = v
	}

	for k, v := range m2 {
		m[k] = v
	}

	return
}

const defaultFormat = "%s"

// HTTPToMap Converts HTTP to Matcher Map
func HTTPToMap(resp *http.Response, body, headers string, duration time.Duration, format string) (m map[string]interface{}) {
	m = make(map[string]interface{})

	if format == "" {
		format = defaultFormat
	}

	m[fmt.Sprintf(format, "content_length")] = resp.ContentLength
	m[fmt.Sprintf(format, "status_code")] = resp.StatusCode

	for k, v := range resp.Header {
		k = strings.ToLower(strings.TrimSpace(strings.ReplaceAll(k, "-", "_")))
		m[fmt.Sprintf(format, k)] = strings.Join(v, " ")
	}

	m[fmt.Sprintf(format, "all_headers")] = headers
	m[fmt.Sprintf(format, "body")] = body

	if r, err := httputil.DumpResponse(resp, true); err == nil {
		m[fmt.Sprintf(format, "raw")] = string(r)
	}

	// Converts duration to seconds (floating point) for DSL syntax
	m[fmt.Sprintf(format, "duration")] = duration.Seconds()

	return m
}

// DNSToMap Converts DNS to Matcher Map
func DNSToMap(msg *dns.Msg, format string) (m map[string]interface{}) {
	m = make(map[string]interface{})

	if format == "" {
		format = defaultFormat
	}

	m[fmt.Sprintf(format, "rcode")] = msg.Rcode

	var qs string

	for _, question := range msg.Question {
		qs += fmt.Sprintln(question.String())
	}

	m[fmt.Sprintf(format, "question")] = qs

	var exs string
	for _, extra := range msg.Extra {
		exs += fmt.Sprintln(extra.String())
	}

	m[fmt.Sprintf(format, "extra")] = exs

	var ans string
	for _, answer := range msg.Answer {
		ans += fmt.Sprintln(answer.String())
	}

	m[fmt.Sprintf(format, "answer")] = ans

	var nss string
	for _, ns := range msg.Ns {
		nss += fmt.Sprintln(ns.String())
	}

	m[fmt.Sprintf(format, "ns")] = nss
	m[fmt.Sprintf(format, "raw")] = msg.String()

	return m
}

func HTTPRequesToMap(req *http.Request) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	var headers string
	for k, v := range req.Header {
		k = strings.ToLower(strings.TrimSpace(strings.ReplaceAll(k, "-", "_")))
		vv := strings.Join(v, " ")
		m[k] = strings.Join(v, " ")
		headers += fmt.Sprintf("%s: %s", k, vv)
	}

	m["all_headers"] = headers

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	m["body"] = body

	reqdump, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil, err
	}

	m["raw"] = string(reqdump)

	return m, nil
}

func HTTPResponseToMap(resp *http.Response) (map[string]interface{}, error) {
	m := make(map[string]interface{})

	m["content_length"] = resp.ContentLength
	m["status_code"] = resp.StatusCode
	var headers string
	for k, v := range resp.Header {
		k = strings.ToLower(strings.TrimSpace(strings.ReplaceAll(k, "-", "_")))
		vv := strings.Join(v, " ")
		m[k] = vv
		headers += fmt.Sprintf("%s: %s", k, vv)
	}
	m["all_headers"] = headers

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	m["all_headers"] = headers
	m["body"] = body

	if r, err := httputil.DumpResponse(resp, true); err == nil {
		m["raw"] = string(r)
	}

	return m, nil
}
