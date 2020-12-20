package middlewares

import (
	"github.com/azukiazusa1/youtube-manager-go/databases"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type DatabaseClient struct {
	DB *gorm.DB
}

func DatabaseService() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := databases.Connect()
			d := DatabaseClient{DB: session}

			defer d.DB.Close()

			d.DB.LogMode(true)

			s.Set("dbs", &d)

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}