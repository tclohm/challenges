class Circle {
	constructor(x_, y_) {
		this.x = x_;
		this.y = y_;
		this.r = 1;
		this.growing = true;
	}

	show() {
		const ctx = document.getElementById('canvas').getContext('2d');
		ctx.strokeStyle = 'white';
		ctx.lineWidth = 2;
		ctx.ellipse(this.x, this.y, this.r, this.r, 0, 0, 2 * Math.PI);
		ctx.stroke();
	}

	grow() {
		if (this.growing) {
			this.r += 1
		}
	}

	edges() {
		return this.x + this.r > width || this.x - this.r < 0 || this.y + this.r > height || this.y - this.r < 0;
	}
}