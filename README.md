# A Not-Very-Useful Program

## Orientation

This program helps you name your Ruby Gems from lyrics of a song ("jammies" is not taken).
You give it an artist and song, and it tells you what words aren't used for existing gems, which words are, and the
percent coverage of your song.

For example, `gemmer -artist="weird" -song="eat it"` will show you that, alas, "tuna" is already used, but "casserole" is not.
~38% of `Eat It` is covered by existing Ruby gems. It looks like this:

![Gemmer Screenshot](/img/gemmer_ss.png)

This program is written in [Go](http://golang.org). You can build it from source or [download a release](https://github.com/marcesher/gemmer/releases)


## Origin Story

I was driving in my car

Listening to "Ice Ice Baby"

And thought:

> "I bet there are some good Ruby Gem names in this song"

## Is this serious?

I used this as a toy project while starting to learn Go. The world doesn't need another TODO or Blog application. It doesn't need this, either, but I got to read Fugazi and Gorilla Biscuits lyrics while building it.

It should behave correctly. However, it is not production-quality software.

## Known limitations

1. Due to severe restrictions on the ChartLyrics API, you can run this program every 30 seconds or so. More frequently and you get errors
1. ChartLyrics doesn't publish many [Avett Brothers](http://www.theavettbrothers.com/) lyrics. Consequently, most suggestions will remain banal. As a consolation, try `gemmer -artist="ll cool j" -song="boomin"`
1. This program does nothing to elevate the nature of gem naming. Were it more ambitious, it might use as its sources Chaucer, Milton, Tennyson, Coleridge, Byron, Keats. It does, however, occasionally surprise with fitting names such as "youngbut"

## Roadmap

This is guaran-damn-teed abandonware.

## Acknowledgements

Thanks to ChartLyrics.com for publishing an API that includes full lyric text