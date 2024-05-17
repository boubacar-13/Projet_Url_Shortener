import axios from "axios";

const API_BASE_URL = "http://localhost:8080";

export const shortenUrl = async (originalUrl) => {
  const response = await axios.post(`${API_BASE_URL}/api/v1/`, { originalUrl });
  return response.data;
};

export const resolveUrl = async (shortUrl) => {
  const response = await axios.get(`${API_BASE_URL}/resolve/${shortUrl}`);
  return response.data;
};
