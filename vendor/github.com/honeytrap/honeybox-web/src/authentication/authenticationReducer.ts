import {
	LoginSuccess,
	Logout, ResumeSessionFailed,
	ResumeSessionSuccess, SignUpFailed, SignUpSuccess
} from './authenticationActions';
import { User } from './interfaces/user';

export interface AuthenticationState {
	authenticated: boolean;
	jwtToken: string;
	signUpErrors: any;
	user: User;
}

export const defaultAuthenticationState: AuthenticationState = {
	authenticated: false,
	jwtToken: null,
	signUpErrors: {},
	user: null
};

export function authenticationReducer(state: AuthenticationState = defaultAuthenticationState, action): AuthenticationState {
	switch (action.type) {
		case LoginSuccess.type:
		case SignUpSuccess.type:
		case ResumeSessionSuccess.type: {
			return {
				...state,
				signUpErrors: {},
				authenticated: true,
				jwtToken: action.payload.jwtToken,
				user: action.payload.user
			};
		}

		case Logout.type: {
			return {
				...state,
				authenticated: false,
				jwtToken: null
			};
		}

		case SignUpFailed.type: {
			return {
				...state,
				signUpErrors: action.payload.errors
			};
		}

		default: {
			return state;
		}
	}
}