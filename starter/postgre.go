package starter

import (
	"database/sql"
	"fmt"
	"forex/systems"
	"forex/utils"
	"strconv"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
)

type Postgres struct {
	PostgresInstance
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

type PostgresInstance struct {
	DB         *gorm.DB
	db         *sql.DB
	BasicModel PostgresModel
	ModelAddrs []interface{}
}

type PostgresModel struct {
	ID          int `gorm:"primary_key" json:"id"`
	CreatedTime int `json:"created_time"`
	EditedTime  int `json:"edited_time"`
	DeletedTime int `json:"deleted_time"`
}

func (m *Postgres) Builder(c *Content) error {
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

func (m *Postgres) CreateDatabase() {
	m.db, m.Error = sql.Open("postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s",
			m.Host,
			strconv.Itoa(m.Port),
			m.Username,
			m.Password),
	)
	stmt := fmt.Sprintf(
		"CREATE DATABASE "+
			"IF NOT EXISTS %s "+
			"WITH ENCODING='UTF8' "+
			"TABLESPACE = pg_default;",
		m.DatabaseName)
	_, _ = m.db.Exec(stmt)
	stmt = fmt.Sprintf(
		"ALTER DATABASE %s "+
			"WITH ENCODING='UTF8' "+
			"TABLESPACE = pg_default;",
		m.DatabaseName)
	_, _ = m.db.Exec(stmt)
}

func (m *Postgres) connectionSetting() {
	m.DB.DB().SetMaxIdleConns(m.MaximumIdleConnection)
	m.DB.DB().SetMaxOpenConns(m.MaximumOpenConnection)
	m.DB.DB().SetConnMaxLifetime(time.Duration(m.MaximumConnectionKeepAliveTime))
}

func (m *Postgres) setCreateCallback() {
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

func (m *Postgres) setUpdateCallback() {
	m.DB.Callback().Update().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		if _, ok := scope.Get("gorm:update_column"); !ok {
			scope.SetColumn("EditedTime", systems.NowInUNIX())
		}
	})
	return
}

func (m *Postgres) setDeleteCallback() {
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

func (m *Postgres) Connector() func() error {
	m.recursionCall(
		func() error {
			m.DB, m.Error = gorm.Open("postgres",
				fmt.Sprintf(
					"host=%s port=%s user=%s dbname=%s password=%s",
					m.Host,
					strconv.Itoa(m.Port),
					m.Username,
					m.DatabaseName,
					m.Password))
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

func (m *Postgres) recursionCall(f func() error, count, duration int, done bool) bool {
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
