package conf

import "fmt"

// FifaConf es la representación del archivo de configuración
type FifaConf struct {
	Regex []regexCase
}

type regexCase struct {
	Name        string
	Description string
	Glob        string
	Pattern     string
}

// Validate verifica que se hayan recibido todos los parámetros necesarios
func (n *FifaConf) Validate() error {
	// @TODO
	return nil
}

func (c regexCase) String() string {
	return fmt.Sprintf(
		"RegexCase [ name: %s, description: %s, glob: %s, pattern: %s ]",
		c.Name,
		c.Description,
		c.Glob,
		c.Pattern)
}
