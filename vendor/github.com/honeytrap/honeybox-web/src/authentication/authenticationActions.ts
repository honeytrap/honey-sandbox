import { createServerAction } from '../shared/helpers/createServerAction';
import { createClientAction } from '../shared/helpers/createClientAction';
import { User } from './interfaces/user';

export const Login = createServerAction<{ username: string, password: string }>('authentication.login');
export const LoginSuccess = createClientAction<{ jwtToken: string, user: User }>('authentication.loginSuccess');
export const LoginFailed = createClientAction('authentication.loginFailed');
export const ResumeSession = createServerAction<{ jwtToken: string }>('authentication.resumeSession');
export const ResumeSessionSuccess = createClientAction<{ jwtToken: string, user: User }>('authentication.resumeSessionSuccess');
export const ResumeSessionFailed = createClientAction('authentication.resumeSessionFailed');
export const SignUp = createServerAction<{ name: string, password: string, company: string, email: string }>('authentication.signUp');
export const SignUpSuccess = createClientAction<{ jwtToken: string, user: User }>('authentication.signUpSuccess');
export const SignUpFailed = createClientAction<{ errors: any }>('authentication.signUpFailed');
export const Logout = createClientAction('authentication.logout');
