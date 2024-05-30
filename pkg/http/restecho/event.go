package restecho

import (
	"net/http"

	"github.com/kiriwill/events-api/pkg/repository/mysql"
	"github.com/kiriwill/events-api/pkg/service"
	"github.com/labstack/echo/v4"
)

func SearchEventHandler(svc *mysql.MysqlRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		var events []service.Event

		date_param := c.QueryParam("date")
		type_param := c.QueryParam("type")

		if type_param != "" && date_param != "" {
			type_param := "%" + type_param + "%"
			date_param := "%" + date_param + "%"
			rows := svc.Db.Where("category like ? AND created_at like ?", type_param, date_param).Find(&events)
			if rows.Error != nil {
				return c.NoContent(http.StatusInternalServerError)
			}
		}

		if type_param != "" && date_param == "" {
			type_param := "%" + type_param + "%"
			rows := svc.Db.Where("category like ?", type_param).Find(&events)
			if rows.Error != nil {
				return c.NoContent(http.StatusInternalServerError)
			}
		}

		if date_param != "" && type_param == "" {
			date_param := "%" + date_param + "%"
			rows := svc.Db.Where("created_at like ? ", date_param).Find(&events)
			if rows.Error != nil {
				return c.NoContent(http.StatusInternalServerError)
			}
		}

		if date_param == "" && type_param == "" {
			rows := svc.Db.Find(&events)
			if rows.Error != nil {
				return c.NoContent(http.StatusInternalServerError)
			}
		}
		return c.JSON(http.StatusOK, events)
	}
}

func GetMetricsHandler(svc *mysql.MysqlRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		var metrics []struct {
			Country string
			Count   int
		}

		rows := svc.Db.Model(&service.Event{}).Select("country, count(country) as count").Group("country").Order("count DESC").Limit(3).Scan(&metrics)
		if rows.Error != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, metrics)
	}
}

func PostEventHandler(svc *mysql.MysqlRepository) echo.HandlerFunc {
	return func(c echo.Context) error {

		event := new(service.Event)
		if err := c.Bind(event); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		if err := c.Validate(event); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		err := svc.Db.Create(event).Error
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, event)
	}
}
