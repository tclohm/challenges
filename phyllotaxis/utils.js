function size(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.fillStyle = color;
	ctx.fillRect(0, 0, width, height);
	return ctx;
}

function rerender(color) {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.clearRect(0,0, width, height)
	size(width, height, color);
}