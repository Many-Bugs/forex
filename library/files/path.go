package files

import (
	"runtime"
	"strings"
)

func GetSplit() (split string) {
	systemType := runtime.GOOS
	split = "/"
	switch systemType {
	case "windows":
		split = "\\"
	case "linux":
		split = "/"
	}
	return
}

func ReplaceSplit(dir string) (ret string) {
	systemType := runtime.GOOS
	ret = dir
	switch systemType {
	case "windows":
		windowsReplaccer := strings.NewReplacer("/", "\\")
		ret = windowsReplaccer.Replace(dir)
	case "linux":
		linuxReplacer := strings.NewReplacer("\\", "/")
		ret = linuxReplacer.Replace(dir)
	}
	return
}

func ReplaceSplitToLinux(dir string) (ret string) {
	linuxReplacer := strings.NewReplacer("\\", "/")
	ret = linuxReplacer.Replace(dir)
	return
}

func ReplaceSplitToWindows(dir string) (ret string) {
	windowsReplaccer := strings.NewReplacer("/", "\\")
	ret = windowsReplaccer.Replace(dir)
	return
}
