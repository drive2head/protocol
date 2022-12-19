const net = require('net');
const TcpClient = require("./tcp-client");

const client = new TcpClient(new net.Socket());
const log = console.log

client.connect(8081, '127.0.0.1', function() {
	console.log('Connected');
	client.write('Hello, server! Love, Client.').catch(log);
}).catch(log)

client.on('data', function(data) {
	console.log('Received: ' + data);
	client.destroy().catch(log);
}).catch(log);

client.on('close', function() {
	console.log('Connection closed');
}).catch(log);
