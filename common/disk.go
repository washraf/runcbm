package common

import (
	"os/exec"
	"strconv"
	"strings"
)

//FindDiskSizeMB
func FindDiskSizeMB(directory string) (int, error) {
	com := exec.Command("du", "-sm", directory)
	res, err := com.CombinedOutput()
	if err != nil {
		return 0, err
	}

	s := strings.Split(string(res), "\t")
	return strconv.Atoi(s[0])
}
