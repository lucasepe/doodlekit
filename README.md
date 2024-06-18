# `doodlekit`

Welcome to `doodlekit`, an open-source Golang framework for creating animated doodles like old school demoscenes and crack intros. 

Whether you're a nostalgic coder or a modern enthusiast, `doodlekit` allows you to craft captivating GIF animations with ease.

> Check out the [lucasepe/doodles](https://github.com/lucasepe/doodles) repository for examples of stunning animations created with `doodlekit`!

## Features

- **Simple API**: easy-to-use functions to create [stunning animations](https://github.com/lucasepe/doodles/GALLERY_1.md)
- **Composable Scenes**: you can compose one or more existing scenes with yours
- **Flexible Configuration**: options to set frame rate, output directory, and more.
- **Open Source**: contribute and expand the framework with your own scenes and animations.
- **GIF Output**: generates GIFs for sharing and showcasing

## Getting Started

### Usage

Here's a basic example to get you started with **`doodlekit`**:

```go
package main

import (
    "context"
    "github.com/lucasepe/doodlekit"
)

type MyScene struct{}

func (s *MyScene) Init(ctx context.Context) {
    // Initialize your scene
    // gc := doodlekit.Canvas(Ctx)
}

func (s *MyScene) Update(ctx context.Context, dt float64) {
    // Update scene logic
    // gc := doodlekit.Canvas(Ctx)
}

func (s *MyScene) Draw(ctx context.Context) {
    // Draw your scene
    // gc := doodlekit.Canvas(Ctx)
    // Use `gc` to draw your stuff
}

func main() {
    scenes := []doodlekit.Scene{
        // here you can add eventually 
        // one or more existing scenes
        &MyScene{},
    }

    loop := doodlekit.NewLoop("my-doodle", 
        doodlekit.FPS(24), 
        doodlekit.StopAfter(15))
    loop.Run(scenes)
}
```

### Creating a Scene

To create your own scene, implement the `doodlekit.Scene` interface:

```go
type Scene interface {
    Init(ctx context.Context)
    Update(ctx context.Context, dt float64)
    Draw(ctx context.Context)
}
```

- **Init**: initialize your scene, load resources, etc.
- **Update**: update the scene logic, animations, positions, etc.
- **Draw**: render your scene, draw shapes, etc.

---

### Acknowledgements

Special thanks to

- the demoscene community and all contributors who keep the spirit of retro computing alive
- the authors of the following libraries for their inspiring work, which helped spark the idea and development of `doodlekit`: 
  - [fogleman/gg](https://github.com/fogleman/gg), [peterhellberg/gfx](https://github.com/peterhellberg/gfx), [Adafruit-GFX-Library](https://github.com/adafruit/Adafruit-GFX-Library), [tinygo-org/tinydraw](https://github.com/tinygo-org/tinydraw), [zachomedia/go-bdf](https://github.com/zachomedia/go-bdf.git)


### License

This project is licensed under the BSD 2-Clause License. See the [LICENSE](LICENSE) file for details.

