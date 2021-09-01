function createCanvas(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.fillStyle = color
	ctx.fillRect(0, 0, width, height)
	return ctx
}

let field;
const rez = 20;
let cols, rows;

function makeField(cols, rows) {
	const arr = []
	for (let i = 0 ; i < rows ; i++) {
		const row = new Array(cols)
		arr.push(row)
		for (let j = 0 ; j < cols ; j++) {
			const square = new Square(i, j);
			arr[i][j] = square
		}
	}
	return arr
}

function setup() {
	const width = 801;
	const height = 801;
	createCanvas(width, height, 'rgba(0, 0, 0, 1)');
	cols = Math.floor(width / rez);
	rows = Math.floor(height / rez);
	field = makeField(cols, rows);
	draw()
}

function Square(x, y) {
	this.x = x
	this.y = y

	this.show = function(number) {
		const ctx = document.getElementById('canvas').getContext('2d');
		if (number > 0.5) {
			ctx.fillStyle = `rgba(0, 0, 0, ${number})`
			ctx.fillRect(x*rez, y*rez, 4, 4)
		} else {
			ctx.clearRect(x*rez, y*rez, 4, 4)
		}
		return ctx
	}
}

function draw() {
	for (let i = 0 ; i < rows ; i++) {
		for (let j = 0 ; j < cols ; j++) {
			let random = Math.random()
			field[i][j].show(random)
		}
	}
}


function main() {
	setup()
	setInterval(draw, 100)
}

main()