package join

import (
	"net"
	"testing"
)

func TestPreferFirst(t *testing.T) {
	a, b, c := net.ParseIP("172.20.0.1"), net.ParseIP("192.168.1.20"), net.ParseIP("10.0.0.5")
	tests := []struct {
		name string
		ips  []net.IP
		pref net.IP
		want []net.IP
	}{
		{"moves to front", []net.IP{a, b, c}, b, []net.IP{b, a, c}},
		{"already first", []net.IP{b, a}, b, []net.IP{b, a}},
		{"absent keeps order", []net.IP{a, b}, c, []net.IP{a, b}},
		{"empty", nil, b, []net.IP{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := preferFirst(tt.ips, tt.pref)
			if len(got) != len(tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
			for i := range got {
				if !got[i].Equal(tt.want[i]) {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			}
		})
	}
}
