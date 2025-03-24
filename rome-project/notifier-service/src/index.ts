import express from 'express';
import { Task } from '../../shared/types';

const app = express();
app.use(express.json());

app.post('/notify', (req, res) => {
  try {
    const task: Task = req.body;
    console.log(`[Reminder] ${task.title} at ${task.scheduledTime}`);
    res.status(200).send('Task received and reminder logged!');
  } catch (error) {
    console.error('Error in Notifier:', error);
    res.status(500).send('Failed to process task.');
  }
});

app.listen(4000, () => console.log('Notifier running on port 4000'));

const x =5

