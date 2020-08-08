package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func setup() (*gorm.DB, func()) {
	db := gormConnect()
	db.AutoMigrate(&User{}, &Dailyreport{}, &Category{}, &Task{})

	return db, func() {
		// teardown
		// remote all rows
		db.Delete(User{})
		db.Delete(Dailyreport{})
		db.Delete(Category{})
		db.Delete(Task{})
		db.Close()
	}
}

func getNow() time.Time {
	now := time.Now()
	dt := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
		0,
		time.Local)
	return dt
}

func TestHandler_SystemInfo(t *testing.T) {

	var actual SystemInfo
	expected := SystemInfo{"0.0.1", "Dailyreport"}

	req := httptest.NewRequest("GET", "/info", nil)
	rec := httptest.NewRecorder()

	router := NewRouter()
	router.ServeHTTP(rec, req)
	err := json.NewDecoder(rec.Body).Decode(&actual)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expected, actual)
}

func TestHandler_作業項目無しの日報を日報idで取得(t *testing.T) {

	db, teardown := setup()
	defer teardown()

	dt := getNow()
	expected := Dailyreport{1, "1234567", "こめんと", []Task{}, dt, nil}
	dummy1 := Dailyreport{2, "21234567", "こめんと2", []Task{}, dt, nil}
	dummy2 := Dailyreport{3, "31234567", "こめんと3", []Task{}, dt, nil}
	db.Create(expected)
	db.Create(dummy1)
	db.Create(dummy2)

	router := NewRouter()
	req := httptest.NewRequest("GET", "/reports/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	json_expected, _ := json.Marshal(expected)
	assert.Equal(t, http.StatusOK, rec.Code)
	json_actual := []byte(strings.TrimRight(rec.Body.String(), "\n"))
	assert.Equal(t, json_expected, json_actual)
}

func TestHandler_作業項目ありの日報を日報idで取得(t *testing.T) {
	db, teardown := setup()
	defer teardown()

	dt := getNow()
	expected := Dailyreport{
		1,
		"1234567",
		"こめんと",
		[]Task{
			Task{1, 1, 1, "たすく", 1.5},
			Task{2, 1, 1, "たすく2", 2.5},
		},
		dt,
		nil,
	}
	dummy1 := Dailyreport{
		2,
		"1234567",
		"こめんと2",
		[]Task{
			Task{3, 2, 1, "たすく3", 3.0},
			Task{4, 2, 1, "たすく4", 4.5},
		},
		dt,
		nil,
	}
	dummy2 := Dailyreport{
		3,
		"31234567",
		"こめんと3",
		[]Task{
			Task{5, 3, 1, "たすく5", 5.5},
			Task{6, 3, 1, "たすく6", 6.0},
		},
		dt,
		nil,
	}
	db.Create(expected)
	db.Create(dummy1)
	db.Create(dummy2)

	router := NewRouter()
	req := httptest.NewRequest("GET", "/reports/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	json_expected, _ := json.Marshal(expected)
	assert.Equal(t, http.StatusOK, rec.Code)
	json_actual := []byte(strings.TrimRight(rec.Body.String(), "\n"))
	assert.Equal(t, json_expected, json_actual)
}

func TestHandler_氏名コード指定で全件取得(t *testing.T) {
	db, teardown := setup()
	defer teardown()

	dt := getNow()
	var expected []Dailyreport
	report1 := Dailyreport{
		1,
		"1234567",
		"こめんと",
		[]Task{
			Task{1, 1, 1, "たすく", 1.5},
			Task{2, 1, 1, "たすく2", 2.5},
		},
		dt,
		nil,
	}
	report2 := Dailyreport{
		2,
		"1234567",
		"こめんと2",
		[]Task{
			Task{3, 2, 1, "たすく3", 3.0},
			Task{4, 2, 1, "たすく4", 4.5},
		},
		dt,
		nil,
	}
	dummy1 := Dailyreport{
		3,
		"31234567",
		"こめんと3",
		[]Task{
			Task{5, 3, 1, "たすく5", 5.5},
			Task{6, 3, 1, "たすく6", 6.0},
		},
		dt,
		nil,
	}
	db.Create(report1)
	db.Create(report2)
	db.Create(dummy1)
	expected = append(expected, report1)
	expected = append(expected, report2)

	//var actual []Dailyreport
	router := NewRouter()
	req := httptest.NewRequest("GET", "/users/1234567/reports", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	json_expected, _ := json.Marshal(expected)
	json_actual := []byte(strings.TrimRight(rec.Body.String(), "\n"))
	assert.Equal(t, json_expected, json_actual)
}

func TestHandler_１タスクの日報を１件登録する(t *testing.T) {
	db, teardown := setup()
	defer teardown()

	report := []byte(`{
		"code": "1234567",
		"comment": "てすとこめんと",
		"tasks": [
			{"category": 1, "task": "テストたすく", "estimate": 1.5}
		]
	}`)

	router := NewRouter()
	req := httptest.NewRequest("POST", `/users/1234567/reports/registry`, strings.NewReader(string(report)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	var actual Dailyreport
	var expected Dailyreport
	err := json.Unmarshal(report, &expected)
	if err != nil {
		fmt.Println(err)
	}
	db.Preload("Tasks").Where("usercode = 1234567").First(&actual)

	assert.Equal(t, expected.Usercode, actual.Usercode)
	assert.Equal(t, expected.Comment, actual.Comment)
	assert.Equal(t, expected.Tasks[0].CategoryId, actual.Tasks[0].CategoryId)
	assert.Equal(t, expected.Tasks[0].Task, actual.Tasks[0].Task)
	assert.Equal(t, expected.Tasks[0].Estimate, actual.Tasks[0].Estimate)
}

func TestHandler_３タスクの日報を１件登録する(t *testing.T) {
	db, teardown := setup()
	defer teardown()

	report := []byte(`{
		"code": "1234567",
		"comment": "てすとこめんと",
		"tasks": [
			{"category": 1, "task": "テストたすく1", "estimate": 1.5},
			{"category": 2, "task": "テストたすく2", "estimate": 1.0},
			{"category": 3, "task": "テストたすく3", "estimate": 1.3}
		]
	}`)

	router := NewRouter()
	req := httptest.NewRequest("POST", `/users/1234567/reports/registry`, strings.NewReader(string(report)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	var actual Dailyreport
	var expected Dailyreport
	err := json.Unmarshal(report, &expected)
	if err != nil {
		fmt.Println(err)
	}
	db.Preload("Tasks").Where("usercode = 1234567").First(&actual)

	assert.Equal(t, expected.Usercode, actual.Usercode)
	assert.Equal(t, expected.Comment, actual.Comment)
	assert.Equal(t, len(expected.Tasks), len(actual.Tasks))
	for i := range expected.Tasks {
		assert.Equal(t, expected.Tasks[i].CategoryId, actual.Tasks[i].CategoryId)
		assert.Equal(t, expected.Tasks[i].Task, actual.Tasks[i].Task)
		assert.Equal(t, expected.Tasks[i].Estimate, actual.Tasks[i].Estimate)
	}
}

func TestHandler_タスク数を変えずに日報の内容を更新する(t *testing.T) {
	db, teardown := setup()
	defer teardown()

	dt := getNow()
	report1 := Dailyreport{
		1,
		"1234567",
		"こめんと",
		[]Task{
			Task{1, 1, 1, "たすく", 1.5},
			Task{2, 1, 1, "たすく2", 2.5},
		},
		dt,
		nil,
	}
	db.Create(report1)

	report1.Comment = "変更後のこめんと"
	report1.Tasks[0].CategoryId = 2
	report1.Tasks[0].Task = "たすくかえた"
	report1.Tasks[1].Estimate = 4.0
	json_str, _ := json.Marshal(report1)
	router := NewRouter()
	req := httptest.NewRequest("PUT", "/reports/1", strings.NewReader(string(json_str)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	expected := report1
	var actual Dailyreport
	db.Preload("Tasks").Where("usercode = 1234567").First(&actual)
	assert.Equal(t, expected.Usercode, actual.Usercode)
	assert.Equal(t, expected.Comment, actual.Comment)
	assert.Equal(t, len(expected.Tasks), len(actual.Tasks))
	for i := range expected.Tasks {
		assert.Equal(t, expected.Tasks[i].CategoryId, actual.Tasks[i].CategoryId)
		assert.Equal(t, expected.Tasks[i].Task, actual.Tasks[i].Task)
		assert.Equal(t, expected.Tasks[i].Estimate, actual.Tasks[i].Estimate)
	}
}

func TestHandler_タスク数を１件削除する(t *testing.T) {
	db, teardown := setup()
	defer teardown()

	dt := getNow()
	report1 := Dailyreport{
		1,
		"1234567",
		"こめんと",
		[]Task{
			Task{1, 1, 1, "たすく", 1.5},
			Task{2, 1, 1, "たすく2", 2.5},
		},
		dt,
		nil,
	}
	db.Create(report1)

	report1.Tasks = []Task{Task{1, 1, 1, "たすく", 1.5}}
	json_str, _ := json.Marshal(report1)
	router := NewRouter()
	req := httptest.NewRequest("PUT", "/reports/1", strings.NewReader(string(json_str)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	expected := report1
	var actual Dailyreport
	db.Preload("Tasks").Where("usercode = 1234567").First(&actual)
	assert.Equal(t, expected.Usercode, actual.Usercode)
	assert.Equal(t, expected.Comment, actual.Comment)
	assert.Equal(t, len(expected.Tasks), len(actual.Tasks))
	for i := range expected.Tasks {
		assert.Equal(t, expected.Tasks[i].CategoryId, actual.Tasks[i].CategoryId)
		assert.Equal(t, expected.Tasks[i].Task, actual.Tasks[i].Task)
		assert.Equal(t, expected.Tasks[i].Estimate, actual.Tasks[i].Estimate)
	}
}

func TestHandler_タスク数を１件追加する(t *testing.T) {
	db, teardown := setup()
	defer teardown()

	dt := getNow()
	report1 := Dailyreport{
		1,
		"1234567",
		"こめんと",
		[]Task{
			Task{1, 1, 1, "たすく", 1.5},
			Task{2, 1, 1, "たすく2", 2.5},
		},
		dt,
		nil,
	}
	db.Create(report1)

	report1.Tasks = []Task{
		Task{1, 1, 1, "たすく", 1.5},
		Task{2, 1, 1, "たすく2", 2.5},
		Task{3, 1, 1, "たすく3", 5.5},
	}
	json_str, _ := json.Marshal(report1)
	router := NewRouter()
	req := httptest.NewRequest("PUT", "/reports/1", strings.NewReader(string(json_str)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	expected := report1
	var actual Dailyreport
	db.Preload("Tasks").Where("usercode = 1234567").First(&actual)
	assert.Equal(t, expected.Usercode, actual.Usercode)
	assert.Equal(t, expected.Comment, actual.Comment)
	assert.Equal(t, len(expected.Tasks), len(actual.Tasks))
	for i := range expected.Tasks {
		assert.Equal(t, expected.Tasks[i].CategoryId, actual.Tasks[i].CategoryId)
		assert.Equal(t, expected.Tasks[i].Task, actual.Tasks[i].Task)
		assert.Equal(t, expected.Tasks[i].Estimate, actual.Tasks[i].Estimate)
	}
}
