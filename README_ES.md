# Template Gin Golang

## Instalación
Para instalar este proyecto es necesario tener docker instalado.
Una vez instalado docker, podemos empezar la instalación.

Ejecutamos este comando para construir el dockerfile, es importante 
```
docker build . -t template-go
```

Después procedemos a ejecutar dicha imagen con:
```
docker run -i -t -p 80:8080 template-go
```

El nombre de la imagen puede ser cambiada, es decir, en vez de poner `template-go`
podemos poner un nombre distinto, lo que si es importante es poner una sintaxis correcta (En minusculas y con guiones en vez de espacios)

`IMPORTANTE: En caso de modificacion de la version en Go.mod, es importante cambiarla tambien en el dockerfile. En el caso de no ser cambiada, puede llegar a dar errores.`
