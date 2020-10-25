// Package openapi provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package openapi

// BoardSide defines model for BoardSide.
type BoardSide struct {

	// The pits of the board side
	Pits   []int64 `json:"pits"`
	Player Player  `json:"player"`

	// The store (big pit) of the board side
	Store int64 `json:"store"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Errors []string `json:"errors"`
}

// Game defines model for Game.
type Game struct {
	BoardSide1 struct {
		// Embedded fields due to inline allOf schema
		// Embedded struct due to allOf(#/components/schemas/BoardSide)
		BoardSide
	} `json:"board_side1"`
	BoardSide2 struct {
		// Embedded fields due to inline allOf schema
		// Embedded struct due to allOf(#/components/schemas/BoardSide)
		BoardSide
	} `json:"board_side2"`

	// The id of the game
	Id string `json:"id"`

	// The result of the game. `0` = `Player1Wins`, `1` = `Player2Wins`, `2` = `Tie`, `3` = `Undefined`.
	Result Result `json:"result"`

	// The turn that needs to play next. `0` = `Player1`, `1` = `Player2`.
	Turn Turn `json:"turn"`
}

// Player defines model for Player.
type Player struct {

	// The name of the player
	Name string `json:"name"`

	// The score of the player
	Score int64 `json:"score"`
}

// Result defines model for Result.
type Result int64

// Turn defines model for Turn.
type Turn int64

// CreateGameJSONBody defines parameters for CreateGame.
type CreateGameJSONBody struct {
	Player1 string `json:"player1"`
	Player2 string `json:"player2"`
}

// UpdateGameJSONBody defines parameters for UpdateGame.
type UpdateGameJSONBody struct {
	PitIndex int64 `json:"pit_index"`
}

// CreateGameRequestBody defines body for CreateGame for application/json ContentType.
type CreateGameJSONRequestBody CreateGameJSONBody

// UpdateGameRequestBody defines body for UpdateGame for application/json ContentType.
type UpdateGameJSONRequestBody UpdateGameJSONBody

