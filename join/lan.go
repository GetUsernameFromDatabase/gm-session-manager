// Package join covers how devices reach the host: LAN address detection and
// the self-signed TLS certificate that names those addresses.
package join

import (
	"fmt"
	"net"
)

// LANAddresses returns the private-range addresses (RFC 1918 and IPv6 ULA) of
// every interface that is up. Loopback and link-local addresses are excluded.
func LANAddresses() ([]net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("list network interfaces: %w", err)
	}
	var ips []net.IP
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, fmt.Errorf("addresses of %s: %w", iface.Name, err)
		}
		for _, a := range addrs {
			if ipNet, ok := a.(*net.IPNet); ok && ipNet.IP.IsPrivate() {
				ips = append(ips, ipNet.IP)
			}
		}
	}
	return ips, nil
}
