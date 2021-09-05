const circles = [];
const width = 801;
const height = 801;
let circle = new Circle(200, 200)

function size(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.fillStyle = color
	ctx.fillRect(0, 0, width, height)
}

function setup() {
	size(width, height, 'black');
	//circles.push(new Circle(200, 200));
}

function draw() {
	const ctx = document.getElementById('canvas').getContext('2d');
	setup();
	// MARK: -- begin path fixed it
	ctx.beginPath();
	circle.show()
	circle.grow()
	// for (const circle of circles) {
	// 	circle.show()
	// 	circle.grow()
	// }
}

function main() {
	setup();
	setInterval(draw, 100);
}
main();