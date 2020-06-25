package core

import (
	"net"
	"net/http"
	"time"
)

const (
	defaultDialTimeout = 30 * time.Second
	defaultTlsTimeout  = 10 * time.Second
	defaultIdleTimeout = 10 * time.Minute
	defaultIdleConns   = 10
)

func defaultHttpClient() *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout: defaultDialTimeout,
		}).DialContext,
		TLSHandshakeTimeout: defaultTlsTimeout,
		MaxIdleConnsPerHost: defaultIdleConns,
		IdleConnTimeout:     defaultIdleTimeout,
	}
	return &http.Client{
		Transport: transport,
	}
}
