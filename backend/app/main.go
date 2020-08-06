package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
	"time"
)

type User struct {
	Id         int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Code       string
	Name       string
	Password   string
	Dept       string
	Mail       string
	Created_at *time.Time
	Updated_at *time.Time
}

type Category struct {
	Id     int `gorm:"primary_key;AUTO_INCREMENT"`
	Level1 int
	Level2 int
	Level3 int
	Name   string
}

type Task struct {
	Id         int `gorm:"primary_key;AUTO_INCREMENT"`
	ReportId   int
	CategoryId int     `json:"category"`
	Task       string  `json:"task"`
	Estimate   float32 `json:"estimate"`
}

type Dailylreport struct {
	Id         int    `gorm:"primary_key;AUTO_INCREMENT"`
	Usercode   string `json:"code"`
	Comment    string `json:"comment" gorm:"type:text"`
	Tasks      []Task `json:"tasks" gorm:"foreignkey:ReportId"`
	Created_at *time.Time
	Updated_at *time.Time
}

type Dbconfig struct {
	user string
	pass string
	name string
	host string
	port string
}

type SearchOption struct {
	Usercode      string     `json:"usercode"`
	Username      string     `json:"username"`
	FromCreatedAt *time.Time `json:"from_created_at"`
	ToCreatedAt   *time.Time `json:"to_created_at"`
}

func gormConnect() *gorm.DB {
	dbconfig := new(Dbconfig)
	dbconfig.user = os.Getenv("DB_USER")
	dbconfig.pass = os.Getenv("DB_PASS")
	dbconfig.name = os.Getenv("DB_NAME")
	dbconfig.host = os.Getenv("DB_HOST")
	dbconfig.port = os.Getenv("DB_PORT")

	dbconfig.name = os.Getenv("DB_NAME")
	connParam := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbconfig.user,
		dbconfig.pass,
		dbconfig.host,
		dbconfig.port,
		dbconfig.name,
	)

	db, err := gorm.Open(
		"mysql",
		connParam,
	)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func createReport(c echo.Context) error {
	fmt.Printf("createReport")
	report := new(Dailylreport)
	if err := c.Bind(report); err != nil {
		return err
	}

	report.Id = 0
	fmt.Printf("%v\n", report)
	db := gormConnect()
	defer db.Close()

	db.Create(report)
	return c.JSON(http.StatusOK, report)
}

func readReport(c echo.Context) error {
	db := gormConnect()
	defer db.Close()

	report := Dailylreport{}
	id := c.Param("id")
	db.Where("id = ?", id).First(&report)
	// FIXME why "Related" does't work ?
	db.Where("report_id = ?", report.Id).Find(&report.Tasks)

	return c.JSON(http.StatusOK, report)
}

func searchReport(c echo.Context) error {
	searchOption := new(SearchOption)
	if err := c.Bind(searchOption); err != nil {
		return err
	}

	fmt.Printf("%v\n", searchOption)
	reports := []Dailylreport{}

	db := gormConnect()
	defer db.Close()
	db.Where("usercode = ?", searchOption.Usercode).Find(&reports)
	return c.JSON(http.StatusOK, reports)
}

func main() {
	db := gormConnect()
	db.AutoMigrate(&User{}, &Dailylreport{}, &Category{}, &Task{})
	defer db.Close()

	e := echo.New()
	e.Use(middleware.CORS())
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins:     []string{"http:///192.168.1.161:18080/"},
	//	AllowHeaders:     []string{"authorization", "Content-Type", "Access-Control-Allow-Origin"},
	//	AllowCredentials: true,
	//	AllowMethods:     []string{http.MethodGet, http.MethodPost},
	//}))

	e.POST("/users/:id/reports/registry", createReport)

	e.GET("/reports/:id", readReport)

	e.POST("/reports/search", searchReport)

	e.Logger.Fatal(e.Start(":3000"))
}
