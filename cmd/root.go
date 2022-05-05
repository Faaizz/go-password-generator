package cmd

import (
	"fmt"
	"os"

	"github.com/faaizz/go-password-generator/business"
	"github.com/faaizz/go-password-generator/router"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "startServer",
	Short: "Start password generator server",
	Long: `
	Start password generator server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return router.Start()
	},
}

var generatePwdCmd = &cobra.Command{
	Use:   "generatePassword",
	Short: "Generate password",
	Long: `
	Generate password`,
	RunE: func(cmd *cobra.Command, args []string) error {
		minLen, err := cmd.Flags().GetInt("min_len")
		if err != nil {
			panic(err)
		}
		charsLen, err := cmd.Flags().GetInt("chars_len")
		if err != nil {
			panic(err)
		}
		numLen, err := cmd.Flags().GetInt("num_len")
		if err != nil {
			panic(err)
		}
		numPwds, err := cmd.Flags().GetInt("num_pwds")
		if err != nil {
			panic(err)
		}

		pwds, err := business.GetPwds(minLen, charsLen, numLen, numPwds)
		if err != nil {
			panic(err)
		}

		for _, pwd := range pwds {
			fmt.Println(pwd)
		}

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().IntP("min_len", "m", 8, "minimum number of characters")
	rootCmd.PersistentFlags().IntP("chars_len", "c", 2, "number of special characters")
	rootCmd.PersistentFlags().IntP("num_len", "n", 2, "number of numeric characters")
	rootCmd.PersistentFlags().IntP("num_pwds", "p", 1, "number of passwords to generate")

	rootCmd.AddCommand(generatePwdCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
