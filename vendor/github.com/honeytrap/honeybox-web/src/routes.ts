import Sandbox from './honeybox/components/sandbox/sandbox';
import LoginForm from './authentication/components/loginForm/loginForm';
import SignUpForm from './authentication/components/signUpForm/signUpForm';

interface RouteConfig {
	path: string;
	exact: boolean;
	name: string;
	component: any;
}

export const authenticatedRoutes: RouteConfig[] = [
	// {
	// 	path: '/',
	// 	exact: true,
	// 	name: 'Home',
	// 	component: Dashboard
	// },
	// {
	// 	path: '/servers',
	// 	exact: true,
	// 	name: 'Servers',
	// 	component: ServersOverview
	// },
	// {
	// 	path: '/agents',
	// 	exact: true,
	// 	name: 'Agents',
	// 	component: AgentsOverview
	// },
	// {
	// 	path: '/collectors',
	// 	exact: true,
	// 	name: 'Collectors',
	// 	component: CollectorsOverview
	// },
	// {
	// 	path: '/analysers',
	// 	exact: true,
	// 	name: 'Analysers',
	// 	component: AnalysersOverview
	// },
	// {
	// 	path: '/alerts',
	// 	exact: true,
	// 	name: 'alerts',
	// 	component: AlertsOverview
	// },
	// {
	// 	path: '/events',
	// 	exact: true,
	// 	name: 'Events',
	// 	component: EventsOverview
	// },
	// {
	// 	path: '/sensors',
	// 	exact: true,
	// 	name: 'Sensors',
	// 	component: SensorsOverview
	// },
	// {
	// 	path: '/attackers',
	// 	exact: true,
	// 	name: 'Attackers',
	// 	component: AttackersOverview
	// },
];

export const notAuthenticatedRoutes: RouteConfig[] = [
	// {
	// 	path: '/login',
	// 	exact: true,
	// 	name: 'Login',
	// 	component: LoginForm
	// },
	// {
	// 	path: '/sign-up',
	// 	exact: true,
	// 	name: 'Sign up',
	// 	component: SignUpForm
	// },
	// {
	// 	path: '/sandbox',
	// 	exact: true,
	// 	name: 'Sandbox',
	// 	component: Sandbox
	// },
];

export const sharedRoutes: RouteConfig[] = [
	{
		path: '/',
		exact: true,
		name: 'Honeybox',
		component: Sandbox
	}
];