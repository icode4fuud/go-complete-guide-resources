// Use in handlers/services instead of log.Println
// You can later swap this to zap or zerolog without touching call sites.
package logging

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "[api] ", log.LstdFlags|log.Lshortfile)
