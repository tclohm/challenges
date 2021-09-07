function Particle(x, y) {
	this.pos = createVector(x, y);
	this.prev = createVector()
	this.vel = Vector.random2D();
	this.acc = createVector();

	this.update = function() {
		this.pos.add(this.vel);
		this.vel.add(this.acc);
		this.acc.mult(0);
	}

	this.show = function() {
		const ctx = document.getElementById('canvas').getContext('2d');
		//ctx.beginPath()
		ctx.lineWidth = 2;
		//ctx.fillStyle = 'yellow';
		//ctx.ellipse(this.pos.x, this.pos.y, 4, 4, 0, 0, 2 * Math.PI);
		//ctx.fill()
		
		if (this.prev.x, this.prev.y) {
			ctx.strokeStyle = 'rgba(255, 255, 255, 0.2)'
			ctx.moveTo(this.pos.x, this.pos.y)
			ctx.lineTo(this.prev.x, this.prev.y)
			ctx.stroke()
		}


		this.prev.x = this.pos.x;
		this.prev.y = this.pos.y;
	}

	this.attracted = function(target) {
		let force = Vector.sub(target, this.pos);
		//console.log("force", force)
		let dsquared = force.magSq();
		dsquared = constrain(dsquared, 25, 500);
		//console.log("dsqr", dsquared)
		let G = 6.67408;
		let strength = G / dsquared;
		force.setMag(strength);
		// console.log("setmag", force)
		this.acc.add(force);
	}
}