class Circle {
	constructor(x_, y_) {
		this.x = x_;
		this.y = y_;
		this.r = 40;
	}

	show() {
		const ctx = document.getElementById('canvas').getContext('2d');
		ctx.beginPath();
		ctx.strokeStyle = 'white';
		ctx.lineWidth = 2;
		ctx.ellipse(this.x, this.y, this.r, this.r, 0, 0, 2 * Math.PI);
		ctx.stroke();
	}

	grow() {
		this.r = this.r + 1;
	}
}