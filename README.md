# FIFA (Flexible and Improved File Analizer)

## Modo de uso

```sh
fifa [-v] <archivo>
```

Para más información consultar [la documentación](./doc/fifa.md).

## Configuración

La configuración por default deberá estar en un archivo `conf.EXTENSION`, siendo
`EXTENSION` cualquiera de las siguientes:

- yaml
- json
- .properties

Un ejemplo de un archivo de configuración yaml es

```yaml
regex:
    - name: id1
        description: Validación 1
        glob: *.txt
        pattern: "Hola\\sMundo"

    - name: id2
        description: Validación 2
        glob: "*"
        pattern: "[a-z][1-9]{2,3}"
```

## Compilación

El archivo `Makefile` provisto construye la aplicación compatible con linux y
windows 32 y 64 bits.  previo a invocarlo se necesitan instalar las 
dependencias, en forma global, o mejor aún local, con glide:

```sh
$ glide i
```

Luego ejecutar

```sh
# compilación únicamente para linux
$ make compile

# compilación multiplataforma (demora más tiempo)
$ make release
```
