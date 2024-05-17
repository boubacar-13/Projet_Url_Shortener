import React, { useState } from "react";
import "./App.css";
import { shortenUrl } from "../../client/src/api";

function App() {
  const [originalUrl, setOriginalUrl] = useState("");
  const [shortUrl, setShortUrl] = useState("");
  const [error, setError] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const result = await shortenUrl(originalUrl);
      setShortUrl(result.shortUrl);
      setError("");
    } catch (err) {
      setError("Erreur lors du raccourcissement de l'URL.");
    }
  };

  return (
    <div className="App">
      <h1>Raccourcisseur d'URL</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={originalUrl}
          onChange={(e) => setOriginalUrl(e.target.value)}
          placeholder="Entrez votre URL"
        />
        <button type="submit">Raccourcir</button>
      </form>
      {shortUrl && (
        <div>
          <h2>URL Raccourcie :</h2>
          <a href={shortUrl} target="_blank" rel="noopener noreferrer">
            {shortUrl}
          </a>
        </div>
      )}
      {error && <p className="error">{error}</p>}
    </div>
  );
}

export default App;
