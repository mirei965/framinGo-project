package framingo

import (
	"fmt"
	"regexp"
	"runtime"
	"time"
)

func (f *Framingo) LoadTime(start time.Time) {
	elapsed := time.Since(start)
	pc, _, _, _ := runtime.Caller(1)
	funcObj := runtime.FuncForPC(pc)
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	f.InfoLog.Println(fmt.Sprintf("Load Time: %s took %s", name, elapsed))
}
