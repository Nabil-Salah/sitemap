# Site Mapper

Site Mapper is a backend application that make a map of all links in website you need. It's like the footer map in most famous websites.
## Features

- Shortest Path to reach any place from home (used bfs algorithm).

- Generate map in json format.

- You put the link of any website you want and specify the max depth.

## Installation

Before proceeding with the installation, ensure you have golang installed on your system, you can download and install from the [go guide](https://go.dev/doc/install)

## Usage

To use Site Mapper run

remove <> and write your values.
```bash
go run . -depth=<max depth reach> -url="<website link>"
```

## Testing website

To test Site provided by default run

remove <> and write your values.
```bash
go run . -depth=<max depth reach> 
```