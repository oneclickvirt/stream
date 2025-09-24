package stream

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCleanStream(t *testing.T) {
	// Test case 1: Empty tempFile should return nil
	err := CleanStream("")
	if err != nil {
		t.Errorf("CleanStream(\"\") should return nil, got %v", err)
	}

	// Test case 2: Test with actual temporary directory
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "streamtest")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Create a file in the temp directory
	tempFile := filepath.Join(tempDir, "test_stream")
	file, err := os.Create(tempFile)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	file.Close()

	// Verify the file exists
	if _, err := os.Stat(tempFile); os.IsNotExist(err) {
		t.Fatalf("Temp file should exist before cleanup")
	}

	// Test CleanStream
	err = CleanStream(tempFile)
	if err != nil {
		t.Errorf("CleanStream should succeed, got error: %v", err)
	}

	// Verify the directory (and file) is removed
	if _, err := os.Stat(tempDir); !os.IsNotExist(err) {
		t.Errorf("Temp directory should be removed after CleanStream")
	}
}