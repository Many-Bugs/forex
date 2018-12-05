package starter

import (
	"fmt"
	"forex/systems"
	"forex/utils"
	"time"

	"github.com/jinzhu/gorm"
)

type Mysql struct {
	SQLInstance
	Username                       string
	Password                       string
	Host                           string
	Port                           int
	DatabaseName                   string
	IsWebService                   bool
	Error                          error
	MaximumIdleConnection          int
	MaximumOpenConnection          int
	MaximumConnectionKeepAliveTime int
}

type SQLInstance struct {
	DB         *gorm.DB
	BasicModel BasicModel
	ModelAddrs []interface{}
}

type BasicModel struct {
	ID          int `gorm:"primary_key" json:"id"`
	CreatedTime int `json:"created_time"`
	EditedTime  int `json:"edited_time"`
	DeletedTime int `json:"deleted_time"`
}

func (m *Mysql) Builder(c *Content) error {
	defer m.Connection()()
	m.connectionsetting()
	m.setCreateCallback()
	m.setUpdateCallback()
	m.setDeleteCallback()
	return nil
}

func (m *Mysql) connectionsetting() {
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
					"UPDATE %v SET %v=%v%v%v",
					scope.QuotedTableName(),
					scope.Quote(deletedOnField.DBName),
					scope.AddToVars(time.Now().Unix()),
					addSpace(scope.CombinedConditionSql()),
					addSpace(extraOption),
				)).Exec()
			} else {
				scope.Raw(fmt.Sprintf(
					"DELETE FROM %v%v%v",
					scope.QuotedTableName(),
					addSpace(scope.CombinedConditionSql()),
					addSpace(extraOption),
				)).Exec()
			}
		}
	})
	return
}

func (m *Mysql) Connection() func() error {
	var err error
	m.DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.DatabaseName))
	if !utils.AssertErr(err) {
		return m.DB.Close
	}
	return nil
}

func addSpace(value string) string {
	if value == "" {
		return value
	} else {
		return " " + value
	}
}

func (m *Mysql) AutoMigrateAddr(obj interface{}) {
	m.ModelAddrs = append(m.ModelAddrs, obj)

	return
}

func (m *Mysql) New() {
	m.ModelAddrs = make([]interface{}, 0)
	return
}

func (m *Mysql) createDatabase() {

}
