class Point {
	constructor(x, y) {
		this.x = x;
		this.y = y;
	}
}

class Rectangle {
	constructor(x, y, width, height) {
		this.x = x;
		this.y = y;
		this.width = width;
		this.height = height;
		this.divided = false;
	}

	contains(point) {
		return (point.x > this.x - this.width &&
				point.x < this.x + this.width &&
				point.y > this.y - this.height &&
				point.y < this.y + this.height); 
	}
}

class QuadTree {
	constructor(boundary, capacity) {
		this.boundary = boundary;
		this.capacity = capacity;
		this.points = [];
	}

	subdivide() {
		const x = this.boundary.x;
		const y = this.boundary.y;
		const width = this.boundary.width;
		const height = this.boundary.height;

		const nw = new Rectangle(x - width/2, y - height/2, width/2, height/2);
		this.northwest = new QuadTree(nw, this.capacity);
		const ne = new Rectangle(x + width/2, y - height/2, width/2, height/2);
		this.northeast = new QuadTree(ne, this.capacity);
		const sw = new Rectangle(x - width/2, y + height/2, width/2, height/2);
		this.southwest = new QuadTree(sw, this.capacity);
		const se = new Rectangle(x + width/2, y + height/2, width/2, height/2);
		this.southeast = new QuadTree(se, this.capacity);
		this.divided = true;
	}

	insert(point) {

		if (!this.boundary.contains(point)) {
			return
		}


		if (this.points.length < this.capacity) {
			this.points.push(point);
		} else {
			if (!this.divided) {
				this.subdivide();
			}
			this.northwest.insert(point);
			this.northeast.insert(point);
			this.southwest.insert(point);
			this.southeast.insert(point);
		}
	}

	show() {
		const ctx = document.getElementById('canvas').getContext('2d');
		ctx.beginPath()
		ctx.strokeStyle = 'white';
		ctx.rect(this.boundary.x, this.boundary.y, this.boundary.width*2, this.boundary.height*2);
		ctx.stroke();

		if (this.divided) {
			this.northwest.show()
			this.northeast.show()
			this.southwest.show()
			this.southeast.show()
		}

		for (const p of this.points) {
			ctx.beginPath()
			ctx.fillStyle = 'yellow';
			ctx.ellipse(p.x, p.y, 2, 2, 0, 0, Math.PI * 2)
			ctx.fill()
		}
	}
}