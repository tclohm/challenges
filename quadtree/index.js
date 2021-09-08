const width = 800;
const height = 800;
function setup() {
	size(width, height, 'rgba(0,0,0,0.8)');
	let boundary = new Rectangle(400, 400, 200, 200);
	let qt = new QuadTree(boundary, 4);

	for (let i = 0 ; i < 500 ; i++) {
		const p = new Point(Math.floor(Math.random() * width), Math.floor(Math.random() * height))
		qt.insert(p)
	}
	
	qt.show()
}

setup()