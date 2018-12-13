package starter

type Starter interface {
	Starter(*Content) error
}

var (
	_ Starter = &Config{}
	_ Starter = &Logger{}
	_ Starter = &App{}
	_ Starter = &Server{}
	_ Starter = &Mysql{}
	_ Starter = &Mongo{}
	_ Starter = &Influx{}
	_ Starter = &Redis{}
)

func (m *Content) DefaultStarter() {

}

func (m *Content) Starter(s Starter) error {
	return nil
}
