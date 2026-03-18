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

