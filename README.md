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