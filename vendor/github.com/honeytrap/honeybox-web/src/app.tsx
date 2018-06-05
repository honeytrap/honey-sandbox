import * as React from 'react';
import { Switch, Redirect, Route } from 'react-router-dom';
import {I18n, translate} from 'react-i18next';
import { withRouter } from 'react-router-dom';
import { authenticatedRoutes, notAuthenticatedRoutes, sharedRoutes } from './routes';
import AuthenticatedRoute from './authentication/components/authenticatedRoute/authenticatedRoute';
import NotAuthenticatedRoute from './authentication/components/notAuthenticatedRoute/notAuthenticatedRoute';
import { connect, Dispatch } from 'react-redux';
import { ResumeSession } from './authentication/authenticationActions';
import { Connect } from './connection/connectionActions';

interface Props {
	dispatch: Dispatch<any>;
}

class App extends React.Component<Props> {
	componentWillMount() {
		const { dispatch } = this.props;

		dispatch(new Connect());

		const token = localStorage.getItem('jwtToken');

		if (token) {
			dispatch(new ResumeSession({
				jwtToken: token
			}));
		}
	}

    render() {
        return (
			<Switch>
				{authenticatedRoutes.map(route =>
					<AuthenticatedRoute path={route.path} component={route.component} exact={route.exact} name={route.name} key={route.path} />
				)}
				{notAuthenticatedRoutes.map(route =>
					<NotAuthenticatedRoute path={route.path} component={route.component} exact={route.exact} name={route.name} key={route.path} />
				)}
				{sharedRoutes.map(route =>
					<Route path={route.path} component={route.component} exact={route.exact} name={route.name} key={route.path} />
				)}
			</Switch>
        );
    }
}

export default withRouter(connect()(App))