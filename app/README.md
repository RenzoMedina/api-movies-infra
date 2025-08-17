
# 🎬 Movie API – by Renzo Medina

Una API REST sencilla construida con Echo en Go, que permite listar y agregar películas. Los datos se almacenan en un archivo local movies.json.
# 🚀 Endpoints

| Método             | Ruta    |Descripción|
| ----------------- | ------------------------ |-------------------------------|
| GET | / | Mensaje de bienvenida |
| GET | /movies | Devuelve la lista de películas |
| POST | /movies |  Agrega una nueva película |



# 📦 Estructura del proyecto

``` plaintext
app-demo/
├── main.go
├── pkg/
│   └── movies.go
├── movies.json
├── go.mod
├── Dockerfile

```

# 📥 Instalación

```bash
git clone https://github.com/RenzoMedina/app-demo.git
cd app-demo
go mod tidy
go run main.go

```
    
# 🐳 Docker

```
docker build -t movie-api .
docker run -p 8080:8080 movie-api

``` 
# 📄 Ejemplo de JSON para POST

``` json
{
  "id": 1,
  "name": "Inception",
  "description": "A mind-bending thriller",
  "category": ["Sci-Fi", "Thriller"]
}


```

## Authors
- Backend Developer & DevOps Jr
- [@renzomedina](https://github.com/RenzoMedina)

