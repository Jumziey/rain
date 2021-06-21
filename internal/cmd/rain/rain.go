package rain

import (
	"strings"

	"github.com/jumziey/rain/internal/config"
	"github.com/spf13/cobra"

	"github.com/jumziey/rain/internal/cmd"
	"github.com/jumziey/rain/internal/cmd/build"
	"github.com/jumziey/rain/internal/cmd/cat"
	consolecmd "github.com/jumziey/rain/internal/cmd/console"
	"github.com/jumziey/rain/internal/cmd/deploy"
	"github.com/jumziey/rain/internal/cmd/diff"
	"github.com/jumziey/rain/internal/cmd/fmt"
	"github.com/jumziey/rain/internal/cmd/info"
	"github.com/jumziey/rain/internal/cmd/logs"
	"github.com/jumziey/rain/internal/cmd/ls"
	"github.com/jumziey/rain/internal/cmd/merge"
	"github.com/jumziey/rain/internal/cmd/rm"
	"github.com/jumziey/rain/internal/cmd/tree"
	"github.com/jumziey/rain/internal/cmd/watch"
	"github.com/jumziey/rain/internal/console"
)

// Cmd is the rain command's entrypoint
var Cmd = &cobra.Command{
	Use:     "rain",
	Long:    "Rain is a command line tool for working with AWS CloudFormation templates and stacks",
	Version: config.VERSION,
}

const usageTemplate = `Usage:{{if .Runnable}}
  <cyan>{{.UseLine}}</>{{end}}{{if .HasAvailableSubCommands}}
  <cyan>{{.CommandPath}}</> [<gray>command</>]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

{{range $group := groups}}{{ $group }}:{{range $c := $.Commands}}{{if $c.IsAvailableCommand}}{{if eq $c.Annotations.Group $group}}
  <cyan>{{rpad $c.Name $c.NamePadding }}</> {{$c.Short}}{{end}}{{end}}{{end}}

{{end}}Other Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}{{if .Annotations.Group}}{{else}}
  <cyan>{{rpad .Name .NamePadding }}</> {{.Short}}{{end}}{{end}}{{end}}{{end}}{{if and .HasParent .HasAvailableFlags}}

Flags:
{{.Flags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}
`

func init() {
	Cmd.AddCommand(build.Cmd)
	Cmd.AddCommand(cat.Cmd)
	Cmd.AddCommand(consolecmd.Cmd)
	Cmd.AddCommand(deploy.Cmd)
	Cmd.AddCommand(diff.Cmd)
	Cmd.AddCommand(fmt.Cmd)
	Cmd.AddCommand(info.Cmd)
	Cmd.AddCommand(logs.Cmd)
	Cmd.AddCommand(ls.Cmd)
	Cmd.AddCommand(merge.Cmd)
	Cmd.AddCommand(rm.Cmd)
	Cmd.AddCommand(tree.Cmd)
	Cmd.AddCommand(watch.Cmd)

	for _, c := range Cmd.Commands() {
		if c.Annotations[cmd.GroupAnnotationLabel] == cmd.StackGroup || c == info.Cmd || c == consolecmd.Cmd {
			c.Flags().StringVarP(&config.Profile, "profile", "p", "", "AWS profile name; read from the AWS CLI configuration file")
			c.Flags().StringVarP(&config.Region, "region", "r", "", "AWS region to use")
		}
	}

	// Customise usage
	Cmd.Annotations = map[string]string{"Groups": strings.Join(cmd.Groups, "|")}

	cobra.AddTemplateFunc("groups", func() []string {
		return cmd.Groups
	})

	oldUsageFunc := Cmd.UsageFunc()
	Cmd.SetUsageFunc(func(c *cobra.Command) error {
		Cmd.SetUsageTemplate(console.Sprint(usageTemplate))
		return oldUsageFunc(c)
	})

	Cmd.PersistentFlags().BoolVarP(&console.NoColour, "no-colour", "", false, "Disable colour output")

	cmd.AddDefaults(Cmd)
}
