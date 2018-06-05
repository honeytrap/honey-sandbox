import ReconnectingWebsocket from 'reconnecting-websocket';
import { Dispatch, Middleware } from 'redux';
import { Connect, Connected } from './connectionActions';
import { getResponse } from './helpers/mockServer';
import { Event, Info, StdErr, StdOut } from '../honeybox/honeyboxActions';

let opened: Promise<ReconnectingWebsocket>;

export const connectionMiddleware: Middleware = ({dispatch}) => next => (action: any) => {
    switch (action.type) {
		case Connect.type: {
			const url = getUrl();
			console.log('Connecting to backend on ', url);

			opened = new Promise(resolve => {
				if (process.env.MOCK_SERVER) {
					return;
				}

				const socket = new ReconnectingWebsocket(url);

				socket.onopen = (event): any => { onOpen(dispatch); resolve(socket) };
				socket.onmessage = (event): any => onMessage(event, dispatch);
				socket.onclose = (event): any => onClose(event, dispatch);
				socket.onerror = (event): any => onClose(event, dispatch);
			});

			break;
		}
	}

	if (action.sendToServer === true) {
            const stringified: string = JSON.stringify(action);

            action = { ...action };
            delete action.sendToServer;

            if (process.env.NODE_ENV !== 'production') {
                console.log('Send', action);
            }

            if (process.env.MOCK_SERVER) {
				const response = getResponse(action);

				if (response) {
					onMessage(response, dispatch);
				}
			} else {
				opened.then(socket => socket.send(stringified));
			}
    }

    return next(action);
};

function onMessage(event, dispatch: Dispatch<any>) {
	const data = JSON.parse(event.data);

	const actionsFromServer = [
		StdOut,
		Event,
		StdErr,
		Info
	];

	const className = actionsFromServer.find(action => action.type === data.type);

	if (className) {
		const action = new (className as any)(data.payload);

		return dispatch(action);
	}

	console.error('Received an unknown action type from the server: ', data);
}

function onOpen(dispatch: Dispatch<any>) {
    console.log('Succesfully connected');
    dispatch(new Connected({ connected: true, errors: null }));
    return;
}

function onClose(event, dispatch: Dispatch<any>) {
    let reason = getCloseMessage(event);
    console.error('Disconnected from socket:', reason);

    dispatch(new Connected({ connected: false, errors: reason }));
}

function getUrl(): string {
    return process.env.SERVER_URI;
}

function getCloseMessage(event: CloseEvent): string {
    switch (event.code) {
        case 1000:
            return 'Normal closure, meaning that the purpose for which the connection was established has been fulfilled.';
        case 1001:
            return 'An endpoint is \'going away\', such as a server going down or a browser having navigated away from a page.';
        case 1002:
            return 'An endpoint is terminating the connection due to a protocol error';
        case 1003:
            return 'An endpoint is terminating the connection because it has received a type of data it cannot accept (e.g., an endpoint that understands only text data MAY send this if it receives a binary message).';
        case 1004:
            return 'Reserved. The specific meaning might be defined in the future.';
        case 1005:
            return 'No status code was actually present.';
        case 1006:
            return 'The connection was closed abnormally, e.g., without sending or receiving a Close control frame';
        case 1007:
            return 'An endpoint is terminating the connection because it has received data within a message that was not consistent with the type of the message (e.g., non-UTF-8 [http://tools.ietf.org/html/rfc3629] data within a text message).';
        case 1008:
            return 'An endpoint is terminating the connection because it has received a message that \'violates its policy\'. This reason is given either if there is no other sutible reason, or if there is a need to hide specific details about the policy.';
        case 1009:
            return 'An endpoint is terminating the connection because it has received a message that is too big for it to process.';
        case 1010: // Note that this status code is not used by the server, because it can fail the WebSocket handshake instead:
            return 'An endpoint (client) is terminating the connection because it has expected the server to negotiate one or more extension, but the server didn\'t return them in the response message of the WebSocket handshake. <br /> Specifically, the extensions that are needed are: ' + event.reason;
        case 1011:
            return 'A server is terminating the connection because it encountered an unexpected condition that prevented it from fulfilling the request.';
        case 1015:
            return 'The connection was closed due to a failure to perform a TLS handshake (e.g., the server certificate can\'t be verified).';
        default:
            return 'Unknown reason';
    }
}