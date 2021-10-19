let a, b;

function setup() {
	size(400, 400, 'rgba(0,0,0,0.8)')

	const canvas = document.getElementById('canvas')
	canvas.addEventListener('mousemove', (e) => {
		a = createVector(e.x, e.y)
	})

	b = createVector(200, 60)
}

function scalarProjection(a,b) {
	const bCopy = b.copy().normalize();
	const sp = a.dot(bCopy);
	bCopy.mult(sp);
	return bCopy;
}

function draw() {
	setup()
	const ctx = document.getElementById('canvas').getContext('2d');
	const pos = createVector(100, 200)
	ctx.strokeStyle = "white"

	
	line(pos.x, pos.y, pos.x + a.x, pos.y + a.y)
	line(pos.x, pos.y, pos.x + b.x, pos.y + b.y)

	circle(pos.x, pos.y, 8)

	const v = scalarProjection(a, b)
	ctx.strokeStyle = "blue"
	line(pos.x, pos.y, pos.x + v.x, pos.y + v.y)

	ctx.strokeStyle = "green"
	line(pos.x + a.x, pos.y + a.y, v.x + pos.x, v.y + pos.y);

	circle(v.x + pos.x, v.y + pos.y, 8)

	circle(pos.x + a.x, pos.y + a.y, 8)

}

function main() {
	setup()
	setInterval(draw, 10)
}

main()

