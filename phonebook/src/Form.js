import React, { useState } from 'react';

export default ({ addUser }) => {

	const [user, setUser] = useState({
		userFirstName: "",
		userLastName: "",
		userPhoneNumber: "",
	});

	const onChange = e => {
		e.preventDefault()
		setUser({ ...user, [e.target.name]: e.target.value })
	}

	return (
		<div style={{ "display": "flex", "flex-direction": "column", "justify-content": "flex-start" }}>
			<div style={{ "margin": "0.2rem", "display": "flex", "justify-content": "flex-start" }}>
				<label style={{ "margin-right": "2.0rem"}}>First Name</label>
				<input 
				name="userFirstName"
				type="text"
				onChange={e => onChange(e)} />
			</div>
			<div style={{ "margin": "0.2rem", "display": "flex", "justify-content": "flex-start" }}>
				<label style={{ "margin-right": "2.1rem"}}>Last Name</label>
				<input 
				name="userLastName"
				type="text"
				onChange={e => onChange(e)} />
			</div>
			<div style={{ "margin": "0.2rem", "display": "flex", "justify-content": "flex-start" }}>
				<label style={{ "margin-right": "0.5rem"}}>Phone Number</label>
				<input 
				name="userPhoneNumber"
				type="text"
				onChange={e => onChange(e)} />
			</div>
			<button
				onClick={e => { addUser(user) }}
				style={{ "backgroundColor": "green", "width": "8rem",  "height": "2rem", "borderRadius": "0.2rem", "borderWidth": "0"}}
			>Add User</button>
		</div>
	)
	
}