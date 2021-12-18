# Mancala

Mancala game implementation in Go as an API.

## Usage

If you want to use the game engine:


```go
import (
"fmt"
"github.com/pablocrivella/mancala/engine"
)

func main() {
game := engine.NewGame("Rick", "Morty")

// Player1 plays
game.PlayTurn(0)
}
```

If you want to use the API go to https://go-mancala.herokuapp.com/docs

To create a new game:

```bash
curl -X POST http://localhost:1323/v1/games -H "Content-Type: application/json" --data '{"player1":"Rick","player2":"Morty"}'
```

To show the state of a game:

```bash
curl http://localhost:1323/v1/games/215663b9-b772-4f21-a69a-befb800741c7
```

To perform the next play:

```bash
curl -X PATCH http://localhost:1323/v1/games/5f489bec-ca14-4eaf-b168-d0bcc5c7e1fd -H "Content-Type: application/json" --data '{"pit_index":0}'
```

### Notes

Games expire after 2 hours.

## License

Copyright 2020 [Pablo Crivella](https://pablocrivella.me).
Read [LICENSE](LICENSE) for details.