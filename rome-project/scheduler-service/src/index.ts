import express from 'express';
import { Task } from '../../shared/types';
import axios from 'axios';
import { v4 as uuidv4 } from 'uuid';

const app = express();
app.use(express.json());

app.post('/create-task', async (req, res) => {
  const { title, description, scheduledTime, createdBy } = req.body;

  const task: Task = {
    id: uuidv4(),
    title,
    description,
    scheduledTime,
    createdBy,
  };

  try {
    await axios.post('http://localhost:4000/notify', task);
    res.status(200).send('Task created and sent to notifier!');
  } catch (err) {
    res.status(500).send('Failed to send task to notifier.');
  }
});

app.listen(3000, () => console.log('Scheduler running on port 3000'));
