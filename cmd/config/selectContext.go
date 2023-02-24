package config

import (
	"sort"

	"github.com/deviceinsight/kafkactl/output"
	"github.com/pkg/errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newSelectContextCmd() *cobra.Command {

	var cmdSelectContext = &cobra.Command{
		Use:     "select-context",
		Aliases: []string{"select"},
		Short:   "interactive select context",
		Long:    `interactive command for selection context`,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {

			contexts := viper.GetStringMap("contexts")

			contextsList := make([]string, 0)

			for k := range contexts {
				contextsList = append(contextsList, k)
			}

			sort.Strings(contextsList)

			var context string

			prompt := &survey.Select{
				Message: "Choose a context:",
				Options: contextsList,
			}

			err := survey.AskOne(prompt, &context)
			if err != nil {
				output.Fail(errors.Wrap(err, "unable to select cluster context"))
			}

			viper.Set("current-context", context)

			if err := viper.WriteConfig(); err != nil {
				output.Fail(errors.Wrap(err, "unable to write config"))
			}
		},
	}

	return cmdSelectContext
}
