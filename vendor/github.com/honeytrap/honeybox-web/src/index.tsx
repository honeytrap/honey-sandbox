import * as React from 'react';
import { Provider, connect } from 'react-redux';
import { compose, createStore, combineReducers, applyMiddleware } from 'redux';
import {
	HashRouter,
	BrowserRouter,
	Redirect,
	Switch,
	Route
} from 'react-router-dom';
import {
	syncHistoryWithStore,
	ConnectedRouter,
	routerReducer,
	routerMiddleware,
	push
} from 'react-router-redux'

import createHistory from 'history/createBrowserHistory'

import { render } from 'react-dom';

import '../scss/style.scss';

import promise from 'redux-promise';
import * as moment from 'moment';

import App from './app';

import i18n from './i18n';

import { connectionMiddleware } from './connection/connectionMiddleware';
import { defaultAppState, rootReducer } from './rootReducer';
import { authenticationMiddleware } from './authentication/authenticationMiddleware';
import { I18nextProvider } from "react-i18next";
import './global.scss';

const clientVersion = process.env.CLIENT_VERSION;

console.log('Honeyfarm client version: ', clientVersion);


const history = createHistory();

function configureStore(initialState) {
	return createStore(
		rootReducer,
		initialState,
		applyMiddleware(
			routerMiddleware(history),
			promise,
			connectionMiddleware,
			authenticationMiddleware
		),
	);
}

if (location.host === 'app.cyberborg.de') {
	i18n.changeLanguage('de');
	moment.locale('de');
} else if (location.host === 'app.cyberborg.nl') {
	i18n.changeLanguage('nl');
	moment.locale('nl');
} else {
	i18n.changeLanguage('en');
	moment.locale('en');
}

const store = configureStore(defaultAppState);


const params = new URLSearchParams(history.location.search);
if (params.get('ref'))
	localStorage.setItem('ref', params.get('ref'));

const load = () => render((
	<Provider store={store}>
		<ConnectedRouter history={history}>
			<I18nextProvider i18n={i18n}>
				<App />
			</I18nextProvider>
		</ConnectedRouter>
	</Provider>
), document.getElementById('root'));

load();
