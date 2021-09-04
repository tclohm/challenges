let width;
let height;
let increment = 0.01;
let start = 0;

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

function createCanvas(width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.fillStyle = color
	ctx.fillRect(0, 0, width, height)
	return ctx
}

function setup() {
	width = 801;
	height = 801;
	createCanvas(width, height, 'rgba(255, 255, 255, 1)');
}

function draw() {
	const ctx = document.getElementById('canvas').getContext('2d');
	let xoff = start;

	ctx.clearRect(0,0, width, height);

	let gradient = ctx.createLinearGradient(50, 100, 120, 0);
	gradient.addColorStop("0", "coral");
	gradient.addColorStop("0.5" ,"pink");
	gradient.addColorStop("1.0", "lightblue");

	setup();

	ctx.beginPath();

	for (let x = 0 ; x < width ; x++) {
		
		const y = noise(xoff) * (height / 2)
		ctx.lineWidth = 15
		ctx.strokeStyle = gradient;

		ctx.lineTo(x, y);
		xoff += increment;
	}
	ctx.stroke();
	start += increment;
}


function main() {
	setup()
	setInterval(draw, 10)
}

main()