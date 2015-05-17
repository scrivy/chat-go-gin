package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.ServeFiles("/*filepath", http.Dir("public"))
    router.Run(":8000")
}
