package gravatarurl

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/url"
	"strconv"
	"strings"
)

// https://docs.gravatar.com/api/avatars/images/#default-image
type DefaultImage string

// https://docs.gravatar.com/api/avatars/images/#rating
type Rating string

// https://docs.gravatar.com/api/avatars/images/#size
type Size uint16

const (
	FourOhFour    DefaultImage = "404"
	Blank         DefaultImage = "blank"
	Colors        DefaultImage = "colors"
	Identicon     DefaultImage = "identicon"
	Initials      DefaultImage = "initials"
	MonsterId     DefaultImage = "monsterid"
	MysteryPerson DefaultImage = "mp"
	Retro         DefaultImage = "retro"
	RoboHash      DefaultImage = "robohash"
	Wavatar       DefaultImage = "wavatar"
)

const (
	G  Rating = "g"
	PG Rating = "pg"
	R  Rating = "r"
	X  Rating = "x"
)

type Options struct {
	DefaultImage
	Rating
	Size
}

func GravatarUrl(identifier string, options Options) string {
	if identifier == "" {
		log.Fatalln("Please specify an identifier, such as an email address")
	}

	if strings.Contains(identifier, "@") {
		identifier = strings.TrimSpace(strings.ToLower(identifier))
	}

	baseUrl, err := url.Parse("https://gravatar.com/avatar/")

	if err != nil {
		log.Fatalln(err)
	}

	hash := sha256.New()
	hash.Write([]byte(identifier))

	baseUrl.Path += hex.EncodeToString(hash.Sum(nil))

	v := url.Values{}

	if options.DefaultImage != "" {
		v.Set("default", string(options.DefaultImage))
	}

	if options.Rating != "" {
		v.Set("rating", string(options.Rating))
	}

	if options.Size != 0 {
		v.Set("size", strconv.Itoa(int(options.Size)))
	}

	baseUrl.RawQuery = v.Encode()

	return baseUrl.String()
}
