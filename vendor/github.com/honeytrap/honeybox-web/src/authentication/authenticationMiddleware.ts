import { push } from 'react-router-redux';
import { createMiddleware } from '../shared/helpers/createMiddleware';
import {
	LoginSuccess,
	Logout,
	ResumeSessionSuccess, SignUpFailed, SignUpSuccess
} from './authenticationActions';
import { actions } from 'react-redux-form';

export const authenticationMiddleware = createMiddleware([
	{
		actionType: LoginSuccess.type,
		afterHandler: (storeAPI, action) => {
			window.localStorage.setItem('jwtToken', action.jwtToken);
			push('/');
		}
	},
	{
		actionType: ResumeSessionSuccess.type,
		afterHandler: (storeAPI, action) => {
			window.localStorage.setItem('jwtToken', action.jwtToken);
			push('/');
		}
	},
	{
		actionType: SignUpSuccess.type,
		afterHandler: (storeAPI, action) => {
			window.localStorage.setItem('jwtToken', action.jwtToken);
			push('/');
		}
	},
	{
		actionType: SignUpFailed.type,
		afterHandler: ({ dispatch, getState }, action) => {
			console.log(action.errors);

			dispatch(actions.setFieldsErrors('user', action.errors))
		}
	},
	{
		actionType: Logout.type,
		afterHandler: (storeAPI, action) => {
			window.localStorage.removeItem('jwtToken');
			push('/login');
		}
	},
]);