function createCanvas(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.fillStyle = color
	ctx.fillRect(0, 0, width, height)
	return ctx
}

const puck = new Puck(400, 400);
const width = 401;
const height = 401;

const left = new Paddle(0+10);
const right = new Paddle(width-20)

document.addEventListener('keyup', function(input) {
	if (input.keyCode === 87) {
			// MARK: -- up
			left.move(0)
		}
		if (input.keyCode === 83) {
			// MARK: -- down
			left.move(0)
		}
		if (input.keyCode === 38) {
			// MARK: -- up
			right.move(0)
		}
		if (input.keyCode === 40) {
			// MARK: -- down
			right.move(0)
		}
})

document.addEventListener('keydown', function(input) {
		if (input.keyCode === 87) {
			// MARK: -- up 
			left.move(-5)
		}
		if (input.keyCode === 83) {
			// MARK: -- down
			left.move(5)
		}
		if (input.keyCode === 38) {
			// MARK: -- up
			right.move(-5)
		}
		if (input.keyCode === 40) {
			// MARK: -- down
			right.move(5)
		}
})

function setup() {
	createCanvas(width, height, 'rgba(0, 0, 0, 1)');
}

function draw() {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.clearRect(0, 0, 800, 800);
	setup()

	left.show()
	right.show()

	left.update()
	right.update()

	puck.update()
	puck.edges(height, width)
	puck.show()

	puck.checkPaddleRight(right);
	puck.checkPaddleLeft(left);
}

function main() {
	setup()
	setInterval(draw, 10)
}

main()