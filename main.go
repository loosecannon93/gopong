package main

import (
    "github.com/banthar/Go-SDL/sdl"
    "./ball"
)

func loadImage(name string) *sdl.Surface {
    image := sdl.Load(name)

    if image == nil {
        panic(sdl.GetError())
    }

    return image

}

const SCREEN_WIDTH = 640
const SCREEN_HEIGHT = 480
func main() {

    if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
        panic(sdl.GetError())
    }

    defer sdl.Quit()
    

    var screen = sdl.SetVideoMode(SCREEN_WIDTH, SCREEN_HEIGHT, 32, 0)

    if screen == nil {
        panic(sdl.GetError())
    }

    sdl.WM_SetCaption("Template", "")

    // create our ball object 
    b := ball.Ball{5,5,1,1,loadImage("ball.png")}
    p1 := ball.Paddle{X:SCREEN_WIDTH/2, Y:50, Image:loadImage("paddle.png") } 
    // main loop 
    for true {

        sdl.EnableKeyRepeat(0,0)
        // Poll For events 
        for ev := sdl.PollEvent(); ev != nil ; ev = sdl.PollEvent() {
            switch e := ev.(type) {
            case *sdl.QuitEvent:
                return
            case *sdl.KeyboardEvent:
                switch  e.Keysym.Sym { 
                    case  sdl.K_LEFT : 
                        p1.MoveLeft()
                    case  sdl.K_RIGHT : 
                        p1.MoveRight()
                }
            default:
            }
        }



        if b.X >= SCREEN_WIDTH || b.X <= 0 { 
            b.Xspeed *= -1 
        }
        if b.Y >= SCREEN_HEIGHT || b.Y <= 0 {
            b.Yspeed *= -1 
        }

        // update objects 
        b.Move()


        // give us a black background 
        screen.FillRect(nil, 0x000000)
        // draw everything 
        screen.Blit(&sdl.Rect{p1.X,p1.Y, 0, 0}, p1.Image, nil)
        screen.Blit(&sdl.Rect{b.X,b.Y, 0, 0}, b.Image, nil)


        screen.Flip()
//        sdl.Delay(25)

    }

}

