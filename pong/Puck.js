class Puck {
	constructor(width, height) {
		this.x = width / 2;
		this.y = height / 2;
		this.xspeed = 2;
		this.yspeed = 1;
		this.radius = 12;
	}

	checkPaddleLeft(paddle) {
		// ball between the paddle
		if (this.y < paddle.y + paddle.height && 
			this.y > paddle.y - paddle.height && 
			this.x - this.radius < paddle.x + paddle.width/2) {
			this.xspeed *= -1
		}
	}

	checkPaddleRight(paddle) {
		// ball between the paddle
		if (this.y < paddle.y + paddle.height && 
			this.y > paddle.y - paddle.height && 
			this.x + this.radius > paddle.x - paddle.width/2) {
			console.log("true")
			this.xspeed *= -1
		}
	}

	show() {
		//const red = Math.floor(Math.random() * 255)
		//const green = Math.floor(Math.random() * 255)
		//const blue = Math.floor(Math.random() * 255)
		const ctx = document.getElementById('canvas').getContext('2d');
		//ctx.fillStyle = `rgba(${red}, ${green}, ${blue}, 1)`;
		ctx.fillStyle = 'rgba(225, 255, 255, 1)';
		ctx.beginPath()
				// 		x, 		y, 	radiusX, 	radiusY, rotation, startAngle, endAngle
		ctx.ellipse(this.x, this.y, this.radius, 		this.radius, 	 0, 	   0,	2 * Math.PI)
		ctx.fill()
	}

	update() {
		this.x = this.x + this.xspeed;
		this.y = this.y + this.yspeed;
	}

	edges(height, width) {
		const ctx = document.getElementById('canvas').getContext('2d');
		if (this.y < 0 || this.y > height) {
			this.yspeed *= -1
		}

		if (this.x < 0 || this.x > width) {
			this.xspeed *= -1
		}
	}
}