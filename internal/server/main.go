package main
import (
"log"
"github.com/Saad-Hussain31/proglog/internal/server"
)
func main() {
srv := server.NewHTTPServer(":8080")
log.Fatal(srv.ListenAndServe())
}

