package tel

import (
	"fmt"
	"net/http"
	"os"
)

func Tel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "123456")
	fmt.Fprintf(w, "789")
	os.Setenv("test", "1234")

	a := os.Getenv("test")
	fmt.Fprintf(w, a)

}
