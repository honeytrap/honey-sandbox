import * as React from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { replace, push } from 'react-router-redux';
import { LocalForm, Errors, Control, actions } from 'react-redux-form';
import { Link } from 'react-router-dom';
import classNames from 'classnames';
import {I18n, translate} from 'react-i18next';
import { compose, Dispatch } from 'redux';

import Modal from 'react-bootstrap/lib/Modal';
import { Button, Card, CardBody, CardGroup, Col, Container, Input, InputGroup, InputGroupAddon, InputGroupText, Row } from 'reactstrap';
import { AppState } from '../../../rootReducer';

import {ResizableBox} from 'react-resizable';
import * as throttle from 'lodash.throttle';

import XTerm from '../xTermComponent/xTermComponent';
import StatusBar from '../statusBar/statusBar';
import * as styles from './sandbox.css';
import Network from '../network/network';
import Syscalls from '../syscalls/syscalls';
import Messages from '../messages/messages';

interface Props {
	dispatch: Dispatch<any>;
	history;
	location;
	t;
	i18n;
}

interface Refs {
    [k: string]: any
}

interface State {
}

class Sandbox extends React.Component<any, State> {
    render() {
        return (
            <I18n ns="translations">
                {(t, {i18n}) => (
                    <div className={styles.layout}>
                        <header className={styles.header}>
                            <StatusBar />
                        </header>

                        <div className={styles.subLayout}>

                            <div className={styles.layoutLeft}>

                                <main className={styles.main}>
                                    <XTerm
                                        addons={ [ 'fit', 'fullscreen', 'search', 'terminado', 'attach' ] }
                                        style={{
                                            overflow: 'hidden',
                                            position: 'relative',
                                            width: '100%',
                                            height: '100%'
                                        }}
                                    />
                                </main>

								<Messages/>
								<Syscalls />
								<Network />
                            </div>

                            <aside className={styles.aside}>
                                <div>Processes</div>
                            </aside>

                        </div>
                    </div>
                )}
            </I18n>
        );
    }
}

const select = (state: AppState) => ({});

export default connect(select)(Sandbox);