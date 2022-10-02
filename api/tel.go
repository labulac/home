package tel

import (
	"fmt"
	"net/http"
)

func Tel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "123456")
}
