package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rossigee/openvpnas-exporter/exporters"
)

func main() {
	var (
		listenAddress     = flag.String("web.listen-address", ":9176", "Address to listen on for web interface and telemetry.")
		metricsPath       = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
		xmlrpcPath        = flag.String("openvpn-as.xmlrpc-path", "/usr/local/openvpn_as/etc/sock/sagent.localroot", "Path of XMLRPC unix domain socket to connect to.")
	)
	flag.Parse()

	log.Printf("Starting OpenVPN AS Exporter\n")
	log.Printf("Listen address: %v\n", *listenAddress)
	log.Printf("Metrics path: %v\n", *metricsPath)
	log.Printf("XML-RPC unix domain path: %v\n", *xmlrpcPath)

	exporter, err := exporters.NewOpenVPNExporter(*xmlrpcPath)
	if err != nil {
		panic(err)
	}
	prometheus.MustRegister(exporter)

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
			<head><title>OpenVPN AS Exporter</title></head>
			<body>
			<h1>OpenVPN AS Exporter</h1>
			<p><a href='` + *metricsPath + `'>Metrics</a></p>
			</body>
			</html>`))
	})
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
