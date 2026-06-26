package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Game constants
const (
	boardWidth  = 10
	boardHeight = 20

	cellWidth  = 2 // ~square blocks in terminal (chars are ~2x taller than wide)
	cellHeight = 1

	previewCols = 4
	previewRows = 4
)

// Piece types 1-7
const (
	Empty = 0
	I     = 1
	O     = 2
	T     = 3
	S     = 4
	Z     = 5
	J     = 6
	L     = 7
)

// Tetromino definitions: [piece][rotation][block][x,y]
var tetrominoes = [][][][2]int{
	{}, // index 0 unused
	// I
	{
		{{0, 1}, {1, 1}, {2, 1}, {3, 1}},
		{{2, 0}, {2, 1}, {2, 2}, {2, 3}},
		{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
		{{1, 0}, {1, 1}, {1, 2}, {1, 3}},
	},
	// O
	{
		{{1, 0}, {2, 0}, {1, 1}, {2, 1}},
		{{1, 0}, {2, 0}, {1, 1}, {2, 1}},
		{{1, 0}, {2, 0}, {1, 1}, {2, 1}},
		{{1, 0}, {2, 0}, {1, 1}, {2, 1}},
	},
	// T
	{
		{{1, 0}, {0, 1}, {1, 1}, {2, 1}},
		{{1, 0}, {1, 1}, {2, 1}, {1, 2}},
		{{0, 1}, {1, 1}, {2, 1}, {1, 2}},
		{{1, 0}, {0, 1}, {1, 1}, {1, 2}},
	},
	// S
	{
		{{1, 0}, {2, 0}, {0, 1}, {1, 1}},
		{{1, 0}, {1, 1}, {2, 1}, {2, 2}},
		{{1, 1}, {2, 1}, {0, 2}, {1, 2}},
		{{0, 0}, {0, 1}, {1, 1}, {1, 2}},
	},
	// Z
	{
		{{0, 0}, {1, 0}, {1, 1}, {2, 1}},
		{{2, 0}, {1, 1}, {2, 1}, {1, 2}},
		{{0, 1}, {1, 1}, {1, 2}, {2, 2}},
		{{1, 0}, {0, 1}, {1, 1}, {0, 2}},
	},
	// J
	{
		{{0, 0}, {0, 1}, {1, 1}, {2, 1}},
		{{1, 0}, {2, 0}, {1, 1}, {1, 2}},
		{{0, 1}, {1, 1}, {2, 1}, {2, 2}},
		{{1, 0}, {1, 1}, {0, 2}, {1, 2}},
	},
	// L
	{
		{{2, 0}, {0, 1}, {1, 1}, {2, 1}},
		{{1, 0}, {1, 1}, {1, 2}, {2, 2}},
		{{0, 1}, {1, 1}, {2, 1}, {0, 2}},
		{{0, 0}, {1, 0}, {1, 1}, {1, 2}},
	},
}

// Colors for each piece type
var pieceColors = map[int]tcell.Color{
	I: tcell.ColorAqua,
	O: tcell.ColorYellow,
	T: tcell.ColorFuchsia,
	S: tcell.ColorLime,
	Z: tcell.ColorRed,
	J: tcell.ColorBlue,
	L: tcell.ColorDarkOrange,
}

var pieceNames = map[int]string{
	I: "I-piece",
	O: "O-piece",
	T: "T-piece",
	S: "S-piece",
	Z: "Z-piece",
	J: "J-piece",
	L: "L-piece",
}

func boardPixelWidth() int  { return boardWidth * cellWidth }
func boardPixelHeight() int { return boardHeight * cellHeight }

const (
	layoutGap         = 4
	titleAboveBoard   = 2
	sidebarNextOffset = 15
)

var controlLines = []string{
	"Controls:",
	"← →   Move",
	"↓     Soft drop",
	"↑ / Z Rotate",
	"Space Hard drop",
	"P     Pause",
	"Q     Quit",
}

func sidebarWidth() int {
	w := previewCols*cellWidth + 2
	for _, line := range controlLines {
		if len(line) > w {
			w = len(line)
		}
	}
	if len("NEXT") > w {
		w = len("NEXT")
	}
	return w
}

func layoutWidth() int {
	return boardPixelWidth() + 2 + layoutGap + sidebarWidth()
}

func layoutHeight() int {
	previewH := previewRows*cellHeight + 2
	sidebarBottom := sidebarNextOffset + previewH + 4 + len(controlLines)
	return titleAboveBoard + max(boardPixelHeight()+2, sidebarBottom)
}

func drawString(s tcell.Screen, x, y int, text string, st tcell.Style) {
	for i, ch := range text {
		s.SetContent(x+i, y, ch, nil, st)
	}
}

func drawBlock(s tcell.Screen, x, y int, color tcell.Color, ghost bool) {
	for dy := 0; dy < cellHeight; dy++ {
		for dx := 0; dx < cellWidth; dx++ {
			px, py := x+dx, y+dy
			if ghost {
				s.SetContent(px, py, '░', nil, tcell.StyleDefault.Foreground(tcell.ColorDarkGray))
				continue
			}
			st := tcell.StyleDefault.Background(color).Foreground(color)
			s.SetContent(px, py, '█', nil, st)
		}
	}
}

func pieceBounds(piece, rot int) (minX, maxX, minY, maxY int) {
	minX, minY = 4, 4
	maxX, maxY = -1, -1
	for _, b := range tetrominoes[piece][rot] {
		if b[0] < minX {
			minX = b[0]
		}
		if b[0] > maxX {
			maxX = b[0]
		}
		if b[1] < minY {
			minY = b[1]
		}
		if b[1] > maxY {
			maxY = b[1]
		}
	}
	return
}

func drawBox(s tcell.Screen, x, y, w, h int, st tcell.Style) {
	s.SetContent(x, y, '┌', nil, st)
	for i := 1; i < w-1; i++ {
		s.SetContent(x+i, y, '─', nil, st)
	}
	s.SetContent(x+w-1, y, '┐', nil, st)

	for row := 1; row < h-1; row++ {
		s.SetContent(x, y+row, '│', nil, st)
		s.SetContent(x+w-1, y+row, '│', nil, st)
	}

	s.SetContent(x, y+h-1, '└', nil, st)
	for i := 1; i < w-1; i++ {
		s.SetContent(x+i, y+h-1, '─', nil, st)
	}
	s.SetContent(x+w-1, y+h-1, '┘', nil, st)
}

type Game struct {
	screen tcell.Screen
	board  [boardHeight][boardWidth]int

	currentPiece int
	currX, currY int
	currRot      int

	nextPiece int

	score int
	lines int
	level int

	gameOver  bool
	paused    bool
	lastDrop  time.Time
	dropSpeed time.Duration

	quit bool
}

func NewGame(s tcell.Screen) *Game {
	g := &Game{
		screen:    s,
		level:     1,
		dropSpeed: 500 * time.Millisecond,
	}
	g.nextPiece = g.randomPiece()
	g.spawnNewPiece()
	g.lastDrop = time.Now()
	return g
}

func (g *Game) randomPiece() int {
	return rand.Intn(7) + 1
}

func (g *Game) spawnNewPiece() {
	g.currentPiece = g.nextPiece
	g.nextPiece = g.randomPiece()
	g.currX = boardWidth/2 - 2
	g.currY = 0
	g.currRot = 0

	// Adjust spawn position per piece type
	switch g.currentPiece {
	case O:
		g.currX = boardWidth/2 - 1
	case I:
		g.currY = -1
	}

	if g.collides(g.currentPiece, g.currX, g.currY, g.currRot) {
		g.gameOver = true
	}
}

func (g *Game) collides(piece, x, y, rot int) bool {
	shape := tetrominoes[piece][rot]
	for _, block := range shape {
		bx := x + block[0]
		by := y + block[1]

		if bx < 0 || bx >= boardWidth || by >= boardHeight {
			return true
		}
		if by >= 0 && g.board[by][bx] != Empty {
			return true
		}
	}
	return false
}

// ghostY returns the Y position where the piece would land with hard drop
func (g *Game) ghostY() int {
	y := g.currY
	for !g.collides(g.currentPiece, g.currX, y+1, g.currRot) {
		y++
	}
	return y
}

func (g *Game) move(dx, dy int) bool {
	if g.paused || g.gameOver {
		return false
	}
	newX := g.currX + dx
	newY := g.currY + dy
	if !g.collides(g.currentPiece, newX, newY, g.currRot) {
		g.currX = newX
		g.currY = newY
		return true
	}
	return false
}

func (g *Game) rotate() bool {
	if g.paused || g.gameOver {
		return false
	}
	newRot := (g.currRot + 1) % 4
	if !g.collides(g.currentPiece, g.currX, g.currY, newRot) {
		g.currRot = newRot
		return true
	}
	// Try wall kicks (simple)
	kicks := []int{-1, 1, -2, 2}
	for _, k := range kicks {
		if !g.collides(g.currentPiece, g.currX+k, g.currY, newRot) {
			g.currX += k
			g.currRot = newRot
			return true
		}
	}
	return false
}

func (g *Game) hardDrop() {
	if g.paused || g.gameOver {
		return
	}
	for g.move(0, 1) {
		g.score += 2 // small bonus per cell dropped
	}
	g.lockPiece()
}

func (g *Game) lockPiece() {
	shape := tetrominoes[g.currentPiece][g.currRot]
	for _, block := range shape {
		bx := g.currX + block[0]
		by := g.currY + block[1]
		if by >= 0 && by < boardHeight && bx >= 0 && bx < boardWidth {
			g.board[by][bx] = g.currentPiece
		}
	}
	g.clearLines()
	g.spawnNewPiece()
}

func (g *Game) clearLines() {
	cleared := 0
	for y := 0; y < boardHeight; y++ {
		full := true
		for x := 0; x < boardWidth; x++ {
			if g.board[y][x] == Empty {
				full = false
				break
			}
		}
		if full {
			cleared++
			// Shift everything down
			for yy := y; yy > 0; yy-- {
				for x := 0; x < boardWidth; x++ {
					g.board[yy][x] = g.board[yy-1][x]
				}
			}
			// Clear top row
			for x := 0; x < boardWidth; x++ {
				g.board[0][x] = Empty
			}
			// Re-check same row after shift
			y--
		}
	}

	if cleared > 0 {
		playLineClearSound(cleared)
		g.lines += cleared
		// Classic scoring
		scoreAdd := 0
		switch cleared {
		case 1:
			scoreAdd = 100
		case 2:
			scoreAdd = 300
		case 3:
			scoreAdd = 500
		case 4:
			scoreAdd = 800
		}
		g.score += scoreAdd * g.level

		// Level up every 10 lines
		newLevel := g.lines/10 + 1
		if newLevel > g.level {
			g.level = newLevel
			ms := 500 - (g.level-1)*40
			if ms < 100 {
				ms = 100
			}
			g.dropSpeed = time.Duration(ms) * time.Millisecond
		}
	}
}

func (g *Game) softDrop() {
	if g.move(0, 1) {
		g.score += 1
	}
}

func (g *Game) update() {
	if g.paused || g.gameOver {
		return
	}
	now := time.Now()
	if now.Sub(g.lastDrop) >= g.dropSpeed {
		if !g.move(0, 1) {
			g.lockPiece()
		}
		g.lastDrop = now
	}
}

func (g *Game) drawBoard(offsetX, offsetY int, borderStyle tcell.Style) {
	drawBox(g.screen, offsetX-1, offsetY-1, boardPixelWidth()+2, boardPixelHeight()+2, borderStyle)

	wellStyle := tcell.StyleDefault.Background(tcell.ColorBlack)
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			px := offsetX + x*cellWidth
			py := offsetY + y*cellHeight
			for dy := 0; dy < cellHeight; dy++ {
				for dx := 0; dx < cellWidth; dx++ {
					g.screen.SetContent(px+dx, py+dy, ' ', nil, wellStyle)
				}
			}
			cell := g.board[y][x]
			if cell != Empty {
				drawBlock(g.screen, px, py, pieceColors[cell], false)
			}
		}
	}
}

func (g *Game) drawActivePiece(offsetX, offsetY int) {
	if g.gameOver {
		return
	}

	gy := g.ghostY()
	if gy != g.currY {
		for _, b := range tetrominoes[g.currentPiece][g.currRot] {
			bx := g.currX + b[0]
			by := gy + b[1]
			if by >= 0 && by < boardHeight && bx >= 0 && bx < boardWidth {
				drawBlock(g.screen, offsetX+bx*cellWidth, offsetY+by*cellHeight, pieceColors[g.currentPiece], true)
			}
		}
	}

	for _, b := range tetrominoes[g.currentPiece][g.currRot] {
		bx := g.currX + b[0]
		by := g.currY + b[1]
		if by >= 0 && by < boardHeight && bx >= 0 && bx < boardWidth {
			drawBlock(g.screen, offsetX+bx*cellWidth, offsetY+by*cellHeight, pieceColors[g.currentPiece], false)
		}
	}
}

func (g *Game) drawNextPreview(sideX, startY int, borderStyle tcell.Style) int {
	previewPixelW := previewCols*cellWidth + 2
	previewPixelH := previewRows*cellHeight + 2

	drawString(g.screen, sideX, startY, "NEXT", borderStyle)
	drawBox(g.screen, sideX, startY+1, previewPixelW, previewPixelH, borderStyle)

	innerX := sideX + 1
	innerY := startY + 2
	minX, maxX, minY, maxY := pieceBounds(g.nextPiece, 0)
	pieceW := maxX - minX + 1
	pieceH := maxY - minY + 1
	adjX := (previewCols - pieceW) / 2
	adjY := (previewRows - pieceH) / 2

	for _, b := range tetrominoes[g.nextPiece][0] {
		px := innerX + (b[0]-minX+adjX)*cellWidth
		py := innerY + (b[1]-minY+adjY)*cellHeight
		drawBlock(g.screen, px, py, pieceColors[g.nextPiece], false)
	}

	labelY := startY + previewPixelH + 2
	drawString(g.screen, sideX, labelY, pieceNames[g.nextPiece], tcell.StyleDefault.Foreground(pieceColors[g.nextPiece]))
	return labelY + 2
}

func (g *Game) draw() {
	s := g.screen
	s.Clear()

	style := tcell.StyleDefault
	borderStyle := style.Foreground(tcell.ColorWhite)

	termW, termH := s.Size()
	layoutW := layoutWidth()
	layoutH := layoutHeight()
	originX := max(0, (termW-layoutW)/2)
	originY := max(0, (termH-layoutH)/2)

	offsetX := originX
	offsetY := originY + titleAboveBoard

	title := "TETRIS"
	drawString(s, originX+(layoutW-len(title))/2, originY, title, style.Bold(true))

	g.drawBoard(offsetX, offsetY, borderStyle)
	g.drawActivePiece(offsetX, offsetY)

	sideX := offsetX + boardPixelWidth() + layoutGap

	drawString(s, sideX, offsetY, "SCORE", borderStyle)
	drawString(s, sideX, offsetY+2, fmt.Sprintf("%d", g.score), style.Bold(true))

	drawString(s, sideX, offsetY+5, "LINES", borderStyle)
	drawString(s, sideX, offsetY+7, fmt.Sprintf("%d", g.lines), style)

	drawString(s, sideX, offsetY+10, "LEVEL", borderStyle)
	drawString(s, sideX, offsetY+12, fmt.Sprintf("%d", g.level), style)

	ctrlY := g.drawNextPreview(sideX, offsetY+sidebarNextOffset, borderStyle)
	for i, line := range controlLines {
		drawString(s, sideX, ctrlY+i, line, style)
	}

	centerX := offsetX + boardPixelWidth()/2 - 3
	centerY := offsetY + boardPixelHeight()/2
	if g.paused {
		drawString(s, centerX, centerY, "PAUSED", style.Bold(true).Foreground(tcell.ColorYellow))
	}
	if g.gameOver {
		drawString(s, centerX-1, centerY-1, "GAME OVER", style.Bold(true).Foreground(tcell.ColorRed))
		msg := "Press R to restart or Q to quit"
		drawString(s, offsetX+(boardPixelWidth()-len(msg))/2, centerY+2, msg, style)
	}

	s.Show()
}

func (g *Game) handleInput(ev tcell.Event) bool {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape, tcell.KeyCtrlC:
			g.quit = true
			return true
		case tcell.KeyLeft:
			g.move(-1, 0)
		case tcell.KeyRight:
			g.move(1, 0)
		case tcell.KeyDown:
			g.softDrop()
		case tcell.KeyUp:
			g.rotate()
		case tcell.KeyRune:
			switch ev.Rune() {
			case 'q', 'Q':
				g.quit = true
				return true
			case 'p', 'P':
				if !g.gameOver {
					g.paused = !g.paused
				}
			case 'z', 'Z', 'x', 'X':
				g.rotate()
			case ' ':
				g.hardDrop()
			case 'r', 'R':
				if g.gameOver {
					g.reset()
				}
			}
		}
	case *tcell.EventResize:
		g.screen.Sync()
	}
	return false
}

func (g *Game) reset() {
	g.board = [boardHeight][boardWidth]int{}
	g.score = 0
	g.lines = 0
	g.level = 1
	g.dropSpeed = 500 * time.Millisecond
	g.gameOver = false
	g.paused = false
	g.nextPiece = g.randomPiece()
	g.spawnNewPiece()
	g.lastDrop = time.Now()
}

func run() error {
	initAudio()
	defer closeAudio()

	s, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	if err := s.Init(); err != nil {
		return err
	}
	defer s.Fini()

	s.SetStyle(tcell.StyleDefault.Background(tcell.ColorDefault))
	s.Clear()

	game := NewGame(s)

	// Ticker for game updates
	ticker := time.NewTicker(16 * time.Millisecond) // ~60 FPS for smooth input feel
	defer ticker.Stop()

	// Event channel
	eventChan := make(chan tcell.Event, 8)
	go func() {
		for {
			ev := s.PollEvent()
			if ev == nil {
				return
			}
			select {
			case eventChan <- ev:
			default:
			}
		}
	}()

	for !game.quit {
		select {
		case ev := <-eventChan:
			if game.handleInput(ev) {
				// quit requested
			}
		case <-ticker.C:
			if !game.paused && !game.gameOver {
				game.update()
			}
			game.draw()
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
