function createCanvas(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.fillStyle = color
	ctx.fillRect(0, 0, width, height)
	return ctx
}

const puck = new Puck(400, 400);
const width = 401;
const height = 401;

function setup() {
	createCanvas(width, height, 'rgba(0, 0, 0, 1)');
}

function draw() {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.clearRect(0, 0, 800, 800);
	setup()
	puck.update()
	puck.edges(height, width)
	puck.show()
}

function main() {
	setup()
	setInterval(draw, 20)
}

main()