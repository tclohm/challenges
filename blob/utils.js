function randomRange(min, max) {
	return Math.floor(Math.random() * (max - min + 1) + min)
}

function size(x, y, width, height, color) {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.beginPath()
	ctx.fillStyle = color;
	ctx.fillRect(x, y, width, height);
	return ctx;
}

const line = function(x1, y1, x2, y2) {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.beginPath();
	ctx.lineWidth = 4
	ctx.moveTo(x1, y1);
	ctx.lineTo(x2, y2);
	ctx.stroke()
	return ctx;
}

const circle = function(x, y, radius) {
	const ctx = document.getElementById('canvas').getContext('2d');
	ctx.beginPath();
	ctx.fillStyle = 'white'
	ctx.strokeStyle = 'black'
	ctx.lineWidth = 4
	ctx.ellipse(x, y, radius, radius, 0, 0, Math.PI * 2)
	ctx.fill()
	ctx.stroke()
	return ctx;
}

const constrain = function(n, low, high) {
  return Math.max(Math.min(n, high), low);
};

function translate(x, y) {
	const ctx = document.getElementById('canvas').getContext('2d')
	if (x instanceof Vector) {
		y = x.y
		x = x.x
	}
	ctx.translate(x, y)
}

function createVector(x, y, z) {
	return new Vector(x,y,z);
}


function Vector(x, y, z) {
	this.x = x || 0;
	this.y = y || 0;
	this.z = z || 0;
}

Vector.mult = function mult(v, n, target) {
	if (!target) {
		target = v.copy();
	} else {
		target.set(v);
	}
	target.mult(n);
	return target
}

Vector.prototype.copy = function copy() {
	return new Vector(this.x, this.y, this.z);
}

Vector.prototype.set = function set(x, y, z) {
	this.x = x.x || 0;
    this.y = x.y || 0;
    this.z = x.z || 0;
    return this;
}

Vector.prototype.normalize = function normalize() {
	const len = this.mag();
	const x = (1 / len)
	if (len !== 0) this.mult(x);
	return this;
}

Vector.prototype.magSq = function magSq() {
	const x = this.x;
	const y = this.y;
	const z = this.z;
	return x * x + y * y + z * z;
};

Vector.prototype.setMag = function setMag(n) {
	return this.normalize().mult(n);
}

Vector.prototype.mag = function mag() {
	return Math.sqrt(this.magSq());
}

Vector.prototype.mult = function mult(x, y, z) {
	if (x instanceof Vector) {
		if ( 
			Number.isFinite(x.x) &&
			Number.isFinite(x.y) &&
			Number.isFinite(x.z) &&
			typeof x.x === 'number' &&
			typeof x.y === 'number' &&
			typeof x.z === 'number'
		) {
			this.x *= x.x;
			this.y *= x.y;
			this.z *= x.z;
		} else {
			console.warn('Vector.prototype.mult: x contains components that either undefined or not finite numbers')
		}
		return this
	}

	const vectorComponents = [...arguments];
	if (
		vectorComponents.every(element => Number.isFinite(element)) &&
		vectorComponents.every(element => typeof element === 'number')
	) {
		if (arguments.length === 1) {
			this.x *= x;
			this.y *= x;
			this.z *= x;
		}
		if (arguments.length === 2) {
			this.x *= x;
			this.y *= y;
		}
		if (arguments.length === 3) {
			this.x *= x;
			this.y *= y;
			this.z *= z;
		}
  	} else {
	    console.warn(
	      'Vector.prototype.mult:',
	      'x, y, or z arguments are either undefined or not a finite number'
	    );
  	}
  	return this;
}

Vector.sub = function sub(v1, v2, target) {
	if (!target) {
    target = v1.copy();
    if (arguments.length === 3) {
     console.warn(
        'The target parameter is undefined, it should be of type Vector',
        'Vector.sub'
      );
    }
  } else {
    target.set(v1);
  }
  target.sub(v2);
  return target;
}

Vector.prototype.sub = function sub(x, y, z) {
	if (x instanceof Vector) {
		this.x -= x.x || 0;
	    this.y -= x.y || 0;
	    this.z -= x.z || 0;
	    return this;
	}
	this.x -= x || 0;
    this.y -= y || 0;
    this.z -= z || 0;
    return this;
}

Vector.random2D = function random2D() {
  return this.fromAngle(Math.random() * Math.PI * 2);
};

Vector.fromAngle = function fromAngle(angle, length) {
  if (typeof length === 'undefined') {
    length = 1;
  }
  return new Vector(length * Math.cos(angle), length * Math.sin(angle), 0);
};

Vector.prototype.add = function add(x, y, z) {
	if (x instanceof Vector) {
    this.x += x.x || 0;
    this.y += x.y || 0;
    this.z += x.z || 0;
    return this;
  }
  if (x instanceof Array) {
    this.x += x[0] || 0;
    this.y += x[1] || 0;
    this.z += x[2] || 0;
    return this;
  }
  this.x += x || 0;
  this.y += y || 0;
  this.z += z || 0;
  return this;
}

Vector.add = function add(v1, v2, target) {
  if (!target) {
    target = v1.copy();
    if (arguments.length === 3) {
      console.warn(
        'The target parameter is undefined, it should be of type Vector',
        'Vector.add'
      );
    }
  } else {
    target.set(v1);
  }
  target.add(v2);
  return target;
};

Vector.prototype.dot = function dot(x, y, z) {
  if (x instanceof Vector) {
    return this.dot(x.x, x.y, x.z);
  }
  return this.x * (x || 0) + this.y * (y || 0) + this.z * (z || 0);
};


Vector.prototype.limit = function limit(max) {
	const mSq = this.magSq();
  if (mSq > max * max) {
    this.div(Math.sqrt(mSq)) //normalize it
      .mult(max);
  }
  return this;
}