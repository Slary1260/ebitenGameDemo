package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	ModeTitle int = iota
	ModeGame
	ModeOver
)

var (
	titleArcadeFont font.Face
	arcadeFont      font.Face
)

type Game struct {
	mode    int
	input   *Input
	cfg     *Config
	ship    *Ship
	bullets map[*Bullet]struct{}
	aliens  map[*Alien]struct{}

	// 如果击杀所有外星人则游戏胜利，有3个外星人移出屏幕外或者碰撞到飞船则游戏失败
	failCount int
	overMsg   string
}

func NewGame() (*Game, error) {
	cfg := loadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)

	ship, err := NewShip(cfg.ScreenWidth, cfg.ScreenHeight)
	if err != nil {
		return nil, err
	}

	g := &Game{
		mode:    ModeTitle,
		input:   &Input{msg: "Hello, World!"},
		cfg:     cfg,
		ship:    ship,
		bullets: map[*Bullet]struct{}{},
		aliens:  map[*Alien]struct{}{},
	}

	g.init()

	return g, nil
}

func (g *Game) init() {
	g.createAliens()
	g.CreateFonts()
}

// 默认ebiten游戏是60帧，即每秒更新60次。
// 该方法主要用来更新游戏的逻辑状态，例如子弹位置更新。
// 注意到Update方法的返回值为error类型，当Update方法返回一个非空的error值时，游戏停止。
func (g *Game) Update() error {
	switch g.mode {
	case ModeTitle:
		// 按下空格或鼠标左键游戏开始
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.mode = ModeGame
		}
	case ModeGame:
		g.input.Update(g)

		// 子弹移动
		for bullet := range g.bullets {
			bullet.Update()
		}

		// 清除屏幕外的子弹
		for bullet := range g.bullets {
			if bullet.OutOfScreen() {
				delete(g.bullets, bullet)
			}
		}

		// 外星人移动
		for alien := range g.aliens {
			alien.Update()
		}

		// 检查碰撞
		g.CheckCollision()

		for alien := range g.aliens {
			if alien.OutOfScreen(g.cfg) {
				g.failCount++
				delete(g.aliens, alien)
				continue
			}

			if CheckCollision(alien, g.ship) {
				g.failCount++
				delete(g.aliens, alien)
				continue
			}
		}

		if g.failCount >= 3 {
			g.overMsg = "Game Over!"
		} else if len(g.aliens) == 0 {
			g.overMsg = "You Win!"
		}

		if len(g.overMsg) > 0 {
			g.mode = ModeOver
			g.aliens = make(map[*Alien]struct{})
			g.bullets = make(map[*Bullet]struct{})
		}

	case ModeOver:
		// 按下空格或鼠标左键即重新开始游戏
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.init()
			g.mode = ModeTitle
		}
	}

	return nil
}

// 每帧（frame）调用。
// 帧是渲染使用的一个时间单位，依赖显示器的刷新率。如果显示器的刷新率为60Hz，Draw将会每秒被调用60次。
// 由于调用Draw方法前，screen会被重置，故DebugPrint每次都需要调用。
func (g *Game) Draw(screen *ebiten.Image) {
	// 更换背景颜色
	screen.Fill(g.cfg.BgColor)

	// 显示文本
	// ebitenutil.DebugPrint(screen, g.input.msg)

	var titleTexts []string
	var texts []string
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"ALIEN INVASION"}
		texts = []string{"", "", "", "", "", "", "", "PRESS SPACE KEY", "", "OR LEFT MOUSE"}
	case ModeGame:
		// 显示小飞机
		g.ship.Draw(screen, g.cfg)

		for bullet := range g.bullets {
			bullet.Draw(screen)
		}

		for alien := range g.aliens {
			alien.Draw(screen)
		}

	case ModeOver:
		texts = []string{"", "GAME OVER!"}
	}

	for i, l := range titleTexts {
		x := (g.cfg.ScreenWidth - len(l)*g.cfg.TitleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*g.cfg.TitleFontSize, color.White)
	}

	for i, l := range texts {
		x := (g.cfg.ScreenWidth - len(l)*g.cfg.FontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*g.cfg.FontSize, color.White)
	}

}

// 该方法接收游戏窗口的尺寸作为参数，返回游戏的逻辑屏幕大小。
// 例子中游戏窗口大小为(640, 480)，Layout返回的逻辑大小为(320, 240)，所以显示会放大1倍。
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth, g.cfg.ScreenHeight
}

func (g *Game) addBullet(bullet *Bullet) {
	g.bullets[bullet] = struct{}{}
}

// 创建一组外星人
func (g *Game) createAliens() error {
	alien, err := NewAlien(g.cfg)
	if err != nil {
		return err
	}

	// 左右各留一个外星人宽度的空间
	availableSpaceX := g.cfg.ScreenWidth - 2*alien.width
	// 两个外星人之间留一个外星人宽度的空间
	numAliens := availableSpaceX / (2 * alien.width)

	for row := 0; row < 2; row++ {
		for i := 0; i < numAliens; i++ {
			alien, err = NewAlien(g.cfg)
			if err != nil {
				return err
			}

			alien.x = float64(alien.width + 2*alien.width*i)
			alien.y = float64(alien.height*row) * 1.5
			g.addAlien(alien)
		}
	}

	return nil
}

func (g *Game) addAlien(alien *Alien) {
	g.aliens[alien] = struct{}{}
}

func (g *Game) CheckCollision() {
	for alien := range g.aliens {
		for bullet := range g.bullets {
			if CheckCollision(bullet, alien) {
				delete(g.aliens, alien)
				delete(g.bullets, bullet)
			}
		}
	}
}

func (g *Game) CreateFonts() error {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		return err
	}

	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.TitleFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}

	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.FontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}

	return nil
}
