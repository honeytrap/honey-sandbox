import * as React from 'react';
import { Redirect, Route } from 'react-router-dom';
import { connect } from 'react-redux';
import { AppState } from '../../../rootReducer';

interface Props {
	authenticated: boolean;
	path: string;
	exact: boolean;
	component: React.Component;
	name: string;
}

class AuthenticatedRoute extends React.Component<Props> {
	render() {
		const { authenticated, ...rest } = this.props;

		if (!authenticated) {
			return <Redirect to="/login" />;
		}

		return <Route {...rest} />
	}
}

const select = (state: AppState) => ({
	authenticated: state.authentication.authenticated
});

export default connect(select)(AuthenticatedRoute);