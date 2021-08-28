function createCanvas(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.fillStyle = color
	ctx.fillRect(0, 0, width, height)
	return ctx
}

function make2DArray(cols, rows) {
	let array = new Array(cols);
	for (let i = 0 ; i < array.length ; i++) {
		array[i] = new Array(rows);
	}
	return array
}

function Cell(i, j, w) {
	this.i = i;
	this.j = j;
	this.mine = false;
	this.revealed = false;
	this.x = i * w;
	this.y = j * w;
	this.w = w;
	this.neighborCount = 0

	this.show = () => {

		const ctx = document.getElementById('canvas').getContext('2d');

		if (this.revealed) {
			if (this.mine) {
				ctx.fillStyle = 'rgba(225, 23, 148, 1)';
				ctx.beginPath()
				// 			x, 					y, 					radiusX, 	radiusY, 	rotation, startAngle, endAngle
				ctx.ellipse(this.x + this.w/2, this.y + this.w/2, this.w * 0.2, this.w * 0.2, 0, 		0, 			2 * Math.PI)
				ctx.fill()
			} else {
				ctx.fillStyle = 'rgba(119, 200, 225, 0.2)';
				ctx.fillRect(this.x, this.y, this.w, this.w);
				if (this.neighborCount > 0) {
					ctx.fillStyle = 'white'
					ctx.font = '20px Helvetica';
					ctx.fillText(this.neighborCount, this.x + 15, this.y + 28);
				}
			}
		} else {
			ctx.strokeStyle = 'rgba(200, 200, 200, 1)';
			ctx.strokeRect(this.x, this.y, this.w, this.w);
			ctx.save()
		}
	}

	this.contains = (x, y) => {
		if ( (x > this.x && x < this.x + this.w) && (y > this.y && y < this.y + this.w) ) {
			return true
		}
		return false
	}

	this.reveal = () => {
		this.revealed = true;

		const ctx = document.getElementById('canvas').getContext('2d');
		ctx.fillStyle = 'rgba(120, 255, 187, 0.38)';
		ctx.fillRect(this.x, this.y, this.w, this.w);
		ctx.restore()
		setTimeout(() => {
			if (this.neighborCount == 0) {
			// flood fill time
				this.floodFill();
			}
		}, 100)
	}

	this.countNeighbors = () => {
		if (this.mine) {
			this.neighborCount = -1;
			return;
		}
		let total = 0;

		for (let xoff = -1 ; xoff <= 1 ; xoff++) {
			for (let yoff = -1 ; yoff <= 1 ; yoff++) {
				let i = this.i + xoff
				let j = this.j + yoff
				if (i > -1 && i < cols && j > -1 && j < rows) {
					let neighbor = grid[i][j];
					if (neighbor.mine) {
						total++;
					}
				}
			}
		}
		this.neighborCount = total;
	}

	this.floodFill = () => {
		for (let xoff = -1 ; xoff <= 1 ; xoff++) {
			for (let yoff = -1 ; yoff <= 1 ; yoff++) {
				let i = this.i + xoff;
				let j = this.j + yoff;
				if (i > -1 && i < cols && j > -1 && j < rows) {
					let neighbor = grid[i][j];
					if (!neighbor.mine && !neighbor.revealed) {
						neighbor.reveal()
					}
				}
			}
		}
	}
}

let grid;
let cols; 
let rows;
const w = 40;
const totalMines = 4;

function setup() {
	const width = 401;
	const height = 401;
	createCanvas(width, height, 'rgba(255, 255, 255, 0.5)');
	cols = Math.floor(width / w);
	rows = Math.floor(height / w);
	grid = make2DArray(cols, rows);

	for (let i = 0 ; i < cols ; i++) {
		for (let j = 0 ; j < rows ; j++) {
			grid[i][j] = new Cell(i, j, w);
		}
	}

	let n = 0;
	while (n < totalMines) {
		let i = Math.floor(Math.random() * cols);
		let j = Math.floor(Math.random() * rows);
		if (!grid[i][j].mine) {
			grid[i][j].mine = true
			n++
		}
	}

	for (let i = 0 ; i < cols ; i++) {
		for (let j = 0 ; j < rows ; j++) {
			grid[i][j].countNeighbors();
		}
	}


	draw()
}

function draw() {
	for (let i = 0 ; i < cols ; i++) {
		for (let j = 0 ; j < rows ; j++) {
			grid[i][j].show();
		}
	}
}

function gameOver() {
	for ( var i = 0 ; i < cols ; i++ ) {
		for ( var j = 0 ; j < rows ; j++ ) {
			grid[i][j].revealed = true
		}
	}
}


function mouseUp(mouseX, mouseY) {
	for ( var i = 0 ; i < cols ; i++ ) {
		for ( var j = 0 ; j < rows ; j++ ) {
			if (grid[i][j].contains(mouseX, mouseY)) {
				grid[i][j].reveal();
				if (grid[i][j].mine) {
					gameOver();
				}
			}
		}
	}
}


function main() {
	setup()
	setInterval(draw, 100)
}

main()