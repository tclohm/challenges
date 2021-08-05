let cols, rows;
let spacing = 30;
let grid;
const path = [];
let spot;

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
	grid = createMatrix(cols, rows);
	for ( let i = 0 ; i < cols ; i++ ) {
		for ( let j = 0 ; j < rows ; j++ ) {
			grid[i][j] = new Spot(i, j)
		}
	}
	spot = grid[0][0]
}

function isValid(i, j) {
	if (i < 0 || i >= cols || j < 0 || j >= rows) {
		return false
	}
	return !grid[i][j].visited;
}

function draw() {
	path.push(spot)
	spot.visited = true

  	console.log("spot", spot)
  	spot = spot.step()

  	if (!spot) {
  		console.log("stuck!")
  	}

  	const ctx = document.getElementById('canvas').getContext('2d');
  	ctx.strokeStyle = 'rgba(255,255,255, 0.5)';
	ctx.lineWidth = spacing * 0.5;
	ctx.fillStyle = 'rgba(255,255,255, 0.5)';
	ctx.beginPath();
	ctx.arc(spot.x * spacing, spot.y * spacing, 1, 0, 2 * Math.PI, true);
	ctx.stroke();

  	for (let spot of path) {
		// working on it
  	}
}

function main() {
	setup();
	setInterval(draw, 1000);
}

main();