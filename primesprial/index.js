let a, b;

function setup() {
	size(400, 400, 'rgba(0,0,0,0.8)')
	const canvas = document.getElementById('canvas');
}

function draw() {
	setup()
	const ctx = document.getElementById('canvas').getContext('2d');
}

function main() {
	setup()
	setInterval(draw, 10)
}

main()

