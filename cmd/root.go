/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bjjsre/amhri/calculate"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "amhri",
	Short: "Tool to calculate heart rate intensities from the Advanced Marathoning book ",
	Long: `Tool to calculate heart rate intensities from the Advanced Marathoning book
	`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: doCalculate,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doCalculate(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("both resting and max HR required")
		return
	}

	resting, err := strconv.ParseFloat(args[0], 32)
	if err != nil {
		fmt.Printf("Failed to convert first arg to float: %w", err)
		return
	}

	max, err := strconv.ParseFloat(args[1], 32)
	if err != nil {
		fmt.Printf("Failed to convert first arg to float: %w", err)
		return
	}

	fmt.Printf("resting: %f max: %f\n", resting, max)
	calculate.Calculate(resting, max)
}
