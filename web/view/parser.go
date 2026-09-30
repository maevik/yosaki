package view

import (
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func MustParseFiles(patterns ...string) *template.Template {
	templates := template.New("")

	for _, pattern := range patterns {
		matches, err := matchFiles(pattern)
		if err != nil {
			panic(pattern + ": " + err.Error())
		}
		if len(matches) == 0 {
			panic("no templates found: " + pattern)
		}

		for _, path := range matches {
			content, err := os.ReadFile(path)
			if err != nil {
				panic(path + ": " + err.Error())
			}

			name := strings.TrimPrefix(path, "tmpl/")
			name = strings.TrimSuffix(name, ".tmpl")

			if _, err := templates.New(name).Parse(string(content)); err != nil {
				panic(path + ": " + err.Error())
			}
		}
	}

	return templates
}

func matchFiles(pattern string) ([]string, error) {
	if !strings.Contains(pattern, "**") {
		return filepath.Glob(pattern)
	}

	parts := strings.SplitN(pattern, "**", 2)
	root := strings.TrimSuffix(parts[0], string(os.PathSeparator))
	if root == "" {
		root = "."
	}
	suffix := strings.TrimPrefix(parts[1], string(os.PathSeparator))

	var matches []string
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}

		ok, err := filepath.Match(suffix, filepath.Base(path))
		if err != nil {
			return err
		}
		if ok {
			matches = append(matches, path)
		}

		return nil
	})
	sort.Strings(matches)
	return matches, err
}
