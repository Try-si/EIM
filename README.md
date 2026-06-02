# Description

EIM (Ebiten Input Manager) est un gestionnaire d'entrées simple pour Ebiten.

## Installation

```bash
go get github.com/Try-si/EIM
```

## Utilisation

### Création d'un fichier de configuration

```json
{
    "Bindings": {
        "jump": 32,
        "run": 16,
        "attack": 1
    }
}
```

### Utilisation dans le code

```go
import "github.com/Try-si/EIM"

func main() {
    im := EIM.NewInputManager("config.json")
    
    // Ajouter une liaison
    im.AddBinding("jump", ebiten.KeySpace)
    
    // Vérifier si une touche est pressée
    if im.IsKeyPressed("jump") {
        // Faire quelque chose
    }
}

// ou

func main() { // si le fichier de configuration est créé
    im := EIM.NewInputManager("config.json")
    
    // Vérifier si une touche est pressée
    if im.IsKeyPressed("jump") {
        // Faire quelque chose
    }
}
```
