package Input

import (
	ETMHelper "github.com/Try-si/ETM/Helper"
	"github.com/hajimehoshi/ebiten/v2"
)

type InputManager struct {
	Bindings map[string]int
}

func NewInputManager() *InputManager {
	im := ETMHelper.Jsontostruct[*InputManager]("config.json")
	return im
}

func (im *InputManager) AddBinding(name string, key int) {
	im.Bindings[name] = key
}

func (im *InputManager) GetBinding(name string) int {
	return im.Bindings[name]
}

func (im *InputManager) SetBinding(name string, key int) {
	im.Bindings[name] = key
}

func (im *InputManager) RemoveBinding(name string) {
	delete(im.Bindings, name)
}

func (im *InputManager) HasBinding(name string) bool {
	_, ok := im.Bindings[name]
	return ok
}

func (im *InputManager) GetBindings() map[string]int {
	return im.Bindings
}

func (im *InputManager) ClearBindings() {
	im.Bindings = make(map[string]int)
}

func (im *InputManager) GetBindingCount() int {
	return len(im.Bindings)
}

func (im *InputManager) IsKeyPressed(binding string) bool {
	key, ok := im.Bindings[binding]
	if !ok {
		return false
	}
	return ebiten.IsKeyPressed(ebiten.Key(key))
}
