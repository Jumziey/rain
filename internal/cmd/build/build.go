package build

import (
	"fmt"
	"github.com/jumziey/rain/internal/cmd"
	"sort"
	"strings"

	"github.com/jumziey/rain/cft/build"
	"github.com/jumziey/rain/cft/format"
	"github.com/jumziey/rain/cft/spec"
	"github.com/spf13/cobra"
)

var buildListFlag = false
var bareTemplate = false
var buildJSON = false

// Cmd is the build command's entrypoint
var Cmd = &cobra.Command{
	Use:                   "build [<resource type>...]",
	Short:                 "Create CloudFormation templates",
	Long:                  "Outputs a CloudFormation template containing the named resource types.",
	Annotations:           cmd.TemplateAnnotation,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		if buildListFlag {
			types := make([]string, 0)
			for t := range spec.Cfn.ResourceTypes {
				types = append(types, t)
			}
			sort.Strings(types)
			fmt.Println(strings.Join(types, "\n"))

			return
		}

		if len(args) == 0 {
			cmd.Help()
			return
		}

		resources := resolveResources(args)

		t, err := build.Template(resources, !bareTemplate)
		if err != nil {
			panic(err)
		}

		out := format.String(t, format.Options{
			JSON: buildJSON,
		})

		fmt.Println(out)
	},
}

func init() {
	Cmd.Flags().BoolVarP(&buildListFlag, "list", "l", false, "List all CloudFormation resource types")
	Cmd.Flags().BoolVarP(&bareTemplate, "bare", "b", false, "Produce a minimal template, omitting all optional resource properties")
	Cmd.Flags().BoolVarP(&buildJSON, "json", "j", false, "Output the template as JSON (default format: YAML)")
}
