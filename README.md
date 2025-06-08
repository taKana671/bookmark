# BookmarkSearch

CLI tool to bookmark websites.

# Environment

* Go 1.23.4
* Windows11

# Usage

After cloning this repository, compile the source code with the `go build` command.  
If not compile, you can use `go run main.go <subcommand> <options>` .  
It is necessary to install Go language beforehand.

```
git clone https://github.com/taKana671/bookmark.git
cd bookmark
go build
```
## subcommands and options

To list available subcommands, type `bookmark --help` or `bookmark -h`.

### bookmark add

Add a bookmark.  
Information on bookmarked sites will be saved in a CSV file created on the same level as this CLI tool.
The information saved includes the datetime of the bookmarking, category, site title, and URL.
If no category is specified, the category is saved as `all`.

```
bookmark add -C <category> -U <url>
// bookmark add --category <category> --url <url>
// if not compile
// go run main.go add -C <category> -U <url>

>bookmark add -C golang -U https://go.dev/
bookmarked on 2025-06-08 12:08:05: golang, The Go Programming Language, https://go.dev/
```
### bookmark search

Search for saved bookmarks.
If neither category nor keyword are specified, all bookmarks will be displayed.
If a keyword is specified, it will determine if the keyword is included in the site's title.

```
bookmark search -C <category> -K <keyword>
// bookmark search --category <category> --keyword <keyword>

>bookmark search
1 [2025-06-08 12:08:05 golang The Go Programming Language https://go.dev/]
2 [2025-06-08 12:15:34 python 3.13.4 Documentation https://docs.python.org/3/

>bookmark search -C golang
1 [2025-06-08 12:08:05 golang The Go Programming Language https://go.dev/]

>bookmark search -K Documentation
1 [2025-06-08 12:15:34 python 3.13.4 Documentation https://docs.python.org/3/]

>bookmark search -C golang -K Go
1 [2025-06-08 12:08:05 golang The Go Programming Language https://go.dev/]

C:\Users\Kanae\Desktop>bookmark search -C SQL
bookmarks not found; category: SQL; keyword:
```

### bookmark open

Open the website by specifying the `bookmark seach` result number.

```
bookmark open -N <number>
// bookmark open --no <number>

>bookmark search
1 [2025-06-08 12:08:05 golang The Go Programming Language https://go.dev/]
2 [2025-06-08 12:15:34 python 3.13.4 Documentation https://docs.python.org/3/]

>bookmark open -N 1
open: https://go.dev/
```

### bookmark delete

Delete the bookmark by specifying the `bookmark seach` result number.

```
bookmark delete -N <number>
// bookmark delete --no <number>

>bookmark search
1 [2025-06-08 12:08:05 golang The Go Programming Language https://go.dev/]
2 [2025-06-08 12:15:34 python 3.13.4 Documentation https://docs.python.org/3/]

>bookmark delete -N 1
deleted a bookmark: The Go Programming Language, https://go.dev/
C:\Users\Kanae\Desktop>bookmark search
1 [2025-06-08 12:15:34 python 3.13.4 Documentation https://docs.python.org/3/]
```



