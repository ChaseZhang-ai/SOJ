package app

import (
	"errors"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"SOJ/internal/config"
)

func TestLoadLanguagesRejectsEmptyCatalogAndMissingDirectory(t *testing.T) {
	t.Run("empty directory defines no languages", func(t *testing.T) {
		dir := t.TempDir()
		cfg := config.Config{LanguagesDir: dir}

		catalog, err := loadLanguages(cfg)
		if err == nil {
			t.Fatal("loadLanguages accepted a directory without language profiles")
		}
		if catalog != nil {
			t.Fatalf("catalog = %v, want nil", catalog)
		}
		if !strings.Contains(err.Error(), dir) {
			t.Fatalf("error %q does not identify the configured directory %q", err, dir)
		}
		if !strings.Contains(err.Error(), "no languages") {
			t.Fatalf("error %q does not mention the absence of languages", err)
		}
		if strings.Contains(err.Error(), "load languages from") {
			t.Fatalf("error %q reported a load failure for a directory that parsed successfully", err)
		}
		if errors.Is(err, fs.ErrNotExist) {
			t.Fatalf("empty catalog error %v was reported as a filesystem error", err)
		}
	})

	t.Run("missing directory reports the filesystem cause", func(t *testing.T) {
		dir := filepath.Join(t.TempDir(), "missing")
		cfg := config.Config{LanguagesDir: dir}

		catalog, err := loadLanguages(cfg)
		if err == nil {
			t.Fatal("loadLanguages accepted a missing directory")
		}
		if catalog != nil {
			t.Fatalf("catalog = %v, want nil", catalog)
		}
		if !strings.Contains(err.Error(), dir) {
			t.Fatalf("error %q does not include the configured path %q", err, dir)
		}
		if !errors.Is(err, fs.ErrNotExist) {
			t.Fatalf("error %v does not preserve fs.ErrNotExist", err)
		}
	})
}
