import React, { useState } from 'react';

const squarestyle = {
	height: '6rem',
	width: '6rem',
	background: 'lightgray',
	margin: '0.1rem',
}

const activestyle = {
	height: '6rem',
	width: '6rem',
	background: 'rgba(0, 0, 0, 0.4)',
	margin: '0.1rem',
	display: 'flex',
	justifyContent: 'center',
}

const shapestyle = {
	display: 'absolute',
	marginLeft: '1.3rem',
	fontSize: '5rem'
}

export default ({ value, onClick }) => {

	return (
		<>
			<div
			onClick={onClick}
			style={squarestyle}
			>
			<span style={shapestyle}>{value}</span>
			</div>
		</>
	)
};