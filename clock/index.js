const width = 400;
const height = 400;

function size(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.fillStyle = color;
	ctx.fillRect(0, 0, width, height);
	return ctx;
}

function render() {
	size(width, height, 'black');
}

function draw() {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.clearRect(0, 0, width, height);
	render()
	const d = new Date()
	const hr = d.getHours()
	const min = d.getMinutes()
	const sec = d.getSeconds()
	ctx.beginPath()
	ctx.fillStyle = 'white'
	ctx.font = '24px helvetica'
	ctx.fillText(`${hr}:${min}:${sec}`, width - 100, height - 20)

	ctx.strokeStyle = 'yellow'
	ctx.lineWidth = 6
	ctx.ellipse(200, 200, 180, 180, 0, 0, 2 * Math.PI);
	ctx.stroke()

}

function main() {
	setInterval(draw, 10);
}
main();