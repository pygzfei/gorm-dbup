[![Build Status](https://github.com/pygzfei/gorm-dbup/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/pygzfei/gorm-dbup/actions?query=branch%3Amaster)
[![doc](https://img.shields.io/badge/go.dev-doc-007d9c?style=flat-square&logo=read-the-docs)](https://pkg.go.dev/github.com/pygzfei/gorm-dbup)
[![codecov](https://codecov.io/gh/pygzfei/gorm-dbup/branch/main/graph/badge.svg?token=RSY21OLN2B)](https://codecov.io/gh/pygzfei/gorm-dbup)
[![Release](https://img.shields.io/github/v/release/pygzfei/gorm-dbup.svg?style=flat-square)](https://github.com/pygzfei/gorm-dbup/releases)
![](https://img.shields.io/badge/license-MIT-green)
### GORM DBup

## Install

```
go get -u github.com/pygzfei/gorm-dbup
```

## Quick start

```go
// init
database, err := gorm.Open(mysql.New(
    mysql.Config{
        DSN:                       "root:123456@tcp(localhost:3306)/test1?charset=utf8mb4&parseTime=True&loc=Local",
        DefaultStringSize:         256,
        DisableDatetimePrecision:  true,
        DontSupportRenameIndex:    true,
        DontSupportRenameColumn:   true,
        SkipInitializeWithVersion: false,
    }))

    if err != nil {
        t.Error(err)
    }
                                        // test1 must be created before running
    err = database.Use(pkg.NewMigration("test1", "./dbup")) //Database, SQLFilesDir
    
    if err != nil {
        t.Error(err)
    }
```