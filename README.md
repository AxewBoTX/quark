# quark

Self host chat app in 1 file

## Tech Stack

- [labstack/echo](https://github.com/labstack/echo) - High performance, minimalist Go web framework
- [golang.org/x/net/websocket](https://pkg.go.dev/golang.org/x/net/websocket) - Golang based client-server WebSocket proptocol implementation
- [go-task/task](https://github.com/go-task/task) - A task runner / simpler Make alternative written in Go
- [embed](https://pkg.go.dev/embed) - File embedding in go programs
- [cosmtrek/air](https://github.com/cosmtrek/air) - Live reload for Go apps
- [charmbracelet/log](https://github.com/charmbracelet/log) - A minimal, colorful Go logging library
- [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions for nice terminal layouts
- [pelletier/go-toml](https://github.com/pelletier/go-toml) - Go library for the TOML file format
- [sqlite](https://pkg.go.dev/modernc.org/sqlite) - Golang SQLite database driver
- [google/uuid](https://github.com/google/uuid) - Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services
- [bun](https://bun.sh/) - Incredibly fast JavaScript runtime, bundler, test runner, and package manager – all in one
- [a-h/templ](https://github.com/a-h/templ) - A language for writing HTML user interfaces in Go
- [tailwindcss](https://tailwindcss.com/) - A utility-first CSS framework for rapid UI development
- [bigskysoftware/htmx](https://github.com/bigskysoftware/htmx) - high power tools for HTML
- [go-resty/resty](https://github.com/go-resty/resty) - Simple HTTP and REST client library for Go
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Provos and Mazières's bcrypt adaptive hashing algorithm implementation

## How to setup

1. Install the required technologies:-
   - [golang:1.22.0](https://tip.golang.org/doc/go1.22)
   - [bun:1.1.10](https://github.com/oven-sh/bun/releases/tag/bun-v1.1.10)
2. Clone the github repository and `cd` into it
   ```
   git clone https://github.com/axewbotx/quark && cd quark
   ```
3. Install [go-task](https://github.com/go-task/task)
   ```
   sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin && export PATH=$PATH:~/.local/bin
   ```
4. Build the project
   ```
   task build
   ```
5. The `quark_server` and `quark_client` binaries will be awailable in newly created `build` directory
