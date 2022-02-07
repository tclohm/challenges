import React from 'react';

export default ({ data }) => {


	return (
		<table style={{ "borderStyle": "solid", "borderWidth": "1px", "borderColor": "green" }}>
			<thead>
			<tr>
				<td>Last Name</td>
				<td>First Name</td>
				<td>Phone Number</td>
			</tr>
			</thead>
			<tbody>
				{data.map((obj, key) => (
					<tr key={key}>
						<td>{obj.userFirstName}</td>
						<td>{obj.userLastName}</td>
						<td>{obj.userPhoneNumber}</td>
					</tr>
				))}
			</tbody>
		</table>
	)
	
}