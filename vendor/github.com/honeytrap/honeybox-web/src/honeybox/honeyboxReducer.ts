import {
	Event,
	Info,
	SetMessageFilter,
	StdErr,
	StdOut
} from './honeyboxActions';
import { StdOutInterface } from './interfaces/stdOutInterface';
import { uniqueId } from 'lodash'
import { Message } from './interfaces/message';

export interface HoneyboxState {
	stdOut: StdOutInterface[];
	messages: Message[];
	messageFilter: string;
	name: string;
}

export const defaultHoneyboxState: HoneyboxState = {
	stdOut: [],
	messages: [],
	messageFilter: '',
	name: ''
};

function b64DecodeUnicode(str) {
	// Going backwards: from bytestream, to percent-encoding, to original string.
	return decodeURIComponent(atob(str).split('').map(function(c) {
		return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
	}).join(''));
}

export function honeyboxReducer(state: HoneyboxState = defaultHoneyboxState, action): HoneyboxState {
	switch (action.type) {
		case StdOut.type: {
			const stdOut: StdOutInterface = {
				id: uniqueId(),
				data: b64DecodeUnicode(action.payload.data).replace(/\r?\n/g, '\r\n'),
				type: 'normal'
			};

			return {
				...state,
				stdOut: state.stdOut.concat([stdOut])
			};
		}

		case StdErr.type: {
			const stdOut: StdOutInterface = {
				id: uniqueId(),
				data: b64DecodeUnicode(action.payload.data).replace(/\r?\n/g, '\r\n'),
				type: 'error'
			};

			return {
				...state,
				stdOut: state.stdOut.concat([stdOut])
			};
		}

		case Event.type: {
			const message: Message = {
				text: action.payload.text
			};

			return {
				...state,
				messages: [message].concat(state.messages)
			};
		}

		case SetMessageFilter.type: {
			return {
				...state,
				messageFilter: action.payload.query
			};
		}

		case Info.type: {
			return {
				...state,
				name: action.payload.name
			};
		}

		default: {
			return state;
		}
	}
}
