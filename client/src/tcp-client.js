class TcpClient {
    constructor(client) {
        this.client = client
    }

    ping() {
        this.client.write('ping')
    }

    connect(port, host, cb) {
        this.client.connect(port, host, cb);
    }

    write(message) {
        try {
            this.client.write(message);
        } catch (e) {
            // TODO implement
            console.log(e)
        }
    }

    destroy() {
        this.client.destroy()
    }

    on(event, cb) {
        this.client.on(event, cb)
    }
}

module.exports = TcpClient
