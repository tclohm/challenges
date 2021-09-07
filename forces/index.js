let attractors = [];
let particles = [];

function setup() {
	size(800, 400, 'rgba(0, 0, 0, 0.9');

	for (let i = 0 ; i < 20 ; i++) {
		const randomX = Math.abs(Math.floor(Math.random() * 800))
		const randomY = Math.abs(Math.floor(Math.random() * 400))
		particles.push(new Particle(randomX, randomY))
	}

	for (let i = 0 ; i < 1 ; i++) {
		//const randomX = Math.abs(Math.floor(Math.random() * 800))
		// const randomY = Math.abs(Math.floor(Math.random() * 400))
		attractors.push(createVector(400, 200))
	}
}

function draw() { 
	const ctx = document.getElementById('canvas').getContext('2d');
	
	// ctx.clearRect(0,0, 800, 400)
	// size(800, 400, 'rgba(0, 0, 0, 0.8');

	for (const attractor of attractors) {
		ctx.beginPath();
		ctx.fillStyle = 'coral';
		ctx.strokeStyle = 'coral';
		ctx.ellipse(attractor.x, attractor.y, 4, 4, 0, 0, 2 * Math.PI);
		ctx.fill()
		ctx.stroke()
	}

	for (const particle of particles) {
		for (const attractor of attractors) {
			particle.attracted(attractor);
		}
		particle.update();
		particle.show();
	}
}

function main() {
	setup()
	setInterval(draw, 10)
}

main();