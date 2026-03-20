package main

import (
		rl "github.com/gen2brain/raylib-go/raylib"
)

type xboss struct {
	img     rl.Rectangle
	rec     rl.Rectangle
	crec    rl.Rectangle
	yt      float32
	xl      float32
	vel     float32
	velX    float32
	velY    float32
	hppause int32
	timer   int32
	hp      int
	hpmax   int
	direc   int
	atkType int
	cnt     rl.Vector2
	off     bool
}

type xmod struct {
	axe            bool
	santa          bool
	snowon         bool
	fireball       bool
	vine           bool
	key            bool
	apple          bool
	planty         bool
	medikit        bool
	wallet         bool
	exitmap        bool
	firetrail      bool
	hppotion       bool
	invisible      bool
	orbital        bool
	chainlightning bool
	recharge       bool
	anchor         bool
	umbrella       bool
	socks          bool
	flood          bool
	peace          bool
	alien          bool
	airstrike      bool
	fireworks      bool
	carrot         bool
	axeN           int
	fireballN      int
	bounceN        int
	keyN           int
	appleN         int
	firetrailN     int
	hppotionN      int
	coffeeN        int
	atkrangeN      int
	atkdmgN        int
	orbitalN       int
	hpringN        int
	armorN         int
	cherryN        int
	cakeN          int
	axeT           int32
	axeT2          int32
	santaT         int32
}

type xmax struct {
	axe      int
	fireball int
	bounce   int
	key      int
	apple    int
	firetrail int
	hppotion int
	coffee   int
	atkrange int
	atkdmg   int
	orbital  int
	hpring   int
	armor    int
	cherry   int
	cake     int
}

type xenemy struct {
	img      rl.Rectangle
	rec      rl.Rectangle
	imgl     rl.Rectangle
	imgr     rl.Rectangle
	crec     rl.Rectangle
	arec     rl.Rectangle
	cnt      rl.Vector2
	ori      rl.Vector2
	ro       float32
	vel      float32
	velX     float32
	velY     float32
	xImg     float32
	xImg2    float32
	fade     float32
	frameNum int
	direc    int
	hp       int
	hpmax    int
	spawnN   int
	fly      bool
	anim     bool
	off      bool
	name     string
	col      rl.Color
	hppause  int32
	T1       int32
}

type xplayer struct {
	cnt             rl.Vector2
	ori             rl.Vector2
	orbital1        rl.Vector2
	orbital2        rl.Vector2
	img             rl.Rectangle
	rec             rl.Rectangle
	crec            rl.Rectangle
	arec            rl.Rectangle
	atkrec          rl.Rectangle
	orbimg1         rl.Rectangle
	orbimg2         rl.Rectangle
	orbrec1         rl.Rectangle
	orbrec2         rl.Rectangle
	atkTimer        int32
	hppause         int32
	slideT          int32
	poisonT         int32
	poisonCollisT   int32
	hppotionT       int32
	peaceT          int32
	waterT          int32
	move            bool
	atk             bool
	slide           bool
	escape          bool
	revived         bool
	poison          bool
	armorHit        bool
	underWater      bool
	siz             float32
	ro              float32
	vel             float32
	imgAtkX         float32
	imgWalkX        float32
	sizImg          float32
	angle           float32
	direc           int
	framesAtk       int
	framesWalk      int
	hp              int
	hpmax           int
	coins           int
	atkDMG          int
	slideDIR        int
	poisonCount     int
	armor           int
	armorMax        int
	rechargeN       int
}

type xblok struct {
	name        string
	desc        string
	img         rl.Rectangle
	rec         rl.Rectangle
	crec        rl.Rectangle
	crec2       rl.Rectangle
	drec        rl.Rectangle
	color       rl.Color
	fade        float32
	cnt         rl.Vector2
	ori         rl.Vector2
	velX        float32
	velY        float32
	vel         float32
	ro          float32
	v2s         []rl.Vector2
	timer       int32
	txtT        int32
	movType     int
	numof       int
	slideDIR    int
	numType     int
	numCoins    int
	shopprice   int
	bump        bool
	onoff       bool
	solid       bool
	onoffswitch bool
	fadeswitch  bool
	shopoff     bool
}

type xcompanion struct {
	img   rl.Rectangle
	imgl  rl.Rectangle
	imgr  rl.Rectangle
	rec   rl.Rectangle
	crec  rl.Rectangle
	hp    int
	hpmax int
	frames int
	vel   float32
	velx  float32
	vely  float32
	cnt   rl.Vector2
	timer int32
}

type xroom struct {
	floor       []xblok
	walls       []xblok
	movBloks    []xblok
	etc         []xblok
	innerBloks  []xblok
	spikes      []xblok
	doorSides   []int
	nextRooms   []int
	visited     bool
	exit        bool
	floorT      rl.Rectangle
	wallT       rl.Rectangle
	enemies     []xenemy
	doorExitRecs []rl.Rectangle
}

type ximg struct {
	img  rl.Rectangle
	rec  rl.Rectangle
	ro   float32
	fade float32
	cnt  rl.Vector2
	ori  rl.Vector2
	col  rl.Color
	off  bool
}

type xkills struct {
	bunnies   int
	bats      int
	mushrooms int
	rocks     int
	slimes    int
	spikehogs int
	ghosts    int
}

type xtxt struct {
	txt   string
	txt2  string
	x     int32
	y     int32
	fade  float32
	col   rl.Color
	onoff bool
}

type xproj struct {
	cnt     rl.Vector2
	drec    rl.Rectangle
	rec     rl.Rectangle
	crec    rl.Rectangle
	img     rl.Rectangle
	ori     rl.Vector2
	onoff   bool
	col     rl.Color
	dmg     int
	bounceN int
	name    string
	ro      float32
	vel     float32
	velx    float32
	vely    float32
	fade    float32
}

type xfx struct {
	timer int32
	onoff bool
	name  string
	cnt   rl.Vector2
	rec   rl.Rectangle
	img   rl.Rectangle
	col   rl.Color
	fade  float32
	recs  []xrec
}

type xrec struct {
	rec   rl.Rectangle
	col   rl.Color
	fade  float32
	velX  float32
	velY  float32
}

type xanim struct {
	xl     float32
	yt     float32
	frames float32
	W      float32
	recTL  rl.Rectangle
}

type GameState struct {
    Core       CoreState
    Render     RenderState
    Audio      AudioState
    Input      InputState
    Player     PlayerState
    Enemies    EnemyState
    Level      LevelState
    Shop       ShopState
    UI         UIState
    FX         FXState
    Companions CompanionState
    Timing     TimingState
    Mario      MarioState
}

type CoreState struct {
    Fps       int32
    Frames    int
    ScrW      int
    ScrH      int
    ScrW32    int32
    ScrH32    int32
    ScrWF32   float32
    ScrHF32   float32
    Cnt       rl.Vector2
    Ori       rl.Vector2
    MouseV2   rl.Vector2
    Mousev2cam rl.Vector2
    Debug     bool
    Pause     bool
}

type RenderState struct {
    Imgs         rl.Texture2D
    Shader       rl.Shader
    Shader2      rl.Shader
    Shader3      rl.Shader
    RenderTarget rl.RenderTexture2D
    Cam2         rl.Camera2D
    ShaderOn     bool
    Shader2On    bool
    Shader3On    bool
    // sprite sheet slices
    Walltiles  []rl.Rectangle
    Floortiles []rl.Rectangle
    Bats       []rl.Rectangle
    Knight     []rl.Rectangle
    Etc        []rl.Rectangle
    Shrines    []rl.Rectangle
    Plants     []rl.Rectangle
    Skulls     []rl.Rectangle
    Candles    []rl.Rectangle
    Signs      []rl.Rectangle
    Splats     []rl.Rectangle
    Statues    []rl.Rectangle
    Mushrooms  []rl.Rectangle
    Alien      []rl.Rectangle
    Patterns   []rl.Rectangle
    Gems       []rl.Rectangle
    Coin       rl.Rectangle
    // animations
    Rabbit1         xanim
    FireballPlayer  xanim
    Burn            xanim
    Star            xanim
    Wateranim       xanim
    PlantBull        xanim
    Spikes          xanim
    Spring          xanim
    Posiongas       xanim
    MushBull        xanim
    Blades          xanim
    Spear           xanim
    Firetrailanim   xanim
    Orbitalanim     xanim
    Floodanim       xanim
    FishR           xanim
    FishL           xanim
    Airstrikeanim   xanim
    Boss1anim       xanim
    Boss2anim       xanim
}
// Done
type AudioState struct {
    Music      rl.Music
    BackMusic  []rl.Music
    Sfx        []rl.Sound
    MusicOn    bool
    Volume     float32
    BgMusicNum int
}
// Done
type InputState struct {
    UseController        bool
    IsController         bool
    ControllerOn         bool
    ControllerDisconnect bool
    ControllerWasOn      bool
    KeypressT            int32
}

type PlayerState struct {
    Pl                     xplayer
    Mods                   xmod
    Max                    xmax
    Kills                  xkills
    PlProj                 []xproj
    Inven                  []xblok
    InvenOn                bool
    HpHitY                 float32
    ReviveY                float32
    WaterY                 float32
    HpHitF                 float32
    ReviveF                float32
    WaterF                 float32
    PlVineRec              rl.Rectangle
    DiedRec                rl.Rectangle
    DiedIMG                rl.Rectangle
    TeleportRoomNum        int
    TeleportRadius         []float32
    StartdmgT              int32
    Escaped                bool
    EscapeRoomFound        bool
    TeleportOn             bool
    PlatkrecOn             bool
    ChainLightingSwingOnOff bool
    Died                   bool
}

type EnemyState struct {
    EnProj     []xproj
    EnSpikes   xenemy
    EnGhost    xenemy
    EnSlime    xenemy
    EnRock     xenemy
    EnMushroom xenemy
}

type LevelState struct {
    Level             []xroom
    LevRec            rl.Rectangle
    LevRecInner       rl.Rectangle
    WallT             rl.Rectangle
    LevW              float32
    BorderWallBlokSiz float32
    LevX              float32
    LevY              float32
    RoomNum           int
    LevBorderBlokNum  int
    LevMap            []rl.Rectangle
    Levelnum          int
    ExitRoomNum       int
    ShopRoomNum       int
    Exited            bool
    NextLevelScreen   bool
    Secs              int
    Mins              int
    MinsEND           int
    SecsEND           int
    RoomChangedTimer  int32
    AnchorT           int32
    RunT              int32
    DiedscrT          int32
    NextlevelT        int32
    LevMapOn          bool
    RoomChanged       bool
    Night             bool
    Flipcam           bool
    Exiton            bool
    ExitLR            bool
    Endgame           bool
    Hardcore          bool
    // end level
    Bosses       []xboss
    Bossnum      int
    EndgameT     int32
    EndPauseT    int32
    EndgopherRec rl.Rectangle
}

// Done
type ShopState struct {
    ShopOn    bool
    ShopExitY float32
    ShopItems []xblok
    ShopNum   int
    ShopExitT int32
}

type UIState struct {
    OptionNum            int
    TxtSize              int32
    OptionT              int32
    OptionsOn            bool
    HpBarsOn             bool
    ArtifactsOn          bool
    ScanLinesOn          bool
    CreditsOn            bool
    HelpOn               bool
    Invincible           bool
    Resettimes           bool
    RestartOn            bool
    OptionsChange        bool
    StartScreen          bool
    Intro                bool
    IntroCount           bool
    IntroT1              int32
    IntroT2              int32
    IntroT3              int32
    IntroF1              float32
    IntroF2              float32
    IntroF3              float32
    FadeBlinkOn          bool
    FadeBlinkOn2         bool
    FadeBlink            float32
    FadeBlink2           float32
    TxtSoldList          []xtxt
    GameTxt              []xtxt
}

type FXState struct {
    Fx              []xfx
    Snow            []ximg
    ScanlineV2      []rl.Vector2
    ChainV2         []rl.Vector2
    ChainLightOn    bool
    ChainLightTimer int32
    Rain            []rl.Rectangle
    FloodRec        rl.Rectangle
    FloodImg        rl.Rectangle
    Fish1           rl.Rectangle
    Fish2           rl.Rectangle
    FishV2          rl.Vector2
    Fish2V2         rl.Vector2
    FishSiz         float32
    FishSiz2        float32
    FishLR          bool
    Fish2LR         bool
    FishRec         rl.Rectangle
    FishRec2        rl.Rectangle
    WaterLR         bool
    WaterUP         bool
    AirstrikeT      int32
    AirstrikebombT  int32
    AirstrikeDir    int
    AirstrikeOn     bool
    AirstrikeV2     []rl.Vector2
    FireworksCnt    rl.Vector2
}

// Done
type CompanionState struct {
    MrPlanty  xcompanion
    MrAlien   xcompanion
    MrCarrot  xcompanion
}

// Done
type TimingState struct {
    Times      []int
    BestTime   bool
    TimesOn    bool
    BestTimesT int32
}

// Done
type MarioState struct {
    MarioOn        bool
    MarioJump      bool
    MarioT         int32
    MarioJumpT     int32
    MarioRecs      []rl.Rectangle
    MarioCoins     []rl.Rectangle
    MarioPL        rl.Rectangle
    MarioScreenRec rl.Rectangle
    PatternRec     rl.Rectangle
    MarioImg       rl.Rectangle
    MarioV2L       rl.Vector2
    MarioV2R       rl.Vector2
    MarioCols      []rl.Color
    MarioCoinOnOff []bool
}
