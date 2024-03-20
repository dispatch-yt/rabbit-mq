const amqp = require('amqplib');

async function produce() {
  const connection = await amqp.connect('amqp://guest:guest@localhost:5672/');
  const channel = await connection.createChannel();
  const queue = 'hello';

  let count = 0;
  setInterval(() => {
    const message = `Message ${count++}`;
    channel.sendToQueue(queue, Buffer.from(message));
    console.log(`[x] Sent ${message}`);
  }, 1000);
}

produce();
