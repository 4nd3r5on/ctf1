package setup

const (
	APP_STAGE_PROD = iota
	APP_STAGE_TEST = iota
	APP_STAGE_DEV  = iota
)

// Used for logging
func stageToString(stage int) string {
	switch stage {
	case APP_STAGE_PROD:
		return "production"
	case APP_STAGE_TEST:
		return "testing"
	case APP_STAGE_DEV:
		return "development"
	default:
		return "unknown"
	}
}

type Config struct {
	httpPort    int
	httpAddress string
}

type App struct {
}

func NewApp(cfg Config) {

}
