const state = {
    OPEN_SESSION: 'OPEN_SESSION',
    PING: 'PING',
    CLOSE_SESSION: 'CLOSE_SESSION'
}

class TcpClient {
    interval = null
    pingTimeout = null

    constructor(client, port, host, cb) {
        this.client = client
        this.port = port
        this.host = host
        this.cb = cb

        this.connect = this.connect.bind(this)
        this.ping = this.ping.bind(this)
        this.write = this.write.bind(this)
        this.destroy = this.destroy.bind(this)
        this.on = this.on.bind(this)
        this.close = this.on.bind(this)
        this.onOk = this.onOk.bind(this)
    }

    async connect() {
        this.interval = null
        try {
            await this.client.connect(this.port, this.host, this.cb);
            await this.write(state.OPEN_SESSION)
            this.interval = setInterval(this.ping, 1000)
        } catch (e) {
            this.interval = null
            console.log(e)
        }
    }

    /** нерабочая фича */
    async ping() {
        console.log('PING')
        try {
            await this.write(state.PING)
            this.pingTimeout = setTimeout(this.close, 5000)
        } catch (e) {
            console.log(e)
            await this.destroy()
        }
    }

    async write(message) {
        try {
            await this.client.write(message);
        } catch (e) {
            console.log(e)
            await this.destroy()
        }
    }

    async close(markups) {
        try {
            await this.client.write(markups);
            await this.write(state.CLOSE_SESSION)
            await this.client.destroy()
        } catch (e) {
            console.log(e)
            await this.destroy()
        }
    }

    async destroy() {
        clearInterval(this.interval)
        this.interval = null
        this.pingTimeout = null
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
            // console.log(e)
            await this.destroy()
        }
    }

    async onOk() {
        console.log('OK')
        this.pingTimeout = null
    }
}

module.exports = TcpClient
