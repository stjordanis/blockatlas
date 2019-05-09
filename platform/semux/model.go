package semux

import "github.com/trustwallet/blockatlas/models"

type Tx struct {
	BlockNumber string        `json:"blockNumber"`
	Data        string        `json:"data"`
	Fee         models.Amount `json:"fee"`
	From        string        `json:"from"`
	Hash        string        `json:"hash"`
	Nonce       string        `json:"nonce"`
	Timestamp   string        `json:"timestamp"`
	To          string        `json:"to"`
	Type        string        `json:"type"`
	Value       models.Amount `json:"value"`
}

type Account struct {
	Address                 string        `json:"address"`
	Available               models.Amount `json:"available"`
	Locked                  models.Amount `json:"locked"`
	Nonce                   string        `json:"nonce"`
	PendingTransactionCount uint64        `json:"pendingTransactionCount"`
	TransactionCount        uint64        `json:"transactionCount"`
}

type GetAccountTransactionsResponse struct {
	Message string `json:"message"`
	Result  []Tx   `json:"result"`
	Success bool   `json:"success"`
}

type GetAccountResponse struct {
	Message string  `json:"message"`
	Result  Account `json:"result"`
	Success bool    `json:"success"`
}
