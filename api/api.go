package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	//response code
	Code int

	// Account balance
	Balance int64
}

type Error struct {
	Code int

	// Error Message
	Message string
}
