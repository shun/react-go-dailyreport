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

type Dailyreport struct {
	Id         int       `gorm:"primary_key;AUTO_INCREMENT"`
	Usercode   string    `json:"code"`
	Comment    string    `json:"comment" gorm:"type:text"`
	Tasks      []Task    `json:"tasks" gorm:"foreignkey:ReportId;association_foreignkey:Id"`
	Created_at time.Time `sql:"DEFAULT:current_timestamp"`
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

type SystemInfo struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}

//=============================================================================
// Function
//
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
	report := new(Dailyreport)
	usercode := c.Param("usercode")
	if err := c.Bind(report); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	if usercode != report.Usercode {
		fmt.Printf("%s, %s\n", usercode, report.Usercode)
		return c.String(http.StatusBadRequest, "NG")
	}

	report.Id = 0
	db := gormConnect()

	db.Create(report)
	return c.JSON(http.StatusOK, report)
}

func readReport(c echo.Context) error {
	db := gormConnect()
	defer db.Close()

	report := Dailyreport{}
	id := c.Param("id")

	query := db.Preload("Tasks")
	if 0 != len(id) {
		query = query.Where("id = ?", id)
	}

	query.First(&report)

	return c.JSON(http.StatusOK, report)
}

func readReports(c echo.Context) error {
	db := gormConnect()
	defer db.Close()

	reports := []Dailyreport{}
	usercode := c.Param("usercode")

	query := db.Preload("Tasks")
	if 0 != len(usercode) {
		query = query.Where("usercode = ?", usercode)
	}

	query.Find(&reports)

	return c.JSON(http.StatusOK, reports)
}

func searchReport(c echo.Context) error {
	searchOption := new(SearchOption)
	if err := c.Bind(searchOption); err != nil {
		return err
	}

	fmt.Printf("%v\n", searchOption)
	reports := []Dailyreport{}

	db := gormConnect()
	defer db.Close()
	db.Preload("Tasks").Where("usercode = ?", searchOption.Usercode).Find(&reports)
	return c.JSON(http.StatusOK, reports)
}

func readSoftwareInfo(c echo.Context) error {
	info := SystemInfo{"0.0.1", "Dailyreport"}
	return c.JSON(http.StatusOK, info)
}

func updateReport(c echo.Context) error {
	db := gormConnect()
	defer db.Close()

	var report Dailyreport
	tmpreport := new(Dailyreport)

	id := c.Param("id")
	if err := c.Bind(tmpreport); err != nil {
		fmt.Printf("[ERROR]: %v\n", err)
		return err
	}

	db.Preload("Tasks").Where("id = ?", id).Find(&report)
	if report.Comment != tmpreport.Comment {
		db.Model(&report).Update("Comment", tmpreport.Comment)
	}

	if len(report.Tasks) != len(tmpreport.Tasks) {
		db.Delete(report.Tasks)
		report.Tasks = tmpreport.Tasks
		db.Save(&report)

	} else {
		report.Tasks = tmpreport.Tasks
		db.Save(&report)
	}
	return c.JSON(http.StatusOK, report)
}

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins:     []string{"http:///192.168.1.161:18080/"},
	//	AllowHeaders:     []string{"authorization", "Content-Type", "Access-Control-Allow-Origin"},
	//	AllowCredentials: true,
	//	AllowMethods:     []string{http.MethodGet, http.MethodPost},
	//}))

	e.GET("/info", readSoftwareInfo)
	e.POST("/users/:usercode/reports/registry", createReport)
	e.GET("/reports/:id", readReport)
	e.GET("/users/:usercode/reports", readReports)
	e.POST("/reports/search", searchReport)
	e.PUT("/reports/:id", updateReport)

	return e
}

func main() {
	db := gormConnect()
	db.AutoMigrate(&User{}, &Dailyreport{}, &Category{}, &Task{})
	defer db.Close()

	router := NewRouter()
	router.Logger.Fatal(router.Start(":3000"))
}
