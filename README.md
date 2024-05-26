# Projet de Raccourcissement d'URL

## Description

Ce projet est une application de raccourcissement d'URL avec une interface utilisateur pour raccourcir des URL, afficher les statistiques et consulter l'historique des URL raccourcies.

## Structure du Projet

- **api/** : Contient le backend de l'application.
- **client/** : Contient le frontend de l'application.
- **data/** : Contient les données de l'application (fichiers statiques, etc.).
- **db/** : Contient les scripts et configurations pour la base de données.
- **docker-compose.yml** : Fichier de configuration Docker pour orchestrer les différents services.

## Prérequis

- Docker
- Docker Compose

## Installation

1. Clonez le dépôt :
   git clone https://github.com/boubacar-13/Projet_Url_Shortener.git
   cd le-repo

2. Lancez les services avec Docker Compose :
   docker-compose up --build

## Utilisation

Une fois les services démarrés, vous pouvez accéder à l'application à http://localhost:3001.



## Documentation des Sous-Projets
Pour plus de détails sur les sous-projets, consultez les README dans les dossiers respectifs :
Backend API
Frontend Client
Données
Base de Données
