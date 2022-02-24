let width = 500;
let height = 500;
let x = width / 2;
let y = height / 2;

let stepLength = 20;
let step = 1;
let stepNumber = 1;
let state = 0;
let turnCounter = 0;
let lastStep = (width / stepLength) * (height / stepLength);

function setup() {
	size(width, height, 'rgba(0,0,0,0.8)');
}

function draw() {

	if (step == lastStep) {
		return
	}

	const ctx = document.getElementById('canvas').getContext('2d');

	const random = randomColor();

	//text(step, x, y, ctx, random);
	dot(x, y, ctx, random);

	switch (state) {
		case 0:
			x += stepLength;
			break;
		case 1:
			y -= stepLength;
			break
		case 2:
			x -= stepLength;
			break
		case 3:
			y += stepLength;
			break
	}

	if (step % stepNumber == 0) {
		state = (state + 1) % 4;
		turnCounter++;

		if (turnCounter % 2 == 0) {
			stepNumber++;
		}
	}

	step++;
}

function main() {
	setup()
	setInterval(draw, 10)
}

main()

