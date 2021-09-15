const width = 800;
const height = 500;
const color = 'rgba(255, 0, 0, 0.5)'
let n = 0;
let c = 6;

function setup() {
	size(width, height, color)
}

function draw() {
	rerender(color)

	let angle = n * 137.3 //137.5;
	let r = c * Math.sqrt(n);
	let x = r * Math.cos(angle) + width/2;
	let y = r * Math.sin(angle) + height/2;

	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.moveTo(x, y)
	ctx.fillStyle = `hsl(${n % 256}, 255, 255)`
	ctx.ellipse(x, y, 4, 4, 0, 0, 2 * Math.PI)
	ctx.fill()
	ctx.save()
	ctx.restore()

	n++
}

function main() {
	setup()
	setInterval(draw, 10)
}

main()