package main

import (
	"log"
	"fmt"
	"../proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	g := gin.Default()

	// Create get method for Add service
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		// Try to convert param string to Uint
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			// Return error message throught json
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		// Try to convert param string to Uint
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			// Return error message throught json
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
			return
		}

		// Create a request without valid params
		request := &proto.Request{A: int64(a), B: int64(b)}

		// Create a response
		if response, err := client.Add(ctx, request); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	// Create get method for Multiply service
	g.GET("/mult/:a/:b", func(ctx *gin.Context) {
		// Try to convert param string to Uint
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			// Return error message throught json
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		// Try to convert param string to Uint
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			// Return error message throught json
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
			return
		}

		// Create a request without valid params
		request := &proto.Request{A: int64(a), B: int64(b)}

		// Create a response
		if response, err := client.Multiply(ctx, request); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	// Check if client can not run on
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}