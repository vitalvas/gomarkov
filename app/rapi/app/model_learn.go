package app

import (
	"net/http"

	"github.com/labstack/echo"
)

// PostModelLearnBatch godoc
// @Summary Learn
// @Accept json
// @Produce json
// @Param message body app.BatchMessage true "Data for learn"
// @Success 201 {object} app.Message
// @Router /api/model/learn/batch [post]
func (m *Model) PostModelLearnBatch(c echo.Context) error {
	u := new(BatchMessage)
	if err := c.Bind(u); err != nil {
		return err
	}

	if len(u.Messages) > 0 {
		for _, row := range u.Messages {
			m.chain.RawAdd(row)
		}
	}

	return c.JSON(http.StatusCreated, Message{Message: "OK"})
}
