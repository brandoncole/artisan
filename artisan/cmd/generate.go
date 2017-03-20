package cmd

import (
	"os"
	"path"

	"github.com/brandoncole/artisan/artisan/model"
	"github.com/brandoncole/artisan/artisan/utility"
	"github.com/spf13/cobra"

	"strings"
)

var flagManifest string
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates artisan scaffolding",
	Run: func(cmd *cobra.Command, args []string) {

		manifest, err := model.NewManifest(flagManifest)
		if nil != err {
			panic(err)
		}

		err = os.RemoveAll(manifest.Resolve(manifest.Outputs.Path))
		if nil != err {
			panic(err)
		}

		schema := manifest.Resolve(path.Join(manifest.Outputs.Path, "schema"))
		grpc := manifest.Resolve(path.Join(manifest.Outputs.Path, "grpc"))
		err = os.MkdirAll(schema, 0766)
		if nil != err {
			panic(err)
		}
		err = os.MkdirAll(grpc, 0766)
		if nil != err {
			panic(err)
		}

		err = utility.Shell("/bin/sh",
			"-c",
			strings.Join([]string{
				"protoc ",
				"-o",
				path.Join(schema, "description.pb"),
				"--include_imports ",
				"--include_source_info ",
				"--go_out=plugins=grpc:" + grpc,
				"--proto_path=/artisan/include",
				"--proto_path=" + manifest.Resolve(manifest.Inputs.Path),
				"$(find " + manifest.Resolve(manifest.Inputs.Path) + " -iname *.proto)",
			}, " "),
		)
		if nil != err {
			panic(err)
		}

	},
}

func init() {
	generateCmd.Flags().StringVarP(&flagManifest, "file", "f", "", "Specifies the path to the manifest")
	RootCmd.AddCommand(generateCmd)
}
