const API_URL = "http://localhost:8080";

async function fetchUsers() {
	try {
		const endpoint = `${API_URL}/users`;
		const headers = {
			'accept': 'application/json'
		};

		const response = await fetch(endpoint, {headers});
		if (!response.ok) {
			throw new Error(`Error fetching users: ${response.status}`);
		}

		const data = await response.json();
		return data;
	} catch(error) {
		console.error(`An error occurred: ${error}`);
		return null;
	}
}

window.addEventListener("DOMContentLoaded", async () => {


	try {
		const data = await fetchUsers();
		let records = '';
		for (let record of data) {
			const user = record['user'];
			const links = record['links'];
			console.log(links);
			let select = `
				<select class="">
					<option selected value="">Choose Leave</option>`;

			Object.entries(links).forEach(([key, value]) => {
				select += `<option value="${value}">${key}</option>`;
			});
			select += '</select>';

			records  += `
				<tr>
					<td>${user["id"]}</td>
					<td>${user["name"]}</td>
					<td>${user["gender"]}</td>
					<td>${select}</td>
				</tr>`;
		}
		document.getElementById("users").innerHTML = records;
	} catch (e) {
		console.error(e);
	}
});
