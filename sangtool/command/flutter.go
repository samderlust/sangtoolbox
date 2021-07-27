package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"com.samderlust/sangtoolbox/sangtool/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	templateFlag = "template"
)

//FlutterCreate create flutter project and customized folders
func FlutterCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flutter_create <name>",
		Short: "create a new Flutter project",
		Long:  `create a new Flutter project and template folder`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			cwd, err := os.Getwd()
			if err != nil {
				return errors.Wrap(err, "failed to Getwd")
			}

			template, _ := cmd.Flags().GetString(templateFlag)
			path := filepath.Join(cwd, name)

			_, b, _, _ := runtime.Caller(0)
			basepath := filepath.Join(filepath.Dir(b), "../..")
			templatePath := filepath.Join(basepath, fmt.Sprintf("sangtool/templates/%s.json", template))

			if err := os.Mkdir(path, 0777); err != nil {
				// return errors.Wrap(err, fmt.Sprintf("failed to create project %s", name))
				e := errors.Wrap(err, fmt.Sprintf("failed to create project %s", name))
				fmt.Println(e)
			}

			templFile, err := ioutil.ReadFile(templatePath)

			if err != nil {
				return errors.Wrap(err, "failed to read template file")
			}

			var data interface{}
			err = json.Unmarshal(templFile, &data)

			if err != nil {
				return errors.Wrap(err, "reading json err")
			}

			if err := utils.CreateDirsRecursive(data, path); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringP(templateFlag, "t", "example", "the tamplate that will be use, default to example")

	return cmd

}
