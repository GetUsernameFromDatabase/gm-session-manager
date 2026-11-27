// Package join covers how devices reach the host: LAN address detection, the
// self-signed TLS certificate that names those addresses, and the redirect of
// plain HTTP to HTTPS on the same port.
package join

import (
	"context"
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

// PreferDefaultRoute moves the address the default route leaves from to the
// front: Hyper-V and WSL vEthernet adapters are private and up, and can come
// first. On any error the list is unchanged.
func PreferDefaultRoute(ips []net.IP) []net.IP {
	conn, err := (&net.Dialer{}).DialContext(context.Background(), "udp4", "192.0.2.1:9") // UDP connect sends nothing
	if err != nil {
		return ips
	}
	defer func() { _ = conn.Close() }()
	addr, ok := conn.LocalAddr().(*net.UDPAddr)
	if !ok {
		return ips
	}
	return preferFirst(ips, addr.IP)
}

// preferFirst returns ips with the address equal to pref first, order otherwise kept.
func preferFirst(ips []net.IP, pref net.IP) []net.IP {
	out := make([]net.IP, 0, len(ips))
	for _, ip := range ips {
		if ip.Equal(pref) {
			out = append(out, ip)
		}
	}
	for _, ip := range ips {
		if !ip.Equal(pref) {
			out = append(out, ip)
		}
	}
	return out
}
