package app

import (
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/vitalvas/gomarkov/libmarkov"
)

// GetModelDump godoc
// @Summary Dump current model
// @Accept json
// @Produce json
// @Success 200 {object} libmarkov.ChainJSON
// @Router /api/model/dump [get]
func (m *Model) GetModelDump(c echo.Context) error {
	return c.JSON(http.StatusOK, m.chain)
}
