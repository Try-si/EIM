package EInputManager

import (
	ETMHelper "github.com/Try-si/ETE/ETEHelper"
	"github.com/hajimehoshi/ebiten/v2"
)

type InputManager struct {
	Bindings     map[string]int
	previousKeys map[ebiten.Key]bool
}

func NewInputManager(ConfigPath string) *InputManager {
	im := ETMHelper.JsonToStruct[*InputManager](ConfigPath)
	im.previousKeys = make(map[ebiten.Key]bool)
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

func (im *InputManager) IsKeyJustPressed(binding string) bool {
	key, ok := im.Bindings[binding]
	if !ok {
		return false
	}
	currentPressed := ebiten.IsKeyPressed(ebiten.Key(key))
	wasPressed := im.previousKeys[ebiten.Key(key)]
	return currentPressed && !wasPressed
}

func (im *InputManager) Update() {
	for key := range im.previousKeys {
		im.previousKeys[key] = ebiten.IsKeyPressed(key)
	}
	for _, key := range im.Bindings {
		k := ebiten.Key(key)
		if _, exists := im.previousKeys[k]; !exists {
			im.previousKeys[k] = ebiten.IsKeyPressed(k)
		}
	}
}
