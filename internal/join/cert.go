package join

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"slices"
	"time"
)

// Certificate files inside the directory passed to LoadOrCreateCert.
const (
	CertFile = "cert.pem"
	KeyFile  = "key.pem"
)

const (
	certValidity = 397 * 24 * time.Hour
	// renewBefore regenerates a certificate this close to expiry, so it never
	// lapses in the middle of a game night.
	renewBefore = 30 * 24 * time.Hour
	certOrg     = "gm-session-manager"
)

// localNames are always in the certificate alongside the LAN addresses.
var (
	localDNSNames = []string{"localhost"}
	localIPs      = []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback}
)

// LoadOrCreateCert returns the certificate stored in dir if it is still valid
// at now and covers every address in ips. Otherwise it generates a new
// self-signed certificate, writes it to dir and returns it. The bool reports
// whether a new certificate was created.
func LoadOrCreateCert(dir string, ips []net.IP, now time.Time) (tls.Certificate, bool, error) {
	certPath, keyPath := filepath.Join(dir, CertFile), filepath.Join(dir, KeyFile)

	// Missing, corrupt, expiring or outdated files are all simply replaced:
	// the certificate holds no data worth keeping.
	if existing, err := tls.LoadX509KeyPair(certPath, keyPath); err == nil && usable(existing.Leaf, ips, now) {
		return existing, false, nil
	}

	certPEM, keyPEM, err := generate(ips, now)
	if err != nil {
		return tls.Certificate{}, false, err
	}
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return tls.Certificate{}, false, fmt.Errorf("create certificate dir: %w", err)
	}
	if err := os.WriteFile(keyPath, keyPEM, 0o600); err != nil {
		return tls.Certificate{}, false, fmt.Errorf("write %s: %w", keyPath, err)
	}
	if err := os.WriteFile(certPath, certPEM, 0o600); err != nil {
		return tls.Certificate{}, false, fmt.Errorf("write %s: %w", certPath, err)
	}
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return tls.Certificate{}, false, fmt.Errorf("load generated certificate: %w", err)
	}
	return cert, true, nil
}

// Fingerprint is the SHA-256 of the certificate, in the colon-free hex form
// browsers show in their certificate viewer.
func Fingerprint(cert tls.Certificate) string {
	sum := sha256.Sum256(cert.Leaf.Raw)
	return hex.EncodeToString(sum[:])
}

func usable(leaf *x509.Certificate, ips []net.IP, now time.Time) bool {
	if leaf == nil || now.Before(leaf.NotBefore) || now.Add(renewBefore).After(leaf.NotAfter) {
		return false
	}
	for _, ip := range slices.Concat(localIPs, ips) {
		if leaf.VerifyHostname(ip.String()) != nil {
			return false
		}
	}
	for _, name := range localDNSNames {
		if leaf.VerifyHostname(name) != nil {
			return false
		}
	}
	return true
}

func generate(ips []net.IP, now time.Time) (certPEM, keyPEM []byte, err error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("generate key: %w", err)
	}
	serial, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, nil, fmt.Errorf("generate serial: %w", err)
	}

	dnsNames := slices.Clone(localDNSNames)
	if host, err := os.Hostname(); err == nil && host != "" {
		dnsNames = append(dnsNames, host)
	}
	tmpl := &x509.Certificate{
		SerialNumber:          serial,
		Subject:               pkix.Name{Organization: []string{certOrg}, CommonName: dnsNames[len(dnsNames)-1]},
		NotBefore:             now.Add(-time.Hour), // tolerate small clock skew on phones
		NotAfter:              now.Add(certValidity),
		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              dnsNames,
		IPAddresses:           slices.Concat(localIPs, ips),
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		return nil, nil, fmt.Errorf("create certificate: %w", err)
	}
	keyDER, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return nil, nil, fmt.Errorf("marshal key: %w", err)
	}
	return pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}),
		pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyDER}), nil
}
