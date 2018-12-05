package starter

type Builder interface {
	Builder(*Content) error
}

type Starter interface {
	Starter(*Content) error
}

var (
	_ Builder = &Config{}
	_ Builder = &Logger{}
	_ Builder = &App{}
	_ Builder = &Server{}
	_ Builder = &Mysql{}
	_ Builder = &Mongo{}
	_ Builder = &Influx{}
	_ Builder = &Redis{}
	_ Builder = &Crawler{}
)
