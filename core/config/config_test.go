package config

import (
	"os"
	"testing"
	"time"
)

func TestNewConfigDefaults(t *testing.T) {
	cfg := New()

	if cfg.SaveInterval != 10*time.Second {
		t.Errorf("expected SaveInterval to be %v, got %v", 10*time.Second, cfg.SaveInterval)
	}

	if cfg.Filename != "sessionData.json" {
		t.Errorf("expected Filename to be %v, got %v", "sessionData.json", cfg.Filename)
	}

	if cfg.Port != ":1378" {
		t.Errorf("expected Port to be %v, got %v", ":1378", cfg.Port)
	}
}

func TestNewConfigWithEnvVariables(t *testing.T) {
	os.Setenv("SAVE_INTERVAL", "15s")
	os.Setenv("FILENAME", "testData.json")
	os.Setenv("PORT", ":8080")

	cfg := New()

	expectedSaveInterval, _ := time.ParseDuration("15s")
	if cfg.SaveInterval != expectedSaveInterval {
		t.Errorf("expected SaveInterval to be %v, got %v", expectedSaveInterval, cfg.SaveInterval)
	}

	if cfg.Filename != "testData.json" {
		t.Errorf("expected Filename to be %v, got %v", "testData.json", cfg.Filename)
	}

	if cfg.Port != ":8080" {
		t.Errorf("expected Port to be %v, got %v", ":8080", cfg.Port)
	}

	os.Unsetenv("SAVE_INTERVAL")
	os.Unsetenv("FILENAME")
	os.Unsetenv("PORT")
}
