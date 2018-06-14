package main
import (
  "fmt"
  "bytes"
  "math/rand"
  "time"
)

type Direction int

const (
  NORTH Direction = 0
  EAST Direction = 1
  SOUTH Direction = 2
  WEST Direction = 3
)

type Ant struct{
  x int
  y int
  moves int
  direction Direction
}

func (a Ant) String() string {
  return fmt.Sprintf("(x %d, y %d, moves %d, dir %d)", a.x,a.y,a.moves,a.direction)
}
type Board struct{
  rows int
  cols int
  squaresOn int
  squaresOff int
  board [][]bool
}
type Game struct{
  b *Board
  langstons_ant *Ant
  w, h int
}
func (a Game) String() string {
  var buf bytes.Buffer
	for x := 0; x < a.w; x++ {
		for y := 0; y < a.h; y++ {
      if x==a.langstons_ant.x && y==a.langstons_ant.y{
        buf.WriteByte(byte('o'))
      }else if a.b.board[x][y]{
        buf.WriteByte(byte('x'))
      }else{
        buf.WriteByte(byte('-'))
      }

		}
		buf.WriteByte(byte('\n'))
	}
	return buf.String()
}

func inits(width, height int) Game {
  a := make([][]bool, width)
  for i := range a {
      a[i] = make([]bool, height)
  }
  for x, row := range a{
    for y, _ := range row{
      a[x][y] = false
    }
  }
  board := &Board{width, height, 0, width*height, a}
  ant := &Ant{width/2,height/2, 0, EAST}
  g := Game{b: board, w:width, h:height, langstons_ant: ant}
  return g
}
func boundsFix(x, y, w, h  int) ( int, int, bool){
  atWall := true
  switch{
    case x < 0:
      x=0
    case y < 0:
      y=0
    case x > w:
      x=w
    case y > h:
      y=h
    default:
      atWall = false
  }
  return x, y, atWall
}
func (g *Game) run(){

  for{
    var isBlackSquare bool

    x,y := g.langstons_ant.x, g.langstons_ant.y
    if g.b.board[x][y] == true{
      isBlackSquare = true
    }else{
      isBlackSquare = false
    }

    if isBlackSquare{
      g.b.board[x][y] = false
      g.langstons_ant.direction =  (g.langstons_ant.direction - 1) % 4
      if g.langstons_ant.direction < 0{
        g.langstons_ant.direction = 3
      }
    }
    if !isBlackSquare{
      g.b.board[x][y] = true
      g.langstons_ant.direction = (g.langstons_ant.direction + 1) % 4
      if g.langstons_ant.direction > 3{
        g.langstons_ant.direction = 0
      }
    }
    var nextx,nexty int
    var atWall bool
    switch g.langstons_ant.direction{
      case NORTH:
        nextx,nexty,atWall = boundsFix(g.langstons_ant.x, g.langstons_ant.y+1, g.w, g.h)
      case EAST:
        nextx,nexty,atWall = boundsFix(g.langstons_ant.x+1, g.langstons_ant.y, g.w, g.h)
      case SOUTH:
        nextx,nexty,atWall = boundsFix(g.langstons_ant.x, g.langstons_ant.y-1, g.w, g.h)
      case WEST:
        nextx,nexty,atWall = boundsFix(g.langstons_ant.x-1, g.langstons_ant.y, g.w, g.h)
    }

    if atWall{
      g.langstons_ant.direction = Direction(rand.Intn(4))
    }else{
      g.langstons_ant.x = nextx
      g.langstons_ant.y = nexty
    }
    fmt.Println(*g)
    time.Sleep(10 * time.Millisecond)
  }

}
func main() {
  g := inits(50,50)
  g.run()
}
