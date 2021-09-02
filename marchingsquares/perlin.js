'use strict';AbortController
const perlin = {
	random_vector: () => {
		const theta = Math.random() * 2 * Math.PI;
		return { x: Math.cos(theta), y: Math.sin(theta) }
	},
	dot_product_grid: (x, y, vx, vy) => {
		let g_vector;
		const d_vector = { x: x - vx, y: y - vy }
		if (this.gradients[[vx, vy]]) {
			g_vector = this.gradients[[vx, vy]];
		} else {
			g_vector = this.random_vector();
			this.gradients[[vx, vy]] = g_vector;
		}
		return d_vector.x * g_vector.x + d_vector.y * g_vector.y;
	},
	smootherstep: (x) => {
		return 6 * x**5 - 15 * x**4 + 10 * x**3;
	},
	interpolation: (x, a, b) => {
		return a + this.smootherstep(x) * (b - a)
	},
	seed: () => {
		this.gradients = {};
		this.memory = {};
	},
	get: (x, y) => {
		if (this.memory.hasOwnProperty([x, y])) return this.memory[[x, y]]
		const xf = Math.floor(x);
		const yf = Math.floor(y);
		// interpolate
		const tr = this.dot_product_grid(x, y, xf+1, yf);
		const tl = this.dot_product_grid(x, y, xf, 	 yf);
		const bl = this.dot_product_grid(x, y, xf, 	 yf+1);
		const br = this.dot_product_grid(x, y, xf+1, yf+1);

		const xt = this.interpolation(x-xf, tl, tr);
		const xb = this.interpolation(x-xf, bl, br);
		const v = this.interpolation(y-yf, xt, xb);
		this.memory[[x,y]] = v;
		return v;
	}
}
perlin.seed();