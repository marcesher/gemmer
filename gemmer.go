package main

import (
	"flag"
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/marcesher/gemmer/fetcher"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

type SortedByLen struct{ sort.StringSlice }

func (p SortedByLen) Less(i, j int) bool { return len(p.StringSlice[i]) < len(p.StringSlice[j]) }

//NewSetFromStrings converts an array of strings into a Set
func NewSetFromStrings(s []string) mapset.Set {
	set := mapset.NewSet()
	for _, item := range s {
		set.Add(item)
	}
	return set
}

// NewStringsFromSet converts a Set into an array of strings
// TODO: there has got to be an easier way to get a keyset out of this thing
func NewStringsFromSet(s mapset.Set) []string {
	var keys []string
	for k := range s {
		str, _ := k.(string)
		keys = append(keys, str)
	}
	return keys
}

//Coverage returns a set of words that are not currently used for gem names. It also returns the percentage of your song that is covered by existing gem names
func Coverage(lyrics, gems []string) ([]string, []string, float64) {
	lyricset := NewSetFromStrings(lyrics)
	gemset := NewSetFromStrings(gems)
	diff := lyricset.Difference(gemset)
	intersect := lyricset.Intersect(gemset)
	return NewStringsFromSet(diff), NewStringsFromSet(intersect), (float64(intersect.Cardinality()) / float64(lyricset.Cardinality()) * 100)
}

//PrepareText lowercases and strips punctuation for any input text (lyrics, poem, tax code)
//Hilarity ensues here. For example, in Special Ed's "I Got it Made", "I'm kinda young--but my tongue speaks maturity" turns into "youngbut", which is a great gem name
func PrepareText(text string) string {
	re := regexp.MustCompile("[^a-zA-Z0-0\\s]")
	clean := re.ReplaceAllString(text, "")
	return strings.ToLower(clean)
}

func main() {
	artist := flag.String("artist", "vanilla ice", "The song artist")
	song := flag.String("song", "ice ice baby", "The song")
	flag.Parse()

	lyric, err := fetcher.SearchWithLuck(*artist, *song)

	if err != nil {
		fmt.Printf(`Most unfortunately, an error has occurred.
			Here's the full error message, which will probably not be useful: %v
			If it's a 'connection reset' error, it's because ChartLyrics limits you
			to one request every 30 seconds or so. ***Slow your roll, son.***
			`, err)
		return
	}

	if len(lyric.LyricText) == 0 {
		fmt.Printf("Could not find a matching song for artist [%v] and song [%v]. Try -artist='vanilla ice' -song='ice ice baby' or -artist='operation ivy' -song='freeze up'\n\n", *artist, *song)
		return
	}
	clean := PrepareText(lyric.LyricText)
	gemlist, _ := ioutil.ReadFile("gemlist.txt")
	diff, intersect, cov := Coverage(strings.Fields(clean), strings.Fields(string(gemlist)))
	fmt.Printf("\n\nSource: Artist - %v;  Song - %v; Url - %v\n\n", lyric.Artist, lyric.Song, lyric.Url)
	fmt.Printf("Your song is %f%% covered by the existing gem list\n\n", cov)
	fmt.Printf("Words found in both the song lyrics and existing gems: %v\n\n", intersect)
	fmt.Printf("Try out these names for your next gem!\n\n")
	sort.Sort(sort.Reverse(SortedByLen(SortedByLen{diff[0:]})))
	fmt.Printf("%v\n\n", diff)
}
