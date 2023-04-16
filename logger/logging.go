package logger

import (
	"fmt"
	"os"
)

// TI writes text [INFO] for logging
func TI() string {
	return "\x1b[92m[INFO]\x1b[0m"
}

// TW writes text [WARNING] for logging
func TW() string {
	return "\x1b[33m[WARNING]\x1b[0m"
}

// TE writes text [ERROR] for logging
func TE() string {
	return "\x1b[91m[ERROR]\x1b[0m"
}

// DebugInfo represents debug info
func DebugInfo(a ...interface{}) {
	fmt.Fprintln(os.Stderr, append([]interface{}{TI()}, a...)...)
	//fmt.Println(TI(), a)
}

// DebugWarning represents debug warning
func DebugWarning(a ...interface{}) {
	fmt.Fprintln(os.Stderr, append([]interface{}{TW()}, a...)...)
}

// DebugError represents debug error
func DebugError(a ...interface{}) {
	fmt.Fprintln(os.Stderr, append([]interface{}{TE()}, a...)...)
	//fmt.Println(TE(), a)
}
