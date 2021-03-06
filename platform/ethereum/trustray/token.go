package trustray

import (
	"github.com/trustwallet/golibs/types"
)

func (c *Client) GetTokenList(address string, coinIndex uint) ([]types.Token, error) {
	account, err := c.GetTokens(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(account.Docs, coinIndex), nil
}

// NormalizeToken converts a Ethereum token into the generic model
func NormalizeToken(srcToken *Contract, coinIndex uint) types.Token {
	tokenType := types.GetEthereumTokenTypeByIndex(coinIndex)

	return types.Token{
		Name:     srcToken.Name,
		Symbol:   srcToken.Symbol,
		TokenID:  srcToken.Address,
		Coin:     coinIndex,
		Decimals: srcToken.Decimals,
		Type:     tokenType,
	}
}

// NormalizeTxs converts multiple Ethereum tokens
func NormalizeTokens(srcTokens []Contract, coinIndex uint) []types.Token {
	assets := make([]types.Token, 0)
	for _, srcToken := range srcTokens {
		token := NormalizeToken(&srcToken, coinIndex)
		assets = append(assets, token)
	}
	return assets
}
