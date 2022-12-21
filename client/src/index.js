const net = require('net');
const { exit } = require('process');
const TcpClient = require("./tcp-client");

const client = new TcpClient(new net.Socket(), 8081, '127.0.0.1', function() {
	console.log('Connected');
});
const log = console.log

const markup = { id: null, data: [] };

function isJson(str) {
    try {
        JSON.parse(str);
    } catch (e) {
        return false;
    }
    return true;
}

function generateRandomData() {
	const CHARACTERS  = 'abcdefghijklmnopqrstuvwxyz';
	const N = CHARACTERS.length;

	const intervalSeconds = 2;

	setInterval(() => {
		randomChar = CHARACTERS.charAt(Math.floor(Math.random() * N))
		markup.data.push(randomChar)
		console.log('markup:', markup)
	}, intervalSeconds * 1000);
}

function sendMarkup() {
	const buf = Buffer.from(JSON.stringify(markup))
	client.write(buf)

	console.log('MARKUP WAS SENT!')
}

client.connect().catch(log)

let markups = []

client.on('data', function(_data) {
	const data = _data.toString();

	if (isJson(data)) {
		parsedMarkup = JSON.parse(data)
		markup.id = parsedMarkup.id
		markup.data = parsedMarkup.data

		generateRandomData()
	} else {
		command = data
		switch (command) {
			case 'REQUEST_DATA': {
				sendMarkup()

				break
			}
			case 'OK': {
				client.onOk()

				break
			}
			default: {
				console.log(`Received command '${command}' is not supported`)

				break
			}
		}
	}

	// client.destroy().catch(log);
}).catch(log);

client.on('close', function() {
	console.log('Connection closed');
	exit()
}).catch(log);
