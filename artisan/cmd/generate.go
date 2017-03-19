package cmd

import (
	"os"
	"path"

	"github.com/brandoncole/artisan/artisan/model"
	"github.com/spf13/cobra"
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
		err = os.MkdirAll(schema, 0700)
		if nil != err {
			panic(err)
		}
		err = os.MkdirAll(grpc, 0700)
		if nil != err {
			panic(err)
		}

	},
}

func init() {
	generateCmd.Flags().StringVarP(&flagManifest, "file", "f", "", "Specifies the path to the manifest")
	RootCmd.AddCommand(generateCmd)
}
