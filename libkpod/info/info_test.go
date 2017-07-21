package info

import (
	"testing"

	"golang.org/x/net/context"
)

func TestRegisterInfo(t *testing.T) {
	defer resetInfoGivers()
	firstCount := len(InfoGivers())
	RegisterInfoGiver(func(ctx context.Context) (name string, info map[string]interface{}, err error) {
		return "test", nil, nil
	})
	newCount := len(InfoGivers())

	diff := newCount - firstCount
	if diff != 1 {
		t.Errorf("expected only one new InfoGiver, but got %d", diff)
	}
}
