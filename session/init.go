package session

// The logger for the session.
import "github.com/Laur1nMartins/revel/logger"

var sessionLog logger.MultiLogger

func InitSession(coreLogger logger.MultiLogger) {
	sessionLog = coreLogger.New("section", "session")
}
