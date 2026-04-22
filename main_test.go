package main

import "testing"

func TestGetFileType(t *testing.T) {
	tests := []struct {
		filename string
		expected string
	}{
		{"image.jpg", "jpg"},
		{"video.mp4", "mp4"},
		{"document.pdf", "pdf"},
		{"noextension", ""},
		{"archive.tar.gz", "gz"},
	}

	for _, tt := range tests {
		result := getFileType(tt.filename)
		if result != tt.expected {
			t.Errorf("getFileType(%s) = %s; want %s", tt.filename, result, tt.expected)
		}
	}
}
