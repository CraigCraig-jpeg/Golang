/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// waiterCmd represents the waiter command
var waiterCmd = &cobra.Command{
	Use:   "waiter",
	Short: "Work out tip for a waiter",
	Long: `work out tip for a waiter`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("waiter called")
		Add(money)
	},
}

func init() {
	rootCmd.AddCommand(waiterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// waiterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// waiterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	waiterCmd.Flags().Int( &money, "Amount", "a", "", "amount to add to")
}

var money int

func Add(money2 int) {
	fmt.Println(money2)
}