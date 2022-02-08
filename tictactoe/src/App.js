import React, { useState, useEffect } from 'react';
import Board from './Board';
import NotificationBoard from './NotificationBoard';

export default () => {

	const winners = squares => {
		const lines = [
			[0, 1, 2],
			[3, 4, 5],
			[6, 7, 8],
			[0, 3, 6],
			[1, 4, 7],
			[2, 5, 8],
			[0, 4, 8],
			[2, 4, 6]
		];

		for (let i = 0 ; i < lines.length ; i++) {
			const [a, b, c] = lines[i];
			if (squares[a] && squares[a] === squares[b] && squares[a] === squares[c]) {
				return squares[a];
			}
		}
		return null;
	}

	const [board, setBoard] = useState(Array(9).fill(null));
	const [xIsNext, setXisNext] = useState(true);

	const winner = winners(board);

	const handleClick = i => {
		const boardCopy = [...board]
		if (winner || boardCopy[i]) return;
		boardCopy[i] = xIsNext ? "X" : "O";
		setBoard(boardCopy)
		setXisNext(!xIsNext);
	}

	const reset = () => {
		setBoard(Array(9).fill(null))
		setXisNext(true)
	}

	return (
		<>
			<NotificationBoard winner={winner} xIsNext={xIsNext} reset={reset} />
			<Board squares={board} onClick={handleClick} />
		</>
	)
};