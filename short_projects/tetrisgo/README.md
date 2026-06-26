# Terminal Tetris (Go)

A classic Tetris game that runs in your terminal, built with Go and [tcell](https://github.com/gdamore/tcell).

## Features

- All 7 classic tetrominoes with rotation + basic wall kicks
- Ghost piece (hard drop preview)
- Next piece preview
- Line clearing with classic scoring (100/300/500/800)
- Progressive speed (levels)
- Pause / restart

## Controls

| Key          | Action       |
|--------------|--------------|
| ← →          | Move left/right |
| ↓            | Soft drop    |
| ↑ or Z or X  | Rotate       |
| Space        | Hard drop    |
| P            | Pause        |
| R            | Restart (after game over) |
| Q or Esc     | Quit         |

## Run

```bash
go run .
# or
go build -o tetris && ./tetris
```

Requires a terminal that supports color and cursor addressing (most modern terminals).

Enjoy!
