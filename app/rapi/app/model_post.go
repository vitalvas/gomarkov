package app

import (
	"net/http"

	"github.com/labstack/echo"
)

// PostModel godoc
// @Summary Learn
// @Accept json
// @Produce json
// @Param message body app.Message true "Data for learn"
// @Success 200 {object} app.Message
// @Router /api/model/learn [post]
func (m *Model) PostModel(c echo.Context) error {
	u := new(Message)
	if err := c.Bind(u); err != nil {
		return err
	}

	if len(u.Message) > 0 {
		m.chain.RawAdd(u.Message)
		return c.JSON(http.StatusCreated, u)
	}

	return c.JSON(http.StatusNoContent, u)
}
