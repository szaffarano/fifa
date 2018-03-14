package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/szaffarano/fifa/conf"
)

var (
	verbose bool
	cfgFile string

	fifaConf conf.FifaConf

	rootCmd = &cobra.Command{
		Use:   "fifa",
		Short: "Flexible and Improved File Analizer",
		Long: `Parsea y analiza archivos en búsqueda de patrones, según 
		configuración`,
	}
)

// Execute es el punto de entrada de la cli
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initconfig)

	// inicializa flags globales
	rootCmd.
		PersistentFlags().
		BoolVarP(
			&verbose,
			"verbose",
			"v",
			false,
			"Imprime información extra")

	rootCmd.
		PersistentFlags().
		StringVarP(
			&cfgFile,
			"conf",
			"c",
			"",
			`Archivo de configuración.  Soporta formato yaml, json, .properties
			y hcl.  Si no se especifica, por default buscará los archivos
			- $HOME/.fifa/conf.[yaml|properties|json]
			- ./conf.[yaml|properties|json]
			
			Un ejemplo de configuración en formato yaml sería:

			regex:
				- name: id1
				  description: Validación 1
				  glob: *.txt
				  pattern: "Hola\\sMundo"

				- name: id2
				  description: Validación 2
				  glob: "*"
				  pattern: "[a-z][1-9]{2,3}"
		 	`)
}

// initconfig lee la configuración
func initconfig() {
	config := viper.New()

	if cfgFile != "" {
		ext := filepath.Ext(cfgFile)
		if len(ext) > 0 {
			config.SetConfigFile(cfgFile)
			config.SetConfigType(ext[1:])
		} else {
			panic(fmt.Errorf("El archivo de configuración debe tener extensión: %s", cfgFile))
		}
	} else {
		config.SetConfigName("fifa")
	}

	config.AddConfigPath("$HOME/.fifa")
	config.AddConfigPath(".")

	config.AutomaticEnv()
	config.SetEnvPrefix("FIFA")

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	config.Unmarshal(&fifaConf)

	if err := fifaConf.Validate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else if verbose {
		fmt.Println("Usando configuración:", config.ConfigFileUsed())

	}
}
