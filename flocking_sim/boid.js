class Boid {
	constructor() {
		this.position = createVector(Math.floor(Math.random() * width), Math.floor(Math.random() * height));
		this.velocity = Vector.random2D();
		this.velocity.setMag(Math.random() * 4);
		this.acceleration = createVector();
		this.maxForce = 0.2;
		this.maxSpeed = 4;
	}

	edges() {
		if (this.position.x > width) {
			this.position.x = 0;
		} else if (this.position.x < 0) {
			this.position.x = width;
		}

		if (this.position.y > height) {
			this.position.y = 0;
		} else if (this.position.y < 0) {
			this.position.y = height;
		}
	}

	align(boids) {
		let perceptionRadius = 50;
		let steering = createVector();
		let total = 0;
		for (const other of boids) {
			const distance = dist(
				this.position.x, 
				this.position.y, 
				other.position.x, 
				other.position.y
			);
			if (other != this && distance < perceptionRadius) {
				steering.add(other.velocity);
				total++;
			}
		}
		if (total > 0) {
			steering.div(total);
			steering.setMag(this.maxSpeed);
			steering.sub(this.velocity);
			steering.limit(this.maxForce);
			
		}
		return steering 
	}

	cohesion(boids) {
		let perceptionRadius = 50;
		let steering = createVector();
		let total = 0;
		for (const other of boids) {
			const distance = dist(
				this.position.x, 
				this.position.y, 
				other.position.x, 
				other.position.y
			);
			if (other != this && distance < perceptionRadius) {
				steering.add(other.position);
				total++;
			}
		}
		if (total > 0) {
			steering.div(total);
			steering.sub(this.position);
			steering.setMag(this.maxSpeed);
			steering.sub(this.velocity);
			steering.limit(this.maxForce);
			
		}
		return steering 
	}

	separation(boids) {
		let perceptionRadius = 50;
		let steering = createVector();
		let total = 0;
		for (const other of boids) {
			const distance = dist(
				this.position.x, 
				this.position.y, 
				other.position.x, 
				other.position.y
			);
			if (other != this && distance < perceptionRadius) {
				const difference = Vector.sub(this.position, other.position);
				difference.div(difference);
				steering.add(difference);
				total++;
			}
		}
		if (total > 0) {
			steering.div(total);
			steering.setMag(this.maxSpeed);
			steering.sub(this.velocity);
			steering.limit(this.maxForce);
			
		}
		return steering 
	}

	flock(boids) {
		const alignment = this.align(boids)
		const cohesion = this.cohesion(boids)
		const separation = this.separation(boids);
		this.acceleration.add(separation);
		this.acceleration.add(alignment);
		this.acceleration.add(cohesion);
	}

	update() {
		this.position.add(this.velocity);
		this.velocity.add(this.acceleration);
		this.velocity.limit(this.maxSpeed);
		this.acceleration.mult(0);
	}

	show() {
		const ctx = document.getElementById('canvas').getContext('2d');
		ctx.beginPath();
		ctx.lineWidth = 4;
		ctx.strokeStyle = 'coral';
		ctx.ellipse(this.position.x, this.position.y, 1, 1, 0, 0, Math.PI * 2)
		ctx.stroke();
	}
}