package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	_store := InitStore()
	testStoreService = _store
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUId := "eefe8123-e963-4360-bf33-dd37c54ff803"
	shortURL := "Jsz4k57oAX"

	SaveUrlMapping(shortURL, initialLink, userUUId)

	assert.Equal(t, initialLink, RetrieveInitialUrl(shortURL))
}