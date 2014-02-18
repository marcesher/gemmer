/*
Uses the ChartLyrics API: http://www.chartlyrics.com/api.aspx

On the plus side, this API returns full lyrics. Unfortunately, it also severely throttles requests. You can run
this program once every 30 seconds or so; otherwise, you get 'connection reset' errors

Search by lyric text: http://api.chartlyrics.com/apiv1.asmx/SearchLyricText?lyricText=bob dylan sittin on top of the world
Get lyric by id: http://api.chartlyrics.com/apiv1.asmx/GetLyric?lyricId=387&lyricCheckSum=3b229d07319ac6babf0b0126e4140d21
Search lyric direct... you rely on chartlyrics' guesses: http://api.chartlyrics.com/apiv1.asmx/SearchLyricDirect?artist=eric&song=follow%20the%20leader
*/

package fetcher

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
)

const API_URL = "http://api.chartlyrics.com/apiv1.asmx"
const SEARCH_URL = API_URL + "/SearchLyricText?lyricText="
const SEARCH_DIRECT_URL = API_URL + "/SearchLyricDirect?"
const GET_URL = API_URL + "/GetLyric?"

type LyricId struct {
	LyricId       string
	LyricChecksum string
}

type Lyric struct {
	LyricId
	Song      string `xml:"LyricSong"`
	Artist    string `xml:"LyricArtist"`
	Url       string `xml:"LyricUrl"`
	LyricText string `xml:"Lyric"`
}

type LyricSearchItem struct {
	LyricId
	Artist string
	Song   string
	Url    string `xml:"SongUrl"`
}

type LyricSearch struct {
	SearchResults []LyricSearchItem `xml:"SearchLyricResult"`
}

//TODO: DRY this thing.

//SearchWithLuck accepts an artist and song and returns the first match, whose relevance is decided by ChartLyrics. Returns a Lyric struct. Will throw a 'connection reset' error if called more than once every 30 seconds or so
func SearchWithLuck(artist, song string) (Lyric, error) {
	url := SEARCH_DIRECT_URL + "artist=" + url.QueryEscape(artist) + "&song=" + url.QueryEscape(song)
	res, err := http.Get(url)
	if err != nil {
		return Lyric{}, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Lyric{}, err
	}
	lyric := Lyric{}
	err = xml.Unmarshal(body, &lyric)
	if err != nil {
		return Lyric{}, err
	}
	return lyric, nil
}

//Search and Get are not currently usable b/c of the restrictions ChartLyrics places on its API usage (can't search then get)
func SearchLyrics(search_terms string) LyricSearch {
	url := SEARCH_URL + url.QueryEscape(search_terms)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	results := LyricSearch{}
	err = xml.Unmarshal(body, &results)
	if err != nil {
		panic(err)
	}
	return results
}

//Search and Get are not currently usable b/c of the restrictions ChartLyrics places on its API usage (can't search then get)
func GetLyric(id LyricId) Lyric {
	url := GET_URL + "lyricid=" + id.LyricId + "&lyricCheckSum=" + id.LyricChecksum
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	lyric := Lyric{}
	err = xml.Unmarshal(body, &lyric)
	if err != nil {
		panic(err)
	}
	return lyric
}

//helpers while I build this silly thing

func SearchWithLuckFromFile() Lyric {
	xmldata, err := ioutil.ReadFile("../mock_endpoints/search_lyric_direct.xml")
	if err != nil {
		panic(err)
	}
	lyric := Lyric{}
	err = xml.Unmarshal(xmldata, &lyric)
	if err != nil {
		panic(err)
	}
	return lyric
}

func SearchLyricsFromFile() LyricSearch {
	xmldata, err := ioutil.ReadFile("../mock_endpoints/search.xml")
	if err != nil {
		panic(err)
	}
	results := LyricSearch{}
	err = xml.Unmarshal(xmldata, &results)
	if err != nil {
		panic(err)
	}
	return results
}

func GetLyricFromFile() Lyric {
	xmldata, err := ioutil.ReadFile("../mock_endpoints/lyric.xml")
	if err != nil {
		panic(err)
	}
	lyric := Lyric{}
	err = xml.Unmarshal(xmldata, &lyric)
	if err != nil {
		panic(err)
	}
	return lyric
}
