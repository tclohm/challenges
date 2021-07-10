let cols, rows;
let w = 24;
let grid = [];

let current;

const stack = [];

function createCanvas(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.fillStyle = color
	ctx.fillRect(0, 0, width, height)
	return ctx
}

function setup() {
	createCanvas(800, 800, 'rgba(220, 220, 220, 0.5)')
	cols = Math.floor(800/w)
	rows = Math.floor(800/w)

	for (let j = 0 ; j < rows ; j++) {
		for (let i = 0 ; i < cols ; i++) {
			let cell = new Cell(i, j)
			grid.push(cell)
		}
	}

	current = grid[0]
}

function draw() {
	current.visited = true

	for (let i = 0 ; i < grid.length ; i++) {
		grid[i].show()
	}

	current.highlight()

	let next = current.checkNeighbors()
	if (next) {
		next.visited = true

		stack.push(current)
		removeWalls(current, next)
		current = next

	} else if (stack.length > 0) {

		current = stack.pop()
		current.backtracked = true

	}

}

function index(i, j) {
	if (i < 0 || j < 0 || i > cols-1 || j > rows-1) {
		return -1
	}
	return i + j * cols
}

function Cell(i, j) {
	this.i = i
	this.j = j
	// top, right, bottom, left
	this.walls = [true, true, true, true]
	this.visited = false
	this.backtracked = false

	this.checkNeighbors = function () {
		const neighbors = []
		const top 	= grid[index(i, j-1)]
		const right 	= grid[index(i+1, j)]
		const bottom 	= grid[index(i, j+1)]
		const left 	= grid[index(i-1, j)]

		if (top && !top.visited) {
			neighbors.push(top)
		}

		if (right && !right.visited) {
			neighbors.push(right)
		}

		if (bottom && !bottom.visited) {
			neighbors.push(bottom)
		}

		if (left && !left.visited) {
			neighbors.push(left)
		}

		if (neighbors.length > 0) {
			let r = Math.floor(Math.random() * neighbors.length)
			return neighbors[r]
		} else {
			return undefined
		}
	}

	this.show = function () {
		let x = this.i*w
		let y = this.j*w
		const ctx = document.getElementById('canvas').getContext('2d')

		if (this.walls[0]) {
			line(x  , 	y, x+w,	y)
		}
		
		if (this.walls[1]) {
			line(x+w, 	y, x+w,	y+w)
		}
		
		if (this.walls[2]) {
			line(x+w, y+w, 	 x,	y+w)
		}

		if (this.walls[3]) {
			line(x  , y+w, 	 x,	y)
		}

		if (this.visited) {
			ctx.fillStyle = 'rgba(225, 102, 204, 0.5)'
			ctx.fillRect(x, y, w, w)
		}

		if (this.backtracked) {
			ctx.fillStyle = 'rgba(128, 0, 128, 0.5)'
			ctx.fillRect(x, y, w, w)
		}

		return ctx
	}

	this.highlight = function () {
		const x = this.i*w
		const y = this.j*w
		const ctx = document.getElementById('canvas').getContext('2d')
		ctx.fillStyle = 'yellow'
		ctx.fillRect(x, y, w, w)
		return ctx
	}

}

function line(startX, startY, endX, endY) {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.strokeStyle = 'white'
	ctx.lineWidth = 1
	ctx.fillStyle = 'rgba(0,0,200,0)'
	ctx.beginPath()
	ctx.moveTo(startX, startY)
	ctx.lineTo(endX, endY)
	ctx.stroke()
}

function removeWalls(a, b) {
	const x = a.i - b.i
	if (x === 1) {
		a.walls[3] = false
		b.walls[1] = false
	} else if (x === -1) {
		a.walls[1] = false
		b.walls[3] = false
	}
	const y = a.j - b.j
	if (y === 1) {
		a.walls[0] = false
		b.walls[2] = false
	} else if (y === -1) {
		a.walls[2] = false
		b.walls[0] = false
	}
}

function main() {
	setup()
	setInterval(draw, 100)
}

main()