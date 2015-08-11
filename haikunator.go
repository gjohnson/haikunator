package haikunator

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var adjectives = []string{
	"autumn", "hidden", "bitter", "misty", "silent", "empty", "dry", "dark", "summer",
	"icy", "delicate", "quiet", "white", "cool", "spring", "winter", "patient",
	"twilight", "dawn", "crimson", "wispy", "weathered", "blue", "billowing",
	"broken", "cold", "damp", "falling", "frosty", "green", "long", "late", "lingering",
	"bold", "little", "morning", "muddy", "old", "red", "rough", "still", "small",
	"sparkling", "throbbing", "shy", "wandering", "withered", "wild", "black",
	"young", "holy", "solitary", "fragrant", "aged", "snowy", "proud", "floral",
	"restless", "divine", "polished", "ancient", "purple", "lively", "nameless",
}

var nouns = []string{
	"waterfall", "river", "breeze", "moon", "rain", "wind", "sea", "morning",
	"snow", "lake", "sunset", "pine", "shadow", "leaf", "dawn", "glitter", "forest",
	"hill", "cloud", "meadow", "sun", "glade", "bird", "brook", "butterfly",
	"bush", "dew", "dust", "field", "fire", "flower", "firefly", "feather", "grass",
	"haze", "mountain", "night", "pond", "darkness", "snowflake", "silence",
	"sound", "sky", "shape", "surf", "thunder", "violet", "water", "wildflower",
	"wave", "water", "resonance", "sun", "wood", "dream", "cherry", "tree", "fog",
	"frost", "voice", "paper", "frog", "smoke", "star",
}

func Haikunate() (string, error) {
	a, err := randomAdjective()
	if err != nil {
		return "", err
	}

	b, err := randomNoun()
	if err != nil {
		return "", err
	}

	c, err := randomNumber()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s-%d", a, b, c), nil
}

func randomAdjective() (string, error) {
	n, err := randomRange(0, len(adjectives))
	if err != nil {
		return "", err
	}
	return adjectives[n], nil
}

func randomNoun() (string, error) {
	n, err := randomRange(0, len(nouns))
	if err != nil {
		return "", err
	}
	return nouns[n], nil
}

func randomNumber() (int, error) {
	n, err := randomRange(1000, 9999)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func randomRange(min, max int) (int, error) {
	bytes, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0, err
	}
	return min + int(bytes.Int64()), nil
}
