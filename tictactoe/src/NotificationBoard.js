import React from 'react';

const container = {
	padding: '1rem',
	display: "flex",
	flexDirection: "column",
	width: "8rem"
}

const winnerstyle = {
	background: 'aqua',
	display: 'flex',
	justifyContent: 'center'
}

export default ({ winner, xIsNext, reset }) => {

	return (
		<div style={container}>
			{winner ? <span style={winnerstyle}>Winner: {winner}</span> : "Next Player: " + (xIsNext ? "X" : "O")}
			<button onClick={reset} >Reset</button>
		</div>
	)
};