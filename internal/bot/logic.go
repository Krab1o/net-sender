package bot

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

var sizes = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}

func formatFileSize(s float64) string {
	base := 1024.0
	unitsLimit := len(sizes)
	i := 0
	for s >= base && i < unitsLimit {
		s = s / base
		i++
	}

	f := "%.0f %s"
	if i > 1 {
		f = "%.2f %s"
	}

	return fmt.Sprintf(f, s, sizes[i])
}

func parseLogin(msg string) (string, bool) {
	expr, err := regexp.Compile(`\/set_login\s[a-zA-Z0-9-_]{2,}`)
	if (err != nil) {
		log.Println(err)
	}
	
	if index := expr.FindStringIndex(msg); index != nil {
		//We take second argument right after "command"
		login := strings.Split(msg[index[0]:index[1]], " ")[1]
		return login, true
	}
	return "", false
}