package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

// Result es el resultado del parseo
type Result struct {
	File   string
	Passed bool
	Detail []ItemResult
}

// ItemResult es el resultado de cada análisis
type ItemResult struct {
	CaseName        string
	CaseDescription string
	Passed          bool
}

var (
	parseCmd = &cobra.Command{
		Use:   "parse",
		Short: "Parsea archivos",
		Long:  `Parsea archivos`,
		Run: func(cmd *cobra.Command, args []string) {
			results := make([]Result, 0)

			if len(args) == 0 {
				fmt.Fprintf(os.Stderr, "No se recibieron archivos a procesar\n")
				os.Exit(0)
			}

			for _, f := range args {
				stat, err := os.Stat(f)
				if os.IsNotExist(err) {
					fmt.Fprintf(os.Stderr, "%s: No existe", f)
					continue
				} else if !stat.Mode().IsRegular() {
					fmt.Fprintf(os.Stderr, "%s: No es  un archivo", f)
					continue
				}

				detail, passed := processFile(f)
				result := Result{
					File:   f,
					Passed: passed,
					Detail: detail,
				}

				results = append(results, result)

			}
			jsonResult, _ := json.Marshal(results)
			fmt.Println(string(jsonResult))

		},
	}
)

func processFile(path string) ([]ItemResult, bool) {
	result := make([]ItemResult, 0)
	passed := true

	for _, r := range fifaConf.Regex {
		_, filename := filepath.Split(path)
		if matched, _ := filepath.Match(r.Glob, filename); matched {
			if verbose {
				fmt.Println(fmt.Sprintf("Procesando %s con %s", filename, r.Glob))
			}
			re, err := regexp.Compile(r.Pattern)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Expresión regular inválida: %s", r.Pattern)
				os.Exit(1)
			}
			if data, err := ioutil.ReadFile(path); err != nil {
				fmt.Fprintf(os.Stderr, "Error leyendo archivo: %s", path)
				os.Exit(1)
			} else {
				result = append(result, ItemResult{
					CaseName:        r.Name,
					CaseDescription: r.Description,
					Passed:          !re.Match(data),
				})
				passed = !re.Match(data) && passed
			}
		}

	}
	return result, passed

}

func init() {
	rootCmd.AddCommand(parseCmd)
}
