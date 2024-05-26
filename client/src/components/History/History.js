import React, { useEffect, useState } from "react";
import axios from "axios";
import "./History.css";

const History = () => {
  const [history, setHistory] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchHistory = async () => {
      try {
        const res = await axios.get("http://localhost:3000/api/v1/history");
        setHistory(res.data);
        setError(null);
      } catch (err) {
        setError(
          err.response
            ? err.response.data
            : "Erreur lors de la récupération de l'historique"
        );
      }
    };
    fetchHistory();
  }, []);

  return (
    <div className="container">
      <h1>Historique des URL raccourcies</h1>
      {error && <div className="error">{error}</div>}
      <div className="table-container">
        <table>
          <thead>
            <tr>
              <th>URL Raccourcie</th>
              <th>URL Longue</th>
              <th>Date de Création</th>
            </tr>
          </thead>
          <tbody>
            {history.map((item, index) => (
              <tr key={index}>
                <td>
                  <a
                    href={item.short_url}
                    target="_blank"
                    rel="noopener noreferrer"
                  >
                    {item.short_url}
                  </a>
                </td>
                <td>{item.long_url}</td>
                <td>{item.created_at}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default History;
