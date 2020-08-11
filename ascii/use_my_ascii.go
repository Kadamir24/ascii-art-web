package ascii

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func AsciiCall(text, font string) (string, int) {
	var output string
	if !check(text) {
		return "No Russian", http.StatusBadRequest
	}
	if font == "" {
		font = "standard"
	}
	if text == "" {
		return "Please enter text", http.StatusBadRequest
	}
	file, err := ioutil.ReadFile("templates/assets/fonts/" + font + ".txt")
	if err != nil {
		return "500 Internal Server Error", http.StatusInternalServerError
	}
	strFile := strings.Split(string(file), "\n")
	strText := strings.Split(text, "\\n")

	for _, str := range strText {
		for i := 1; i < 9; i++ {
			for _, ch := range str {
				output += (strFile[(int(ch)-32)*9+i])
			}
			output += "\n"
		}
	}
	return output, http.StatusOK
}

func check(str string) bool {
	var re = regexp.MustCompile(`[а-яА-Я]`)
	if len(re.FindAllString(str, -1)) != 0 {
		return false
	}
	return true
}
