package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// Version es la version de la aplicación
	Version = "desconocida"

	// BuildDate indica la fecha de construcción
	BuildDate = "desconocida"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Imprime la versión de fifa",
	Long:  `Imprime la versión de fifa`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("fifa %s\n", Version)
		fmt.Printf("  Fecha de construcción: %s\n", BuildDate)
		fmt.Printf("  Compilado con: %s\n", runtime.Version())
	},
}
