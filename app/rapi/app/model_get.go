package app

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/vitalvas/gomarkov/libmarkov"
)

// GetModel godoc
// @Summary Generate data from learned source
// @Accept json
// @Produce json
// @Success 200 {object} app.Message
// @Router /api/model [get]
func (m *Model) GetModel(c echo.Context) error {
	order := m.chain.Order
	tokens := make([]string, 0)

	for i := 0; i < order; i++ {
		tokens = append(tokens, libmarkov.StartToken)
	}

	for tokens[len(tokens)-1] != libmarkov.EndToken {
		next, _ := m.chain.Generate(tokens[(len(tokens) - order):])
		tokens = append(tokens, next)
	}

	msg := strings.Join(tokens[order:len(tokens)-1], "")

	return c.JSON(http.StatusOK, Message{
		Message: msg,
	})
}
