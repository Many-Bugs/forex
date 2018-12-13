package starter

type Router interface {
	Router(*Server)
}

var (
	_ Router = &Config{}
	_ Router = &Logger{}
	_ Router = &App{}
	_ Router = &Mysql{}
	_ Router = &Mongo{}
	_ Router = &Influx{}
)
