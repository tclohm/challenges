let cols, rows;
let w = 40;
let grid = [];

function createCanvas(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.fillStyle = color
	ctx.fillRect(0, 0, width, height)
	return ctx
}

function setup() {
	createCanvas(800, 800, 'black')
	cols = Math.floor(800/w)
	rows = Math.floor(800/w)

	for (let j = 0 ; j < rows ; j++) {
		for (let i = 0 ; i < cols ; i++) {
			let cell = new Cell(i, j)
			grid.push(cell)
		}
	}
}

function draw() {
	for (let i = 0 ; i < grid.length ; i++) {
		grid[i].show()
	}
}

function Cell(i, j) {
	this.i = i
	this.j = j

	this.show = function () {
		let x = this.i*w
		let y = this.j*w
		const ctx = document.getElementById('canvas').getContext('2d')

		line(x  , 	y, x+w,	y)
		
		line(x+w, 	y, x+w,	y+w)
		
		line(x+w, y+w, x  ,	y+w)
		
		line(x  , 	y, x+w,	y)

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

function main() {
	setup()
	draw()
}

main()