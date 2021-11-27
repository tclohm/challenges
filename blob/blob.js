function Blob(x, y, r) {
	this.pos = createVector(x,y)
	this.r = r;

	this.show = function () {
		circle(this.pos.x, this.pos.y, this.r)
	}

	this.update = function () {
		let velocity = createVector(mouseX-width/2, mouseY-height/2)
		velocity.setMag(3)
		this.pos.add(velocity)
	}
}