package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/ssh"
	"net"
	"sync"
	"time"
	"wm-infoflow-api-go/conf"
)

// NewMySQLByDSN connects to MySQL by DSN
func NewMySQLByDSN(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mysql: %s", err)
	}

	db.DB().SetConnMaxLifetime(60 * time.Second)
	db.DB().SetMaxOpenConns(2000)
	db.DB().SetMaxIdleConns(100)


	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	return db, nil
}

// GenMySQLDSN generates DSN for MySQL
func GenMySQLDSN(cfg conf.MySQL) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&timeout=30s&parseTime=true",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}

// NewMySQL connects to MySQL with config struct
func NewMySQL(cfg conf.MySQL) (*gorm.DB, error) {
	return NewMySQLByDSN(GenMySQLDSN(cfg))
}

var mySQL = make(map[string]*gorm.DB)
var mySQLLock = sync.Mutex{}

// GetMySQL create MySQL connection instance
func GetMySQL(cfg conf.MySQL) (*gorm.DB, error) {
	dsn := GenMySQLDSN(cfg)
	if db, ok := mySQL[dsn]; ok {
		return db, nil
	}

	mySQLLock.Lock()
	defer mySQLLock.Unlock()
	if db, ok := mySQL[dsn]; ok {
		return db, nil
	}

	if db, err := NewMySQLByDSN(dsn); err == nil {
		mySQL[dsn] = db
		return mySQL[dsn], nil
	} else {
		return nil, err
	}
}

type SSHDialer struct {
	*ssh.Client
}

func (dialer *SSHDialer) Dial(addr string) (net.Conn, error) {
	return dialer.Client.Dial("tcp", addr)
}


func updateTimeStampForCreateCallback(scope *gorm.Scope) {

	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// 注册更新钩子在持久化之前
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", time.Now().Unix())
	}
}

// 注册删除钩子在删除之前
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
