package helper

import (
	"strconv"
	"strings"
)

func ExtractArtistIDFromURL(path string) (int, error) {
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]

	artistID, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return artistID, nil
}
