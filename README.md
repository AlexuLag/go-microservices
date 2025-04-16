# Plataforma de Microservicios

Esta plataforma implementa una arquitectura de microservicios usando Go y Docker.

## Servicios incluidos

1. **Broker Service**: Actúa como punto de entrada para la comunicación entre servicios.
2. **Authentication Service**: Maneja la autenticación de usuarios.
3. **Logger Service**: Registra eventos del sistema.
4. **Mail Service**: Gestiona el envío de correos electrónicos.
5. **Listener Service**: Escucha y reacciona a eventos.
6. **Front-end**: Interfaz de usuario para interactuar con los servicios.

## Requisitos previos

- Go 1.22.0 o superior
- Docker y Docker Compose
- Make (para entornos Unix) o `make.exe` (para Windows)
- protoc (compilador de Protocol Buffers, para gRPC)

## Obtención de dependencias

Cada microservicio tiene sus propias dependencias. Para instalarlas, utiliza los siguientes comandos:

### Instalación global de herramientas gRPC (si se utiliza gRPC)

```bash
# Instalar el generador de código para protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Instalar el generador de código para gRPC
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### Para cada servicio

Navega al directorio del servicio y ejecuta:

```bash
cd [service-directory]
go mod tidy
```

Por ejemplo:

```bash
cd broker-service
go mod tidy
```

Repite este proceso para cada uno de los servicios (authentication-service, logger-service, mail-service, listener-service, front-end).

## Comandos Make

El proyecto incluye un Makefile con los siguientes comandos útiles:

### Iniciar todos los servicios

```bash
cd project
make up
```

Este comando inicia todos los contenedores sin reconstruirlos.

### Construir e iniciar todos los servicios

```bash
cd project
make up_build
```

Este comando detiene los contenedores si están en ejecución, reconstruye todos los binarios de los servicios y reinicia los contenedores.

### Detener todos los servicios

```bash
cd project
make down
```

### Construir servicios individuales

```bash
# Construir el servicio broker
make build_broker

# Construir el servicio de autenticación
make build_auth

# Construir el servicio de logging
make build_logger

# Construir el servicio de correo
make build_mail

# Construir el servicio de escucha
make build_listener

# Construir el front-end
make build_front
```

### Iniciar el front-end localmente

```bash
make start
```

### Detener el front-end local

```bash
make stop
```

## Estructura de comunicación

Los servicios se comunican entre sí utilizando diferentes mecanismos:

1. **HTTP/REST**: Comunicación básica entre servicios
2. **RPC**: Llamadas a procedimientos remotos para operaciones síncronas
3. **RabbitMQ**: Sistema de mensajería para comunicación asíncrona usando patrón topic
4. **gRPC**: Comunicación eficiente basada en Protocol Buffers (donde se implementa)

## Notas importantes

1. En entornos Windows, utiliza `project/Makefile.windows` en lugar del Makefile estándar.
2. Todos los binarios se construyen como ejecutables Linux para su uso en contenedores Docker.
3. Las bases de datos y otros servicios de infraestructura se gestionan a través de Docker Compose.

## Solución de problemas

Si experimentas problemas al ejecutar los servicios:

1. Verifica que los puertos necesarios estén disponibles.
2. Asegúrate de que MongoDB, RabbitMQ y otros servicios externos estén accesibles.
3. Comprueba los logs de Docker para identificar errores específicos:
   ```bash
   docker-compose logs [service-name]
   ``` 