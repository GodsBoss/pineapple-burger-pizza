package game

import (
	"github.com/GodsBoss/gggg/v2/pkg/dom"
	"github.com/GodsBoss/gggg/v2/pkg/event/keyboard"
	"github.com/GodsBoss/gggg/v2/pkg/event/mouse"
	"github.com/GodsBoss/gggg/v2/pkg/event/tick"
	"github.com/GodsBoss/gggg/v2/pkg/game"
)

func New(img *dom.Image) *Game {
	tmpl := game.Template[*data]{
		CreateData: func() *data {
			return &data{}
		},
	}

	spriteMap, keys := createSpriteMap(img)

	// Define states.
	title := tmpl.AddState()
	playing := tmpl.AddState()
	gameOver := tmpl.AddState()

	// Configure states.

	title.
		SetInitHandler(initTitle).
		SetKeyboardHandler(createReceiveTitle(playing.ID())).
		SetTickHandler(createReceiveTickEventTitle())

	playing.
		SetInitHandler(initPlaying).
		SetKeyboardHandler(createReceiveKeyEventPlaying()).
		SetMouseHandler(createReceiveMouseEventPlaying()).
		SetTickHandler(createReceiveTickEventPlaying(gameOver.ID()))

	gameOver.
		SetInitHandler(initGameOver).
		SetKeyboardHandler(createReceiveKeyEventGameOver(title.ID())).
		SetTickHandler(createReceiveTickEventGameOver())

	instance, _ := tmpl.NewInstance()

	tm := newTextManager(spriteMap, charSprites)

	r := &renderer{
		scaler: createScaler(),
	}

	r.
		AddStateRenderer(titleState, stateRendererFunc(renderTitle(spriteMap, keys, tm))).
		AddStateRenderer(playingState, stateRendererFunc(renderPlaying(spriteMap, keys, tm))).
		AddStateRenderer(gameOverState, stateRendererFunc(renderGameOver(spriteMap, keys, tm)))

	return &Game{
		data:     instance,
		renderer: r,
	}
}

type Game struct {
	data     *game.Instance[*data]
	renderer *renderer
}

func (g *Game) TicksPerSecond() int {
	return 20
}

func (g *Game) SetOutput(ctx2d *dom.Context2D) {
	g.renderer.SetOutput(ctx2d)
}

func (g *Game) Render() {
	g.renderer.Render(g.data.Data())
}

func (g *Game) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, scaleX, scaleY float64) {
	return g.renderer.Scale(availableWidth, availableHeight)
}

func (g *Game) Tick(ms int) {
	event := tick.Event{
		MsSinceLastTick: ms,
	}
	g.data.ReceiveTickEvent(event)
}

func (g *Game) ReceiveKeyEvent(event keyboard.Event) {
	g.data.ReceiveKeyEvent(event)
}

func (g *Game) ReceiveMouseEvent(event mouse.Event) {
	g.data.ReceiveMouseEvent(event)
}
