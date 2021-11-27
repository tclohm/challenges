function Blob(x, y, r) {
	this.pos = createVector(x,y)
	this.r = r;

	this.show = function () {
		circle(this.pos.x, this.pos.y, this.r)
	}

	this.update = function () {
		let velocity = createVector(mouseX-width/2, mouseY-height/2)
		velocity.setMag(1)
		this.pos.add(velocity)
	}

	this.eats = function (other) {
		const d = Vector.dist(this.pos, other.pos)
		if (d < this.r + other.r) {
			const sum = Math.PI * this.r * this.r + Math.PI * other.r * other.r
			this.r = Math.sqrt(sum / Math.PI)
			return true
		}
		return false
	}
}