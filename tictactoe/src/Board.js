import React, { useState } from 'react';
import Square from './Square';

const boardstyle = {
	"display": "grid",
	"margin": "0 auto",
	"height": "20rem",
	"width": "20rem",
	"gridTemplate": "repeat(3, 1fr) / repeat(3, 1fr)"
}

export default ({ squares, onClick }) => {
	
	return (
		<div style={boardstyle}>
			{squares.map((value, i) => (
				<Square key={i} value={value} onClick={() => onClick(i)} />
			))}
		</div>
	)
}