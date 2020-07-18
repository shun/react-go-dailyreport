# What this ?

This repository is sinmple Daily report system.

---

# Architecture

- frontend: react
- backend: golang(Echo)
- DB: Mariadb

---

# Function

## Common
- [ ] Authentification

---

## For User
- [x] Registory daily report
- [ ] Re-edit daily report

---

## For Administrator
- [ ] List up daily report
- [x] Browse daily report
- Search daily report
    - [x] Search by Usercode
    - [x] Search by Username
    - [ ] Search by from created date
    - [ ] Search by to created date
- [ ] Add comment

---

# API

| # | method | name | description |
|:-:|:-:|:-|:-|
|1|POST|createReport|register the report data|
|2|GET|readReport|read the report data|
---

# Class Diagram

```plantuml
@startuml

namespace Dailyreport-sys {
    namespace entity {
        Dailyreport "1"*-right-"*" Task
        Task "1"--"1" Category
        User "1" *-right- "*" Dailyreport

        class Category {
            Id: int
            Level1: int
            Level2: int
            Level3: int
            Name: string
        }

        class Dailyreport
        {
            Id: int
            Usercode : string
            Comment : string
            Tasks : []Task
            Created_at : Time
            Updated_at : Time
        }

        class Task
        {
            Id: int
            ReportId : int
            CategoryId : int
            Task : string
            Estimate : float32
        }

        class User {
            Id: int
            Code : string
            Name : string
            Password : string
            Dept: string
            Mail : string
            Created_at : Time
            Updated_at : Time

        }

    }

}

@enduml
```

<style>
h1 {
background-color: #ff000;
}
</style>