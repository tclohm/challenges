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

	if (Math.random(1) < 0.5) {
		this.mine = true;
	} else {
		this.mine = false;
	}

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
				ctx.fillStyle = 'rgba(200, 200, 200, 1)';
				ctx.fillRect(this.x, this.y, this.w, this.w);
			}
		} else {
			ctx.strokeStyle = 'rgba(200, 200, 200, 1)';
			ctx.strokeRect(this.x, this.y, this.w, this.w);
		}
	}

	this.contains = (x, y) => {
		if ( (x > this.x && x < this.x + this.w) && (y > this.y && y < this.y + this.w) ) {
			console.log("contains y", y, "<", this.y, "<", this.y + w);
			console.log("contains x", x, "<", this.x, "<", this.x + w);
			return true
		}
		return false
	}

	this.reveal = () => {
		this.revealed = true
	}

	this.countNeighbors = () => {
		if (this.mine) {
			return -1;
		}
		let total = 0;

		for (let i = -1 ; i <= 1 ; i++) {
			for (let j = -1 ; j <= 1 ; j++) {
				let neighbor = grid[this.i + i][this.j + j];
				if (neighbor.mine) {
					total++;
				}
			}
		}
		this.neighborCount = total;
	}
}

let grid;

let cols; 
let rows;
const w = 40;

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


function mouseUp(mouseX, mouseY) {
	for ( var i = 0 ; i < cols ; i++ ) {
		for ( var j = 0 ; j < rows ; j++ ) {
			if (grid[i][j].contains(mouseX, mouseY)) {
				grid[i][j].reveal();
			}
		}
	}
}


function main() {
	setup()
	setInterval(draw, 100)
}

main()