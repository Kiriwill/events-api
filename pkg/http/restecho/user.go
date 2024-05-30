package restecho

import (
	"net/http"

	"github.com/kiriwill/events-api/pkg/repository/mysql"
	"github.com/kiriwill/events-api/pkg/service"
	"github.com/labstack/echo/v4"
)

func PostUserHandler(svc *mysql.MysqlRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		req := new(service.User)
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		if err := c.Validate(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		req.Password, err = EncriptPassword(req.Password)
		if err != nil {
			return err
		}

		rows := svc.Db.Create(req)
		if rows.RowsAffected == 0 {
			return c.NoContent(http.StatusBadRequest)
		}

		if rows.Error != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusCreated, struct {
			ID uint
		}{
			ID: req.ID,
		})
	}
}
