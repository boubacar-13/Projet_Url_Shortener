# Backend API

## Description

Cette partie du projet gère les opérations backend, y compris la gestion des URL raccourcies, le suivi des clics et l'interaction avec la base de données.

## Fonctionnalités

- Raccourcissement d'URLs
- Limitation du taux d'utilisation
- Calcule des statistiques des URLs raccourcies
- Création de l'historique des URLs raccourcies par l'utilisateur
- En cours de développement : Authentification des utilisateurs (inscription et connexion)

## Lancement en Développement

Pour lancer le serveur en mode développement :

```bash
  go run main
```

## API Endpoints

POST /api/v1/ : Crée une nouvelle URL raccourcie.
GET /api/v1/stats : Récupère les statistiques d'utilisation.
GET /api/v1/history : Récupère l'historique d'utilisation.
