# Establece la imagen base como golang:alpine, que es una imagen liviana de Go basada en Alpine Linux.
FROM golang:alpine

# Agrega una etiqueta al Dockerfile para especificar el responsable o mantenedor del archivo.
LABEL maintainer="cristian bustamante"

# Install git.
# Ejecuta comandos en el sistema operativo Alpine Linux dentro del contenedor para actualizar el índice de paquetes (apk update).
# Instala git (necesario para obtener las dependencias), bash y build-base (paquetes requeridos para construir algunas bibliotecas Go).

RUN apk update && apk add --no-cache git bash build-base

# Crea un directorio /app dentro del contenedor y lo establece como el directorio de trabajo actual. 
# Es el lugar donde se copiarán los archivos del proyecto.
RUN mkdir /app
WORKDIR /app

# Copia solo el archivo go.mod y go.sum primero.
# Esto permite que Docker utilice la caché de capas para las dependencias si no se han modificado.
COPY go.mod .
COPY go.sum .

# Descarga las dependencias del proyecto utilizando el comando go mod download.
# Esto aprovecha la caché de las dependencias si no se han modificado.
RUN go mod download

# Copia el resto de los archivos del proyecto.
COPY . .

# Compila el código fuente del proyecto Go en un archivo ejecutable llamado /build.
# El comando go build compila el código y crea un ejecutable en la ubicación especificada por la opción -o.
RUN go build -o /build

# Expone el puerto 8080 del contenedor.
# Esto indica que el contenedor está escuchando en el puerto 8080 y permite la comunicación a través de ese puerto.
EXPOSE 8080

# Define el comando que se ejecutará cuando se inicie el contenedor.
# En este caso, se ejecuta el ejecutable /build generado anteriormente.
CMD ["/build"]