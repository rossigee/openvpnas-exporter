# OpenVPN AS metrics exporter for Prometheus

A Prometheus exporter that makes calls to the XML RPC unix domain socket exposed by the [OpenVPN AS](https://openvpn.net/) service, and generates metrics from the responses.

Current metrics include:

* Server status
  * Number of connected clients.
* Subscription status
  * Number of current, fallback and maximum concurrent connections.
  * Subscription last updated timestamp.

That's all I've needed and had time to implement so far.

Based on [openvpn-exporter](https://github.com/kumina/openvpn_exporter).


## Exposed metrics example

```
# HELP openvpnas_server_connected_clients Number Of Connected Clients
# TYPE openvpnas_server_connected_clients gauge
openvpnas_server_connected_clients 1
# HELP openvpnas_subscription_current_client_connections Number of client connections currently being used from the OpenVPN subscription.
# TYPE openvpnas_subscription_current_client_connections gauge
openvpnas_subscription_current_client_connections 1
# HELP openvpnas_subscription_fallback_client_connections Number of fallback connections in use on the OpenVPN subscription.
# TYPE openvpnas_subscription_fallback_client_connections gauge
openvpnas_subscription_fallback_client_connections 2
# HELP openvpnas_subscription_maximum_client_connections Maximum number of client connections allowed by the OpenVPN subscription.
# TYPE openvpnas_subscription_maximum_client_connections gauge
openvpnas_subscription_maximum_client_connections 100
# HELP openvpnas_subscription_status_update_time_seconds UNIX timestamp at which the OpenVPN subscription status was last updated.
# TYPE openvpnas_subscription_status_update_time_seconds gauge
openvpnas_subscription_status_update_time_seconds 1.636795645e+09
# HELP openvpnas_up Whether scraping OpenVPN's metrics was successful.
# TYPE openvpnas_up gauge
openvpnas_up 1
```

## Usage

Usage of `openvpnas_exporter`:

```sh
  -openvpn-as.xmlrpc-path string
    	Path at which the XML-RPC unix domain socket file can be found. (default "/usr/local/openvpn_as/etc/sock/sagent.localroot")
  -web.listen-address string
    	Address to listen on for web interface and telemetry. (default ":9176")
  -web.telemetry-path string
    	Path under which to expose metrics. (default "/metrics")
```

It appears to only run as 'root' user. The XML-RPC request fails otherwise.

## Get a standalone executable binary

You can download the pre-compiled binaries from the [releases page](https://github.com/rossigee/openvpnas-exporter/releases).
