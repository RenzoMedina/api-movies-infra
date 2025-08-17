# 🎬 API Movies Infra – Infraestructura para Movie API

Este repositorio contiene la infraestructura como código [IaC](https://github.com/RenzoMedina/api-movies-infra/tree/main/infra) para desplegar la [Movie API](https://github.com/RenzoMedina/api-movies-infra/tree/main/app), una API REST construida en Go con Echo. La infraestructura se define con Terraform y se automatiza con GitHub Actions.


# 🔐 Requisitos

| Secreto             | Descripción    |
| ----------------- | ------------------------ |
| AWS_ACCESS_KEY_ID | Access Key de AWS |
| AWS_SECRET_ACCESS_KEY | Secret Key de AWS |
| DOCKER_USERNAME | Token de Docker tu user name | 
| DOCKER_PASSWORD |  Token de Docker tu password |  


# 📦 Estructura del proyecto

``` plaintext
api-movies-infra/ 
├── app/ # Archivos Go con ECHO
├── infra/ # Archivos Terraform (VPC, EC2) 
├── Dockerfile # Imagen para la API 
├── .github/workflows/ # Workflows de CI/CD 
└── README.md

```

# 🧨 Destrucción automática

```bash
El workflow Destroy Infraestructura:

con la rama "destroy"

```
    
## Authors
- Backend Developer & DevOps Jr
- [@renzomedina](https://github.com/RenzoMedina)

