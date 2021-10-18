let a, b;

function setup() {
	size(400, 400, 'rgba(0,0,0,0.8)')

	a = createVector(100, -60)
	b = createVector(200, 60)
}

function scalarProjection(a,b) {
	const bCopy = b.copy().normalize();
	const sp = a.dot(bCopy);
	bCopy.mult(sp);
	return bCopy;
}

function draw() {
	const ctx = document.getElementById('canvas').getContext('2d');
	const pos = createVector(100, 200)
	ctx.strokeStyle = "white"
	line(pos.x, pos.y, pos.x + a.x, pos.y + a.y)
	line(pos.x, pos.y, pos.x + b.x, pos.y + b.y)

	circle(pos.x, pos.y, 8)

	const v = scalarProjection(a, b)
	ctx.strokeStyle = "blue"
	line(pos.x, pos.y, pos.x + v.x, pos.y + v.y)
}

setup()
draw()