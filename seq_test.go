package cclib

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestFindImageSequences(t *testing.T) {
	// Create a virtual filesystem with test files
	fs := fstest.MapFS{
		"shot_001.jpg":     &fstest.MapFile{Data: []byte{}},
		"shot_002.jpg":     &fstest.MapFile{Data: []byte{}},
		"shot_003.jpg":     &fstest.MapFile{Data: []byte{}},
		"other-0001.png":   &fstest.MapFile{Data: []byte{}},
		"other-0002.png":   &fstest.MapFile{Data: []byte{}},
		"test.1.exr":       &fstest.MapFile{Data: []byte{}},
		"test.2.exr":       &fstest.MapFile{Data: []byte{}},
		"test.3.exr":       &fstest.MapFile{Data: []byte{}},
		"single.jpg":       &fstest.MapFile{Data: []byte{}},
		"notasequence.txt": &fstest.MapFile{Data: []byte{}},
	}

	sequences, err := FindImageSequences(fs)
	if err != nil {
		t.Fatalf("FindImageSequences failed: %v", err)
	}

	// Expected sequences
	expected := []SeqInfo{
		{
			BaseName:   "shot",
			Delimiter:  "_",
			Padding:    3,
			StartFrame: 1,
			EndFrame:   3,
			Extension:  "jpg",
		},
		{
			BaseName:   "other",
			Delimiter:  "-",
			Padding:    4,
			StartFrame: 1,
			EndFrame:   2,
			Extension:  "png",
		},
		{
			BaseName:   "test",
			Delimiter:  ".",
			Padding:    1,
			StartFrame: 1,
			EndFrame:   3,
			Extension:  "exr",
		},
	}

	// Sort both slices to ensure consistent comparison
	if len(sequences) != len(expected) {
		t.Errorf("Expected %d sequences, got %d", len(expected), len(sequences))
		return
	}

	// Compare each sequence
	for i, exp := range expected {
		found := false
		for _, seq := range sequences {
			if reflect.DeepEqual(exp, seq) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Sequence %d not found: %+v", i, exp)
		}
	}
}
