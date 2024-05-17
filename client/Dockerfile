# Utiliser l'image officielle de Node.js comme image de base
FROM node:14

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers package.json et package-lock.json
COPY package*.json ./

# Installer les dépendances
RUN npm install

# Copier le reste des fichiers du projet
COPY . .

# Exposer le port que l'application utilise
EXPOSE 3000

# Commande pour démarrer l'application
CMD ["npm", "start"]
