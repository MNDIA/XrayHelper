package trojan

import (
	"XrayHelper/main/serial"
	"strconv"
	"strings"
)

// getMuxObjectXray get xray MuxObject
func getMuxObjectXray(enabled bool) serial.OrderedMap {
	var mux serial.OrderedMap
	mux.Set("enabled", enabled)
	return mux
}

// getTrojanSettingsObjectXray get xray Trojan SettingsObject
func getTrojanSettingsObjectXray(trojan *Trojan) serial.OrderedMap {
	var serverArray serial.OrderedArray
	var server serial.OrderedMap
	server.Set("address", trojan.Server)
	port, _ := strconv.Atoi(trojan.Port)
	server.Set("port", port)
	server.Set("password", trojan.Password)
	server.Set("level", 0)
	serverArray = append(serverArray, server)

	var settingsObject serial.OrderedMap
	settingsObject.Set("servers", serverArray)
	return settingsObject
}

// getStreamSettingsObjectXray get xray StreamSettingsObject
func getStreamSettingsObjectXray(trojan *Trojan) serial.OrderedMap {
	var streamSettingsObject serial.OrderedMap
	streamSettingsObject.Set("network", trojan.Network)
	switch trojan.Network {
	case "tcp":
		var tcpSettingsObject serial.OrderedMap
		var headerObject serial.OrderedMap
		switch trojan.Type {
		case "http":
			headerObject.Set("type", trojan.Type)
			if len(trojan.Host) > 0 {
				var requestObject serial.OrderedMap
				var headers serial.OrderedMap
				var host serial.OrderedArray
				host = append(host, trojan.Host)
				var connection serial.OrderedArray
				connection = append(connection, "keep-alive")
				var acceptEncoding serial.OrderedArray
				acceptEncoding = append(acceptEncoding, "gzip, deflate")
				var userAgent serial.OrderedArray
				userAgent = append(userAgent, "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
					"Mozilla/5.0 (iPhone; CPU iPhone OS 16_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Mobile/15E148 Safari/604.1")
				headers.Set("Host", host)
				headers.Set("Connection", connection)
				headers.Set("Pragma", "no-cache")
				headers.Set("Accept-Encoding", acceptEncoding)
				headers.Set("User-Agent", userAgent)
				requestObject.Set("headers", headers)
				headerObject.Set("request", requestObject)
			}
		default:
			headerObject.Set("type", "none")
		}
		tcpSettingsObject.Set("header", headerObject)
		streamSettingsObject.Set("tcpSettings", tcpSettingsObject)
	case "kcp":
		var kcpSettingsObject serial.OrderedMap
		if len(trojan.Type) > 0 {
			var headerObject serial.OrderedMap
			headerObject.Set("type", trojan.Type)
			kcpSettingsObject.Set("header", headerObject)
		}
		kcpSettingsObject.Set("congestion", false)
		kcpSettingsObject.Set("downlinkCapacity", 100)
		kcpSettingsObject.Set("mtu", 1350)
		kcpSettingsObject.Set("readBufferSize", 1)
		if len(trojan.Path) > 0 {
			kcpSettingsObject.Set("seed", trojan.Path)
		}
		kcpSettingsObject.Set("tti", 50)
		kcpSettingsObject.Set("uplinkCapacity", 12)
		kcpSettingsObject.Set("writeBufferSize", 1)
		streamSettingsObject.Set("kcpSettings", kcpSettingsObject)
	case "ws":
		var wsSettingsObject serial.OrderedMap
		if len(trojan.Host) > 0 {
			var headersObject serial.OrderedMap
			headersObject.Set("Host", trojan.Host)
			wsSettingsObject.Set("headers", headersObject)
		}
		if len(trojan.Path) > 0 {
			wsSettingsObject.Set("path", trojan.Path)
		}
		streamSettingsObject.Set("wsSettings", wsSettingsObject)
	case "http", "h2":
		var httpSettingsObject serial.OrderedMap
		if len(trojan.Host) > 0 {
			var host serial.OrderedArray
			host = append(host, trojan.Host)
			httpSettingsObject.Set("host", host)
		}
		if len(trojan.Path) > 0 {
			httpSettingsObject.Set("path", trojan.Path)
		}
		streamSettingsObject.Set("httpSettings", httpSettingsObject)
	case "httpupgrade":
		var httpupgradeSettingsObject serial.OrderedMap
		if len(trojan.Host) > 0 {
			var host serial.OrderedArray
			host = append(host, trojan.Host)
			httpupgradeSettingsObject.Set("host", host)
		}
		if len(trojan.Path) > 0 {
			httpupgradeSettingsObject.Set("path", trojan.Path)
		}
		streamSettingsObject.Set("httpupgrade", httpupgradeSettingsObject)
	case "quic":
		var quicSettingsObject serial.OrderedMap
		if len(trojan.Type) > 0 {
			var headerObject serial.OrderedMap
			headerObject.Set("type", trojan.Type)
			quicSettingsObject.Set("header", headerObject)
		}
		if len(trojan.Path) > 0 {
			quicSettingsObject.Set("key", trojan.Path)
		}
		if len(trojan.Host) > 0 {
			quicSettingsObject.Set("security", trojan.Host)
		}
		streamSettingsObject.Set("quicSettings", quicSettingsObject)
	case "grpc":
		var grpcSettingsObject serial.OrderedMap
		if trojan.Type == "multi" {
			grpcSettingsObject.Set("multiMode", true)
		} else {
			grpcSettingsObject.Set("multiMode", false)
		}
		if len(trojan.Host) > 0 {
			grpcSettingsObject.Set("authority", trojan.Host)
		}
		if len(trojan.Path) > 0 {
			grpcSettingsObject.Set("serviceName", trojan.Path)
		}
		streamSettingsObject.Set("grpcSettings", grpcSettingsObject)
	}
	streamSettingsObject.Set("security", trojan.Security)
	switch trojan.Security {
	case "tls":
		var tlsSettingsObject serial.OrderedMap
		var alpn serial.OrderedArray
		alpnSlice := strings.Split(trojan.Alpn, ",")
		for _, v := range alpnSlice {
			if len(v) > 0 {
				alpn = append(alpn, v)
				tlsSettingsObject.Set("alpn", alpn)
			}
		}
		tlsSettingsObject.Set("allowInsecure", false)
		if len(trojan.FingerPrint) > 0 {
			tlsSettingsObject.Set("fingerprint", trojan.FingerPrint)
		}
		if len(trojan.Sni) > 0 {
			tlsSettingsObject.Set("serverName", trojan.Sni)
		}
		streamSettingsObject.Set("tlsSettings", tlsSettingsObject)
	case "reality":
		var realitySettingsObject serial.OrderedMap
		realitySettingsObject.Set("allowInsecure", false)
		if len(trojan.FingerPrint) > 0 {
			realitySettingsObject.Set("fingerprint", trojan.FingerPrint)
		}
		if len(trojan.Sni) > 0 {
			realitySettingsObject.Set("serverName", trojan.Sni)
		}
		realitySettingsObject.Set("publicKey", trojan.PublicKey)
		realitySettingsObject.Set("shortId", trojan.ShortId)
		realitySettingsObject.Set("spiderX", trojan.SpiderX)
		streamSettingsObject.Set("realitySettings", realitySettingsObject)
	}
	var sockoptObject serial.OrderedMap
	sockoptObject.Set("domainStrategy", "UseIP")
	streamSettingsObject.Set("sockopt", sockoptObject)
	return streamSettingsObject
}
