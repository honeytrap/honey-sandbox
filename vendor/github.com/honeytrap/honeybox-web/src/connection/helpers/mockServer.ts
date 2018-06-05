interface Response {
	willRespond: Function;
	type: string;
	[property: string]: any;
}

export const responses: Response[] = [

];

export function getResponse(action) {
	const response = responses.find(response => response.willRespond(action));

	if (!response) {
		return null;
	}

	const responseCopy = {
		...response
	};

	delete responseCopy.willRespond;

	console.log('Mocking server response: ', responseCopy);

	return {
		data: JSON.stringify(responseCopy)
	};
}