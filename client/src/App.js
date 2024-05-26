import React from "react";
import { BrowserRouter, Routes, Route, Link } from "react-router-dom";
import Stats from "./components/Stats/Stats";
import Shorten from "./components/Shorten/Shorten";
import History from "./components/History/History";
import "./App.css";

const App = () => {
  return (
    <>
      <BrowserRouter>
        <div className="container">
          <nav>
            <ul>
              <li>
                <Link to="/shorten">Raccourcir une URL</Link>
              </li>
              <li>
                <Link to="/stats">Voir les statistiques</Link>
              </li>
              <li>
                <Link to="/history">Voir l'historique</Link>
              </li>
            </ul>
          </nav>
          <Routes>
            <Route path="/shorten" element={<Shorten />} />
            <Route path="/stats" element={<Stats />} />
            <Route path="/history" element={<History />} />
          </Routes>
        </div>
      </BrowserRouter>
    </>
  );
};

export default App;
