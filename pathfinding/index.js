const height = 400;
const width = 400;
const cols = 15;
const rows = 15;

const grid = new Array(cols);

const openset = [];
const closedset = [];
let path = [];

let start;
let end;
let current;

function dist(x1, y1, x2, y2) {
	// eucledian distance
	// return Math.sqrt((Math.pow(x1-x2,2))+(Math.pow(y1-y2,2)));
	// mahattan distance aka taxi cab
	return Math.abs(x1 - x2) + Math.abs(y1  - y2)
}

function heuristic(a, b) {
	return dist(a.x, a.y, b.x, b.y)
}

function removeFromArray(array, element) {
	const index = array.indexOf(element)
	array.splice(index, 1)
	// Return a copy, doens't mutate the original
	// return array.filter(el => {
	// 	return el != element
	// })
}

function Spot(i, j) {
	this.x = i;
	this.y = j;
	this.f = 0;
	this.g = 0;
	this.h = 0;
	this.previous = undefined;
	this.neighbors = [];
	this.wall = false;

	if (Math.random() < 0.3) {
		this.wall = true;
	}

	this.width = width / cols;
	this.height = height / rows;

	this.show = (color) => {
		const ctx = document.getElementById('canvas').getContext('2d');
		ctx.beginPath();
		ctx.lineWidth = 1;
		ctx.fillStyle = color;
		if (this.wall) {
			ctx.fillStyle = 'black'
		}
		ctx.rect(this.x*this.width, this.y*this.height, this.width-1, this.height-1);
		ctx.fill()
		return ctx;
	}

	this.addNeighbors = (grid) => {
		const x = this.x;
		const y = this.y;
		if (x < cols - 1) 	this.neighbors.push(grid[x + 1][y])
		if (x > 0) 			this.neighbors.push(grid[x - 1][y])
		if (y < rows - 1) 	this.neighbors.push(grid[x    ][y + 1])
		if (y > 0) 			this.neighbors.push(grid[x    ][y - 1])
	}


}

function size(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.fillStyle = color;
	ctx.fillRect(0, 0, width, height);
	return ctx;
}

// return 0 if searching
// return 1 if reached goal
// return -1 if no solution
function setup() {
	size(width, height, 'black');

	for (let i = 0 ; i < cols ; i++) {
		grid[i] = new Array(rows);
	}

	for (let i = 0 ; i < cols ; i++) {
		for (let j = 0 ; j < rows ; j++) {
			grid[i][j] = new Spot(i, j);
		}
	}

	for (let i = 0 ; i < cols ; i++) {
		for (let j = 0 ; j < rows ; j++) {
			grid[i][j].addNeighbors(grid);
		}
	}

	start = grid[0][0];
	end = grid[cols-1][rows-1];
	start.wall = false;
	end.wall = false;

	openset.push(start);
}

function draw() {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.clearRect(0, 0, width, height);
	
	size(width, height, 'black');

	for (let i = 0 ; i < cols ; i++) {
		for (let j = 0 ; j < rows ; j++) {
			grid[i][j].show('white');
		}
	}

	pathfinding()

	for (let i = 0 ; i < closedset.length ; i++) {
		closedset[i].show('coral');
	}

	for (let i = 0 ; i < openset.length ; i++) {
		openset[i].show('lightgreen');
	}

	// find the path -- backtracking
	path = [];
	let tmp = current;
	path.push(tmp);

	while (tmp.previous) {
		path.push(tmp.previous);
		tmp = tmp.previous;
	}


	for (const marks of path) {
		marks.show('purple');
	}
}

function pathfinding() {

	if (openset.length > 0) {
		// keep moving
		let winner = 0;
		for (let i = 0 ; i < openset.length ; i++) {
			if (openset[i].f < openset[winner].f) {
				winner = i;
			}
		}

		current = openset[winner];


		if (current == end) {
			console.log('done');
		}

		removeFromArray(openset, current);
		closedset.push(current);

		const neighbors = current.neighbors;

		for (const neighbor of neighbors) {

			if (!closedset.includes(neighbor) && !neighbor.wall) {

				let tmpG = current.g + heuristic(neighbor, current)

				if (!openset.includes(neighbor)) {
					openset.push(neighbor)
				} else if (tmpG >= neighbor.g) {
					continue;
				}

				neighbor.g = tmpG;
				neighbor.h = heuristic(neighbor, end);
				neighbor.f = neighbor.g + neighbor.h;
				neighbor.previous = current;
			}
		}
	} else {
		// no solution
		console.log("no solution");
	}

}

function main() {
	setup();
	setInterval(draw, 100)
}
main();