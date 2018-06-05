import * as React from 'react';
import { Badge, DropdownItem, DropdownMenu, DropdownToggle, Nav, NavItem, NavLink, Container } from 'reactstrap';
import { AppAsideToggler, AppHeaderDropdown, AppNavbarBrand, AppSidebarToggler } from '@coreui/react';
import logo from './bee.png';
import {Link} from 'react-router-dom';
import LogoutButton from '../../../authentication/components/logoutButton/logoutButton';
import { connect } from 'react-redux';
import { AppState } from '../../../rootReducer';
import { User } from '../../../authentication/interfaces/user';

interface Props {
	user: User;
}

class Header extends React.Component<Props> {
	render() {
		const { user } = this.props;

		return (
			<React.Fragment>
				<Container fluid>
					<AppSidebarToggler className="d-lg-none" display="md" mobile />
					<AppNavbarBrand
						full={{ src: logo, width: 55, height: 35, alt: 'Honeyfarm' }}
						minimized={{ src: logo, width: 30, height: 30, alt: 'Honeyfarm' }}
					/>

					<Nav className="d-md-down-none" navbar>
						<NavItem className="px-3">
							<Link to="/" className="nav-link">Dashboard</Link>
						</NavItem>
						<NavItem className="px-3">
							<NavLink href="/kibana/app/kibana">Kibana</NavLink>
						</NavItem>
						<NavItem className="px-3">
							<NavLink href="https://demo.marija.io/">Exploration</NavLink>
						</NavItem>
					</Nav>
					<Nav className="ml-auto" navbar>
						<NavItem className="px-3">
							<Link to="/account" className="nav-link">
								<i className="icon-user" />
								{user.name}
							</Link>
						</NavItem>
						<NavItem className="d-md-down-none">
							<LogoutButton/>
						</NavItem>
					</Nav>
				</Container>
			</React.Fragment>
		);
	}
}

const select = (state: AppState) => ({
	user: state.authentication.user
});

export default connect(select)(Header);