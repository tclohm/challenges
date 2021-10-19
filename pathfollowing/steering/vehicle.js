function findProjection(pos, a, b) {
  let v1 = p5.Vector.sub(a, pos);
  let v2 = p5.Vector.sub(b, pos);
  v2.normalize();
  let sp = v1.dot(v2);
  v2.mult(sp);
  v2.add(pos);
  return v2;
}

class Vehicle {
  constructor(x, y) {
    this.pos = createVector(x, y);
    this.vel = createVector(0, 0);
    this.acc = createVector(0, 0);
    this.maxSpeed = 6;
    this.maxForce = 0.1;
    this.r = 16;
  }

  follow(path) {
    // Path following algorithm here!!

    // Step 1 calculate future position
    const ctx = document.getElementById('canvas').getContext('2d');
    let future = this.vel.copy();
    future.mult(20);
    future.add(this.pos);
    ctx.fillStyle = 'red'
    ctx.strokeStyle = 'clear'
    circle(future.x, future.y, 16);

    let target = findProjection(path.start, future, path.end);
    ctx.fillStyle = 'green'
    ctx.strokeStyle = 'clear'
    circle(target.x, target.y, 16);

    let d = p5.Vector.dist(future, target);
    if (d > path.radius) {
      return this.seek(target);
    } else {
      return createVector(0, 0);
    }
  }

  seek(target, arrival = false) {
    let force = p5.Vector.sub(target, this.pos);
    let desiredSpeed = this.maxSpeed;
    if (arrival) {
      let slowRadius = 100;
      let distance = force.mag();
      if (distance < slowRadius) {
        desiredSpeed = map(distance, 0, slowRadius, 0, this.maxSpeed);
      }
    }
    force.setMag(desiredSpeed);
    force.sub(this.vel);
    force.limit(this.maxForce);
    return force;
  }

  applyForce(force) {
    this.acc.add(force);
  }

  update() {
    this.vel.add(this.acc);
    this.vel.limit(this.maxSpeed);
    this.pos.add(this.vel);
    this.acc.set(0, 0);
  }

  show() {
    const ctx = document.getElementById('canvas').getContext('2d');
    ctx.strokeStyle = 'white'
    ctx.lineWidth = '2';
    ctx.stroke();
    ctx.fillStyle = 'white'
    ctx.fill()
    ctx.translate(this.pos.x, this.pos.y);
    ctx.rotate(this.vel.heading());
    ctx.triangle(-this.r, -this.r / 2, -this.r, this.r / 2, this.r, 0);
  }

  edges() {
    if (this.pos.x > width + this.r) {
      this.pos.x = -this.r;
    } else if (this.pos.x < -this.r) {
      this.pos.x = width + this.r;
    }
    if (this.pos.y > height + this.r) {
      this.pos.y = -this.r;
    } else if (this.pos.y < -this.r) {
      this.pos.y = height + this.r;
    }
  }
}