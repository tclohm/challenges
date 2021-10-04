const width = 400;
const height = 400;

function map(n, start1, stop1, start2, stop2, withinBounds) {
	const newval = (n - start1) / (stop1 - start1) * (stop2 - start2) + start2;
	if (!withinBounds) {
		return newval;
	}
	if (start2 < stop2) {
		return constrain(newval, start2, stop2);
	} else {
		return constrain(newval, stop2, start2);
	}
}

function constrain(n, low, high) {
	return Math.max(Math.min(n, high), low);
}

function lerp(start, stop, amt) {
	return amt * (stop - start) + start
}

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

	ctx.strokeStyle = 'lightcoral'
	ctx.lineWidth = 12
	//n, start1, stop1, start2, stop2, withinBounds
	const seconds = map(sec, 0, 60, 0, Math.PI * 2)

	
	ctx.beginPath()
	ctx.arc(200, 200, 180, 0, seconds)
	ctx.stroke()

	ctx.strokeStyle = 'LawnGreen'
	ctx.lineWidth = 12
	//n, start1, stop1, start2, stop2, withinBounds
	const minutes = map(min, 0, 60, 0, Math.PI * 2)
	
	ctx.beginPath()
	ctx.arc(200, 200, 160, 0, minutes)
	ctx.stroke()

	ctx.strokeStyle = 'LemonChiffon'
	ctx.lineWidth = 12
	//n, start1, stop1, start2, stop2, withinBounds
	const hours = map(hr, 0, 60, 0, Math.PI * 2)
	
	ctx.beginPath()
	ctx.arc(200, 200, 140, 0, hours)
	ctx.stroke()

	window.requestAnimationFrame(draw)

}

function loop() {
	draw()
	window.requestAnimationFrame(loop)
}

window.requestAnimationFrame(loop)