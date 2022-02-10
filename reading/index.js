const fs = require('fs');
const { EventEmitter } = require('events');

// Push
class Reader extends EventEmitter {
	constructor(filename, chunksize = 1024) {
		this.filename = filename;
		this.chunksize = chunksize;
		this.offset = 0;
	}

	read() {
		fs.read(this.filename, this.offset, this.chunksize, (err, chunk) => {
			if (err) {
				this.emit('error', err);
			} else {
				this.emit('data', chunk);
				if (chunk !== null) {
					this.offset += chunk.length;
					this.read();
				}
			}
		});
	}
}

const reader = new Reader('/somefile');
reader.on('data', chunk => asyncHandle(chunk));
reader.read()


// Pull for reading
class AnotherReader {
	constructor(filename, chunksize = 1024) {
		this.filename = filename;
		this.chunksize = chunksize;
		this.offset = 0;
	}

	read() {
		return new Promise((resolve, reject) => {
			fs.read(this,filename, this.offset, this.chunksize, (err, chunk) => {
				if (err) {
					return reject(err);
				}
				if (chunk !== null) {
					this.offset += chunk.length;
				}
				resolve(chunk);
			});
		});
	}
}

const newReader = new AnotherReader('/somefile');
while (true) {
	const chunk = await reader.read();
	await asyncHandle(chunk)
}