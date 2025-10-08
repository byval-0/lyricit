# Lyricit

Lyricit is a simple command-line tool that displays lyrics from an LRC file in a karaoke-style. It's written in Go and supports two modes: simple line-by-line display and a character-by-character streaming mode.

## Features

*   Parses LRC files with timestamps.
*   Displays lyrics synchronized with the timestamps.
*   Simple mode: Prints the entire lyric line at the scheduled time.
*   Stream mode: Prints each character of the lyric with a calculated delay to simulate a karaoke effect.

## Installation

To build the project, you need to have Go installed.

1.  Clone the repository:
    ```bash
    git clone https://github.com/byval/lyricit.git
    cd lyricit
    ```
2.  Build the executable:
    ```bash
    go build
    ```

## Usage

Run lyricit from the command line, providing the path to an LRC file.

### Simple Mode

This is the default mode. It will print each line of the lyrics at the time specified in the LRC file.

```bash
./lyricit sample.lrc
```

### Stream Mode

Enable stream mode with the `-stream` flag. This will print the lyrics character by character, giving a karaoke-like effect.

```bash
./lyricit -stream sample.lrc
```

### Notification Mode

Enable notification mode with the `-notify` flag. This will send a desktop notification for each lyric line. The notification title will be formatted as "{artist} - {title}", which are extracted from the LRC file.

```bash
./lyricit -notify sample.lrc
```

## Example

Given a `sample.lrc` file with the following content:

```
[ar:Test Artist]
[ti:Test Title]
[00:01.00]Hello
[00:02.50]World
[00:04.00]This is a sample LRC file.
[00:06.00]Enjoy!
```

Running `./lyricit -stream sample.lrc` will display:

*   At 1 second: "Hello" (character by character)
*   At 2.5 seconds: "World" (character by character)
*   At 4 seconds: "This is a sample LRC file." (character by character)
*   At 6 seconds: "Enjoy!" (character by character)

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## License

This project is open source and available under the [MIT License](LICENSE).
