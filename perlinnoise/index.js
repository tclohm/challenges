function createCanvas(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.fillStyle = color
	ctx.fillRect(0, 0, width, height)
	return ctx
}
let width;
let height
function setup() {
	width = 801;
	height = 801;
	createCanvas(width, height, 'rgba(0, 0, 0, 1)');
}

let xoff = 4;

// MARK: Thanks P5JS! from their docs
const constrain = function(n, low, high) {
  return Math.max(Math.min(n, high), low);
};

// MARK: Thanks P5JS! from their docs
const map = function(n, start1, stop1, start2, stop2, withinBounds) {
  const newval = (n - start1) / (stop1 - start1) * (stop2 - start2) + start2;
  if (!withinBounds) {
    return newval;
  }
  if (start2 < stop2) {
    return this.constrain(newval, start2, stop2);
  } else {
    return this.constrain(newval, stop2, start2);
  }
};

function draw() {
	xoff = xoff + 0.02;
	let n = noise(xoff) * width;
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.clearRect(0,0, width, height)
	setup()
	let x = map(noise(xoff), 0, 1, 0, width);
	xoff += 0.02;
	const middleY = height / 2;
	ctx.fillStyle = 'white';
	ctx.beginPath();
	ctx.ellipse(x, middleY, 24, 24, 0, 0, 2 * Math.PI)
	ctx.fill()
}


function main() {
	setup()
	setInterval(draw, 20)
}

main()