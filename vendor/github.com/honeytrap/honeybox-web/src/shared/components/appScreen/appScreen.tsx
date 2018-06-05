import * as React from 'react';
import navigation from './_nav';
import {
	AppAside,
	AppBreadcrumb,
	AppFooter,
	AppHeader,
	AppSidebar,
	AppSidebarFooter,
	AppSidebarForm,
	AppSidebarHeader,
	AppSidebarMinimizer,
	AppSidebarNav
} from '@coreui/react';
import {I18n, translate} from 'react-i18next';
import { Container } from 'reactstrap';
import Footer from '../footer/footer';
import Aside from '../aside/aside';
import Header from '../header/header';

interface Props {
	children;
}

export function AppScreen(props: Props) {
	return (
		<I18n ns="translations">
			{
				(t, {i18n}) => (
					<div className="app">
						<AppHeader fixed>
							<Header />
						</AppHeader>
						<div className="app-body">
							<AppSidebar fixed display="lg">
								<AppSidebarHeader />
								<AppSidebarForm />
								<AppSidebarNav navConfig={navigation} />
								<AppSidebarFooter />
								<AppSidebarMinimizer />
							</AppSidebar>
							<main className="main">
								<ol className="breadcrumb" />
								<Container fluid>
									<div className="animated fadeIn">
										{props.children}
									</div>
								</Container>
							</main>
							<AppAside fixed hidden>
								<Aside />
							</AppAside>
						</div>
						<AppFooter>
							<Footer />
						</AppFooter>
					</div>
				)
			}
		</I18n>
	);
}