## fifa version

Imprime la versión de fifa

### Synopsis


Imprime la versión de fifa

```
fifa version [flags]
```

### Options

```
  -h, --help   help for version
```

### Options inherited from parent commands

```
  -c, --conf string   Archivo de configuración.  Soporta formato yaml, json, .properties
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
		 	
  -v, --verbose       Imprime información extra
```

### SEE ALSO
* [fifa](fifa.md)	 - Flexible and Improved File Analizer

###### Auto generated by spf13/cobra on 14-Mar-2018
