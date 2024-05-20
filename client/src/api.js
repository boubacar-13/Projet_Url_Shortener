import axios from "axios";

// const API_BASE_URL = process.env.API_URL; //golang server url
const API_BASE_URL = "http://localhost:3000"; //golang server url

export const shortenUrl = async (originalUrl) => {
  const response = await axios.post(`${API_BASE_URL}/api/v1/`, {
    originalUrl,
  });
  console.log("La réponse :" + response);
  return response.data;
};

export const resolveUrl = async (shortUrl) => {
  const response = await axios.get(`${API_BASE_URL}/resolve/${shortUrl}`);
  return response.data;
};
