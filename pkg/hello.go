package pkg

import (
	"fmt"
	"net/http"
)

//HelloWord e a mensagem que aparece
func HelloWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "chamou o hello word")
}
