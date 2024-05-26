import React, { useEffect, useState } from "react";
import axios from "axios";
import "./Stats.css";

const Stats = () => {
  const [stats, setStats] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const res = await axios.get("http://localhost:3000/api/v1/stats");
        setStats(res.data);
        setError(null);
      } catch (err) {
        setError(
          err.response
            ? err.response.data
            : "Erreur lors de la récupération des statistiques"
        );
      }
    };
    fetchStats();
  }, []);

  return (
    <div className="container">
      <h1>Statistiques</h1>
      {error && <div className="error">{error}</div>}
      {stats ? (
        stats.total_urls > 0 ? (
          <div className="stats">
            <h2>Total des URLs raccourcies: {stats.total_urls}</h2>
            <h3>Clics par URL:</h3>
            <ul>
              {Object.entries(stats.clicks).map(([shortURL, clicks]) => (
                <li key={shortURL}>
                  {shortURL}: {clicks} clics
                </li>
              ))}
            </ul>
          </div>
        ) : (
          <div className="no-stats">
            Aucune URL n'a été raccourcie pour le moment.
          </div>
        )
      ) : (
        <div>Chargement des statistiques...</div>
      )}
    </div>
  );
};

export default Stats;
