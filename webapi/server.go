package webapi

import (
	"fmt"
	"net/http"
)

func StartServer(port string) error {
	http.HandleFunc("/findSCC", handleFindSCC)
	http.HandleFunc("/findBridge", handleFindBridge)
	fmt.Println("Starting web server at http://localhost:" + port)
	return http.ListenAndServe(":" + port, nil)
}