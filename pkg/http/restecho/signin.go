package restecho

import (
	"net/http"

	"github.com/kiriwill/events-api/pkg/repository/mysql"
	"github.com/kiriwill/events-api/pkg/service"
	"github.com/labstack/echo/v4"
)

type UserCredentials struct {
	Email string `json:"email" validate:"required"`
	Pass  string `json:"password" validate:"required"`
}

func SignInHandler(svc *mysql.MysqlRepository, auth *auth) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(UserCredentials)
		var user service.User

		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		if err := c.Validate(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		rows := svc.Db.Where("email like ?", req.Email).Find(&user)
		if rows.Error != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		userExist, err := ComparePassword(user.Password, req.Pass)
		if err != nil {
			return err
		}

		if userExist {
			claims := auth.NewUserClaim()
			token, err := auth.GenerateToken(claims)
			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, echo.Map{
				"token": token,
			})
		}

		return c.NoContent(http.StatusBadRequest)
	}
}
