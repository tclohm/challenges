function size(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.fillStyle = color;
	ctx.fillRect(0, 0, width, height);
	return ctx;
}

function randomColor() {
	let R = (Math.floor(Math.random() * 255)),
    	G = (Math.floor(Math.random() * 255)),
    	B = (Math.floor(Math.random() * 255)),
    color = 'rgb(' + R + ', ' + G + ',' + B + ')';
    return color
}

function text(num, x, y, ctx, color) {
	ctx.font = '10px Helvetica';
	ctx.textAlign = 'center';
	ctx.moveTo(x, y);
	ctx.fillStyle = color;
	ctx.fillText(num, x, y);
	return ctx;
}

function square(x, y, ctx, color) {
	ctx.fillStyle = color;
	ctx.fillRect(x, y, 10, 10);
	return ctx;
}

function dot(x, y, ctx, color) {
	ctx.fillStyle = color;
	ctx.beginPath();
	ctx.arc(x, y, 7, 0, 2 * Math.PI);
	ctx.fill();
	return ctx;
}