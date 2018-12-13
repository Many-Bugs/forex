package starter

import (
	"database/sql"
	"fmt"
	"forex/systems"
	"forex/utils"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

type Mysql struct {
	MysqlInstance
	Username                       string
	Password                       string
	Host                           string
	Port                           int
	DatabaseName                   string
	IsWebService                   bool
	Error                          error
	MaximumIdleConnection          int
	MaximumOpenConnection          int
	MaximumConnectionRetry         int
	MinimumRetryDuration           int
	MaximumConnectionKeepAliveTime int
}

type MysqlInstance struct {
	DB         *gorm.DB
	db         *sql.DB
	BasicModel MysqlModel
	ModelAddrs []interface{}
}

type MysqlModel struct {
	ID          int `gorm:"primary_key" json:"id"`
	CreatedTime int `json:"created_time"`
	EditedTime  int `json:"edited_time"`
	DeletedTime int `json:"deleted_time"`
}

func (m *Mysql) Builder(c *Content) error {
	m.CreateDatabase()
	if close := m.Connector(); close != nil {
		defer close()
		m.connectionSetting()
		m.setCreateCallback()
		m.setUpdateCallback()
		m.setDeleteCallback()
	}
	return nil
}

func (m *Mysql) CreateDatabase() {
	m.db, _ = sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/",
			m.Username,
			m.Password,
			m.Host,
			m.Port),
	)
	stmt := fmt.Sprintf(
		"CREATE DATABASE "+
			"IF NOT EXISTS %s "+
			"CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;",
		m.DatabaseName)
	_, _ = m.db.Exec(stmt)
	stmt = fmt.Sprintf(
		"ALTER DATABASE %s "+
			"CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;",
		m.DatabaseName)
	_, _ = m.db.Exec(stmt)
}

func (m *Mysql) connectionSetting() {
	m.DB.DB().SetMaxIdleConns(m.MaximumIdleConnection)
	m.DB.DB().SetMaxOpenConns(m.MaximumOpenConnection)
	m.DB.DB().SetConnMaxLifetime(time.Duration(m.MaximumConnectionKeepAliveTime))
}

func (m *Mysql) setCreateCallback() {
	m.DB.Callback().Create().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		if !scope.HasError() {
			now := systems.NowInUNIX()
			if createTimeField, ok := scope.FieldByName("CreatedTime"); ok {
				if createTimeField.IsBlank {
					createTimeField.Set(now)
				}
			}
			if modifyTimeField, ok := scope.FieldByName("EditedTime"); ok {
				if modifyTimeField.IsBlank {
					modifyTimeField.Set(now)
				}
			}
		}
	})
	return
}

func (m *Mysql) setUpdateCallback() {
	m.DB.Callback().Update().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		if _, ok := scope.Get("gorm:update_column"); !ok {
			scope.SetColumn("EditedTime", systems.NowInUNIX())
		}
	})
	return
}

func (m *Mysql) setDeleteCallback() {
	m.DB.Callback().Delete().Replace("gorm:delete", func(scope *gorm.Scope) {
		if !scope.HasError() {
			var extraOption string
			if str, ok := scope.Get("gorm:delete_option"); ok {
				extraOption = fmt.Sprint(str)
			}
			if deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedTime"); !scope.Search.Unscoped && hasDeletedOnField {
				scope.Raw(fmt.Sprintf(
					"UPDATE %v SET %v=%v %v %v",
					scope.QuotedTableName(),
					scope.Quote(deletedOnField.DBName),
					scope.AddToVars(time.Now().Unix()),
					scope.CombinedConditionSql(),
					extraOption,
				)).Exec()
			} else {
				scope.Raw(fmt.Sprintf(
					"DELETE FROM %v %v %v",
					scope.QuotedTableName(),
					scope.CombinedConditionSql(),
					extraOption,
				)).Exec()
			}
		}
	})
	return
}

func (m *Mysql) Connector() func() error {
	m.recursionCall(
		func() error {
			m.DB, m.Error = gorm.Open("mysql",
				fmt.Sprintf(
					"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
					m.Username,
					m.Password,
					m.Host+":"+strconv.Itoa(m.Port),
					m.DatabaseName))
			return m.Error
		},
		m.MaximumConnectionRetry,
		m.MinimumRetryDuration,
		false)
	if !utils.AssertErr(m.Error) {
		return m.DB.Close
	}
	return nil
}

func (m *Mysql) recursionCall(f func() error, count, duration int, done bool) bool {
	if !done {
		m.Error = f()
		count--
	}
	if count > 0 && m.Error == nil {
		return true
	} else if count == 0 && m.Error != nil {
		return true
	} else {
		time.Sleep(time.Duration(duration) * time.Second)
	}
	return m.recursionCall(f, count, duration, false)
}

func (m *Mysql) AutoMigrateByAddr(obj interface{}) {
	m.ModelAddrs = append(m.ModelAddrs, obj)
	for _, model := range m.ModelAddrs {
		defer m.Connector()()
		m.DB.AutoMigrate(model)
	}
	return
}

func (m *Mysql) New() {
	m.ModelAddrs = make([]interface{}, 0)
	return
}

func (m *Mysql) Starter(c *Content) error {
	return nil
}

func (m *Mysql) Router(s *Server) {
	return
}
