import {
	CloseError,
	Connected,
	ErrorAction,
	LoadingDone
} from './connectionActions';

export interface ConnectionState {
    errors: string;
    connected: boolean;
    loading: boolean;
}

export const defaultConnectionState: ConnectionState = {
    connected: false,
    errors: null,
    loading: true
};

export function connectionReducer(state: ConnectionState = defaultConnectionState, action): ConnectionState {
    switch (action.type) {
        case Connected.type: {
            return {
                ...state,
                connected: action.payload.connected,
                errors: action.payload.errors
            };
        }

        case ErrorAction.type: {
            return {
                ...state,
                errors: action.payload.errors
            }
        }

        case CloseError.type: {
            return {
                ...state,
                errors: undefined
            };
        }

        case LoadingDone.type: {
            return {
                ...state,
                loading: false
            };
        }

        default: {
            return state;
        }
    }
}