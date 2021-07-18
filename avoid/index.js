let cols, rows;
let spacing = 5;
let grid;

function createMatrix(cols, rows) {
	const arr = new Array(cols)
	for (let i = 0; i < arr.length; i++) {
		arr[i] = new Array(rows)
	}

}

function createCanvas(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.fillStyle = color
	ctx.fillRect(0, 0, width, height)
	return ctx
}

function setup() {
	const height = 400
	const width = 400
	createCanvas(height, width, 'rgba(220, 220, 220, 0.5)')
	cols = Math.floor(height/spacing)
	rows = Math.floor(width/spacing)

	
}

function draw() {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.strokeStyle = 'white'
	ctx.lineWidth = spacing * 0.5
	ctx.fillStyle = 'rgba(255,255,255, 0.5)'
	ctx.beginPath();
  	ctx.arc(x * spacing, y * spacing, 1, 0, 2 * Math.PI, true);
  	ctx.stroke();

  	const rand = Math.floor(Math.random() * 4)
  	switch (rand) {
  		case 0:
  			x += 1
  			break
  		case 1:
  			x -= 1
  			break
  		case 2:
  			y += 1
  			break
  		case 3:
  			y -= 1
  			break
  	}
}

function main() {
	setup()
	setInterval(draw, 100)
}

main()