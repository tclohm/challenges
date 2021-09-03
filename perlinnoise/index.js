function createCanvas(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.fillStyle = color
	ctx.fillRect(0, 0, width, height)
	return ctx
}

function setup() {
	const width = 801;
	const height = 801;
	createCanvas(width, height, 'rgba(0, 0, 0, 1)');
}

function draw() {
	const noise = perlin
	console.log(noise)
}


function main() {
	setup()
	draw()
}

main()