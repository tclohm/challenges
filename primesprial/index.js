let width = 500;
let height = 500;
let x = width / 2;
let y = height / 2;


// MARK: -- space between steps
let stepLength = 20;
// MARK: -- steps taken
let step = 1;
// MARK: -- how many steps we need before turning
let numberOfSteps = 1;
// MARK: -- state will help us figure out our direction we need to move
let state = 0;
// MARK: -- everytime we turn we increase the counter
let turnCounter = 0;
let lastStep = (width / stepLength) * (height / stepLength);

function setup() {
	size(width, height, 'rgba(0,0,0,0.8)');
}

function draw() {
	// MARK: -- if we reach the last step, stop
	if (step == lastStep) {
		return
	}

	const ctx = document.getElementById('canvas').getContext('2d');

	const random = randomColor();

	//text(step, x, y, ctx, random);
	dot(x, y, ctx, random);

	// MARK: -- move our dot / text depending on where state is
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

	// MARK: -- if step / numberOfSteps remainder is 0
	//				we change our state, making sure it stays with our 4 cases
	//				increase our turn counter
	//				
	//				every two turns increase the number of steps

	if (step % numberOfSteps == 0) {
		state = (state + 1) % 4;
		turnCounter++;

		if (turnCounter % 2 == 0) {
			numberOfSteps++;
		}
	}

	step++;
}

function main() {
	setup()
	setInterval(draw, 100)
}

main()

