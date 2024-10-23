# Utiliser l'image de base Go
FROM golang:1.20-alpine

# Créer un dossier pour l'application
WORKDIR /app

# Copier le go.mod et go.sum pour installer les dépendances en avance
COPY go.mod go.sum ./

# Installer les dépendances
RUN go mod download

# Copier le reste des fichiers de l'application
COPY . .

# Installer `swag` pour la génération de documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Générer la documentation Swagger
RUN swag init

# Installer `air` pour le hot reload
RUN go install github.com/cosmtrek/air@latest

# Commande par défaut pour lancer `air`
CMD ["air"]
