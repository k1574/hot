package engine

import (
	"container/list"
	"time"
	//_ "image/png"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"hot/m/engine/camera"
	"hot/m/engine/transform"
	"hot/m/engine/vector"
	"hot/m/engine/behaviorer"
)

type Engine struct {
	lastTime time.Time
	Objects *list.List
	DT float64
	WinCfg pixelgl.WindowConfig
	Win *pixelgl.Window
	Cam *camera.Camera
}

func
(eng *Engine)update(){
		eng.Win.Clear(colornames.Whitesmoke)
		eng.setNewDT()
		for e := eng.Objects.Front() ; e != nil ; e = e.Next() {
			o := e.Value.(behaviorer.Behaviorer)
			o.Update()

			od := o.GetO()
			if od == nil {
				continue
			}

			finmat := pixel.IM.ScaledXY(pixel.ZV, od.T.S).
				Rotated(vector.Z, od.T.R).
				Moved(od.T.P).
				Rotated(eng.Cam.T.P.Add(eng.Win.Bounds().Center()),
					eng.Cam.T.R).
				Moved(vector.Z.Sub(eng.Cam.T.P)).
				ScaledXY(eng.Cam.T.P.Add(eng.Win.Bounds().Center()),
					eng.Cam.T.S)

			if od.S != nil {
				od.S.Draw(eng.Win, finmat)
			}
		}
		eng.Win.Update()
}

func
(eng *Engine)setNewDT(){
	eng.DT = time.Since(eng.lastTime).Seconds()
	eng.lastTime = time.Now()
}

func
(eng *Engine)AddBehaviorer(v behaviorer.Behaviorer) {
	eng.Objects.PushBack(v)
	v.Start()
}

func
New(cfg pixelgl.WindowConfig) (*Engine) {
	eng := Engine {
		Objects: list.New(),
		WinCfg: cfg,
		Cam: camera.New(
			transform.New(
				vector.New(1, 1),
				vector.New(1, 1),
				0),
			),
	}

	return &eng
}

func
(eng *Engine)run() {
	var err error

	eng.Win, err = pixelgl.NewWindow(eng.WinCfg)
	if err != nil {
		panic(err)
	}
	eng.Win.SetSmooth(true)

	eng.lastTime = time.Now()
	for !eng.Win.Closed() {
		eng.update()
	}
}

func
(eng *Engine)Run() {
	pixelgl.Run(eng.run)
}

