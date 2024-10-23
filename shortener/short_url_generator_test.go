package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	originalLink_1 := "https://campaigntracker.io/blog/the-power-of-long-urls-unveiling-the-secrets-behind-lengthy-web-addresses"
	shortLink_1 := GenerateShortLink(originalLink_1, UserId)

	originalLink_2 := "https://docs.imperva.com/bundle/on-premises-knowledgebase-reference-guide/page/abnormally_long_url.htm"
	shortLink_2 := GenerateShortLink(originalLink_2, UserId)

	originalLink_3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink_3 := GenerateShortLink(originalLink_3, UserId)


	assert.Equal(t, shortLink_1, "a7M4N2")
	assert.Equal(t, shortLink_2, "7SV2oR")
	assert.Equal(t, shortLink_3, "dhZTay")
}