package db

import (
	"context"
	"errors"
	"testing"

	"gm-session-manager/db/sqlc"
)

func TestStoreInTx(t *testing.T) {
	errAbort := errors.New("abort")
	tests := []struct {
		name      string
		fnErr     error
		wantCount int
	}{
		{"commit on nil", nil, 1},
		{"rollback on error", errAbort, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := t.Context()
			store := NewStore(openTemp(t))

			err := store.InTx(ctx, func(q *sqlc.Queries) error {
				if _, err := q.CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "Test", System: "dnd5e", Edition: "2024"}); err != nil {
					return err
				}
				return tt.fnErr
			})
			if !errors.Is(err, tt.fnErr) {
				t.Fatalf("InTx error = %v, want %v", err, tt.fnErr)
			}
			if got := countCampaigns(ctx, t, store); got != tt.wantCount {
				t.Errorf("campaigns = %d, want %d", got, tt.wantCount)
			}
		})
	}
}

func countCampaigns(ctx context.Context, t *testing.T, store *Store) int {
	t.Helper()
	camps, err := store.ListCampaigns(ctx)
	if err != nil {
		t.Fatal(err)
	}
	return len(camps)
}
