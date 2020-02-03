package util

import (
	"testing"
)

const mapfile1 = "../../storage/sample_map"

func TestBuildMap(t *testing.T) {

	cities2, world2, err2 := MakeMap(mapfile1)
	if err2 != nil {
		t.Fatalf("Unable to open test map %s", mapfile1)
	}

	if len(cities2) != 5 {
		t.Fatalf("Number of loaded cities differ, got: %d, want: %d", len(cities2), 5)
	}

	if len(world2) != 5 {
		t.Fatalf("Number of mapped cities with names differ, got: %d, want: %d", len(world2), 5)
	}
}
