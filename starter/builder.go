package starter

type Builder interface {
	Builder() error
}

var (
	_ Builder = &Config{}
	_ Builder = &Logger{}
	_ Builder = &App{}
	_ Builder = &Server{}
)
