package route

import "testing"

func TestPath(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		ids     []any
		want    string
	}{
		{"none", "/live", nil, "/live"},
		{"root", "/", nil, "/"},
		{"one", "/notes/:note", []any{7}, "/notes/7"},
		{"two", "/images/:image/:action", []any{int64(3), "up"}, "/images/3/up"},
		{"param in the middle", "/blocks/:block/steps", []any{12}, "/blocks/12/steps"},
		{"escaped", "/compendium/find/:kind/:slug", []any{"monster", "a b/c?"}, "/compendium/find/monster/a%20b%2Fc%3F"},
		{"catch-all", "/assets/*path", []any{"js app.js"}, "/assets/js%20app.js"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Path(tt.pattern, tt.ids...); got != tt.want {
				t.Errorf("Path(%q, %v) = %q, want %q", tt.pattern, tt.ids, got, tt.want)
			}
		})
	}
}

func TestPathPanicsOnWrongIDCount(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		ids     []any
	}{
		{"too few", "/images/:image/:action", []any{1}},
		{"none for a param", "/notes/:note", nil},
		{"too many", "/notes/:note", []any{1, 2}},
		{"too many for none", "/live", []any{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Errorf("Path(%q, %v) did not panic", tt.pattern, tt.ids)
				}
			}()
			Path(tt.pattern, tt.ids...)
		})
	}
}
