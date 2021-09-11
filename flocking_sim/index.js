const flock = [];
const height = 360;
const width = 640;
const color = "rgba(50, 50, 255, 0.8)";

function setup() {
	size(width, height, color);
	for (let i = 0 ; i < 100 ; i++) {
		flock.push(new Boid());
	}
	
}

function draw() {
	rerender(color);
	for (let boid of flock) {
		boid.edges();
		boid.flock(flock);
		boid.update();
		boid.show();
	}
}

function main() {
	setup();
	setInterval(draw, 10)
}

main();