package fetcher

import (
	"fmt"
	"testing"
)

func TestSearchWithLuck(t *testing.T) {
	lyric, err := SearchWithLuck("eric", "follow the leader")

	if err != nil {
		t.Error("Error should not be nil but was %v", err)
	}
	fmt.Printf("Identifier: %#v\n", lyric.LyricId)
	fmt.Printf("Source: Artist - %v;  Song - %v; Url - %v\n\n", lyric.Artist, lyric.Song, lyric.Url)

	if lyric.Artist != "Eric B. & Rakim" || lyric.Song != "Follow the Leader" || lyric.LyricId.LyricId != "19141" {
		t.Error("Search result not as expected. Expected a Lyric for Eric B. and Rakim Follow the leader but instead got this bullllllshit: %v", lyric)
	}
	if len(lyric.LyricText) == 0 {
		t.Error("Lyric Text is empty. Should've been the best rhyme ever but was not.")
	}
}

//SearchLyrics not currently in use in the application
//func TestSearchLyrics(t *testing.T) {
//	results := SearchLyrics("ice ice baby too cold too cold")
//	fmt.Println(results.SearchResults)
//	if len(results.SearchResults) == 0 {
//		t.Error("Should not be empty")
//	}
//	for _, result := range results.SearchResults {
//		fmt.Printf("Identifier: %#v\n", result.LyricId)
//		fmt.Printf("Source: Artist - %v;  Song - %v; Url - %v\n\n", result.Artist, result.Song, result.Url)
//	}
//}

//sanity checkers while I build this thing
func TestGetLyricFromFile(t *testing.T) {
	lyric := GetLyricFromFile()
	fmt.Println(lyric)
	fmt.Printf("Identifier: %#v\n", lyric.LyricId)
	fmt.Printf("Source: Artist - %v;  Song - %v; Url - %v\n\n", lyric.Artist, lyric.Song, lyric.Url)
}

func TestSearchWithLuckFromFile(t *testing.T) {
	lyric := SearchWithLuckFromFile()
	fmt.Println(lyric)
	fmt.Printf("Identifier: %#v\n", lyric.LyricId)
	fmt.Printf("Source: Artist - %v;  Song - %v; Url - %v\n\n", lyric.Artist, lyric.Song, lyric.Url)
}

func TestSearchLyricsFromFile(t *testing.T) {
	results := SearchLyricsFromFile()
	for _, result := range results.SearchResults {
		fmt.Printf("Identifier: %#v\n", result.LyricId)
		fmt.Printf("Source: Artist - %v;  Song - %v; Url - %v\n\n", result.Artist, result.Song, result.Url)
	}
}
