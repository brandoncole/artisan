package cmd

import (
	"fmt"
	"github.com/brandoncole/artisan/artisan/api"
	"github.com/brandoncole/artisan/artisan/model"
	"github.com/brandoncole/artisan/artisan/utility"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path"
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

		data, err := ioutil.ReadFile(path.Join(schema, "description.pb"))
		if nil != err {
			panic(err)
		}

		description := &descriptor.FileDescriptorSet{}
		err = proto.Unmarshal(data, description)
		if nil != err {
			panic(err)
		}

		for _, item := range description.File {
			fmt.Printf("Package %s %s\n",
				item.GetName(),
				item.GetPackage(),
			)
			for _, svc := range item.GetService() {
				fmt.Printf("  Service %s %d %d\n",
					svc.GetName(),
					len(svc.GetOptions().GetUninterpretedOption()),
					len(svc.GetOptions().ExtensionRangeArray()),
				)
			}
			for _, msg := range item.GetMessageType() {
				fmt.Printf("  Message %s\n",
					msg.GetName(),
				)
				for _, field := range msg.GetField() {
					fmt.Printf("    Field %s %s\n",
						field.GetName(),
						field.GetTypeName(),
					)
					if nil != field.GetOptions() {
						v, err := proto.GetExtension(field.GetOptions(), api.E_StringIsDate)
						if nil == err {
							fmt.Printf("      Date: %v!\n", *(v.(*bool)))
						}
					}
				}
			}
		}

	},
}

func init() {
	generateCmd.Flags().StringVarP(&flagManifest, "file", "f", "", "Specifies the path to the manifest")
	RootCmd.AddCommand(generateCmd)
}
