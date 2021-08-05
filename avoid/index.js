// right, left, down, up
const allOptions = [
	{ dx:  1,  dy: 0  }, 
	{ dx: -1,  dy: 0  },
	{ dx:  0,  dy: 1  },
	{ dx:  0,  dy: -1 }];
let x, y;
let cols, rows;
let spacing = 30;
let grid;

function createMatrix(cols, rows) {
	let arr = new Array(cols);
	for (let i = 0; i < arr.length; i++) {
		arr[i] = new Array(rows);
	}
	return arr;
}

function createCanvas(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.fillStyle = color;
	ctx.fillRect(0, 0, width, height);
	return ctx;
}

function setup() {
	const height = 800;
	const width = 800;
	createCanvas(height, width, 'rgba(100, 100, 100, 1)');
	cols = Math.floor(height/spacing);
	rows = Math.floor(width/spacing);
	x = cols / 2;
	y = rows / 2;
	grid = createMatrix(cols, rows);
	for ( let i = 0 ; i < cols ; i++ ) {
		for ( let j = 0 ; j < rows ; j++ ) {
			grid[i][j] = false
		}
	}

	grid[x][y] = true
}

function isValid(i, j) {
	if (i < 0 || i >= cols || j < 0 || j >= rows) {
		return false
	}
	return !grid[i][j]
}

function draw() {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.strokeStyle = 'rgba(255,255,255, 0.5)';
	ctx.lineWidth = spacing * 0.5;
	ctx.fillStyle = 'rgba(255,255,255, 0.5)';
	ctx.beginPath();
  	ctx.arc(x * spacing, y * spacing, 1, 0, 2 * Math.PI, true);
  	ctx.stroke();

  	const options = allOptions.filter(option => {
  		if (isValid(x + option.dx, y + option.dy)) {
  			return option
  		}
  	})

  	if (options.length > 0) {
  		const random = Math.floor(Math.random() * options.length)
	  	const step = options[random]

	  	x += step.dx
	  	y += step.dy
	  	grid[x][y] = true
  	} else {
  		console.log("STOP!")
  		return
  	}
}

function main() {
	setup();
	setInterval(draw, 100);
}

main();