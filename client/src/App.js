import React, { useState } from "react";
import axios from "axios";
import "./App.css";

const App = () => {
  const [url, setUrl] = useState("");
  const [short, setShort] = useState("");
  const [expiry, setExpiry] = useState(24); // default to 24 hours
  const [response, setResponse] = useState(null);
  const [error, setError] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const res = await axios.post("http://localhost:3000/api/v1/", {
        url,
        short,
        expiry,
      });
      setResponse(res.data);
      setError(null);
    } catch (err) {
      setError(
        err.response
          ? err.response.data
          : "Erreur lors du raccourcissement de l'URL"
      );
    }
  };

  return (
    <div className="container">
      <h1>Raccourcisseur d'URL</h1>
      <form onSubmit={handleSubmit} className="form">
        <div className="form-group">
          <label>URL:</label>
          <input
            type="text"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Personnaliser l'URL:</label>
          <input
            type="text"
            value={short}
            onChange={(e) => setShort(e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>Doit expirer dans (nombre d'heures) :</label>
          <input
            type="number"
            value={expiry}
            onChange={(e) => setExpiry(Number(e.target.value))}
            min="1"
            required
          />
        </div>
        <button type="submit" className="btn">
          Raccourcir l'URL
        </button>
      </form>
      {response && (
        <div className="response">
          <h2>URL raccourcie</h2>
          <p>
            <span className="responsePara">URL initiale:</span> {response.url}
          </p>
          <p>
            <span className="responsePara">URL raccourcie:</span>{" "}
          </p>
          <a
            className="responsePara"
            href={response.short}
            target="_blank"
            rel="noopener noreferrer"
          >
            {response.short}
          </a>
          <br />
          <p>
            <span className="responsePara">Expire dans:</span> {response.expiry}{" "}
            heures
          </p>
          <p>
            <span className="responsePara">
              Quota d'utilisations restant de l'outil de raccourcissement:
            </span>{" "}
            {response.rate_limit}
          </p>
          <p>
            <span className="responsePara">
              Réinitialisation du quota d'utilisation dans:
            </span>{" "}
            {response.rate_limit_reset} minutes
          </p>
        </div>
      )}
      {error && <div className="error">{error}</div>}
    </div>
  );
};

export default App;
