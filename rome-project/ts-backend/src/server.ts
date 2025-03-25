import express, { Request, Response } from "express";
import axios from "axios";
import dotenv from "dotenv";
import { paths } from "../gen/api";
import path from "path";

// Load .env from ts-backend/.env
dotenv.config({ path: path.resolve(__dirname, "../.env") });

const app = express();
const PORT = process.env.PORT || 5002;
const GO_BACKEND_URL = process.env.GO_BACKEND_URL || "http://localhost:5001/log";

app.use(express.json());

// Message type from OpenAPI
type MessageType =
  paths["/log"]["post"]["requestBody"]["content"]["application/json"];

// Send message from TS to Go
app.post(
  "/send-to-go",
  async (req: Request<{}, {}, MessageType>, res: Response) => {
    const message = req.body;

    try {
      await axios.post(GO_BACKEND_URL, message);
      console.log("Message forwarded to Go:", message);
      res.status(200).json({ status: "forwarded" });
    } catch (error) {
      console.error("Error forwarding message to Go:", error);
      res.status(500).json({ error: "Failed to forward message" });
    }
  }
);

// Receive message from Go
app.post("/log", (req: Request<{}, {}, MessageType>, res: Response) => {
  const { from, to, message, date } = req.body;
  console.log(`Message from ${from} to ${to} @ ${date}: ${message}`);
  res.status(200).json({ status: "received" });
});

app.listen(PORT, () => {
  console.log(`TypeScript backend is running on port ${PORT}`);
});
