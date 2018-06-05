import {
	connectionReducer,
	ConnectionState, defaultConnectionState
} from './connection/connectionReducer';
import {
	authenticationReducer,
	AuthenticationState, defaultAuthenticationState
} from './authentication/authenticationReducer';
import { combineReducers } from 'redux';
import {
	defaultHoneyboxState,
	honeyboxReducer,
	HoneyboxState
} from './honeybox/honeyboxReducer';

export const rootReducer = combineReducers({
	sessions: connectionReducer,
	authentication: authenticationReducer,
	connection: connectionReducer,
	honeyfarm: honeyboxReducer
});

export interface AppState {
	authentication: AuthenticationState;
	connection: ConnectionState;
	honeyfarm: HoneyboxState;
}

export const defaultAppState: AppState = {
	authentication: defaultAuthenticationState,
	connection: defaultConnectionState,
	honeyfarm: defaultHoneyboxState
};