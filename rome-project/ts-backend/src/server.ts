import express, { Request, Response } from "express";
import axios from "axios";
import dotenv from "dotenv";
import { paths } from "../gen/api";
import path from "path";
import { PrismaClient } from "@prisma/client";

const prisma = new PrismaClient();

// Load .env from ts-backend/.env
dotenv.config({ path: path.resolve(__dirname, "../.env") });

const app = express();
const PORT = process.env.PORT || 5002;
const GO_BACKEND_URL =
  process.env.GO_BACKEND_URL || "http://localhost:5001/log";

app.use(express.json());

// Message type from OpenAPI
type MessageType =
  paths["/log"]["post"]["requestBody"]["content"]["application/json"];

// Send message from TS to Go
app.post("/send", async (req: Request<{}, {}, MessageType>, res: Response) => {
  const message = req.body;

  try {
    await axios.post(GO_BACKEND_URL, message);
    console.log("Message forwarded to Go:", message);
    res.status(200).json({ status: "forwarded" });
  } catch (error) {
    console.error("Error forwarding message to Go:", error);
    res.status(500).json({ error: "Failed to forward message" });
  }
});

// Receive message from Go
app.post("/log", async (req: Request<{}, {}, MessageType>, res: Response) => {
  const { from, to, message, date } = req.body;
  console.log(`Message from ${from} to ${to} @ ${date}: ${message}`);

  await prisma.message.create({
    data: {
      sender: from,
      receiver: to,
      message,
      timestamp: new Date(date)
    }
  });

  res.status(200).json({ status: "received" });
});

app.get("/logs", async (_req, res) => {
  const logs = await prisma.message.findMany({
    orderBy: { timestamp: "desc" },
    take: 10
  });
  res.json(logs);
});

app.listen(PORT, () => {
  console.log(`TypeScript backend is running on port ${PORT}`);
});
