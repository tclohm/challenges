function make2DArray(cols, rows) {
	let array = new Array(cols);
	for (let i = 0 ; i < array.length ; i++) {
		array[i] = new Array(rows);
	}
	return array
}

function Cell(x, y, w) {
	this.mine = true;
	this.revealed = true;
	this.x = x;
	this.y = y;
	this.w = w;

	if (Math.random(1) < 0.5) {
		this.mine = true;
	} else {
		this.mind = false;
	}

	this.show = () => {
		createRect(this.x, this.y, this.w, this.w)
		if (this.revealed) {
			if (this.mine) {
				ctx.arc(this.x + this.w/2, this.y + this.w/2, this.w * 0.5)
			}
		}
	}

	this.contains = (x, y) => {
		return ( x > this.x  && x < this.x + this.w && y > this.y && y < this.y + this.w )
	}

	this.reveal = () => {
		this.revealed = true
	}
}

let grid;

let cols; 
let rows;

function setup() {
	createCanvas(201, 201);
	cols = Math.floor(width / w);
	rows = Math.floor(height / w);
	grid = make2DArray(cols, row);

	for (let i = 0 ; i < cols ; i++) {
		for (let j = 0 ; j < rows ; j++) {
			grid[i][j] = new Cell(i * w, j * w, w);
		}
	}

}

function draw() {
	for (let i = 0 ; i < cols ; i++) {
		for (let j = 0 ; j < rows ; j++) {
			grid[i][j].show();
		}
	}
}


function mouseUp() {
	for ( var i = 0 ; i < cols ; i++ ) {
		for ( var j = 0 ; j < rows ; j++ ) {
			if (grid[i][j].contains(mouseX, mouseY)) {
				grid[i][j].reveal()
			}
		}
	}
}