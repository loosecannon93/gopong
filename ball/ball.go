package ball

import ( 
    "github.com/banthar/Go-SDL/sdl"
    )

type Ball struct {
    X int16
    Y int16
    Xspeed int16
    Yspeed int16
    Image *sdl.Surface
}

func (b *Ball) Move() {
    b.X += b.Xspeed
    b.Y += b.Yspeed
}

