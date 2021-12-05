# BFF Elentari - Proyecto Kümelen

## Descripción del Proyecto

"Kümelen" nace desde la necesidad de implementar, con ayuda de elementos tecnológicos, nuevos canales de comunicación
que le permitan a los estudiantes de la UTEM conocer e informarse sobre los servicios estudiantiles que pueden apoyar
en las dimensiones extraacadémicas de su vida universitaria.

## Descripción del artefacto

El artefacto "Elentari" es una API con finalidad de Backend for Frontend que procesa y envía la información necesaria a
la [aplicación móvil del proyecto](https://github.com/DiegoSepuSoto/manwe).

## Variables de entorno

Las variables de entorno necesarias para la ejecución del artefacto son:

APP_VERSION
ENV
PORT
ILUVATAR_CMS_HOST
ILUVATAR_TOKEN
ILUVATAR_HOST

## Entorno de compilación/ejecución

Para compilar el código es necesario configurar un ambiente de Go en su versión 1.15

La ejecución del mismo es totalmente agnóstica, dado que de la compilación es generado un archivo binario ejecutable

## Generación/Visualización de documentación

Para la generación/actualización de la documentación del proyecto se requiere ejecutar el siguiente comando:

```console
foo@bar:~$ swag init -g src/main.go
```

Luego, para su visualización, es necesario entrar a la siguiente ruta:

http://servidor.com/swagger/index.html

### Nombre en clave

El nombre en clave utilizado para este artefacto, "Elentari", viene de la mitología creada por el autor J.R.R. Tolkien,
donde Elentari es uno de los Valär más importantes creados por Ilúvatar. [Ver Más](https://es.wikipedia.org/wiki/Varda)