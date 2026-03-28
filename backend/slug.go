package main

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"
)

var slugCleanupPattern = regexp.MustCompile(`[^a-z0-9]+`)

func slugify(input string) string {
	base := strings.ToLower(strings.TrimSpace(input))
	base = slugCleanupPattern.ReplaceAllString(base, "-")
	base = strings.Trim(base, "-")
	if base == "" {
		return "post"
	}

	return base
}

func generateUniquePostSlug(db *sql.DB, source string, excludePostID int) (string, error) {
	baseSlug := slugify(source)
	candidate := baseSlug
	suffix := 2

	for {
		var count int
		if excludePostID > 0 {
			err := db.QueryRow(
				"SELECT COUNT(1) FROM posts WHERE slug = ? AND id != ?",
				candidate,
				excludePostID,
			).Scan(&count)
			if err != nil {
				return "", err
			}
		} else {
			err := db.QueryRow("SELECT COUNT(1) FROM posts WHERE slug = ?", candidate).Scan(&count)
			if err != nil {
				return "", err
			}
		}

		if count == 0 {
			return candidate, nil
		}

		candidate = fmt.Sprintf("%s-%d", baseSlug, suffix)
		suffix++
	}
}
