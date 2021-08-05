// right, left, down, up
const allOptions = [
	{ dx:  1,  dy:  0, tried: false }, 
	{ dx: -1,  dy:  0, tried: false },
	{ dx:  0,  dy:  1, tried: false },
	{ dx:  0,  dy: -1, tried: false }
];

class Spot {
	constructor(i, j) {
		this.i = i;
		this.j = j;
		this.x = i * spacing;
		this.y = j * spacing;
		this.options = [...allOptions]
		this.visited = false;
	}

	step() {
		const options = allOptions.filter(option => {
  			if (isValid(this.i + option.dx, this.j + option.dy)) {
  				return option
  			}
  		})

	  	if (options.length > 0) {
	  		const random = Math.floor(Math.random() * options.length)
		  	const step = options[random]
		  	return grid[this.i + step.dx][this.j + step.dy]
  		} else {
  			return undefined
  		}
	}
}