package game

import (
	"nowhere-home/src/overlays"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	GetName() string
	Update() error
	Draw(screen *ebiten.Image)
}

type Scene_Manager struct {
	current Scene
	next    Scene
	fadeDir int //1 = fading in, -1 fading out
}

func (s *Scene_Manager) Update() error {
	count := overlays.UpdateFade(s.fadeDir)

	if s.fadeDir == 1 && count >= overlays.FadeAlphaMaxCount {
		if s.next == nil {
			s.fadeDir = 0
		} else {
			s.fadeDir = -1
			s.current = s.next
			s.next = nil
		}
	} else if s.fadeDir == -1 && count <= 0 {
		panic(1)
	}

	return s.current.Update()
}

func (s *Scene_Manager) Draw(screen *ebiten.Image) {
	s.current.Draw(screen)
	overlays.DrawFade(screen)
}

func (s *Scene_Manager) GoTo(scene Scene) {
	s.fadeDir = 1
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
	}
}