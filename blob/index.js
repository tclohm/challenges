let blob;
let mouseX;
let mouseY;
const blobs = [];
let width = 600;
let height = 600;

function setup() {
	const ctx = document.getElementById('canvas').getContext('2d')
	size(0, 0, height, width, 'rgba(0,0,0,1)')
	blob = new Blob(width/2, height/2, 30);
	for (let i = 0 ; i < 100 ; i++) {
		const x = randomRange(-width, width*2)
		const y = randomRange(-height, height*2)
		blobs[i] = new Blob(x, y, 10)
	}
}

function draw() {
	const ctx = document.getElementById('canvas').getContext('2d')
	ctx.beginPath()
	ctx.clearRect(width/2-blob.pos.x, height/2-blob.pos.y, height, width)
	ctx.setTransform(1, 0, 0, 1, 0, 0)
	translate(width/2-blob.pos.x, height/2-blob.pos.y)
	size(0, 0, height, width, 'rgba(0,0,0,1)')
	
	blob.show()
	blob.update()

	
	for (const b of blobs) {
		b.show()
	}

	
}

function main() {
	setup()

	canvas.addEventListener('mousemove', e => {
		mouseX = e.clientX
		mouseY = e.clientY
	})

	setInterval(draw, 5)
}

main()