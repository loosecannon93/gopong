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

func main() {

    if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
        panic(sdl.GetError())
    }

    defer sdl.Quit()

    var screen = sdl.SetVideoMode(640, 480, 32, 0)

    if screen == nil {
        panic(sdl.GetError())
    }

    sdl.WM_SetCaption("Template", "")

    // create our ball object 
    b := ball.Ball{5,5,1,1,loadImage("image.jpg")}

    // main loop 
    for true {


        // Poll For events 
        for ev := sdl.PollEvent(); ev != nil ; ev = sdl.PollEvent() {
            switch ev.(type) {
            case *sdl.QuitEvent:
                return
            default:
            }
        }



        // update objects 
        b.Move()


        // give us a black background 
        screen.FillRect(nil, 0x000000)
        // draw everything 
        screen.Blit(&sdl.Rect{b.X,b.Y, 50, 50}, b.Image, nil)


        screen.Flip()
//        sdl.Delay(25)

    }

}

