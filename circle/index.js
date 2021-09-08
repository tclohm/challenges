const circles = [];
const width = 801;
const height = 801;

function dist(x1, y1, x2, y2) {
	return Math.sqrt((Math.pow(x1-x2,2))+(Math.pow(y1-y2,2)));
}

function size(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.fillStyle = color;
	ctx.fillRect(0, 0, width, height);
	return ctx;
}

function setup() {
	size(width, height, 'black');
	const circle = newCircle();
	if (circle != null) {
		circles.push(circle);
	} 
}

function draw() {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.clearRect(0, 0, width, height);
	setup();
	for (const circle of circles) {

		if (circle.growing) {
			if (circle.edges()) {
				circle.growing = false;
			} else {
				for (let other of circles) {
					if (circle !== other) {
						const d = dist(circle.x, circle.y, other.x, other.y);
						if (d - 3 < circle.r + other.r) {
							circle.growing = false;
							break
						}
					}
				}
			}
		}
		ctx.beginPath();
		circle.show();
		circle.grow();
	}
}

function newCircle() {
	const x = Math.random() * width;
	const y = Math.random() * height;

	let valid = true;
	for (const circle of circles) {
		const distance = dist(x, y, circle.x, circle.y)
		if (distance < circle.r) {
			valid = false;
			break;
		}
	}

	if (valid) {
		return new Circle(x, y)
	} else {
		return null;
	}
}

function main() {
	setup();
	setInterval(draw, 1);
}
main();