class Paddle {
	constructor(x) {
		this.x = x;
		this.width = 10;
		this.height = 80;
		this.y = height/2;
		this.ychange = 0;
	}

	show() {
		const ctx = document.getElementById('canvas').getContext('2d');
		ctx.fillStyle = 'rgba(225, 255, 255, 1)';
		ctx.beginPath()
		ctx.rect(this.x, this.y, this.width, this.height)
		ctx.fill()
	}

	update() {
		this.y += this.ychange
		if (this.y > 310.5) {
			this.y = 309.5
		}
		if (this.y < 10.5) {
			this.y = 11.5
		}
	}

	move(steps) {
		this.ychange = steps;
	}
}