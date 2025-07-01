package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/spf13/cobra"
)

func Decode[T any](r io.Reader) (T, error) {
	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

func RegisterStaticCompletions(cmd *cobra.Command, flag string, options []string) {
	err := cmd.RegisterFlagCompletionFunc(flag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return options, cobra.ShellCompDirectiveNoFileComp
	})

	if err != nil {
		log.Printf("warning: could not register completion for --%s: %v", flag, err)
	}
}
