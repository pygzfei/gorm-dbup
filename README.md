### GORM DBup

## Install

```
go get -u github.com/pygzfei/gorm-dbup
```

## Quick start
```
database, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN:                       "DNS",
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}))

	if err != nil {
		t.Error(err)
	}
	                                    //database, sql files dir
	err = database.Use(pkg.NewMigration("test", "./dbup"))

	if err != nil {
		t.Error(err)
	}
```