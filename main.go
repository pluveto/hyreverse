package main

import (
	"context"
	"fmt"
	"github.com/akamensky/argparse"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const VersionNumber = "0.1"
const HTTPS = "https"
const HTTP = "http"

type Option struct {
	LocalProtocol string
	LocalAddress  string
	CertFile      *string
	KeyFile       *string
	RemoteUrl     string
}

func getOptions() Option {
	programName := filepath.Base(os.Args[0])
	parser := argparse.NewParser(programName, "simple and handy reverse proxy for http/https services")
	local := parser.String("l", "local", &argparse.Options{Required: true, Help: "Local address to host on. Example: http://localhost:8080"})
	remote := parser.String("r", "remote", &argparse.Options{Required: true, Help: "Remote address. Example: https://www.google.com"})
	cert := parser.String("c", "cert", &argparse.Options{Required: false, Help: "Cert file path"})
	key := parser.String("k", "key", &argparse.Options{Required: false, Help: "Key file path"})
	err := parser.Parse(os.Args)

	if err != nil {
		print(parser.Usage(err))
		os.Exit(255)
	}
	localProtocol, localAddress := parseLocal(*local)

	return Option{
		LocalProtocol: localProtocol,
		LocalAddress:  localAddress,
		CertFile:      cert,
		KeyFile:       key,
		RemoteUrl:     *remote,
	}
}

func parseLocal(s string) (localProtocol, localAddress string) {
	const httpsPrefix = "https://"
	const httpPrefix = "http://"

	if strings.HasPrefix(s, httpsPrefix) {
		localProtocol = HTTPS
		localAddress = s[len(httpsPrefix):]
	} else if strings.HasPrefix(s, httpPrefix) {
		localProtocol = HTTP
		localAddress = s[len(httpPrefix):]
	} else {
		localProtocol = HTTP
		localAddress = s
		log.Println("warning: using http protocol for local service")
	}

	localAddress = strings.TrimRight(localAddress, "/")

	return
}

func main() {
	option := getOptions()
	fmt.Printf(
		"|  |      __   ___       ___  __   __   ___ \n"+
			"|__| \\ / |__) |__  \\  / |__  |__) /__` |__  \n"+
			"|  |  |  |  \\ |___  \\/  |___ |  \\ .__/ |___ \n"+
			"Version: %v. Listening on %v://%v -> %v\n", VersionNumber,
		option.LocalProtocol, option.LocalAddress, option.RemoteUrl,
	)
	ListenAndServe(option)

}

func noContextCancellationErrors(rw http.ResponseWriter, req *http.Request, err error) {
	if err != context.Canceled {
		log.Printf("http: proxy error: %v", err)
	}
	rw.WriteHeader(http.StatusBadGateway)
}

func ListenAndServe(option Option) {
	targetUrl, err := url.Parse(option.RemoteUrl)
	if err != nil {
		log.Panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
	proxy.ErrorHandler = noContextCancellationErrors
	if option.LocalProtocol == HTTPS && nil != option.CertFile && nil != option.KeyFile {
		err = http.ListenAndServeTLS(option.LocalAddress, *option.CertFile, *option.KeyFile, proxy)
	} else {
		err = http.ListenAndServe(option.LocalAddress, proxy)
	}
	if err != nil {
		log.Panic(err)
	}
}
