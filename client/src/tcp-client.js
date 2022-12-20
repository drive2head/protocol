class TcpClient {
    interval = null

    constructor(client) {
        this.client = client
    }

    /** нерабочая фича */
    // async ping() {
    //     try {
    //         await this.write('ping')
    //     } catch (e) {
    //         console.log('trying to reconnect')
    //         await this.connect()
    //         console.log(e)
    //     }
    // }

    async connect(port, host, cb) {
        this.interval = null
        try {
            await this.client.connect(port, host, cb);
            /** поэтому закомментил */
            // this.interval = setInterval(this.ping, 5000)
        } catch (e) {
            this.interval = null
            console.log(e)
        }
    }

    async write(message) {
        try {
            await this.client.write(message);
        } catch (e) {
            console.log(e)
        }
    }

    async destroy() {
        try {
            await this.client.destroy()
            this.interval = null
        } catch (e) {
            console.log(e)
        }
    }

    async on(event, cb) {
        try {
            await this.client.on(event, cb)
        } catch (e) {
            console.log(e)
        }
    }
}

module.exports = TcpClient
