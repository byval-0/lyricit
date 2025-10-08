package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
)

type LyricLine struct {
	Time  time.Duration
	Lyric string
}

func parseLRC(file *os.File) (string, string, []LyricLine, error) {
	var lines []LyricLine
	var artist, title string
	scanner := bufio.NewScanner(file)
	metaRe := regexp.MustCompile(`\[(ar|ti):(.*)\]`)
	lyricRe := regexp.MustCompile(`\[(\d{2}):(\d{2})\.(\d{2})\](.*)`)

	for scanner.Scan() {
		line := scanner.Text()
		if metaMatch := metaRe.FindStringSubmatch(line); len(metaMatch) == 3 {
			switch metaMatch[1] {
			case "ar":
				artist = strings.TrimSpace(metaMatch[2])
			case "ti":
				title = strings.TrimSpace(metaMatch[2])
			}
		} else if lyricMatch := lyricRe.FindStringSubmatch(line); len(lyricMatch) == 5 {
			min, _ := strconv.Atoi(lyricMatch[1])
			sec, _ := strconv.Atoi(lyricMatch[2])
			cs, _ := strconv.Atoi(lyricMatch[3])
			lyric := strings.TrimSpace(lyricMatch[4])

			duration := time.Duration(min)*time.Minute + time.Duration(sec)*time.Second + time.Duration(cs)*10*time.Millisecond
			lines = append(lines, LyricLine{Time: duration, Lyric: lyric})
		}
	}

	return artist, title, lines, scanner.Err()
}

func main() {
	streamMode := flag.Bool("stream", false, "Enable streaming mode")
	notifyMode := flag.Bool("notify", false, "Enable notification mode")
	beeep.AppName = "Lyricit"
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run main.go [-stream] [-notify] <file.lrc>")
		return
	}

	filePath := flag.Arg(0)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	artist, title, lyrics, err := parseLRC(file)
	if err != nil {
		fmt.Printf("Error parsing LRC file: %v\n", err)
		return
	}

	if len(lyrics) == 0 {
		fmt.Println("No lyrics found in the file.")
		return
	}

	startTime := time.Now()
	for i, line := range lyrics {
		time.Sleep(line.Time - (time.Since(startTime)))

		if *streamMode {
			var nextLineTime time.Duration
			if i+1 < len(lyrics) {
				nextLineTime = lyrics[i+1].Time
			} else {
				nextLineTime = line.Time + time.Second
			}

			duration := nextLineTime - line.Time
			lyricLength := len(line.Lyric)
			charDelay := time.Duration(0)
			if lyricLength != 0 {
				charDelay = duration / time.Duration(len(line.Lyric))
			}

			if charDelay > 50*time.Millisecond {
				charDelay = 50 * time.Millisecond
			}

			for _, char := range line.Lyric {
				fmt.Printf("%c", char)
				time.Sleep(charDelay)
			}
			fmt.Println()
		} else {
			fmt.Println(line.Lyric)
		}

		if *notifyMode {
			notificationTitle := fmt.Sprintf("%s - %s", artist, title)
			err := beeep.Notify(notificationTitle, line.Lyric, "")
			if err != nil {
				fmt.Printf("Error sending notification: %v\n", err)
			}
		}
	}
}
