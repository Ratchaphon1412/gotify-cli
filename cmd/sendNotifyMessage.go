/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
)

// sendNotifyMessageCmd represents the sendNotifyMessage command
var sendNotifyMessageCmd = &cobra.Command{
	Use:   "sendNotifyMessage",
	Short: "Send a message to a gotify server.",
	Long: `Send a message to a gotify server.
	It is a simple tool that allows you to send messages to a gotify server using the command line.
	
	For example:
	
	$ gotify-cli sendNotifyMessage "Hello, World!" host:port
	
	This will send the message "Hello, World!" to the gotify server running at host:port.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			// text
			text := args[0]

			fmt.Println("Text is:", text)

			//url
			url := args[1]
			// send socket tcp to 0.0.0.0 8998
			conn, err := net.Dial("tcp", url)
			if err != nil {
				fmt.Println("Error connecting:", err.Error())
				os.Exit(1)
			}

			// send text content to server
			_, err = conn.Write([]byte(text))

			if err != nil {
				log.Println("Error sending message:", err)
				return
			}

			fmt.Println("Text  sent, waiting for response...")

			// receive response from server
			response := make([]byte, 1024)
			_, err = conn.Read(response)
			if err != nil {
				log.Println("Error reading response:", err)
				return
			}

			fmt.Println("Response:", string(response))
			defer conn.Close()

		} else {
			fmt.Println("No file path provided.")
		}
	},
}

func init() {
	rootCmd.AddCommand(sendNotifyMessageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendNotifyMessageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendNotifyMessageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
