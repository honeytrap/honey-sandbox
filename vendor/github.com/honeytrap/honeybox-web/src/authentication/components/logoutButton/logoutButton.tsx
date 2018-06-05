import * as React from 'react';
import { connect, Dispatch } from 'react-redux';
import { NavLink } from 'reactstrap';
import { Logout } from '../../authenticationActions';

interface Props {
	dispatch: Dispatch<any>;
}

class LogoutButton extends React.Component<Props> {
	logout = () => {
		const { dispatch } = this.props;

		dispatch(new Logout());
	};

	render() {
		return (
			<NavLink href="#" onClick={this.logout}>Log out</NavLink>
		)
	}
}

export default connect()(LogoutButton);