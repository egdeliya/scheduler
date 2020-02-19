package startup

type Config struct {
	Port string `env:"PORT" envDefault:"8088"`
}
