# 🗒️ Sticky Notes (Go + HTMX)

Realtime sticky note board using Go, HTMX, and SSE.

## Setup

**Mac/Linux**

1. Install Go 1.22+
2. Run `go install github.com/air-verse/air@latest`
3. Add `~/go/bin` to your PATH if needed
4. Start in `/sticky-notes`
5. In the project root, run `air` to start the server

**Windows**

1. Install Go from golang.org
2. Run `go install github.com/air-verse/air@latest`
3. Add `C:\Users\YourName\go\bin` to your system PATH
4. Start in `\sticky-notes`
5. Run `.\air.ps1` from PowerShell

## Structure

- `cmd/web` — app entry
- `internal/notes` — note model + store
- `internal/server` — routes + SSE
- `templates/` — HTML

## Run

Visit `http://localhost:8080` in your browser.  
Open multiple tabs to see realtime syncing.
