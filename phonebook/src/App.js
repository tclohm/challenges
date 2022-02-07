import React, { useState } from 'react';
import Form from './Form';
import Entries from './Entries';

export default () => {

	const [phonebook, setPhonebook] = useState([{
		userFirstName: "Taylor",
		userLastName: "Lohman",
		userPhoneNumber: "132456789",
	}])

	const addUser = (user) => {
		setPhonebook([...phonebook, user].sort((person1, person2) => {
			if (person1.userLastName === person2.userLastName) {
				return person1.userFirstName > person2.userFirstName
			}
			return person1.userLastName > person2.userLastName
		}));
	}

	return (
		<>
			<Form addUser={addUser} />
			<Entries data={phonebook} />
		</>
	)
}