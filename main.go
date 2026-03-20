package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// 10.2023-001 - BITTY KNIGHT

// NOTE ALL THE MARK: NOTES ARE FOR VS CODE MARK JUMPING AND CAN BE IGNORED


// Constants - these stay as package-level vars since they're computed constants
var (
    bsU   = float32(16)
    bsU2  = bsU * 2
    bsU3  = bsU * 3
    bsU4  = bsU * 4
    bsU5  = bsU * 5
    bsU6  = bsU * 6
    bsU7  = bsU * 7
    bsU8  = bsU * 8
    bsU9  = bsU * 9
    bsU10 = bsU * 10
    bsU11 = bsU * 11
    bsU12 = bsU * 12

    bsUi   = 16
    bsU2i  = bsUi * 2
    bsU3i  = bsUi * 3
    bsU4i  = bsUi * 4
    bsU5i  = bsUi * 5
    bsU6i  = bsUi * 6
    bsU7i  = bsUi * 7
    bsU8i  = bsUi * 8
    bsU9i  = bsUi * 9
    bsU10i = bsUi * 10

    bsUi32   = int32(16)
    bsU2i32  = bsUi32 * 2
    bsU3i32  = bsUi32 * 3
    bsU4i32  = bsUi32 * 4
    bsU5i32  = bsUi32 * 5
    bsU6i32  = bsUi32 * 6
    bsU7i32  = bsUi32 * 7
    bsU8i32  = bsUi32 * 8
    bsU9i32  = bsUi32 * 9
    bsU10i32 = bsUi32 * 10

    txU   = int32(10)
    txU2  = txU * 2
    txU3  = txU * 3
    txU4  = txU * 4
    txU5  = txU * 5
    txU6  = txU * 6
    txU7  = txU * 7
    txU8  = txU * 8
    txU9  = txU * 9
    txU10 = txU * 10
)

var gs = &GameState{
	Core: CoreState{
		Fps: 60,
		Ori: rl.NewVector2(0, 0),
	},
	Audio: AudioState{
		Volume: 0.2,
	},
	UI: UIState{
		TxtSize:    txU2,
		FadeBlink:  0.5,
		FadeBlink2: 0.5,
	},
	Input: InputState{
		ControllerOn: true,
	},
	Render: RenderState{
		Coin: rl.NewRectangle(1120, 250, 16, 16),
	},
	Level: LevelState{
		LevW: 720,
		BorderWallBlokSiz: bsU,
		Levelnum: 1,
	},
	Player: PlayerState{
		HpHitF:  1,
		ReviveF: 1,
		WaterF:  1,
	},
}

// MARK: DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW
func drawcam() { //MARK:DRAW CAM

	//INVEN OPTIONS SHOP ON
	if gs.UI.OptionsOn {
		drawOptions()
	} else if gs.Timing.TimesOn {
		drawTimes()
	} else if gs.Shop.ShopOn {
		drawShop()
	} else if gs.Mario.MarioOn {
		drawUpMario()
	} else if gs.Level.Endgame {
		drawEndGame()
	} else if gs.Level.NextLevelScreen {
		drawnextlevelscreen()
	} else if gs.Player.Died {
		drawDied()
	}

	//DRAW GAME LEVEL
	if !gs.Core.Pause {

		//FLOORS
		dFloors := gs.Level.Level[gs.Level.RoomNum].floor
		for a := 0; a < len(dFloors); a++ {
			drawBlok(dFloors[a], false, false, 0)
		}
		//rl.DrawRectangleRec(gs.Level.LevRec, rl.Fade(rl.Black, 0.5))

		//SPIKES
		if len(gs.Level.Level[gs.Level.RoomNum].spikes) > 0 {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].spikes); a++ {
				shadowRec := gs.Level.Level[gs.Level.RoomNum].spikes[a].rec
				shadowRec.X -= 5
				shadowRec.Y += 5
				rl.DrawTexturePro(gs.Render.Imgs, gs.Level.Level[gs.Level.RoomNum].spikes[a].img, shadowRec, gs.Core.Ori, 0, rl.Fade(rl.Black, 0.8))
				col := ranRed()
				rl.DrawTexturePro(gs.Render.Imgs, gs.Level.Level[gs.Level.RoomNum].spikes[a].img, gs.Level.Level[gs.Level.RoomNum].spikes[a].rec, gs.Core.Ori, 0, col)

				if gs.Core.Frames%3 == 0 {
					gs.Level.Level[gs.Level.RoomNum].spikes[a].img.X += 32
					if gs.Level.Level[gs.Level.RoomNum].spikes[a].img.X >= gs.Render.Spikes.xl+gs.Render.Spikes.frames*32 {
						gs.Level.Level[gs.Level.RoomNum].spikes[a].img.X = gs.Render.Spikes.xl
					}
				}

				//CHECK PLAYER SPIKES COLLISION
				if rl.CheckCollisionRecs(gs.Player.Pl.crec, gs.Level.Level[gs.Level.RoomNum].spikes[a].rec) {
					hitPL(0, 2)
				}
			}
		}

		//WALLS
		dWalls := gs.Level.Level[gs.Level.RoomNum].walls
		for a := 0; a < len(dWalls); a++ {
			drawBlok(dWalls[a], false, false, 0)
			//WALL BLOCKNUMS
			if gs.Core.Debug {
				rl.DrawText(fmt.Sprint(a), dWalls[a].rec.ToInt32().X+4, dWalls[a].rec.ToInt32().Y+4, txU, rl.White)
			}
		}

		//INNER BLOKS
		if len(gs.Level.Level[gs.Level.RoomNum].innerBloks) > 0 {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].innerBloks); a++ {
				drawBlok(gs.Level.Level[gs.Level.RoomNum].innerBloks[a], false, true, 4)
			}
		}

		//ENEMY PROJECTILES
		if len(gs.Enemies.EnProj) > 0 {
			drawUpEnProj()
		}

		//FX
		if len(gs.FX.Fx) > 0 {
			drawupfx()
		}

		//ETC
		if len(gs.Level.Level[gs.Level.RoomNum].etc) > 0 {
			drawUpEtc()
		}

		//PLAYER PROJECTILES
		if len(gs.Player.PlProj) > 0 {
			drawUpPlayerProj()
		}

		//CHAIN LIGHTNING
		if gs.FX.ChainLightOn {
			drawChainLight()
		}

		//ENEMIES
		if len(gs.Level.Level[gs.Level.RoomNum].enemies) > 0 {
			drawUpEnemies()
		}
		//BOSS
		if gs.Level.Levelnum == 6 {
			drawUpBoss()
		}

		//MOVE BLOKS
		dMovBloks := gs.Level.Level[gs.Level.RoomNum].movBloks
		for a := 0; a < len(dMovBloks); a++ {
			drawBlok(dMovBloks[a], true, false, 0)
		}

		//PLAYER
		drawPlayer()

		//GAME TEXT
		if len(gs.UI.GameTxt) > 0 {
			clear := false
			for a := 0; a < len(gs.UI.GameTxt); a++ {
				if gs.UI.GameTxt[a].onoff {
					rl.DrawText(gs.UI.GameTxt[a].txt, gs.UI.GameTxt[a].x, gs.UI.GameTxt[a].y, 20, rl.Fade(gs.UI.GameTxt[a].col, gs.UI.GameTxt[a].fade))

					gs.UI.GameTxt[a].fade -= 0.02
					gs.UI.GameTxt[a].y--
					if gs.UI.GameTxt[a].fade <= 0 {
						gs.UI.GameTxt[a].onoff = false
					}
				} else {
					clear = true
				}
			}
			if clear {
				for a := 0; a < len(gs.UI.GameTxt); a++ {
					if !gs.UI.GameTxt[a].onoff {
						gs.UI.GameTxt = remTxt(gs.UI.GameTxt, a)
					}
				}
			}
		}

		//AIR STRIKE
		if gs.FX.AirstrikeOn {
			drawUpAirStrike()
		}

		//SNOW
		if len(gs.FX.Snow) > 0 {
			clear := false
			for a := 0; a < len(gs.FX.Snow); a++ {
				if !gs.FX.Snow[a].off {
					rl.DrawTexturePro(gs.Render.Imgs, gs.FX.Snow[a].img, gs.FX.Snow[a].rec, gs.FX.Snow[a].ori, gs.FX.Snow[a].ro, rl.Fade(gs.FX.Snow[a].col, gs.FX.Snow[a].fade))

					gs.FX.Snow[a].ro += 2
					gs.FX.Snow[a].rec.Y += 2

					if gs.FX.Snow[a].rec.Y > gs.Core.ScrHF32 {
						gs.FX.Snow[a].off = true
						clear = true
					}
				}
			}
			if clear {
				for a := 0; a < len(gs.FX.Snow); a++ {
					if gs.FX.Snow[a].off {
						gs.FX.Snow = remImg(gs.FX.Snow, a)
					}
				}
			}
		}

		//INVENTORY
		if len(gs.Player.Inven) > 0 {
			drawInven()
		}
		//PLAYER INFO
		drawPlayerInfo()

		if gs.Core.Debug {
			//NEXT ROOM DOOR NUMS
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].doorSides); a++ {
				if gs.Level.Level[gs.Level.RoomNum].doorSides[a] == 1 {
					rl.DrawText("up  "+fmt.Sprint(gs.Level.Level[gs.Level.RoomNum].nextRooms[a]), gs.Level.LevRec.ToInt32().X+gs.Level.LevRec.ToInt32().Width/2, gs.Level.LevRec.ToInt32().Y+bsUi32, txU2, rl.White)
				}
				if gs.Level.Level[gs.Level.RoomNum].doorSides[a] == 2 {
					rl.DrawText("right"+fmt.Sprint(gs.Level.Level[gs.Level.RoomNum].nextRooms[a]), gs.Level.LevRec.ToInt32().X+gs.Level.LevRec.ToInt32().Width-bsU2i32, gs.Level.LevRec.ToInt32().Y+gs.Level.LevRec.ToInt32().Width/2, txU2, rl.White)
				}
				if gs.Level.Level[gs.Level.RoomNum].doorSides[a] == 3 {
					rl.DrawText("down  "+fmt.Sprint(gs.Level.Level[gs.Level.RoomNum].nextRooms[a]), gs.Level.LevRec.ToInt32().X+gs.Level.LevRec.ToInt32().Width/2, gs.Level.LevRec.ToInt32().Y+gs.Level.LevRec.ToInt32().Width-bsU2i32, txU2, rl.White)
				}
				if gs.Level.Level[gs.Level.RoomNum].doorSides[a] == 4 {
					rl.DrawText("left  "+fmt.Sprint(gs.Level.Level[gs.Level.RoomNum].nextRooms[a]), gs.Level.LevRec.ToInt32().X+bsU2i32, gs.Level.LevRec.ToInt32().Y+gs.Level.LevRec.ToInt32().Width/2, txU2, rl.White)
				}
			}
			//DOOR EXIT RECS

			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].doorExitRecs); a++ {
				rl.DrawRectangleLinesEx(gs.Level.Level[gs.Level.RoomNum].doorExitRecs[a], 0.5, rl.Magenta)
			}

			//LEVEL BORDER RECS
			rl.DrawRectangleLinesEx(gs.Level.LevRec, 2, rl.Green)
			rl.DrawRectangleLinesEx(gs.Level.LevRecInner, 2, rl.Magenta)
		}
	}

	//ARTIFACTS
	if gs.UI.ArtifactsOn {
		num := 100

		for {
			x := (gs.Level.LevX - bsU6) + rF32(0, gs.Level.LevW+bsU12)
			y := gs.Level.LevY + rF32(0, gs.Level.LevW)
			siz := rF32(1, 3)
			rec := rl.NewRectangle(x, y, siz, siz)
			rl.DrawRectangleRec(rec, rl.Black)

			num--
			if num == 0 {
				break
			}
		}
	}

}
func drawnextlevelscreen() { //MARK:DRAW NEXT LEVEL SCREEN

	txt := "prepare for level " + fmt.Sprint(gs.Level.Levelnum)
	txtlen := rl.MeasureText(txt, 40)
	x := int32(gs.Core.Cnt.X) - txtlen/2
	y := int32(gs.Core.Cnt.Y - 20)

	rl.DrawText(txt, x, y, 40, rl.White)

	txt = "press space or button to continue"
	txtlen = rl.MeasureText(txt, 20)
	x = int32(gs.Core.Cnt.X) - txtlen/2
	y = int32(gs.Core.Cnt.Y + 30)

	rl.DrawText(txt, x, y, 20, rl.White)

	if gs.Level.NextlevelT > 0 {
		gs.Level.NextlevelT--
	} else {
		gs.Player.StartdmgT = gs.Core.Fps * 5
		if rl.IsKeyPressed(rl.KeySpace) {
			gs.Level.NextLevelScreen = false
			gs.Core.Pause = false
		}
		if gs.Input.UseController {
			if rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
				gs.Level.NextLevelScreen = false
				gs.Core.Pause = false
			}
		}
	}

}
func drawEndGame() { //MARK:DRAW END GAME
	gs.Core.Pause = true
	rl.DrawRectangle(0, 0, gs.Core.ScrW32, gs.Core.ScrH32, rl.Black)
	gs.Render.ShaderOn = false
	if gs.Level.EndPauseT > 0 {
		gs.Level.EndPauseT--
	}
	if gs.Level.EndgameT > 0 {
		gs.Level.EndgameT--
		if gs.Level.EndgopherRec.Y > gs.Level.LevRec.Y+gs.Level.LevRec.Height-gs.Level.EndgopherRec.Height {
			gs.Level.EndgopherRec.Y -= 2
		}
	}

	rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[55], gs.Level.EndgopherRec, rl.Vector2Zero(), 0, rl.White)
	rl.PlaySound(gs.Audio.Sfx[19])
	txt := "you"
	txtlen := rl.MeasureText(txt, txU8)
	rl.DrawText(txt, int32(gs.Core.Cnt.X)-txtlen/2-3, int32(gs.Core.Cnt.Y)-txU8+3, txU8, rl.Black)
	rl.DrawText(txt, int32(gs.Core.Cnt.X)-txtlen/2, int32(gs.Core.Cnt.Y)-txU8, txU8, rl.White)
	txt = "win"
	txtlen = rl.MeasureText(txt, txU8)
	rl.DrawText(txt, int32(gs.Core.Cnt.X)-txtlen/2-3, int32(gs.Core.Cnt.Y)+txU+3, txU8, rl.Black)
	rl.DrawText(txt, int32(gs.Core.Cnt.X)-txtlen/2, int32(gs.Core.Cnt.Y)+txU, txU8, rl.White)

	if gs.Input.KeypressT == 0 {
		if gs.Level.EndPauseT == 0 {
			if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) || gs.Level.EndgameT == 0 {
				gs.Level.Endgame = false
				gs.Level.EndgameT = 0
				gs.Level.EndgopherRec = rl.NewRectangle(gs.Core.Cnt.X-bsU4, gs.Level.LevRec.Y+gs.Level.LevRec.Height, bsU8, bsU8)
				restartgame()
			}
		}
	}
}
func drawShop() { //MARK:DRAW SHOP

	rl.DrawRectangle(0, 0, gs.Core.ScrW32, gs.Core.ScrH32, rl.Black)

	if rl.IsKeyPressed(rl.KeyA) || rl.GetGamepadAxisMovement(0, 0) < 0 || rl.IsGamepadButtonDown(0, 4) {
		if gs.UI.OptionT == 0 {
			gs.UI.OptionT = gs.Core.Fps / 5
			gs.Shop.ShopNum--
			if gs.Shop.ShopNum < 0 {
				gs.Shop.ShopNum = 4
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyD) || rl.GetGamepadAxisMovement(0, 0) > 0 || rl.IsGamepadButtonDown(0, 2) {
		if gs.UI.OptionT == 0 {
			gs.UI.OptionT = gs.Core.Fps / 5
			gs.Shop.ShopNum++
			if gs.Shop.ShopNum > 4 {
				gs.Shop.ShopNum = 0
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyS) || rl.GetGamepadAxisMovement(0, 1) > 0 || rl.IsGamepadButtonDown(0, 3) {
		if gs.UI.OptionT == 0 {
			gs.UI.OptionT = gs.Core.Fps / 5
			if gs.Shop.ShopNum == 0 {
				gs.Shop.ShopNum = 2
			} else if gs.Shop.ShopNum == 2 {
				gs.Shop.ShopNum = 4
			} else if gs.Shop.ShopNum == 4 {
				gs.Shop.ShopNum = 0
			}
			if gs.Shop.ShopNum == 1 {
				gs.Shop.ShopNum = 3
			} else if gs.Shop.ShopNum == 3 {
				gs.Shop.ShopNum = 4
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyW) || rl.GetGamepadAxisMovement(0, 1) < 0 || rl.IsGamepadButtonDown(0, 1) {
		if gs.UI.OptionT == 0 {
			gs.UI.OptionT = gs.Core.Fps / 5
			if gs.Shop.ShopNum == 0 {
				gs.Shop.ShopNum = 4
			} else if gs.Shop.ShopNum == 2 {
				gs.Shop.ShopNum = 0
			}
			if gs.Shop.ShopNum == 1 {
				gs.Shop.ShopNum = 4
			} else if gs.Shop.ShopNum == 3 {
				gs.Shop.ShopNum = 1
			}
			if gs.Shop.ShopNum == 4 {
				gs.Shop.ShopNum = 0
			}
		}
	}
	if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {

		if gs.Shop.ShopNum == 4 {
			gs.Shop.ShopExitT = gs.Core.Fps * 2
			gs.Shop.ShopOn = false
			gs.Player.Pl.cnt.Y = gs.Shop.ShopExitY
			upPlayerRec()
			gs.Core.Pause = false
		} else {
			if gs.Player.Mods.wallet {
				gs.Player.Mods.wallet = false
				clearinven("wallet")
				gs.Shop.ShopItems[gs.Shop.ShopNum].shopoff = true
				addshopitem(gs.Shop.ShopNum)
				rl.PlaySound(gs.Audio.Sfx[21])
			} else if !gs.Shop.ShopItems[gs.Shop.ShopNum].shopoff && gs.Player.Pl.coins >= gs.Shop.ShopItems[gs.Shop.ShopNum].shopprice {
				if !gs.Shop.ShopItems[gs.Shop.ShopNum].shopoff {
					gs.Shop.ShopItems[gs.Shop.ShopNum].shopoff = true
					gs.Player.Pl.coins -= gs.Shop.ShopItems[gs.Shop.ShopNum].shopprice
				}
				gs.Shop.ShopItems[gs.Shop.ShopNum].shopoff = true
				addshopitem(gs.Shop.ShopNum)
				rl.PlaySound(gs.Audio.Sfx[21])
			} else {
				rl.PlaySound(gs.Audio.Sfx[22])
			}
		}
	}

	txt := "shop"
	txtlen := rl.MeasureText(txt, txU5)
	txtx := int32(gs.Core.Cnt.X) - txtlen/2
	txty := int32(gs.Level.LevY) + txU
	rl.DrawText(txt, txtx, txty, txU5, rl.White)

	siz := bsU5
	y := gs.Level.LevY + bsU6
	x := gs.Core.Cnt.X
	x -= siz * 2

	rec := rl.NewRectangle(x, y, siz, siz)
	if gs.Shop.ShopNum == 0 {
		if gs.Shop.ShopItems[0].shopoff {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Red, gs.UI.FadeBlink))
		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Green, gs.UI.FadeBlink))
		}
	}
	if gs.Shop.ShopItems[0].shopoff {
		col := ranRed()
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[0].img, rec, rl.Vector2Zero(), 0, col)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[0].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(col, 0.2))
	} else {

		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[0].img, rec, rl.Vector2Zero(), 0, gs.Shop.ShopItems[0].color)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[0].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(gs.Shop.ShopItems[0].color, 0.2))
		coinx := rec.X + rec.Width + bsU
		coiny := rec.Y + bsU
		txtlen = rl.MeasureText("x"+fmt.Sprint(gs.Shop.ShopItems[0].shopprice), gs.UI.TxtSize)
		txtx = int32(coinx+siz/4) - txtlen/2
		txty = int32(coiny+siz/2) + bsUi32/3
		rl.DrawText("x"+fmt.Sprint(gs.Shop.ShopItems[0].shopprice), txtx, txty, gs.UI.TxtSize, rl.White)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Coin, rl.NewRectangle(coinx, coiny, siz/2, siz/2), gs.Core.Ori, 0, rl.White)
	}
	txtlen = rl.MeasureText(gs.Shop.ShopItems[0].name, gs.UI.TxtSize)
	rl.DrawText(gs.Shop.ShopItems[0].name, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, rec.ToInt32().Y+rec.ToInt32().Height+bsUi32/4, gs.UI.TxtSize, rl.White)

	rec.X += (siz * 2) + siz/2
	if gs.Shop.ShopNum == 1 {
		if gs.Shop.ShopItems[1].shopoff {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Red, gs.UI.FadeBlink))
		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Green, gs.UI.FadeBlink))
		}
	}

	if gs.Shop.ShopItems[1].shopoff {
		col := ranRed()
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[1].img, rec, rl.Vector2Zero(), 0, col)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[1].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(col, 0.2))
	} else {
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[1].img, rec, rl.Vector2Zero(), 0, gs.Shop.ShopItems[1].color)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[1].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(gs.Shop.ShopItems[1].color, 0.2))
		coinx := rec.X + rec.Width + bsU
		coiny := rec.Y + bsU
		txtlen = rl.MeasureText("x"+fmt.Sprint(gs.Shop.ShopItems[1].shopprice), gs.UI.TxtSize)
		txtx = int32(coinx+siz/4) - txtlen/2
		txty = int32(coiny+siz/2) + bsUi32/3
		rl.DrawText("x"+fmt.Sprint(gs.Shop.ShopItems[1].shopprice), txtx, txty, gs.UI.TxtSize, rl.White)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Coin, rl.NewRectangle(coinx, coiny, siz/2, siz/2), gs.Core.Ori, 0, rl.White)
	}
	txtlen = rl.MeasureText(gs.Shop.ShopItems[1].name, gs.UI.TxtSize)
	rl.DrawText(gs.Shop.ShopItems[1].name, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, rec.ToInt32().Y+rec.ToInt32().Height+bsUi32/4, gs.UI.TxtSize, rl.White)

	rec.Y += siz * 2
	if gs.Shop.ShopNum == 3 {
		if gs.Shop.ShopItems[3].shopoff {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Red, gs.UI.FadeBlink))
		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Green, gs.UI.FadeBlink))
		}
	}

	if gs.Shop.ShopItems[3].shopoff {
		col := ranRed()
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[3].img, rec, rl.Vector2Zero(), 0, col)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[3].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(col, 0.2))
	} else {
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[3].img, rec, rl.Vector2Zero(), 0, gs.Shop.ShopItems[3].color)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[3].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(gs.Shop.ShopItems[3].color, 0.2))
		coinx := rec.X + rec.Width + bsU
		coiny := rec.Y + bsU
		txtlen = rl.MeasureText("x"+fmt.Sprint(gs.Shop.ShopItems[3].shopprice), gs.UI.TxtSize)
		txtx = int32(coinx+siz/4) - txtlen/2
		txty = int32(coiny+siz/2) + bsUi32/3
		rl.DrawText("x"+fmt.Sprint(gs.Shop.ShopItems[3].shopprice), txtx, txty, gs.UI.TxtSize, rl.White)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Coin, rl.NewRectangle(coinx, coiny, siz/2, siz/2), gs.Core.Ori, 0, rl.White)
	}
	txtlen = rl.MeasureText(gs.Shop.ShopItems[3].name, gs.UI.TxtSize)
	rl.DrawText(gs.Shop.ShopItems[3].name, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, rec.ToInt32().Y+rec.ToInt32().Height+bsUi32/4, gs.UI.TxtSize, rl.White)

	rec.X -= (siz * 2) + siz/2
	if gs.Shop.ShopNum == 2 {
		if gs.Shop.ShopItems[2].shopoff {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Red, gs.UI.FadeBlink))
		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Green, gs.UI.FadeBlink))
		}
	}

	if gs.Shop.ShopItems[2].shopoff {
		col := ranRed()
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[2].img, rec, rl.Vector2Zero(), 0, col)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[2].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(col, 0.2))
	} else {
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[2].img, rec, rl.Vector2Zero(), 0, gs.Shop.ShopItems[2].color)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Shop.ShopItems[2].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(gs.Shop.ShopItems[2].color, 0.2))
		coinx := rec.X + rec.Width + bsU
		coiny := rec.Y + bsU
		txtlen = rl.MeasureText("x"+fmt.Sprint(gs.Shop.ShopItems[2].shopprice), gs.UI.TxtSize)
		txtx = int32(coinx+siz/4) - txtlen/2
		txty = int32(coiny+siz/2) + bsUi32/3
		rl.DrawText("x"+fmt.Sprint(gs.Shop.ShopItems[2].shopprice), txtx, txty, gs.UI.TxtSize, rl.White)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Coin, rl.NewRectangle(coinx, coiny, siz/2, siz/2), gs.Core.Ori, 0, rl.White)
	}
	txtlen = rl.MeasureText(gs.Shop.ShopItems[2].name, gs.UI.TxtSize)
	rl.DrawText(gs.Shop.ShopItems[2].name, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, rec.ToInt32().Y+rec.ToInt32().Height+bsUi32/4, gs.UI.TxtSize, rl.White)

	//WALLET
	if gs.Player.Mods.wallet {
		walletx := gs.Core.Cnt.X - siz*2 + siz/8
		wallety := rec.Y + rec.Height + bsU4 + siz/8
		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[11], rl.NewRectangle(walletx, wallety, siz-siz/4, siz-siz/4), gs.Core.Ori, 0, ranBrown())
	}
	//PL COINS
	coinx := gs.Core.Cnt.X - siz
	coiny := rec.Y + rec.Height + bsU4
	rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Coin, rl.NewRectangle(coinx, coiny, siz, siz), gs.Core.Ori, 0, rl.White)

	txtx = int32(coinx + siz)
	txty = int32(coiny + siz/4)
	rl.DrawText("x"+fmt.Sprint(gs.Player.Pl.coins), txtx, txty, gs.UI.TxtSize*2, rl.White)

	if gs.Core.Frames%6 == 0 {
		gs.Render.Coin.X += 16
		if gs.Render.Coin.X >= 1200 {
			gs.Render.Coin.X = 1120
		}
	}

	//EXIT
	txtlen = rl.MeasureText("exit", gs.UI.TxtSize*2)
	txtx = int32(gs.Core.Cnt.X) - txtlen/2
	txty = int32(coiny + siz + bsU2)

	wid := float32(txtlen) + bsU2
	heig := float32(gs.UI.TxtSize*2) + bsU/2

	rec = rl.NewRectangle(gs.Core.Cnt.X-wid/2, float32(txty)-bsU/4, wid, heig)

	if gs.Shop.ShopNum == 4 {
		rl.DrawRectangleRec(rec, ranRed())
	} else {
		rl.DrawRectangleLinesEx(rec, 2, ranCol())
	}

	rl.DrawText("exit", txtx-2, txty+2, gs.UI.TxtSize*2, rl.Black)
	rl.DrawText("exit", txtx, txty, gs.UI.TxtSize*2, rl.White)

}
func drawHelp() { //MARK:DRAW HELP

	rl.DrawRectangle(0, 0, gs.Core.ScrW32, gs.Core.ScrH32, rl.Black)
	txt := "help"
	txtlen := rl.MeasureText(txt, txU5)
	txtx := int32(gs.Core.Cnt.X) - txtlen/2
	txty := int32(gs.Level.LevY) + txU
	rl.DrawText(txt, txtx, txty, txU5, rl.White)

	txtx = int32(gs.Core.Cnt.X - bsU9)
	txty += txU7

	txt = "five levels collect power ups"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "kill enemies avoid traps"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "last level defeat boss"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txty += gs.UI.TxtSize + txU/4
	txt = "WASD keys / xbox left stick/dpad > move"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "SPACE key / xbox a > attack"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "TAB key / xbox y > inventory"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "RIGHT CTRL key / xbox b > map"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "ESC key / xbox menu > options/exit"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "END key > exits game"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)

	if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
		gs.UI.HelpOn = false
	}

}
func drawDied() { //MARK:DRAW DIED

	rl.DrawRectangle(0, 0, gs.Core.ScrW32, gs.Core.ScrH32, rl.Black)

	if gs.Level.DiedscrT > 0 {
		gs.Level.DiedscrT--
	}

	rl.DrawTexturePro(gs.Render.Imgs, gs.Player.DiedIMG, gs.Player.DiedRec, rl.Vector2Zero(), 0, ranRed())
	rl.DrawTexturePro(gs.Render.Imgs, gs.Player.DiedIMG, BlurRec(gs.Player.DiedRec, 10), rl.Vector2Zero(), 0, rl.Fade(ranRed(), rF32(0.1, 0.4)))
	gs.Player.DiedRec.X -= 2
	gs.Player.DiedRec.Y -= 2
	gs.Player.DiedRec.Width += 4
	gs.Player.DiedRec.Height += 4

	if gs.Player.DiedRec.Y <= gs.Level.LevRecInner.Y || rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) && gs.Level.DiedscrT == 0 {
		gs.Player.Died = false
		gs.Timing.BestTime = false
		gs.Timing.TimesOn = true
		gs.Timing.BestTimesT = gs.Core.Fps
	}
	txt := "you"
	txtlen := rl.MeasureText(txt, txU8)
	rl.DrawText(txt, int32(gs.Core.Cnt.X)-txtlen/2-3, int32(gs.Core.Cnt.Y)-txU8+3, txU8, rl.Black)
	rl.DrawText(txt, int32(gs.Core.Cnt.X)-txtlen/2, int32(gs.Core.Cnt.Y)-txU8, txU8, rl.White)
	txt = "died"
	txtlen = rl.MeasureText(txt, txU8)
	rl.DrawText(txt, int32(gs.Core.Cnt.X)-txtlen/2-3, int32(gs.Core.Cnt.Y)+txU+3, txU8, rl.Black)
	rl.DrawText(txt, int32(gs.Core.Cnt.X)-txtlen/2, int32(gs.Core.Cnt.Y)+txU, txU8, rl.White)

	txt = "new best time"
	txtlen = rl.MeasureText(txt, txU4)
	rl.DrawText(txt, int32(gs.Core.Cnt.X)-txtlen/2-3, gs.Core.ScrH32-txU5+3, txU4, rl.Black)
	rl.DrawText(txt, int32(gs.Core.Cnt.X)-txtlen/2, gs.Core.ScrH32-txU5, txU4, rl.White)

}
func drawTimes() { //MARK:DRAW TIMES

	if gs.Timing.BestTimesT > 0 {
		gs.Timing.BestTimesT--
	}

	rl.DrawRectangle(0, 0, gs.Core.ScrW32, gs.Core.ScrH32, rl.Black)
	txt := "best times"
	txtlen := rl.MeasureText(txt, txU5)
	txtx := int32(gs.Core.Cnt.X) - txtlen/2
	txty := int32(gs.Level.LevY) + txU
	rl.DrawText(txt, txtx, txty, txU5, rl.White)

	txty += txU7

	for i := 0; i < len(gs.Timing.Times); i++ {
		minutes, seconds := gs.Timing.Times[i]/60, gs.Timing.Times[i]%60
		minTXT := fmt.Sprint(minutes)
		secsTXT := fmt.Sprint(seconds)
		if seconds < 10 {
			if seconds == 0 {
				secsTXT = "00"
			} else {
				secsTXT = "0" + secsTXT
			}
		}
		if minutes < 10 {
			if minutes == 0 {
				minTXT = "00"
			} else {
				minTXT = "0" + minTXT
			}
		}
		timesTXT := minTXT + ":" + secsTXT
		txtlen := rl.MeasureText(timesTXT, gs.UI.TxtSize*2)
		rl.DrawText(timesTXT, int32(gs.Core.Cnt.X)-txtlen/2, txty, gs.UI.TxtSize*2, rl.White)
		txty += gs.UI.TxtSize*2 + txU/2
	}

	if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
		if gs.Timing.BestTimesT <= 0 {
			gs.Timing.TimesOn = false
			if !gs.UI.OptionsOn {
				restartgame()
			}
		}
	}

}

func drawCredits() { //MARK:DRAW CREDITS

	rl.DrawRectangle(0, 0, gs.Core.ScrW32, gs.Core.ScrH32, rl.Black)
	txt := "credits"
	txtlen := rl.MeasureText(txt, txU5)
	txtx := int32(gs.Core.Cnt.X) - txtlen/2
	txty := int32(gs.Level.LevY) + txU
	rl.DrawText(txt, txtx, txty, txU5, rl.White)

	txtx = int32(gs.Core.Cnt.X - bsU9)
	txty += txU7

	txt = "kenney.nl"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "laredgames.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "pixelfelix.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "piiixl.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "stealthix.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "rad-potato.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "pixelfrog-assets.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "free-game-assets.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "bdragon1727.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "kamioo.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "luquigames.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "nebelstern.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "bit-by-bit-sound.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "ironchestgames.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "pixeljad.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "magory.itch.io"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "opengameart.org/users/subspaceaudio"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
	txty += gs.UI.TxtSize + txU/4
	txt = "opengameart.org/users/rubberduck"
	rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)

	if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
		gs.UI.CreditsOn = false
	}

}
func drawresettimes() { //MARK:DRAW RESET TIMES
	txt := "reset times"
	txtlen := rl.MeasureText(txt, txU5)
	txtx := int32(gs.Core.Cnt.X) - txtlen/2
	txty := int32(gs.Core.Cnt.Y) - txU8
	rl.DrawText(txt, txtx, txty, txU5, rl.White)
	txty += txU6
	rec := rl.NewRectangle(gs.Core.Cnt.X-bsU3, float32(txty), bsU3, bsU2)
	recX := rec.ToInt32().X
	if gs.Level.ExitLR {
		rec.X += rec.Width
		if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
			for i := 0; i < len(gs.Timing.Times); i++ {
				gs.Timing.Times[i] = 600
			}
			savetimes()
			gs.UI.Resettimes = false
		}
		rl.DrawRectangleRec(rec, rl.Green)
	} else {
		rl.DrawRectangleRec(rec, rl.Red)
		if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
			gs.UI.Resettimes = false
		}
	}

	txtlen = rl.MeasureText("Y", txU3)
	txtx = recX + rec.ToInt32().Width + rec.ToInt32().Width/2 - txtlen/2
	txty = rec.ToInt32().Y + 2
	rl.DrawText("Y", txtx-2, txty+2, txU3, rl.Black)
	rl.DrawText("Y", txtx, txty, txU3, rl.White)

	txtlen = rl.MeasureText("N", txU3)
	txtx = recX + rec.ToInt32().Width/2 - txtlen/2
	txty = rec.ToInt32().Y + 2
	rl.DrawText("N", txtx-2, txty+2, txU3, rl.Black)
	rl.DrawText("N", txtx, txty, txU3, rl.White)

	if rl.IsKeyPressed(rl.KeyA) || rl.GetGamepadAxisMovement(0, 0) < 0 && rl.GetGamepadAxisMovement(0, 0) > -0.3 || rl.IsGamepadButtonDown(0, 4) {
		if gs.UI.OptionT == 0 {
			gs.UI.OptionT = gs.Core.Fps / 5
			gs.Level.ExitLR = !gs.Level.ExitLR
		}
	} else if rl.IsKeyPressed(rl.KeyD) || rl.GetGamepadAxisMovement(0, 0) > 0 && rl.GetGamepadAxisMovement(0, 0) < 0.3 || rl.IsGamepadButtonDown(0, 2) {
		if gs.UI.OptionT == 0 {
			gs.UI.OptionT = gs.Core.Fps / 5
			gs.Level.ExitLR = !gs.Level.ExitLR
		}
	}

	if rl.IsKeyPressed(rl.KeyY) {
		for i := 0; i < len(gs.Timing.Times); i++ {
			gs.Timing.Times[i] = 600
		}
		savetimes()
		gs.UI.Resettimes = false
	} else if rl.IsKeyPressed(rl.KeyN) {
		gs.UI.Resettimes = false
	}

}
func drawrestartconfirm() { //MARK:DRAW RESTART CONFIRM
	txt := "restart"
	txtlen := rl.MeasureText(txt, txU5)
	txtx := int32(gs.Core.Cnt.X) - txtlen/2
	txty := int32(gs.Core.Cnt.Y) - txU8
	rl.DrawText(txt, txtx, txty, txU5, rl.White)
	txty += txU6
	rec := rl.NewRectangle(gs.Core.Cnt.X-bsU3, float32(txty), bsU3, bsU2)
	recX := rec.ToInt32().X
	if gs.Level.ExitLR {
		rec.X += rec.Width
		if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
			restartgame()
			gs.UI.OptionsOn = false
			gs.UI.RestartOn = false
		}
		rl.DrawRectangleRec(rec, rl.Green)
	} else {
		rl.DrawRectangleRec(rec, rl.Red)
		if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
			gs.UI.RestartOn = false
		}
	}

	txtlen = rl.MeasureText("Y", txU3)
	txtx = recX + rec.ToInt32().Width + rec.ToInt32().Width/2 - txtlen/2
	txty = rec.ToInt32().Y + 2
	rl.DrawText("Y", txtx-2, txty+2, txU3, rl.Black)
	rl.DrawText("Y", txtx, txty, txU3, rl.White)

	txtlen = rl.MeasureText("N", txU3)
	txtx = recX + rec.ToInt32().Width/2 - txtlen/2
	txty = rec.ToInt32().Y + 2
	rl.DrawText("N", txtx-2, txty+2, txU3, rl.Black)
	rl.DrawText("N", txtx, txty, txU3, rl.White)

	if rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyLeft) || rl.GetGamepadAxisMovement(0, 0) < 0 && rl.GetGamepadAxisMovement(0, 0) > -0.3 || rl.IsGamepadButtonDown(0, 4) {
		if gs.UI.OptionT == 0 {
			gs.UI.OptionT = gs.Core.Fps / 5
			gs.Level.ExitLR = !gs.Level.ExitLR
		}
	} else if rl.IsKeyPressed(rl.KeyD) || rl.IsKeyPressed(rl.KeyRight) || rl.GetGamepadAxisMovement(0, 0) > 0 && rl.GetGamepadAxisMovement(0, 0) < 0.3 || rl.IsGamepadButtonDown(0, 2) {
		if gs.UI.OptionT == 0 {
			gs.UI.OptionT = gs.Core.Fps / 5
			gs.Level.ExitLR = !gs.Level.ExitLR
		}
	}

	if rl.IsKeyPressed(rl.KeyY) {
		restartgame()
		gs.UI.OptionsOn = false
		gs.UI.RestartOn = false
	} else if rl.IsKeyPressed(rl.KeyN) {
		gs.UI.RestartOn = false
	}

}
func drawExit() { //MARK:DRAW EXIT

	txt := "exit"
	txtlen := rl.MeasureText(txt, txU5)
	txtx := int32(gs.Core.Cnt.X) - txtlen/2
	txty := int32(gs.Core.Cnt.Y) - txU8
	rl.DrawText(txt, txtx, txty, txU5, rl.White)
	txty += txU6
	rec := rl.NewRectangle(float32(txtx), float32(txty), bsU3, bsU2)
	recX := rec.ToInt32().X
	if gs.Level.ExitLR {
		rec.X += rec.Width
		if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
			exitgame()
		}
		rl.DrawRectangleRec(rec, rl.Green)
	} else {
		rl.DrawRectangleRec(rec, rl.Red)
		if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
			gs.Level.Exiton = false
		}
	}

	txtlen = rl.MeasureText("Y", txU3)
	txtx = recX + rec.ToInt32().Width + rec.ToInt32().Width/2 - txtlen/2
	txty = rec.ToInt32().Y + 2
	rl.DrawText("Y", txtx-2, txty+2, txU3, rl.Black)
	rl.DrawText("Y", txtx, txty, txU3, rl.White)

	txtlen = rl.MeasureText("N", txU3)
	txtx = recX + rec.ToInt32().Width/2 - txtlen/2
	txty = rec.ToInt32().Y + 2
	rl.DrawText("N", txtx-2, txty+2, txU3, rl.Black)
	rl.DrawText("N", txtx, txty, txU3, rl.White)

	if rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyLeft) || rl.GetGamepadAxisMovement(0, 0) < 0 && rl.GetGamepadAxisMovement(0, 0) > -0.3 || rl.IsGamepadButtonDown(0, 4) {
		if gs.UI.OptionT == 0 {
			gs.UI.OptionT = gs.Core.Fps / 5
			gs.Level.ExitLR = !gs.Level.ExitLR
		}
	} else if rl.IsKeyPressed(rl.KeyD) || rl.IsKeyPressed(rl.KeyRight) || rl.GetGamepadAxisMovement(0, 0) > 0 && rl.GetGamepadAxisMovement(0, 0) < 0.3 || rl.IsGamepadButtonDown(0, 2) {
		if gs.UI.OptionT == 0 {
			gs.UI.OptionT = gs.Core.Fps / 5
			gs.Level.ExitLR = !gs.Level.ExitLR
		}
	}

	if rl.IsKeyPressed(rl.KeyY) {
		exitgame()
	} else if rl.IsKeyPressed(rl.KeyN) {
		gs.Level.Exiton = false
	}

}
func drawOptions() { //MARK:DRAW OPTIONS
	//rl.ShowCursor()

	rl.DrawRectangle(0, 0, gs.Core.ScrW32, gs.Core.ScrH32, rl.Black)
	if gs.UI.CreditsOn {
		drawCredits()
	} else if gs.Timing.TimesOn {
		drawTimes()
	} else if gs.UI.HelpOn {
		drawHelp()
	} else if gs.Level.Exiton {
		drawExit()
	} else if gs.UI.Resettimes {
		drawresettimes()
	} else if gs.UI.RestartOn {
		drawrestartconfirm()
	} else {

		txt := "options"
		txtlen := rl.MeasureText(txt, txU5)
		txtx := int32(gs.Core.Cnt.X) - txtlen/2
		txty := int32(gs.Level.LevY) + txU
		rl.DrawText(txt, txtx, txty, txU5, rl.White)

		txty += txU7
		txtx = int32(gs.Core.Cnt.X - bsU7)
		onoffx := txtx + int32(gs.Level.LevRec.Width/3) - gs.UI.TxtSize*2

		rec := rl.NewRectangle(float32(txtx)-bsU/2, float32(txty)-bsU/4, bsU*17, bsU2-bsU/4)
		rec.Y += float32(gs.UI.OptionNum) * float32(gs.UI.TxtSize+gs.UI.TxtSize/2)
		if gs.UI.OptionNum == 10 || gs.UI.OptionNum == 11 || gs.UI.OptionNum == 12 || gs.UI.OptionNum == 13 || gs.UI.OptionNum == 14 || gs.UI.OptionNum == 15 || gs.UI.OptionNum == 16 {
			rec.Y += float32(gs.UI.TxtSize + gs.UI.TxtSize/2)
		}

		//KEYS GAMEPAD INP
		if rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyUp) || rl.GetGamepadAxisMovement(0, 1) < 0 || rl.IsGamepadButtonDown(0, 1) {
			if gs.UI.OptionT == 0 {
				gs.UI.OptionT = gs.Core.Fps / 5
				gs.UI.OptionNum--
				if gs.UI.OptionNum < 0 {
					gs.UI.OptionNum = 16
				}
			}
		}
		if rl.IsKeyPressed(rl.KeyS) || rl.IsKeyPressed(rl.KeyDown) || rl.GetGamepadAxisMovement(0, 1) > 0 || rl.IsGamepadButtonDown(0, 3) {
			if gs.UI.OptionT == 0 {
				gs.UI.OptionT = gs.Core.Fps / 5
				gs.UI.OptionNum++
				if gs.UI.OptionNum > 16 {
					gs.UI.OptionNum = 0
				}
			}
		}
		if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
			switch gs.UI.OptionNum {
			case 0:
				gs.UI.HpBarsOn = !gs.UI.HpBarsOn
			case 1:
				gs.UI.ScanLinesOn = !gs.UI.ScanLinesOn
			case 2:
				gs.UI.ArtifactsOn = !gs.UI.ArtifactsOn
			case 3:
				gs.Render.ShaderOn = !gs.Render.ShaderOn
			case 4:
				gs.Player.PlatkrecOn = !gs.Player.PlatkrecOn
			case 5:
				gs.UI.Invincible = !gs.UI.Invincible
			case 6:
				if gs.Input.IsController && gs.Input.UseController {
					gs.Input.ControllerOn = false
					gs.Input.UseController = false
				} else if gs.Input.IsController && !gs.Input.UseController {
					gs.Input.ControllerOn = true
					gs.Input.UseController = true
				} else if !gs.Input.IsController {
					gs.Input.ControllerOn = false
					gs.Input.UseController = false
				}
			case 7:
				if gs.Audio.MusicOn {
					gs.Audio.MusicOn = false
				} else {
					rl.StopMusicStream(gs.Audio.Music)
					gs.Audio.Music = gs.Audio.BackMusic[gs.Audio.BgMusicNum]
					gs.Audio.Music.Looping = true
					gs.Audio.MusicOn = true
					rl.PlayMusicStream(gs.Audio.Music)
					gs.Audio.MusicOn = true
				}

			case 10:
				gs.UI.RestartOn = true
			case 11:
				gs.Timing.TimesOn = true
			case 12:
				gs.UI.HelpOn = true
			case 13:
				gs.UI.CreditsOn = true
			case 14:
				gs.UI.Resettimes = true
			case 15:
				gs.Level.Hardcore = !gs.Level.Hardcore
				restartgame()
				gs.UI.OptionsOn = false
			case 16:
				gs.Level.Exiton = true
				gs.Level.ExitLR = false
			}
			gs.UI.OptionsChange = true
		}

		//OPTIONS LIST
		rl.DrawRectangleRec(rec, rl.Fade(ranCol(), gs.UI.FadeBlink2))

		txt = "hp bars"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		gs.UI.HpBarsOn = onoff(onoffx, txty, float32(gs.UI.TxtSize), gs.UI.HpBarsOn)
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2

		txt = "scan lines"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		gs.UI.ScanLinesOn = onoff(onoffx, txty, float32(gs.UI.TxtSize), gs.UI.ScanLinesOn)
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2

		txt = "pixel artifacts"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		gs.UI.ArtifactsOn = onoff(onoffx, txty, float32(gs.UI.TxtSize), gs.UI.ArtifactsOn)
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2

		txt = "bloom 'fuzzy'"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		gs.Render.ShaderOn = onoff(onoffx, txty, float32(gs.UI.TxtSize), gs.Render.ShaderOn)
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2

		txt = "player atk range"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		gs.Player.PlatkrecOn = onoff(onoffx, txty, float32(gs.UI.TxtSize), gs.Player.PlatkrecOn)
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2

		txt = "gs.UI.Invincible"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		gs.UI.Invincible = onoff(onoffx, txty, float32(gs.UI.TxtSize), gs.UI.Invincible)
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2

		txt = "use controller"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		gs.Input.UseController = onoff(onoffx, txty, float32(gs.UI.TxtSize), gs.Input.UseController)
		if !gs.Input.IsController && gs.UI.OptionNum == 6 {
			txt = "no controller detected"
			rl.DrawText(txt, txtx, txty+(gs.UI.TxtSize+gs.UI.TxtSize/2)*4, gs.UI.TxtSize, ranCol())
		}
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2
		txt = "music"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		gs.Audio.MusicOn = onoff(onoffx, txty, float32(gs.UI.TxtSize), gs.Audio.MusicOn)
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2

		txt = "music track"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		gs.Audio.BgMusicNum = int(updownswitch(onoffx, txty, float32(gs.UI.TxtSize), float32(gs.Audio.BgMusicNum), 2))
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2

		//CHANGE MUSIC TRACK
		if gs.UI.OptionNum == 8 {
			txt = "use left / right to adjust"
			rl.DrawText(txt, txtx, txty+gs.UI.TxtSize+gs.UI.TxtSize/2, gs.UI.TxtSize, ranCol())
			if rl.IsKeyPressed(rl.KeyD) || rl.IsKeyPressed(rl.KeyRight) || rl.GetGamepadAxisMovement(0, 0) > 0 || rl.IsGamepadButtonDown(0, 2) {
				if gs.UI.OptionT == 0 {
					gs.UI.OptionT = gs.Core.Fps / 5
					gs.Audio.BgMusicNum++
					if gs.Audio.BgMusicNum > 2 {
						gs.Audio.BgMusicNum = 0
					}
					rl.StopMusicStream(gs.Audio.Music)
					gs.Audio.Music = gs.Audio.BackMusic[gs.Audio.BgMusicNum]
					gs.Audio.Music.Looping = true
					gs.Audio.MusicOn = true
					rl.PlayMusicStream(gs.Audio.Music)
					gs.UI.OptionsChange = true
				}
			} else if rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyLeft) || rl.GetGamepadAxisMovement(0, 0) < 0 || rl.IsGamepadButtonDown(0, 4) {
				if gs.UI.OptionT == 0 {
					gs.UI.OptionT = gs.Core.Fps / 5
					gs.Audio.BgMusicNum--
					if gs.Audio.BgMusicNum < 0 {
						gs.Audio.BgMusicNum = 2
					}
					rl.StopMusicStream(gs.Audio.Music)
					gs.Audio.Music = gs.Audio.BackMusic[gs.Audio.BgMusicNum]
					gs.Audio.Music.Looping = true
					gs.Audio.MusicOn = true
					rl.PlayMusicStream(gs.Audio.Music)
					gs.UI.OptionsChange = true
				}
			}
		}

		txt = "volume - 0 is off"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		gs.Audio.Volume = updownswitch(onoffx, txty, float32(gs.UI.TxtSize), gs.Audio.Volume, 1)
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2

		//VOLUME LR
		if gs.UI.OptionNum == 9 {
			txt = "use left / right to adjust"
			rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, ranCol())
			if rl.IsKeyPressed(rl.KeyD) || rl.IsKeyPressed(rl.KeyRight) || rl.GetGamepadAxisMovement(0, 0) > 0 || rl.IsGamepadButtonDown(0, 2) {
				if gs.UI.OptionT == 0 {
					gs.UI.OptionT = gs.Core.Fps / 5
					if gs.Audio.Volume < 1 {
						gs.Audio.Volume += 0.1
					}
					gs.UI.OptionsChange = true
				}
			} else if rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyLeft) || rl.GetGamepadAxisMovement(0, 0) < 0 || rl.IsGamepadButtonDown(0, 4) {
				if gs.UI.OptionT == 0 {
					gs.UI.OptionT = gs.Core.Fps / 5
					if gs.Audio.Volume > 0 {
						gs.Audio.Volume -= 0.1
					}
					if gs.Audio.Volume < 0 {
						gs.Audio.Volume = 0
					}
					gs.UI.OptionsChange = true
				}
			}
		}

		txty += gs.UI.TxtSize + gs.UI.TxtSize/2
		txt = "restart game"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2
		txt = "best times"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2
		txt = "help"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2
		txt = "credits"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2
		txt = "reset times"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2
		txt = "hardcore"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		gs.Level.Hardcore = onoff(onoffx, txty, float32(gs.UI.TxtSize), gs.Level.Hardcore)
		if gs.UI.OptionNum == 15 {
			txt = "more enemies > game will restart"
			rl.DrawText(txt, txtx, txty-(gs.UI.TxtSize+gs.UI.TxtSize/2)*6, gs.UI.TxtSize, ranCol())
		}
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2
		txt = "exit"
		rl.DrawText(txt, txtx, txty, gs.UI.TxtSize, rl.White)
		txty += gs.UI.TxtSize + gs.UI.TxtSize/2

	}

}
func drawUpMario() { //MARK:DRAW UP MARIO

	//TIMER
	rl.DrawText(fmt.Sprint(gs.Mario.MarioT), gs.Mario.MarioScreenRec.ToInt32().X+bsUi32, gs.Mario.MarioScreenRec.ToInt32().Y+bsUi32, gs.UI.TxtSize, rl.White)
	gs.Mario.MarioT--

	//INP
	if rl.IsKeyDown(rl.KeyD) || rl.GetGamepadAxisMovement(0, 0) > 0 || rl.IsGamepadButtonDown(0, 2) {
		gs.Mario.MarioPL.X += 8
		gs.Mario.MarioV2L.X += 8
		gs.Mario.MarioV2R.X += 8
		gs.Mario.MarioImg.Y = gs.Render.Knight[0].Y

	} else if rl.IsKeyDown(rl.KeyA) || rl.GetGamepadAxisMovement(0, 0) < 0 || rl.IsGamepadButtonDown(0, 4) {
		gs.Mario.MarioPL.X -= 8
		gs.Mario.MarioV2L.X -= 8
		gs.Mario.MarioV2R.X -= 8
		gs.Mario.MarioImg.Y = gs.Render.Knight[2].Y
	}
	if rl.IsKeyPressed(rl.KeyW) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
		if !gs.Mario.MarioJump {
			gs.Mario.MarioJump = true
			gs.Mario.MarioJumpT = gs.Core.Fps / 3
		}
	}

	//DRAW PATTERN
	x := gs.Mario.MarioScreenRec.X
	y := gs.Mario.MarioScreenRec.Y
	siz := bsU10
	for {
		rl.DrawTexturePro(gs.Render.Imgs, gs.Mario.PatternRec, rl.NewRectangle(x, y, siz, siz), gs.Core.Ori, 0, rl.Fade(ranCol(), 0.05))
		x += siz
		if x >= gs.Mario.MarioScreenRec.X+gs.Mario.MarioScreenRec.Width {
			x = gs.Mario.MarioScreenRec.X
			y += siz
		}
		if y >= gs.Mario.MarioScreenRec.Y+gs.Mario.MarioScreenRec.Height {
			break
		}
	}

	//DRAW BLOKS
	for a := 0; a < len(gs.Mario.MarioRecs); a++ {
		rl.DrawTexturePro(gs.Render.Imgs, gs.Level.WallT, gs.Mario.MarioRecs[a], gs.Core.Ori, 0, gs.Mario.MarioCols[a])
		rl.DrawTexturePro(gs.Render.Imgs, gs.Level.WallT, BlurRec(gs.Mario.MarioRecs[a], 2), gs.Core.Ori, 0, rl.Fade(gs.Mario.MarioCols[a], 0.2))
	}

	//DRAW COINS
	for i := 0; i < len(gs.Mario.MarioCoinOnOff); i++ {
		if gs.Mario.MarioCoinOnOff[i] {
			rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Coin, gs.Mario.MarioCoins[i], gs.Core.Ori, 0, rl.White)
			if gs.Mario.MarioT%6 == 0 {
				gs.Render.Coin.X += 16
				if gs.Render.Coin.X >= 1200 {
					gs.Render.Coin.X = 1120
				}
				if rl.CheckCollisionRecs(gs.Mario.MarioPL, gs.Mario.MarioCoins[i]) {
					if gs.Mario.MarioCoinOnOff[i] {
						gs.Player.Pl.coins++
						gs.Mario.MarioCoinOnOff[i] = false
						rl.PlaySound(gs.Audio.Sfx[18])
					}
				}
			}
		}

	}

	//DRAW PLAYER
	drec := gs.Mario.MarioPL
	drec.X -= drec.Width / 4
	drec.Width += drec.Width / 2
	drec.Y -= drec.Height / 4
	drec.Height += drec.Height / 2
	//ori2 := rl.NewVector2(drec.Width/2,drec.Height/2)
	rl.DrawTexturePro(gs.Render.Imgs, gs.Mario.MarioImg, drec, gs.Core.Ori, 0, rl.White)
	if gs.Core.Debug {
		rl.DrawRectangleLinesEx(gs.Mario.MarioPL, 1, rl.White)
	}

	if gs.Core.Frames%4 == 0 {
		gs.Mario.MarioImg.X += gs.Player.Pl.sizImg
	}
	if gs.Mario.MarioImg.X > gs.Player.Pl.imgWalkX+(float32(gs.Player.Pl.framesWalk-1)*gs.Player.Pl.sizImg) {
		gs.Mario.MarioImg.X = gs.Player.Pl.imgWalkX
	}

	//JUMP FALL
	if gs.Mario.MarioJumpT > 0 {
		gs.Mario.MarioJumpT--
		gs.Mario.MarioPL.Y -= 12
		gs.Mario.MarioV2L.Y -= 12
		gs.Mario.MarioV2R.Y -= 12
	} else {
		collides := false
		for a := 0; a < len(gs.Mario.MarioRecs); a++ {
			if rl.CheckCollisionPointRec(gs.Mario.MarioV2L, gs.Mario.MarioRecs[a]) || rl.CheckCollisionPointRec(gs.Mario.MarioV2R, gs.Mario.MarioRecs[a]) {
				collides = true
				gs.Mario.MarioJump = false
			}
		}
		if !collides {
			gs.Mario.MarioPL.Y += 8
			gs.Mario.MarioV2L.Y += 8
			gs.Mario.MarioV2R.Y += 8
		}

	}

	//EXIT SCREEN
	if gs.Mario.MarioT == 0 || gs.Mario.MarioPL.X+gs.Mario.MarioPL.Width < gs.Mario.MarioScreenRec.X || gs.Mario.MarioPL.X > gs.Mario.MarioScreenRec.X+gs.Mario.MarioScreenRec.Width {
		gs.Mario.MarioOn = false
		gs.Core.Pause = false
	}

}
func drawInven() { //MARK:DRAW INVENTORY

	x := gs.Level.LevX - bsU2
	y := gs.Level.LevY + bsU/2
	siz := bsU + bsU/2

	for a := 0; a < len(gs.Player.Inven); a++ {
		rec := rl.NewRectangle(x, y, siz, siz)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Player.Inven[a].img, rec, gs.Core.Ori, 0, gs.Player.Inven[a].color)

		rl.DrawTexturePro(gs.Render.Imgs, gs.Player.Inven[a].img, BlurRec(rec, 1), gs.Core.Ori, 0, rl.Fade(gs.Player.Inven[a].color, rF32(0.1, 0.2)))

		if gs.Player.Inven[a].numof > 1 {
			txtx := int32(x+siz) - txU/2
			txty := int32(y+siz) - txU/2
			rl.DrawText(fmt.Sprint(gs.Player.Inven[a].numof), txtx-2, txty-2, txU, rl.Black)
			rl.DrawText(fmt.Sprint(gs.Player.Inven[a].numof), txtx, txty, txU, rl.White)
		}

		y += siz + bsU/4
	}

}
func drawInvenDetail() { //MARK:DRAW INVENTORY DETAIL

	if len(gs.Player.Inven) > 0 {

		x := gs.Level.LevX
		y := gs.Level.LevY + bsU/4
		siz := bsU

		txtcntr := int32(gs.Core.Cnt.X)
		txty := int32(y) + txU/4

		for a := 0; a < len(gs.Player.Inven); a++ {

			txt := gs.Player.Inven[a].name + " - " + gs.Player.Inven[a].desc
			txtlen := rl.MeasureText(txt, txU)

			txtx := txtcntr - txtlen/2
			x = float32(txtx) - (siz + bsU/2)

			rl.DrawText(txt, txtx-1, txty+1, txU, rl.Black)
			rl.DrawText(txt, txtx, txty, txU, rl.White)

			rl.DrawTexturePro(gs.Render.Imgs, gs.Player.Inven[a].img, rl.NewRectangle(x, y, siz, siz), gs.Core.Ori, 0, gs.Player.Inven[a].color)

			if gs.Player.Inven[a].numof > 1 {
				txtx2 := int32(x+siz) - txU/2
				txty2 := int32(y+siz) - txU/2
				rl.DrawText(fmt.Sprint(gs.Player.Inven[a].numof), txtx2-2, txty2-2, txU, rl.Black)
				rl.DrawText(fmt.Sprint(gs.Player.Inven[a].numof), txtx2, txty2, txU, rl.White)
			}

			y += siz + bsU/4
			txty += int32(siz + bsU/4)

		}

	} else {

		txt := "you have nothing..."
		txtlen := rl.MeasureText(txt, txU8)
		rl.DrawText(txt, int32(gs.Core.Cnt.X)-txtlen/2, int32(gs.Core.Cnt.Y)-txU8, txU8, rl.White)
		txt = "find something"
		txtlen = rl.MeasureText(txt, txU8)
		rl.DrawText(txt, int32(gs.Core.Cnt.X)-txtlen/2, int32(gs.Core.Cnt.Y)+txU, txU8, rl.White)
	}

}
func drawPlayerInfo() { //MARK:DRAW PLAYER INFO

	//TIMER
	y := gs.Level.LevY + bsU/4
	txtx := int32(gs.Level.LevRec.X + gs.Level.LevRec.Width + bsU/2)
	txty := int32(y)

	minT := "0"
	if gs.Level.Mins == 0 {
		minT = "00"
	} else if gs.Level.Mins < 10 {
		minT = "0" + fmt.Sprint(gs.Level.Mins)
	} else {
		minT = fmt.Sprint(gs.Level.Mins)
	}
	secsT := "00"
	if gs.Level.Secs > 0 && gs.Level.Secs < 10 {
		secsT = "0" + fmt.Sprint(gs.Level.Secs)
	} else if gs.Level.Secs > 9 {
		secsT = fmt.Sprint(gs.Level.Secs)
	}
	txt := minT + ":" + secsT
	rl.DrawText(txt, txtx, txty, txU3, rl.White)

	//HP ARMOR
	x := gs.Level.LevX + gs.Level.LevW + bsU/2
	y = gs.Level.LevY + bsU2 + bsU/2
	siz := bsU + bsU/2

	for a := 0; a < gs.Player.Pl.hpmax; a++ {
		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[2], rl.NewRectangle(x, y, siz, siz), gs.Core.Ori, 0, rl.DarkGray)
		y += siz
	}
	yorig := y

	y = gs.Level.LevY + bsU2 + bsU/2
	for a := 0; a < gs.Player.Pl.hp; a++ {
		rec := rl.NewRectangle(x, y, siz, siz)
		if gs.Player.Pl.poison {
			rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[2], rec, gs.Core.Ori, 0, rl.Green)
			rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[2], BlurRec(rec, 2), gs.Core.Ori, 0, rl.Fade(rl.Green, rF32(0.1, 0.3)))
		} else {
			rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[2], rec, gs.Core.Ori, 0, rl.Red)
			rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[2], BlurRec(rec, 2), gs.Core.Ori, 0, rl.Fade(rl.Red, rF32(0.1, 0.3)))
		}
		y += siz
	}

	y = yorig + bsU/2

	for a := 0; a < gs.Player.Pl.armorMax; a++ {
		rec := rl.NewRectangle(x, y, siz, siz)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[38], rec, gs.Core.Ori, 0, rl.DarkGray)
		y += siz
	}

	y = yorig + bsU/2

	for a := 0; a < gs.Player.Pl.armor; a++ {
		rec := rl.NewRectangle(x, y, siz, siz)
		col := ranCyan()
		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[38], rec, gs.Core.Ori, 0, col)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[38], BlurRec(rec, 2), gs.Core.Ori, 0, rl.Fade(col, rF32(0.1, 0.3)))
		y += siz

	}

	//COINS
	y = gs.Level.LevY + gs.Level.LevW - (bsU2 + bsU/2)
	siz = bsU2
	rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Coin, rl.NewRectangle(x, y, siz, siz), gs.Core.Ori, 0, rl.White)
	if gs.Core.Frames%6 == 0 {
		gs.Render.Coin.X += 16
		if gs.Render.Coin.X >= 1200 {
			gs.Render.Coin.X = 1120
		}
	}

	txt = "x "
	txtx = int32(x+siz) + bsUi32/2
	txty = int32(y)
	rl.DrawText(txt, txtx, txty, txU3, rl.White)
	txtlen := rl.MeasureText(txt, txU3)
	txtx += txtlen
	txty += 2
	txt = fmt.Sprint(gs.Player.Pl.coins)
	rl.DrawText(txt, txtx, txty, txU3, rl.White)

	//TXT SOLD
	if len(gs.UI.TxtSoldList) > 0 {
		clear := false
		for a := 0; a < len(gs.UI.TxtSoldList); a++ {
			if gs.UI.TxtSoldList[a].onoff {
				rl.DrawText(gs.UI.TxtSoldList[a].txt, gs.UI.TxtSoldList[a].x, gs.UI.TxtSoldList[a].y, txU2, rl.Fade(gs.UI.TxtSoldList[a].col, gs.UI.TxtSoldList[a].fade))

				rl.DrawText(gs.UI.TxtSoldList[a].txt2, gs.UI.TxtSoldList[a].x, gs.UI.TxtSoldList[a].y+txU2, txU2, rl.Fade(gs.UI.TxtSoldList[a].col, gs.UI.TxtSoldList[a].fade))

				gs.UI.TxtSoldList[a].y--
				gs.UI.TxtSoldList[a].fade -= 0.01
				if gs.UI.TxtSoldList[a].fade <= 0 {
					gs.UI.TxtSoldList[a].onoff = false
				}
			} else {
				clear = true
			}
		}

		if clear {
			for a := 0; a < len(gs.UI.TxtSoldList); a++ {
				if !gs.UI.TxtSoldList[a].onoff {
					gs.UI.TxtSoldList = remTxt(gs.UI.TxtSoldList, a)
				}
			}
		}
	}

}
func drawChainLight() { //MARK:DRAW CHAIN LIGHTNING

	for a := 1; a < len(gs.FX.ChainV2); a++ {
		rl.DrawLineEx(gs.FX.ChainV2[a], gs.FX.ChainV2[a-1], rF32(2, 12), rl.Fade(ranCyan(), rF32(0.2, 0.5)))
	}
	gs.FX.ChainV2 = nil
	for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].enemies); a++ {
		if gs.Level.Level[gs.Level.RoomNum].enemies[a].cnt.X > gs.Level.LevRecInner.X && gs.Level.Level[gs.Level.RoomNum].enemies[a].cnt.Y > gs.Level.LevRecInner.Y {
			gs.FX.ChainV2 = append(gs.FX.ChainV2, gs.Level.Level[gs.Level.RoomNum].enemies[a].cnt)
		}
	}
	gs.FX.ChainLightTimer--

	if gs.FX.ChainLightTimer <= 0 {
		for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].enemies); a++ {
			gs.Level.Level[gs.Level.RoomNum].enemies[a].hppause = gs.Core.Fps / 2
			gs.Level.Level[gs.Level.RoomNum].enemies[a].hp -= 1
			if gs.Level.Level[gs.Level.RoomNum].enemies[a].hp <= 0 {
				cntr := gs.Level.Level[gs.Level.RoomNum].enemies[a].cnt
				addkill(a)
				gs.Level.Level[gs.Level.RoomNum].enemies[a].off = true
				makeFX(2, cntr)
			} else {
				playenemyhit()
			}
		}
		gs.FX.ChainLightOn = false
	}

}
func drawPlayer() { //MARK:DRAW PLAYER

	drawRec := gs.Player.Pl.rec
	drawRec.X += gs.Player.Pl.rec.Width / 2
	drawRec.Y += gs.Player.Pl.rec.Height / 2
	shadowRec := drawRec
	shadowRec.X -= 7
	shadowRec.Y += 7

	//NOT MOVING BOUNCE
	if !gs.Player.Pl.move && !gs.Player.Pl.atk {
		if roll18() == 18 {
			drawRec.Y -= 2
			shadowRec.Y -= 2
		}
	}

	//ATK REC
	if gs.Player.PlatkrecOn && gs.Player.Pl.atkTimer > 0 {
		rl.DrawRectangleRec(gs.Player.Pl.atkrec, rl.Fade(darkRed(), 0.3))
	}

	//ESCAPE VINE
	if gs.Player.Pl.escape {
		siz := bsU2
		gs.Player.PlVineRec = rl.NewRectangle(gs.Player.Pl.cnt.X-siz/2, gs.Level.LevRec.Y, siz, gs.Player.Pl.rec.Y-gs.Level.LevRec.Y)
		if gs.Core.Debug {
			rl.DrawRectangleLinesEx(gs.Player.PlVineRec, 0.5, rl.Red)
		}
		y := gs.Player.PlVineRec.Y
		x := gs.Player.PlVineRec.X
		for {
			rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[6], rl.NewRectangle(x, y, siz, siz), gs.Core.Ori, 0, rl.DarkGreen)
			y += siz
			if y > gs.Player.PlVineRec.Y+gs.Player.PlVineRec.Height {
				break
			}
		}
	}
	//DRAW BUBBLES
	if gs.Player.Mods.flood && gs.Player.Pl.underWater {
		waterRec := rl.NewRectangle((gs.Player.Pl.crec.X+gs.Player.Pl.crec.Width/2)-bsU/2, gs.Player.Pl.crec.Y-gs.Player.WaterY, bsU, bsU)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.PlantBull.recTL, waterRec, gs.Core.Ori, 0, rl.Fade(ranCyan(), gs.Player.WaterF))
		//rl.DrawRectangleRec(waterRec,ranCol())
		waterRec2 := waterRec
		change := rF32(2, bsU/2)
		waterRec2.Width -= change
		waterRec2.Height -= change

		if gs.FX.WaterLR {
			waterRec2.X += bsU
		} else {
			waterRec2.X -= bsU
		}
		if gs.FX.WaterUP {
			waterRec2.Y += bsU
		} else {
			waterRec2.Y -= bsU
		}

		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.PlantBull.recTL, waterRec2, gs.Core.Ori, 0, rl.Fade(ranCyan(), gs.Player.WaterF))

		waterRec.X += rF32(-2, 2)
		gs.Player.WaterY += rF32(2, 5)
		gs.Player.WaterF -= 0.02
		if gs.Player.WaterY > bsU10 {
			gs.Player.WaterY = 0
			gs.Player.WaterF = 1
			gs.FX.WaterLR = flipcoin()
			gs.FX.WaterUP = flipcoin()
		}

		if roll18() == 18 {
			rl.PlaySound(gs.Audio.Sfx[24])
		}

	}
	//DRAW HP HIT
	if gs.Player.Pl.hppause != 0 && !gs.Player.Pl.revived && gs.Player.Pl.armorHit { //ARMOR
		hpRec := rl.NewRectangle((gs.Player.Pl.crec.X+gs.Player.Pl.crec.Width/2)-bsU/2, gs.Player.Pl.crec.Y-gs.Player.HpHitY, bsU, bsU)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[38], hpRec, gs.Core.Ori, 0, rl.Fade(ranCyan(), gs.Player.HpHitF))
		gs.Player.HpHitY += 2
		gs.Player.HpHitF -= 0.05
	} else if gs.Player.Pl.hppause != 0 && !gs.Player.Pl.revived && !gs.Player.Pl.armorHit { //HP
		hpRec := rl.NewRectangle((gs.Player.Pl.crec.X+gs.Player.Pl.crec.Width/2)-bsU/2, gs.Player.Pl.crec.Y-gs.Player.HpHitY, bsU, bsU)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[2], hpRec, gs.Core.Ori, 0, rl.Fade(rl.Red, gs.Player.HpHitF))
		gs.Player.HpHitY += 2
		gs.Player.HpHitF -= 0.05
	} else if gs.Player.Pl.hppause != 0 && gs.Player.Pl.revived { //REVIVED
		txty := int32(gs.Player.Pl.rec.Y-gs.Player.ReviveY) - txU2
		txt := "revived"
		txtlen := rl.MeasureText(txt, txU2)
		txtx := int32(gs.Player.Pl.rec.X+gs.Player.Pl.rec.Width/2) - txtlen/2
		rl.DrawText(txt, txtx, txty, txU2, rl.Fade(rl.White, gs.Player.ReviveF))
		gs.Player.ReviveY += 2
		gs.Player.ReviveF -= 0.05
	}
	if gs.Player.Pl.peaceT > 0 {
		peceRec := rl.NewRectangle((gs.Player.Pl.crec.X+gs.Player.Pl.crec.Width/2)-bsU/2, gs.Player.Pl.crec.Y-(bsU), bsU, bsU)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[46], peceRec, gs.Core.Ori, 0, rl.White)
	}
	//DRAW PLAYER IMG
	rl.DrawTexturePro(gs.Render.Imgs, gs.Player.Pl.img, shadowRec, gs.Player.Pl.ori, gs.Player.Pl.ro, rl.Fade(rl.Black, 0.7))
	if gs.Player.Pl.hppause > 0 {
		rl.DrawTexturePro(gs.Render.Imgs, gs.Player.Pl.img, drawRec, gs.Player.Pl.ori, gs.Player.Pl.ro, ranCol())
	} else {
		rl.DrawTexturePro(gs.Render.Imgs, gs.Player.Pl.img, drawRec, gs.Player.Pl.ori, gs.Player.Pl.ro, rl.White)
	}
	//ORBITAL
	if gs.Player.Mods.orbital {
		if gs.Player.Mods.orbitalN >= 1 {
			if gs.Level.RoomChanged {
				gs.Player.Pl.orbital1 = rl.NewVector2(gs.Player.Pl.cnt.X+bsU4, gs.Player.Pl.cnt.Y+bsU4)
			} else {
				gs.Player.Pl.angle = gs.Player.Pl.angle * (math.Pi / 180)
				newx := float32(math.Cos(float64(gs.Player.Pl.angle)))*(gs.Player.Pl.orbital1.X-gs.Player.Pl.cnt.X) - float32(math.Sin(float64(gs.Player.Pl.angle)))*(gs.Player.Pl.orbital1.Y-gs.Player.Pl.cnt.Y) + gs.Player.Pl.cnt.X
				newy := float32(math.Sin(float64(gs.Player.Pl.angle)))*(gs.Player.Pl.orbital1.X-gs.Player.Pl.cnt.X) + float32(math.Cos(float64(gs.Player.Pl.angle)))*(gs.Player.Pl.orbital1.Y-gs.Player.Pl.cnt.Y) + gs.Player.Pl.cnt.Y
				gs.Player.Pl.orbital1 = rl.NewVector2(newx, newy)
				gs.Player.Pl.angle += 4
				if getabs(gs.Player.Pl.orbital1.X-gs.Player.Pl.cnt.X) > bsU4 {
					if gs.Player.Pl.orbital1.X > gs.Player.Pl.cnt.X {
						gs.Player.Pl.orbital1.X -= bsU / 8
					} else {
						gs.Player.Pl.orbital1.X += bsU / 8
					}
				}
				if getabs(gs.Player.Pl.orbital1.Y-gs.Player.Pl.cnt.Y) > bsU4 {
					if gs.Player.Pl.orbital1.Y > gs.Player.Pl.cnt.Y {
						gs.Player.Pl.orbital1.Y -= bsU / 8
					} else {
						gs.Player.Pl.orbital1.Y += bsU / 8
					}
				}
			}

			siz := bsU + bsU/2
			gs.Player.Pl.orbrec1 = rl.NewRectangle(gs.Player.Pl.orbital1.X-siz/2, gs.Player.Pl.orbital1.Y-siz/2, siz, siz)
			rl.DrawTexturePro(gs.Render.Imgs, gs.Player.Pl.orbimg1, gs.Player.Pl.orbrec1, gs.Core.Ori, 0, rl.White)

			if gs.Core.Frames%3 == 0 {
				gs.Player.Pl.orbimg1.X += gs.Render.Orbitalanim.W
				if gs.Player.Pl.orbimg1.X > gs.Render.Orbitalanim.xl+gs.Render.Orbitalanim.frames*gs.Render.Orbitalanim.W {
					gs.Player.Pl.orbimg1.X = gs.Render.Orbitalanim.xl
				}
			}

		}
		if gs.Player.Mods.orbitalN == 2 {
			if gs.Level.RoomChanged {
				gs.Player.Pl.orbital2 = rl.NewVector2(gs.Player.Pl.cnt.X-bsU7, gs.Player.Pl.cnt.Y-bsU7)
			} else {
				gs.Player.Pl.angle = gs.Player.Pl.angle * (math.Pi / 180)
				newx := float32(math.Cos(float64(gs.Player.Pl.angle)))*(gs.Player.Pl.orbital2.X-gs.Player.Pl.cnt.X) - float32(math.Sin(float64(gs.Player.Pl.angle)))*(gs.Player.Pl.orbital2.Y-gs.Player.Pl.cnt.Y) + gs.Player.Pl.cnt.X
				newy := float32(math.Sin(float64(gs.Player.Pl.angle)))*(gs.Player.Pl.orbital2.X-gs.Player.Pl.cnt.X) + float32(math.Cos(float64(gs.Player.Pl.angle)))*(gs.Player.Pl.orbital2.Y-gs.Player.Pl.cnt.Y) + gs.Player.Pl.cnt.Y
				gs.Player.Pl.orbital2 = rl.NewVector2(newx, newy)
				gs.Player.Pl.angle += 6

				if getabs(gs.Player.Pl.orbital2.X-gs.Player.Pl.cnt.X) > bsU7 {
					if gs.Player.Pl.orbital2.X > gs.Player.Pl.cnt.X {
						gs.Player.Pl.orbital2.X -= bsU / 8
					} else {
						gs.Player.Pl.orbital2.X += bsU / 8
					}
				}
				if getabs(gs.Player.Pl.orbital2.Y-gs.Player.Pl.cnt.Y) > bsU7 {
					if gs.Player.Pl.orbital2.Y > gs.Player.Pl.cnt.Y {
						gs.Player.Pl.orbital2.Y -= bsU / 8
					} else {
						gs.Player.Pl.orbital2.Y += bsU / 8
					}
				}
			}

			siz := bsU + bsU/2
			gs.Player.Pl.orbrec2 = rl.NewRectangle(gs.Player.Pl.orbital2.X-siz/2, gs.Player.Pl.orbital2.Y-siz/2, siz, siz)
			rl.DrawTexturePro(gs.Render.Imgs, gs.Player.Pl.orbimg2, gs.Player.Pl.orbrec2, gs.Core.Ori, 0, rl.White)

			if gs.Core.Frames%3 == 0 {
				gs.Player.Pl.orbimg2.X += gs.Render.Orbitalanim.W
				if gs.Player.Pl.orbimg2.X > gs.Render.Orbitalanim.xl+gs.Render.Orbitalanim.frames*gs.Render.Orbitalanim.W {
					gs.Player.Pl.orbimg2.X = gs.Render.Orbitalanim.xl
				}
			}
		}

	}

	//COMPANIONS
	if gs.Player.Mods.planty || gs.Player.Mods.alien || gs.Player.Mods.carrot {
		drawUpCompanions()
	}

	//NIGHT REC
	if gs.Level.Night {
		rec := rl.NewRectangle(gs.Player.Pl.cnt.X-gs.Render.Etc[25].Width/2, gs.Player.Pl.cnt.Y-gs.Render.Etc[25].Height/2, gs.Render.Etc[25].Width, gs.Render.Etc[25].Height)
		rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[25], rec, gs.Core.Ori, 0, rl.White)

		recT := rl.NewRectangle(gs.Level.LevRec.X, gs.Level.LevRec.Y, gs.Level.LevRec.Width, rec.Y-gs.Level.LevRec.Y)
		recB := rl.NewRectangle(gs.Level.LevRec.X, rec.Y+rec.Height, gs.Level.LevRec.Width, (gs.Level.LevRec.Y+gs.Level.LevRec.Height)-(rec.Y+rec.Height))
		recL := rl.NewRectangle(gs.Level.LevRec.X, rec.Y, rec.X-gs.Level.LevRec.X, rec.Height)
		recR := rl.NewRectangle(rec.X+rec.Width, rec.Y, (gs.Level.LevRec.X+gs.Level.LevRec.Width)-(rec.X+rec.Width), rec.Height)

		rl.DrawRectangleRec(recT, rl.Fade(rl.Black, 0.6))
		rl.DrawRectangleRec(recB, rl.Fade(rl.Black, 0.6))
		rl.DrawRectangleRec(recL, rl.Fade(rl.Black, 0.6))
		rl.DrawRectangleRec(recR, rl.Fade(rl.Black, 0.6))

	}

	//DEBUG
	if gs.Core.Debug {
		rl.DrawRectangleLinesEx(gs.Player.Pl.arec, 1, rl.Red)
		rl.DrawRectangleLinesEx(gs.Player.Pl.crec, 1, rl.Blue)
		rl.DrawRectangleLinesEx(gs.Player.Pl.atkrec, 1, rl.White)
		rl.DrawPixelV(gs.Player.Pl.cnt, rl.Red)
	}

}

func drawUpCompanions() { //MARK:DRAW UP COMPANIONS

	if gs.Player.Mods.carrot {

		gs.Companions.MrCarrot.timer--
		if gs.Companions.MrCarrot.timer == 0 {
			gs.Companions.MrCarrot.timer = gs.Core.Fps * rI32(1, 5)
			makeProjectile("gs.Companions.MrCarrot")
		}
		//MOVE
		if checkNextMove(gs.Companions.MrCarrot.rec, gs.Companions.MrCarrot.velx, gs.Companions.MrCarrot.vely, false) {
			gs.Companions.MrCarrot.rec.X += gs.Companions.MrCarrot.velx
			gs.Companions.MrCarrot.rec.Y += gs.Companions.MrCarrot.vely
		} else {
			gs.Companions.MrCarrot.velx = rF32(-gs.Companions.MrCarrot.vel, gs.Companions.MrCarrot.vel)
			gs.Companions.MrCarrot.vely = rF32(-gs.Companions.MrCarrot.vel, gs.Companions.MrCarrot.vel)
		}

		//IMG

		shadowRec := gs.Companions.MrCarrot.rec
		shadowRec.X -= 5
		shadowRec.Y += 5
		if gs.Companions.MrCarrot.velx > 0 {
			rl.DrawTexturePro(gs.Render.Imgs, gs.Companions.MrCarrot.imgr, shadowRec, gs.Core.Ori, 0, rl.Fade(rl.Black, 0.8))
			rl.DrawTexturePro(gs.Render.Imgs, gs.Companions.MrCarrot.imgr, gs.Companions.MrCarrot.rec, gs.Core.Ori, 0, rl.White)
		} else {
			rl.DrawTexturePro(gs.Render.Imgs, gs.Companions.MrCarrot.imgl, shadowRec, gs.Core.Ori, 0, rl.Fade(rl.Black, 0.8))
			rl.DrawTexturePro(gs.Render.Imgs, gs.Companions.MrCarrot.imgl, gs.Companions.MrCarrot.rec, gs.Core.Ori, 0, rl.White)
		}

		//ANIM
		if gs.Core.Frames%6 == 0 {
			gs.Companions.MrCarrot.imgl.X += gs.Companions.MrCarrot.imgl.Width
			if gs.Companions.MrCarrot.imgl.X > gs.Companions.MrCarrot.imgl.Width*float32(gs.Companions.MrCarrot.frames) {
				gs.Companions.MrCarrot.imgl.X = 0
			}
			gs.Companions.MrCarrot.imgr.X += gs.Companions.MrCarrot.imgr.Width
			if gs.Companions.MrCarrot.imgr.X > 228+(gs.Companions.MrCarrot.imgr.Width*float32(gs.Companions.MrCarrot.frames)) {
				gs.Companions.MrCarrot.imgr.X = 228
			}
		}
	}

	if gs.Player.Mods.alien {
		gs.Companions.MrAlien.timer--
		if gs.Companions.MrAlien.timer == 0 {
			gs.Companions.MrAlien.timer = gs.Core.Fps * rI32(3, 8)
			makeProjectile("gs.Companions.MrAlien")
		}
		//MOVE
		if checkNextMove(gs.Companions.MrAlien.rec, gs.Companions.MrAlien.velx, gs.Companions.MrAlien.vely, false) {
			gs.Companions.MrAlien.rec.X += gs.Companions.MrAlien.velx
			gs.Companions.MrAlien.rec.Y += gs.Companions.MrAlien.vely
		} else {
			gs.Companions.MrAlien.velx = rF32(-gs.Companions.MrAlien.vel, gs.Companions.MrAlien.vel)
			gs.Companions.MrAlien.vely = rF32(-gs.Companions.MrAlien.vel, gs.Companions.MrAlien.vel)
		}

		//IMG
		shadowRec := gs.Companions.MrAlien.rec
		shadowRec.X -= 5
		shadowRec.Y += 5
		rl.DrawTexturePro(gs.Render.Imgs, gs.Companions.MrAlien.img, shadowRec, gs.Core.Ori, 0, rl.Fade(rl.Black, 0.8))
		rl.DrawTexturePro(gs.Render.Imgs, gs.Companions.MrAlien.img, gs.Companions.MrAlien.rec, gs.Core.Ori, 0, rl.White)

		if getabs(gs.Companions.MrAlien.velx) > getabs(gs.Companions.MrAlien.vely) {
			if gs.Companions.MrAlien.velx > 0 {
				gs.Companions.MrAlien.img.Y = 898
			} else {
				gs.Companions.MrAlien.img.Y = 834
			}
		} else {
			if gs.Companions.MrAlien.vely > 0 {
				gs.Companions.MrAlien.img.Y = 770
			} else {
				gs.Companions.MrAlien.img.Y = 962
			}
		}

		if gs.Core.Frames%3 == 0 {
			gs.Companions.MrAlien.img.X += gs.Companions.MrAlien.img.Width
			if gs.Companions.MrAlien.img.X >= 1200 {
				gs.Companions.MrAlien.img.X = 1008
			}
		}
	}

	if gs.Player.Mods.planty {
		//MOVE
		if checkNextMove(gs.Companions.MrPlanty.rec, gs.Companions.MrPlanty.velx, gs.Companions.MrPlanty.vely, false) {
			gs.Companions.MrPlanty.rec.X += gs.Companions.MrPlanty.velx
			gs.Companions.MrPlanty.rec.Y += gs.Companions.MrPlanty.vely
		} else {
			gs.Companions.MrPlanty.velx = rF32(-gs.Companions.MrPlanty.vel, gs.Companions.MrPlanty.vel)
			gs.Companions.MrPlanty.vely = rF32(-gs.Companions.MrPlanty.vel, gs.Companions.MrPlanty.vel)
		}
		gs.Companions.MrPlanty.cnt = rl.NewVector2(gs.Companions.MrPlanty.rec.Width/2, gs.Companions.MrPlanty.rec.Height/2)

		//IMG
		shadowRec := gs.Companions.MrPlanty.rec
		shadowRec.X -= 5
		shadowRec.Y += 5
		if gs.Companions.MrPlanty.velx > 0 {
			rl.DrawTexturePro(gs.Render.Imgs, gs.Companions.MrPlanty.imgr, shadowRec, gs.Core.Ori, 0, rl.Fade(rl.Black, 0.8))
			rl.DrawTexturePro(gs.Render.Imgs, gs.Companions.MrPlanty.imgr, gs.Companions.MrPlanty.rec, gs.Core.Ori, 0, rl.White)
		} else {
			rl.DrawTexturePro(gs.Render.Imgs, gs.Companions.MrPlanty.imgl, shadowRec, gs.Core.Ori, 0, rl.Fade(rl.Black, 0.8))
			rl.DrawTexturePro(gs.Render.Imgs, gs.Companions.MrPlanty.imgl, gs.Companions.MrPlanty.rec, gs.Core.Ori, 0, rl.White)
		}

		//ANIM
		if gs.Core.Frames%6 == 0 {
			gs.Companions.MrPlanty.imgl.X += gs.Companions.MrPlanty.imgl.Width
			if gs.Companions.MrPlanty.imgl.X > gs.Companions.MrPlanty.imgl.Width*float32(gs.Companions.MrPlanty.frames) {
				gs.Companions.MrPlanty.imgl.X = 0
			}
			gs.Companions.MrPlanty.imgr.X += gs.Companions.MrPlanty.imgr.Width
			if gs.Companions.MrPlanty.imgr.X > 352+(gs.Companions.MrPlanty.imgr.Width*float32(gs.Companions.MrPlanty.frames)) {
				gs.Companions.MrPlanty.imgr.X = 352
			}
		}

		//BULLETS
		if gs.Core.Frames%30 == 0 {
			makeProjectile("plantbull")
		}

	}

}

func drawnocamBG() { //MARK:DRAW NO CAM BACKGROUND

}

func drawnocam() { //MARK:DRAW NO CAM

	//INTRO
	if gs.UI.Intro && !gs.UI.OptionsOn {
		if gs.UI.IntroT1 > 0 {
			gs.Render.ShaderOn = false
			x := gs.Core.Cnt.X - gs.Render.Etc[56].Width/2
			y := gs.Core.Cnt.Y - gs.Render.Etc[56].Height/2
			rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[56], rl.NewRectangle(x, y, gs.Render.Etc[56].Width, gs.Render.Etc[56].Height), gs.Core.Ori, 0, rl.Fade(rl.White, gs.UI.IntroF1))
			if gs.UI.IntroF1 < 1 {
				gs.UI.IntroF1 += 0.01
			}

			if gs.UI.IntroF1 > 0.5 {
				txt := "raylib.com"
				txtlen := rl.MeasureText(txt, 20)
				txtx := int32(gs.Core.Cnt.X) - txtlen/2
				txty := int32(gs.Core.Cnt.Y + (gs.Render.Etc[56].Height / 2) + bsU/2)
				rl.DrawText(txt, txtx, txty, 20, rl.White)
			}
			gs.UI.IntroT1--
			if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) {
				gs.UI.IntroT1 = 0
			}
		} else if gs.UI.IntroT2 > 0 {
			gs.Render.ShaderOn = false
			x := gs.Core.Cnt.X - gs.Render.Etc[55].Width/2
			y := gs.Core.Cnt.Y - gs.Render.Etc[55].Height/2
			rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[55], rl.NewRectangle(x, y, gs.Render.Etc[55].Width, gs.Render.Etc[55].Height), gs.Core.Ori, 0, rl.Fade(rl.White, gs.UI.IntroF2))
			if gs.UI.IntroF2 < 1 {
				gs.UI.IntroF2 += 0.01
			}
			if gs.UI.IntroF2 > 0.5 {
				txt := "go.dev"
				txtlen := rl.MeasureText(txt, 20)
				txtx := int32(gs.Core.Cnt.X) - txtlen/2
				txty := int32(gs.Core.Cnt.Y + (gs.Render.Etc[55].Height / 2) + bsU/2)
				rl.DrawText(txt, txtx, txty, 20, rl.White)
			}
			gs.UI.IntroT2--
			if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) {
				gs.UI.IntroT2 = 0

			}
		} else {
			gs.Render.ShaderOn = true

			if gs.UI.IntroCount {

				gs.UI.IntroT3--
				if gs.UI.IntroT3 > gs.Core.Fps*2 {
					txt := "3"
					txtlen := rl.MeasureText(txt, 200)
					txty := int32(gs.Core.Cnt.Y) - 100
					txtx := int32(gs.Core.Cnt.X) - txtlen/2
					rl.DrawText(txt, txtx, txty, 200, rl.Green)
				} else if gs.UI.IntroT3 > gs.Core.Fps {
					txt := "2"
					txtlen := rl.MeasureText(txt, 200)
					txty := int32(gs.Core.Cnt.Y) - 100
					txtx := int32(gs.Core.Cnt.X) - txtlen/2
					rl.DrawText(txt, txtx, txty, 200, rl.Green)
				} else if gs.UI.IntroT3 > 0 {
					txt := "1"
					txtlen := rl.MeasureText(txt, 200)
					txty := int32(gs.Core.Cnt.Y) - 100
					txtx := int32(gs.Core.Cnt.X) - txtlen/2
					rl.DrawText(txt, txtx, txty, 200, rl.Green)
				} else {
					gs.UI.Intro = false
					gs.Core.Pause = false
					gs.Level.RunT = 0
					gs.Level.Mins = 0
					gs.Level.Secs = 0
					if gs.Audio.MusicOn {
						rl.PlayMusicStream(gs.Audio.Music)
					}
				}

			} else {
				x := gs.Core.Cnt.X - gs.Render.Etc[54].Width/2
				y := gs.Core.Cnt.Y - gs.Render.Etc[54].Height/2
				rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[54], rl.NewRectangle(x, y, gs.Render.Etc[54].Width, gs.Render.Etc[54].Height), gs.Core.Ori, 0, rl.Fade(rl.White, gs.UI.IntroF3))
				if gs.UI.IntroF3 < 1 {
					gs.UI.IntroF3 += 0.01
				} else {
					txt := "press space or a button to start"
					txtlen := rl.MeasureText(txt, 20)
					txty := int32(y+gs.Render.Etc[54].Height) + bsU2i32
					txtx := int32(gs.Core.Cnt.X) - txtlen/2
					rl.DrawText(txt, txtx, txty, 20, rl.Green)

					txt = "press esc or menu button for options"
					txtlen = rl.MeasureText(txt, 20)
					txty += 30
					txtx = int32(gs.Core.Cnt.X) - txtlen/2
					rl.DrawText(txt, txtx, txty, 20, rl.Green)
				}
				if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) {
					gs.UI.IntroCount = true
					gs.Player.StartdmgT = gs.Core.Fps * 7
					rl.PlaySound(gs.Audio.Sfx[13])
				}

				if gs.UI.IntroF3 > 0.5 {
					txt := "unklnik.com"
					txtlen := rl.MeasureText(txt, 40)
					txtx := int32(gs.Core.Cnt.X) - txtlen/2
					txty := int32(gs.Core.ScrH) - 50
					rl.DrawText(txt, txtx, txty, 40, rl.Green)
				}
			}

		}

	}

	//FLOOD
	if gs.Player.Mods.flood {
		rl.DrawRectangleRec(gs.FX.FloodRec, rl.Fade(rl.SkyBlue, rF32(0.07, 0.12)))

		//WATER ANIM
		siz := bsU2
		x := gs.FX.FloodRec.X
		y := gs.FX.FloodRec.Y - siz/2
		for {
			rec := rl.NewRectangle(x, y, siz, siz)
			rl.DrawTexturePro(gs.Render.Imgs, gs.FX.FloodImg, rec, gs.Core.Ori, 0, rl.SkyBlue)
			x += siz
			if x >= gs.Core.ScrWF32 {
				break
			}
		}
		if gs.Core.Frames%3 == 0 {
			gs.FX.FloodImg.X += gs.Render.Floodanim.W
			if gs.FX.FloodImg.X > gs.Render.Floodanim.xl+gs.Render.Floodanim.frames*gs.Render.Floodanim.W {
				gs.FX.FloodImg.X = gs.Render.Floodanim.xl
			}
		}

		//FISH
		rec := rl.NewRectangle(gs.FX.FishV2.X-gs.FX.FishSiz/2, gs.FX.FishV2.Y-gs.FX.FishSiz/2, gs.FX.FishSiz, gs.FX.FishSiz)
		rl.DrawTexturePro(gs.Render.Imgs, gs.FX.Fish1, rec, gs.Core.Ori, 0, rl.SkyBlue)
		rec = rl.NewRectangle(gs.FX.Fish2V2.X-gs.FX.FishSiz2/2, gs.FX.Fish2V2.Y-gs.FX.FishSiz2/2, gs.FX.FishSiz2, gs.FX.FishSiz2)
		rl.DrawTexturePro(gs.Render.Imgs, gs.FX.Fish2, rec, gs.Core.Ori, 0, rl.SkyBlue)

		if gs.Core.Frames%10 == 0 {
			if gs.FX.FishLR {
				gs.FX.Fish1.X += gs.Render.FishL.W
				if gs.FX.Fish1.X > gs.Render.FishL.xl+gs.Render.FishL.frames*gs.Render.FishL.W {
					gs.FX.Fish1.X = gs.Render.FishL.xl
				}
			} else {
				gs.FX.Fish1.X += gs.Render.FishR.W
				if gs.FX.Fish1.X > gs.Render.FishR.xl+gs.Render.FishR.frames*gs.Render.FishR.W {
					gs.FX.Fish1.X = gs.Render.FishR.xl
				}
			}
			if gs.FX.Fish2LR {
				gs.FX.Fish2.X += gs.Render.FishR.W
				if gs.FX.Fish2.X > gs.Render.FishR.xl+gs.Render.FishR.frames*gs.Render.FishR.W {
					gs.FX.Fish2.X = gs.Render.FishR.xl
				}
			} else {
				gs.FX.Fish2.X += gs.Render.FishL.W
				if gs.FX.Fish2.X > gs.Render.FishL.xl+gs.Render.FishL.frames*gs.Render.FishL.W {
					gs.FX.Fish2.X = gs.Render.FishL.xl
				}
			}

		}

	}

	//RAIN
	if gs.Player.Mods.umbrella {
		for a := 0; a < len(gs.FX.Rain); a++ {
			rl.DrawRectangleRec(gs.FX.Rain[a], rl.Fade(ranCyan(), rF32(0.4, 0.7)))
			gs.FX.Rain[a].Y += 8
			if gs.FX.Rain[a].Y > gs.Core.ScrHF32 {
				gs.FX.Rain[a].Y = rF32(-gs.Core.ScrHF32, -bsU)
			}
		}

	}

	//TELEPORT
	if gs.Player.TeleportOn {
		for a := 0; a < len(gs.Player.TeleportRadius); a++ {
			rl.DrawCircleLines(int32(gs.Core.Cnt.X), int32(gs.Core.Cnt.Y), gs.Player.TeleportRadius[a], ranCol())
			gs.Player.TeleportRadius[a] -= bsU
			if gs.Player.TeleportRadius[a] <= 0 {
				gs.Player.TeleportOn = false
				gs.Player.Pl.cnt = rl.NewVector2(gs.Level.LevRecInner.X+gs.Level.LevRecInner.Width/2, gs.Level.LevRecInner.Y+bsU3)
				upPlayerRec()
				for i := 0; i < len(gs.Level.Level[gs.Player.TeleportRoomNum].innerBloks); i++ {
					if rl.CheckCollisionRecs(gs.Player.Pl.crec, gs.Level.Level[gs.Player.TeleportRoomNum].innerBloks[i].rec) {
						gs.Player.Pl.rec.X = gs.Level.Level[gs.Player.TeleportRoomNum].innerBloks[i].rec.X + gs.Level.Level[gs.Player.TeleportRoomNum].innerBloks[i].rec.Width + bsU/4
						upPlayerRec()
						break
					}
				}

				cntCompanion := gs.Player.Pl.cnt
				if gs.Player.Mods.carrot {
					gs.Companions.MrCarrot.rec = rl.NewRectangle(cntCompanion.X-gs.Companions.MrCarrot.rec.Width/2, cntCompanion.Y-gs.Companions.MrCarrot.rec.Width/2, gs.Companions.MrCarrot.rec.Width, gs.Companions.MrCarrot.rec.Width)
				}
				if gs.Player.Mods.alien {
					gs.Companions.MrAlien.rec = rl.NewRectangle(cntCompanion.X-gs.Companions.MrAlien.rec.Width/2, cntCompanion.Y-gs.Companions.MrAlien.rec.Width/2, gs.Companions.MrAlien.rec.Width, gs.Companions.MrAlien.rec.Width)
				}
				if gs.Player.Mods.planty {
					gs.Companions.MrPlanty.rec = rl.NewRectangle(cntCompanion.X-gs.Companions.MrPlanty.rec.Width/2, cntCompanion.Y-gs.Companions.MrPlanty.rec.Width/2, gs.Companions.MrPlanty.rec.Width, gs.Companions.MrPlanty.rec.Width)
				}

				gs.Level.RoomNum = gs.Player.TeleportRoomNum
				break
			}
		}
	}

	//SCANLINES
	if gs.UI.ScanLinesOn {
		for a := 0; a < len(gs.FX.ScanlineV2); a++ {
			v2 := gs.FX.ScanlineV2[a]
			v2.X += gs.Core.ScrWF32
			rl.DrawLineEx(gs.FX.ScanlineV2[a], v2, 1, rl.Fade(rl.Black, 0.5))
			gs.FX.ScanlineV2[a].Y++
			if gs.FX.ScanlineV2[a].Y > gs.Core.ScrHF32+2 {
				gs.FX.ScanlineV2[a].Y = 0
			}
		}
	}

	//MARK: DRAW MAP
	if gs.Level.LevMapOn {
		txt := "level " + fmt.Sprint(gs.Level.Levelnum)
		txtlen := rl.MeasureText(txt, txU4)
		txtx := int32(gs.Core.Cnt.X) - txtlen/2
		txty := txU
		rl.DrawText(txt, txtx, txty, txU4, rl.White)

		for a := 0; a < len(gs.Level.LevMap); a++ {
			if gs.Core.Debug {
				rl.DrawText(fmt.Sprint(a), gs.Level.LevMap[a].ToInt32().X+4, gs.Level.LevMap[a].ToInt32().Y+4, txU, rl.White)
			}

			if gs.Level.ShopRoomNum == a && gs.Level.RoomNum == a {
				rec := gs.Level.LevMap[a]
				rec.Width = rec.Width / 2
				rl.DrawRectangleRec(rec, rl.Fade(rl.Green, 0.2))
				rec.X += rec.Width
				rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
			} else if gs.Level.ExitRoomNum == a && gs.Level.RoomNum == a {
				rec := gs.Level.LevMap[a]
				rec.Width = rec.Width / 2
				rl.DrawRectangleRec(rec, rl.Fade(rl.Green, 0.2))
				rec.X += rec.Width
				rl.DrawRectangleRec(rec, rl.Fade(rl.Yellow, 0.2))
			} else if gs.Level.ExitRoomNum == a && gs.Player.Mods.exitmap || gs.Level.ExitRoomNum == a && gs.Level.Level[a].visited {
				rl.DrawRectangleRec(gs.Level.LevMap[a], rl.Fade(rl.Yellow, 0.2))
			} else if gs.Level.ShopRoomNum == a {
				rl.DrawRectangleRec(gs.Level.LevMap[a], rl.Fade(rl.Magenta, 0.2))
			} else if gs.Level.RoomNum == a {
				rl.DrawRectangleRec(gs.Level.LevMap[a], rl.Fade(rl.Green, 0.2))
			} else {
				if gs.Level.Level[a].visited {
					rl.DrawRectangleRec(gs.Level.LevMap[a], rl.Fade(rl.Blue, 0.2))
				} else {
					rl.DrawRectangleRec(gs.Level.LevMap[a], rl.Fade(rl.Red, 0.2))
				}
			}
			rl.DrawRectangleLinesEx(gs.Level.LevMap[a], 1, rl.Black)
		}

		txt = "player visited shop exit"
		txtlen = rl.MeasureText(txt, txU4)
		txtx = int32(gs.Core.Cnt.X) - txtlen/2
		txty = gs.Core.ScrH32 - txU5
		txtlen = rl.MeasureText("player ", txU4)
		rl.DrawText("player ", txtx, txty, txU4, rl.Green)
		txtx += txtlen
		txtlen = rl.MeasureText("visited ", txU4)
		rl.DrawText("visited ", txtx, txty, txU4, rl.Blue)
		txtx += txtlen
		txtlen = rl.MeasureText("shop ", txU4)
		rl.DrawText("shop ", txtx, txty, txU4, rl.Magenta)
		txtx += txtlen
		txtlen = rl.MeasureText("exit ", txU4)
		rl.DrawText("exit ", txtx, txty, txU4, rl.Yellow)

	}

}
func drawnoRender() { //MARK:DRAW NO RENDER

	if gs.Player.InvenOn {
		rl.BeginMode2D(gs.Render.Cam2)

		drawInvenDetail()

		rl.EndMode2D()

	}

	if gs.Core.Debug {
		drawDebug()
	}

}
func drawDebug() { //MARK:DRAW DEBUG

	siderec := rl.NewRectangle(0, 0, 300, gs.Core.ScrHF32)
	rl.DrawRectangleRec(siderec, rl.Fade(darkRed(), 0.3))

	txtX, txtY := txU, txU

	rl.DrawText("pl.cnt.X"+" "+fmt.Sprint(gs.Player.Pl.cnt.X), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("pl.cnt.Y"+" "+fmt.Sprint(gs.Player.Pl.cnt.Y), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("levBorderBlokNum"+" "+fmt.Sprint(gs.Level.LevBorderBlokNum), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("len FX"+" "+fmt.Sprint(len(gs.FX.Fx)), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("axeT2"+" "+fmt.Sprint(gs.Player.Mods.axeT2), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("santaT"+" "+fmt.Sprint(gs.Player.Mods.santaT), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("snowOn"+" "+fmt.Sprint(gs.Player.Mods.snowon), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("len(enProj)"+" "+fmt.Sprint(len(gs.Enemies.EnProj)), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("len(plProj)"+" "+fmt.Sprint(len(gs.Player.PlProj)), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("mods.fireballN"+" "+fmt.Sprint(gs.Player.Mods.fireballN), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("mods.bounceN"+" "+fmt.Sprint(gs.Player.Mods.bounceN), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("pl.atkDMG"+" "+fmt.Sprint(gs.Player.Pl.atkDMG), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("mods.orbitalN"+" "+fmt.Sprint(gs.Player.Mods.orbitalN), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("mods.alien"+" "+fmt.Sprint(gs.Player.Mods.alien), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("mods.carrot"+" "+fmt.Sprint(gs.Player.Mods.carrot), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("cam2.Zoom"+" "+fmt.Sprint(gs.Render.Cam2.Zoom), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("gs.Input.IsController"+" "+fmt.Sprint(gs.Input.IsController), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("gs.Input.UseController"+" "+fmt.Sprint(gs.Input.UseController), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("GamepadAxisMovement 0"+" "+fmt.Sprint(rl.GetGamepadAxisMovement(0, 0)), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("GamepadAxisMovement 1"+" "+fmt.Sprint(rl.GetGamepadAxisMovement(0, 1)), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("bosses[bossnum].timer"+" "+fmt.Sprint(gs.Level.Bosses[gs.Level.Bossnum].timer), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("gs.Shop.ShopItems[0].shopoff"+" "+fmt.Sprint(gs.Shop.ShopItems[0].shopoff), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("floodRec.Y"+" "+fmt.Sprint(gs.FX.FloodRec.Y), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("pl.crec.Y"+" "+fmt.Sprint(gs.Player.Pl.crec.Y), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("minsEND"+" "+fmt.Sprint(gs.Level.MinsEND), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("secsEND"+" "+fmt.Sprint(gs.Level.SecsEND), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("len(gs.Mario.MarioCoins)"+" "+fmt.Sprint(len(gs.Mario.MarioCoins)), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("hardcore"+" "+fmt.Sprint(gs.Level.Hardcore), txtX, txtY, txU, rl.White)
	txtY += txU

}
func drawBlokDrec(blok xblok, shadow, blur bool, blurDist float32) { //MARK:DRAW BLOK DREC

	if shadow {
		shadowRec := blok.drec
		shadowRec.X -= 5
		shadowRec.Y += 5
		rl.DrawTexturePro(gs.Render.Imgs, blok.img, shadowRec, rl.NewVector2(blok.drec.Width/2, blok.drec.Height/2), blok.ro, rl.Fade(rl.Black, 0.8))
	}

	rl.DrawTexturePro(gs.Render.Imgs, blok.img, blok.drec, rl.NewVector2(blok.drec.Width/2, blok.drec.Height/2), blok.ro, rl.Fade(blok.color, blok.fade))
	if blur {
		blurRec := blok.drec
		blurRec.X -= rF32(-blurDist, blurDist)
		blurRec.Y -= rF32(-blurDist, blurDist)
		rl.DrawTexturePro(gs.Render.Imgs, blok.img, blurRec, rl.NewVector2(blok.drec.Width/2, blok.drec.Height/2), blok.ro, rl.Fade(blok.color, rF32(0.05, 0.2)))
	}

	if gs.Core.Debug {
		rl.DrawRectangleLinesEx(blok.rec, 0.5, rl.Green)
		rl.DrawRectangleLinesEx(blok.crec, 1, rl.Blue)
		rl.DrawRectangleLinesEx(blok.crec2, 2, rl.Red)
		rl.DrawCircleV(blok.cnt, 4, rl.Red)
		rl.DrawText(blok.name, blok.rec.ToInt32().X, blok.rec.ToInt32().Y-10, 10, rl.White)
		rl.DrawText(fmt.Sprint(blok.timer), blok.rec.ToInt32().X, blok.rec.ToInt32().Y+blok.rec.ToInt32().Height, 10, rl.White)
	}

}
func drawBlok(blok xblok, shadow, blur bool, blurDist float32) { //MARK:DRAW BLOK

	if shadow {
		shadowRec := blok.rec
		shadowRec.X -= 5
		shadowRec.Y += 5
		rl.DrawTexturePro(gs.Render.Imgs, blok.img, shadowRec, rl.NewVector2(0, 0), blok.ro, rl.Fade(rl.Black, 0.8))
	}

	//CANDLE LIGHT
	if blok.name == "candle" {
		radius := rF32(bsU3, bsU5)
		rl.DrawCircleGradient(int32(blok.cnt.X), int32(blok.cnt.Y), radius, rl.Fade(ranYellow(), rF32(0.05, 0.3)), rl.Blank)
	}

	//DRAW BLOK IMG
	rl.DrawTexturePro(gs.Render.Imgs, blok.img, blok.rec, rl.NewVector2(0, 0), blok.ro, rl.Fade(blok.color, blok.fade))
	if blur {
		if blok.name == "skull" || blok.name == "candle" { //HALF BLUR
			blurRec := blok.rec
			blurRec.X -= rF32(-blurDist/2, blurDist/2)
			blurRec.Y -= rF32(-blurDist/2, blurDist/2)
			rl.DrawTexturePro(gs.Render.Imgs, blok.img, blurRec, rl.NewVector2(0, 0), blok.ro, rl.Fade(blok.color, rF32(0.05, 0.2)))
		} else { //FULL BLUR
			blurRec := blok.rec
			blurRec.X -= rF32(-blurDist, blurDist)
			blurRec.Y -= rF32(-blurDist, blurDist)
			rl.DrawTexturePro(gs.Render.Imgs, blok.img, blurRec, rl.NewVector2(0, 0), blok.ro, rl.Fade(blok.color, rF32(0.05, 0.2)))
		}
	}

	if gs.Core.Debug {
		rl.DrawRectangleLinesEx(blok.rec, 0.5, rl.Green)
		rl.DrawRectangleLinesEx(blok.crec, 2, rl.Blue)
		rl.DrawRectangleLinesEx(blok.crec2, 2, rl.Red)
		rl.DrawText(blok.name, blok.rec.ToInt32().X, blok.rec.ToInt32().Y-10, 10, rl.White)
	}

}
func drawUpEnProj() { //MARK:DRAW UP ENEMY PROJECTILES

	clear := false
	for a := 0; a < len(gs.Enemies.EnProj); a++ {

		if gs.Enemies.EnProj[a].onoff {
			//	rl.DrawRectangleRec(gs.Enemies.EnProj[a].rec, gs.Enemies.EnProj[a].col)

			if gs.Enemies.EnProj[a].name == "mushbull" {
				shadowrec := makeDrec(gs.Enemies.EnProj[a].rec)
				shadowrec.X -= 5
				shadowrec.Y += 5
				rl.DrawTexturePro(gs.Render.Imgs, gs.Enemies.EnProj[a].img, shadowrec, origin(gs.Enemies.EnProj[a].rec), gs.Enemies.EnProj[a].ro, rl.Fade(rl.Black, 0.7))
				if flipcoin() {
					rl.DrawTexturePro(gs.Render.Imgs, gs.Enemies.EnProj[a].img, makeDrec(gs.Enemies.EnProj[a].rec), origin(gs.Enemies.EnProj[a].rec), gs.Enemies.EnProj[a].ro, ranCyan())
				} else {
					rl.DrawTexturePro(gs.Render.Imgs, gs.Enemies.EnProj[a].img, makeDrec(gs.Enemies.EnProj[a].rec), origin(gs.Enemies.EnProj[a].rec), gs.Enemies.EnProj[a].ro, ranRed())
				}
			} else if gs.Enemies.EnProj[a].name == "boss3" {
				rl.DrawTexturePro(gs.Render.Imgs, gs.Enemies.EnProj[a].img, makeDrec(gs.Enemies.EnProj[a].rec), gs.Enemies.EnProj[a].ori, gs.Enemies.EnProj[a].ro, gs.Enemies.EnProj[a].col)

			} else if gs.Enemies.EnProj[a].name == "ninja" {

				rl.DrawTexturePro(gs.Render.Imgs, gs.Enemies.EnProj[a].img, makeDrec(gs.Enemies.EnProj[a].rec), rl.NewVector2(gs.Enemies.EnProj[a].rec.Width/2, gs.Enemies.EnProj[a].rec.Height/2), gs.Enemies.EnProj[a].ro, gs.Enemies.EnProj[a].col)

			} else {
				rl.DrawTexturePro(gs.Render.Imgs, gs.Enemies.EnProj[a].img, gs.Enemies.EnProj[a].rec, gs.Core.Ori, gs.Enemies.EnProj[a].ro, gs.Enemies.EnProj[a].col)
				if gs.Enemies.EnProj[a].name == "boss2" {
					rl.DrawTexturePro(gs.Render.Imgs, gs.Enemies.EnProj[a].img, BlurRec(gs.Enemies.EnProj[a].rec, 7), gs.Core.Ori, gs.Enemies.EnProj[a].ro, rl.Fade(gs.Enemies.EnProj[a].col, 0.5))
				}
			}

			switch gs.Enemies.EnProj[a].name {
			case "boss3":
				gs.Enemies.EnProj[a].ro += 12
				if gs.Core.Frames%4 == 0 {
					if gs.Enemies.EnProj[a].rec.Width < bsU8 {
						gs.Enemies.EnProj[a].rec.X -= 4
						gs.Enemies.EnProj[a].rec.Y -= 4
						gs.Enemies.EnProj[a].rec.Width += 8
						gs.Enemies.EnProj[a].rec.Height += 8
						gs.Enemies.EnProj[a].ori = rl.NewVector2(gs.Enemies.EnProj[a].rec.Width/2, gs.Enemies.EnProj[a].rec.Height/2)
					}
				}
				if roll18() == 18 {
					gs.Enemies.EnProj[a].velx, gs.Enemies.EnProj[a].vely = moveFollow(gs.Enemies.EnProj[a].cnt, gs.Player.Pl.cnt, gs.Enemies.EnProj[a].vel)
				}
			case "boss2":
				if gs.Core.Frames%4 == 0 {
					gs.Enemies.EnProj[a].img.X += gs.Enemies.EnProj[a].img.Width
					if gs.Enemies.EnProj[a].img.X > gs.Render.Boss2anim.xl+(float32(gs.Render.Boss2anim.frames)*gs.Enemies.EnProj[a].img.Width) {
						gs.Enemies.EnProj[a].img.X = gs.Render.Boss2anim.xl
					}
				}
			case "boss1":
				if roll12() == 12 {
					gs.Enemies.EnProj[a].velx, gs.Enemies.EnProj[a].vely = moveFollow(gs.Enemies.EnProj[a].cnt, gs.Player.Pl.cnt, gs.Enemies.EnProj[a].vel)
				}

				if gs.Core.Frames%3 == 0 {
					gs.Enemies.EnProj[a].img.X += gs.Enemies.EnProj[a].img.Width
					if gs.Enemies.EnProj[a].img.X > gs.Render.Boss1anim.xl+(float32(gs.Render.Boss1anim.frames)*gs.Enemies.EnProj[a].img.Width) {
						gs.Enemies.EnProj[a].img.X = gs.Render.Boss1anim.xl
					}
				}
			case "mushbull":
				if gs.Core.Frames%2 == 0 {
					gs.Enemies.EnProj[a].img.X += gs.Enemies.EnProj[a].img.Width
					if gs.Enemies.EnProj[a].img.X > gs.Render.MushBull.xl+(float32(gs.Render.MushBull.frames)*gs.Enemies.EnProj[a].img.Width) {
						gs.Enemies.EnProj[a].img.X = gs.Render.MushBull.xl
					}
				}

				if roll18() == 18 {
					gs.Enemies.EnProj[a].velx, gs.Enemies.EnProj[a].vely = moveFollow(gs.Enemies.EnProj[a].cnt, gs.Player.Pl.cnt, gs.Enemies.EnProj[a].vel)
				}
			case "ninja":
				gs.Enemies.EnProj[a].ro += bsU / 2
			}

			if gs.Core.Debug {
				rl.DrawRectangleLinesEx(gs.Enemies.EnProj[a].rec, 0.5, rl.Red)
				rl.DrawRectangleLinesEx(gs.Enemies.EnProj[a].crec, 0.5, rl.White)
			}

			if checkNextMove(gs.Enemies.EnProj[a].rec, gs.Enemies.EnProj[a].velx, gs.Enemies.EnProj[a].vely, true) {
				gs.Enemies.EnProj[a].cnt.X += gs.Enemies.EnProj[a].velx
				gs.Enemies.EnProj[a].cnt.Y += gs.Enemies.EnProj[a].vely
				gs.Enemies.EnProj[a].rec = rl.NewRectangle(gs.Enemies.EnProj[a].cnt.X-gs.Enemies.EnProj[a].rec.Width/2, gs.Enemies.EnProj[a].cnt.Y-gs.Enemies.EnProj[a].rec.Height/2, gs.Enemies.EnProj[a].rec.Width, gs.Enemies.EnProj[a].rec.Height)
				switch gs.Enemies.EnProj[a].name {
				case "boss1":
					gs.Enemies.EnProj[a].crec = gs.Enemies.EnProj[a].rec
					gs.Enemies.EnProj[a].crec.X += gs.Enemies.EnProj[a].rec.Width / 4
					gs.Enemies.EnProj[a].crec.Y += gs.Enemies.EnProj[a].rec.Height / 4
					gs.Enemies.EnProj[a].crec.Width = gs.Enemies.EnProj[a].rec.Width / 2
					gs.Enemies.EnProj[a].crec.Height = gs.Enemies.EnProj[a].rec.Height / 2
				}
			} else {
				gs.Enemies.EnProj[a].onoff = false
			}

			if !rl.CheckCollisionPointRec(gs.Enemies.EnProj[a].cnt, gs.Level.LevRecInner) {
				gs.Enemies.EnProj[a].onoff = false
			}

			if gs.Enemies.EnProj[a].name == "boss1" {
				if rl.CheckCollisionRecs(gs.Enemies.EnProj[a].crec, gs.Player.Pl.crec) {
					hitPL(a, 1)
				}
			} else {
				if rl.CheckCollisionRecs(gs.Enemies.EnProj[a].rec, gs.Player.Pl.crec) {
					hitPL(a, 1)
				}
			}

		} else {
			clear = true
		}
	}

	if clear {
		for a := 0; a < len(gs.Enemies.EnProj); a++ {
			if !gs.Enemies.EnProj[a].onoff {
				gs.Enemies.EnProj = remProj(gs.Enemies.EnProj, a)
			}
		}
	}

}
func drawUpAirStrike() { //MARK:DRAW UP AIR STRIKE

	for a := 0; a < len(gs.FX.AirstrikeV2); a++ {
		siz := bsU2
		rec := rl.NewRectangle(gs.FX.AirstrikeV2[a].X-siz/2, gs.FX.AirstrikeV2[a].Y-siz/2, siz, siz)
		switch gs.FX.AirstrikeDir {
		case 1:
			rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[48], makeDrec(rec), origin(rec), 180, rl.White)
			gs.FX.AirstrikeV2[a].Y += 7
			if a == 1 {
				if gs.FX.AirstrikeV2[a].Y > gs.Level.LevRec.Y+gs.Level.LevRec.Width {
					gs.FX.AirstrikeOn = false
				}
			}
		case 2:
			rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[48], makeDrec(rec), origin(rec), 270, rl.White)
			gs.FX.AirstrikeV2[a].X -= 7
			if a == 1 {
				if gs.FX.AirstrikeV2[a].X < gs.Level.LevRec.X-rec.Width {
					gs.FX.AirstrikeOn = false
				}
			}
		case 3:
			rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[48], makeDrec(rec), origin(rec), 0, rl.White)
			gs.FX.AirstrikeV2[a].Y -= 7
			if a == 1 {
				if gs.FX.AirstrikeV2[a].Y < gs.Level.LevRec.Y-rec.Width {
					gs.FX.AirstrikeOn = false
				}
			}
		case 4:
			rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[48], makeDrec(rec), origin(rec), 90, rl.White)
			gs.FX.AirstrikeV2[a].X += 7
			if a == 1 {
				if gs.FX.AirstrikeV2[a].X > gs.Level.LevRec.X+gs.Level.LevRec.Width {
					gs.FX.AirstrikeOn = false
				}
			}
		}
	}

	gs.FX.AirstrikebombT--
	if gs.FX.AirstrikebombT <= 0 && rl.CheckCollisionPointRec(gs.FX.AirstrikeV2[0], gs.Level.LevRecInner) {

		siz := bsU8

		switch gs.FX.AirstrikeDir {
		case 1:
			zblok := makeBlokGenNoRecNoCntr()
			zblok.cnt = gs.FX.AirstrikeV2[0]
			zblok.cnt.X += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.name = "airbomb"
			zblok.img = gs.Render.Airstrikeanim.recTL
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
			zblok.cnt = gs.FX.AirstrikeV2[1]
			zblok.cnt.X += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
			zblok.cnt = gs.FX.AirstrikeV2[2]
			zblok.cnt.X += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
		case 2:
			zblok := makeBlokGenNoRecNoCntr()
			zblok.cnt = gs.FX.AirstrikeV2[0]
			zblok.cnt.Y += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.name = "airbomb"
			zblok.img = gs.Render.Airstrikeanim.recTL
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
			zblok.cnt = gs.FX.AirstrikeV2[1]
			zblok.cnt.Y += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
			zblok.cnt = gs.FX.AirstrikeV2[2]
			zblok.cnt.Y += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
		case 3:
			zblok := makeBlokGenNoRecNoCntr()
			zblok.cnt = gs.FX.AirstrikeV2[0]
			zblok.cnt.X += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.name = "airbomb"
			zblok.img = gs.Render.Airstrikeanim.recTL
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
			zblok.cnt = gs.FX.AirstrikeV2[1]
			zblok.cnt.X += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
			zblok.cnt = gs.FX.AirstrikeV2[2]
			zblok.cnt.X += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
		case 4:
			zblok := makeBlokGenNoRecNoCntr()
			zblok.cnt = gs.FX.AirstrikeV2[0]
			zblok.cnt.Y += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.name = "airbomb"
			zblok.img = gs.Render.Airstrikeanim.recTL
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
			zblok.cnt = gs.FX.AirstrikeV2[1]
			zblok.cnt.Y += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
			zblok.cnt = gs.FX.AirstrikeV2[2]
			zblok.cnt.Y += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
		}

		gs.FX.AirstrikebombT = rI32(int(gs.Core.Fps/4), int(gs.Core.Fps*2))
	}

}
func drawUpPlayerProj() { //MARK:DRAW UP PLAYER PROJECTILES

	//MOVE DRAW
	clear := false
	for a := 0; a < len(gs.Player.PlProj); a++ {
		if gs.Player.PlProj[a].onoff {
			shadowRec := gs.Player.PlProj[a].drec
			shadowRec.X -= 5
			shadowRec.Y += 5
			rl.DrawTexturePro(gs.Render.Imgs, gs.Player.PlProj[a].img, shadowRec, gs.Player.PlProj[a].ori, gs.Player.PlProj[a].ro, rl.Fade(rl.Black, 0.8))
			rl.DrawTexturePro(gs.Render.Imgs, gs.Player.PlProj[a].img, gs.Player.PlProj[a].drec, gs.Player.PlProj[a].ori, gs.Player.PlProj[a].ro, rl.Fade(gs.Player.PlProj[a].col, gs.Player.PlProj[a].fade))

			if gs.Core.Debug {
				rl.DrawRectangleLinesEx(gs.Player.PlProj[a].rec, 0.5, rl.Magenta)
			}

			gs.Player.PlProj[a].cnt.X += gs.Player.PlProj[a].velx
			gs.Player.PlProj[a].cnt.Y += gs.Player.PlProj[a].vely
			gs.Player.PlProj[a].rec = rl.NewRectangle(gs.Player.PlProj[a].cnt.X-gs.Player.PlProj[a].rec.Width/2, gs.Player.PlProj[a].cnt.Y-gs.Player.PlProj[a].rec.Height/2, gs.Player.PlProj[a].rec.Width, gs.Player.PlProj[a].rec.Height)
			gs.Player.PlProj[a].drec = gs.Player.PlProj[a].rec
			gs.Player.PlProj[a].drec.X += gs.Player.PlProj[a].rec.Width / 2
			gs.Player.PlProj[a].drec.Y += gs.Player.PlProj[a].rec.Height / 2

			switch gs.Player.PlProj[a].name {
			case "gs.Companions.MrCarrot":
				gs.Player.PlProj[a].ro += 8
			case "plantbull":
				if gs.Core.Frames%4 == 0 {
					gs.Player.PlProj[a].img.X += gs.Render.PlantBull.recTL.Width
					if gs.Player.PlProj[a].img.X >= gs.Render.PlantBull.xl+(gs.Render.PlantBull.recTL.Width*gs.Render.PlantBull.frames) {
						gs.Player.PlProj[a].img.X = gs.Render.PlantBull.xl
					}
				}
			case "fireball":
				if gs.Core.Frames%3 == 0 {
					gs.Player.PlProj[a].img.X += 17
					if gs.Player.PlProj[a].img.X >= gs.Render.FireballPlayer.xl+(17*gs.Render.FireballPlayer.frames) {
						gs.Player.PlProj[a].img.X = gs.Render.FireballPlayer.xl
					}
				}
			}

			//COLLISION ETC
			if len(gs.Level.Level[gs.Level.RoomNum].etc) > 0 {
				for b := 0; b < len(gs.Level.Level[gs.Level.RoomNum].etc); b++ {
					if gs.Level.Level[gs.Level.RoomNum].etc[b].onoff {
						if rl.CheckCollisionRecs(gs.Level.Level[gs.Level.RoomNum].etc[b].rec, gs.Player.PlProj[a].rec) {
							switch gs.Level.Level[gs.Level.RoomNum].etc[b].name {
							case "powerupBlok":
								destroyPowerupBlok(b)
								makeFX(3, gs.Level.Level[gs.Level.RoomNum].etc[b].cnt)
								gs.Level.Level[gs.Level.RoomNum].etc[b].onoff = false
								rl.PlaySound(gs.Audio.Sfx[6])
							case "oilbarrel":
								makeFX(4, gs.Level.Level[gs.Level.RoomNum].etc[b].cnt)
								gs.Level.Level[gs.Level.RoomNum].etc[b].onoff = false
								rl.PlaySound(gs.Audio.Sfx[10])
								rl.PlaySound(gs.Audio.Sfx[5])
							}
						}
					}
				}
			}

			//COLLISION BOUNDARY
			if !rl.CheckCollisionPointRec(gs.Player.PlProj[a].cnt, gs.Level.LevRecInner) {

				if gs.Player.PlProj[a].bounceN > 0 {
					gs.Player.PlProj[a].bounceN--
					gs.Player.PlProj[a].velx *= -1
					gs.Player.PlProj[a].vely *= -1
					if gs.Player.PlProj[a].name == "fireball" || gs.Player.PlProj[a].name == "fireworks" {
						gs.Player.PlProj[a].ro += 180
					}
				} else {
					gs.Player.PlProj[a].onoff = false
				}
			}
			//COLLISION INNER RECS
			if gs.Player.PlProj[a].onoff {
				for b := 0; b < len(gs.Level.Level[gs.Level.RoomNum].innerBloks); b++ {
					if rl.CheckCollisionRecs(gs.Player.PlProj[a].rec, gs.Level.Level[gs.Level.RoomNum].innerBloks[b].rec) {
						if gs.Player.PlProj[a].bounceN > 0 {
							gs.Player.PlProj[a].bounceN--
							gs.Player.PlProj[a].velx *= -1
							gs.Player.PlProj[a].vely *= -1
							if gs.Player.PlProj[a].name == "fireball" || gs.Player.PlProj[a].name == "fireworks" {
								gs.Player.PlProj[a].ro += 180
							}
						} else {
							gs.Player.PlProj[a].onoff = false
						}
					}
				}
			}

			switch gs.Player.PlProj[a].name {
			case "axe":
				gs.Player.PlProj[a].ro += 8
			}

		} else {
			clear = true
		}

	}

	if clear {
		for a := 0; a < len(gs.Player.PlProj); a++ {
			if !gs.Player.PlProj[a].onoff {
				gs.Player.PlProj = remProj(gs.Player.PlProj, a)
			}
		}
	}

	//CHECK ENEMY PROJ COLLIS
	if len(gs.Level.Level[gs.Level.RoomNum].enemies) > 0 {
		for a := 0; a < len(gs.Player.PlProj); a++ {
			if gs.Player.PlProj[a].onoff {
				for b := 0; b < len(gs.Level.Level[gs.Level.RoomNum].enemies); b++ {
					if rl.CheckCollisionRecs(gs.Player.PlProj[a].rec, gs.Level.Level[gs.Level.RoomNum].enemies[b].rec) && gs.Level.Level[gs.Level.RoomNum].enemies[b].hppause == 0 {
						gs.Level.Level[gs.Level.RoomNum].enemies[b].hppause = gs.Core.Fps / 2
						gs.Level.Level[gs.Level.RoomNum].enemies[b].hp -= gs.Player.PlProj[a].dmg
						playenemyhit()
						if gs.Level.Level[gs.Level.RoomNum].enemies[b].hp <= 0 {
							cntr := gs.Level.Level[gs.Level.RoomNum].enemies[b].cnt
							gs.Level.Level[gs.Level.RoomNum].enemies[b].off = true
							addkill(b)
							makeFX(2, cntr)
						}

					}
				}

			}
		}
	}
	//CHECK BOSS PROJ COLLIS
	if gs.Level.Levelnum == 6 {
		for a := 0; a < len(gs.Player.PlProj); a++ {
			if gs.Player.PlProj[a].onoff {

				if rl.CheckCollisionRecs(gs.Player.PlProj[a].rec, gs.Level.Bosses[gs.Level.Bossnum].crec) && gs.Level.Bosses[gs.Level.Bossnum].hppause == 0 {
					gs.Level.Bosses[gs.Level.Bossnum].hppause = gs.Core.Fps
					gs.Level.Bosses[gs.Level.Bossnum].hp -= gs.Player.PlProj[a].dmg
					if gs.Level.Bosses[gs.Level.Bossnum].hp <= 0 {
						cntr := gs.Level.Bosses[gs.Level.Bossnum].cnt
						gs.Level.Bosses[gs.Level.Bossnum].off = true
						makeFX(2, cntr)
						if !gs.Level.Endgame {
							gs.Core.Pause = true
							gs.Level.MinsEND = gs.Level.Mins
							gs.Level.SecsEND = gs.Level.Secs
							addtime()
							gs.Level.EndgopherRec = rl.NewRectangle(gs.Core.Cnt.X-bsU4, gs.Level.LevRec.Y+gs.Level.LevRec.Height, bsU8, bsU8)
							gs.Level.EndgameT = gs.Core.Fps * 3
							gs.Level.Endgame = true
						}
					}
					rl.PlaySound(gs.Audio.Sfx[11])
				}

			}
		}

	}

}
func drawUpEtc() { //MARK:DRAW UP ETC

	clear := false
	for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].etc); a++ {

		if gs.Level.Level[gs.Level.RoomNum].etc[a].onoff {

			//TIMERS
			if gs.Level.Level[gs.Level.RoomNum].etc[a].timer > 0 {
				gs.Level.Level[gs.Level.RoomNum].etc[a].timer--
			}
			if gs.Level.Level[gs.Level.RoomNum].etc[a].txtT > 0 {
				gs.Level.Level[gs.Level.RoomNum].etc[a].txtT--
			}
			//SHOP ARROWS
			if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "shop" {
				rl.DrawTriangle(gs.Level.Level[gs.Level.RoomNum].etc[a].v2s[4], gs.Level.Level[gs.Level.RoomNum].etc[a].v2s[3], gs.Level.Level[gs.Level.RoomNum].etc[a].v2s[5], rl.Fade(rl.SkyBlue, 0.8))

				if gs.Core.Frames%3 == 0 {
					gs.Level.Level[gs.Level.RoomNum].etc[a].v2s[3].Y--
					gs.Level.Level[gs.Level.RoomNum].etc[a].v2s[4].Y--
					gs.Level.Level[gs.Level.RoomNum].etc[a].v2s[5].Y--
				}

				if gs.Level.Level[gs.Level.RoomNum].etc[a].v2s[4].Y < gs.Level.Level[gs.Level.RoomNum].etc[a].rec.Y+gs.Level.Level[gs.Level.RoomNum].etc[a].rec.Height {

					gs.Level.Level[gs.Level.RoomNum].etc[a].v2s[3] = gs.Level.Level[gs.Level.RoomNum].etc[a].v2s[0]
					gs.Level.Level[gs.Level.RoomNum].etc[a].v2s[4] = gs.Level.Level[gs.Level.RoomNum].etc[a].v2s[1]
					gs.Level.Level[gs.Level.RoomNum].etc[a].v2s[5] = gs.Level.Level[gs.Level.RoomNum].etc[a].v2s[2]
				}
			}
			//DRAW BLOK
			if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "turret" {
				drawBlokDrec(gs.Level.Level[gs.Level.RoomNum].etc[a], true, false, 0)

				gs.Level.Level[gs.Level.RoomNum].etc[a].ro++
				if gs.Level.Level[gs.Level.RoomNum].etc[a].timer <= 0 {
					gs.Level.Level[gs.Level.RoomNum].etc[a].timer = rI32(1, 3) * gs.Core.Fps
					makeProjectileEnemy(1, gs.Level.Level[gs.Level.RoomNum].etc[a].cnt)
				}
			} else if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "spring" {

				drawBlokDrec(gs.Level.Level[gs.Level.RoomNum].etc[a], false, true, 4)

				if rl.CheckCollisionRecs(gs.Player.Pl.crec, gs.Level.Level[gs.Level.RoomNum].etc[a].crec) && gs.Level.Level[gs.Level.RoomNum].etc[a].timer == 0 {
					gs.Level.Level[gs.Level.RoomNum].etc[a].timer = gs.Core.Fps * 3
					gs.Level.Level[gs.Level.RoomNum].etc[a].onoffswitch = true

					gs.Player.Pl.slide = true
					gs.Player.Pl.slideT = gs.Core.Fps / 2
					gs.Player.Pl.slideDIR = gs.Level.Level[gs.Level.RoomNum].etc[a].slideDIR

					rl.PlaySound(gs.Audio.Sfx[3])
				}
				if gs.Level.Level[gs.Level.RoomNum].etc[a].onoffswitch {
					if gs.Core.Frames%4 == 0 {
						gs.Level.Level[gs.Level.RoomNum].etc[a].img.X += gs.Render.Spring.W
						if gs.Level.Level[gs.Level.RoomNum].etc[a].img.X > gs.Render.Spring.xl+gs.Render.Spring.frames*gs.Render.Spring.W {
							gs.Level.Level[gs.Level.RoomNum].etc[a].img.X = gs.Render.Spring.xl
							gs.Level.Level[gs.Level.RoomNum].etc[a].onoffswitch = false
						}
					}
				} else {
					if gs.Level.Level[gs.Level.RoomNum].etc[a].img.X > gs.Render.Spring.xl {
						if gs.Core.Frames%4 == 0 {
							gs.Level.Level[gs.Level.RoomNum].etc[a].img.X -= gs.Render.Spring.W
							if gs.Level.Level[gs.Level.RoomNum].etc[a].img.X < gs.Render.Spring.xl {
								gs.Level.Level[gs.Level.RoomNum].etc[a].img.X = gs.Render.Spring.xl
							}
						}
					}
				}

			} else if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "gascloudtrap" {
				drawBlokDrec(gs.Level.Level[gs.Level.RoomNum].etc[a], false, true, 4)
				gs.Level.Level[gs.Level.RoomNum].etc[a].ro++

				if rl.CheckCollisionRecs(gs.Player.Pl.crec, gs.Level.Level[gs.Level.RoomNum].etc[a].rec) && gs.Level.Level[gs.Level.RoomNum].etc[a].timer == 0 {
					gs.Level.Level[gs.Level.RoomNum].etc[a].timer = gs.Core.Fps * 3
					makegascloud(gs.Level.Level[gs.Level.RoomNum].etc[a].cnt)
					rl.PlaySound(gs.Audio.Sfx[20])
				}

			} else if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "gascloud" {

				rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[58], gs.Level.Level[gs.Level.RoomNum].etc[a].drec, gs.Level.Level[gs.Level.RoomNum].etc[a].ori, gs.Level.Level[gs.Level.RoomNum].etc[a].ro, rl.Fade(ranGreen(), rF32(0.3, 0.8)))

				if rl.CheckCollisionRecs(gs.Player.Pl.crec, gs.Level.Level[gs.Level.RoomNum].etc[a].crec) && !gs.Player.Pl.poison && gs.Player.Pl.poisonCollisT == 0 {
					gs.Player.Pl.poisonCollisT = gs.Core.Fps * 3
					if gs.Player.Mods.apple {
						gs.Player.Mods.appleN--
						if gs.Player.Mods.appleN == 0 {
							gs.Player.Mods.apple = false
							clearinven("apple")
						}
					} else {
						gs.Player.Pl.poison = true
						gs.Player.Pl.poisonCount = 2
						gs.Player.Pl.poisonT = gs.Core.Fps * 3
						txtHere("poisoned", gs.Player.Pl.rec)
						rl.PlaySound(gs.Audio.Sfx[23])
					}
				}

				if gs.Core.Frames%3 == 0 {
					gs.Level.Level[gs.Level.RoomNum].etc[a].img.X += gs.Render.Posiongas.W
					if gs.Level.Level[gs.Level.RoomNum].etc[a].img.X > gs.Render.Posiongas.xl+gs.Render.Posiongas.frames*gs.Render.Posiongas.W {
						gs.Level.Level[gs.Level.RoomNum].etc[a].img.X = gs.Render.Posiongas.xl
					}
				}

				rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[58], BlurRec(gs.Level.Level[gs.Level.RoomNum].etc[a].drec, 5), gs.Level.Level[gs.Level.RoomNum].etc[a].ori, gs.Level.Level[gs.Level.RoomNum].etc[a].ro, rl.Fade(ranGreen(), rF32(0.1, 0.3)))
				rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[58], BlurRec(gs.Level.Level[gs.Level.RoomNum].etc[a].drec, 7), gs.Level.Level[gs.Level.RoomNum].etc[a].ori, gs.Level.Level[gs.Level.RoomNum].etc[a].ro, rl.Fade(ranGreen(), rF32(0.1, 0.3)))

				if gs.Core.Debug {
					rl.DrawRectangleLinesEx(gs.Level.Level[gs.Level.RoomNum].etc[a].crec, 0.5, rl.White)
				}

				gs.Level.Level[gs.Level.RoomNum].etc[a].ro += 5

				if checknextmoveV2innerRec(gs.Level.Level[gs.Level.RoomNum].etc[a].cnt, gs.Level.Level[gs.Level.RoomNum].etc[a].velX, gs.Level.Level[gs.Level.RoomNum].etc[a].velY) {
					gs.Level.Level[gs.Level.RoomNum].etc[a].rec.X += gs.Level.Level[gs.Level.RoomNum].etc[a].velX
					gs.Level.Level[gs.Level.RoomNum].etc[a].rec.Y += gs.Level.Level[gs.Level.RoomNum].etc[a].velY
					gs.Level.Level[gs.Level.RoomNum].etc[a].drec.X += gs.Level.Level[gs.Level.RoomNum].etc[a].velX
					gs.Level.Level[gs.Level.RoomNum].etc[a].drec.Y += gs.Level.Level[gs.Level.RoomNum].etc[a].velY
					gs.Level.Level[gs.Level.RoomNum].etc[a].crec.X += gs.Level.Level[gs.Level.RoomNum].etc[a].velX
					gs.Level.Level[gs.Level.RoomNum].etc[a].crec.Y += gs.Level.Level[gs.Level.RoomNum].etc[a].velY
					gs.Level.Level[gs.Level.RoomNum].etc[a].cnt.X += gs.Level.Level[gs.Level.RoomNum].etc[a].velX
					gs.Level.Level[gs.Level.RoomNum].etc[a].cnt.Y += gs.Level.Level[gs.Level.RoomNum].etc[a].velY
				} else {
					gs.Level.Level[gs.Level.RoomNum].etc[a].velX = rF32(-gs.Level.Level[gs.Level.RoomNum].etc[a].vel, gs.Level.Level[gs.Level.RoomNum].etc[a].vel)
					gs.Level.Level[gs.Level.RoomNum].etc[a].velY = rF32(-gs.Level.Level[gs.Level.RoomNum].etc[a].vel, gs.Level.Level[gs.Level.RoomNum].etc[a].vel)
				}

			} else if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "slimetrail" || gs.Level.Level[gs.Level.RoomNum].etc[a].name == "playerblood" {

				drawBlok(gs.Level.Level[gs.Level.RoomNum].etc[a], false, true, 7)

				gs.Level.Level[gs.Level.RoomNum].etc[a].fade -= 0.005
				if gs.Level.Level[gs.Level.RoomNum].etc[a].fade <= 0 {
					gs.Level.Level[gs.Level.RoomNum].etc[a].onoff = false
				}

				if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "slimetrail" {
					if rl.CheckCollisionRecs(gs.Player.Pl.crec, gs.Level.Level[gs.Level.RoomNum].etc[a].crec) && !gs.Player.Pl.poison && gs.Player.Pl.poisonCollisT == 0 {
						gs.Player.Pl.poisonCollisT = gs.Core.Fps * 3
						if gs.Player.Mods.apple {
							gs.Player.Mods.appleN--
							if gs.Player.Mods.appleN == 0 {
								gs.Player.Mods.apple = false
								clearinven("apple")
							}
						} else {
							gs.Player.Pl.poison = true
							gs.Player.Pl.poisonCount = 2
							gs.Player.Pl.poisonT = gs.Core.Fps * 3
							txtHere("poisoned", gs.Player.Pl.rec)
						}
					}
				}

			} else if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "footprints" {

				drawBlokDrec(gs.Level.Level[gs.Level.RoomNum].etc[a], false, true, 4)

				gs.Level.Level[gs.Level.RoomNum].etc[a].fade -= 0.01
				if gs.Level.Level[gs.Level.RoomNum].etc[a].fade <= 0 {
					gs.Level.Level[gs.Level.RoomNum].etc[a].onoff = false
				}

			} else if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "blades" {
				drawBlok(gs.Level.Level[gs.Level.RoomNum].etc[a], true, true, 4)
				if gs.Core.Frames%2 == 0 {
					gs.Level.Level[gs.Level.RoomNum].etc[a].img.X += gs.Render.Blades.W
					if gs.Level.Level[gs.Level.RoomNum].etc[a].img.X > gs.Render.Blades.xl+gs.Render.Blades.frames*gs.Render.Blades.W {
						gs.Level.Level[gs.Level.RoomNum].etc[a].img.X = gs.Render.Blades.xl
						gs.Level.Level[gs.Level.RoomNum].etc[a].onoffswitch = false
					}
				}
				if checkNextMove(gs.Level.Level[gs.Level.RoomNum].etc[a].rec, gs.Level.Level[gs.Level.RoomNum].etc[a].velX, gs.Level.Level[gs.Level.RoomNum].etc[a].velY, true) {
					gs.Level.Level[gs.Level.RoomNum].etc[a].rec.X += gs.Level.Level[gs.Level.RoomNum].etc[a].velX
					gs.Level.Level[gs.Level.RoomNum].etc[a].rec.Y += gs.Level.Level[gs.Level.RoomNum].etc[a].velY
					gs.Level.Level[gs.Level.RoomNum].etc[a].cnt = rl.NewVector2(gs.Level.Level[gs.Level.RoomNum].etc[a].rec.X+gs.Level.Level[gs.Level.RoomNum].etc[a].rec.Width/2, gs.Level.Level[gs.Level.RoomNum].etc[a].rec.Y+gs.Level.Level[gs.Level.RoomNum].etc[a].rec.Width/2)
				} else {
					gs.Level.Level[gs.Level.RoomNum].etc[a].velX = rF32(-gs.Level.Level[gs.Level.RoomNum].etc[a].velX, gs.Level.Level[gs.Level.RoomNum].etc[a].velX)
					gs.Level.Level[gs.Level.RoomNum].etc[a].velY = rF32(-gs.Level.Level[gs.Level.RoomNum].etc[a].velY, gs.Level.Level[gs.Level.RoomNum].etc[a].velY)
				}

				if rl.CheckCollisionRecs(gs.Player.Pl.crec, gs.Level.Level[gs.Level.RoomNum].etc[a].rec) {
					if !rl.IsSoundPlaying(gs.Audio.Sfx[30]) {
						rl.PlaySound(gs.Audio.Sfx[30])
					}
					hitPL(0, 2)
				}

			} else if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "spear" {
				drawBlokDrec(gs.Level.Level[gs.Level.RoomNum].etc[a], false, true, 4)

				if rl.CheckCollisionRecs(gs.Player.Pl.crec, gs.Level.Level[gs.Level.RoomNum].etc[a].crec) {
					hitPL(0, 2)
					if gs.Core.Frames%4 == 0 {
						if gs.Level.Level[gs.Level.RoomNum].etc[a].img.X == gs.Render.Spear.xl {
							rl.PlaySound(gs.Audio.Sfx[26])
						}
						gs.Level.Level[gs.Level.RoomNum].etc[a].img.X += gs.Render.Spear.W
						if gs.Level.Level[gs.Level.RoomNum].etc[a].img.X > gs.Render.Spear.xl+gs.Render.Spear.frames*gs.Render.Spear.W {
							gs.Level.Level[gs.Level.RoomNum].etc[a].img.X = gs.Render.Spear.xl
						}
					}

				} else {
					if gs.Level.Level[gs.Level.RoomNum].etc[a].img.X != gs.Render.Spear.xl {
						if gs.Core.Frames%4 == 0 {
							gs.Level.Level[gs.Level.RoomNum].etc[a].img.X -= gs.Render.Spear.W
						}
					}
				}

			} else if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "powerupBlok" {
				drawBlok(gs.Level.Level[gs.Level.RoomNum].etc[a], true, false, 4)
			} else if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "flamefiretrail" {
				drawBlok(gs.Level.Level[gs.Level.RoomNum].etc[a], false, true, 4)
				if gs.Core.Frames%3 == 0 {
					gs.Level.Level[gs.Level.RoomNum].etc[a].img.X += gs.Render.Firetrailanim.W
					if gs.Level.Level[gs.Level.RoomNum].etc[a].img.X > gs.Render.Firetrailanim.xl+gs.Render.Firetrailanim.frames*gs.Render.Firetrailanim.W {
						gs.Level.Level[gs.Level.RoomNum].etc[a].img.X = gs.Render.Firetrailanim.xl
					}
				}

				gs.Level.Level[gs.Level.RoomNum].etc[a].fade -= 0.005
				if gs.Level.Level[gs.Level.RoomNum].etc[a].fade <= 0 {
					gs.Level.Level[gs.Level.RoomNum].etc[a].onoff = false
				}

			} else if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "airbomb" {
				drawBlok(gs.Level.Level[gs.Level.RoomNum].etc[a], false, true, 4)
				if gs.Core.Frames%5 == 0 {
					if gs.Level.Level[gs.Level.RoomNum].etc[a].img.X == gs.Render.Airstrikeanim.xl {
						rl.PlaySound(gs.Audio.Sfx[28])
					}
					gs.Level.Level[gs.Level.RoomNum].etc[a].img.X += gs.Render.Airstrikeanim.W
					if gs.Level.Level[gs.Level.RoomNum].etc[a].img.X > gs.Render.Airstrikeanim.xl+gs.Render.Airstrikeanim.frames*gs.Render.Airstrikeanim.W {
						gs.Level.Level[gs.Level.RoomNum].etc[a].img.X = gs.Render.Airstrikeanim.xl
					}
				}

				gs.Level.Level[gs.Level.RoomNum].etc[a].fade -= 0.03
				if gs.Level.Level[gs.Level.RoomNum].etc[a].fade <= 0 {
					gs.Level.Level[gs.Level.RoomNum].etc[a].onoff = false
				}

			} else { //DRAW OTHER BLOKS

				drawBlok(gs.Level.Level[gs.Level.RoomNum].etc[a], true, true, 4)

				if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "exit" {
					txt := "exit"
					txtlen := rl.MeasureText(txt, 20)
					txtx := int32(gs.Level.Level[gs.Level.RoomNum].etc[a].rec.X+gs.Level.Level[gs.Level.RoomNum].etc[a].rec.Width/2) - txtlen/2
					txtx++
					txty := int32(gs.Level.Level[gs.Level.RoomNum].etc[a].rec.Y) - 18
					rl.DrawText(txt, txtx-4, txty+4, 20, rl.Black)
					rl.DrawText(txt, txtx, txty, 20, ranCol())

					if rl.CheckCollisionRecs(gs.Level.Level[gs.Level.RoomNum].etc[a].rec, gs.Player.Pl.crec) {
						rl.PlaySound(gs.Audio.Sfx[15])
						gs.Level.Exited = true
					}
				}

				if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "chest" && gs.Level.Level[gs.Level.RoomNum].etc[a].img.X < 493 {

					if rl.CheckCollisionRecs(gs.Player.Pl.arec, gs.Level.Level[gs.Level.RoomNum].etc[a].crec) && gs.Player.Mods.key && !gs.Level.Level[gs.Level.RoomNum].etc[a].onoffswitch {
						rl.PlaySound(gs.Audio.Sfx[17])
						gs.Level.Level[gs.Level.RoomNum].etc[a].onoffswitch = true
						gs.Player.Mods.keyN--
						if gs.Player.Mods.keyN == 0 {
							gs.Player.Mods.key = false
						}
						delInven("key")
						makechestitem(a)
					} else if rl.CheckCollisionRecs(gs.Player.Pl.arec, gs.Level.Level[gs.Level.RoomNum].etc[a].crec) && !gs.Player.Mods.key && !gs.Level.Level[gs.Level.RoomNum].etc[a].onoffswitch && gs.Level.Level[gs.Level.RoomNum].etc[a].txtT == 0 {
						txtHere("locked", gs.Level.Level[gs.Level.RoomNum].etc[a].rec)
						gs.Level.Level[gs.Level.RoomNum].etc[a].txtT = gs.Core.Fps * 3
						rl.PlaySound(gs.Audio.Sfx[25])
					}

					if gs.Level.Level[gs.Level.RoomNum].etc[a].onoffswitch {
						if gs.Core.Frames%6 == 0 {
							gs.Level.Level[gs.Level.RoomNum].etc[a].img.X += 24
						}
					}

				}
			}
			//COLLISIONS PLAYER CREC > ETC REC
			if rl.CheckCollisionRecs(gs.Player.Pl.crec, gs.Level.Level[gs.Level.RoomNum].etc[a].rec) {
				switch gs.Level.Level[gs.Level.RoomNum].etc[a].name {

				case "gem":
					if gs.Level.Level[gs.Level.RoomNum].etc[a].onoff {
						gs.Player.Pl.coins += gs.Level.Level[gs.Level.RoomNum].etc[a].numCoins
						gs.Level.Level[gs.Level.RoomNum].etc[a].onoff = false
						txtAddCoinMulti()
					}

				case "switch": //MARK: SWITCHES ON/OFF
					if gs.Level.Level[gs.Level.RoomNum].etc[a].timer == 0 {
						rl.PlaySound(gs.Audio.Sfx[12])
						gs.Level.Level[gs.Level.RoomNum].etc[a].timer = gs.Core.Fps * 2
						gs.Level.Level[gs.Level.RoomNum].etc[a].onoffswitch = !gs.Level.Level[gs.Level.RoomNum].etc[a].onoffswitch
						if gs.Level.Level[gs.Level.RoomNum].etc[a].img == gs.Render.Etc[21] {
							gs.Level.Level[gs.Level.RoomNum].etc[a].img = gs.Render.Etc[22]
						} else {
							gs.Level.Level[gs.Level.RoomNum].etc[a].img = gs.Render.Etc[21]
						}

						switch gs.Level.Level[gs.Level.RoomNum].etc[a].numType {
						case 1:
							makeSwitchArrows()
						case 2:
							gs.Level.Night = !gs.Level.Night
						case 3:
							gs.Level.Flipcam = !gs.Level.Flipcam
							cams()
						case 4:
							gs.Render.Shader2On = !gs.Render.Shader2On
							if gs.Render.Shader3On {
								gs.Render.Shader3On = false
							}
						case 5:
							gs.Render.Shader3On = !gs.Render.Shader3On
							if gs.Render.Shader2On {
								gs.Render.Shader2On = false
							}
						case 6:
							if gs.Level.Level[gs.Level.RoomNum].etc[a].numCoins > 0 {
								gs.Level.Level[gs.Level.RoomNum].etc[a].numCoins--
								txtAddCoin()
							}

						}
					}

				//MARK:INGAME COLLECT INVEN
				case "health potion", "throwing axe", "recharge hp", "medi kit", "wallet", "mr planty", "apple", "key", "vine", "bounce", "santa", "fireball", "map", "firetrail", "invisible", "coffee", "teleport", "attack range", "attack damage", "orbital", "chain lightning", "health ring", "armor", "recharge", "anchor", "umbrella", "moldy socks", "cherry", "fish", "birthday cake", "peace", "mr alien", "air strike", "fireworks", "mr carrot", "mario":
					collectInven(a)
					gs.Level.Level[gs.Level.RoomNum].etc[a].onoff = false
				}
			}
			//COLLISIONS PLAYER CREC > ETC CREC2
			if rl.CheckCollisionRecs(gs.Player.Pl.crec, gs.Level.Level[gs.Level.RoomNum].etc[a].crec2) && gs.Level.Level[gs.Level.RoomNum].etc[a].name == "shop" && !gs.Player.Pl.escape && gs.Shop.ShopExitT == 0 {
				rl.PlaySound(gs.Audio.Sfx[16])
				gs.Shop.ShopExitY = gs.Level.Level[gs.Level.RoomNum].etc[a].rec.Y + gs.Level.Level[gs.Level.RoomNum].etc[a].rec.Height + bsU2
				gs.Shop.ShopOn = true
				gs.Core.Pause = true
			}

			//COLLISIONS PLAYER ATTACK AREA
			if rl.CheckCollisionRecs(gs.Player.Pl.atkrec, gs.Level.Level[gs.Level.RoomNum].etc[a].rec) {

				switch gs.Level.Level[gs.Level.RoomNum].etc[a].name {
				case "oilbarrel":
					if gs.Player.Pl.atk {
						makeFX(4, gs.Level.Level[gs.Level.RoomNum].etc[a].cnt)
						gs.Level.Level[gs.Level.RoomNum].etc[a].onoff = false
						rl.PlaySound(gs.Audio.Sfx[10])
						rl.PlaySound(gs.Audio.Sfx[5])
					}
				case "powerupBlok":
					if gs.Player.Pl.atk {
						if gs.Player.Mods.fireworks {
							gs.FX.FireworksCnt = gs.Level.Level[gs.Level.RoomNum].etc[a].cnt
							makeProjectile("fireworks")
						}
						destroyPowerupBlok(a)
						gs.Level.Level[gs.Level.RoomNum].etc[a].onoff = false
						rl.PlaySound(gs.Audio.Sfx[6])
					}
				}
			}
		} else {
			clear = true
		}
	}

	if clear {
		for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].etc); a++ {
			if !gs.Level.Level[gs.Level.RoomNum].etc[a].onoff {
				gs.Level.Level[gs.Level.RoomNum].etc = remBlok(gs.Level.Level[gs.Level.RoomNum].etc, a)
			}
		}

	}

	if gs.Level.Exited {
		gs.Core.Pause = true
		rl.DrawRectangle(0, 0, gs.Core.ScrW32, gs.Core.ScrH32, rl.Black)
		makenewlevel()
		gs.Level.Exited = false
	}

}
func drawupfx() { //MARK:DRAW UP FX

	clear := false
	for a := 0; a < len(gs.FX.Fx); a++ {

		switch gs.FX.Fx[a].name {

		case "fxBurnWoodBarrel", "fxBurnOilBarrel":
			col := ranOrange()
			rl.DrawTexturePro(gs.Render.Imgs, gs.FX.Fx[a].img, gs.FX.Fx[a].rec, gs.Core.Ori, 0, col)
			rl.DrawTexturePro(gs.Render.Imgs, gs.FX.Fx[a].img, BlurRec(gs.FX.Fx[a].rec, 4), gs.Core.Ori, 0, rl.Fade(col, rF32(0.2, 0.5)))
			if gs.Core.Frames%3 == 0 {
				gs.FX.Fx[a].img.X += 17
				if gs.FX.Fx[a].img.X >= gs.Render.Burn.xl+(17*gs.Render.Burn.frames) {
					gs.FX.Fx[a].img.X = gs.Render.Burn.xl
				}
			}

			if rl.CheckCollisionRecs(gs.Player.Pl.crec, gs.FX.Fx[a].rec) {
				hitPL(0, 2)
			}

			if gs.Core.Debug {
				rl.DrawRectangleLinesEx(gs.FX.Fx[a].rec, 0.5, rl.Blue)
			}

		case "fxEnemy":

			for b := 0; b < len(gs.FX.Fx[a].recs); b++ {
				v2 := rl.NewVector2(gs.FX.Fx[a].recs[b].rec.X+gs.FX.Fx[a].recs[b].rec.Width/2, gs.FX.Fx[a].recs[b].rec.Y+gs.FX.Fx[a].recs[b].rec.Height/2)
				rl.DrawCircleV(v2, gs.FX.Fx[a].recs[b].rec.Width/2, rl.Fade(gs.FX.Fx[a].recs[b].col, gs.FX.Fx[a].recs[b].fade))
				gs.FX.Fx[a].recs[b].rec.X += gs.FX.Fx[a].recs[b].velX
				gs.FX.Fx[a].recs[b].rec.Y += gs.FX.Fx[a].recs[b].velY
				gs.FX.Fx[a].recs[b].fade -= 0.05
				gs.FX.Fx[a].recs[b].rec.Width += 0.1
				gs.FX.Fx[a].recs[b].rec.Height += 0.1
			}

		case "fxBarrel":

			for b := 0; b < len(gs.FX.Fx[a].recs); b++ {
				rl.DrawRectangleRec(gs.FX.Fx[a].recs[b].rec, rl.Fade(gs.FX.Fx[a].recs[b].col, gs.FX.Fx[a].recs[b].fade))
				gs.FX.Fx[a].recs[b].rec.X += gs.FX.Fx[a].recs[b].velX
				gs.FX.Fx[a].recs[b].rec.Y += gs.FX.Fx[a].recs[b].velY
				gs.FX.Fx[a].recs[b].fade -= 0.05
				gs.FX.Fx[a].recs[b].rec.Width -= 0.05
				gs.FX.Fx[a].recs[b].rec.Height -= 0.05
			}

		}

		if gs.FX.Fx[a].onoff {

			gs.FX.Fx[a].timer--
			if gs.FX.Fx[a].timer == 0 {
				gs.FX.Fx[a].onoff = false
			}

		} else {
			clear = true
		}

	}

	if clear {
		for a := 0; a < len(gs.FX.Fx); a++ {
			if !gs.FX.Fx[a].onoff {
				gs.FX.Fx = remFX(gs.FX.Fx, a)
			}
		}
	}

}
func drawUpEnemies() { //MARK:DRAW UP ENEMIES

	clear := false
	for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].enemies); a++ {

		if !gs.Level.Level[gs.Level.RoomNum].enemies[a].off {

			//TIMERS
			if gs.Level.Level[gs.Level.RoomNum].enemies[a].T1 > 0 {
				gs.Level.Level[gs.Level.RoomNum].enemies[a].T1--
			}
			if gs.Level.Level[gs.Level.RoomNum].enemies[a].hppause > 0 {
				gs.Level.Level[gs.Level.RoomNum].enemies[a].hppause--
			}

			//PLAYER ENEMY CREC COLLIS
			if rl.CheckCollisionRecs(gs.Player.Pl.crec, gs.Level.Level[gs.Level.RoomNum].enemies[a].crec) && gs.Player.Pl.hppause == 0 {
				hitPL(0, 2)
			}

			shadowRec := gs.Level.Level[gs.Level.RoomNum].enemies[a].rec
			shadowRec.X -= 5
			shadowRec.Y += 5
			rl.DrawTexturePro(gs.Render.Imgs, gs.Level.Level[gs.Level.RoomNum].enemies[a].img, shadowRec, gs.Core.Ori, gs.Level.Level[gs.Level.RoomNum].enemies[a].ro, rl.Fade(rl.Black, 0.8))

			if gs.Level.Level[gs.Level.RoomNum].enemies[a].hppause > 0 {
				rl.DrawTexturePro(gs.Render.Imgs, gs.Level.Level[gs.Level.RoomNum].enemies[a].img, gs.Level.Level[gs.Level.RoomNum].enemies[a].rec, gs.Core.Ori, gs.Level.Level[gs.Level.RoomNum].enemies[a].ro, rl.Fade(ranCol(), gs.Level.Level[gs.Level.RoomNum].enemies[a].fade))
			} else {
				rl.DrawTexturePro(gs.Render.Imgs, gs.Level.Level[gs.Level.RoomNum].enemies[a].img, gs.Level.Level[gs.Level.RoomNum].enemies[a].rec, gs.Core.Ori, gs.Level.Level[gs.Level.RoomNum].enemies[a].ro, rl.Fade(gs.Level.Level[gs.Level.RoomNum].enemies[a].col, gs.Level.Level[gs.Level.RoomNum].enemies[a].fade))
			}

			if gs.Level.AnchorT > 0 {
				siz := bsU
				rec := rl.NewRectangle(gs.Level.Level[gs.Level.RoomNum].enemies[a].crec.X+gs.Level.Level[gs.Level.RoomNum].enemies[a].crec.Width/2-siz/2, gs.Level.Level[gs.Level.RoomNum].enemies[a].crec.Y+gs.Level.Level[gs.Level.RoomNum].enemies[a].crec.Height/2-siz/2, siz, siz)
				rec.Y -= siz * 2
				rl.DrawTexturePro(gs.Render.Imgs, gs.Render.Etc[39], rec, gs.Core.Ori, 0, ranCyan())
			}

			if gs.Core.Debug {
				rl.DrawRectangleLinesEx(gs.Level.Level[gs.Level.RoomNum].enemies[a].rec, 1, rl.Red)
				rl.DrawRectangleLinesEx(gs.Level.Level[gs.Level.RoomNum].enemies[a].crec, 1, rl.White)
				rl.DrawRectangleLinesEx(gs.Level.Level[gs.Level.RoomNum].enemies[a].arec, 1, rl.Magenta)
				rl.DrawCircleV(gs.Level.Level[gs.Level.RoomNum].enemies[a].cnt, 4, rl.Red)
			}

			//ANIM
			switch gs.Level.Level[gs.Level.RoomNum].enemies[a].name {

			case "ghost", "slime", "rock":
				if gs.Core.Frames%4 == 0 {
					if gs.Level.Level[gs.Level.RoomNum].enemies[a].imgr.X < gs.Level.Level[gs.Level.RoomNum].enemies[a].xImg2+(gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Width*float32(gs.Level.Level[gs.Level.RoomNum].enemies[a].frameNum)) {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].imgr.X += gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Width
					} else {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].imgr.X = gs.Level.Level[gs.Level.RoomNum].enemies[a].xImg2
					}
					if gs.Level.Level[gs.Level.RoomNum].enemies[a].imgl.X < gs.Level.Level[gs.Level.RoomNum].enemies[a].xImg+(gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Width*float32(gs.Level.Level[gs.Level.RoomNum].enemies[a].frameNum)) {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].imgl.X += gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Width
					} else {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].imgl.X = gs.Level.Level[gs.Level.RoomNum].enemies[a].xImg
					}
				}

				if gs.Level.Level[gs.Level.RoomNum].enemies[a].velX > 0 {
					gs.Level.Level[gs.Level.RoomNum].enemies[a].img = gs.Level.Level[gs.Level.RoomNum].enemies[a].imgr
				} else {
					gs.Level.Level[gs.Level.RoomNum].enemies[a].img = gs.Level.Level[gs.Level.RoomNum].enemies[a].imgl
				}

			case "mushroom":
				if gs.Core.Frames%3 == 0 {
					if gs.Level.Level[gs.Level.RoomNum].enemies[a].imgr.X < gs.Level.Level[gs.Level.RoomNum].enemies[a].xImg2+(gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Width*float32(gs.Level.Level[gs.Level.RoomNum].enemies[a].frameNum)) {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].imgr.X += gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Width
					} else {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].imgr.X = gs.Level.Level[gs.Level.RoomNum].enemies[a].xImg2
					}
					if gs.Level.Level[gs.Level.RoomNum].enemies[a].imgl.X < gs.Level.Level[gs.Level.RoomNum].enemies[a].xImg+(gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Width*float32(gs.Level.Level[gs.Level.RoomNum].enemies[a].frameNum)) {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].imgl.X += gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Width
					} else {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].imgl.X = gs.Level.Level[gs.Level.RoomNum].enemies[a].xImg
					}
				}

				if gs.Level.Level[gs.Level.RoomNum].enemies[a].velX > 0 {
					gs.Level.Level[gs.Level.RoomNum].enemies[a].img = gs.Level.Level[gs.Level.RoomNum].enemies[a].imgr
				} else {
					gs.Level.Level[gs.Level.RoomNum].enemies[a].img = gs.Level.Level[gs.Level.RoomNum].enemies[a].imgl
				}

			case "spikehog":
				if gs.Core.Frames%4 == 0 {
					if gs.Level.Level[gs.Level.RoomNum].enemies[a].img.X < gs.Level.Level[gs.Level.RoomNum].enemies[a].xImg+(gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Width*float32(gs.Level.Level[gs.Level.RoomNum].enemies[a].frameNum)) {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].img.X += gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Width
					} else {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].img.X = gs.Level.Level[gs.Level.RoomNum].enemies[a].xImg
					}
				}

			case "rabbit1":
				if gs.Core.Frames%8 == 0 {
					if gs.Level.Level[gs.Level.RoomNum].enemies[a].img.X < gs.Render.Rabbit1.xl+(gs.Render.Rabbit1.frames*gs.Render.Rabbit1.W) {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].img.X += gs.Render.Rabbit1.W
					} else {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].img.X = gs.Render.Rabbit1.xl
					}
				}
				switch gs.Level.Level[gs.Level.RoomNum].enemies[a].direc {
				case 1:
					gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Y = gs.Render.Rabbit1.yt + gs.Render.Rabbit1.W
				case 2:
					gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Y = gs.Render.Rabbit1.yt + gs.Render.Rabbit1.W*2
				case 3:
					gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Y = gs.Render.Rabbit1.yt
				case 4:
					gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Y = gs.Render.Rabbit1.yt + gs.Render.Rabbit1.W*3
				}
			case "bat":
				if gs.Core.Frames%8 == 0 {
					if gs.Level.Level[gs.Level.RoomNum].enemies[a].img.X < gs.Level.Level[gs.Level.RoomNum].enemies[a].xImg+(gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Width*float32(gs.Level.Level[gs.Level.RoomNum].enemies[a].frameNum)) {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].img.X += gs.Level.Level[gs.Level.RoomNum].enemies[a].img.Width
					} else {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].img.X = gs.Level.Level[gs.Level.RoomNum].enemies[a].xImg
					}
				}
			}

			//ENEMY HP BAR
			if gs.UI.HpBarsOn {
				hpX := gs.Level.Level[gs.Level.RoomNum].enemies[a].rec.X + gs.Level.Level[gs.Level.RoomNum].enemies[a].rec.Width/2
				siz := float32(4)
				wid := float32(gs.Level.Level[gs.Level.RoomNum].enemies[a].hpmax) * (siz + 1)
				hpX -= wid / 2
				hpY := gs.Level.Level[gs.Level.RoomNum].enemies[a].rec.Y + gs.Level.Level[gs.Level.RoomNum].enemies[a].rec.Height + 5

				rec := rl.NewRectangle(hpX, hpY, siz, siz)
				for b := 0; b < gs.Level.Level[gs.Level.RoomNum].enemies[a].hpmax; b++ {
					rl.DrawRectangleLinesEx(rec, 1, rl.White)
					rec.X += siz + 1
				}
				rec = rl.NewRectangle(hpX, hpY, siz, siz)
				for b := 0; b < gs.Level.Level[gs.Level.RoomNum].enemies[a].hp; b++ {
					rl.DrawRectangleRec(rec, rl.Red)
					rec.X += siz + 1
				}
			}

			//MOVE
			if gs.Level.AnchorT == 0 {
				moveenemy(a)
			}

			//MARK:PLAYER ENEMY ATTACK
			if rl.CheckCollisionRecs(gs.Player.Pl.atkrec, gs.Level.Level[gs.Level.RoomNum].enemies[a].rec) {
				if gs.Level.Level[gs.Level.RoomNum].enemies[a].hppause == 0 {
					if gs.Player.Pl.atk {
						gs.Level.Level[gs.Level.RoomNum].enemies[a].hppause = gs.Core.Fps / 2
						gs.Level.Level[gs.Level.RoomNum].enemies[a].hp -= gs.Player.Pl.atkDMG
						if gs.Level.Level[gs.Level.RoomNum].enemies[a].hp <= 0 {
							cntr := gs.Level.Level[gs.Level.RoomNum].enemies[a].cnt
							addkill(a)
							gs.Level.Level[gs.Level.RoomNum].enemies[a].off = true
							makeFX(2, cntr)
						} else {
							playenemyhit()
						}

						if gs.Player.Mods.chainlightning && !gs.Player.ChainLightingSwingOnOff {
							gs.Player.ChainLightingSwingOnOff = true
							if roll6() > 2 {
								makeChainLightning()

							}
						}
					}
				}
			}

		} else {
			clear = true
		}
	}

	if clear {
		for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].enemies); a++ {
			if gs.Level.Level[gs.Level.RoomNum].enemies[a].off {
				gs.Level.Level[gs.Level.RoomNum].enemies = remEnemy(gs.Level.Level[gs.Level.RoomNum].enemies, a)
			}
		}
	}

}
func drawUpBoss() { //MARK: DRAW UP BOSS

	//IMG
	shadowrec := gs.Level.Bosses[gs.Level.Bossnum].rec
	shadowrec.X -= 4
	shadowrec.Y += 4
	rl.DrawTexturePro(gs.Render.Imgs, gs.Level.Bosses[gs.Level.Bossnum].img, shadowrec, gs.Core.Ori, 0, rl.Fade(rl.Black, 0.7))

	if gs.Level.Bosses[gs.Level.Bossnum].hppause > 0 {
		rl.DrawTexturePro(gs.Render.Imgs, gs.Level.Bosses[gs.Level.Bossnum].img, gs.Level.Bosses[gs.Level.Bossnum].rec, gs.Core.Ori, 0, ranCol())
	} else {
		rl.DrawTexturePro(gs.Render.Imgs, gs.Level.Bosses[gs.Level.Bossnum].img, gs.Level.Bosses[gs.Level.Bossnum].rec, gs.Core.Ori, 0, rl.White)
	}

	if gs.Core.Frames%6 == 0 {
		gs.Level.Bosses[gs.Level.Bossnum].img.X += 48
		if gs.Level.Bosses[gs.Level.Bossnum].img.X > gs.Level.Bosses[gs.Level.Bossnum].xl+gs.Level.Bosses[gs.Level.Bossnum].img.Width*2 {
			gs.Level.Bosses[gs.Level.Bossnum].img.X = gs.Level.Bosses[gs.Level.Bossnum].xl
		}
	}
	switch gs.Level.Bosses[gs.Level.Bossnum].direc {
	case 1:
		gs.Level.Bosses[gs.Level.Bossnum].img.Y = gs.Level.Bosses[gs.Level.Bossnum].yt + gs.Level.Bosses[gs.Level.Bossnum].img.Height*3
	case 2:
		gs.Level.Bosses[gs.Level.Bossnum].img.Y = gs.Level.Bosses[gs.Level.Bossnum].yt + gs.Level.Bosses[gs.Level.Bossnum].img.Height*2
	case 3:
		gs.Level.Bosses[gs.Level.Bossnum].img.Y = gs.Level.Bosses[gs.Level.Bossnum].yt
	case 4:
		gs.Level.Bosses[gs.Level.Bossnum].img.Y = gs.Level.Bosses[gs.Level.Bossnum].yt + gs.Level.Bosses[gs.Level.Bossnum].img.Height
	}
	if gs.Core.Debug {
		rl.DrawRectangleLinesEx(gs.Level.Bosses[gs.Level.Bossnum].rec, 1, rl.White)
		rl.DrawRectangleLinesEx(gs.Level.Bosses[gs.Level.Bossnum].crec, 1, rl.White)
	}

	//HP BARS
	if gs.UI.HpBarsOn {
		hpX := gs.Level.Bosses[gs.Level.Bossnum].rec.X + gs.Level.Bosses[gs.Level.Bossnum].rec.Width/2
		siz := float32(4)
		wid := float32(gs.Level.Bosses[gs.Level.Bossnum].hpmax) * (siz + 1)
		hpX -= wid / 2
		hpY := gs.Level.Bosses[gs.Level.Bossnum].rec.Y + gs.Level.Bosses[gs.Level.Bossnum].rec.Height + 5

		rec := rl.NewRectangle(hpX, hpY, siz, siz)
		for b := 0; b < gs.Level.Bosses[gs.Level.Bossnum].hpmax; b++ {
			rl.DrawRectangleLinesEx(rec, 1, rl.White)
			rec.X += siz + 1
		}
		rec = rl.NewRectangle(hpX, hpY, siz, siz)
		for b := 0; b < gs.Level.Bosses[gs.Level.Bossnum].hp; b++ {
			rl.DrawRectangleRec(rec, rl.Red)
			rec.X += siz + 1
		}
	}

	//TIMERS
	gs.Level.Bosses[gs.Level.Bossnum].timer--
	if gs.Level.Bosses[gs.Level.Bossnum].timer <= 0 {
		gs.Level.Bosses[gs.Level.Bossnum].timer = gs.Core.Fps * rI32(1, 4)
		switch gs.Level.Bosses[gs.Level.Bossnum].atkType {
		case 1:
			makeProjectileEnemy(7, gs.Level.Bosses[gs.Level.Bossnum].cnt)
		case 2:
			makeProjectileEnemy(8, gs.Level.Bosses[gs.Level.Bossnum].cnt)
		case 3:
			makeProjectileEnemy(9, gs.Level.Bosses[gs.Level.Bossnum].cnt)
		}
	}
	if gs.Level.Bosses[gs.Level.Bossnum].hppause > 0 {
		gs.Level.Bosses[gs.Level.Bossnum].hppause--
	}

	//MARK:PLAYER BOSS ATTACK
	if rl.CheckCollisionRecs(gs.Player.Pl.atkrec, gs.Level.Bosses[gs.Level.Bossnum].crec) {
		if gs.Level.Bosses[gs.Level.Bossnum].hppause == 0 {
			if gs.Player.Pl.atk {
				gs.Level.Bosses[gs.Level.Bossnum].hppause = gs.Core.Fps
				gs.Level.Bosses[gs.Level.Bossnum].hp -= gs.Player.Pl.atkDMG
				if gs.Level.Bosses[gs.Level.Bossnum].hp <= 0 {
					cntr := gs.Level.Bosses[gs.Level.Bossnum].cnt
					gs.Level.Bosses[gs.Level.Bossnum].off = true
					makeFX(2, cntr)
					if !gs.Level.Endgame {
						gs.Level.MinsEND = gs.Level.Mins
						gs.Level.SecsEND = gs.Level.Secs
						addtime()
						gs.Level.Endgame = true
					}
				}
				rl.PlaySound(gs.Audio.Sfx[29])
			}
		}
	}

	//MOVE
	moveboss()

}

// MARK: CHECK CHECK CHECK CHECK CHECK CHECK CHECK CHECK CHECK CHECK CHECK CHECK CHECK CHECK
func checkcontroller() { //MARK:CHECK CONTROLLER

	gs.Input.IsController = rl.IsGamepadAvailable(0)
	if gs.Input.IsController && gs.Input.ControllerOn {
		gs.Input.UseController = true
	} else if !gs.Input.IsController {
		gs.Input.UseController = false
		gs.Input.ControllerDisconnect = true
	}

	if gs.Input.IsController {
		gs.Input.ControllerOn = true
		gs.Input.ControllerDisconnect = false
		gs.Input.ControllerWasOn = true
	}

	if gs.Input.ControllerDisconnect && gs.Input.ControllerWasOn && !gs.Core.Pause {
		gs.Input.ControllerOn = false
		gs.Input.ControllerWasOn = false
		gs.UI.OptionsOn = true
		gs.UI.OptionNum = 0
		gs.Core.Pause = true
	}
}
func checknextmoveV2innerRec(cnt rl.Vector2, velx, velxy float32) bool { //MARK:CHECK NEXT MOVE V2 POINT INNER REC
	canmove := true
	cnt.X += velx
	cnt.Y += velxy
	if !rl.CheckCollisionPointRec(cnt, gs.Level.LevRecInner) {
		canmove = false
	}
	return canmove
}

func checkNextMove(rec rl.Rectangle, velx, vely float32, destroyEtc bool) bool { //MARK:CHECK NEXT MOVE

	canmove := true

	nextRec := rec
	nextRec.X += velx
	nextRec.Y += vely

	tl, tr, br, bl := rl.NewVector2(nextRec.X, nextRec.Y), rl.NewVector2(nextRec.X+nextRec.Width, nextRec.Y), rl.NewVector2(nextRec.X+nextRec.Width, nextRec.Y+nextRec.Height), rl.NewVector2(nextRec.X, nextRec.Y+nextRec.Height)

	//CHECK INNER REC

	if !rl.CheckCollisionPointRec(tl, gs.Level.LevRecInner) || !rl.CheckCollisionPointRec(tr, gs.Level.LevRecInner) || !rl.CheckCollisionPointRec(br, gs.Level.LevRecInner) || !rl.CheckCollisionPointRec(bl, gs.Level.LevRecInner) {
		canmove = false
	}

	//CHECK MOVE BLOCKS
	if canmove {
		if len(gs.Level.Level[gs.Level.RoomNum].movBloks) > 0 {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].movBloks); a++ {
				if rl.CheckCollisionRecs(nextRec, gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec) {
					canmove = false
				}
			}
		}
	}

	//CHECK ETC
	if canmove {
		if len(gs.Level.Level[gs.Level.RoomNum].etc) > 0 {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].etc); a++ {
				if gs.Level.Level[gs.Level.RoomNum].etc[a].solid && gs.Level.Level[gs.Level.RoomNum].etc[a].name != "turret" {
					if rl.CheckCollisionRecs(gs.Level.Level[gs.Level.RoomNum].etc[a].rec, nextRec) {
						canmove = false
						if destroyEtc {
							switch gs.Level.Level[gs.Level.RoomNum].etc[a].name {
							case "powerupBlok":
								destroyPowerupBlok(a)
								makeFX(3, gs.Level.Level[gs.Level.RoomNum].etc[a].cnt)
								gs.Level.Level[gs.Level.RoomNum].etc[a].onoff = false
								rl.PlaySound(gs.Audio.Sfx[6])
							case "oilbarrel":
								makeFX(4, gs.Level.Level[gs.Level.RoomNum].etc[a].cnt)
								gs.Level.Level[gs.Level.RoomNum].etc[a].onoff = false
								rl.PlaySound(gs.Audio.Sfx[10])
								rl.PlaySound(gs.Audio.Sfx[5])
							}
						}
					}
				}
			}
		}
	}

	//CHECK INNER BLOKS
	if canmove {
		if len(gs.Level.Level[gs.Level.RoomNum].innerBloks) > 0 {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].innerBloks); a++ {
				if gs.Level.Level[gs.Level.RoomNum].innerBloks[a].solid {
					if rl.CheckCollisionRecs(gs.Level.Level[gs.Level.RoomNum].innerBloks[a].crec, nextRec) {
						canmove = false
					}
				}
			}
		}
	}

	return canmove

}
func checkInnerBloksExits(roomN int, rec rl.Rectangle) bool { //MARK:CHECK INNER BLOK EXITS

	canadd := true

	for a := 0; a < len(gs.Level.Level[roomN].doorExitRecs); a++ {
		if rl.CheckCollisionRecs(rec, gs.Level.Level[roomN].doorExitRecs[a]) {
			canadd = false
		}
	}

	if canadd {
		for a := 0; a < len(gs.Level.Level[roomN].innerBloks); a++ {
			if rl.CheckCollisionRecs(rec, gs.Level.Level[roomN].innerBloks[a].rec) {
				canadd = false
			}
		}
	}

	return canadd
}
func checkMoveBlok(blokNum int) bool { //MARK:CHECK MOVE BLOK

	canmove := true
	checkBlok := gs.Level.Level[gs.Level.RoomNum].movBloks[blokNum]
	nextRec := checkBlok.rec
	if checkBlok.velX != 0 {
		nextRec.X += checkBlok.velX
	}
	if checkBlok.velY != 0 {
		nextRec.Y += checkBlok.velY
	}

	//CHECK PLAYER
	if rl.CheckCollisionRecs(nextRec, gs.Player.Pl.crec) {
		canmove = false
	}

	//CHECK MOVE BLOKS
	if canmove {
		if len(gs.Level.Level[gs.Level.RoomNum].movBloks) > 1 {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].movBloks); a++ {
				if blokNum != a {
					if rl.CheckCollisionRecs(nextRec, gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec) {
						canmove = false
					}
				}
			}
		}
	}

	//CHECK INNER BLOKS
	if canmove {
		if len(gs.Level.Level[gs.Level.RoomNum].innerBloks) > 0 {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].innerBloks); a++ {
				if rl.CheckCollisionRecs(nextRec, gs.Level.Level[gs.Level.RoomNum].innerBloks[a].crec) {
					canmove = false
				}
			}
		}
	}
	//CHECK BOUNDARY
	if canmove {
		checkV1 := rl.NewVector2(nextRec.X, nextRec.Y)
		checkV2 := checkV1
		checkV2.X += nextRec.Width
		checkV3 := checkV2
		checkV3.Y += nextRec.Height
		checkV4 := checkV1
		checkV4.Y += nextRec.Height

		if !rl.CheckCollisionPointRec(checkV1, gs.Level.LevRecInner) || !rl.CheckCollisionPointRec(checkV2, gs.Level.LevRecInner) || !rl.CheckCollisionPointRec(checkV3, gs.Level.LevRecInner) || !rl.CheckCollisionPointRec(checkV4, gs.Level.LevRecInner) {
			canmove = false
		}
	}

	return canmove

}
func checkplayermove(direc int) bool { //MARK:CHECK PLAYER MOVE

	canmove := true
	nextRec := gs.Player.Pl.crec

	switch direc {
	case 1:
		nextRec.Y -= gs.Player.Pl.vel
	case 2:
		nextRec.X += gs.Player.Pl.vel
	case 3:
		nextRec.Y += gs.Player.Pl.vel
	case 4:
		nextRec.X -= gs.Player.Pl.vel
	}

	tl, tr, br, bl := rl.NewVector2(nextRec.X, nextRec.Y), rl.NewVector2(nextRec.X+nextRec.Width, nextRec.Y), rl.NewVector2(nextRec.X+nextRec.Width, nextRec.Y+nextRec.Height), rl.NewVector2(nextRec.X, nextRec.Y+nextRec.Height)

	//CHECK BOUNDARY WALLS
	dWalls := gs.Level.Level[gs.Level.RoomNum].walls
	for a := 0; a < len(dWalls); a++ {
		if rl.CheckCollisionPointRec(tl, dWalls[a].rec) || rl.CheckCollisionPointRec(tr, dWalls[a].rec) || rl.CheckCollisionPointRec(br, dWalls[a].rec) || rl.CheckCollisionPointRec(bl, dWalls[a].rec) {
			canmove = false
		}
		if rl.CheckCollisionRecs(nextRec, dWalls[a].rec) {
			canmove = false
		}
	}

	//CHECK MOVE BLOCKS
	if canmove {
		if len(gs.Level.Level[gs.Level.RoomNum].movBloks) > 0 {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].movBloks); a++ {
				if rl.CheckCollisionRecs(nextRec, gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec) {
					canmove = false
					if gs.Level.Level[gs.Level.RoomNum].movBloks[a].bump {
						switch direc {
						case 1:
							gs.Level.Level[gs.Level.RoomNum].movBloks[a].velY += -gs.Player.Pl.vel / 8
						case 2:
							gs.Level.Level[gs.Level.RoomNum].movBloks[a].velX += gs.Player.Pl.vel / 8
						case 3:
							gs.Level.Level[gs.Level.RoomNum].movBloks[a].velY += gs.Player.Pl.vel / 8
						case 4:
							gs.Level.Level[gs.Level.RoomNum].movBloks[a].velX -= gs.Player.Pl.vel / 8
						}
					}
				}

			}
		}
	}

	//CHECK ETC
	if canmove {
		if len(gs.Level.Level[gs.Level.RoomNum].etc) > 0 {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].etc); a++ {
				if gs.Level.Level[gs.Level.RoomNum].etc[a].solid {
					if rl.CheckCollisionRecs(gs.Level.Level[gs.Level.RoomNum].etc[a].rec, nextRec) {
						canmove = false
					}
				}
			}
		}
	}

	//CHECK INNER BLOKS
	if canmove {
		if len(gs.Level.Level[gs.Level.RoomNum].innerBloks) > 0 {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].innerBloks); a++ {
				if gs.Level.Level[gs.Level.RoomNum].innerBloks[a].solid {
					if rl.CheckCollisionRecs(gs.Level.Level[gs.Level.RoomNum].innerBloks[a].crec, nextRec) {
						canmove = false
					}
				}
			}
		}

	}

	return canmove
}

// MARK: FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND
func findRecPoswithSpacing(wid, space float32, numRoom int) (rec rl.Rectangle, found bool) { //MARK: FIND REC POS WITH SPACING

	countbreak := 500
	found = true

	wid2 := wid + (space * 2)

	for {
		rec = rl.NewRectangle(rF32(gs.Level.LevRecInner.X, gs.Level.LevRecInner.X+gs.Level.LevRecInner.Width-wid2), rF32(gs.Level.LevRecInner.Y, gs.Level.LevRecInner.Y+gs.Level.LevRecInner.Width-wid2), wid2, wid2)
		canadd := true

		for a := 0; a < len(gs.Level.Level[numRoom].doorExitRecs); a++ {
			if rl.CheckCollisionRecs(rec, gs.Level.Level[numRoom].doorExitRecs[a]) {
				canadd = false
			}
		}
		if canadd {
			for a := 0; a < len(gs.Level.Level[numRoom].movBloks); a++ {
				if rl.CheckCollisionRecs(rec, gs.Level.Level[numRoom].movBloks[a].rec) {
					canadd = false
				}
			}
		}
		if canadd {
			for a := 0; a < len(gs.Level.Level[numRoom].innerBloks); a++ {
				if rl.CheckCollisionRecs(rec, gs.Level.Level[numRoom].innerBloks[a].rec) {
					canadd = false
				}
			}
		}
		if canadd {
			for a := 0; a < len(gs.Level.Level[numRoom].etc); a++ {
				if rl.CheckCollisionRecs(rec, gs.Level.Level[numRoom].etc[a].rec) && gs.Level.Level[numRoom].etc[a].solid {
					canadd = false
				}
			}
		}

		countbreak--
		if countbreak == 0 || canadd {

			rec.X += space
			rec.Y += space
			rec.Width -= space * 2
			rec.Height -= space * 2

			if countbreak == 0 {
				found = false
			}
			break
		}

	}

	return rec, found
}
func findRecPos(wid float32, numRoom int) (rec rl.Rectangle, found bool) { //MARK: FIND REC POS

	countbreak := 500
	found = true

	for {
		rec = rl.NewRectangle(rF32(gs.Level.LevRecInner.X, gs.Level.LevRecInner.X+gs.Level.LevRecInner.Width-wid), rF32(gs.Level.LevRecInner.Y, gs.Level.LevRecInner.Y+gs.Level.LevRecInner.Width-wid), wid, wid)
		canadd := true
		for a := 0; a < len(gs.Level.Level[numRoom].movBloks); a++ {
			if rl.CheckCollisionRecs(rec, gs.Level.Level[numRoom].movBloks[a].rec) {
				canadd = false
			}
		}
		if canadd {
			for a := 0; a < len(gs.Level.Level[numRoom].innerBloks); a++ {
				if rl.CheckCollisionRecs(rec, gs.Level.Level[numRoom].innerBloks[a].rec) {
					canadd = false
				}
			}
		}
		if canadd {
			for a := 0; a < len(gs.Level.Level[numRoom].etc); a++ {
				if rl.CheckCollisionRecs(rec, gs.Level.Level[numRoom].etc[a].rec) && gs.Level.Level[numRoom].etc[a].solid {
					canadd = false
				}
			}
		}

		countbreak--
		if countbreak == 0 || canadd {
			if countbreak == 0 {
				found = false
			}
			break
		}

	}

	return rec, found
}
func findRanCntV2() rl.Vector2 { //MARK: FIND RANDOM CNTR V2
	v2 := gs.Core.Cnt
	wid := gs.Level.LevW / 2
	wid -= bsU2
	v2.X += rF32(-wid, wid)
	v2.Y += rF32(-wid, wid)
	return v2
}
func findRanRecLoc(w, h float32, roomNum int) (tl rl.Vector2) { //MARK: FIND RANDOM RECTANGLE LOCATION

	v2 := rl.Vector2{}
	countbreak := 100
	for {
		v2 = rl.NewVector2(gs.Level.LevRecInner.X+bsU2, gs.Level.LevRecInner.Y+bsU2)
		wid := gs.Level.LevRecInner.Width - bsU4
		heig := wid
		wid -= w
		heig -= h
		v2.X += rF32(0, wid)
		v2.Y += rF32(0, heig)

		checkrec := rl.NewRectangle(v2.X, v2.Y, w, h)
		canadd := true
		for a := 0; a < len(gs.Level.Level[roomNum].doorExitRecs); a++ {
			if rl.CheckCollisionRecs(checkrec, gs.Level.Level[roomNum].doorExitRecs[a]) {
				canadd = false
			}
		}

		countbreak--
		if canadd || countbreak == 0 {
			break
		}
	}

	return v2
}

// MARK: MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE
func moveboss() { //MARK:MOVE BOSS

	checkRec := gs.Level.Bosses[gs.Level.Bossnum].crec
	checkRec.X += gs.Level.Bosses[gs.Level.Bossnum].velX
	checkRec.Y += gs.Level.Bosses[gs.Level.Bossnum].velY

	canmove := true

	tl := rl.NewVector2(checkRec.X, checkRec.Y)
	tr := tl
	tr.X += checkRec.Width
	br := tr
	br.Y += checkRec.Height
	bl := tl
	bl.Y += checkRec.Height

	//CHECK BOUNDARY
	if !rl.CheckCollisionPointRec(tl, gs.Level.LevRecInner) || !rl.CheckCollisionPointRec(tr, gs.Level.LevRecInner) || !rl.CheckCollisionPointRec(br, gs.Level.LevRecInner) || !rl.CheckCollisionPointRec(bl, gs.Level.LevRecInner) {
		canmove = false
	}

	//CHECK INNER BLOKS
	if canmove {
		if len(gs.Level.Level[gs.Level.RoomNum].innerBloks) > 0 {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].innerBloks); a++ {
				if rl.CheckCollisionRecs(gs.Level.Level[gs.Level.RoomNum].innerBloks[a].crec, checkRec) {
					canmove = false
				}
			}
		}
	}

	//MOVE
	if canmove {

		gs.Level.Bosses[gs.Level.Bossnum].rec.X += gs.Level.Bosses[gs.Level.Bossnum].velX
		gs.Level.Bosses[gs.Level.Bossnum].rec.Y += gs.Level.Bosses[gs.Level.Bossnum].velY
		gs.Level.Bosses[gs.Level.Bossnum].crec = gs.Level.Bosses[gs.Level.Bossnum].rec
		gs.Level.Bosses[gs.Level.Bossnum].crec.X += gs.Level.Bosses[gs.Level.Bossnum].rec.Width / 4
		gs.Level.Bosses[gs.Level.Bossnum].crec.Y += gs.Level.Bosses[gs.Level.Bossnum].rec.Height / 5
		gs.Level.Bosses[gs.Level.Bossnum].crec.Width -= gs.Level.Bosses[gs.Level.Bossnum].rec.Width / 2
		gs.Level.Bosses[gs.Level.Bossnum].crec.Height -= gs.Level.Bosses[gs.Level.Bossnum].rec.Height / 5
		gs.Level.Bosses[gs.Level.Bossnum].cnt = rl.NewVector2(gs.Level.Bosses[gs.Level.Bossnum].rec.X+gs.Level.Bosses[gs.Level.Bossnum].rec.Width/2, gs.Level.Bosses[gs.Level.Bossnum].rec.Y+gs.Level.Bosses[gs.Level.Bossnum].rec.Height/2)

		//CHECK SOCKS COLLIS
		if gs.Player.Mods.socks && gs.Level.Bosses[gs.Level.Bossnum].hppause == 0 {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].etc); a++ {
				if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "footprints" && gs.Level.Bosses[gs.Level.Bossnum].hppause == 0 && rl.CheckCollisionRecs(gs.Level.Level[gs.Level.RoomNum].etc[a].rec, gs.Level.Bosses[gs.Level.Bossnum].crec) {
					gs.Level.Bosses[gs.Level.Bossnum].hppause = gs.Core.Fps
					gs.Level.Bosses[gs.Level.Bossnum].hp--
					if gs.Level.Bosses[gs.Level.Bossnum].hp <= 0 {
						cntr := gs.Level.Bosses[gs.Level.Bossnum].cnt
						gs.Level.Bosses[gs.Level.Bossnum].off = true
						makeFX(2, cntr)
						if !gs.Level.Endgame {
							gs.Level.MinsEND = gs.Level.Mins
							gs.Level.SecsEND = gs.Level.Secs
							addtime()
							gs.Level.Endgame = true
						}
					}
					rl.PlaySound(gs.Audio.Sfx[29])
				}
			}
		}
		//CHECK ORBITAL COLLIS
		if gs.Player.Mods.orbital && gs.Level.Bosses[gs.Level.Bossnum].hppause == 0 {
			if rl.CheckCollisionRecs(gs.Player.Pl.orbrec1, gs.Level.Bosses[gs.Level.Bossnum].crec) {
				gs.Level.Bosses[gs.Level.Bossnum].hppause = gs.Core.Fps
				gs.Level.Bosses[gs.Level.Bossnum].hp--
				if gs.Level.Bosses[gs.Level.Bossnum].hp <= 0 {
					cntr := gs.Level.Bosses[gs.Level.Bossnum].cnt
					gs.Level.Bosses[gs.Level.Bossnum].off = true
					makeFX(2, cntr)
					if !gs.Level.Endgame {
						gs.Level.MinsEND = gs.Level.Mins
						gs.Level.SecsEND = gs.Level.Secs
						addtime()
						gs.Level.Endgame = true
					}
					rl.PlaySound(gs.Audio.Sfx[29])
				}
			}

			if gs.Player.Mods.orbitalN == 2 {
				if rl.CheckCollisionRecs(gs.Player.Pl.orbrec2, gs.Level.Bosses[gs.Level.Bossnum].crec) {
					gs.Level.Bosses[gs.Level.Bossnum].hppause = gs.Core.Fps
					gs.Level.Bosses[gs.Level.Bossnum].hp--
					if gs.Level.Bosses[gs.Level.Bossnum].hp <= 0 {
						cntr := gs.Level.Bosses[gs.Level.Bossnum].cnt
						gs.Level.Bosses[gs.Level.Bossnum].off = true
						makeFX(2, cntr)
						if !gs.Level.Endgame {
							gs.Level.MinsEND = gs.Level.Mins
							gs.Level.SecsEND = gs.Level.Secs
							addtime()
							gs.Level.Endgame = true
						}
					}
					rl.PlaySound(gs.Audio.Sfx[29])
				}
			}

		}
		//CHECK AIRSTRIKE COLLIS
		if gs.Player.Mods.airstrike && gs.Level.Bosses[gs.Level.Bossnum].hppause == 0 {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].etc); a++ {
				if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "airbomb" && rl.CheckCollisionRecs(gs.Level.Level[gs.Level.RoomNum].etc[a].crec, gs.Level.Bosses[gs.Level.Bossnum].crec) {
					gs.Level.Bosses[gs.Level.Bossnum].hppause = gs.Core.Fps
					gs.Level.Bosses[gs.Level.Bossnum].hp--
					if gs.Level.Bosses[gs.Level.Bossnum].hp <= 0 {
						cntr := gs.Level.Bosses[gs.Level.Bossnum].cnt
						gs.Level.Bosses[gs.Level.Bossnum].off = true
						makeFX(2, cntr)
						if !gs.Level.Endgame {
							gs.Level.MinsEND = gs.Level.Mins
							gs.Level.SecsEND = gs.Level.Secs
							addtime()
							gs.Level.Endgame = true
						}
						rl.PlaySound(gs.Audio.Sfx[29])
					}
				}
			}
		}
		//CHECK FIRETRAIL COLLIS
		if gs.Player.Mods.firetrail {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].etc); a++ {
				if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "flamefiretrail" && gs.Level.Bosses[gs.Level.Bossnum].hppause == 0 && rl.CheckCollisionRecs(gs.Level.Level[gs.Level.RoomNum].etc[a].rec, gs.Level.Bosses[gs.Level.Bossnum].crec) {
					gs.Level.Bosses[gs.Level.Bossnum].hppause = gs.Core.Fps
					gs.Level.Bosses[gs.Level.Bossnum].hp--
					if gs.Level.Bosses[gs.Level.Bossnum].hp <= 0 && !gs.Level.Endgame {
						gs.Core.Pause = true
						gs.Level.EndPauseT = gs.Core.Fps * 5
						cntr := gs.Level.Bosses[gs.Level.Bossnum].cnt
						gs.Level.Bosses[gs.Level.Bossnum].off = true
						makeFX(2, cntr)
						if !gs.Level.Endgame {
							gs.Level.MinsEND = gs.Level.Mins
							gs.Level.SecsEND = gs.Level.Secs
							addtime()
							gs.Level.Endgame = true
						}
					}
					rl.PlaySound(gs.Audio.Sfx[29])
				}
			}
		}

	} else {
		gs.Level.Bosses[gs.Level.Bossnum].velX = rF32(-gs.Level.Bosses[gs.Level.Bossnum].vel, gs.Level.Bosses[gs.Level.Bossnum].vel)
		gs.Level.Bosses[gs.Level.Bossnum].velY = rF32(-gs.Level.Bosses[gs.Level.Bossnum].vel, gs.Level.Bosses[gs.Level.Bossnum].vel)
	}

	//FIND DIREC FOR ANIM
	if getabs(gs.Level.Bosses[gs.Level.Bossnum].velX) > getabs(gs.Level.Bosses[gs.Level.Bossnum].velY) {
		if gs.Level.Bosses[gs.Level.Bossnum].velX > 0 {
			gs.Level.Bosses[gs.Level.Bossnum].direc = 2
		} else {
			gs.Level.Bosses[gs.Level.Bossnum].direc = 4
		}
	} else {
		if gs.Level.Bosses[gs.Level.Bossnum].velY > 0 {
			gs.Level.Bosses[gs.Level.Bossnum].direc = 3
		} else {
			gs.Level.Bosses[gs.Level.Bossnum].direc = 1
		}

	}

}

func moveFollow(v2, targetV2 rl.Vector2, vel float32) (velx, vely float32) { //MARK:MOVE FOLLOW

	xdiff := absdiff(v2.X, targetV2.X)
	ydiff := absdiff(v2.Y, targetV2.Y)

	if xdiff > ydiff {
		velx = vel
		vely = ydiff / (xdiff / vel)
		if v2.X > targetV2.X {
			velx = -velx
		}
		if v2.Y > targetV2.Y {
			vely = -vely
		}
	} else {
		vely = vel
		velx = xdiff / (ydiff / vel)
		if v2.X > targetV2.X {
			velx = -velx
		}
		if v2.Y > targetV2.Y {
			vely = -vely
		}
	}

	return velx, vely
}
func moveenemy(num int) { //MARK:MOVE ENEMY

	checkRec := gs.Level.Level[gs.Level.RoomNum].enemies[num].rec
	checkRec.X += gs.Level.Level[gs.Level.RoomNum].enemies[num].velX
	checkRec.Y += gs.Level.Level[gs.Level.RoomNum].enemies[num].velY

	canmove := true

	tl := rl.NewVector2(checkRec.X, checkRec.Y)
	tr := tl
	tr.X += checkRec.Width
	br := tr
	br.Y += checkRec.Height
	bl := tl
	bl.Y += checkRec.Height

	//CHECK BOUNDARY
	if !rl.CheckCollisionPointRec(tl, gs.Level.LevRecInner) || !rl.CheckCollisionPointRec(tr, gs.Level.LevRecInner) || !rl.CheckCollisionPointRec(br, gs.Level.LevRecInner) || !rl.CheckCollisionPointRec(bl, gs.Level.LevRecInner) {
		canmove = false
	}

	//CHECK INNER BLOKS
	if canmove {
		if len(gs.Level.Level[gs.Level.RoomNum].innerBloks) > 0 {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].innerBloks); a++ {
				if rl.CheckCollisionRecs(gs.Level.Level[gs.Level.RoomNum].innerBloks[a].crec, checkRec) {
					canmove = false
				}
			}
		}
	}

	//MOVE
	if canmove {

		gs.Level.Level[gs.Level.RoomNum].enemies[num].rec = checkRec
		gs.Level.Level[gs.Level.RoomNum].enemies[num].crec = gs.Level.Level[gs.Level.RoomNum].enemies[num].rec

		//MOVE COLLIS REC & AREA REC
		switch gs.Level.Level[gs.Level.RoomNum].enemies[num].name {
		case "mushroom":
			gs.Level.Level[gs.Level.RoomNum].enemies[num].crec = gs.Level.Level[gs.Level.RoomNum].enemies[num].rec
			gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Y += gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Height / 2
			gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Height -= gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Height / 2
		case "slime", "rock":
			gs.Level.Level[gs.Level.RoomNum].enemies[num].crec = gs.Level.Level[gs.Level.RoomNum].enemies[num].rec
			gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Y += gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Height / 3
			gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Height -= gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Height / 3
		case "spikehog":
			gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Height = gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Height / 2
			gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Y += gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Height
		case "ghost":
			gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Y += gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Height / 3
			gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Height -= gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Height / 3
			gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.X += gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Width / 8
			gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Width -= gs.Level.Level[gs.Level.RoomNum].enemies[num].crec.Width / 4
			gs.Level.Level[gs.Level.RoomNum].enemies[num].arec = gs.Level.Level[gs.Level.RoomNum].enemies[num].crec
			gs.Level.Level[gs.Level.RoomNum].enemies[num].arec.X -= gs.Level.Level[gs.Level.RoomNum].enemies[num].arec.Width * 2
			gs.Level.Level[gs.Level.RoomNum].enemies[num].arec.Y -= gs.Level.Level[gs.Level.RoomNum].enemies[num].arec.Width * 2
			gs.Level.Level[gs.Level.RoomNum].enemies[num].arec.Width = gs.Level.Level[gs.Level.RoomNum].enemies[num].arec.Width * 5
			gs.Level.Level[gs.Level.RoomNum].enemies[num].arec.Height = gs.Level.Level[gs.Level.RoomNum].enemies[num].arec.Height * 5

		}

		gs.Level.Level[gs.Level.RoomNum].enemies[num].cnt = rl.NewVector2(checkRec.X+checkRec.Width/2, checkRec.Y+checkRec.Height/2)

		//CHECK SOCKS COLLIS
		if gs.Player.Mods.socks && gs.Level.Level[gs.Level.RoomNum].enemies[num].hppause == 0 && !gs.Level.Level[gs.Level.RoomNum].enemies[num].fly {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].etc); a++ {
				if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "footprints" && gs.Level.Level[gs.Level.RoomNum].enemies[num].hppause == 0 && rl.CheckCollisionRecs(gs.Level.Level[gs.Level.RoomNum].etc[a].rec, gs.Level.Level[gs.Level.RoomNum].enemies[num].crec) {
					gs.Level.Level[gs.Level.RoomNum].enemies[num].hppause = gs.Core.Fps / 2
					gs.Level.Level[gs.Level.RoomNum].enemies[num].hp--
					if gs.Level.Level[gs.Level.RoomNum].enemies[num].hp <= 0 {
						cntr := gs.Level.Level[gs.Level.RoomNum].enemies[num].cnt
						addkill(num)
						gs.Level.Level[gs.Level.RoomNum].enemies[num].off = true
						makeFX(2, cntr)
					} else {
						playenemyhit()
					}
				}
			}
		}
		//CHECK ORBITAL COLLIS
		if gs.Player.Mods.orbital && gs.Level.Level[gs.Level.RoomNum].enemies[num].hppause == 0 {

			if rl.CheckCollisionRecs(gs.Player.Pl.orbrec1, gs.Level.Level[gs.Level.RoomNum].enemies[num].crec) {
				gs.Level.Level[gs.Level.RoomNum].enemies[num].hppause = gs.Core.Fps / 2
				gs.Level.Level[gs.Level.RoomNum].enemies[num].hp--
				if gs.Level.Level[gs.Level.RoomNum].enemies[num].hp <= 0 {
					cntr := gs.Level.Level[gs.Level.RoomNum].enemies[num].cnt
					addkill(num)
					gs.Level.Level[gs.Level.RoomNum].enemies[num].off = true
					makeFX(2, cntr)
				} else {
					playenemyhit()
				}
			}

			if gs.Player.Mods.orbitalN == 2 {
				if rl.CheckCollisionRecs(gs.Player.Pl.orbrec2, gs.Level.Level[gs.Level.RoomNum].enemies[num].crec) {
					gs.Level.Level[gs.Level.RoomNum].enemies[num].hppause = gs.Core.Fps / 2
					gs.Level.Level[gs.Level.RoomNum].enemies[num].hp--
					if gs.Level.Level[gs.Level.RoomNum].enemies[num].hp <= 0 {
						cntr := gs.Level.Level[gs.Level.RoomNum].enemies[num].cnt
						addkill(num)
						gs.Level.Level[gs.Level.RoomNum].enemies[num].off = true
						makeFX(2, cntr)
					} else {
						if flipcoin() {
							rl.PlaySound(gs.Audio.Sfx[1])
						} else {
							rl.PlaySound(gs.Audio.Sfx[2])
						}
					}
				}
			}

		}
		//CHECK AIRSTRIKE COLLIS
		if gs.Player.Mods.airstrike {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].etc); a++ {
				if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "airbomb" && rl.CheckCollisionRecs(gs.Level.Level[gs.Level.RoomNum].etc[a].crec, gs.Level.Level[gs.Level.RoomNum].enemies[num].crec) {
					gs.Level.Level[gs.Level.RoomNum].enemies[num].hp = 0
					cntr := gs.Level.Level[gs.Level.RoomNum].enemies[num].cnt
					addkill(num)
					gs.Level.Level[gs.Level.RoomNum].enemies[num].off = true
					makeFX(2, cntr)
				}
			}
		}
		//CHECK FIRETRAIL COLLIS
		if gs.Player.Mods.firetrail && !gs.Level.Level[gs.Level.RoomNum].enemies[num].fly {
			for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].etc); a++ {
				if gs.Level.Level[gs.Level.RoomNum].etc[a].name == "flamefiretrail" && gs.Level.Level[gs.Level.RoomNum].enemies[num].hppause == 0 && rl.CheckCollisionRecs(gs.Level.Level[gs.Level.RoomNum].etc[a].rec, gs.Level.Level[gs.Level.RoomNum].enemies[num].crec) {
					gs.Level.Level[gs.Level.RoomNum].enemies[num].hppause = gs.Core.Fps / 2
					gs.Level.Level[gs.Level.RoomNum].enemies[num].hp--
					if gs.Level.Level[gs.Level.RoomNum].enemies[num].hp <= 0 {
						cntr := gs.Level.Level[gs.Level.RoomNum].enemies[num].cnt
						addkill(num)
						gs.Level.Level[gs.Level.RoomNum].enemies[num].off = true
						makeFX(2, cntr)
					} else {
						playenemyhit()
					}
				}
			}
		}
		//PLAYER ENEMY AREC COLLIS
		if !gs.Player.Mods.invisible {
			if rl.CheckCollisionRecs(gs.Player.Pl.rec, gs.Level.Level[gs.Level.RoomNum].enemies[num].arec) {
				if gs.Level.Level[gs.Level.RoomNum].enemies[num].name == "ghost" {
					gs.Level.Level[gs.Level.RoomNum].enemies[num].velX, gs.Level.Level[gs.Level.RoomNum].enemies[num].velY = moveFollow(gs.Level.Level[gs.Level.RoomNum].enemies[num].cnt, gs.Player.Pl.cnt, gs.Level.Level[gs.Level.RoomNum].enemies[num].vel)
				}
			}
		}

		switch gs.Level.Level[gs.Level.RoomNum].enemies[num].name {

		case "mushroom":
			if gs.Level.Level[gs.Level.RoomNum].enemies[num].T1 == 0 {
				gs.Level.Level[gs.Level.RoomNum].enemies[num].T1 = rI32(int(gs.Core.Fps*2), int(gs.Core.Fps*5))
				zproj := xproj{}
				zproj.img = gs.Render.MushBull.recTL
				zproj.onoff = true
				zproj.cnt = gs.Level.Level[gs.Level.RoomNum].enemies[num].cnt
				zproj.vel = bsU / 5
				siz := bsU2
				zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
				zproj.velx, zproj.vely = moveFollow(gs.Level.Level[gs.Level.RoomNum].enemies[num].cnt, gs.Player.Pl.cnt, zproj.vel)
				zproj.fade = 1
				zproj.col = ranRed()
				zproj.name = "mushbull"
				gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)
			}

		case "rock":
			if gs.Level.Level[gs.Level.RoomNum].enemies[num].T1 == 0 && gs.Level.Level[gs.Level.RoomNum].enemies[num].spawnN > 0 {
				gs.Level.Level[gs.Level.RoomNum].enemies[num].spawnN--
				gs.Level.Level[gs.Level.RoomNum].enemies[num].T1 = rI32(int(gs.Core.Fps*2), int(gs.Core.Fps*5))
				zen := xenemy{}
				zen = gs.Enemies.EnRock
				zen.spawnN = 0
				zen.rec = gs.Level.Level[gs.Level.RoomNum].enemies[num].rec
				zen.crec = zen.rec
				zen.crec.Y += zen.crec.Height / 3
				zen.crec.Height -= zen.crec.Height / 3
				zen.velX = rF32(-zen.vel, zen.vel)
				zen.velY = rF32(-zen.vel, zen.vel)
				gs.Level.Level[gs.Level.RoomNum].enemies = append(gs.Level.Level[gs.Level.RoomNum].enemies, zen)
			}
			if roll18() == 18 {
				gs.Level.Level[gs.Level.RoomNum].enemies[num].velX = rF32(-gs.Level.Level[gs.Level.RoomNum].enemies[num].vel, gs.Level.Level[gs.Level.RoomNum].enemies[num].vel)
				gs.Level.Level[gs.Level.RoomNum].enemies[num].velY = rF32(-gs.Level.Level[gs.Level.RoomNum].enemies[num].vel, gs.Level.Level[gs.Level.RoomNum].enemies[num].vel)
			}
		case "slime":
			if gs.Level.Level[gs.Level.RoomNum].enemies[num].T1 == 0 {
				gs.Level.Level[gs.Level.RoomNum].enemies[num].T1 = rI32(int(gs.Core.Fps/2), int(gs.Core.Fps*2))
				zblok := xblok{}
				zblok.rec = gs.Level.Level[gs.Level.RoomNum].enemies[num].rec
				zblok.crec = zblok.rec
				zblok.crec.X += zblok.crec.Width / 8
				zblok.crec.Y += zblok.crec.Width / 8
				zblok.crec.Width -= zblok.crec.Width / 4
				zblok.crec.Height -= zblok.crec.Height / 4
				zblok.cnt = gs.Level.Level[gs.Level.RoomNum].enemies[num].cnt
				zblok.img = gs.Render.Splats[rInt(0, len(gs.Render.Splats))]
				zblok.color = ranGreen()
				zblok.fade = 0.7
				zblok.onoff = true
				zblok.name = "slimetrail"
				gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
			}
			if roll36() == 36 {
				gs.Level.Level[gs.Level.RoomNum].enemies[num].velX = rF32(-gs.Level.Level[gs.Level.RoomNum].enemies[num].vel, gs.Level.Level[gs.Level.RoomNum].enemies[num].vel)
				gs.Level.Level[gs.Level.RoomNum].enemies[num].velY = rF32(-gs.Level.Level[gs.Level.RoomNum].enemies[num].vel, gs.Level.Level[gs.Level.RoomNum].enemies[num].vel)
			}
		case "spikehog":
			if roll18() == 18 {
				gs.Level.Level[gs.Level.RoomNum].enemies[num].velX = 0
				gs.Level.Level[gs.Level.RoomNum].enemies[num].velY = 0
				if flipcoin() {
					gs.Level.Level[gs.Level.RoomNum].enemies[num].velX = gs.Level.Level[gs.Level.RoomNum].enemies[num].vel
					if flipcoin() {
						gs.Level.Level[gs.Level.RoomNum].enemies[num].velX *= -1
					}
				} else {
					gs.Level.Level[gs.Level.RoomNum].enemies[num].velY = gs.Level.Level[gs.Level.RoomNum].enemies[num].vel
					if flipcoin() {
						gs.Level.Level[gs.Level.RoomNum].enemies[num].velY *= -1
					}
				}
			}
		}

	} else {
		switch gs.Level.Level[gs.Level.RoomNum].enemies[num].name {

		case "ghost", "slime", "rock", "mushroom":
			gs.Level.Level[gs.Level.RoomNum].enemies[num].velX = rF32(-gs.Level.Level[gs.Level.RoomNum].enemies[num].vel, gs.Level.Level[gs.Level.RoomNum].enemies[num].vel)
			gs.Level.Level[gs.Level.RoomNum].enemies[num].velY = rF32(-gs.Level.Level[gs.Level.RoomNum].enemies[num].vel, gs.Level.Level[gs.Level.RoomNum].enemies[num].vel)
		case "spikehog":
			gs.Level.Level[gs.Level.RoomNum].enemies[num].velX = 0
			gs.Level.Level[gs.Level.RoomNum].enemies[num].velY = 0
			if flipcoin() {
				gs.Level.Level[gs.Level.RoomNum].enemies[num].velX = gs.Level.Level[gs.Level.RoomNum].enemies[num].vel
				if flipcoin() {
					gs.Level.Level[gs.Level.RoomNum].enemies[num].velX *= -1
				}
			} else {
				gs.Level.Level[gs.Level.RoomNum].enemies[num].velY = gs.Level.Level[gs.Level.RoomNum].enemies[num].vel
				if flipcoin() {
					gs.Level.Level[gs.Level.RoomNum].enemies[num].velY *= -1
				}
			}
		case "bat", "rabbit1":
			if gs.Level.Level[gs.Level.RoomNum].enemies[num].velX > 0 {
				gs.Level.Level[gs.Level.RoomNum].enemies[num].velX = rF32(-gs.Level.Level[gs.Level.RoomNum].enemies[num].vel, 0)
			} else if gs.Level.Level[gs.Level.RoomNum].enemies[num].velX < 0 {
				gs.Level.Level[gs.Level.RoomNum].enemies[num].velX = rF32(0, gs.Level.Level[gs.Level.RoomNum].enemies[num].vel)
			}

			if gs.Level.Level[gs.Level.RoomNum].enemies[num].velY > 0 {
				gs.Level.Level[gs.Level.RoomNum].enemies[num].velY = rF32(-gs.Level.Level[gs.Level.RoomNum].enemies[num].vel, 0)
			} else if gs.Level.Level[gs.Level.RoomNum].enemies[num].velY < 0 {
				gs.Level.Level[gs.Level.RoomNum].enemies[num].velY = rF32(0, gs.Level.Level[gs.Level.RoomNum].enemies[num].vel)
			}
		}
	}

	//FIND DIREC FOR ANIM
	if getabs(gs.Level.Level[gs.Level.RoomNum].enemies[num].velX) > getabs(gs.Level.Level[gs.Level.RoomNum].enemies[num].velY) {
		if gs.Level.Level[gs.Level.RoomNum].enemies[num].velX > 0 {
			gs.Level.Level[gs.Level.RoomNum].enemies[num].direc = 2
		} else {
			gs.Level.Level[gs.Level.RoomNum].enemies[num].direc = 4
		}
	} else {
		if gs.Level.Level[gs.Level.RoomNum].enemies[num].velY > 0 {
			gs.Level.Level[gs.Level.RoomNum].enemies[num].direc = 3
		} else {
			gs.Level.Level[gs.Level.RoomNum].enemies[num].direc = 1
		}
	}

}

func movebloks() { //MARK:MOVE BLOKS

	for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].movBloks); a++ {

		if checkMoveBlok(a) {

			gs.Level.Level[gs.Level.RoomNum].movBloks[a].cnt.X += gs.Level.Level[gs.Level.RoomNum].movBloks[a].velX
			gs.Level.Level[gs.Level.RoomNum].movBloks[a].cnt.Y += gs.Level.Level[gs.Level.RoomNum].movBloks[a].velY
			gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec = rl.NewRectangle(gs.Level.Level[gs.Level.RoomNum].movBloks[a].cnt.X-gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec.Width/2, gs.Level.Level[gs.Level.RoomNum].movBloks[a].cnt.Y-gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec.Height/2, gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec.Width, gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec.Height)

		} else {
			switch gs.Level.Level[gs.Level.RoomNum].movBloks[a].movType {
			case 1, 2: //LR UD
				if gs.Level.Level[gs.Level.RoomNum].movBloks[a].velX > 0 {
					gs.Level.Level[gs.Level.RoomNum].movBloks[a].velX = -gs.Level.Level[gs.Level.RoomNum].movBloks[a].velX
				} else if gs.Level.Level[gs.Level.RoomNum].movBloks[a].velX < 0 {
					gs.Level.Level[gs.Level.RoomNum].movBloks[a].velX = getabs(gs.Level.Level[gs.Level.RoomNum].movBloks[a].velX)
				}
				if gs.Level.Level[gs.Level.RoomNum].movBloks[a].velY > 0 {
					gs.Level.Level[gs.Level.RoomNum].movBloks[a].velY = -gs.Level.Level[gs.Level.RoomNum].movBloks[a].velY
				} else if gs.Level.Level[gs.Level.RoomNum].movBloks[a].velY < 0 {
					gs.Level.Level[gs.Level.RoomNum].movBloks[a].velY = getabs(gs.Level.Level[gs.Level.RoomNum].movBloks[a].velY)
				}
			}
		}

		if rl.CheckCollisionRecs(gs.Player.Pl.crec, gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec) {
			gs.Level.Level[gs.Level.RoomNum].movBloks[a].cnt.X -= gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec.Width * 2
			gs.Level.Level[gs.Level.RoomNum].movBloks[a].cnt.Y -= gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec.Width * 2
			gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec = rl.NewRectangle(gs.Level.Level[gs.Level.RoomNum].movBloks[a].cnt.X-gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec.Width/2, gs.Level.Level[gs.Level.RoomNum].movBloks[a].cnt.Y-gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec.Height/2, gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec.Width, gs.Level.Level[gs.Level.RoomNum].movBloks[a].rec.Height)
		}
	}

}

// MARK: ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC
func savesettings() { //MARK:SAVE SETTINGS

	f, err := os.Create("etc/st.000")
	if err != nil {
		fmt.Println(err)
		return
	}

	settingsTXT := ""

	if gs.UI.HpBarsOn {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}
	if gs.UI.ScanLinesOn {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}
	if gs.UI.ArtifactsOn {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}
	if gs.Render.ShaderOn {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}
	if gs.Player.PlatkrecOn {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}
	if gs.UI.Invincible {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}
	if gs.Input.UseController {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}
	if gs.Audio.MusicOn {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}

	musicTXT := fmt.Sprint(gs.Audio.BgMusicNum)
	settingsTXT = settingsTXT + musicTXT + ","

	voltxt := fmt.Sprintf("%.0f", gs.Audio.Volume*10)
	settingsTXT = settingsTXT + voltxt + ","

	if gs.Level.Hardcore {
		settingsTXT = settingsTXT + "1"
	} else {
		settingsTXT = settingsTXT + "0"
	}

	_, err = f.WriteString(settingsTXT)
	if err != nil {
		fmt.Printf("Failed to write settings: %s", err)
	}
}

func addtime() { //MARK:ADD TIME

	totaltime := gs.Level.MinsEND*60 + gs.Level.SecsEND

	checktime := totaltime
	canadd := false
	changenum := 0

	for i := 0; i < len(gs.Timing.Times); i++ {
		if checktime < gs.Timing.Times[i] {
			canadd = true
			checktime = gs.Timing.Times[i]
			changenum = i
		}
	}
	if canadd {
		gs.Timing.Times[changenum] = totaltime
		sort.Ints(gs.Timing.Times)
		gs.Timing.BestTime = true
		savetimes()
	}

}
func restartgame() { //MARK: RESTART GAME

	for a := 0; a < len(gs.Level.Level); a++ {
		gs.Level.Level[a].etc = nil
		gs.Level.Level[a].enemies = nil
		gs.Level.Level[a].doorExitRecs = nil
		gs.Level.Level[a].doorSides = nil
		gs.Level.Level[a].floor = nil
		gs.Level.Level[a].innerBloks = nil
		gs.Level.Level[a].movBloks = nil
		gs.Level.Level[a].nextRooms = nil
		gs.Level.Level[a].spikes = nil
		gs.Level.Level[a].visited = false
		gs.Level.Level[a].walls = nil
	}

	gs.Player.Kills = xkills{}
	gs.Player.Inven = []xblok{}
	gs.Player.Mods = xmod{}
	gs.FX.AirstrikeT = 0
	gs.FX.AirstrikeOn = false
	gs.FX.Fx = nil
	gs.Player.PlProj = nil
	gs.Enemies.EnProj = nil
	gs.Level.Flipcam = false
	gs.Render.Cam2.Rotation = 0

	gs.Player.Pl.armor = 0
	gs.Player.Pl.coins = 0
	gs.Player.Mods.armorN = 0
	gs.Player.Pl.armorMax = 0
	gs.Player.Pl.armor = 0
	gs.FX.FloodRec.Y = gs.Core.ScrHF32 + bsU

	gs.Player.Pl.atk, gs.Player.Pl.slide, gs.Player.Pl.escape, gs.Player.Pl.revived, gs.Player.Pl.poison = false, false, false, false, false

	gs.Level.Night = false

	gs.Level.Level = nil
	gs.Render.Shader2On = false
	gs.Render.Shader3On = false

	gs.Level.Levelnum = 1

	makeplayer()
	makelevel()

	gs.Player.Pl.cnt = gs.Core.Cnt
	gs.Level.RoomNum = 0

	gs.UI.IntroCount = true
	gs.UI.IntroT3 = gs.Core.Fps * 3

	gs.UI.Intro = true
	gs.Player.StartdmgT = gs.Core.Fps * 7
	rl.PlaySound(gs.Audio.Sfx[13])

}
func savetimes() { //MARK: SAVE TIMES

	f, err := os.Create("etc/sc.000")
	if err != nil {
		fmt.Println(err)
		return
	}

	scoresTXT := ""
	for i := 0; i < len(gs.Timing.Times); i++ {
		if i == len(gs.Timing.Times)-1 {
			scoresTXT = scoresTXT + fmt.Sprint(gs.Timing.Times[i])
		} else {
			scoresTXT = scoresTXT + fmt.Sprint(gs.Timing.Times[i]) + ","
		}
	}

	_, err = f.WriteString(scoresTXT)
	if err != nil {
		fmt.Printf("Failed to write times: %s", err)
	}
}

func exitgame() { //MARK: EXIT GAME

	if gs.UI.OptionsChange {
		savesettings()
	}
	savetimes()

	rl.EndDrawing()
	unload()
	rl.CloseAudioDevice()
	rl.CloseWindow()

}
func addshopitem(num int) { //MARK: ADD SHOP ITEM

	//CHECK FOR SAME ADD TO INVEN
	sold := false
	if len(gs.Player.Inven) > 0 {
		foundsame := false
		for a := 0; a < len(gs.Player.Inven); a++ {
			if gs.Shop.ShopItems[num].name == gs.Player.Inven[a].name {

				switch gs.Shop.ShopItems[num].name {
				case "fireworks":
					gs.Player.Pl.coins++
					txtSold("fireworks")
					sold = true
				case "peace":
					gs.Player.Pl.coins++
					txtSold("peace")
					sold = true
				case "anchor":
					gs.Player.Pl.coins++
					txtSold("anchor")
					sold = true
				case "recharge":
					gs.Player.Pl.coins++
					txtSold("recharge")
					sold = true
				case "orbital":
					if gs.Player.Mods.orbitalN < gs.Player.Max.orbital {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("orbital")
						sold = true
					}
				case "coffee":
					if gs.Player.Mods.coffeeN < gs.Player.Max.coffee {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("coffee")
						sold = true
					}
				case "invisible":
					gs.Player.Pl.coins++
					txtSold("invisible")
					sold = true
				case "health potion":
					if gs.Player.Mods.hppotionN < gs.Player.Max.hppotion {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("health potion")
						sold = true
					}
				case "firetrail":
					if gs.Player.Mods.firetrailN < gs.Player.Max.firetrail {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("firetrail")
						sold = true
					}
				case "map":
					gs.Player.Pl.coins++
					txtSold("map")
					sold = true
				case "apple":
					if gs.Player.Mods.appleN < gs.Player.Max.apple {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("apple")
						sold = true
					}
				case "key":
					if gs.Player.Mods.keyN < gs.Player.Max.key {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("key")
						sold = true
					}
				case "bounce":
					if gs.Player.Mods.bounceN < gs.Player.Max.bounce {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("bounce")
						sold = true
					}
				case "fireball":
					if gs.Player.Mods.fireballN < gs.Player.Max.fireball {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("fireball")
						sold = true
					}
				case "throwing axe":
					if gs.Player.Mods.axeN < gs.Player.Max.axe {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("axe")
						sold = true
					}
				}
				foundsame = true
			}
		}
		if !foundsame {
			gs.Player.Inven = append(gs.Player.Inven, gs.Shop.ShopItems[num])
			rl.PlaySound(gs.Audio.Sfx[8])
		}
	} else {
		gs.Player.Inven = append(gs.Player.Inven, gs.Shop.ShopItems[num])
		rl.PlaySound(gs.Audio.Sfx[8])
	}

	//UP MODS
	if !sold {
		switch gs.Shop.ShopItems[num].name {
		case "fireworks":
			gs.Player.Mods.fireworks = true
		case "peace":
			gs.Player.Mods.peace = true
		case "anchor":
			gs.Player.Mods.anchor = true
		case "recharge":
			gs.Player.Mods.recharge = true
		case "orbital":
			gs.Player.Mods.orbital = true
			if gs.Player.Mods.orbitalN < gs.Player.Max.orbital {
				gs.Player.Mods.orbitalN++
				if gs.Player.Mods.orbitalN == 1 {
					gs.Player.Pl.orbital1 = rl.NewVector2(gs.Player.Pl.cnt.X+bsU4, gs.Player.Pl.cnt.Y+bsU4)
				}
				if gs.Player.Mods.orbitalN == 2 {
					gs.Player.Pl.orbital2 = rl.NewVector2(gs.Player.Pl.cnt.X-bsU7, gs.Player.Pl.cnt.Y-bsU7)
				}
			}
		case "coffee":
			if gs.Player.Mods.coffeeN < gs.Player.Max.coffee {
				gs.Player.Mods.coffeeN++
				gs.Player.Pl.vel++
			}
		case "invisible":
			gs.Player.Mods.invisible = true
		case "health potion":
			gs.Player.Mods.hppotion = true
			if gs.Player.Mods.hppotionN < gs.Player.Max.hppotion {
				gs.Player.Mods.hppotionN++
			}
		case "firetrail":
			gs.Player.Mods.firetrail = true
			if gs.Player.Mods.firetrailN < gs.Player.Max.firetrail {
				gs.Player.Mods.firetrailN++
			}
		case "bounce":
			if gs.Player.Mods.bounceN < gs.Player.Max.bounce {
				gs.Player.Mods.bounceN++
			}
		case "map":
			gs.Player.Mods.exitmap = true
		case "apple":
			gs.Player.Mods.apple = true
			if gs.Player.Mods.appleN < gs.Player.Max.apple {
				gs.Player.Mods.appleN++
			}
		case "key":
			gs.Player.Mods.key = true
			if gs.Player.Mods.keyN < gs.Player.Max.key {
				gs.Player.Mods.keyN++
			}
		case "fireball":
			gs.Player.Mods.fireball = true
			if gs.Player.Mods.fireballN < gs.Player.Max.fireball {
				gs.Player.Mods.fireballN++
			}
		case "santa":
			gs.Player.Mods.santa = true
			gs.Player.Mods.santaT = rI32(7, 21) * gs.Core.Fps
		case "throwing axe":
			gs.Player.Mods.axe = true
			if gs.Player.Mods.axeN < gs.Player.Max.axe {
				gs.Player.Mods.axeN++
				gs.Player.Mods.axeT = (int32(gs.Player.Max.axe) * gs.Core.Fps) - (int32(gs.Player.Mods.axeN) * gs.Core.Fps)
			}
			if gs.Player.Mods.axeT < gs.Core.Fps {
				gs.Player.Mods.axeT = gs.Core.Fps
			}

		}
	}
}

func playenemyhit() { //MARK:PLAYER ENEMY HIT SOUND
	if flipcoin() {
		rl.PlaySound(gs.Audio.Sfx[1])
	} else {
		rl.PlaySound(gs.Audio.Sfx[2])
	}
}
func updownswitch(x32, y32 int32, siz, value float32, numType int) float32 { //MARK:UP DOWN SWITCH

	x := float32(x32)
	y := float32(y32)

	rec := rl.NewRectangle(x, y, siz*3, siz)

	rec2 := rec
	rec2.Width = rec2.Width / 3

	rec3 := rec2
	rec3.X += rec3.Width * 2
	rec3.Width -= 2
	rec3.X += 2
	rec2.Width -= 2

	rl.DrawRectangleRec(rec, rl.Black)
	rl.DrawRectangleRec(rec2, rl.Red)
	rl.DrawRectangleRec(rec3, rl.Green)

	rl.DrawRectangleLinesEx(rec, 1, rl.White)

	switch numType {
	case 2:
		txt := fmt.Sprint(value)
		txtlen := rl.MeasureText(txt, gs.UI.TxtSize)

		rl.DrawText(txt, (rec.ToInt32().X+rec.ToInt32().Width/2)-txtlen/2, rec.ToInt32().Y+1, gs.UI.TxtSize, rl.White)

	case 1:
		txt := fmt.Sprintf("%.0f", value*10)
		txtlen := rl.MeasureText(txt, gs.UI.TxtSize)
		rl.DrawText(txt, (rec.ToInt32().X+rec.ToInt32().Width/2)-txtlen/2, rec.ToInt32().Y+1, gs.UI.TxtSize, rl.White)
		upvolume()
	}

	return value
}

func destroyPowerupBlok(blokNum int) { //MARK:DESTROY POWERUP BLOK

	makeFX(1, gs.Level.Level[gs.Level.RoomNum].etc[blokNum].cnt)

	choose := rInt(1, 36)
	//choose = 6

	zblok := makeBlokGeneric(bsU+bsU/2, gs.Level.Level[gs.Level.RoomNum].etc[blokNum].cnt)
	zblok.onoff = true
	zblok.numof = 1

	switch choose {
	case 35: //MARIO
		zblok.name = "mario"
		zblok.desc = "collect extra coins"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[51]
	case 34: //MR CARROT
		zblok.name = "mr carrot"
		zblok.desc = "your friend the root vegetable"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[50]
	case 33: //FIREWORKS
		zblok.name = "fireworks"
		zblok.desc = "shoot fireworks when activating powerup block"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[49]
	case 32: //AIR STRIKE
		zblok.name = "air strike"
		zblok.desc = "support from above at random intervals"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[48]
	case 31: //ALIEN
		zblok.name = "mr alien"
		zblok.desc = "a strange companion"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[47]
	case 30: //PEACE
		zblok.name = "peace"
		zblok.desc = "take no damage for 2 seconds on entering room"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[46]
	case 29: //CAKE
		zblok.name = "birthday cake"
		zblok.desc = "get a random present"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[45]
	case 28: //FLOOD
		zblok.name = "fish"
		zblok.desc = "a not quite biblical flood"
		zblok.color = ranCyan()
		zblok.img = gs.Render.Etc[44]
	case 27: //CHERRY
		zblok.name = "cherry"
		zblok.desc = "coin jackpot - random coin amount added"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[43]
	case 26: //SOCKS
		zblok.name = "moldy socks"
		zblok.desc = "leaves a trail of damaging footprints"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[42]
	case 25: //UMBRELLA
		zblok.name = "umbrella"
		zblok.desc = "rain when you don't need it"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[40]
	case 24: //ANCHOR
		zblok.name = "anchor"
		zblok.desc = "enemies pause for 2 seconds on entering room"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[39]
	case 23: //RECHARGE
		zblok.name = "recharge"
		zblok.desc = "only works with armor - recharges 1 armor every 2 rooms"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[12]
	case 22: //ARMOR
		zblok.name = "armor"
		zblok.desc = "protection that can be recharged"
		zblok.color = ranCyan()
		zblok.img = gs.Render.Etc[38]
	case 21: //HP RING
		zblok.name = "health ring"
		zblok.desc = "adds another hp heart"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[37]
	case 20: //CHAIN LIGHTNING
		zblok.name = "chain lightning"
		zblok.desc = "chance to damage all enemies on screen"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[36]
	case 19: //ORBITAL
		zblok.name = "orbital"
		zblok.desc = "erratic revolving orbs that damage enemies"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[35]
	case 18: //ATTACK DAMAGE
		zblok.name = "attack damage"
		zblok.desc = "increases damage of sword swing"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[34]
	case 17: //ATTACK RANGE
		zblok.name = "attack range"
		zblok.desc = "increases range of sword swing"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[33]
	case 16: //TELEPORT
		zblok.name = "teleport"
		zblok.desc = "transport to another room - destroyed on use"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[32]
	case 15: //COFFEE
		zblok.name = "coffee"
		zblok.desc = "move faster - collect more = faster movement"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[31]
	case 14: //INVISIBLE
		zblok.name = "invisible"
		zblok.desc = "enemies will not follow you"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[30]
	case 13: //FIRE TRAIL
		zblok.name = "firetrail"
		zblok.desc = "trail of fire - does not effect flying enemies"
		zblok.color = ranOrange()
		zblok.img = gs.Render.Etc[29]
	case 12: //FIREBALL
		zblok.name = "fireball"
		zblok.desc = "fires on attack - collect more = more fireballs"
		zblok.color = ranOrange()
		zblok.img = gs.Render.FireballPlayer.recTL
	case 11: //MAP
		zblok.name = "map"
		zblok.desc = "reveals location of exit room"
		zblok.color = ranGrey()
		zblok.img = gs.Render.Etc[28]
	case 10: //WALLET
		zblok.name = "wallet"
		zblok.desc = "purchase one shop item free - destroyed on use"
		zblok.color = rl.Brown
		zblok.img = gs.Render.Etc[11]
	case 9: //MEDI KIT
		zblok.name = "medi kit"
		zblok.desc = "resurrect from death - destroyed on use"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[10]
	case 8: //PLANT COMPANION
		zblok.name = "mr planty"
		zblok.desc = "a companion to assist"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[9]
	case 7: //APPLE
		zblok.name = "apple"
		zblok.desc = "prevents poisoning - destroyed on use"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[8]
	case 6: //KEY
		zblok.name = "key"
		zblok.desc = "open locked chests"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[7]
	case 5: //ESCAPE VINE
		zblok.name = "vine"
		zblok.desc = "automatically escape room at low hp - destroyed on use"
		zblok.color = rl.DarkGreen
		zblok.img = gs.Render.Etc[6]
	case 4: //BOUNCE PROJECTILE
		zblok.name = "bounce"
		zblok.desc = "projectiles bounce > collect more = more bounces"
		zblok.color = rl.Yellow
		zblok.img = gs.Render.Etc[5]
	case 3: //SANTA
		zblok.name = "santa"
		zblok.desc = "snow when you don't need it"
		zblok.color = rl.White
		zblok.img = gs.Render.Etc[4]
	case 2: //THROWING AXE
		zblok.name = "throwing axe"
		zblok.desc = "fires at interval > collect more = faster fire rate"
		zblok.color = rl.SkyBlue
		zblok.img = gs.Render.Etc[3]
	case 1: //HP POTION
		zblok.name = "health potion"
		zblok.desc = "automatically used when health < 2"
		zblok.color = rl.Red
		zblok.img = gs.Render.Etc[1]

	}

	gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)

}
func collectInven(blokNum int) { //MARK:COLLECT INVENTORY

	//CHECK FOR SAME ADD TO INVEN
	sold := false
	if len(gs.Player.Inven) > 0 {
		foundsame := false
		for a := 0; a < len(gs.Player.Inven); a++ {
			if gs.Level.Level[gs.Level.RoomNum].etc[blokNum].name == gs.Player.Inven[a].name {

				switch gs.Level.Level[gs.Level.RoomNum].etc[blokNum].name {
				case "mr carrot":
					gs.Player.Pl.coins++
					txtSold("mr carrot")
					sold = true
				case "fireworks":
					gs.Player.Pl.coins++
					txtSold("fireworks")
					sold = true
				case "air strike":
					gs.Player.Pl.coins++
					txtSold("air strike")
					sold = true
				case "mr alien":
					gs.Player.Pl.coins++
					txtSold("mr alien")
					sold = true
				case "peace":
					gs.Player.Pl.coins++
					txtSold("peace")
					sold = true
				case "birthday cake":
					if gs.Player.Mods.cakeN < gs.Player.Max.cake {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("birthday cake")
						sold = true
					}
				case "fish":
					gs.Player.Pl.coins++
					txtSold("fish")
					sold = true
				case "cherry":
					if gs.Player.Mods.cherryN < gs.Player.Max.cherry {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("cherry")
						sold = true
					}
				case "socks":
					gs.Player.Pl.coins++
					txtSold("socks")
					sold = true
				case "umbrella":
					gs.Player.Pl.coins++
					txtSold("umbrella")
					sold = true
				case "anchor":
					gs.Player.Pl.coins++
					txtSold("anchor")
					sold = true
				case "recharge":
					gs.Player.Pl.coins++
					txtSold("recharge")
					sold = true
				case "armor":
					if gs.Player.Mods.armorN < gs.Player.Max.armor {
						gs.Player.Inven[a].numof++
						if gs.Player.Pl.armor < gs.Player.Pl.armorMax {
							gs.Player.Pl.armor++
						}
						rl.PlaySound(gs.Audio.Sfx[8])
					} else if gs.Player.Mods.armorN == gs.Player.Max.armor && gs.Player.Pl.armor < gs.Player.Pl.armorMax {
						gs.Player.Pl.armor++
					} else {
						gs.Player.Pl.coins++
						txtSold("armor")
						sold = true
					}
				case "health ring":
					if gs.Player.Mods.hpringN < gs.Player.Max.hpring {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("health ring")
						sold = true
					}
				case "chain lightning":
					gs.Player.Pl.coins++
					txtSold("chain lightning")
					sold = true
				case "orbital":
					if gs.Player.Mods.orbitalN < gs.Player.Max.orbital {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("orbital")
						sold = true
					}
				case "attack damage":
					if gs.Player.Mods.atkdmgN < gs.Player.Max.atkdmg {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("attack damage")
						sold = true
					}
				case "attack range":
					if gs.Player.Mods.atkrangeN < gs.Player.Max.atkrange {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("attack range")
						sold = true
					}
				case "coffee":
					if gs.Player.Mods.coffeeN < gs.Player.Max.coffee {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("coffee")
						sold = true
					}
				case "invisible":
					gs.Player.Pl.coins++
					txtSold("invisible")
					sold = true
				case "health potion":
					if gs.Player.Mods.hppotionN < gs.Player.Max.hppotion {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("health potion")
						sold = true
					}
				case "firetrail":
					if gs.Player.Mods.firetrailN < gs.Player.Max.firetrail {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("firetrail")
						sold = true
					}
				case "map":
					gs.Player.Pl.coins++
					txtSold("map")
					sold = true
				case "wallet":
					gs.Player.Pl.coins++
					txtSold("wallet")
					sold = true
				case "medi kit":
					gs.Player.Pl.coins++
					txtSold("medi kit")
					sold = true
				case "mr planty":
					gs.Player.Pl.coins++
					txtSold("mr planty")
					sold = true
				case "apple":
					if gs.Player.Mods.appleN < gs.Player.Max.apple {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("apple")
						sold = true
					}
				case "key":
					if gs.Player.Mods.keyN < gs.Player.Max.key {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("key")
						sold = true
					}
				case "vine":
					gs.Player.Pl.coins++
					txtSold("vine")
					sold = true
				case "bounce":
					if gs.Player.Mods.bounceN < gs.Player.Max.bounce {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("bounce")
						sold = true
					}
				case "fireball":
					if gs.Player.Mods.fireballN < gs.Player.Max.fireball {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("fireball")
						sold = true
					}
				case "santa":
					gs.Player.Pl.coins++
					txtSold("santa")
					sold = true
				case "throwing axe":
					if gs.Player.Mods.axeN < gs.Player.Max.axe {
						gs.Player.Inven[a].numof++
						rl.PlaySound(gs.Audio.Sfx[8])
					} else {
						gs.Player.Pl.coins++
						txtSold("axe")
						sold = true
					}
				}
				foundsame = true
			}
		}
		if !foundsame && gs.Level.Level[gs.Level.RoomNum].etc[blokNum].name != "teleport" {
			gs.Player.Inven = append(gs.Player.Inven, gs.Level.Level[gs.Level.RoomNum].etc[blokNum])
			rl.PlaySound(gs.Audio.Sfx[8])
		}
	} else {
		if gs.Level.Level[gs.Level.RoomNum].etc[blokNum].name != "teleport" {
			gs.Player.Inven = append(gs.Player.Inven, gs.Level.Level[gs.Level.RoomNum].etc[blokNum])
			rl.PlaySound(gs.Audio.Sfx[8])
		}
	}

	//UP MODS
	if !sold {
		switch gs.Level.Level[gs.Level.RoomNum].etc[blokNum].name {
		case "mario":
			makemario()
			gs.Core.Pause = true
			gs.Mario.MarioOn = true
		case "mr carrot":
			if !gs.Player.Mods.alien && !gs.Player.Mods.planty {
				gs.Player.Mods.carrot = true
				gs.Companions.MrCarrot.rec = rl.NewRectangle(gs.Player.Pl.cnt.X-gs.Companions.MrCarrot.rec.Width/2, gs.Player.Pl.cnt.Y-gs.Companions.MrCarrot.rec.Width/2, gs.Companions.MrCarrot.rec.Width, gs.Companions.MrCarrot.rec.Width)
			} else {
				txtCompanion()
			}
		case "fireworks":
			gs.Player.Mods.fireworks = true
		case "air strike":
			gs.Player.Mods.airstrike = true
			gs.FX.AirstrikeT = gs.Core.Fps * rI32(3, 8)
		case "mr alien":
			if !gs.Player.Mods.carrot && !gs.Player.Mods.planty {
				gs.Player.Mods.alien = true
				gs.Companions.MrAlien.rec = rl.NewRectangle(gs.Player.Pl.cnt.X-gs.Companions.MrAlien.rec.Width/2, gs.Player.Pl.cnt.Y-gs.Companions.MrAlien.rec.Width/2, gs.Companions.MrAlien.rec.Width, gs.Companions.MrAlien.rec.Width)
			} else {
				txtCompanion()
			}

		case "peace":
			gs.Player.Mods.peace = true
		case "birthday cake":
			if gs.Player.Mods.cakeN < gs.Player.Max.cake {
				gs.Player.Mods.cakeN++
				birthdaycake()
			}
		case "fish":
			gs.Player.Mods.flood = true
			gs.FX.FloodRec = rl.NewRectangle(0, gs.Core.ScrHF32+bsU, gs.Core.ScrWF32, gs.Core.ScrHF32)
			makefish()
		case "cherry":
			if gs.Player.Mods.cherryN < gs.Player.Max.cherry {
				gs.Player.Mods.cherryN++
				num2 := rInt(1, 6)
				gs.Player.Pl.coins += num2
				txtAddCoinMulti()
			}
		case "moldy socks":
			gs.Player.Mods.socks = true
		case "umbrella":
			gs.Player.Mods.umbrella = true
		case "anchor":
			gs.Player.Mods.anchor = true
		case "recharge":
			gs.Player.Mods.recharge = true
		case "armor":
			if gs.Player.Mods.armorN < gs.Player.Max.armor {
				gs.Player.Mods.armorN++
				gs.Player.Pl.armorMax++
				gs.Player.Pl.armor++
			}
		case "health ring":
			if gs.Player.Mods.hpringN < gs.Player.Max.hpring {
				gs.Player.Mods.hpringN++
				gs.Player.Pl.hpmax++
				if gs.Player.Pl.hp < gs.Player.Pl.hpmax {
					gs.Player.Pl.hp++
				}
			}
		case "chain lightning":
			gs.Player.Mods.chainlightning = true
		case "orbital":
			gs.Player.Mods.orbital = true
			if gs.Player.Mods.orbitalN < gs.Player.Max.orbital {
				gs.Player.Mods.orbitalN++
				if gs.Player.Mods.orbitalN == 1 {
					gs.Player.Pl.orbital1 = rl.NewVector2(gs.Player.Pl.cnt.X+bsU4, gs.Player.Pl.cnt.Y+bsU4)
				}
				if gs.Player.Mods.orbitalN == 2 {
					gs.Player.Pl.orbital2 = rl.NewVector2(gs.Player.Pl.cnt.X-bsU7, gs.Player.Pl.cnt.Y-bsU7)
				}
			}
		case "attack damage":
			if gs.Player.Mods.atkdmgN < gs.Player.Max.atkdmg {
				gs.Player.Mods.atkdmgN++
			}
			gs.Player.Pl.atkDMG++
		case "attack range":
			if gs.Player.Mods.atkrangeN < gs.Player.Max.atkrange {
				gs.Player.Mods.atkrangeN++
			}
			gs.Player.Pl.atkrec.X -= bsU
			gs.Player.Pl.atkrec.Y -= bsU
			gs.Player.Pl.atkrec.Width += bsU2
			gs.Player.Pl.atkrec.Height += bsU2
		case "teleport":
			gs.Player.Pl.hppause = gs.Core.Fps * 5
			maketeleport()
			gs.Player.TeleportRoomNum = rInt(0, len(gs.Level.Level))
			gs.Player.TeleportOn = true
			rl.PlaySound(gs.Audio.Sfx[9])
		case "coffee":
			if gs.Player.Mods.coffeeN < gs.Player.Max.coffee {
				gs.Player.Mods.coffeeN++
				gs.Player.Pl.vel++
			}
		case "invisible":
			gs.Player.Mods.invisible = true
		case "health potion":
			gs.Player.Mods.hppotion = true
			if gs.Player.Mods.hppotionN < gs.Player.Max.hppotion {
				gs.Player.Mods.hppotionN++
			}
		case "firetrail":
			gs.Player.Mods.firetrail = true
			if gs.Player.Mods.firetrailN < gs.Player.Max.firetrail {
				gs.Player.Mods.firetrailN++
			}
		case "bounce":
			if gs.Player.Mods.bounceN < gs.Player.Max.bounce {
				gs.Player.Mods.bounceN++
			}
		case "map":
			gs.Player.Mods.exitmap = true
		case "wallet":
			gs.Player.Mods.wallet = true
		case "medi kit":
			gs.Player.Mods.medikit = true
		case "mr planty":
			if !gs.Player.Mods.alien && !gs.Player.Mods.carrot {
				gs.Player.Mods.planty = true
				gs.Companions.MrPlanty.rec = rl.NewRectangle(gs.Player.Pl.cnt.X-gs.Companions.MrPlanty.rec.Width/2, gs.Player.Pl.cnt.Y-gs.Companions.MrPlanty.rec.Width/2, gs.Companions.MrPlanty.rec.Width, gs.Companions.MrPlanty.rec.Width)
			} else {
				txtCompanion()
			}

		case "apple":
			gs.Player.Mods.apple = true
			if gs.Player.Mods.appleN < gs.Player.Max.apple {
				gs.Player.Mods.appleN++
			}
		case "key":
			gs.Player.Mods.key = true
			if gs.Player.Mods.keyN < gs.Player.Max.key {
				gs.Player.Mods.keyN++
			}
		case "vine":
			gs.Player.Mods.vine = true
		case "fireball":
			gs.Player.Mods.fireball = true
			if gs.Player.Mods.fireballN < gs.Player.Max.fireball {
				gs.Player.Mods.fireballN++
			}
		case "santa":
			gs.Player.Mods.santa = true
			gs.Player.Mods.santaT = rI32(7, 21) * gs.Core.Fps
		case "throwing axe":
			gs.Player.Mods.axe = true
			if gs.Player.Mods.axeN < gs.Player.Max.axe {
				gs.Player.Mods.axeN++
				gs.Player.Mods.axeT = (int32(gs.Player.Max.axe) * gs.Core.Fps) - (int32(gs.Player.Mods.axeN) * gs.Core.Fps)
			}
			if gs.Player.Mods.axeT < gs.Core.Fps {
				gs.Player.Mods.axeT = gs.Core.Fps
			}

		}
	}

}
func birthdaycake() { //MARK:BIRTHDAY CAKE

	found := false
	for {

		choose := rInt(1, 6)
		switch choose {
		case 1:
			if gs.Player.Pl.hp < gs.Player.Pl.hpmax {
				gs.Player.Pl.hp = gs.Player.Pl.hpmax
				found = true
			}
		case 2:
			if gs.Player.Pl.hpmax < 10 {
				gs.Player.Pl.hpmax++
				found = true
			}
		case 3:
			if gs.Player.Mods.armorN == 0 {
				if gs.Player.Mods.armorN < gs.Player.Max.armor {
					gs.Player.Mods.armorN++
					gs.Player.Pl.armorMax++
					gs.Player.Pl.armor++
					found = true
				}
			}
		case 4:
			if gs.Player.Mods.armorN > 0 && gs.Player.Mods.armorN < gs.Player.Max.armor {
				gs.Player.Mods.armorN++
				gs.Player.Pl.armorMax++
				gs.Player.Pl.armor++
				found = true
			}
		case 5:
			if gs.Player.Pl.armor > 0 && gs.Player.Pl.armor < gs.Player.Pl.armorMax {
				gs.Player.Pl.armor = gs.Player.Pl.armorMax
				found = true
			}
		case 6:
			num2 := rInt(1, 6)
			gs.Player.Pl.coins += num2
			txtAddCoinMulti()
			found = true
		case 7:
			if gs.Player.Pl.poison {
				gs.Player.Pl.poisonT = 0
				gs.Player.Pl.poisonCollisT = 0
				if gs.Player.Pl.hp < gs.Player.Pl.hpmax {
					gs.Player.Pl.hp++
				}
			}
		}

		if found {
			break
		}

	}

}
func txtHere(txt string, rec rl.Rectangle) { //MARK:TEXT HERE

	ztxt := xtxt{}
	ztxt.txt = txt
	txtlen := rl.MeasureText(txt, 20)
	x := int32(rec.X+rec.Width/2) - txtlen/2
	y := int32(rec.Y - 20)
	ztxt.fade = 1
	ztxt.col = rl.White
	if txt == "poisoned" {
		ztxt.col = rl.Green
	}
	ztxt.x = x
	ztxt.y = y
	ztxt.onoff = true

	gs.UI.GameTxt = append(gs.UI.GameTxt, ztxt)

}
func delInven(name string) { //MARK:DEL INVEN
	clear := false
	for a := 0; a < len(gs.Player.Inven); a++ {
		if name == gs.Player.Inven[a].name {
			gs.Player.Inven[a].numof--
			if gs.Player.Inven[a].numof == 0 {
				clear = true
			}
		}
	}

	if clear {
		clearinven(name)
	}

}
func onoff(x32, y32 int32, siz float32, name bool) bool { //MARK: ON/OFF

	x := float32(x32)
	y := float32(y32)

	rec := rl.NewRectangle(x, y, siz*2, siz)
	rl.DrawRectangleRec(rec, rl.Black)
	rec2 := rec
	rec2.Width = rec2.Width / 2
	if name {
		rec2.X += siz
		rl.DrawRectangleRec(rec2, rl.Green)
	} else {
		rl.DrawRectangleRec(rec2, rl.Red)
	}

	rl.DrawRectangleLinesEx(rec, 1, rl.White)

	if rl.CheckCollisionPointRec(gs.Core.Mousev2cam, rec) {

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			name = !name
		}
	}

	return name

}
func clearinven(name string) { //MARK:CLEAR INVEN

	num := 0
	clear := false

	for a := 0; a < len(gs.Player.Inven); a++ {
		if gs.Player.Inven[a].name == name {
			num = a
			clear = true
		}
	}
	if clear {
		gs.Player.Inven = remBlok(gs.Player.Inven, num)
	}

}
func escapeplayer() { //MARK:ESCAPE PLAYER

	if gs.Player.Pl.cnt.Y > gs.Level.LevRecInner.Y && !gs.Player.Escaped {
		gs.Player.Pl.cnt.Y -= bsU / 4
		if gs.Player.Pl.cnt.X > gs.Core.Cnt.X {
			gs.Player.Pl.cnt.X -= bsU / 4
		} else if gs.Player.Pl.cnt.X < gs.Core.Cnt.X {
			gs.Player.Pl.cnt.X += bsU / 4
		}

		upPlayerRec()
		if gs.Player.Pl.cnt.Y <= gs.Level.LevRecInner.Y {
			gs.Player.Escaped = true
			upRoomChange()
			gs.Player.Pl.hp = 2
		}

	}

	if gs.Player.Escaped {
		if !gs.Player.EscapeRoomFound {
			choose := 0
			for {
				choose = rInt(0, len(gs.Level.Level))
				if choose != gs.Level.RoomNum {
					break
				}
			}
			gs.Level.RoomNum = choose
			gs.Player.EscapeRoomFound = true
		}
		if gs.Player.Pl.cnt.Y < gs.Level.LevRecInner.Y+bsU3 {
			gs.Player.Pl.cnt.Y += bsU / 4
		}
		if gs.Player.Pl.cnt.X > gs.Core.Cnt.X {
			gs.Player.Pl.cnt.X -= bsU / 4
		} else if gs.Player.Pl.cnt.X < gs.Core.Cnt.X {
			gs.Player.Pl.cnt.X += bsU / 4
		}
		upPlayerRec()

		if gs.Player.Pl.cnt.Y >= gs.Level.LevRecInner.Y+bsU3 {
			gs.Player.Pl.escape = false
		}
	}

}
func cleanlevel() { //MARK:CLEAN LEVEL

	//EMPTY ETC BLOKS
	for a := 0; a < len(gs.Level.Level); a++ {

		var blokstoClear []int
		for b := 0; b < len(gs.Level.Level[a].etc); b++ {
			if gs.Level.Level[a].etc[b].name == "" {
				blokstoClear = append(blokstoClear, b)
			}
		}
		if len(blokstoClear) > 0 {
			for b := 0; b < len(blokstoClear); b++ {
				gs.Level.Level[a].etc = remBlok(gs.Level.Level[a].etc, blokstoClear[b])
			}
		}

	}

	//START ROOM CENTER BLOK CLEAR
	num := 0
	clear := false
	checkrec := rl.NewRectangle(gs.Core.Cnt.X-bsU4, gs.Core.Cnt.Y-bsU4, bsU8, bsU8)
	if len(gs.Level.Level[0].innerBloks) > 0 {
		for a := 0; a < len(gs.Level.Level[0].innerBloks); a++ {
			if rl.CheckCollisionRecs(checkrec, gs.Level.Level[0].innerBloks[a].rec) {
				num = a
				clear = true
			}
		}
	}
	if clear {
		gs.Level.Level[0].innerBloks = remBlok(gs.Level.Level[0].innerBloks, num)
	}

	//CLEAR ETC CENTER
	for i := 0; i < len(gs.Level.Level[0].etc); i++ {
		checkrec := rl.NewRectangle(gs.Core.Cnt.X-bsU2, gs.Core.Cnt.Y-bsU2, bsU4, bsU4)
		if rl.CheckCollisionRecs(checkrec, gs.Level.Level[0].etc[i].rec) {
			gs.Level.Level[0].etc[i].onoff = false
		}
	}

	//CLEAR MOVE BLOCKS CENTER
	for i := 0; i < len(gs.Level.Level[0].movBloks); i++ {
		checkrec := rl.NewRectangle(gs.Core.Cnt.X-bsU2, gs.Core.Cnt.Y-bsU2, bsU4, bsU4)
		if rl.CheckCollisionRecs(checkrec, gs.Level.Level[0].movBloks[i].rec) {
			gs.Level.Level[0].movBloks[i].rec.X -= bsU4
		}
	}

}
func hitPL(numEnProj, numType int) { //MARK:HIT PLAYER

	if gs.Player.Pl.hppause == 0 && gs.Player.Pl.peaceT == 0 && gs.Player.StartdmgT == 0 {
		gs.Player.Pl.hppause = gs.Core.Fps * 2
		gs.Player.HpHitY = 0
		gs.Player.HpHitF = 1

		if gs.Player.Pl.armor > 0 {
			gs.Player.Pl.armor--
			gs.Player.Pl.armorHit = true
		} else {
			switch numType {
			case 2: // BURN POISON ENEMY COLLIS WATER
				if !gs.UI.Invincible {
					gs.Player.Pl.hp--

				}
			case 1: //ENEMY PROJECTILE
				if !gs.UI.Invincible {
					gs.Player.Pl.hp -= gs.Enemies.EnProj[numEnProj].dmg

				}
			}
			zblok := xblok{}
			zblok.rec = gs.Player.Pl.rec
			zblok.cnt = gs.Player.Pl.cnt
			zblok.img = gs.Render.Splats[rInt(0, len(gs.Render.Splats))]
			zblok.color = ranRed()
			zblok.fade = 0.5
			zblok.onoff = true
			zblok.name = "playerblood"
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
			if gs.Player.Pl.hp <= 0 && !gs.UI.Invincible {
				gs.Player.Pl.hp = 0
				if gs.Player.Mods.medikit {
					gs.Player.Pl.hp = gs.Player.Pl.hpmax
					gs.Player.Pl.revived = true
					gs.Player.Pl.hppause = gs.Core.Fps * 3
					gs.Player.ReviveY = 0
					gs.Player.ReviveF = 1
					gs.Player.Mods.medikit = false
					clearinven("medi kit")
				} else {
					gs.Level.DiedscrT = gs.Core.Fps * 3
					gs.Core.Pause = true
					gs.Player.Died = true
					gs.Player.DiedRec = rl.NewRectangle(gs.Core.Cnt.X-bsU, gs.Core.Cnt.Y-bsU, bsU2, bsU2)
					gs.Player.DiedIMG = gs.Render.Splats[rInt(0, len(gs.Render.Splats))]
				}

			}
			rl.PlaySound(gs.Audio.Sfx[14])

		}
	}

}

func addkill(num int) { //MARK:ADD KILL
	switch gs.Level.Level[gs.Level.RoomNum].enemies[num].name {
	case "rabbit1":
		gs.Player.Kills.bunnies++
	case "bat":
		gs.Player.Kills.bats++
	case "mushroom":
		gs.Player.Kills.mushrooms++
	case "ghost":
		gs.Player.Kills.ghosts++
	case "spikehog":
		gs.Player.Kills.spikehogs++
	case "rock":
		gs.Player.Kills.rocks++
	case "slime":
		gs.Player.Kills.slimes++
	}
	rl.PlaySound(gs.Audio.Sfx[7])
}
func txtSold(name string) { //MARK:TEXT SOLD
	rl.PlaySound(gs.Audio.Sfx[18])
	ztxt := xtxt{}
	ztxt.onoff = true
	ztxt.col = rl.White
	ztxt.fade = 1
	ztxt.y = int32(gs.Level.LevY+gs.Level.LevW) - bsU5i32
	ztxt.x = int32(gs.Level.LevX+gs.Level.LevW) + bsUi32

	ztxt.txt = name + " max"
	ztxt.txt2 = "extra sold"

	gs.UI.TxtSoldList = append(gs.UI.TxtSoldList, ztxt)

}
func txtCompanion() { //MARK:TEXT COMPANION
	ztxt := xtxt{}
	ztxt.onoff = true
	ztxt.col = rl.White
	ztxt.fade = 1
	ztxt.y = int32(gs.Level.LevY+gs.Level.LevW) - bsU5i32
	ztxt.x = int32(gs.Level.LevX+gs.Level.LevW) + bsUi32

	ztxt.txt = "1 companion"
	ztxt.txt2 = "allowed"

	gs.UI.TxtSoldList = append(gs.UI.TxtSoldList, ztxt)
}
func txtAddCoin() { //MARK:TEXT ADD 1 COIN
	rl.PlaySound(gs.Audio.Sfx[18])
	ztxt := xtxt{}
	ztxt.onoff = true
	ztxt.col = rl.White
	ztxt.fade = 1
	ztxt.y = int32(gs.Level.LevY+gs.Level.LevW) - bsU5i32
	ztxt.x = int32(gs.Level.LevX+gs.Level.LevW) + bsUi32

	ztxt.txt = "+1 coin"

	gs.UI.TxtSoldList = append(gs.UI.TxtSoldList, ztxt)
	gs.Player.Pl.coins++

}
func txtAddCoinMulti() { //MARK:TEXT ADD MULTIPLE COINS
	rl.PlaySound(gs.Audio.Sfx[18])
	ztxt := xtxt{}
	ztxt.onoff = true
	ztxt.col = rl.White
	ztxt.fade = 1
	ztxt.y = int32(gs.Level.LevY+gs.Level.LevW) - bsU5i32
	ztxt.x = int32(gs.Level.LevX+gs.Level.LevW) + bsUi32

	ztxt.txt = "jackpot"

	gs.UI.TxtSoldList = append(gs.UI.TxtSoldList, ztxt)
	gs.Player.Pl.coins++

}

// MARK: UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP
func up() { //MARK:UP

	if !gs.Core.Pause {
		uplevel()
		upplayer()
		timers()
	}

	checkcontroller()

	if gs.Input.KeypressT > 0 {
		gs.Input.KeypressT--
	}
	if gs.UI.OptionT > 0 {
		gs.UI.OptionT--
	}
	if gs.Player.StartdmgT > 0 {
		gs.Player.StartdmgT--
	}

	inp()
	gs.Core.MouseV2 = rl.GetMousePosition()
	gs.Core.Mousev2cam = rl.GetScreenToWorld2D(gs.Core.MouseV2, gs.Render.Cam2)

	if gs.Audio.MusicOn {
		upaudio()
	}

	if gs.UI.FadeBlinkOn {
		if gs.UI.FadeBlink < 0.5 {
			gs.UI.FadeBlink += 0.05
		} else {
			gs.UI.FadeBlinkOn = false
		}
	} else {
		if gs.UI.FadeBlink > 0.1 {
			gs.UI.FadeBlink -= 0.05
		} else {
			gs.UI.FadeBlinkOn = true
		}
	}

	if gs.UI.FadeBlinkOn2 {
		if gs.UI.FadeBlink2 < 0.3 {
			gs.UI.FadeBlink2 += 0.03
		} else {
			gs.UI.FadeBlinkOn2 = false
		}
	} else {
		if gs.UI.FadeBlink2 > 0.1 {
			gs.UI.FadeBlink2 -= 0.03
		} else {
			gs.UI.FadeBlinkOn2 = true
		}
	}

}
func upaudio() { //MARK:UP AUDIO
	rl.UpdateMusicStream(gs.Audio.Music)
}
func upvolume() { //MARK:UP VOLUME

	rl.SetMusicVolume(gs.Audio.Music, gs.Audio.Volume)

	for a := 0; a < len(gs.Audio.Sfx); a++ {
		rl.SetSoundVolume(gs.Audio.Sfx[a], gs.Audio.Volume+0.1)
	}

	rl.SetMasterVolume(gs.Audio.Volume)

}
func upPlayerMods() { //MARK:UP PLAYER MODS

	//AIR STRIKE
	if gs.Player.Mods.airstrike {
		gs.FX.AirstrikeT--
		if gs.FX.AirstrikeT <= 0 {
			gs.FX.AirstrikeT = gs.Core.Fps * rI32(3, 8)
			makeairstrike()
		}

	}

	//FLOOD
	if gs.Player.Mods.flood {
		if gs.Mario.MarioOn {
			gs.FX.FloodRec.Y = gs.Core.ScrHF32 + bsU
		} else if gs.Level.Levelnum == 6 {
			gs.Mario.MarioOn = false
		} else {
			levrecV2WorldtoScreen := rl.GetWorldToScreen2D(rl.NewVector2(gs.Level.LevRecInner.X, gs.Level.LevRecInner.Y), gs.Render.Cam2)
			if gs.FX.FloodRec.Y > levrecV2WorldtoScreen.Y+bsU12 {
				gs.FX.FloodRec.Y -= 4
			} else {
				if gs.FX.FishLR {
					gs.FX.FishV2.X -= 5
					if gs.FX.FishV2.X < -gs.FX.FishSiz {
						gs.FX.FishLR = false
						gs.FX.Fish1 = gs.Render.FishR.recTL
						gs.FX.FishV2.Y = rF32(gs.Core.ScrHF32/3, gs.Core.ScrHF32)
					}
				} else {
					gs.FX.FishV2.X += 7
					if gs.FX.FishV2.X > gs.Core.ScrWF32 {
						gs.FX.FishLR = true
						gs.FX.Fish1 = gs.Render.FishL.recTL
						gs.FX.FishV2.Y = rF32(gs.Core.ScrHF32/3, gs.Core.ScrHF32)
					}
				}
				if gs.FX.Fish2LR {
					gs.FX.Fish2V2.X += 4
					if gs.FX.Fish2V2.X > gs.Core.ScrWF32 {
						gs.FX.Fish2LR = false
						gs.FX.Fish2 = gs.Render.FishL.recTL
						gs.FX.Fish2V2.Y = rF32(gs.Core.ScrHF32/3, gs.Core.ScrHF32)
					}
				} else {
					gs.FX.Fish2V2.X -= 7
					if gs.FX.Fish2V2.X <= 0 {
						gs.FX.Fish2LR = true
						gs.FX.Fish2 = gs.Render.FishR.recTL
						gs.FX.Fish2V2.Y = rF32(gs.Core.ScrHF32/3, gs.Core.ScrHF32)
					}
				}
			}
		}
	}
	//SOCKS
	if gs.Player.Mods.socks {
		sockstimer := 30
		if gs.Core.Frames%sockstimer == 0 {
			zblok := makeBlokGenNoRecNoCntr()
			zblok.fade = 0.7
			zblok.color = ranGreen()
			zblok.cnt = gs.Player.Pl.cnt
			siz := bsU2
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y, siz, siz)
			zblok.cnt = makeCnt(zblok)
			zblok.drec = makeDrec(zblok.rec)
			zblok.img = gs.Render.Etc[52]
			zblok.name = "footprints"
			switch gs.Player.Pl.direc {
			case 2:
				zblok.ro = 90
			case 3:
				zblok.ro = 180
			case 4:
				zblok.ro = 270
			}
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
		}
	}

	//FIRETRAIL
	if gs.Player.Mods.firetrail {

		flametimer := 60
		if gs.Player.Mods.firetrailN == 2 {
			flametimer = 30
		} else if gs.Player.Mods.firetrailN == 3 {
			flametimer = 15
		}

		if gs.Core.Frames%flametimer == 0 {
			zblok := makeBlokGenNoRecNoCntr()
			zblok.color = ranOrange()
			zblok.cnt = gs.Player.Pl.cnt
			siz := bsU2
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y, siz, siz)
			zblok.img = gs.Render.Firetrailanim.recTL
			zblok.name = "flamefiretrail"
			gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
		}

	}

	//HP POTION
	if gs.Player.Mods.hppotion && gs.Player.Pl.hppotionT == 0 {
		if gs.Player.Pl.hp <= 2 {
			gs.Player.Pl.hppotionT = gs.Core.Fps
			gs.Player.Pl.hp = gs.Player.Pl.hpmax
			gs.Player.Mods.hppotionN--
			gs.Player.Pl.hppause = gs.Core.Fps * 3
			if gs.Player.Mods.hppotionN <= 0 {
				gs.Player.Mods.hppotion = false
				clearinven("health potion")
			}
		}
	}

	//ESCAPE VINE
	if gs.Player.Pl.hp == 1 && gs.Player.Mods.vine && gs.Level.Levelnum != 6 {
		rl.PlaySound(gs.Audio.Sfx[27])
		gs.Player.Mods.vine = false
		clearinven("vine")
		gs.Player.Escaped = false
		gs.Player.EscapeRoomFound = false
		gs.Player.Pl.escape = true
		gs.Player.Pl.hppause = gs.Core.Fps * 5
	}

	//SANTA
	if gs.Player.Mods.santa {
		gs.Player.Mods.santaT--
		if gs.Player.Mods.santaT <= 0 {
			if gs.Player.Mods.snowon {
				gs.Player.Mods.snowon = false
			} else {
				makesnow()
				gs.Player.Mods.snowon = true
			}
			gs.Player.Mods.santaT = rI32(7, 21) * gs.Core.Fps
		}
	}
	//AXE
	if gs.Player.Mods.axe {
		gs.Player.Mods.axeT2--
		if gs.Player.Mods.axeT2 <= 0 {
			gs.Player.Mods.axeT2 = gs.Player.Mods.axeT
			makeProjectile("axe")
		}
	}

}

func upRoomChange() { //MARK:UP ROOM CHANGE

	gs.Player.PlProj = nil
	gs.Enemies.EnProj = nil
	gs.FX.Fx = nil

	//FLOOD
	if gs.Player.Mods.flood {
		gs.Player.Pl.underWater = false
		gs.FX.FloodRec = rl.NewRectangle(0, gs.Core.ScrHF32+bsU, gs.Core.ScrWF32, gs.Core.ScrHF32)
		makefish()
	}
	//ARMOR RECHARGE
	if gs.Player.Mods.recharge && gs.Player.Mods.armorN > 0 {
		gs.Player.Pl.rechargeN++
		if gs.Player.Pl.rechargeN == 2 {
			gs.Player.Pl.rechargeN = 0
			if gs.Player.Pl.armor < gs.Player.Pl.armorMax {
				gs.Player.Pl.armor++
			}
		}
	}

	//PEACE PAUSE
	if gs.Player.Mods.peace {
		gs.Player.Pl.peaceT = gs.Core.Fps * 2
	}
	//ANCHOR PAUSE
	if gs.Player.Mods.anchor {
		gs.Level.AnchorT = gs.Core.Fps * 2
	}

}

func uplevel() { //MARK:UP LEVEL

	movebloks()

}
func upplayer() { //MARK:UP PLAYER

	//TIMERS
	if gs.Player.Mods.flood {

		v2 := rl.GetScreenToWorld2D(rl.NewVector2(gs.FX.FloodRec.Y, gs.FX.FloodRec.Y), gs.Render.Cam2)
		if gs.Player.Pl.crec.Y > v2.Y {
			gs.Player.Pl.underWater = true
			gs.Player.Pl.waterT++
			if gs.Player.Pl.waterT == gs.Core.Fps*3 {
				hitPL(0, 2)
				gs.Player.Pl.waterT = 0
			}
		} else {
			gs.Player.Pl.waterT = 0
			gs.Player.Pl.underWater = false
			gs.Player.WaterY = 0
		}
	}
	if gs.Player.Pl.peaceT > 0 {
		gs.Player.Pl.peaceT--
	}
	if gs.Player.Pl.hppotionT > 0 {
		gs.Player.Pl.hppotionT--
	}
	if gs.Player.Pl.hppause > 0 {
		gs.Player.Pl.hppause--
		if gs.Player.Pl.hppause == 1 && gs.Player.Pl.revived {
			gs.Player.Pl.revived = false
		}
		if gs.Player.Pl.hppause == 1 && gs.Player.Pl.armorHit {
			gs.Player.Pl.armorHit = false
		}
	}
	if gs.Player.Pl.poisonCollisT > 0 {
		gs.Player.Pl.poisonCollisT--
	}

	if gs.Player.Pl.poison {
		gs.Player.Pl.poisonT--
		if gs.Player.Pl.poisonT == 0 {
			hitPL(0, 2)
			gs.Player.Pl.poisonCount--
			if gs.Player.Pl.poisonCount == 0 {
				gs.Player.Pl.poisonT = 0
				gs.Player.Pl.poison = false
			} else {
				gs.Player.Pl.poisonT = gs.Core.Fps * 3
			}
		}
	}

	//ESCAPE
	if gs.Player.Pl.escape {
		escapeplayer()
	}

	//SLIDE
	if gs.Player.Pl.slide {
		switch gs.Player.Pl.slideDIR {
		case 1:
			if checkplayermove(1) {
				gs.Player.Pl.cnt.Y -= gs.Player.Pl.vel * 2
			}
		case 2:
			if checkplayermove(2) {
				gs.Player.Pl.cnt.X += gs.Player.Pl.vel * 2
			}
		case 3:
			if checkplayermove(3) {
				gs.Player.Pl.cnt.Y += gs.Player.Pl.vel * 2
			}
		case 4:
			if checkplayermove(4) {
				gs.Player.Pl.cnt.X -= gs.Player.Pl.vel * 2
			}
		}
		upPlayerRec()
		gs.Player.Pl.slideT--
		if gs.Player.Pl.slideT <= 0 {
			gs.Player.Pl.slide = false
		}
	}

	//UP MODS
	upPlayerMods()

	//UP IMG
	switch gs.Player.Pl.direc {
	case 1:
		gs.Player.Pl.img.Y = gs.Render.Knight[1].Y
	case 2:
		gs.Player.Pl.img.Y = gs.Render.Knight[0].Y
	case 3:
		gs.Player.Pl.img.Y = gs.Render.Knight[3].Y
	case 4:
		gs.Player.Pl.img.Y = gs.Render.Knight[2].Y
	}

	if !gs.Player.Pl.move && !gs.Player.Pl.atk {
		gs.Player.Pl.img.X = gs.Player.Pl.imgWalkX
	} else if gs.Player.Pl.move && !gs.Player.Pl.atk {
		if gs.Core.Frames%4 == 0 {
			gs.Player.Pl.img.X += gs.Player.Pl.sizImg
		}
		if gs.Player.Pl.img.X > gs.Player.Pl.imgWalkX+(float32(gs.Player.Pl.framesWalk-1)*gs.Player.Pl.sizImg) {
			gs.Player.Pl.img.X = gs.Player.Pl.imgWalkX
		}
	} else if !gs.Player.Pl.move && gs.Player.Pl.atk {
		if gs.Core.Frames%4 == 0 {
			gs.Player.Pl.img.X += gs.Player.Pl.sizImg
		}
		if gs.Player.Pl.img.X > gs.Player.Pl.imgAtkX+(float32(gs.Player.Pl.framesAtk-1)*gs.Player.Pl.sizImg) {
			gs.Player.Pl.img.X = gs.Player.Pl.imgAtkX
		}
	}

	//FIND NEXT ROOM ON MOVEMENT
	if !gs.Level.RoomChanged {
		if !rl.CheckCollisionPointRec(gs.Player.Pl.cnt, gs.Level.LevRec) {
			gs.Level.RoomChanged = true
			gs.Level.RoomChangedTimer = gs.Core.Fps / 2
			upRoomChange()
			cntCompanion := gs.Player.Pl.cnt
			if gs.Player.Pl.cnt.X <= gs.Level.LevX {
				for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].doorSides); a++ {
					if gs.Level.Level[gs.Level.RoomNum].doorSides[a] == 4 {
						gs.Level.RoomNum = gs.Level.Level[gs.Level.RoomNum].nextRooms[a]
						gs.Level.Level[gs.Level.RoomNum].visited = true
						break
					}
				}
				gs.Player.Pl.cnt.X = gs.Level.LevRecInner.X + gs.Level.LevRecInner.Width - bsU
				cntCompanion = gs.Player.Pl.cnt
				cntCompanion.X -= bsU2
			} else if gs.Player.Pl.cnt.X >= gs.Level.LevX+gs.Level.LevW {
				for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].doorSides); a++ {
					if gs.Level.Level[gs.Level.RoomNum].doorSides[a] == 2 {
						gs.Level.RoomNum = gs.Level.Level[gs.Level.RoomNum].nextRooms[a]
						gs.Level.Level[gs.Level.RoomNum].visited = true
						break
					}
				}
				gs.Player.Pl.cnt.X = gs.Level.LevRecInner.X + bsU
				cntCompanion = gs.Player.Pl.cnt
				cntCompanion.X += bsU2
			} else if gs.Player.Pl.cnt.Y <= gs.Level.LevY {
				for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].doorSides); a++ {
					if gs.Level.Level[gs.Level.RoomNum].doorSides[a] == 1 {
						gs.Level.RoomNum = gs.Level.Level[gs.Level.RoomNum].nextRooms[a]
						gs.Level.Level[gs.Level.RoomNum].visited = true
						break
					}
				}
				gs.Player.Pl.cnt.Y = gs.Level.LevRecInner.Y + gs.Level.LevRecInner.Width - bsU
				cntCompanion = gs.Player.Pl.cnt
				cntCompanion.Y -= bsU2
			} else if gs.Player.Pl.cnt.Y >= gs.Level.LevY+gs.Level.LevW {
				for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].doorSides); a++ {
					if gs.Level.Level[gs.Level.RoomNum].doorSides[a] == 3 {
						gs.Level.RoomNum = gs.Level.Level[gs.Level.RoomNum].nextRooms[a]
						gs.Level.Level[gs.Level.RoomNum].visited = true
						break
					}
				}
				gs.Player.Pl.cnt.Y = gs.Level.LevRecInner.Y + bsU
				cntCompanion = gs.Player.Pl.cnt
				cntCompanion.Y += bsU2
			}

			if gs.Player.Mods.carrot {
				gs.Companions.MrCarrot.rec = rl.NewRectangle(cntCompanion.X-gs.Companions.MrCarrot.rec.Width/2, cntCompanion.Y-gs.Companions.MrCarrot.rec.Width/2, gs.Companions.MrCarrot.rec.Width, gs.Companions.MrCarrot.rec.Width)
			}
			if gs.Player.Mods.alien {
				gs.Companions.MrAlien.rec = rl.NewRectangle(cntCompanion.X-gs.Companions.MrAlien.rec.Width/2, cntCompanion.Y-gs.Companions.MrAlien.rec.Width/2, gs.Companions.MrAlien.rec.Width, gs.Companions.MrAlien.rec.Width)
			}
			if gs.Player.Mods.planty {
				gs.Companions.MrPlanty.rec = rl.NewRectangle(cntCompanion.X-gs.Companions.MrPlanty.rec.Width/2, cntCompanion.Y-gs.Companions.MrPlanty.rec.Width/2, gs.Companions.MrPlanty.rec.Width, gs.Companions.MrPlanty.rec.Width)
			}
		}
	}

	//TIMERS
	if gs.Player.Pl.atkTimer > 0 {
		gs.Player.Pl.atkTimer--
		if gs.Player.Pl.atkTimer == 1 {
			gs.Player.Pl.img.X = gs.Player.Pl.imgWalkX
			gs.Player.Pl.atk = false
		}
	}
}
func upPlayerRec() { //MARK:UP PLAYER REC CENTER CHANGED
	gs.Player.Pl.rec = rl.NewRectangle(gs.Player.Pl.cnt.X-gs.Player.Pl.rec.Width/2, gs.Player.Pl.cnt.Y-gs.Player.Pl.rec.Width/2, gs.Player.Pl.rec.Width, gs.Player.Pl.rec.Width)
	gs.Player.Pl.crec = rl.NewRectangle(gs.Player.Pl.cnt.X-gs.Player.Pl.crec.Width/2, gs.Player.Pl.cnt.Y-gs.Player.Pl.crec.Height/2, gs.Player.Pl.crec.Width, gs.Player.Pl.crec.Height)
	gs.Player.Pl.arec = rl.NewRectangle(gs.Player.Pl.cnt.X-gs.Player.Pl.arec.Width/2, gs.Player.Pl.cnt.Y-gs.Player.Pl.arec.Height/2, gs.Player.Pl.arec.Width, gs.Player.Pl.arec.Height)
	gs.Player.Pl.atkrec = rl.NewRectangle(gs.Player.Pl.cnt.X-gs.Player.Pl.atkrec.Width/2, gs.Player.Pl.cnt.Y-gs.Player.Pl.atkrec.Height/2, gs.Player.Pl.atkrec.Width, gs.Player.Pl.atkrec.Height)
}

// MARK:MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE
func makeaudio() { //MARK:MAKE AUDIO

	gs.Audio.BackMusic = append(gs.Audio.BackMusic, rl.LoadMusicStream("audio/1.ogg"))  //0
	gs.Audio.BackMusic = append(gs.Audio.BackMusic, rl.LoadMusicStream("audio/2.ogg"))  //1
	gs.Audio.BackMusic = append(gs.Audio.BackMusic, rl.LoadMusicStream("audio/16.ogg")) //2

	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/3.ogg"))  //0 PLAYER SWORD
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/4.ogg"))  //1 ENEMY HIT1
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/5.ogg"))  //2 ENEMY HIT2
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/6.ogg"))  //3 SPRING
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/7.ogg"))  //4 SPEAR
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/8.ogg"))  //5 OIL BARREL BURN
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/9.ogg"))  //6 POWER UP BLOCK DESTROY
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/10.ogg")) //7 ENEMY DEATH
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/11.ogg")) //8 ITEM PICKUP
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/12.ogg")) //9 TELEPORT
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/13.ogg")) //10 OIL BARREL EXPLODE
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/14.ogg")) //11 LIGHTNING
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/15.ogg")) //12 SWITCH
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/17.ogg")) //13 COUNTDOWN
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/18.ogg")) //14 PLAYER HIT
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/19.ogg")) //15 STAIRS
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/20.ogg")) //16 SHOP DOOR
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/21.ogg")) //17 CHEST OPENING
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/22.ogg")) //18 COIN ADDED
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/23.ogg")) //19 WIN GAME
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/24.ogg")) //20 GAS TRAP
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/25.ogg")) //21 SHOP PURCHASE
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/26.ogg")) //22d NO MONEY SHOP PURCHASE
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/27.ogg")) //23 POISONED
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/28.ogg")) //24 UNDERWATER
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/29.ogg")) //25 LOCKED CHEST
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/30.ogg")) //26 SPEAR
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/31.ogg")) //27 VINE
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/32.ogg")) //28 AIR STRIKE EPLOSION
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/33.ogg")) //29 BOSS HIT
	gs.Audio.Sfx = append(gs.Audio.Sfx, rl.LoadSound("audio/34.ogg")) //30 CIRCULAR SAW

	upvolume()

}
func makechestitem(num int) { //MARK:MAKE CHEST ITEM

	newcnt := gs.Level.Level[gs.Level.RoomNum].etc[num].cnt
	newcnt.Y -= bsU2
	zblok := makeBlokGeneric(bsU+bsU/2, newcnt)
	choose := rInt(0, len(gs.Render.Gems))
	zblok.onoff = true
	zblok.numof = 1
	zblok.name = "gem"
	zblok.color = rl.White
	zblok.img = gs.Render.Gems[choose]
	switch choose {
	case 0:
		zblok.numCoins = 5
	case 1:
		zblok.numCoins = 6
	case 2:
		zblok.numCoins = 7
	case 3:
		zblok.numCoins = 8
	case 4:
		zblok.numCoins = 9
	case 5:
		zblok.numCoins = 10
	case 6:
		zblok.numCoins = 11
	case 7:
		zblok.numCoins = 12
	}

	gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)
}
func maketimes() { //MARK: MAKE TIMES

	contents, err := os.ReadFile("etc/sc.000")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	txtTimes := strings.Split(string(contents), ",")

	for i := 0; i < len(txtTimes); i++ {
		num, _ := strconv.Atoi(txtTimes[i])
		gs.Timing.Times = append(gs.Timing.Times, num)
		sort.Ints(gs.Timing.Times)
	}

}
func makesettings() { //MARK: MAKE SETTINGS

	contents, err := os.ReadFile("etc/st.000")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	txtSettings := strings.Split(string(contents), ",")

	if txtSettings[0] == "1" {
		gs.UI.HpBarsOn = true
	} else {
		gs.UI.HpBarsOn = false
	}
	if txtSettings[1] == "1" {
		gs.UI.ScanLinesOn = true
	} else {
		gs.UI.ScanLinesOn = false
	}
	if txtSettings[2] == "1" {
		gs.UI.ArtifactsOn = true
	} else {
		gs.UI.ArtifactsOn = false
	}
	if txtSettings[3] == "1" {
		gs.Render.ShaderOn = true
	} else {
		gs.Render.ShaderOn = false
	}
	if txtSettings[4] == "1" {
		gs.Player.PlatkrecOn = true
	} else {
		gs.Player.PlatkrecOn = false
	}
	if txtSettings[5] == "1" {
		gs.UI.Invincible = true
	} else {
		gs.UI.Invincible = false
	}
	if txtSettings[6] == "1" {
		gs.Input.UseController = true
		gs.Input.ControllerOn = true
	} else {
		gs.Input.UseController = true
		gs.Input.ControllerOn = true
	}
	if txtSettings[7] == "1" {
		gs.Audio.MusicOn = true
	} else {
		gs.Audio.MusicOn = false
	}
	if txtSettings[10] == "1" {
		gs.Level.Hardcore = true
	} else {
		gs.Level.Hardcore = false
	}

	//BG MUSIC
	num, _ := strconv.Atoi(txtSettings[8])
	gs.Audio.BgMusicNum = num
	//VOLUME
	num, _ = strconv.Atoi(txtSettings[9])
	gs.Audio.Volume = float32(num) / 10

}
func makeshop() { //MARK: MAKE SHOP

	gs.Shop.ShopItems = nil
	gs.Shop.ShopNum = 0
	countbreak := 100

	//SHOP IMG
	for {
		gs.Level.ShopRoomNum = rInt(0, len(gs.Level.Level))

		zblok := xblok{}
		zblok.name = "shop"
		zblok.solid = true
		zblok.onoff = true
		zblok.img = gs.Render.Etc[14]
		zblok.cnt = gs.Core.Cnt
		zblok.cnt.Y -= bsU8

		siz := bsU8

		canadd := true
		zblok.rec, canadd = findRecPoswithSpacing(siz, bsU4, gs.Level.ShopRoomNum)
		zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
		zblok.crec = zblok.rec
		zblok.crec.X += zblok.crec.Width / 8
		zblok.crec.Width = (zblok.crec.Width / 4) * 3
		zblok.crec2 = zblok.crec
		zblok.crec2.Width += bsU
		zblok.crec2.Height += bsU
		zblok.crec2.X -= bsU / 2
		zblok.crec2.Y -= bsU / 2

		zblok.color = rl.White
		zblok.fade = 1

		siz = bsU / 2
		v1 := rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Height+bsU)
		v2 := v1
		v2.Y += siz
		v2.X += siz / 2
		v3 := v2
		v3.X -= siz

		zblok.v2s = append(zblok.v2s, v1)
		zblok.v2s = append(zblok.v2s, v2)
		zblok.v2s = append(zblok.v2s, v3)
		zblok.v2s = append(zblok.v2s, v1)
		zblok.v2s = append(zblok.v2s, v2)
		zblok.v2s = append(zblok.v2s, v3)

		rl.DrawTriangle(v2, v1, v3, rl.Red)

		countbreak--
		if canadd || countbreak == 0 {
			gs.Level.Level[gs.Level.ShopRoomNum].etc = append(gs.Level.Level[gs.Level.ShopRoomNum].etc, zblok)
			break
		}
	}

	//SHOP ITEMS
	zblok := xblok{}
	zblok.fade = 1
	zblok.onoff = true
	num := 4

	for num > 0 {

		choose := rInt(1, 16)

		switch choose {
		case 15: //FIREWORKS
			zblok.name = "fireworks"
			zblok.desc = "shoot fireworks when activating powerup block"
			zblok.color = rl.White
			zblok.img = gs.Render.Etc[49]
		case 14: //PEACE
			zblok.name = "peace"
			zblok.desc = "take no damage for 2 seconds on entering room"
			zblok.color = rl.White
			zblok.img = gs.Render.Etc[46]
		case 13: //ANCHOR
			zblok.name = "anchor"
			zblok.desc = "enemies pause for 2 seconds on entering room"
			zblok.color = rl.White
			zblok.img = gs.Render.Etc[39]
		case 12: //RECHARGE
			zblok.name = "recharge"
			zblok.desc = "only works with armor - recharges 1 armor every 2 rooms"
			zblok.color = rl.White
			zblok.img = gs.Render.Etc[12]
		case 11: //ORBITAL
			zblok.name = "orbital"
			zblok.desc = "erratic revolving orbs that damage enemies"
			zblok.color = rl.White
			zblok.img = gs.Render.Etc[35]
		case 10: //COFFEE
			zblok.name = "coffee"
			zblok.desc = "move faster - collect more = faster movement"
			zblok.color = rl.White
			zblok.img = gs.Render.Etc[31]
		case 9: //INVISIBLE
			zblok.name = "invisible"
			zblok.desc = "enemies will not follow you"
			zblok.color = rl.White
			zblok.img = gs.Render.Etc[30]
		case 8: //FIRE TRAIL
			zblok.name = "firetrail"
			zblok.desc = "trail of fire - does not effect flying enemies"
			zblok.color = ranOrange()
			zblok.img = gs.Render.Etc[29]
		case 7: //FIREBALL
			zblok.name = "fireball"
			zblok.desc = "fires on attack - collect more = more fireballs"
			zblok.color = ranOrange()
			zblok.img = gs.Render.FireballPlayer.recTL
		case 6: //MAP
			zblok.name = "map"
			zblok.desc = "reveals location of exit room"
			zblok.color = ranGrey()
			zblok.img = gs.Render.Etc[28]
		case 5: //APPLE
			zblok.name = "apple"
			zblok.desc = "prevents poisoning - destroyed on use"
			zblok.color = rl.White
			zblok.img = gs.Render.Etc[8]
		case 4: //SKELETON KEY
			zblok.name = "key"
			zblok.desc = "open locked things"
			zblok.color = rl.White
			zblok.img = gs.Render.Etc[7]
		case 3: //BOUNCE PROJECTILE
			zblok.name = "bounce"
			zblok.desc = "projectiles bounce > collect more = more bounces"
			zblok.color = rl.Yellow
			zblok.img = gs.Render.Etc[5]
		case 2: //THROWING AXE
			zblok.name = "throwing axe"
			zblok.desc = "fires at interval > collect more = faster fire rate"
			zblok.color = rl.SkyBlue
			zblok.img = gs.Render.Etc[3]
		case 1: //HP POTION
			zblok.name = "health potion"
			zblok.desc = "automatically used when health < 2"
			zblok.color = rl.Red
			zblok.img = gs.Render.Etc[1]
		}

		if len(gs.Shop.ShopItems) > 0 {
			canadd := true
			for a := 0; a < len(gs.Shop.ShopItems); a++ {
				if gs.Shop.ShopItems[a].name == zblok.name {
					canadd = false
				}
			}
			if canadd {
				zblok.shopprice = rInt(2, 8)
				gs.Shop.ShopItems = append(gs.Shop.ShopItems, zblok)
				num--
			}
		} else {
			zblok.shopprice = rInt(2, 8)
			gs.Shop.ShopItems = append(gs.Shop.ShopItems, zblok)
			num--
		}

	}

}

func makeProjectileEnemy(num int, cnt rl.Vector2) { //MARK:MAKE PROJECTILE ENEMY

	switch num {

	case 1: //TURRET

		siz := bsU
		zproj := xproj{}
		zproj.name = "ninja"
		zproj.cnt = cnt
		//zproj.cnt.X -= bsU / 2
		//zproj.cnt.Y -= bsU / 2
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.vel = bsU / 4
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		zproj.col = ranCyan()
		zproj.dmg = 1
		zproj.fade = 1
		zproj.img = gs.Render.Etc[19]
		zproj.onoff = true

		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)

	case 9: //BOSS ATK 3
		siz := bsU3
		zproj := xproj{}
		zproj.name = "boss3"
		zproj.cnt = cnt
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.ori = rl.NewVector2(zproj.rec.Width/2, zproj.rec.Height/2)
		zproj.vel = bsU / 2
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		zproj.col = rl.White
		zproj.dmg = 1
		zproj.fade = 1
		zproj.img = gs.Render.Etc[57]
		zproj.onoff = true
		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)

	case 8: //BOSS ATK 2
		siz := bsU3
		zproj := xproj{}
		zproj.name = "boss2"
		zproj.cnt = cnt
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.vel = bsU / 2
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		zproj.col = rl.White
		zproj.dmg = 1
		zproj.fade = 1
		zproj.img = gs.Render.Boss2anim.recTL
		zproj.onoff = true
		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)

	case 7: //BOSS ATK 1
		siz := bsU3
		zproj := xproj{}
		zproj.name = "boss1"
		zproj.cnt = cnt
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.crec = zproj.rec
		zproj.crec.X += zproj.rec.Width / 4
		zproj.crec.Y += zproj.rec.Height / 4
		zproj.crec.Width = zproj.rec.Width / 2
		zproj.crec.Height = zproj.rec.Height / 2
		zproj.vel = bsU / 2
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		zproj.col = rl.White
		zproj.dmg = 1
		zproj.fade = 1
		zproj.img = gs.Render.Boss1anim.recTL
		zproj.onoff = true
		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)

	case 2, 3, 4, 5: //SWITCH ARROWS

		siz := bsU2
		zproj := xproj{}
		zproj.name = "switcharrow"
		zproj.cnt = cnt
		//zproj.cnt.X -= bsU / 2
		//zproj.cnt.Y -= bsU / 2
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.vel = bsU / 2
		if num == 2 {
			zproj.vely = zproj.vel
			zproj.ro = 135
		} else if num == 3 {
			zproj.velx = zproj.vel
			zproj.ro = 45
		} else if num == 4 {
			zproj.vely = -zproj.vel
			zproj.ro = -45
		} else if num == 5 {
			zproj.velx = -zproj.vel
			zproj.ro = -135
		}
		zproj.col = ranOrange()
		zproj.dmg = 1
		zproj.fade = 1
		zproj.img = gs.Render.Etc[24]
		zproj.onoff = true

		gs.Enemies.EnProj = append(gs.Enemies.EnProj, zproj)

	}

}
func makeendlevel() { //MARK: MAKE END LEVEL

	gs.Level.Level = nil

	gs.Level.RoomNum = 0

	//MAKE LEVEL ROOMS

	gs.Level.Bossnum = rInt(0, len(gs.Level.Bosses))

	countedBorderBlocks := false

	floorT := gs.Render.Floortiles[rInt(0, len(gs.Render.Floortiles))]
	gs.Level.WallT = gs.Render.Walltiles[rInt(0, len(gs.Render.Walltiles))]

	zroom := xroom{}
	zroom.floorT = floorT
	zroom.wallT = gs.Level.WallT

	//BOUNDARY WALLS
	x := gs.Level.LevX
	y := gs.Level.LevY

	for x < gs.Level.LevX+gs.Level.LevW {
		if !countedBorderBlocks {
			gs.Level.LevBorderBlokNum++
		}
		zblok := xblok{}
		zblok.fade = 1
		zblok.img = zroom.wallT
		switch gs.Level.Levelnum {
		case 1:
			zblok.color = ranBlue()
		case 2:
			zblok.color = ranBrown()
		case 3:
			zblok.color = ranOrange()
		case 4:
			zblok.color = ranDarkBlue()
		case 5:
			zblok.color = ranCol()
		case 6:
			zblok.color = ranRed()
		}
		zblok.rec = rl.NewRectangle(x, y, gs.Level.BorderWallBlokSiz, gs.Level.BorderWallBlokSiz)
		zroom.walls = append(zroom.walls, zblok)
		zblok.rec.Y = gs.Level.LevY + gs.Level.LevW - gs.Level.BorderWallBlokSiz
		switch gs.Level.Levelnum {
		case 1:
			zblok.color = ranBlue()
		case 2:
			zblok.color = ranBrown()
		case 3:
			zblok.color = ranOrange()
		case 4:
			zblok.color = ranDarkBlue()
		case 5:
			zblok.color = ranCol()
		case 6:
			zblok.color = ranRed()
		}
		zroom.walls = append(zroom.walls, zblok)
		x += gs.Level.BorderWallBlokSiz
	}
	if !countedBorderBlocks {
		countedBorderBlocks = true
	}
	x = gs.Level.LevX
	y = gs.Level.LevY + gs.Level.BorderWallBlokSiz
	for y < gs.Level.LevY+gs.Level.LevW-gs.Level.BorderWallBlokSiz {
		zblok := xblok{}
		zblok.fade = 1
		zblok.img = zroom.wallT
		switch gs.Level.Levelnum {
		case 1:
			zblok.color = ranBlue()
		case 2:
			zblok.color = ranBrown()
		case 3:
			zblok.color = ranOrange()
		case 4:
			zblok.color = ranDarkBlue()
		case 5:
			zblok.color = ranCol()
		case 6:
			zblok.color = ranRed()
		}
		zblok.rec = rl.NewRectangle(x, y, gs.Level.BorderWallBlokSiz, gs.Level.BorderWallBlokSiz)
		zroom.walls = append(zroom.walls, zblok)
		zblok.rec.X = gs.Level.LevX + gs.Level.LevW - gs.Level.BorderWallBlokSiz
		switch gs.Level.Levelnum {
		case 1:
			zblok.color = ranBlue()
		case 2:
			zblok.color = ranBrown()
		case 3:
			zblok.color = ranOrange()
		case 4:
			zblok.color = ranDarkBlue()
		case 5:
			zblok.color = ranCol()
		case 6:
			zblok.color = ranRed()
		}
		zroom.walls = append(zroom.walls, zblok)
		y += gs.Level.BorderWallBlokSiz
	}

	//FLOOR
	x = gs.Level.LevX
	y = gs.Level.LevY
	siz := bsU3
	zblok := xblok{}
	zblok.img = zroom.floorT
	for {
		zblok.rec = rl.NewRectangle(x, y, siz, siz)
		zblok.fade = rF32(0.1, 0.25)
		switch gs.Level.Levelnum {
		case 1:
			zblok.color = ranRed()
		case 2:
			zblok.color = ranDarkGreen()
		case 3:
			zblok.color = ranDarkBlue()
		case 4:
			zblok.color = ranDarkGrey()
		case 5:
			zblok.color = ranDarkBlue()
		case 6:
			zblok.color = ranRed()
		}
		zroom.floor = append(zroom.floor, zblok)

		x += siz
		if x >= gs.Level.LevX+gs.Level.LevW {
			x = gs.Level.LevX
			y += siz
		}
		if y >= gs.Level.LevY+gs.Level.LevW {
			break
		}
	}

	gs.Level.Level = append(gs.Level.Level, zroom)

	//SWITCHES
	if flipcoin() {
		zblok := makeBlokGenNoRecNoCntr()
		canadd := true
		zblok.rec, canadd = findRecPoswithSpacing(bsU2, bsU/2, 0)
		zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
		zblok.name = "switch"
		zblok.numType = roll6()
		if zblok.numType == 6 {
			zblok.numCoins = rInt(3, 11)
		}
		zblok.onoffswitch = flipcoin()
		zblok.color = rl.SkyBlue
		if zblok.onoffswitch {
			zblok.img = gs.Render.Etc[21]
		} else {
			zblok.img = gs.Render.Etc[22]
		}
		if canadd {
			gs.Level.Level[0].etc = append(gs.Level.Level[0].etc, zblok)
		}
	}
	if flipcoin() {
		zblok := makeBlokGenNoRecNoCntr()
		canadd := true
		zblok.rec, canadd = findRecPoswithSpacing(bsU2, bsU/2, 0)
		zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
		zblok.name = "switch"
		zblok.numType = roll6()
		if zblok.numType == 6 {
			zblok.numCoins = rInt(3, 11)
		}
		zblok.onoffswitch = flipcoin()
		zblok.color = rl.SkyBlue
		if zblok.onoffswitch {
			zblok.img = gs.Render.Etc[21]
		} else {
			zblok.img = gs.Render.Etc[22]
		}
		if canadd {
			gs.Level.Level[0].etc = append(gs.Level.Level[0].etc, zblok)
		}
	}

	//SKULLS
	num := rInt(1, 5)
	for {
		zblok := makeBlokGenNoRecNoCntr()
		siz := rF32((bsU/4)*3, bsU+bsU/2)
		canadd := true
		zblok.rec, canadd = findRecPos(siz, 0)
		zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
		zblok.name = "skull"
		zblok.fade = rF32(0.3, 0.6)
		zblok.color = ranGrey()
		zblok.img = gs.Render.Skulls[rInt(0, len(gs.Render.Skulls))]
		if canadd {
			gs.Level.Level[0].etc = append(gs.Level.Level[0].etc, zblok)
			num--
		}
		if num <= 0 {
			break
		}
	}

	//OIL BARRELS
	num = rInt(2, 7)
	for {
		zblok := makeBlokGenNoRecNoCntr()
		canadd := true
		zblok.rec, canadd = findRecPoswithSpacing(bsU2, bsU/2, 0)
		zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
		zblok.solid = true
		zblok.onoff = true
		zblok.name = "oilbarrel"
		zblok.color = rl.DarkGreen
		zblok.img = gs.Render.Etc[20]
		if canadd {
			gs.Level.Level[0].etc = append(gs.Level.Level[0].etc, zblok)
			num--
		}

		if num <= 0 {
			break
		}
	}

	makeInnerBloks()
	makemovebloks()
	makeblades()
	maketurrets()
	makespears()
	makespikes()

	//REMOVE BLOKS COLLIDING WITH BOSS REC
	checkrec := rl.NewRectangle(gs.Core.Cnt.X-bsU4, gs.Level.LevRecInner.Y+bsU2, bsU8, bsU8)
	numcollis := 0
	found := false
	for i := 0; i < len(gs.Level.Level[0].innerBloks); i++ {
		if rl.CheckCollisionRecs(checkrec, gs.Level.Level[0].innerBloks[i].rec) {
			numcollis = i
			found = true
		}
	}
	if found {
		gs.Level.Level[0].innerBloks = remBlok(gs.Level.Level[0].innerBloks, numcollis)
	}

	cleanlevel()

}
func makelevel() { //MARK:MAKE LEVEL

	gs.Level.Level = nil

	//MAKE LEVEL ROOMS
	numRooms := rInt(7, 13)
	orignumRooms := numRooms
	countedBorderBlocks := false

	floorT := gs.Render.Floortiles[rInt(0, len(gs.Render.Floortiles))]
	gs.Level.WallT = gs.Render.Walltiles[rInt(0, len(gs.Render.Walltiles))]

	for numRooms > 0 {
		zroom := xroom{}
		zroom.floorT = floorT
		zroom.wallT = gs.Level.WallT

		//BOUNDARY WALLS
		x := gs.Level.LevX
		y := gs.Level.LevY

		for x < gs.Level.LevX+gs.Level.LevW {
			if !countedBorderBlocks {
				gs.Level.LevBorderBlokNum++
			}
			zblok := xblok{}
			zblok.fade = 1
			zblok.img = zroom.wallT
			switch gs.Level.Levelnum {
			case 1:
				zblok.color = ranBlue()
			case 2:
				zblok.color = ranBrown()
			case 3:
				zblok.color = ranOrange()
			case 4:
				zblok.color = ranDarkBlue()
			case 5:
				zblok.color = ranCol()
			}
			zblok.rec = rl.NewRectangle(x, y, gs.Level.BorderWallBlokSiz, gs.Level.BorderWallBlokSiz)
			zroom.walls = append(zroom.walls, zblok)
			zblok.rec.Y = gs.Level.LevY + gs.Level.LevW - gs.Level.BorderWallBlokSiz
			switch gs.Level.Levelnum {
			case 1:
				zblok.color = ranBlue()
			case 2:
				zblok.color = ranBrown()
			case 3:
				zblok.color = ranOrange()
			case 4:
				zblok.color = ranDarkBlue()
			case 5:
				zblok.color = ranCol()
			}
			zroom.walls = append(zroom.walls, zblok)
			x += gs.Level.BorderWallBlokSiz
		}
		if !countedBorderBlocks {
			countedBorderBlocks = true
		}
		x = gs.Level.LevX
		y = gs.Level.LevY + gs.Level.BorderWallBlokSiz
		for y < gs.Level.LevY+gs.Level.LevW-gs.Level.BorderWallBlokSiz {
			zblok := xblok{}
			zblok.fade = 1
			zblok.img = zroom.wallT
			switch gs.Level.Levelnum {
			case 1:
				zblok.color = ranBlue()
			case 2:
				zblok.color = ranBrown()
			case 3:
				zblok.color = ranOrange()
			case 4:
				zblok.color = ranDarkBlue()
			case 5:
				zblok.color = ranCol()
			}
			zblok.rec = rl.NewRectangle(x, y, gs.Level.BorderWallBlokSiz, gs.Level.BorderWallBlokSiz)
			zroom.walls = append(zroom.walls, zblok)
			zblok.rec.X = gs.Level.LevX + gs.Level.LevW - gs.Level.BorderWallBlokSiz
			switch gs.Level.Levelnum {
			case 1:
				zblok.color = ranBlue()
			case 2:
				zblok.color = ranBrown()
			case 3:
				zblok.color = ranOrange()
			case 4:
				zblok.color = ranDarkBlue()
			case 5:
				zblok.color = ranCol()
			}
			zroom.walls = append(zroom.walls, zblok)
			y += gs.Level.BorderWallBlokSiz
		}

		//FLOOR
		x = gs.Level.LevX
		y = gs.Level.LevY
		siz := bsU3
		zblok := xblok{}
		zblok.img = zroom.floorT
		for {
			zblok.rec = rl.NewRectangle(x, y, siz, siz)
			zblok.fade = rF32(0.1, 0.25)
			switch gs.Level.Levelnum {
			case 1:
				zblok.color = ranRed()
			case 2:
				zblok.color = ranDarkGreen()
			case 3:
				zblok.color = ranDarkBlue()
			case 4:
				zblok.color = ranDarkGrey()
			case 5:
				zblok.color = ranDarkBlue()
			}
			zroom.floor = append(zroom.floor, zblok)

			x += siz
			if x >= gs.Level.LevX+gs.Level.LevW {
				x = gs.Level.LevX
				y += siz
			}
			if y >= gs.Level.LevY+gs.Level.LevW {
				break
			}
		}

		gs.Level.Level = append(gs.Level.Level, zroom)

		numRooms--
	}
	numRooms = orignumRooms

	//MAKE LEVEL MAP
	gs.Level.LevMap = nil
	mapRecSize := float32(96)
	rec := rl.NewRectangle(gs.Core.Cnt.X-mapRecSize/2, gs.Core.Cnt.Y-mapRecSize/2, mapRecSize, mapRecSize)
	gs.Level.LevMap = append(gs.Level.LevMap, rec)
	numRooms--
	countbreak := 0
	for numRooms > 0 {

		choose := gs.Level.LevMap[0]
		if len(gs.Level.LevMap) > 1 {
			choose = gs.Level.LevMap[rInt(0, len(gs.Level.LevMap))]
		}

		side := rInt(1, 5)
		checkV2 := rl.NewVector2(choose.X+choose.Width/2, choose.Y+choose.Width/2)
		switch side {
		case 1:
			checkV2.Y -= mapRecSize
		case 2:
			checkV2.X += mapRecSize
		case 3:
			checkV2.Y += mapRecSize
		case 4:
			checkV2.X -= mapRecSize
		}

		canadd := true
		for a := 0; a < len(gs.Level.LevMap); a++ {
			if rl.CheckCollisionPointRec(checkV2, gs.Level.LevMap[a]) {
				canadd = false
			}
		}

		if canadd {
			switch side {
			case 1:
				rec = choose
				rec.Y -= mapRecSize
			case 2:
				rec = choose
				rec.X += mapRecSize
			case 3:
				rec = choose
				rec.Y += mapRecSize
			case 4:
				rec = choose
				rec.X -= mapRecSize
			}

			gs.Level.LevMap = append(gs.Level.LevMap, rec)
			numRooms--
		} else {
			countbreak++
		}

		if countbreak == 1000 {
			break
		}
	}

	//FIND DOORS
	for a := 0; a < len(gs.Level.LevMap); a++ {

		checkV2 := rl.NewVector2(gs.Level.LevMap[a].X+gs.Level.LevMap[a].Width/2, gs.Level.LevMap[a].Y+gs.Level.LevMap[a].Width/2)
		ocheckV2 := checkV2
		checkV2.Y -= gs.Level.LevMap[a].Width
		for b := 0; b < len(gs.Level.LevMap); b++ {
			if a != b {
				if rl.CheckCollisionPointRec(checkV2, gs.Level.LevMap[b]) {
					gs.Level.Level[a].doorSides = append(gs.Level.Level[a].doorSides, 1)
					gs.Level.Level[a].nextRooms = append(gs.Level.Level[a].nextRooms, b)
				}
			}
		}
		checkV2 = ocheckV2
		checkV2.X += gs.Level.LevMap[a].Width
		for b := 0; b < len(gs.Level.LevMap); b++ {
			if a != b {
				if rl.CheckCollisionPointRec(checkV2, gs.Level.LevMap[b]) {
					gs.Level.Level[a].doorSides = append(gs.Level.Level[a].doorSides, 2)
					gs.Level.Level[a].nextRooms = append(gs.Level.Level[a].nextRooms, b)
				}
			}
		}
		checkV2 = ocheckV2
		checkV2.Y += gs.Level.LevMap[a].Width
		for b := 0; b < len(gs.Level.LevMap); b++ {
			if a != b {
				if rl.CheckCollisionPointRec(checkV2, gs.Level.LevMap[b]) {
					gs.Level.Level[a].doorSides = append(gs.Level.Level[a].doorSides, 3)
					gs.Level.Level[a].nextRooms = append(gs.Level.Level[a].nextRooms, b)
				}
			}
		}
		checkV2 = ocheckV2
		checkV2.X -= gs.Level.LevMap[a].Width
		for b := 0; b < len(gs.Level.LevMap); b++ {
			if a != b {
				if rl.CheckCollisionPointRec(checkV2, gs.Level.LevMap[b]) {
					gs.Level.Level[a].doorSides = append(gs.Level.Level[a].doorSides, 4)
					gs.Level.Level[a].nextRooms = append(gs.Level.Level[a].nextRooms, b)
				}
			}
		}

	}

	//REMOVE DOOR BLOKS
	for a := 0; a < len(gs.Level.Level); a++ {

		for b := 0; b < len(gs.Level.Level[a].doorSides); b++ {

			if gs.Level.Level[a].doorSides[b] == 1 {
				checkV2 := rl.NewVector2(gs.Level.LevX+gs.Level.LevW/2, gs.Level.LevY+gs.Level.BorderWallBlokSiz/2)
				checkV2L1 := checkV2
				checkV2L1.X -= gs.Level.BorderWallBlokSiz
				checkV2L2 := checkV2L1
				checkV2L2.X -= gs.Level.BorderWallBlokSiz
				checkV2R1 := checkV2
				checkV2R1.X += gs.Level.BorderWallBlokSiz
				checkV2R2 := checkV2R1
				checkV2R2.X += gs.Level.BorderWallBlokSiz

				exitRec := rl.NewRectangle(checkV2L2.X-gs.Level.BorderWallBlokSiz/2, checkV2L2.Y-gs.Level.BorderWallBlokSiz/2, gs.Level.BorderWallBlokSiz*5, bsU5)
				gs.Level.Level[a].doorExitRecs = append(gs.Level.Level[a].doorExitRecs, exitRec)

				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L1, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L2, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R1, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R2, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
			}
			if gs.Level.Level[a].doorSides[b] == 2 {
				checkV2 := rl.NewVector2(gs.Level.LevX+gs.Level.LevW-gs.Level.BorderWallBlokSiz/2, gs.Level.LevY+gs.Level.LevW/2)
				checkV2L1 := checkV2
				checkV2L1.Y -= gs.Level.BorderWallBlokSiz
				checkV2L2 := checkV2L1
				checkV2L2.Y -= gs.Level.BorderWallBlokSiz
				checkV2R1 := checkV2
				checkV2R1.Y += gs.Level.BorderWallBlokSiz
				checkV2R2 := checkV2R1
				checkV2R2.Y += gs.Level.BorderWallBlokSiz

				exitRec := rl.NewRectangle((checkV2L2.X+gs.Level.BorderWallBlokSiz/2)-bsU5, checkV2L2.Y-gs.Level.BorderWallBlokSiz/2, bsU5, gs.Level.BorderWallBlokSiz*5)
				gs.Level.Level[a].doorExitRecs = append(gs.Level.Level[a].doorExitRecs, exitRec)

				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L1, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L2, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R1, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R2, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
			}

			if gs.Level.Level[a].doorSides[b] == 3 {
				checkV2 := rl.NewVector2(gs.Level.LevX+gs.Level.LevW/2, gs.Level.LevY+gs.Level.LevW-gs.Level.BorderWallBlokSiz/2)
				checkV2L1 := checkV2
				checkV2L1.X -= gs.Level.BorderWallBlokSiz
				checkV2L2 := checkV2L1
				checkV2L2.X -= gs.Level.BorderWallBlokSiz
				checkV2R1 := checkV2
				checkV2R1.X += gs.Level.BorderWallBlokSiz
				checkV2R2 := checkV2R1
				checkV2R2.X += gs.Level.BorderWallBlokSiz

				exitRec := rl.NewRectangle(checkV2L2.X-gs.Level.BorderWallBlokSiz/2, (checkV2L2.Y+gs.Level.BorderWallBlokSiz/2)-bsU5, gs.Level.BorderWallBlokSiz*5, bsU5)
				gs.Level.Level[a].doorExitRecs = append(gs.Level.Level[a].doorExitRecs, exitRec)

				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L1, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L2, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R1, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R2, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
			}

			if gs.Level.Level[a].doorSides[b] == 4 {
				checkV2 := rl.NewVector2(gs.Level.LevX+gs.Level.BorderWallBlokSiz/2, gs.Level.LevY+gs.Level.LevW/2)
				checkV2L1 := checkV2
				checkV2L1.Y -= gs.Level.BorderWallBlokSiz
				checkV2L2 := checkV2L1
				checkV2L2.Y -= gs.Level.BorderWallBlokSiz
				checkV2R1 := checkV2
				checkV2R1.Y += gs.Level.BorderWallBlokSiz
				checkV2R2 := checkV2R1
				checkV2R2.Y += gs.Level.BorderWallBlokSiz

				exitRec := rl.NewRectangle((checkV2L2.X - gs.Level.BorderWallBlokSiz/2), checkV2L2.Y-gs.Level.BorderWallBlokSiz/2, bsU5, gs.Level.BorderWallBlokSiz*5)
				gs.Level.Level[a].doorExitRecs = append(gs.Level.Level[a].doorExitRecs, exitRec)

				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L1, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L2, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R1, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
				for c := 0; c < len(gs.Level.Level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R2, gs.Level.Level[a].walls[c].rec) {
						gs.Level.Level[a].walls = remBlok(gs.Level.Level[a].walls, c)
					}
				}
			}

		}

	}

	gs.Level.Level[0].visited = true

	makeInnerBloks()
	makemovebloks()
	makestatues()
	makeshop()
	makeetc()

	if gs.Level.Levelnum > 3 {
		makeblades()
		maketurrets()
	}
	if gs.Level.Levelnum > 2 {
		makespears()
	}
	if gs.Level.Levelnum > 1 {
		makespikes()
	}

	makeenemies()

	makeexit()

	cleanlevel()

}
func makenewlevel() { //MARK:MAKE NEW LEVEL

	gs.Level.NextLevelScreen = true
	gs.Level.Levelnum++

	for a := 0; a < len(gs.Level.Level); a++ {
		gs.Level.Level[a].etc = nil
		gs.Level.Level[a].enemies = nil
		gs.Level.Level[a].doorExitRecs = nil
		gs.Level.Level[a].doorSides = nil
		gs.Level.Level[a].floor = nil
		gs.Level.Level[a].innerBloks = nil
		gs.Level.Level[a].movBloks = nil
		gs.Level.Level[a].nextRooms = nil
		gs.Level.Level[a].spikes = nil
		gs.Level.Level[a].visited = false
		gs.Level.Level[a].walls = nil
	}

	gs.Level.Level = nil
	gs.FX.Fx = nil
	gs.Player.PlProj = nil
	gs.Enemies.EnProj = nil
	gs.FX.FloodRec.Y = gs.Core.ScrHF32 + bsU

	if gs.Level.Levelnum == 6 {
		gs.Player.Mods.planty = false
		gs.Player.Mods.carrot = false
		gs.Player.Mods.alien = false
		gs.Player.Mods.vine = false
		gs.Player.Mods.airstrike = false
		makeendlevel()
	} else {
		makelevel()
	}
	gs.Level.NextlevelT = gs.Core.Fps

	gs.Player.Pl.cnt = gs.Core.Cnt
	cntCompanion := gs.Player.Pl.cnt

	if gs.Player.Mods.carrot {
		gs.Companions.MrCarrot.rec = rl.NewRectangle(cntCompanion.X-gs.Companions.MrCarrot.rec.Width/2, cntCompanion.Y-gs.Companions.MrCarrot.rec.Width/2, gs.Companions.MrCarrot.rec.Width, gs.Companions.MrCarrot.rec.Width)
	}
	if gs.Player.Mods.alien {
		gs.Companions.MrAlien.rec = rl.NewRectangle(cntCompanion.X-gs.Companions.MrAlien.rec.Width/2, cntCompanion.Y-gs.Companions.MrAlien.rec.Width/2, gs.Companions.MrAlien.rec.Width, gs.Companions.MrAlien.rec.Width)
	}
	if gs.Player.Mods.planty {
		gs.Companions.MrPlanty.rec = rl.NewRectangle(cntCompanion.X-gs.Companions.MrPlanty.rec.Width/2, cntCompanion.Y-gs.Companions.MrPlanty.rec.Width/2, gs.Companions.MrPlanty.rec.Width, gs.Companions.MrPlanty.rec.Width)
	}

	gs.Level.RoomNum = 0

}
func makeInnerBloks() { //MARK: MAKE INNER BLOKS

	for a := 0; a < len(gs.Level.Level); a++ {

		num := rInt(3, 7)

		for num > 0 {
			countbreak := 100
			for {
				siz := rF32(bsU4, bsU8)
				zblok := xblok{}
				tl := findRanRecLoc(siz, siz, a)
				zblok.rec = rl.NewRectangle(tl.X, tl.Y, siz, siz)
				zblok.img = gs.Level.WallT
				switch gs.Level.Levelnum {
				case 1:
					zblok.color = ranBlue()
				case 2:
					zblok.color = ranBrown()
				case 3:
					zblok.color = ranOrange()
				case 4:
					zblok.color = ranDarkBlue()
				case 5:
					zblok.color = ranCol()
				case 6:
					zblok.color = ranOrange()
				}
				zblok.fade = 1
				zblok.crec = zblok.rec
				zblok.solid = true

				canadd := true
				for b := 0; b < len(gs.Level.Level[a].innerBloks); b++ {
					if rl.CheckCollisionRecs(zblok.rec, gs.Level.Level[a].innerBloks[b].rec) {
						canadd = false
					}
					if a == 0 {
						checkrec := rl.NewRectangle(gs.Core.Cnt.X-bsU4, gs.Core.Cnt.Y-bsU4, bsU8, bsU8)
						if rl.CheckCollisionRecs(zblok.rec, checkrec) {
							canadd = false
						}
					}
				}
				countbreak--

				if canadd || countbreak == 0 {
					gs.Level.Level[a].innerBloks = append(gs.Level.Level[a].innerBloks, zblok)
					break
				}
			}
			num--
		}
	}
}

func makemario() { //MARK:MAKE MARIO

	gs.Mario.MarioRecs = nil
	gs.Mario.MarioCols = nil
	gs.Mario.MarioCoins = nil
	gs.Mario.MarioCoinOnOff = nil
	gs.Mario.MarioT = gs.Core.Fps * 5

	//IMG
	gs.Mario.MarioImg = gs.Render.Knight[0]

	//BORDER REC
	gs.Mario.MarioScreenRec = rl.NewRectangle(gs.Core.Cnt.X-gs.Core.ScrWF32/(gs.Render.Cam2.Zoom*2), gs.Core.Cnt.Y-gs.Core.ScrHF32/(gs.Render.Cam2.Zoom*2), gs.Core.ScrWF32/gs.Render.Cam2.Zoom, gs.Core.ScrHF32/gs.Render.Cam2.Zoom)

	//BACK PATTERN
	gs.Mario.PatternRec = gs.Render.Patterns[rInt(0, len(gs.Render.Patterns))]

	//FLOOR
	siz := bsU4
	x := float32(0)
	y := gs.Level.LevRecInner.Y + gs.Level.LevRec.Width - (siz + bsU)
	for {
		gs.Mario.MarioRecs = append(gs.Mario.MarioRecs, rl.NewRectangle(x, y, siz, siz))
		gs.Mario.MarioCols = append(gs.Mario.MarioCols, ranBrown())
		if roll6() > 4 {
			gs.Mario.MarioCoins = append(gs.Mario.MarioCoins, rl.NewRectangle(x, y-siz, siz/2, siz/2))
			gs.Mario.MarioCoinOnOff = append(gs.Mario.MarioCoinOnOff, true)
		}
		x += siz
		if x >= gs.Core.ScrWF32 {
			break
		}
	}

	//PLAYER REC
	gs.Mario.MarioPL = rl.NewRectangle(gs.Core.ScrWF32/2, y-siz, siz, siz)
	gs.Mario.MarioV2L = rl.NewVector2(gs.Mario.MarioPL.X, gs.Mario.MarioPL.Y+gs.Mario.MarioPL.Width+2)
	gs.Mario.MarioV2R = rl.NewVector2(gs.Mario.MarioPL.X+gs.Mario.MarioPL.Width, gs.Mario.MarioPL.Y+gs.Mario.MarioPL.Width+2)

	//PLATFORMS
	y -= siz * 3
	num := float32(rInt(5, 9))
	x = gs.Mario.MarioScreenRec.X + rF32(0, (gs.Mario.MarioScreenRec.Width/2)-(num*siz))
	for num > 0 {
		gs.Mario.MarioRecs = append(gs.Mario.MarioRecs, rl.NewRectangle(x, y, siz, siz))
		gs.Mario.MarioCols = append(gs.Mario.MarioCols, ranGreen())
		if roll6() > 4 {
			gs.Mario.MarioCoins = append(gs.Mario.MarioCoins, rl.NewRectangle(x, y-siz, siz/2, siz/2))
			gs.Mario.MarioCoinOnOff = append(gs.Mario.MarioCoinOnOff, true)
		}
		x += siz
		num--
	}
	num = float32(rInt(5, 9))
	x = gs.Mario.MarioScreenRec.X + gs.Mario.MarioScreenRec.Width/2 + rF32(0, (gs.Mario.MarioScreenRec.Width/2)-(num*siz))
	for num > 0 {
		gs.Mario.MarioRecs = append(gs.Mario.MarioRecs, rl.NewRectangle(x, y, siz, siz))
		gs.Mario.MarioCols = append(gs.Mario.MarioCols, ranOrange())
		if roll6() > 4 {
			gs.Mario.MarioCoins = append(gs.Mario.MarioCoins, rl.NewRectangle(x, y-siz, siz/2, siz/2))
			gs.Mario.MarioCoinOnOff = append(gs.Mario.MarioCoinOnOff, true)
		}
		x += siz
		num--
	}
	y -= siz * 3
	num = float32(rInt(5, 9))
	x = gs.Mario.MarioScreenRec.X + rF32(0, (gs.Mario.MarioScreenRec.Width/2)-(num*siz))
	for num > 0 {
		gs.Mario.MarioRecs = append(gs.Mario.MarioRecs, rl.NewRectangle(x, y, siz, siz))
		gs.Mario.MarioCols = append(gs.Mario.MarioCols, ranCyan())
		if roll6() > 4 {
			gs.Mario.MarioCoins = append(gs.Mario.MarioCoins, rl.NewRectangle(x, y-siz, siz/2, siz/2))
			gs.Mario.MarioCoinOnOff = append(gs.Mario.MarioCoinOnOff, true)
		}
		x += siz
		num--
	}
	num = float32(rInt(5, 9))
	x = gs.Mario.MarioScreenRec.X + gs.Mario.MarioScreenRec.Width/2 + rF32(0, (gs.Mario.MarioScreenRec.Width/2)-(num*siz))
	for num > 0 {
		gs.Mario.MarioRecs = append(gs.Mario.MarioRecs, rl.NewRectangle(x, y, siz, siz))
		gs.Mario.MarioCols = append(gs.Mario.MarioCols, ranPink())
		if roll6() > 4 {
			gs.Mario.MarioCoins = append(gs.Mario.MarioCoins, rl.NewRectangle(x, y-siz, siz/2, siz/2))
			gs.Mario.MarioCoinOnOff = append(gs.Mario.MarioCoinOnOff, true)
		}
		x += siz
		num--
	}

}
func makeairstrike() { //MARK:MAKE AIR STRIKE

	gs.FX.AirstrikeV2 = nil
	gs.FX.AirstrikeDir = rInt(1, 5)
	//gs.FX.AirstrikeDir = 4

	switch gs.FX.AirstrikeDir {
	case 1:
		v2 := rl.NewVector2(rF32(gs.Level.LevRecInner.X+bsU4, gs.Level.LevRecInner.X+gs.Level.LevRecInner.Width-bsU8), gs.Level.LevRecInner.Y-bsU4)
		gs.FX.AirstrikeV2 = append(gs.FX.AirstrikeV2, v2)
		v2.Y -= bsU4
		v2.X -= bsU4
		gs.FX.AirstrikeV2 = append(gs.FX.AirstrikeV2, v2)
		v2.X += bsU8
		gs.FX.AirstrikeV2 = append(gs.FX.AirstrikeV2, v2)
	case 2:
		v2 := rl.NewVector2(gs.Level.LevRecInner.X+gs.Level.LevRecInner.Width+bsU4, rF32(gs.Level.LevRecInner.Y+bsU4, gs.Level.LevRecInner.Y+gs.Level.LevRecInner.Width-bsU8))
		gs.FX.AirstrikeV2 = append(gs.FX.AirstrikeV2, v2)
		v2.Y -= bsU4
		v2.X += bsU4
		gs.FX.AirstrikeV2 = append(gs.FX.AirstrikeV2, v2)
		v2.Y += bsU8
		gs.FX.AirstrikeV2 = append(gs.FX.AirstrikeV2, v2)
	case 3:
		v2 := rl.NewVector2(rF32(gs.Level.LevRecInner.X+bsU4, gs.Level.LevRecInner.X+gs.Level.LevRecInner.Width-bsU8), gs.Level.LevRecInner.Y+gs.Level.LevRecInner.Width+bsU4)
		gs.FX.AirstrikeV2 = append(gs.FX.AirstrikeV2, v2)
		v2.Y += bsU4
		v2.X -= bsU4
		gs.FX.AirstrikeV2 = append(gs.FX.AirstrikeV2, v2)
		v2.X += bsU8
		gs.FX.AirstrikeV2 = append(gs.FX.AirstrikeV2, v2)
	case 4:
		v2 := rl.NewVector2(gs.Level.LevRecInner.X-bsU4, rF32(gs.Level.LevRecInner.Y+bsU4, gs.Level.LevRecInner.Y+gs.Level.LevRecInner.Width-bsU8))
		gs.FX.AirstrikeV2 = append(gs.FX.AirstrikeV2, v2)
		v2.Y -= bsU4
		v2.X -= bsU4
		gs.FX.AirstrikeV2 = append(gs.FX.AirstrikeV2, v2)
		v2.Y += bsU8
		gs.FX.AirstrikeV2 = append(gs.FX.AirstrikeV2, v2)
	}

	gs.FX.AirstrikebombT = rI32(int(gs.Core.Fps/4), int(gs.Core.Fps*2))
	gs.FX.AirstrikeOn = true
}
func makefish() { //MARK:MAKE FISH

	gs.FX.FishV2 = rl.Vector2{}
	gs.FX.Fish2V2 = rl.Vector2{}

	gs.FX.FishSiz = rF32(bsU4, bsU7)
	gs.FX.FishSiz2 = rF32(bsU4, bsU7)

	gs.FX.FishV2.X = -gs.FX.FishSiz
	gs.FX.Fish2V2.X = gs.Core.ScrWF32 + gs.FX.FishSiz2

	gs.FX.FishV2.Y = rF32(gs.Core.ScrHF32/3, gs.Core.ScrHF32)
	gs.FX.Fish2V2.Y = rF32(gs.Core.ScrHF32/3, gs.Core.ScrHF32)

	gs.FX.FishRec = rl.NewRectangle(gs.FX.FishV2.X, gs.FX.FishV2.Y, gs.FX.FishSiz, gs.FX.FishSiz)
	gs.FX.FishRec2 = rl.NewRectangle(gs.FX.Fish2V2.X, gs.FX.Fish2V2.Y, gs.FX.FishSiz2, gs.FX.FishSiz2)

	gs.FX.Fish1 = gs.Render.FishR.recTL
	gs.FX.Fish2 = gs.Render.FishL.recTL

}
func makeplayer() { //MARK:MAKE PLAYER

	gs.Player.Pl.atkDMG = 1
	gs.Player.Pl.cnt = gs.Core.Cnt
	gs.Player.Pl.hp = 5
	gs.Player.Pl.hpmax = 5
	gs.Player.Pl.vel = 4
	gs.Player.Pl.siz = 72
	gs.Player.Pl.rec = rl.NewRectangle(gs.Player.Pl.cnt.X-gs.Player.Pl.siz/2, gs.Player.Pl.cnt.Y-gs.Player.Pl.siz/2, gs.Player.Pl.siz, gs.Player.Pl.siz)
	gs.Player.Pl.crec = gs.Player.Pl.rec
	gs.Player.Pl.arec = gs.Player.Pl.rec
	gs.Player.Pl.atkrec = gs.Player.Pl.rec
	gs.Player.Pl.atkrec.X -= bsU
	gs.Player.Pl.atkrec.Y -= bsU
	gs.Player.Pl.atkrec.Width += bsU2
	gs.Player.Pl.atkrec.Height += bsU2
	gs.Player.Pl.crec.X += gs.Player.Pl.crec.Width / 3
	gs.Player.Pl.crec.Y += gs.Player.Pl.crec.Height / 3
	gs.Player.Pl.crec.Width = gs.Player.Pl.crec.Width / 3
	gs.Player.Pl.crec.Height = gs.Player.Pl.crec.Height / 2
	gs.Player.Pl.framesAtk = 6
	gs.Player.Pl.framesWalk = 8
	gs.Player.Pl.sizImg = 32
	gs.Player.Pl.img = gs.Render.Knight[0]
	gs.Player.Pl.ori = rl.NewVector2(gs.Player.Pl.rec.Width/2, gs.Player.Pl.rec.Height/2)
	gs.Player.Pl.orbimg1 = gs.Render.Orbitalanim.recTL
	gs.Player.Pl.orbimg2 = gs.Player.Pl.orbimg1

	gs.Player.Max.axe = 10
	gs.Player.Max.fireball = 8
	gs.Player.Max.bounce = 4
	gs.Player.Max.key = 3
	gs.Player.Max.apple = 3
	gs.Player.Max.firetrail = 3
	gs.Player.Max.hppotion = 3
	gs.Player.Max.coffee = 5
	gs.Player.Max.atkrange = 3
	gs.Player.Max.atkdmg = 2
	gs.Player.Max.orbital = 2
	gs.Player.Max.hpring = 3
	gs.Player.Max.armor = 3
	gs.Player.Max.cherry = 99
	gs.Player.Max.cake = 99

}
func makeEnemyTypes() { //MARK: MAKE ENEMY TYPES

	siz := bsU4

	gs.Enemies.EnSpikes.img = rl.NewRectangle(0, 404, 44, 44)
	gs.Enemies.EnSpikes.fade = 1
	gs.Enemies.EnSpikes.col = rl.White
	gs.Enemies.EnSpikes.xImg = 0
	gs.Enemies.EnSpikes.frameNum = 15
	gs.Enemies.EnSpikes.vel = bsU / 5
	gs.Enemies.EnSpikes.name = "spikehog"
	gs.Enemies.EnSpikes.hp = 5
	gs.Enemies.EnSpikes.hpmax = gs.Enemies.EnSpikes.hp
	gs.Enemies.EnSpikes.rec = rl.NewRectangle(0, 0, siz, siz)

	gs.Enemies.EnGhost.imgl = rl.NewRectangle(4, 506, 44, 44)
	gs.Enemies.EnGhost.imgr = rl.NewRectangle(454, 506, 44, 44)
	gs.Enemies.EnGhost.img = gs.Enemies.EnGhost.imgr
	gs.Enemies.EnGhost.col = rl.White
	gs.Enemies.EnGhost.fade = 1
	gs.Enemies.EnGhost.xImg = 4
	gs.Enemies.EnGhost.xImg2 = 454
	gs.Enemies.EnGhost.frameNum = 9
	gs.Enemies.EnGhost.vel = bsU / 5
	gs.Enemies.EnGhost.name = "ghost"
	gs.Enemies.EnGhost.hp = 4
	gs.Enemies.EnGhost.hpmax = gs.Enemies.EnGhost.hp
	gs.Enemies.EnGhost.rec = rl.NewRectangle(0, 0, siz, siz)

	gs.Enemies.EnSlime.imgl = rl.NewRectangle(4, 556, 44, 44)
	gs.Enemies.EnSlime.imgr = rl.NewRectangle(458, 556, 44, 44)
	gs.Enemies.EnSlime.img = gs.Enemies.EnSlime.imgr
	gs.Enemies.EnSlime.col = rl.White
	gs.Enemies.EnSlime.fade = 1
	gs.Enemies.EnSlime.xImg = 4
	gs.Enemies.EnSlime.xImg2 = 458
	gs.Enemies.EnSlime.frameNum = 9
	gs.Enemies.EnSlime.vel = bsU / 5
	gs.Enemies.EnSlime.name = "slime"
	gs.Enemies.EnSlime.hp = 5
	gs.Enemies.EnSlime.hpmax = gs.Enemies.EnSlime.hp
	gs.Enemies.EnSlime.rec = rl.NewRectangle(0, 0, siz, siz)

	gs.Enemies.EnRock.imgl = rl.NewRectangle(4, 608, 22, 22)
	gs.Enemies.EnRock.imgr = rl.NewRectangle(318, 608, 22, 22)
	gs.Enemies.EnRock.img = gs.Enemies.EnRock.imgr
	gs.Enemies.EnRock.col = rl.White
	gs.Enemies.EnRock.fade = 1
	gs.Enemies.EnRock.xImg = 4
	gs.Enemies.EnRock.xImg2 = 318
	gs.Enemies.EnRock.frameNum = 13
	gs.Enemies.EnRock.vel = bsU / 4
	gs.Enemies.EnRock.name = "rock"
	gs.Enemies.EnRock.hp = 3
	gs.Enemies.EnRock.hpmax = gs.Enemies.EnRock.hp
	gs.Enemies.EnRock.rec = rl.NewRectangle(0, 0, bsU3, bsU3)

	gs.Enemies.EnMushroom.imgl = rl.NewRectangle(4, 634, 32, 32)
	gs.Enemies.EnMushroom.imgr = rl.NewRectangle(4, 668, 32, 32)
	gs.Enemies.EnMushroom.img = gs.Enemies.EnMushroom.imgr
	gs.Enemies.EnMushroom.col = rl.White
	gs.Enemies.EnMushroom.fade = 1
	gs.Enemies.EnMushroom.xImg = 4
	gs.Enemies.EnMushroom.xImg2 = 4
	gs.Enemies.EnMushroom.frameNum = 13
	gs.Enemies.EnMushroom.vel = bsU / 5
	gs.Enemies.EnMushroom.name = "mushroom"
	gs.Enemies.EnMushroom.hp = 5
	gs.Enemies.EnMushroom.hpmax = gs.Enemies.EnMushroom.hp
	gs.Enemies.EnMushroom.rec = rl.NewRectangle(0, 0, siz, siz)

}
func makeChainLightning() { //MARK:MAKE CHAIN LIGHTNING

	gs.FX.ChainV2 = nil
	if len(gs.Level.Level[gs.Level.RoomNum].enemies) > 1 {
		for a := 0; a < len(gs.Level.Level[gs.Level.RoomNum].enemies); a++ {
			if gs.Level.Level[gs.Level.RoomNum].enemies[a].cnt.X > gs.Level.LevRecInner.X && gs.Level.Level[gs.Level.RoomNum].enemies[a].cnt.Y > gs.Level.LevRecInner.Y {
				gs.FX.ChainV2 = append(gs.FX.ChainV2, gs.Level.Level[gs.Level.RoomNum].enemies[a].cnt)
			}
		}
		gs.FX.ChainLightTimer = gs.Core.Fps / 3
		gs.FX.ChainLightOn = true
		rl.PlaySound(gs.Audio.Sfx[11])
	}
}
func makeenemies() { //MARK:MAKE ENEMIES

	makeEnemyTypes()

	zen := xenemy{}
	siz := bsU2

	for a := 0; a < len(gs.Level.Level); a++ {

		if gs.Level.Levelnum > 4 {
			if flipcoin() { //MUSHROOM
				zen = xenemy{}
				zen = gs.Enemies.EnMushroom
				zen.T1 = rI32(int(gs.Core.Fps*2), int(gs.Core.Fps*5))
				canadd := true
				zen.rec, canadd = findRecPos(gs.Enemies.EnMushroom.rec.Width, a)
				zen.crec = zen.rec
				zen.crec.Y += zen.crec.Height / 2
				zen.crec.Height -= zen.crec.Height / 2
				zen.velX = rF32(-zen.vel, zen.vel)
				zen.velY = rF32(-zen.vel, zen.vel)
				if canadd {
					gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
				}
			}
			if gs.Level.Hardcore {
				if flipcoin() { //MUSHROOM
					zen = xenemy{}
					zen = gs.Enemies.EnMushroom
					zen.T1 = rI32(int(gs.Core.Fps*2), int(gs.Core.Fps*5))
					canadd := true
					zen.rec, canadd = findRecPos(gs.Enemies.EnMushroom.rec.Width, a)
					zen.crec = zen.rec
					zen.crec.Y += zen.crec.Height / 2
					zen.crec.Height -= zen.crec.Height / 2
					zen.velX = rF32(-zen.vel, zen.vel)
					zen.velY = rF32(-zen.vel, zen.vel)
					if canadd {
						gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
					}
				}
				if flipcoin() { //MUSHROOM
					zen = xenemy{}
					zen = gs.Enemies.EnMushroom
					zen.T1 = rI32(int(gs.Core.Fps*2), int(gs.Core.Fps*5))
					canadd := true
					zen.rec, canadd = findRecPos(gs.Enemies.EnMushroom.rec.Width, a)
					zen.crec = zen.rec
					zen.crec.Y += zen.crec.Height / 2
					zen.crec.Height -= zen.crec.Height / 2
					zen.velX = rF32(-zen.vel, zen.vel)
					zen.velY = rF32(-zen.vel, zen.vel)
					if canadd {
						gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
					}
				}
			}
		}
		if gs.Level.Levelnum > 3 {

			if flipcoin() { //GHOST
				zen = xenemy{}
				zen = gs.Enemies.EnGhost
				zen.fly = true
				canadd := true
				zen.rec, canadd = findRecPos(gs.Enemies.EnGhost.rec.Width, a)
				zen.crec = zen.rec
				zen.crec.Y += zen.crec.Height / 3
				zen.crec.Height -= zen.crec.Height / 3
				zen.crec.X += zen.crec.Width / 8
				zen.crec.Width -= zen.crec.Width / 4
				zen.arec = zen.crec
				zen.arec.X -= zen.arec.Width * 2
				zen.arec.Y -= zen.arec.Width * 2
				zen.arec.Width = zen.arec.Width * 5
				zen.arec.Height = zen.arec.Height * 5

				zen.velX = rF32(-zen.vel, zen.vel)
				zen.velY = rF32(-zen.vel, zen.vel)

				if canadd {
					gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
				}
			}
			if gs.Level.Hardcore {
				if flipcoin() { //GHOST
					zen = xenemy{}
					zen = gs.Enemies.EnGhost
					zen.fly = true
					canadd := true
					zen.rec, canadd = findRecPos(gs.Enemies.EnGhost.rec.Width, a)
					zen.crec = zen.rec
					zen.crec.Y += zen.crec.Height / 3
					zen.crec.Height -= zen.crec.Height / 3
					zen.crec.X += zen.crec.Width / 8
					zen.crec.Width -= zen.crec.Width / 4
					zen.arec = zen.crec
					zen.arec.X -= zen.arec.Width * 2
					zen.arec.Y -= zen.arec.Width * 2
					zen.arec.Width = zen.arec.Width * 5
					zen.arec.Height = zen.arec.Height * 5

					zen.velX = rF32(-zen.vel, zen.vel)
					zen.velY = rF32(-zen.vel, zen.vel)

					if canadd {
						gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
					}
				}
				if flipcoin() { //GHOST
					zen = xenemy{}
					zen = gs.Enemies.EnGhost
					zen.fly = true
					canadd := true
					zen.rec, canadd = findRecPos(gs.Enemies.EnGhost.rec.Width, a)
					zen.crec = zen.rec
					zen.crec.Y += zen.crec.Height / 3
					zen.crec.Height -= zen.crec.Height / 3
					zen.crec.X += zen.crec.Width / 8
					zen.crec.Width -= zen.crec.Width / 4
					zen.arec = zen.crec
					zen.arec.X -= zen.arec.Width * 2
					zen.arec.Y -= zen.arec.Width * 2
					zen.arec.Width = zen.arec.Width * 5
					zen.arec.Height = zen.arec.Height * 5

					zen.velX = rF32(-zen.vel, zen.vel)
					zen.velY = rF32(-zen.vel, zen.vel)

					if canadd {
						gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
					}
				}

			}

		}
		if gs.Level.Levelnum > 1 {

			if flipcoin() { //SLIME
				zen = xenemy{}
				zen = gs.Enemies.EnSlime
				zen.T1 = rI32(int(gs.Core.Fps/2), int(gs.Core.Fps*2))
				canadd := true
				zen.rec, canadd = findRecPos(gs.Enemies.EnSlime.rec.Width, a)
				zen.crec = zen.rec
				zen.crec.Y += zen.crec.Height / 3
				zen.crec.Height -= zen.crec.Height / 3
				zen.velX = rF32(-zen.vel, zen.vel)
				zen.velY = rF32(-zen.vel, zen.vel)
				if canadd {
					gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
				}
			}

			if gs.Level.Hardcore {
				if flipcoin() { //SLIME
					zen = xenemy{}
					zen = gs.Enemies.EnSlime
					zen.T1 = rI32(int(gs.Core.Fps/2), int(gs.Core.Fps*2))
					canadd := true
					zen.rec, canadd = findRecPos(gs.Enemies.EnSlime.rec.Width, a)
					zen.crec = zen.rec
					zen.crec.Y += zen.crec.Height / 3
					zen.crec.Height -= zen.crec.Height / 3
					zen.velX = rF32(-zen.vel, zen.vel)
					zen.velY = rF32(-zen.vel, zen.vel)
					if canadd {
						gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
					}
				}
				if flipcoin() { //SLIME
					zen = xenemy{}
					zen = gs.Enemies.EnSlime
					zen.T1 = rI32(int(gs.Core.Fps/2), int(gs.Core.Fps*2))
					canadd := true
					zen.rec, canadd = findRecPos(gs.Enemies.EnSlime.rec.Width, a)
					zen.crec = zen.rec
					zen.crec.Y += zen.crec.Height / 3
					zen.crec.Height -= zen.crec.Height / 3
					zen.velX = rF32(-zen.vel, zen.vel)
					zen.velY = rF32(-zen.vel, zen.vel)
					if canadd {
						gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
					}
				}
			}
		}

		if flipcoin() { //ROCK
			zen = xenemy{}
			zen = gs.Enemies.EnRock
			zen.spawnN = rInt(2, 11)
			zen.T1 = rI32(int(gs.Core.Fps*2), int(gs.Core.Fps*5))
			canadd := true
			zen.rec, canadd = findRecPos(gs.Enemies.EnRock.rec.Width, a)
			zen.crec = zen.rec
			zen.crec.Y += zen.crec.Height / 3
			zen.crec.Height -= zen.crec.Height / 3
			zen.velX = rF32(-zen.vel, zen.vel)
			zen.velY = rF32(-zen.vel, zen.vel)
			if canadd {
				gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
			}
		} else { //SPIKEHOG
			zen = xenemy{}
			zen = gs.Enemies.EnSpikes
			canadd := true
			zen.rec, canadd = findRecPos(gs.Enemies.EnSpikes.rec.Width, a)
			zen.crec = zen.rec
			zen.crec.Height = zen.crec.Height / 2
			zen.crec.Y += zen.crec.Height
			if flipcoin() {
				zen.velX = zen.vel
				if flipcoin() {
					zen.velX *= -1
				}
			} else {
				zen.velY = zen.vel
				if flipcoin() {
					zen.velY *= -1
				}
			}
			if canadd {
				gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
			}
		}
		if gs.Level.Hardcore {
			if flipcoin() { //ROCK
				zen = xenemy{}
				zen = gs.Enemies.EnRock
				zen.spawnN = rInt(2, 11)
				zen.T1 = rI32(int(gs.Core.Fps*2), int(gs.Core.Fps*5))
				canadd := true
				zen.rec, canadd = findRecPos(gs.Enemies.EnRock.rec.Width, a)
				zen.crec = zen.rec
				zen.crec.Y += zen.crec.Height / 3
				zen.crec.Height -= zen.crec.Height / 3
				zen.velX = rF32(-zen.vel, zen.vel)
				zen.velY = rF32(-zen.vel, zen.vel)
				if canadd {
					gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
				}
			} else { //SPIKEHOG
				zen = xenemy{}
				zen = gs.Enemies.EnSpikes
				canadd := true
				zen.rec, canadd = findRecPos(gs.Enemies.EnSpikes.rec.Width, a)
				zen.crec = zen.rec
				zen.crec.Height = zen.crec.Height / 2
				zen.crec.Y += zen.crec.Height
				if flipcoin() {
					zen.velX = zen.vel
					if flipcoin() {
						zen.velX *= -1
					}
				} else {
					zen.velY = zen.vel
					if flipcoin() {
						zen.velY *= -1
					}
				}
				if canadd {
					gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
				}
			}
		}

		//BATS
		if flipcoin() {
			zen.name = "bat"
			zen.hp = 2
			zen.hpmax = zen.hp
			zen.fly = true
			zen.vel = bsU / 3
			zen.velX = rF32(-zen.vel, zen.vel)
			zen.velY = rF32(-zen.vel, zen.vel)
			zen.img = gs.Render.Bats[rInt(0, len(gs.Render.Bats))]
			zen.fade = 1
			zen.col = rl.White
			zen.frameNum = 3
			zen.xImg = zen.img.X
			canadd := true
			zen.rec, canadd = findRecPos(siz, a)
			zen.ori = rl.NewVector2(0, 0)
			zen.crec = zen.rec
			if canadd {
				gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
			}

			if flipcoin() {
				zen.velX = rF32(-zen.vel, zen.vel)
				zen.velY = rF32(-zen.vel, zen.vel)
				zen.img = gs.Render.Bats[rInt(0, len(gs.Render.Bats))]
				zen.xImg = zen.img.X
				canadd = true
				zen.rec, canadd = findRecPos(siz, a)
				zen.crec = zen.rec
				if canadd {
					gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
				}
			}
		}

		//RABBITS
		if flipcoin() {
			zen = xenemy{}
			zen.hp = 3
			zen.hpmax = zen.hp
			zen.name = "rabbit1"
			zen.anim = true
			zen.vel = bsU / 4
			zen.velX = rF32(-zen.vel, zen.vel)
			zen.velY = rF32(-zen.vel, zen.vel)
			zen.img = gs.Render.Rabbit1.recTL
			zen.fade = 1
			zen.col = ranCol()
			canadd := true
			zen.rec, canadd = findRecPos(siz, a)
			zen.ori = rl.NewVector2(0, 0)
			zen.rec = rl.NewRectangle(zen.cnt.X-siz/2, zen.cnt.Y-siz/2, siz, siz)
			zen.crec = zen.rec
			if canadd {
				if zen.rec.X > gs.Level.LevRecInner.X && zen.rec.Y > gs.Level.LevRecInner.Y {
					gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
				}
			}

			if flipcoin() {
				zen.velX = rF32(-zen.vel, zen.vel)
				zen.velY = rF32(-zen.vel, zen.vel)
				zen.col = ranCol()
				canadd := true
				zen.rec, canadd = findRecPos(siz, a)
				zen.crec = zen.rec
				if canadd {
					if zen.rec.X > gs.Level.LevRecInner.X && zen.rec.Y > gs.Level.LevRecInner.Y {
						gs.Level.Level[a].enemies = append(gs.Level.Level[a].enemies, zen)
					}
				}
			}
		}

	}

}
func maketeleport() { //MARK:MAKE TELEPORT
	gs.Player.TeleportRadius = nil
	length := gs.Core.ScrWF32 / 2
	for a := 0; a < 10; a++ {
		gs.Player.TeleportRadius = append(gs.Player.TeleportRadius, length)
		length -= bsU4
	}

}
func makegascloud(cnt rl.Vector2) { //MARK:MAKE GAS CLOUD

	siz := bsU10

	zblok := xblok{}
	zblok.fade = 1
	zblok.cnt = cnt
	zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
	zblok.crec = zblok.rec
	zblok.crec.X += zblok.crec.Width / 4
	zblok.crec.Y += zblok.crec.Width / 4
	zblok.crec.Width = zblok.crec.Width / 2
	zblok.crec.Height = zblok.crec.Height / 2
	zblok.drec = zblok.rec
	zblok.drec.X += zblok.rec.Width / 2
	zblok.drec.Y += zblok.rec.Height / 2
	zblok.ori = rl.NewVector2(zblok.rec.Width/2, zblok.rec.Width/2)
	zblok.name = "gascloud"
	zblok.onoff = true
	zblok.color = ranGreen()
	zblok.img = gs.Render.Posiongas.recTL
	zblok.vel = bsU / 4
	zblok.velX = rF32(-zblok.vel, zblok.vel)
	zblok.velY = rF32(-zblok.vel, zblok.vel)
	gs.Level.Level[gs.Level.RoomNum].etc = append(gs.Level.Level[gs.Level.RoomNum].etc, zblok)

}
func makeSwitchArrows() { //MARK:MAKE SWITCH ARROWS

	choose := rInt(1, 5)

	switch choose {
	case 1:
		x := gs.Level.LevRecInner.X + bsU
		y := gs.Level.LevRecInner.Y
		for {
			makeProjectileEnemy(2, rl.NewVector2(x+bsU, y+bsU))
			x += bsU2
			if x > gs.Level.LevRecInner.X+gs.Level.LevRecInner.Width-bsU2 {
				break
			}
		}

	case 2:
		x := gs.Level.LevRecInner.X + gs.Level.LevRecInner.Width - bsU2
		y := gs.Level.LevRecInner.Y + bsU
		for {
			makeProjectileEnemy(5, rl.NewVector2(x+bsU, y+bsU))
			y += bsU2
			if y > gs.Level.LevRecInner.Y+gs.Level.LevRecInner.Width-bsU2 {
				break
			}
		}

	case 3:
		x := gs.Level.LevRecInner.X + bsU
		y := gs.Level.LevRecInner.Y + gs.Level.LevRecInner.Width - bsU2
		for {
			makeProjectileEnemy(4, rl.NewVector2(x+bsU, y+bsU))
			x += bsU2
			if x > gs.Level.LevRecInner.X+gs.Level.LevRecInner.Width-bsU2 {
				break
			}
		}

	case 4:
		x := gs.Level.LevRecInner.X
		y := gs.Level.LevRecInner.Y + bsU
		for {
			makeProjectileEnemy(3, rl.NewVector2(x+bsU, y+bsU))
			y += bsU2
			if y > gs.Level.LevRecInner.Y+gs.Level.LevRecInner.Width-bsU2 {
				break
			}
		}

	}

}

func makeexit() { //MARK:MAKE EXIT

	for {
		gs.Level.ExitRoomNum = rInt(1, len(gs.Level.Level))
		if gs.Level.ExitRoomNum != gs.Level.ShopRoomNum {
			gs.Level.Level[gs.Level.ExitRoomNum].exit = true
			break
		}
	}

	zblok := makeBlokGenNoRecNoCntr()
	zblok.img = gs.Render.Etc[53]
	zblok.name = "exit"
	zblok.color = ranBrown()
	siz := bsU3
	for {
		canadd := true
		zblok.rec, canadd = findRecPoswithSpacing(siz, bsU2, gs.Level.ExitRoomNum)
		if canadd {
			break
		}
	}

	gs.Level.Level[gs.Level.ExitRoomNum].etc = append(gs.Level.Level[gs.Level.ExitRoomNum].etc, zblok)

}
func makestatues() { //MARK:MAKE STATUES

	for a := 0; a < len(gs.Level.Level); a++ {

		if roll6() > 4 {
			zblok := makeBlokGenNoRecNoCntr()
			siz := rF32(bsU3, bsU5)
			found := true
			zblok.rec, found = findRecPoswithSpacing(siz, bsU2, a)
			zblok.img = gs.Render.Statues[rInt(0, len(gs.Render.Statues))]
			zblok.name = "statue"
			zblok.solid = true

			if found {
				gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
			}
		}
	}

}
func makespears() { //MARK:MAKE SPEARS

	for a := 0; a < len(gs.Level.Level); a++ {

		zblok := makeBlokGenNoRecNoCntr()
		multi := float32(2)
		choose := rInt(1, 5)
		switch choose {
		case 1:
			zblok.ro = 180
			zblok.rec = rl.NewRectangle(gs.Level.LevRecInner.X, gs.Level.LevRecInner.Y, multi*gs.Render.Spear.recTL.Width, multi*gs.Render.Spear.recTL.Height)
			for {
				zblok.rec.X += rF32(0, gs.Level.LevRecInner.Width-bsU2)
				if checkInnerBloksExits(a, zblok.rec) {
					break
				}
			}
			zblok.drec = zblok.rec
			zblok.drec.X += zblok.drec.Width / 2
			zblok.drec.Y += zblok.drec.Height / 2
			zblok.crec = zblok.rec
			zblok.ori = rl.NewVector2(zblok.drec.Width/2, zblok.drec.Height/2)
		case 2:
			zblok.ro = 270
			zblok.rec = rl.NewRectangle(gs.Level.LevRecInner.X+gs.Level.LevRecInner.Width-bsU-(multi*gs.Render.Spear.recTL.Height)/2, gs.Level.LevRecInner.Y+bsU-(multi*gs.Render.Spear.recTL.Height)/2, multi*gs.Render.Spear.recTL.Width, multi*gs.Render.Spear.recTL.Height)

			for {
				zblok.rec.Y += rF32(0, gs.Level.LevRecInner.Width-bsU2)
				if checkInnerBloksExits(a, zblok.rec) {
					break
				}
			}

			zblok.drec = zblok.rec
			zblok.drec.X += zblok.drec.Width / 2
			zblok.drec.Y += zblok.drec.Height / 2
			zblok.crec = rl.NewRectangle(gs.Level.LevRecInner.X+gs.Level.LevRecInner.Width-multi*gs.Render.Spear.recTL.Height, (zblok.rec.Y+zblok.rec.Height/2)-bsU, multi*gs.Render.Spear.recTL.Height, multi*gs.Render.Spear.recTL.Width)
			zblok.ori = rl.NewVector2(zblok.drec.Width/2, zblok.drec.Height/2)

		case 3:
			zblok.ro = 0
			zblok.rec = rl.NewRectangle(gs.Level.LevRecInner.X, gs.Level.LevRecInner.Y+gs.Level.LevRecInner.Height-(multi*gs.Render.Spear.recTL.Height), multi*gs.Render.Spear.recTL.Width, multi*gs.Render.Spear.recTL.Height)
			for {
				zblok.rec.X += rF32(0, gs.Level.LevRecInner.Width-bsU2)
				if checkInnerBloksExits(a, zblok.rec) {
					break
				}
			}
			zblok.drec = zblok.rec
			zblok.drec.X += zblok.drec.Width / 2
			zblok.drec.Y += zblok.drec.Height / 2
			zblok.crec = zblok.rec
			zblok.ori = rl.NewVector2(zblok.drec.Width/2, zblok.drec.Height/2)

		case 4:
			zblok.ro = 90
			zblok.rec = rl.NewRectangle(gs.Level.LevRecInner.X+bsU3, gs.Level.LevRecInner.Y+bsU-(multi*gs.Render.Spear.recTL.Height)/2, multi*gs.Render.Spear.recTL.Width, multi*gs.Render.Spear.recTL.Height)

			for {
				zblok.rec.Y += rF32(0, gs.Level.LevRecInner.Width-bsU2)
				if checkInnerBloksExits(a, zblok.rec) {
					break
				}
			}

			zblok.drec = zblok.rec
			zblok.drec.X += zblok.drec.Width / 2
			zblok.drec.Y += zblok.drec.Height / 2
			zblok.crec = rl.NewRectangle(zblok.rec.X-bsU3, (zblok.rec.Y+zblok.rec.Height/2)-bsU, multi*gs.Render.Spear.recTL.Height, multi*gs.Render.Spear.recTL.Width)
			zblok.ori = rl.NewVector2(zblok.drec.Width/2, zblok.drec.Height/2)
		}

		zblok.img = gs.Render.Spear.recTL
		zblok.name = "spear"
		if zblok.rec.X > gs.Level.LevRecInner.X && zblok.rec.X < gs.Level.LevRecInner.Width+gs.Level.LevRecInner.X && zblok.rec.Y > gs.Level.LevRecInner.Y && zblok.rec.Y < gs.Level.LevRecInner.Y+gs.Level.LevRecInner.Width {
			gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
		}

	}
}
func makeblades() { //MARK:MAKE BLADES

	for a := 0; a < len(gs.Level.Level); a++ {

		if flipcoin() {
			siz := rF32(bsU2, bsU4)
			zblok := makeBlokGenNoRecNoCntr()
			found := true
			zblok.rec, found = findRecPos(siz, a)
			zblok.img = gs.Render.Blades.recTL
			zblok.name = "blades"
			zblok.vel = bsU / 4
			zblok.velX = rF32(-zblok.vel, zblok.vel)
			zblok.velY = rF32(-zblok.vel, zblok.vel)

			if found {
				gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
			}
		}
	}
}
func makeshrines() { //MARK:MAKE SHRINES

	choose := rInt(0, len(gs.Level.Level))

	siz := rF32(bsU3, bsU5)
	zblok := xblok{}
	for {
		zblok = makeBlokGenRandom(siz)
		if checkInnerBloksExits(choose, zblok.rec) {
			break
		}
	}

	zblok.img = gs.Render.Shrines[rInt(0, len(gs.Render.Shrines))]
	zblok.name = "shrine"
	zblok.color = ranCol()
	zblok.fade = 0.2
	zblok.solid = true
	zblok.onoff = true
	gs.Level.Level[choose].etc = append(gs.Level.Level[choose].etc, zblok)

}
func makeetc() { //MARK:MAKE ETC

	for a := 0; a < len(gs.Level.Level); a++ {

		//GAS CLOUD TRAPS
		if roll6() == 6 {
			zblok := makeBlokGenNoRecNoCntr()
			canadd := true
			zblok.rec, canadd = findRecPoswithSpacing(bsU2, bsU/2, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.drec = zblok.rec
			zblok.drec.X += zblok.rec.Width / 2
			zblok.drec.Y += zblok.rec.Height / 2
			zblok.ori = rl.NewVector2(zblok.rec.Width/2, zblok.rec.Width/2)
			zblok.name = "gascloudtrap"
			zblok.color = rl.SkyBlue
			zblok.img = gs.Render.Etc[26]
			if canadd {
				gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
			}
		}

		//SWITCHES
		if roll6() == 6 {
			zblok := makeBlokGenNoRecNoCntr()
			canadd := true
			zblok.rec, canadd = findRecPoswithSpacing(bsU2, bsU/2, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.name = "switch"
			zblok.numType = roll6()
			if zblok.numType == 6 {
				zblok.numCoins = rInt(3, 11)
			}
			zblok.onoffswitch = flipcoin()
			zblok.color = rl.SkyBlue
			if zblok.onoffswitch {
				zblok.img = gs.Render.Etc[21]
			} else {
				zblok.img = gs.Render.Etc[22]
			}
			if canadd {
				gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
			}
		}
		//CHESTS
		if roll6() == 6 {

			siz := bsU2 + bsU/2
			zblok := makeBlokGenNoRecNoCntr()
			canadd := true
			zblok.rec, canadd = findRecPoswithSpacing(siz, bsU/2, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.name = "chest"
			zblok.fade = rF32(0.5, 0.8)
			zblok.color = ranOrange()
			zblok.crec = zblok.rec
			zblok.crec.Width += bsU
			zblok.crec.Height += bsU
			zblok.crec.X -= bsU / 2
			zblok.crec.Y -= bsU / 2
			zblok.img = gs.Render.Etc[23]
			if canadd {
				gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
			}
		}
		//SKULLS
		num := rInt(1, 5)
		for {
			zblok := makeBlokGenNoRecNoCntr()
			siz := rF32((bsU/4)*3, bsU+bsU/2)
			canadd := true
			zblok.rec, canadd = findRecPos(siz, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.name = "skull"
			zblok.fade = rF32(0.3, 0.6)
			zblok.color = ranGrey()
			zblok.img = gs.Render.Skulls[rInt(0, len(gs.Render.Skulls))]
			if canadd {
				gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
				num--
			}
			if num <= 0 {
				break
			}
		}

		//CANDLES
		num = rInt(0, 3)
		for {
			zblok := makeBlokGenNoRecNoCntr()
			siz := rF32(bsU+bsU/2, bsU2)
			canadd := true
			zblok.rec, canadd = findRecPos(siz, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.onoff = true
			zblok.name = "candle"
			zblok.fade = rF32(0.5, 0.8)
			zblok.color = rl.White
			zblok.img = gs.Render.Candles[rInt(0, len(gs.Render.Candles))]
			if canadd {
				gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
				num--
			}

			if num <= 0 {
				break
			}
		}

		//SIGNS
		num = rInt(0, 2)
		for {
			zblok := makeBlokGenNoRecNoCntr()
			siz := rF32(bsU+bsU/2, bsU2)
			canadd := true
			zblok.rec, canadd = findRecPos(siz, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.name = "sign"
			zblok.fade = rF32(0.4, 0.7)
			zblok.color = ranBrown()
			zblok.img = gs.Render.Signs[rInt(0, len(gs.Render.Signs))]
			if canadd {
				gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
				num--
			}

			if num <= 0 {
				break
			}
		}

		//PLANTS
		num = rInt(1, 5)
		for {
			zblok := makeBlokGenNoRecNoCntr()
			siz := rF32(bsU+bsU/2, bsU2)
			canadd := true
			zblok.rec, canadd = findRecPos(siz, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.name = "plant"
			zblok.fade = rF32(0.3, 0.6)
			zblok.color = ranGreen()
			zblok.img = gs.Render.Plants[rInt(0, len(gs.Render.Plants))]
			if canadd {
				gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
				num--
			}

			if num <= 0 {
				break
			}
		}

		//MUSHROOMS
		num = rInt(1, 5)
		for {
			zblok := makeBlokGenNoRecNoCntr()
			siz := rF32(bsU, bsU+bsU/2)
			canadd := true
			zblok.rec, canadd = findRecPos(siz, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.name = "mushroom"
			zblok.fade = rF32(0.3, 0.6)
			zblok.color = rl.White
			zblok.img = gs.Render.Mushrooms[rInt(0, len(gs.Render.Mushrooms))]
			if canadd {
				gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
				num--
			}

			if num <= 0 {
				break
			}
		}
		//POWERUP BLOK
		if roll6() > 2 {
			zblok := makeBlokGenNoRecNoCntr()
			canadd := true
			zblok.rec, canadd = findRecPoswithSpacing(bsU+bsU/2, bsU/4, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.solid = true
			zblok.name = "powerupBlok"
			zblok.color = brightYellow()
			zblok.img = gs.Render.Etc[27]
			if canadd {
				gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)

			}
		}

		//OIL BARRELS
		num = rInt(0, 3)
		for {
			zblok := makeBlokGenNoRecNoCntr()
			canadd := true
			zblok.rec, canadd = findRecPoswithSpacing(bsU2, bsU/2, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.solid = true
			zblok.onoff = true
			zblok.name = "oilbarrel"
			zblok.color = rl.DarkGreen
			zblok.img = gs.Render.Etc[20]
			if canadd {
				gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
				num--
			}

			if num <= 0 {
				break
			}
		}

		//SPRINGS
		if flipcoin() {
			num := rInt(1, 4)
			siz := bsU2
			zblok := xblok{}
			zblok.onoff = true
			zblok.name = "spring"
			zblok.fade = 1
			zblok.color = rl.White
			zblok.rec = rl.NewRectangle(0, 0, siz, siz)
			zblok.img = gs.Render.Spring.recTL

			addtolevel := false

			for num > 0 {
				choose := rInt(1, 5)
				switch choose {
				case 1:
					zblok.rec.Y = gs.Level.LevRecInner.Y
					for {
						zblok.rec.X = gs.Level.LevRecInner.X + rF32(0, gs.Level.LevRecInner.Width-bsU2)
						canadd := true
						for b := 0; b < len(gs.Level.Level[a].doorExitRecs); b++ {
							if rl.CheckCollisionRecs(zblok.rec, gs.Level.Level[a].doorExitRecs[b]) {
								canadd = false
							}
						}
						if canadd {
							addtolevel = true
							break
						}
					}
					zblok.ro = 180
					zblok.slideDIR = 3
				case 2:
					zblok.rec.X = gs.Level.LevRecInner.X + gs.Level.LevRecInner.Width - siz
					for {
						zblok.rec.Y = gs.Level.LevRecInner.Y + rF32(0, gs.Level.LevRecInner.Width-bsU2)
						canadd := true
						for b := 0; b < len(gs.Level.Level[a].doorExitRecs); b++ {
							if rl.CheckCollisionRecs(zblok.rec, gs.Level.Level[a].doorExitRecs[b]) {
								canadd = false
							}
						}
						if canadd {
							addtolevel = true
							break
						}
					}
					zblok.ro = 270
					zblok.slideDIR = 4
				case 3:
					zblok.rec.Y = gs.Level.LevRecInner.Y + gs.Level.LevRecInner.Width - siz
					for {
						zblok.rec.X = gs.Level.LevRecInner.X + rF32(0, gs.Level.LevRecInner.Width-bsU2)
						canadd := true
						for b := 0; b < len(gs.Level.Level[a].doorExitRecs); b++ {
							if rl.CheckCollisionRecs(zblok.rec, gs.Level.Level[a].doorExitRecs[b]) {
								canadd = false
							}
						}
						if canadd {
							addtolevel = true
							break
						}
					}
					zblok.ro = 0
					zblok.slideDIR = 1
				case 4:
					zblok.rec.X = gs.Level.LevRecInner.X
					for {
						zblok.rec.Y = gs.Level.LevRecInner.Y + rF32(0, gs.Level.LevRecInner.Width-bsU2)
						canadd := true
						for b := 0; b < len(gs.Level.Level[a].doorExitRecs); b++ {
							if rl.CheckCollisionRecs(zblok.rec, gs.Level.Level[a].doorExitRecs[b]) {
								canadd = false
							}
						}
						if canadd {
							addtolevel = true
							break
						}
					}
					zblok.ro = 90
					zblok.slideDIR = 2
				}

				if addtolevel {
					zblok.drec = zblok.rec
					zblok.drec.X += zblok.rec.Width / 2
					zblok.drec.Y += zblok.rec.Width / 2
					zblok.crec = zblok.rec

					gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)

					num--
				}
			}
		}

		//SPRING BLOKS
		if flipcoin() {
			siz := bsU2
			zblok := xblok{}
			zblok = makeBlokGenNoRecNoCntr()

			canadd := true
			zblok.rec, canadd = findRecPoswithSpacing(siz, bsU2, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.solid = true
			zblok.name = "springblok"
			zblok.fade = 1
			switch gs.Level.Levelnum {
			case 1:
				zblok.color = ranBlue()
			case 2:
				zblok.color = ranBrown()
			case 3:
				zblok.color = ranOrange()
			case 4:
				zblok.color = ranDarkBlue()
			case 5:
				zblok.color = ranCol()
			}
			zblok.img = gs.Level.WallT
			if canadd {
				gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)

				zblok.name = "spring"
				zblok.solid = false
				zblok.color = rl.White
				zblok.img = gs.Render.Spring.recTL

				if flipcoin() {
					zblok.rec.X += siz
					zblok.slideDIR = 2
					zblok.ro = 90
					zblok.drec = zblok.rec
					zblok.drec.X += zblok.rec.Width / 2
					zblok.drec.Y += zblok.rec.Width / 2
					zblok.crec = zblok.rec
					gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
				} else {
					zblok.rec.X += siz
				}

				if flipcoin() {
					zblok.rec.X -= siz * 2
					zblok.slideDIR = 4
					zblok.ro = 270
					zblok.drec = zblok.rec
					zblok.drec.X += zblok.rec.Width / 2
					zblok.drec.Y += zblok.rec.Width / 2
					zblok.crec = zblok.rec
					gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
				} else {
					zblok.rec.X -= siz * 2
				}

				if flipcoin() {
					zblok.rec.Y += siz
					zblok.rec.X += siz
					zblok.slideDIR = 3
					zblok.ro = 180
					zblok.drec = zblok.rec
					zblok.drec.X += zblok.rec.Width / 2
					zblok.drec.Y += zblok.rec.Width / 2
					zblok.crec = zblok.rec
					gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
				} else {
					zblok.rec.Y += siz
					zblok.rec.X += siz
				}

				if flipcoin() {
					zblok.rec.Y -= siz * 2
					zblok.slideDIR = 1
					zblok.ro = 0
					zblok.drec = zblok.rec
					zblok.drec.X += zblok.rec.Width / 2
					zblok.drec.Y += zblok.rec.Width / 2
					zblok.crec = zblok.rec
					gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
				} else {
					zblok.rec.Y -= siz * 2
				}
			}
		}
	}
}
func BlurRec(rec rl.Rectangle, dist float32) rl.Rectangle { //MARK:MAKE BLUR REC
	rec.X += rF32(-dist, dist)
	rec.Y += rF32(-dist, dist)
	return rec
}

func makeDrec(rec rl.Rectangle) rl.Rectangle { //MARK:MAKE DREC
	rec.X += rec.Width / 2
	rec.Y += rec.Height / 2
	return rec

}
func origin(rec rl.Rectangle) rl.Vector2 { //MARK:MAKE ORIGIN
	return rl.NewVector2(rec.Width/2, rec.Height/2)
}

func makesnow() { //MARK:MAKE SNOW

	num := rInt(50, 100)

	for {

		x := rF32(0, gs.Core.ScrWF32)
		y := gs.Level.LevY - rF32(bsU, gs.Core.ScrHF32)
		siz := rF32(bsU, bsU3)

		zimg := ximg{}
		zimg.rec = rl.NewRectangle(x, y, siz, siz)
		zimg.cnt = rl.NewVector2(zimg.rec.X+zimg.rec.Width/2, zimg.rec.Y+zimg.rec.Height/2)
		choose := rInt(1, 4)
		switch choose {
		case 1:
			zimg.img = gs.Render.Etc[15]
		case 2:
			zimg.img = gs.Render.Etc[16]
		case 3:
			zimg.img = gs.Render.Etc[17]
		}
		zimg.ori = rl.NewVector2(zimg.rec.Width/2, zimg.rec.Height/2)
		zimg.ro = rF32(0, 360)
		zimg.col = rl.White
		zimg.fade = rF32(0.3, 0.7)
		gs.FX.Snow = append(gs.FX.Snow, zimg)

		num--
		if num == 0 {
			break
		}
	}

}

func makeProjectile(name string) { //MARK:MAKE PROJECTILE

	zproj := xproj{}
	zproj.name = name
	zproj.cnt = gs.Player.Pl.cnt
	zproj.onoff = true
	zproj.vel = bsU / 2
	zproj.fade = 1
	zproj.dmg = 1
	zproj.bounceN = gs.Player.Mods.bounceN

	siz := bsU + bsU/2

	switch name {
	case "gs.Companions.MrCarrot":
		siz = bsU2
		zproj.cnt = rl.NewVector2(gs.Companions.MrCarrot.rec.X+gs.Companions.MrCarrot.rec.Width/2, gs.Companions.MrCarrot.rec.Y+gs.Companions.MrCarrot.rec.Height/2)
		zproj.rec = rl.NewRectangle(gs.Companions.MrCarrot.cnt.X-siz/2, gs.Companions.MrCarrot.cnt.Y-siz/2, siz, siz)
		zproj.drec = zproj.rec
		zproj.drec.X += zproj.rec.Width / 2
		zproj.drec.Y += zproj.rec.Height / 2
		zproj.ori = rl.NewVector2(zproj.rec.Width/2, zproj.rec.Height/2)
		zproj.img = gs.Render.Etc[50]
		zproj.col = rl.White
		zproj.vel = bsU / 4

		choose := rInt(1, 9)

		switch choose {
		case 1:
			zproj.velx = -zproj.vel
			zproj.vely = -zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		case 2:
			zproj.vely = -zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		case 3:
			zproj.velx = +zproj.vel
			zproj.vely = -zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		case 4:
			zproj.velx = zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		case 5:
			zproj.velx = +zproj.vel
			zproj.vely = +zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		case 6:
			zproj.vely = +zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		case 7:
			zproj.velx = -zproj.vel
			zproj.vely = +zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		case 8:
			zproj.velx = -zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		}

		choose2 := rInt(1, 9)
		for {
			choose = rInt(1, 9)
			if choose != choose2 {
				break
			}
		}

		switch choose2 {
		case 1:
			zproj.velx = -zproj.vel
			zproj.vely = -zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		case 2:
			zproj.vely = -zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		case 3:
			zproj.velx = +zproj.vel
			zproj.vely = -zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		case 4:
			zproj.velx = zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		case 5:
			zproj.velx = +zproj.vel
			zproj.vely = +zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		case 6:
			zproj.vely = +zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		case 7:
			zproj.velx = -zproj.vel
			zproj.vely = +zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		case 8:
			zproj.velx = -zproj.vel
			gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		}

	case "fireworks":
		siz = bsU2
		zproj.cnt = gs.FX.FireworksCnt
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.drec = zproj.rec
		zproj.drec.X += zproj.rec.Width / 2
		zproj.drec.Y += zproj.rec.Height / 2
		zproj.ori = rl.NewVector2(zproj.rec.Width/2, zproj.rec.Height/2)
		zproj.img = gs.Render.Etc[49]
		zproj.col = rl.White

		zproj.velx = bsU / 2
		zproj.vely = -bsU / 2
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		zproj.velx = bsU / 2
		zproj.vely = +bsU / 2
		zproj.ro = 90
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		zproj.velx = -bsU / 2
		zproj.vely = -bsU / 2
		zproj.ro = 270
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		zproj.velx = -bsU / 2
		zproj.vely = +bsU / 2
		zproj.ro = 180
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)

	case "gs.Companions.MrAlien":
		siz = bsU
		zproj.cnt = rl.NewVector2(gs.Companions.MrAlien.rec.X+gs.Companions.MrAlien.rec.Width/2, gs.Companions.MrAlien.rec.Y+gs.Companions.MrAlien.rec.Height/2)
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.drec = zproj.rec
		zproj.drec.X += zproj.rec.Width / 2
		zproj.drec.Y += zproj.rec.Height / 2
		zproj.ori = rl.NewVector2(zproj.rec.Width/2, zproj.rec.Height/2)
		zproj.img = gs.Render.PlantBull.recTL
		zproj.col = ranGreen()

		zproj.velx = bsU / 4
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		zproj.velx = -bsU / 4
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		zproj.velx = 0

		zproj.vely = bsU / 4
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		zproj.vely = -bsU / 4
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		zproj.vely = 0

		zproj.velx = bsU / 4
		zproj.vely = -bsU / 4
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		zproj.velx = bsU / 4
		zproj.vely = +bsU / 4
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		zproj.velx = -bsU / 4
		zproj.vely = -bsU / 4
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)
		zproj.velx = -bsU / 4
		zproj.vely = +bsU / 4
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)

	case "plantbull":
		siz = bsU
		zproj.cnt = rl.NewVector2(gs.Companions.MrPlanty.rec.X+gs.Companions.MrPlanty.rec.Width/2, gs.Companions.MrPlanty.rec.Y+gs.Companions.MrPlanty.rec.Height/2)
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.drec = zproj.rec
		zproj.drec.X += zproj.rec.Width / 2
		zproj.drec.Y += zproj.rec.Height / 2
		zproj.ori = rl.NewVector2(zproj.rec.Width/2, zproj.rec.Height/2)
		zproj.drec = zproj.rec
		zproj.img = gs.Render.PlantBull.recTL
		zproj.col = ranGreen()
		if gs.Companions.MrPlanty.velx > 0 {
			zproj.velx = bsU / 4
		} else {
			zproj.velx = -bsU / 4
		}
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)

	case "fireball":
		zproj.img = gs.Render.FireballPlayer.recTL
		zproj.col = ranOrange()

		switch gs.Player.Pl.direc {
		case 1:
			zproj.vely = -zproj.vel
			zproj.ro = 270
		case 2:
			zproj.velx = zproj.vel
		case 3:
			zproj.vely = zproj.vel
			zproj.ro = 90
		case 4:
			zproj.velx = -zproj.vel
			zproj.ro = 180
		default:
			zproj.velx = zproj.vel
		}

		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.drec = zproj.rec
		zproj.drec.X += zproj.rec.Width / 2
		zproj.drec.Y += zproj.rec.Height / 2
		zproj.ori = rl.NewVector2(zproj.rec.Width/2, zproj.rec.Height/2)
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)

		if gs.Player.Mods.fireballN > 1 {
			zproj.col = ranOrange()
			if gs.Player.Mods.fireballN >= 2 {
				zproj.ro += 180
				zproj.velx *= -1
				zproj.vely *= -1
				gs.Player.PlProj = append(gs.Player.PlProj, zproj)
			}
			if gs.Player.Mods.fireballN >= 3 {
				zproj.col = ranOrange()
				zproj.ro += 90
				zproj.vely = 0
				zproj.velx = 0
				switch gs.Player.Pl.direc {
				case 1:
					zproj.velx = zproj.vel
				case 2:
					zproj.vely = zproj.vel
				case 3:
					zproj.velx = -zproj.vel
				case 4:
					zproj.vely = -zproj.vel
				}
				gs.Player.PlProj = append(gs.Player.PlProj, zproj)
			}
			if gs.Player.Mods.fireballN >= 4 {
				zproj.col = ranOrange()
				zproj.ro += 180
				zproj.velx *= -1
				zproj.vely *= -1
				gs.Player.PlProj = append(gs.Player.PlProj, zproj)
			}

			if gs.Player.Mods.fireballN >= 5 {
				zproj.col = ranOrange()
				zproj.ro = 0
				switch gs.Player.Pl.direc {
				case 1:
					zproj.velx = zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 315
				case 2:
					zproj.velx = zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 45
				case 3:
					zproj.velx = -zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 135
				case 4:
					zproj.velx = -zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 225
				}
				gs.Player.PlProj = append(gs.Player.PlProj, zproj)
			}

			if gs.Player.Mods.fireballN >= 6 {
				zproj.col = ranOrange()
				zproj.ro = 0
				switch gs.Player.Pl.direc {
				case 1:
					zproj.velx = -zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 135
				case 2:
					zproj.velx = -zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 225
				case 3:
					zproj.velx = zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 315
				case 4:
					zproj.velx = zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 45
				}
				gs.Player.PlProj = append(gs.Player.PlProj, zproj)
			}

			if gs.Player.Mods.fireballN >= 7 {
				zproj.col = ranOrange()
				zproj.ro = 0
				switch gs.Player.Pl.direc {
				case 1:
					zproj.velx = zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 45
				case 2:
					zproj.velx = -zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 135
				case 3:
					zproj.velx = -zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 225
				case 4:
					zproj.velx = zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 315
				}
				gs.Player.PlProj = append(gs.Player.PlProj, zproj)
			}

			if gs.Player.Mods.fireballN >= 8 {
				zproj.col = ranOrange()
				zproj.ro = 0
				switch gs.Player.Pl.direc {
				case 1:
					zproj.velx = -zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 225
				case 2:
					zproj.velx = zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 315
				case 3:
					zproj.velx = zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 45
				case 4:
					zproj.velx = -zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 135
				}
				gs.Player.PlProj = append(gs.Player.PlProj, zproj)
			}

		}

	case "axe":
		zproj.img = gs.Render.Etc[3]
		zproj.col = rl.SkyBlue

		for {
			zproj.velx = rF32(-zproj.vel, zproj.vel)
			zproj.vely = rF32(-zproj.vel, zproj.vel)
			if getabs(zproj.vely) > zproj.vel/2 || getabs(zproj.velx) > zproj.vel/2 {
				break
			}
		}

		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.drec = zproj.rec
		zproj.drec.X += zproj.rec.Width / 2
		zproj.drec.Y += zproj.rec.Height / 2
		zproj.ori = rl.NewVector2(zproj.rec.Width/2, zproj.rec.Height/2)
		gs.Player.PlProj = append(gs.Player.PlProj, zproj)
	}

}

func makespikes() { //MARK: MAKE SPIKES

	siz := bsU2
	zblok := xblok{}
	zblok.name = "spikes"
	zblok.fade = 1
	zblok.img = gs.Render.Spikes.recTL

	for a := 0; a < len(gs.Level.Level); a++ {
		if flipcoin() {
			zblok.cnt = findRanCntV2()
			num := rInt(10, 30)
			for num > 0 {
				zblok.color = ranBlue()
				zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
				v1 := rl.NewVector2(zblok.rec.X, zblok.rec.Y)
				v2 := v1
				v2.X += siz
				v3 := v2
				v3.Y += siz
				v4 := v1
				v4.Y += siz

				if rl.CheckCollisionPointRec(v1, gs.Level.LevRecInner) && rl.CheckCollisionPointRec(v2, gs.Level.LevRecInner) && rl.CheckCollisionPointRec(v3, gs.Level.LevRecInner) && rl.CheckCollisionPointRec(v4, gs.Level.LevRecInner) {
					gs.Level.Level[a].spikes = append(gs.Level.Level[a].spikes, zblok)
				}

				choose := rInt(1, 5)
				switch choose {
				case 1:
					zblok.cnt.Y -= siz
				case 2:
					zblok.cnt.X += siz
				case 3:
					zblok.cnt.Y += siz
				case 43:
					zblok.cnt.X -= siz
				}

				num--
			}
		}
	}

}
func maketurrets() { //MARK: MAKE TURRETS

	siz := bsU2
	zblok := xblok{}
	zblok.name = "turret"
	zblok.img = gs.Render.Etc[18]
	zblok.solid = true
	zblok.fade = 1
	zblok.onoff = true

	for a := 0; a < len(gs.Level.Level); a++ {

		if flipcoin() {
			num := rInt(1, 4)
			for num > 0 {
				for {
					zblok.cnt = findRanCntV2()
					zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
					if checkInnerBloksExits(a, zblok.rec) {
						break
					}
				}
				zblok.drec = zblok.rec
				zblok.drec.X += zblok.rec.Width / 2
				zblok.drec.Y += zblok.rec.Height / 2
				zblok.ori = rl.NewVector2(zblok.drec.Width, zblok.drec.Height)
				zblok.color = rl.White
				zblok.crec = zblok.rec
				zblok.ro = rF32(0, 360)
				zblok.timer = rI32(1, 3) * gs.Core.Fps
				gs.Level.Level[a].etc = append(gs.Level.Level[a].etc, zblok)
				num--
			}
		}
	}

}

func makeFX(numType int, cnt rl.Vector2) { //MARK: MAKE FX

	zfx := xfx{}
	zfx.onoff = true
	zfx.timer = gs.Core.Fps * 2
	zfx.cnt = cnt

	switch numType {
	case 4:
		zfx.name = "fxBurnOilBarrel"
		zfx.timer = gs.Core.Fps * 4
		siz := bsU2
		cnt2 := cnt
		cnt2.X -= siz
		origX := cnt2.X
		cnt2.Y -= siz
		count := 0
		for a := 0; a < 9; a++ {
			if a != 4 {
				if flipcoin() {
					if rl.CheckCollisionPointRec(cnt2, gs.Level.LevRecInner) {
						zfx.rec = rl.NewRectangle(cnt2.X-siz/2, cnt2.Y-siz/2, siz, siz)
						zfx.img = gs.Render.Burn.recTL
						zfx.col = ranOrange()
						zfx.fade = rF32(0.7, 1.1)
						gs.FX.Fx = append(gs.FX.Fx, zfx)
					}
				}
			}
			cnt2.X += siz
			count++
			if count == 3 {
				count = 0
				cnt2.X = origX
				cnt2.Y += siz
			}
		}

		cnt2 = cnt
		cnt2.X -= siz * 2
		origX = cnt2.X
		cnt2.Y -= siz * 2
		count = 0
		for a := 0; a < 25; a++ {
			if a != 12 {
				if flipcoin() {
					if rl.CheckCollisionPointRec(cnt2, gs.Level.LevRecInner) {
						zfx.rec = rl.NewRectangle(cnt2.X-siz/2, cnt2.Y-siz/2, siz, siz)
						zfx.img = gs.Render.Burn.recTL
						zfx.col = ranOrange()
						zfx.fade = rF32(0.7, 1.1)
						gs.FX.Fx = append(gs.FX.Fx, zfx)
					}
				}
			}
			cnt2.X += siz
			count++
			if count == 5 {
				count = 0
				cnt2.X = origX
				cnt2.Y += siz
			}
		}

	case 3: //BURN WOOD BARREL
		zfx.name = "fxBurnWoodBarrel"
		zfx.timer = gs.Core.Fps * 4
		siz := bsU2
		cnt2 := cnt

		num := rInt(1, 4)

		for num > 0 {
			countbreak := 20
			for {
				choose := rInt(1, 9)
				switch choose {
				case 1:
					cnt2.Y -= bsU2
					cnt2.X -= bsU2
				case 2:
					cnt2.Y -= bsU2
				case 3:
					cnt2.Y -= bsU2
					cnt2.X += bsU2
				case 4:
					cnt2.X += bsU2
				case 5:
					cnt2.Y += bsU2
					cnt2.X += bsU2
				case 6:
					cnt2.Y += bsU2
				case 7:
					cnt2.Y += bsU2
					cnt2.X -= bsU2
				case 8:
					cnt2.X -= bsU2
				}
				if rl.CheckCollisionPointRec(cnt2, gs.Level.LevRecInner) || countbreak == 0 {
					break
				}
				countbreak--
			}

			zfx.rec = rl.NewRectangle(cnt2.X-siz/2, cnt2.Y-siz/2, siz, siz)
			zfx.img = gs.Render.Burn.recTL
			zfx.col = ranOrange()
			zfx.fade = rF32(0.7, 1.1)

			gs.FX.Fx = append(gs.FX.Fx, zfx)
			num--
		}

	case 2: //ENEMY
		zfx.name = "fxEnemy"
		num := 20
		for {

			zrec := xrec{}
			wid := rF32(bsU, bsU3)
			zrec.rec = rl.NewRectangle(cnt.X-wid/2, cnt.Y-wid/2, wid, wid)
			zrec.velX = rF32(-bsU, bsU)
			zrec.velY = rF32(-bsU, bsU)
			zrec.col = ranRed()
			zrec.fade = rF32(0.7, 1.1)

			zfx.recs = append(zfx.recs, zrec)

			num--
			if num == 0 {
				break
			}
		}

		gs.FX.Fx = append(gs.FX.Fx, zfx)
	case 1: //BARREL
		zfx.name = "fxBarrel"
		num := 20
		for {

			zrec := xrec{}
			zrec.rec = rl.NewRectangle(cnt.X, cnt.Y, bsU/2, bsU/2)
			zrec.velX = rF32(-bsU, bsU)
			zrec.velY = rF32(-bsU, bsU)
			zrec.col = rl.Brown
			zrec.fade = 0.8

			zfx.recs = append(zfx.recs, zrec)

			num--
			if num == 0 {
				break
			}
		}

		gs.FX.Fx = append(gs.FX.Fx, zfx)

	}

}
func makeBlokGenNoRecNoCntr() xblok { //MARK:MAKE GENERIC BLOCK NO CENTER NO REC
	zblok := xblok{}
	zblok.fade = 1
	zblok.color = rl.White
	zblok.onoff = true

	return zblok
}
func makeBlokGeneric(siz float32, cnt rl.Vector2) xblok { //MARK:MAKE GENERIC BLOCK

	zblok := xblok{}
	zblok.fade = 1
	zblok.color = ranCol()
	zblok.cnt = cnt
	zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
	return zblok
}
func makeBlokGenRandom(siz float32) xblok { //MARK:MAKE GENERIC BLOCK RANDOM POSITION
	zblok := xblok{}
	zblok.fade = 1
	zblok.color = ranCol()

	x := gs.Level.LevRecInner.X + bsU
	y := gs.Level.LevRecInner.Y + bsU
	x = rF32(x, x+gs.Level.LevRecInner.Width-bsU3)
	y = rF32(y, y+gs.Level.LevRecInner.Width-bsU3)

	zblok.rec = rl.NewRectangle(x, y, siz, siz)
	zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Height/2)
	zblok.crec = zblok.rec

	return zblok

}
func makeBlokGenericRanCentr(siz float32) xblok { //MARK:MAKE GENERIC BLOCK RANDOM CENTRE

	zblok := xblok{}
	zblok.fade = 1
	zblok.color = ranCol()
	for {
		zblok.cnt = findRanCntV2()
		if zblok.cnt.X+siz/2 < gs.Level.LevRec.X+gs.Level.LevRec.Width && zblok.cnt.Y+siz/2 < gs.Level.LevRec.Y+gs.Level.LevRec.Width {
			break
		}
	}

	zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
	zblok.crec = zblok.rec
	return zblok
}
func makeCnt(blok xblok) rl.Vector2 {
	v2 := rl.NewVector2(blok.rec.X+blok.rec.Width/2, blok.rec.Y+blok.rec.Height/2)
	return v2
}
func makemovebloks() { //MARK:MAKE MOVE BLOKS

	zblok := xblok{}
	bloksiz := bsU2
	blokimg := gs.Render.Walltiles[rInt(0, len(gs.Render.Walltiles))]
	zblok.fade = 1

	for a := 0; a < len(gs.Level.Level); a++ {

		bloknum := rInt(0, 5)

		for b := 0; b < bloknum; b++ {
			addtolevel := false
			countbreak := 100
			for {

				zblok = makeBlokGenNoRecNoCntr()

				canadd := true
				zblok.rec, canadd = findRecPos(bloksiz, a)
				zblok.cnt = makeCnt(zblok)

				zblok.color = ranPink()
				zblok.img = blokimg
				zblok.velX = 0
				zblok.velY = 0
				zblok.vel = bsU / 2

				if len(gs.Level.Level[a].movBloks) > 0 {
					for c := 0; c < len(gs.Level.Level[a].movBloks); c++ {
						if rl.CheckCollisionRecs(zblok.rec, gs.Level.Level[a].movBloks[c].rec) {
							canadd = false
						}
					}
				}
				if canadd {
					countbreak = 0
					addtolevel = true
				}
				countbreak--
				if countbreak <= 0 {
					break
				}
			}
			if addtolevel {
				choose := rInt(1, 3)
				switch choose {

				case 2: //UD
					zblok.velY = rF32(zblok.vel/4, zblok.vel)
					if roll12() == 12 {
						zblok.velX = rF32(zblok.vel/4, zblok.vel)
					}
					zblok.movType = 2 //UD
					gs.Level.Level[a].movBloks = append(gs.Level.Level[a].movBloks, zblok)
				case 1: //LR
					zblok.velX = rF32(zblok.vel/4, zblok.vel)
					if roll12() == 12 {
						zblok.velY = rF32(zblok.vel/4, zblok.vel)
					}
					zblok.movType = 1 //LR
					gs.Level.Level[a].movBloks = append(gs.Level.Level[a].movBloks, zblok)
				}
			}

		}

	}

}

func makeshaders() { //MARK:MAKE SHADERS
	gs.Render.Shader = rl.LoadShader("", "shaders/bloom.fs")
	gs.Render.Shader2 = rl.LoadShader("", "shaders/grayscale.fs")
	gs.Render.Shader3 = rl.LoadShader("", "shaders/sobel.fs")
	gs.Render.RenderTarget = rl.LoadRenderTexture(gs.Core.ScrW32, gs.Core.ScrH32)
}
func makecompanions() { //MARK: MAKE COMPANIONS

	//MR CARROT
	siz := bsU3
	gs.Companions.MrCarrot.rec = rl.NewRectangle(gs.Core.Cnt.X, gs.Core.Cnt.Y, siz, siz)
	gs.Companions.MrCarrot.imgl = rl.NewRectangle(0, 248, 38, 38)
	gs.Companions.MrCarrot.imgr = rl.NewRectangle(228, 248, 38, 38)
	gs.Companions.MrCarrot.vel = bsU / 5
	gs.Companions.MrCarrot.velx = rF32(-gs.Companions.MrCarrot.vel, gs.Companions.MrCarrot.vel)
	gs.Companions.MrCarrot.vely = rF32(-gs.Companions.MrCarrot.vel, gs.Companions.MrCarrot.vel)
	gs.Companions.MrCarrot.hp = 5
	gs.Companions.MrCarrot.hpmax = gs.Companions.MrCarrot.hp
	gs.Companions.MrCarrot.frames = 5
	gs.Companions.MrCarrot.timer = gs.Core.Fps * rI32(1, 5)

	//MR PLANTY
	gs.Companions.MrPlanty.rec = rl.NewRectangle(gs.Core.Cnt.X, gs.Core.Cnt.Y, siz, siz)
	gs.Companions.MrPlanty.imgl = rl.NewRectangle(0, 454, 44, 44)
	gs.Companions.MrPlanty.imgr = rl.NewRectangle(352, 454, 44, 44)
	gs.Companions.MrPlanty.vel = bsU / 5
	gs.Companions.MrPlanty.velx = rF32(-gs.Companions.MrPlanty.vel, gs.Companions.MrPlanty.vel)
	gs.Companions.MrPlanty.vely = rF32(-gs.Companions.MrPlanty.vel, gs.Companions.MrPlanty.vel)
	gs.Companions.MrPlanty.hp = 5
	gs.Companions.MrPlanty.hpmax = gs.Companions.MrPlanty.hp
	gs.Companions.MrPlanty.frames = 7

	//ALIEN
	gs.Companions.MrAlien.rec = rl.NewRectangle(gs.Core.Cnt.X, gs.Core.Cnt.Y, siz, siz)
	gs.Companions.MrAlien.img = gs.Render.Alien[0]
	gs.Companions.MrAlien.vel = bsU / 5
	gs.Companions.MrAlien.velx = rF32(-gs.Companions.MrAlien.vel, gs.Companions.MrAlien.vel)
	gs.Companions.MrAlien.vely = rF32(-gs.Companions.MrAlien.vel, gs.Companions.MrAlien.vel)
	gs.Companions.MrAlien.hp = 5
	gs.Companions.MrAlien.hpmax = gs.Companions.MrAlien.hp
	gs.Companions.MrAlien.timer = gs.Core.Fps * rI32(3, 8)

}
func makefxinitial() { //MARK:MAKE FX INITAL

	//RAIN
	num := 300
	for a := 0; a < num; a++ {
		siz := rF32(2, 5)
		rec := rl.NewRectangle(rF32(0, gs.Core.ScrWF32), rF32(-gs.Core.ScrHF32, -bsU), siz, siz)
		gs.FX.Rain = append(gs.FX.Rain, rec)
	}

	//SCAN LINES
	y := float32(-2)
	x := float32(0)
	change := float32(3)

	for {
		gs.FX.ScanlineV2 = append(gs.FX.ScanlineV2, rl.NewVector2(x, y))
		y += change
		if y >= gs.Core.ScrHF32+1 {
			break
		}
	}

}
func makebosses() { //MARK:MAKE BOSSES
	siz := bsU8
	zboss := xboss{}
	zboss.hp = 20
	zboss.timer = gs.Core.Fps * 4
	zboss.hpmax = zboss.hp
	zboss.img = rl.NewRectangle(308, 1561, 48, 48)
	zboss.xl = zboss.img.X
	zboss.yt = zboss.img.Y
	zboss.vel = rF32(bsU/8, bsU/4)
	zboss.velX = rF32(-zboss.vel, zboss.vel)
	zboss.velY = rF32(-zboss.vel, zboss.vel)
	zboss.rec = rl.NewRectangle(gs.Core.Cnt.X-siz/2, gs.Level.LevRecInner.Y+bsU2, siz, siz)
	zboss.cnt = gs.Core.Cnt
	zboss.crec = zboss.rec
	zboss.crec.X += zboss.rec.Width / 4
	zboss.crec.Y += zboss.rec.Height / 5
	zboss.crec.Width -= zboss.rec.Width / 2
	zboss.crec.Height -= zboss.rec.Height / 5
	zboss.atkType = rInt(1, 4)
	gs.Level.Bosses = append(gs.Level.Bosses, zboss)
	zboss.img = rl.NewRectangle(450, 1561, 48, 48)
	zboss.xl = zboss.img.X
	zboss.vel = rF32(bsU/8, bsU/4)
	zboss.velX = rF32(-zboss.vel, zboss.vel)
	zboss.velY = rF32(-zboss.vel, zboss.vel)
	zboss.atkType = rInt(1, 4)
	gs.Level.Bosses = append(gs.Level.Bosses, zboss)
	zboss.img = rl.NewRectangle(593, 1561, 48, 48)
	zboss.xl = zboss.img.X
	zboss.vel = rF32(bsU/8, bsU/4)
	zboss.velX = rF32(-zboss.vel, zboss.vel)
	zboss.velY = rF32(-zboss.vel, zboss.vel)
	zboss.atkType = rInt(1, 4)
	gs.Level.Bosses = append(gs.Level.Bosses, zboss)
}
func makeimgs() { //MARK:MAKE IMGS

	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(2, 36, 18, 18))       // 0 BARREL
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(24, 36, 16, 16))      // 1 HP POTION
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(42, 36, 16, 16))      // 2 PLAYER HP ICON
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(60, 36, 14, 14))      // 3 THROWING AXE
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(77, 36, 18, 18))      // 4 SANTA
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(99, 36, 13, 13))      // 5 BOUNCE PROJECTILE
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(115, 36, 16, 16))     // 6 ESCAPE VINE
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(132, 35, 17, 17))     // 7 SKELETON KEY
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(151, 35, 16, 16))     // 8 APPLE
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(170, 36, 17, 17))     // 9 PLANT COMPANION
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(192, 36, 18, 18))     // 10 MEDI KIT
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(215, 38, 16, 16))     // 11 WALLET
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(235, 38, 15, 15))     // 12 RECHARGE
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(1082, 234, 32, 32))   // 13 SHOP1
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(993, 156, 104, 104))  // 14 SHOP2
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(1132, 228, 17, 17))   // 15 SNOW1
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(1157, 229, 15, 15))   // 16 SNOW2
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(1182, 228, 15, 15))   // 17 SNOW3
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(254, 37, 16, 16))     // 18 TURRET
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(274, 38, 14, 14))     // 19 NINJA STAR
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(292, 36, 16, 16))     // 20 OIL BARREL
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(344, 60, 15, 15))     // 21 SWITCH 1
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(361, 60, 15, 15))     // 22 SWITCH 2
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(445, 58, 17, 17))     // 23 CHEST
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(717, 37, 12, 12))     // 24 SWITCH ARROW
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(998, 414, 200, 200))  // 25 NIGHT REC
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(731, 33, 18, 18))     // 26 SPINNING TRAP
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(256, 16, 16, 16))     // 27 POWERUP BLOK
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(820, 34, 18, 18))     // 28 MAP
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(842, 32, 18, 18))     // 29 FIRE TRAIL
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(863, 35, 16, 16))     // 30 INVISIBILITY
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(883, 33, 17, 17))     // 31 COFFEE
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(905, 34, 18, 18))     // 32 TELEPORT
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(927, 35, 16, 16))     // 33 ATTACK RANGE
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(945, 35, 16, 16))     // 34 ATTACK DAMAGE
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(587, 58, 15, 15))     // 35 ORBITAL
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(607, 58, 16, 16))     // 36 CHAIN LIGHTNING
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(624, 57, 16, 16))     // 37 HP RING
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(641, 57, 16, 16))     // 38 ARMOR SHIELD
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(661, 56, 18, 18))     // 39 ANCHOR PAUSE ENEMY
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(681, 55, 18, 18))     // 40 UMBRELLA RAIN
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(705, 57, 14, 14))     // 41 DASH
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(724, 56, 18, 18))     // 42 SOCKS POISON GAS
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(748, 56, 17, 17))     // 43 CHERRIES
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(770, 58, 15, 15))     // 44 FISH FLOOD
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(789, 57, 18, 18))     // 45 CAKE BIRTHDAY
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(811, 58, 16, 16))     // 46 PEACE NO DAMAGE
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(831, 58, 16, 16))     // 47 ALIEN SPECIAL ENEMY
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(850, 57, 16, 16))     // 48 AIR STRIKE
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(893, 58, 14, 14))     // 49 POWERUP FIREWORKS
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(909, 58, 14, 14))     // 50 CARROT COMPANION
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(870, 55, 18, 18))     // 51 MARIO PLATFORMER LEVEL
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(1132, 80, 64, 64))    // 52 FOOTPRINTS
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(966, 35, 14, 14))     // 53 STAIRCASE
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(41, 1219, 565, 300))  // 54 BITTY KNIGHT LOGO
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(748, 1214, 308, 305)) // 55 GO LOGO
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(10, 1532, 256, 256))  // 56 RAYLIB
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(938, 104, 20, 20))    // 57 BOSS 3 PROJ
	gs.Render.Etc = append(gs.Render.Etc, rl.NewRectangle(841, 818, 128, 128))  // 58 POISON GAS

	gs.Render.Plants = append(gs.Render.Plants, rl.NewRectangle(315, 36, 13, 13))
	gs.Render.Plants = append(gs.Render.Plants, rl.NewRectangle(329, 38, 11, 11))
	gs.Render.Plants = append(gs.Render.Plants, rl.NewRectangle(341, 37, 12, 12))
	gs.Render.Plants = append(gs.Render.Plants, rl.NewRectangle(354, 31, 18, 18))
	gs.Render.Plants = append(gs.Render.Plants, rl.NewRectangle(373, 31, 18, 18))
	gs.Render.Plants = append(gs.Render.Plants, rl.NewRectangle(392, 33, 16, 16))
	gs.Render.Plants = append(gs.Render.Plants, rl.NewRectangle(409, 33, 16, 16))
	gs.Render.Plants = append(gs.Render.Plants, rl.NewRectangle(425, 33, 16, 16))
	gs.Render.Plants = append(gs.Render.Plants, rl.NewRectangle(442, 35, 14, 14))
	gs.Render.Plants = append(gs.Render.Plants, rl.NewRectangle(457, 35, 14, 14))

	x := float32(0)
	y := float32(0)

	//FLOOR
	for {
		gs.Render.Floortiles = append(gs.Render.Floortiles, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x >= 896 {
			break
		}
	}
	//INNER WALLS
	x = 0
	y = 16
	for {
		gs.Render.Walltiles = append(gs.Render.Walltiles, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x >= 256 {
			break
		}
	}
	//BATS
	x = 1
	y = 351
	for {
		gs.Render.Bats = append(gs.Render.Bats, rl.NewRectangle(x, y, 16, 16))
		y += 16
		if y > 383 {
			break
		}
	}
	//KNIGHT MOVE
	x = 935
	gs.Player.Pl.imgWalkX = x
	y = 268
	for a := 0; a < 4; a++ {
		gs.Render.Knight = append(gs.Render.Knight, rl.NewRectangle(x, y, 32, 32))
		y += 32
	}
	//KNIGHT ATTACK
	x = 735
	gs.Player.Pl.imgAtkX = x
	y = 268
	for a := 0; a < 4; a++ {
		gs.Render.Knight = append(gs.Render.Knight, rl.NewRectangle(x, y, 32, 32))
		y += 32
	}
	//BUILDINGS
	x = 0
	y = 132
	for {
		gs.Render.Shrines = append(gs.Render.Shrines, rl.NewRectangle(x, y, 32, 32))
		x += 32
		if x > 64 {
			break
		}
	}
	//SKULLS
	x = 474
	y = 34
	for a := 0; a < 7; a++ {
		gs.Render.Skulls = append(gs.Render.Skulls, rl.NewRectangle(x, y, 16, 16))
		x += 16
	}
	//CANDLES
	x = 586
	y = 34
	for a := 0; a < 3; a++ {
		gs.Render.Candles = append(gs.Render.Candles, rl.NewRectangle(x, y, 16, 16))
		x += 16
	}
	//SIGNS
	x = 634
	y = 34
	for a := 0; a < 5; a++ {
		gs.Render.Signs = append(gs.Render.Signs, rl.NewRectangle(x, y, 16, 16))
		x += 16
	}
	//SPLATS
	x = 0
	y = 1072
	for a := 0; a < 6; a++ {
		gs.Render.Splats = append(gs.Render.Splats, rl.NewRectangle(x, y, 128, 128))
		x += 128
	}
	//STATUES
	x = 0
	y = 175
	for a := 0; a < 16; a++ {
		gs.Render.Statues = append(gs.Render.Statues, rl.NewRectangle(x, y, 48, 48))
		x += 48
	}
	//MUSHROOMS
	x = 751
	y = 34
	for a := 0; a < 4; a++ {
		gs.Render.Mushrooms = append(gs.Render.Mushrooms, rl.NewRectangle(x, y, 16, 16))
		x += 16
	}
	//PATTERNS
	x = 2
	y = 798
	for a := 0; a < 6; a++ {
		gs.Render.Patterns = append(gs.Render.Patterns, rl.NewRectangle(x, y, 128, 128))
		gs.Render.Patterns = append(gs.Render.Patterns, rl.NewRectangle(x, y+128, 128, 128))
		x += 128
	}
	//GEMS
	x = 0
	y = 297
	for a := 0; a < 8; a++ {
		gs.Render.Gems = append(gs.Render.Gems, rl.NewRectangle(x, y, 32, 32))
		x += 32
	}

	//RABBIT1
	gs.Render.Rabbit1.frames = 3
	gs.Render.Rabbit1.xl = 68
	gs.Render.Rabbit1.yt = 335
	gs.Render.Rabbit1.W = 16
	gs.Render.Rabbit1.recTL = rl.NewRectangle(gs.Render.Rabbit1.xl, gs.Render.Rabbit1.yt, gs.Render.Rabbit1.W, gs.Render.Rabbit1.W)

	//FIREBALL PLAYER
	gs.Render.FireballPlayer.frames = 3
	gs.Render.FireballPlayer.xl = 1
	gs.Render.FireballPlayer.yt = 60
	gs.Render.FireballPlayer.W = 16
	gs.Render.FireballPlayer.recTL = rl.NewRectangle(gs.Render.FireballPlayer.xl, gs.Render.FireballPlayer.yt, gs.Render.FireballPlayer.W, gs.Render.FireballPlayer.W)

	//BURN
	gs.Render.Burn.frames = 3
	gs.Render.Burn.xl = 72
	gs.Render.Burn.yt = 60
	gs.Render.Burn.W = 16
	gs.Render.Burn.recTL = rl.NewRectangle(gs.Render.Burn.xl, gs.Render.Burn.yt, gs.Render.Burn.W, gs.Render.Burn.W)

	//STAR
	gs.Render.Star.frames = 3
	gs.Render.Star.xl = 144
	gs.Render.Star.yt = 60
	gs.Render.Star.W = 16
	gs.Render.Star.recTL = rl.NewRectangle(gs.Render.Star.xl, gs.Render.Star.yt, gs.Render.Star.W, gs.Render.Star.W)

	//WATER
	gs.Render.Wateranim.frames = 3
	gs.Render.Wateranim.xl = 208
	gs.Render.Wateranim.yt = 60
	gs.Render.Wateranim.W = 16
	gs.Render.Wateranim.recTL = rl.NewRectangle(gs.Render.Wateranim.xl, gs.Render.Wateranim.yt, gs.Render.Wateranim.W, gs.Render.Wateranim.W)

	//PLANT BULL
	gs.Render.PlantBull.frames = 4
	gs.Render.PlantBull.xl = 276
	gs.Render.PlantBull.yt = 60
	gs.Render.PlantBull.W = 16
	gs.Render.PlantBull.recTL = rl.NewRectangle(gs.Render.PlantBull.xl, gs.Render.PlantBull.yt, gs.Render.PlantBull.W, gs.Render.PlantBull.W)

	//SPIKES
	gs.Render.Spikes.frames = 13
	gs.Render.Spikes.xl = 0
	gs.Render.Spikes.yt = 90
	gs.Render.Spikes.W = 32
	gs.Render.Spikes.recTL = rl.NewRectangle(gs.Render.Spikes.xl, gs.Render.Spikes.yt, gs.Render.Spikes.W, gs.Render.Spikes.W)

	//SPRING
	gs.Render.Spring.frames = 2
	gs.Render.Spring.xl = 382
	gs.Render.Spring.yt = 59
	gs.Render.Spring.W = 16
	gs.Render.Spring.recTL = rl.NewRectangle(gs.Render.Spring.xl, gs.Render.Spring.yt, gs.Render.Spring.W, gs.Render.Spring.W)

	//POISON GAS
	gs.Render.Posiongas.frames = 3
	gs.Render.Posiongas.xl = 686
	gs.Render.Posiongas.yt = 638
	gs.Render.Posiongas.W = 128
	gs.Render.Posiongas.recTL = rl.NewRectangle(gs.Render.Posiongas.xl, gs.Render.Posiongas.yt, gs.Render.Posiongas.W, gs.Render.Posiongas.W)

	//MUSHROOM BULL
	gs.Render.MushBull.frames = 3
	gs.Render.MushBull.xl = 516
	gs.Render.MushBull.yt = 59
	gs.Render.MushBull.W = 16
	gs.Render.MushBull.recTL = rl.NewRectangle(gs.Render.MushBull.xl, gs.Render.MushBull.yt, gs.Render.MushBull.W, gs.Render.MushBull.W)

	//BLADES
	gs.Render.Blades.frames = 7
	gs.Render.Blades.xl = 456
	gs.Render.Blades.yt = 92
	gs.Render.Blades.W = 32
	gs.Render.Blades.recTL = rl.NewRectangle(gs.Render.Blades.xl, gs.Render.Blades.yt, gs.Render.Blades.W, gs.Render.Blades.W)

	//SPEAR
	gs.Render.Spear.frames = 7
	gs.Render.Spear.xl = 1004
	gs.Render.Spear.yt = 5
	gs.Render.Spear.W = 16
	gs.Render.Spear.recTL = rl.NewRectangle(gs.Render.Spear.xl, gs.Render.Spear.yt, gs.Render.Spear.W, 64)

	//FIRETRAIL
	gs.Render.Firetrailanim.frames = 3
	gs.Render.Firetrailanim.xl = 724
	gs.Render.Firetrailanim.yt = 82
	gs.Render.Firetrailanim.W = 16
	gs.Render.Firetrailanim.recTL = rl.NewRectangle(gs.Render.Firetrailanim.xl, gs.Render.Firetrailanim.yt, gs.Render.Firetrailanim.W, 16)

	//ORBITAL
	gs.Render.Orbitalanim.frames = 5
	gs.Render.Orbitalanim.xl = 800
	gs.Render.Orbitalanim.yt = 84
	gs.Render.Orbitalanim.W = 16
	gs.Render.Orbitalanim.recTL = rl.NewRectangle(gs.Render.Orbitalanim.xl, gs.Render.Orbitalanim.yt, gs.Render.Orbitalanim.W, 16)

	//FLOOD
	gs.Render.Floodanim.frames = 3
	gs.Render.Floodanim.xl = 900
	gs.Render.Floodanim.yt = 82
	gs.Render.Floodanim.W = 16
	gs.Render.Floodanim.recTL = rl.NewRectangle(gs.Render.Floodanim.xl, gs.Render.Floodanim.yt, gs.Render.Floodanim.W, 16)
	gs.FX.FloodImg = gs.Render.Floodanim.recTL

	//FISH R
	gs.Render.FishR.frames = 3
	gs.Render.FishR.xl = 724
	gs.Render.FishR.yt = 105
	gs.Render.FishR.W = 16
	gs.Render.FishR.recTL = rl.NewRectangle(gs.Render.FishR.xl, gs.Render.FishR.yt, gs.Render.FishR.W, 16)

	//FISH L
	gs.Render.FishL.frames = 3
	gs.Render.FishL.xl = 794
	gs.Render.FishL.yt = 105
	gs.Render.FishL.W = 16
	gs.Render.FishL.recTL = rl.NewRectangle(gs.Render.FishL.xl, gs.Render.FishL.yt, gs.Render.FishL.W, 16)

	//AIRSTRIKE
	gs.Render.Airstrikeanim.frames = 3
	gs.Render.Airstrikeanim.xl = 664
	gs.Render.Airstrikeanim.yt = 614
	gs.Render.Airstrikeanim.W = 32
	gs.Render.Airstrikeanim.recTL = rl.NewRectangle(gs.Render.Airstrikeanim.xl, gs.Render.Airstrikeanim.yt, gs.Render.Airstrikeanim.W, 32)

	//BOSS1 PROJ
	gs.Render.Boss1anim.frames = 3
	gs.Render.Boss1anim.xl = 977
	gs.Render.Boss1anim.yt = 83
	gs.Render.Boss1anim.W = 16
	gs.Render.Boss1anim.recTL = rl.NewRectangle(gs.Render.Boss1anim.xl, gs.Render.Boss1anim.yt, gs.Render.Boss1anim.W, 16)

	//BOSS2 PROJ
	gs.Render.Boss2anim.frames = 3
	gs.Render.Boss2anim.xl = 866
	gs.Render.Boss2anim.yt = 107
	gs.Render.Boss2anim.W = 16
	gs.Render.Boss2anim.recTL = rl.NewRectangle(gs.Render.Boss2anim.xl, gs.Render.Boss2anim.yt, gs.Render.Boss2anim.W, 16)

	//ALIEN
	x = 1008
	y = 770

	for {
		gs.Render.Alien = append(gs.Render.Alien, rl.NewRectangle(x, y, 64, 64))
		x += 64
		if x >= 1200 {
			x = 1008
			y += 64
		}
		if y >= 1026 {
			break
		}
	}

}

// MARK: CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE
func inp() { //MARK:INP

	//OPTIONS ON

	if rl.IsKeyPressed(rl.KeyEscape) && !gs.UI.Intro && !gs.Mario.MarioOn && !gs.Player.Died && !gs.Level.NextLevelScreen {

		if gs.UI.OptionsOn && !gs.Level.Exiton && !gs.Shop.ShopOn && !gs.Level.LevMapOn && !gs.Player.InvenOn && !gs.UI.CreditsOn && !gs.UI.HelpOn && !gs.Timing.TimesOn {
			if gs.UI.OptionsChange {
				savesettings()
			}
			gs.UI.OptionsOn = false
			gs.UI.CreditsOn = false
			gs.UI.HelpOn = false
			gs.Level.Exiton = false
			gs.Core.Pause = false
		} else if gs.UI.OptionsOn && gs.Level.Exiton {
			gs.Level.Exiton = false
		} else if !gs.UI.OptionsOn && !gs.Shop.ShopOn && !gs.Level.LevMapOn && !gs.Player.InvenOn && !gs.UI.CreditsOn && !gs.UI.HelpOn && !gs.Timing.TimesOn {
			gs.Core.Pause = true
			gs.UI.CreditsOn = false
			gs.UI.HelpOn = false
			gs.Level.Exiton = false
			gs.UI.OptionsOn = true
			gs.UI.OptionsChange = false
			gs.UI.OptionNum = 0
		} else if !gs.UI.OptionsOn && gs.Shop.ShopOn && !gs.Level.LevMapOn && !gs.Player.InvenOn && !gs.UI.CreditsOn && !gs.UI.HelpOn && !gs.Timing.TimesOn {
			gs.Shop.ShopOn = false
			gs.Player.Pl.cnt.Y = gs.Shop.ShopExitY
			upPlayerRec()
			gs.Core.Pause = false
		} else if !gs.UI.OptionsOn && !gs.Shop.ShopOn && gs.Level.LevMapOn && !gs.Player.InvenOn && !gs.UI.CreditsOn && !gs.UI.HelpOn && !gs.Timing.TimesOn {
			gs.Level.LevMapOn = false
			gs.Core.Pause = false
		} else if !gs.UI.OptionsOn && !gs.Shop.ShopOn && !gs.Level.LevMapOn && gs.Player.InvenOn && !gs.UI.CreditsOn && !gs.UI.HelpOn && !gs.Timing.TimesOn {
			gs.Player.InvenOn = false
			gs.Core.Pause = false
		} else if gs.UI.OptionsOn && !gs.Shop.ShopOn && !gs.Level.LevMapOn && !gs.Player.InvenOn && gs.UI.CreditsOn && !gs.UI.HelpOn && !gs.Timing.TimesOn {
			gs.UI.CreditsOn = false
			gs.UI.OptionsOn = true
		} else if gs.UI.OptionsOn && !gs.Shop.ShopOn && !gs.Level.LevMapOn && !gs.Player.InvenOn && !gs.UI.CreditsOn && gs.UI.HelpOn && !gs.Timing.TimesOn {
			gs.UI.HelpOn = false
		} else if !gs.UI.OptionsOn && !gs.Shop.ShopOn && !gs.Level.LevMapOn && !gs.Player.InvenOn && !gs.UI.CreditsOn && !gs.UI.HelpOn && gs.Timing.TimesOn {
			gs.Timing.TimesOn = false
			restartgame()
		}

	} else if gs.UI.Intro {
		if rl.IsKeyPressed(rl.KeyEscape) || rl.IsGamepadButtonPressed(0, rl.GamepadButtonMiddleRight) {
			if gs.UI.OptionsOn {
				if gs.UI.OptionsChange {
					savesettings()
				}
				gs.UI.OptionsOn = false
			} else {
				gs.UI.OptionsOn = true
				gs.UI.OptionsChange = false
				gs.UI.OptionNum = 0
			}
		}

		if gs.UI.OptionsOn && rl.IsGamepadButtonPressed(0, 6) {
			if gs.UI.OptionsChange {
				savesettings()
			}
			gs.UI.OptionsOn = false
		}
	}
	if gs.Input.UseController {
		if rl.IsGamepadButtonPressed(0, rl.GamepadButtonMiddleRight) && !gs.Player.InvenOn && !gs.Shop.ShopOn && !gs.Level.LevMapOn && !gs.UI.Intro && !gs.Mario.MarioOn && !gs.Player.Died && !gs.Level.NextLevelScreen {
			if gs.UI.OptionsOn {
				if gs.UI.OptionsChange {
					savesettings()
				}
				gs.UI.CreditsOn = false
				gs.UI.HelpOn = false
				gs.Level.Exiton = false
				gs.UI.OptionsOn = false
				gs.Core.Pause = false
			} else {
				gs.UI.CreditsOn = false
				gs.UI.HelpOn = false
				gs.Level.Exiton = false
				gs.UI.OptionsChange = false
				gs.UI.OptionsOn = true
				gs.UI.OptionNum = 0
				gs.Core.Pause = true
			}
		}
	}

	if !gs.UI.Intro && !gs.Mario.MarioOn && !gs.Player.Died {

		//INVEN ON
		if rl.IsKeyPressed(rl.KeyTab) && !gs.UI.OptionsOn && !gs.Level.LevMapOn && !gs.Shop.ShopOn && !gs.Level.NextLevelScreen {
			if gs.Player.InvenOn {
				gs.Player.InvenOn = false
				gs.Core.Pause = false
			} else {
				gs.Player.InvenOn = true
				gs.Core.Pause = true
			}
		}

		//MAP ON
		if rl.IsKeyPressed(rl.KeyRightControl) && !gs.UI.OptionsOn && !gs.Player.InvenOn && !gs.Shop.ShopOn && !gs.Level.NextLevelScreen {

			if gs.Level.LevMapOn {
				gs.Level.LevMapOn = false
				gs.Core.Pause = false
			} else {
				gs.Level.LevMapOn = true
				gs.Core.Pause = true
			}
		}
		if gs.Input.UseController {
			if rl.IsGamepadButtonPressed(0, 6) && !gs.UI.OptionsOn && !gs.Player.InvenOn && !gs.Shop.ShopOn && !gs.Timing.TimesOn && !gs.Level.NextLevelScreen && !gs.UI.HelpOn && !gs.UI.CreditsOn {
				if gs.Level.LevMapOn {
					gs.Level.LevMapOn = false
					gs.Core.Pause = false
				} else {
					gs.Level.LevMapOn = true
					gs.Core.Pause = true
				}
			} else if rl.IsGamepadButtonPressed(0, 6) && gs.UI.OptionsOn && !gs.Shop.ShopOn && !gs.Timing.TimesOn && !gs.Level.NextLevelScreen && !gs.UI.HelpOn && !gs.UI.CreditsOn {
				if gs.UI.OptionsChange {
					savesettings()
				}
				gs.UI.OptionsOn = false
				gs.Core.Pause = false

			} else if rl.IsGamepadButtonPressed(0, 6) && gs.Player.InvenOn && !gs.Shop.ShopOn && !gs.Timing.TimesOn && !gs.Level.NextLevelScreen && !gs.UI.HelpOn && !gs.UI.CreditsOn {
				gs.Player.InvenOn = false
				gs.Core.Pause = false
			} else if rl.IsGamepadButtonPressed(0, 6) && gs.UI.OptionsOn && !gs.Player.InvenOn && !gs.Shop.ShopOn && gs.Timing.TimesOn && !gs.Level.NextLevelScreen && !gs.UI.HelpOn && !gs.UI.CreditsOn {
				gs.Timing.TimesOn = false
			} else if rl.IsGamepadButtonPressed(0, 6) && gs.UI.OptionsOn && !gs.Player.InvenOn && !gs.Shop.ShopOn && !gs.Timing.TimesOn && !gs.Level.NextLevelScreen && gs.UI.HelpOn && !gs.UI.CreditsOn {
				gs.UI.HelpOn = false
			} else if rl.IsGamepadButtonPressed(0, 6) && gs.UI.OptionsOn && !gs.Player.InvenOn && !gs.Shop.ShopOn && !gs.Timing.TimesOn && !gs.Level.NextLevelScreen && !gs.UI.HelpOn && gs.UI.CreditsOn {
				gs.UI.CreditsOn = false
			}
			//INVEN ON
			if rl.IsGamepadButtonPressed(0, 5) && !gs.UI.OptionsOn && !gs.Level.LevMapOn && !gs.Shop.ShopOn && !gs.Timing.TimesOn && !gs.Level.NextLevelScreen && !gs.UI.HelpOn && !gs.UI.CreditsOn {
				if gs.Player.InvenOn {
					gs.Player.InvenOn = false
					gs.Core.Pause = false
				} else {
					gs.Player.InvenOn = true
					gs.Core.Pause = true
				}
			}
		}

		//PLAYER
		if !gs.Core.Pause {
			if !gs.Player.Pl.escape {
				if rl.IsKeyPressed(rl.KeySpace) {
					if gs.Player.Pl.atkTimer == 0 {
						rl.PlaySound(gs.Audio.Sfx[0])
						gs.Player.ChainLightingSwingOnOff = false
						gs.Player.Pl.atk = true
						gs.Player.Pl.atkTimer = gs.Core.Fps / 3
						gs.Player.Pl.img.X = gs.Player.Pl.imgAtkX
						if gs.Player.Mods.fireball {
							makeProjectile("fireball")
						}
					}
					gs.Input.KeypressT = gs.Core.Fps / 2
				}
				gs.Player.Pl.move = false

				if !gs.Player.Pl.slide {
					if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
						if checkplayermove(1) {
							gs.Player.Pl.cnt.Y -= gs.Player.Pl.vel
						}
						gs.Player.Pl.direc = 1
						if gs.Player.Pl.atkTimer == 0 {
							gs.Player.Pl.move = true
						}
					}
					if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
						if checkplayermove(3) {
							gs.Player.Pl.cnt.Y += gs.Player.Pl.vel
						}
						gs.Player.Pl.direc = 3
						if gs.Player.Pl.atkTimer == 0 {
							gs.Player.Pl.move = true
						}
					}
					if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
						if checkplayermove(4) {
							gs.Player.Pl.cnt.X -= gs.Player.Pl.vel
						}
						gs.Player.Pl.direc = 4
						if gs.Player.Pl.atkTimer == 0 {
							gs.Player.Pl.move = true
						}
					}
					if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
						if checkplayermove(2) {
							gs.Player.Pl.cnt.X += gs.Player.Pl.vel
						}
						gs.Player.Pl.direc = 2
						if gs.Player.Pl.atkTimer == 0 {
							gs.Player.Pl.move = true
						}
					}
				}

				//CONTROLLER
				if gs.Input.UseController {
					if rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
						if gs.Player.Pl.atkTimer == 0 {
							rl.PlaySound(gs.Audio.Sfx[0])
							gs.Player.ChainLightingSwingOnOff = false
							gs.Player.Pl.atk = true
							gs.Player.Pl.atkTimer = gs.Core.Fps / 3
							gs.Player.Pl.img.X = gs.Player.Pl.imgAtkX
							if gs.Player.Mods.fireball {
								makeProjectile("fireball")
							}
						}
						gs.Input.KeypressT = gs.Core.Fps / 2
					}
					gs.Player.Pl.move = false

					if !gs.Player.Pl.slide {

						if rl.GetGamepadAxisMovement(0, 1) < 0 || rl.IsGamepadButtonDown(0, 1) {
							if checkplayermove(1) {
								gs.Player.Pl.cnt.Y -= gs.Player.Pl.vel
							}
							gs.Player.Pl.direc = 1
							if gs.Player.Pl.atkTimer == 0 {
								gs.Player.Pl.move = true
							}
						}
						if rl.GetGamepadAxisMovement(0, 1) > 0 || rl.IsGamepadButtonDown(0, 3) {
							if checkplayermove(3) {
								gs.Player.Pl.cnt.Y += gs.Player.Pl.vel
							}
							gs.Player.Pl.direc = 3
							if gs.Player.Pl.atkTimer == 0 {
								gs.Player.Pl.move = true
							}
						}

						if rl.GetGamepadAxisMovement(0, 0) < 0 || rl.IsGamepadButtonDown(0, 4) {
							if checkplayermove(4) {
								gs.Player.Pl.cnt.X -= gs.Player.Pl.vel
							}
							gs.Player.Pl.direc = 4
							if gs.Player.Pl.atkTimer == 0 {
								gs.Player.Pl.move = true
							}
						}
						if rl.GetGamepadAxisMovement(0, 0) > 0 || rl.IsGamepadButtonDown(0, 2) {
							if checkplayermove(2) {
								gs.Player.Pl.cnt.X += gs.Player.Pl.vel
							}
							gs.Player.Pl.direc = 2
							if gs.Player.Pl.atkTimer == 0 {
								gs.Player.Pl.move = true
							}
						}

					}
				}

				upPlayerRec()
			}
		}
	}

	/*
		//DEBUG
		if rl.IsKeyPressed(rl.KeyF1) {
			if gs.Core.Debug {
				gs.Core.Debug = false
			} else {
				gs.Core.Debug = true
			}
		}

		//ZOOM
		if rl.IsKeyPressed(rl.KeyKpAdd) {
			if gs.Render.Cam2.Zoom == 1 {
				gs.Render.Cam2.Zoom = 2
			} else if gs.Render.Cam2.Zoom == 2 {
				gs.Render.Cam2.Zoom = 3
			} else if gs.Render.Cam2.Zoom == 3 {
				gs.Render.Cam2.Zoom = 4
			} else if gs.Render.Cam2.Zoom == 4 {
				gs.Render.Cam2.Zoom = 1
			}
			cams()
		}
		if rl.IsKeyPressed(rl.KeyKpSubtract) {
			if gs.Render.Cam2.Zoom == 1 {
				gs.Render.Cam2.Zoom = 4
			} else if gs.Render.Cam2.Zoom == 2 {
				gs.Render.Cam2.Zoom = 1
			} else if gs.Render.Cam2.Zoom == 3 {
				gs.Render.Cam2.Zoom = 2
			} else if gs.Render.Cam2.Zoom == 4 {
				gs.Render.Cam2.Zoom = 3
			}
			cams()
		}
	*/

}

func timers() { //MARK:TIMERS

	if gs.Shop.ShopExitT > 0 {
		gs.Shop.ShopExitT--
	}

	gs.Level.RunT++
	if gs.Level.RunT%gs.Core.Fps == 0 {
		gs.Level.Secs++
		gs.Level.RunT = 0
	}
	if gs.Level.Secs == 60 {
		gs.Level.Secs = 0
		gs.Level.Mins++
	}

	if gs.Level.AnchorT > 0 {
		gs.Level.AnchorT--
	}

	if gs.Level.RoomChangedTimer > 0 {
		gs.Level.RoomChangedTimer--
		if gs.Level.RoomChangedTimer == 1 {
			gs.Level.RoomChanged = false
		}
	}

}
func cams() { //MARK:CAMS

	gs.Render.Cam2.Target = gs.Core.Cnt

	gs.Render.Cam2.Offset.X = gs.Core.ScrWF32 / 2
	gs.Render.Cam2.Offset.Y = gs.Core.ScrHF32 / 2

	if gs.Level.Flipcam && !gs.UI.OptionsOn {
		gs.Render.Cam2.Rotation = 180
	} else if gs.Level.Flipcam && gs.UI.OptionsOn {
		gs.Render.Cam2.Rotation = 0
	} else {
		gs.Render.Cam2.Rotation = 0
	}

}

func initialWindow() { //MARK:INITIAL WINDOW
	rl.SetExitKey(rl.KeyEnd)
	rl.HideCursor()
	gs.Render.Imgs = rl.LoadTexture("img/imgs.png")
	gs.Render.RenderTarget = rl.LoadRenderTexture(gs.Core.ScrW32, gs.Core.ScrH32)

	gs.Core.Pause = true
	gs.UI.Intro = true
	//gs.Render.ShaderOn = true

	//gs.Level.Endgame = true

	gs.UI.HpBarsOn = true
	makesettings()
	makeshaders()
	makeaudio()
	makefxinitial()
	makeimgs()
	makebosses()
	makelevel()
	makecompanions()
	makeplayer()
	maketimes()
	cams()

	gs.Audio.Music = gs.Audio.BackMusic[gs.Audio.BgMusicNum]
	gs.Audio.Music.Looping = true
	rl.SetMasterVolume(4) //CHANGE VOLUME
	rl.SetMusicVolume(gs.Audio.Music, gs.Audio.Volume)
	upvolume()

	gs.Player.DiedRec = rl.NewRectangle(gs.Core.Cnt.X-bsU, gs.Core.Cnt.Y-bsU, bsU2, bsU2)
	gs.Player.DiedIMG = gs.Render.Splats[rInt(0, len(gs.Render.Splats))]
	gs.Level.EndgameT = gs.Core.Fps * 5
	gs.Level.EndgopherRec = rl.NewRectangle(gs.Core.Cnt.X-bsU4, gs.Level.LevRec.Y+gs.Level.LevRec.Height, bsU8, bsU8)

}

func unload() { //MARK:UNLOAD
	rl.UnloadShader(gs.Render.Shader)
	rl.UnloadShader(gs.Render.Shader2)
	rl.UnloadRenderTexture(gs.Render.RenderTarget)
	rl.UnloadTexture(gs.Render.Imgs)

	rl.StopMusicStream(gs.Audio.Music)
	rl.UnloadMusicStream(gs.Audio.Music)
	rl.UnloadMusicStream(gs.Audio.BackMusic[gs.Audio.BgMusicNum])

	for a := 0; a < len(gs.Audio.Sfx); a++ {
		rl.UnloadSound(gs.Audio.Sfx[a])
	}
	rl.CloseAudioDevice()

}

func main() { //MARK:MAIN
  gs.UI.IntroT1 = gs.Core.Fps * 2
  gs.UI.IntroT2 = gs.Core.Fps * 2
  gs.UI.IntroT3 = gs.Core.Fps * 3

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	//rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(1920, 1080, "BITTY KNIGHT - unklnik.com") //GET SCREEN SIZE
	rl.InitAudioDevice()
	rl.SetWindowMonitor(0)
	rl.SetWindowState(rl.FlagBorderlessWindowedMode)
	gs.Core.ScrW, gs.Core.ScrH = rl.GetScreenWidth(), rl.GetScreenHeight()

	//rl.ToggleFullscreen()

	rl.SetWindowSize(gs.Core.ScrW, gs.Core.ScrH)

	gs.Core.ScrW32, gs.Core.ScrH32 = int32(gs.Core.ScrW), int32(gs.Core.ScrH)
	gs.Core.ScrWF32, gs.Core.ScrHF32 = float32(gs.Core.ScrW), float32(gs.Core.ScrH)
	gs.Core.Cnt = rl.NewVector2(gs.Core.ScrWF32/2, gs.Core.ScrHF32/2)
	if gs.Core.ScrH >= 2160 {
		gs.Render.Cam2.Zoom = 3
	} else if gs.Core.ScrH >= 1440 && gs.Core.ScrH < 2160 {
		gs.Render.Cam2.Zoom = 2
	} else if gs.Core.ScrH == 1200 {
		gs.Render.Cam2.Zoom = 1.65
	} else if gs.Core.ScrH > 1050 && gs.Core.ScrH < 1200 {
		gs.Render.Cam2.Zoom = 1.5
	} else if gs.Core.ScrH >= 990 && gs.Core.ScrH <= 1050 {
		gs.Render.Cam2.Zoom = 1.35
	} else if gs.Core.ScrH >= 900 && gs.Core.ScrH < 990 {
		gs.Render.Cam2.Zoom = 1.2
	} else if gs.Core.ScrH >= 720 && gs.Core.ScrH < 900 {
		gs.Render.Cam2.Zoom = 1
	} else if gs.Core.ScrH >= 600 && gs.Core.ScrH < 720 {
		gs.Render.Cam2.Zoom = 0.8
	} else if gs.Core.ScrH < 600 && gs.Core.ScrH > 300 {
		gs.Render.Cam2.Zoom = 0.5
	} else if gs.Core.ScrH < 300 {
		gs.Render.Cam2.Zoom = 0.2
	}

	gs.Level.LevX = gs.Core.Cnt.X - gs.Level.LevW/2
	gs.Level.LevY = gs.Core.Cnt.Y - gs.Level.LevW/2
	gs.Level.LevRec = rl.NewRectangle(gs.Level.LevX, gs.Level.LevY, gs.Level.LevW, gs.Level.LevW)
	gs.Level.LevRecInner = gs.Level.LevRec
	gs.Level.LevRecInner.X += gs.Level.BorderWallBlokSiz
	gs.Level.LevRecInner.Y += gs.Level.BorderWallBlokSiz
	gs.Level.LevRecInner.Width -= gs.Level.BorderWallBlokSiz * 2
	gs.Level.LevRecInner.Height -= gs.Level.BorderWallBlokSiz * 2

	initialWindow() //INITIAL INSIDE WINDOW

	rl.SetTargetFPS(gs.Core.Fps)

	for !rl.WindowShouldClose() {
		gs.Core.Frames++
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if gs.Render.ShaderOn && !gs.Render.Shader2On && !gs.Render.Shader3On { //BLOOM SHADER
			rl.BeginTextureMode(gs.Render.RenderTarget) // Enable drawing to texture
			rl.ClearBackground(rl.Black)
			drawnocamBG()

			rl.BeginMode2D(gs.Render.Cam2)

			drawcam()

			rl.EndMode2D()

			drawnocam()

			rl.EndTextureMode()

			rl.BeginShaderMode(gs.Render.Shader)
			rl.DrawTextureRec(gs.Render.RenderTarget.Texture, rl.NewRectangle(0, 0, float32(gs.Render.RenderTarget.Texture.Width), float32(-gs.Render.RenderTarget.Texture.Height)), rl.NewVector2(0, 0), rl.White)
			rl.EndShaderMode()

		} else if gs.Render.ShaderOn && gs.Render.Shader2On && !gs.Render.Shader3On {
			rl.BeginTextureMode(gs.Render.RenderTarget) // Enable drawing to texture
			rl.ClearBackground(rl.Black)
			drawnocamBG()

			rl.BeginMode2D(gs.Render.Cam2)

			drawcam()

			rl.EndMode2D()

			drawnocam()

			rl.EndTextureMode()

			rl.BeginTextureMode(gs.Render.RenderTarget) // Enable drawing to texture
			rl.BeginShaderMode(gs.Render.Shader)
			rl.DrawTextureRec(gs.Render.RenderTarget.Texture, rl.NewRectangle(0, 0, float32(gs.Render.RenderTarget.Texture.Width), float32(-gs.Render.RenderTarget.Texture.Height)), rl.NewVector2(0, 0), rl.White)
			rl.EndShaderMode()
			rl.EndTextureMode()
			rl.BeginShaderMode(gs.Render.Shader2)
			rl.DrawTextureRec(gs.Render.RenderTarget.Texture, rl.NewRectangle(0, 0, float32(gs.Render.RenderTarget.Texture.Width), float32(-gs.Render.RenderTarget.Texture.Height)), rl.NewVector2(0, 0), rl.White)
			rl.EndShaderMode()
		} else if gs.Render.ShaderOn && !gs.Render.Shader2On && gs.Render.Shader3On {
			rl.BeginTextureMode(gs.Render.RenderTarget) // Enable drawing to texture
			rl.ClearBackground(rl.Black)
			drawnocamBG()

			rl.BeginMode2D(gs.Render.Cam2)

			drawcam()

			rl.EndMode2D()

			drawnocam()

			rl.EndTextureMode()

			rl.BeginTextureMode(gs.Render.RenderTarget) // Enable drawing to texture
			rl.BeginShaderMode(gs.Render.Shader)
			rl.DrawTextureRec(gs.Render.RenderTarget.Texture, rl.NewRectangle(0, 0, float32(gs.Render.RenderTarget.Texture.Width), float32(-gs.Render.RenderTarget.Texture.Height)), rl.NewVector2(0, 0), rl.White)
			rl.EndShaderMode()
			rl.EndTextureMode()
			rl.BeginShaderMode(gs.Render.Shader3)
			rl.DrawTextureRec(gs.Render.RenderTarget.Texture, rl.NewRectangle(0, 0, float32(gs.Render.RenderTarget.Texture.Width), float32(-gs.Render.RenderTarget.Texture.Height)), rl.NewVector2(0, 0), rl.White)
			rl.EndShaderMode()
		} else if !gs.Render.ShaderOn && gs.Render.Shader2On && !gs.Render.Shader3On {
			rl.BeginTextureMode(gs.Render.RenderTarget) // Enable drawing to texture
			rl.ClearBackground(rl.Black)
			drawnocamBG()

			rl.BeginMode2D(gs.Render.Cam2)

			drawcam()

			rl.EndMode2D()

			drawnocam()

			rl.EndTextureMode()

			rl.BeginShaderMode(gs.Render.Shader2)
			rl.DrawTextureRec(gs.Render.RenderTarget.Texture, rl.NewRectangle(0, 0, float32(gs.Render.RenderTarget.Texture.Width), float32(-gs.Render.RenderTarget.Texture.Height)), rl.NewVector2(0, 0), rl.White)
			rl.EndShaderMode()
		} else if !gs.Render.ShaderOn && !gs.Render.Shader2On && gs.Render.Shader3On {
			rl.BeginTextureMode(gs.Render.RenderTarget) // Enable drawing to texture
			rl.ClearBackground(rl.Black)
			drawnocamBG()

			rl.BeginMode2D(gs.Render.Cam2)

			drawcam()

			rl.EndMode2D()

			drawnocam()

			rl.EndTextureMode()

			rl.BeginShaderMode(gs.Render.Shader3)
			rl.DrawTextureRec(gs.Render.RenderTarget.Texture, rl.NewRectangle(0, 0, float32(gs.Render.RenderTarget.Texture.Width), float32(-gs.Render.RenderTarget.Texture.Height)), rl.NewVector2(0, 0), rl.White)
			rl.EndShaderMode()
		} else { //NO BLOOM SHADER
			rl.ClearBackground(rl.Black)
			drawnocamBG()
			rl.BeginMode2D(gs.Render.Cam2)
			drawcam()
			rl.EndMode2D()
			drawnocam()
		}

		drawnoRender()

		up() //UPDATE
		rl.EndDrawing()
	}

	unload() //UNLOAD

	rl.CloseAudioDevice()

	rl.CloseWindow()
}
