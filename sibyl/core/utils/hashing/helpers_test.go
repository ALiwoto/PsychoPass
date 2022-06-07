package hashing_test

import (
	"testing"

	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/hashing"
)

func TestGenerateAccessHash(t *testing.T) {
	h := hashing.GenerateAccessHash()
	if len(h) != hashing.PollingAccessHashSize {
		t.Error("Expected 32 characters, got", len(h))
	}
}
