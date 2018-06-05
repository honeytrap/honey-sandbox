import { createSelector } from 'reselect';
import { AppState } from '../rootReducer';
import { Message } from './interfaces/message';

export const selectFilteredMessages = createSelector(
	(state: AppState) => state.honeyfarm.messages,
	(state: AppState) => state.honeyfarm.messageFilter,

	(messages: Message[], messageFilter: string): Message[] => {
		if (messageFilter.length === 0) {
			return messages;
		}

		return messages.filter(message =>
			message.text.toLowerCase().includes(messageFilter.toLowerCase())
		)
	}

);