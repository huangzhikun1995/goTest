package GoTest

import (
	"go.uber.org/zap"
)

func main() {
	log := zap.NewExample()
	log.Debug("Hello World")
	log.Info("Hello World")
}

