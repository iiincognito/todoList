package cmd

import (
	"GoLess2/httpp"
	"GoLess2/todo"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var port int

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {

		addr := fmt.Sprintf(":%d", port)
		todoList := todo.NewList()
		httpHand := httpp.NewHTTPHandlers(todoList)
		server := httpp.NewHTTPServer(httpHand)
		log.Printf("Server is running on http://localhost%s\n", addr)
		err := server.StartServer(addr)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntVar(
		&port,
		"port",
		9091,
		"port to listen on",
	)
}
