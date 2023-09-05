package setup

import "fmt"

func (a *App) servAddr() string {
	var proto string
	if a.tlsEnabled {
		proto = "https"
	} else {
		proto = "http"
	}

	var domain string
	if a.servDomain != "" {
		domain = a.servDomain
	} else {
		domain = a.servIP
	}

	var port string
	if (a.tlsEnabled && a.httpPort != 433) || (!a.tlsEnabled && a.httpPort != 80) {
		port = fmt.Sprintf(":%d", a.httpPort)
	}
	return fmt.Sprintf("%s://%s%s", proto, domain, port)
}

func (a *App) apiURL() string {
	return a.servAddr() + a.apiPrefix
}
