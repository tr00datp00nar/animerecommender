# :deciduous_tree: animerecommender

This is `animerecommender`. I have put it into a command branch for inclusion into my c Bonzai stateful command tree.

## Installation

If you just want to try it out, grab the release binary with curl and put into your PATH:

```
curl -L https://github.com/tr00datp00nar/animerecommender/releases/latest/download/animerecommender-linux-amd64 -o ~/.local/bin/animerecommender
curl -L https://github.com/tr00datp00nar/animerecommender/releases/latest/download/animerecommender-darwin-amd64 -o ~/.local/bin/animerecommender
curl -L https://github.com/tr00datp00nar/animerecommender/releases/latest/download/animerecommender-darwin-arm64 -o ~/.local/bin/animerecommender
curl -L https://github.com/tr00datp00nar/animerecommender/releases/latest/download/animerecommender-windows-amd64 -o ~/.local/bin/animerecommender
```

Or with `go`:

```shell
go install github.com/tr00datp00nar/animerecommender/cmd/animerecommender@latest
```

Composed

```go
package c

import (
	Z "github.com/rwxrob/bonzai/z"
    "github.com/tr00datp00nar/animerecommender"
)

var Cmd = &Z.Cmd{
	Name:     'c',
    Commands: []*Z.Cmd{help.Cmd, animerecommender.Cmd},
}
```

## Resources

To learn more about Bonzai stateful command trees: https://github.com/rwxrob/bonzai

To see my personal Bonzai stateful command tree: https://github.com/tr00datp00nar/c

To see the original Bonzai stateful command tree z: https://github.com/rwxrob/z
