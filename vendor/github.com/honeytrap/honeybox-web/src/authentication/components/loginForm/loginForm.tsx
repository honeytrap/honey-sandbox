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
import { Login } from '../../authenticationActions';
import { AppState } from '../../../rootReducer';

const mailFormat = /[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]+$/;

interface Props {
	dispatch: Dispatch<any>;
	history;
	location;
	t;
	i18n;
}

interface State {
	username: string;
	password: string;
	borderBottom: string;
	loading: boolean;
	error: boolean;
	showModal: boolean;
}

class LoginForm extends React.Component<any, State> {
	formDispatch;
	state: State = {
		username: '',
		password: '',
		borderBottom:'',
		loading:false,
		error:false,
		showModal: false
	};

	attachDispatch(dispatch) {
		this.formDispatch = dispatch;
	}

	handleSubmit(values) {
		this.setState({error:false});

		const { dispatch } = this.props;

		dispatch(new Login({
			username: values.username,
			password: values.password
		}));
	}

	handleHide() {
		this.setState({ showModal: false });
	}

	render() {
		const { t, i18n } = this.props;

		const usernameInput = (props) => <input autoComplete="username" autoFocus className={`form-control ${props.valid ? 'no-error' : 'has-error'}`} {...props} />;
		const passwordInput = (props) => <input autoComplete="current-password" className={`form-control ${props.valid ? 'no-error' : 'has-error'}`} {...props} />;

		return (
			<I18n ns="translations">
				{
					(t, {i18n}) => (
						<div>
							<div className="app flex-row align-items-center">
								<Container>
									<Row className="justify-content-center">
										<Col md="8">
											<CardGroup>
												<Card className="p-4">
													<CardBody>
														<LocalForm className="cnt_form_inlog"
																   getDispatch={(dispatch) => this.attachDispatch(dispatch)}
																   model="user"
																   initialState={{
																	   username: '',
																	   password: ''
																   }}
																   onSubmit={(values) => this.handleSubmit(values)}
														>
															<Modal show={this.state.showModal} onHide={() => this.handleHide() } className={ classNames({'login-error-popup': true, visible: this.state.error }) }>
																<Modal.Header closeButton className="login-error-header"></Modal.Header>
																<Modal.Body className="login-error-body">
																	De ingevoerde email adres of wachtwoord is onjuist. Controleer de gegevens en probeer het opnieuw.
																	<Errors className="failed_error-msg alert-danger "
																			show="touched"
																			model="user"
																			messages={{
																				isLoginFailed: 'Email adres of wachtwoord is onjuist.',
																			}}
																	/>
																</Modal.Body>
															</Modal>
															<h1>Login</h1>
															<p className="text-muted">Sign In to your account</p>
															<InputGroup className="mb-3">
																<InputGroupAddon addonType="prepend">
																	<InputGroupText>
																		<i className="icon-user"></i>
																	</InputGroupText>
																</InputGroupAddon>
																<Control.text
																	model="user.username"
																	validateOn="change"
																	placeholder={t('pages.login.formUser')}
																	component={ usernameInput }
																	mapProps={{
																		valid: ({fieldValue}) => fieldValue.valid
																	}}
																	validators={{
																		required: (val) => val.length,
																		format: (val) => mailFormat.test(val) || val.length === 0
																	}}
																/>
															</InputGroup>
															<Errors className="error-msg"
																	model="user.username"
																	show="touched"
																	messages={{
																		required: t('pages.login.userError'),
																		format: "Gebruik je emailadres."
																	}}
															/>
															<div className="clearfix"></div>
															<InputGroup className="mb-4">
																<InputGroupAddon addonType="prepend">
																	<InputGroupText>
																		<i className="icon-lock"></i>
																	</InputGroupText>
																</InputGroupAddon>
																<Control.text
																	model="user.password"
																	placeholder={t('pages.login.formPassword')}
																	type="password"
																	validateOn="change"
																	component={ passwordInput }
																	mapProps={{
																		valid: ({fieldValue}) => fieldValue.valid
																	}}
																	validators={{
																		required: (val) => val.length,
																	}}
																/>
															</InputGroup>
															<Errors className="error-msg"
																	model="user.password"
																	show="touched"
																	messages={{
																		required: t('pages.login.passwordError')
																	}}
															/>
															<Row>
																<Col xs="6">
																	<Control.button
																		model="user"
																		className={ classNames({"btn": true, "btn-primary": true}) }
																		disabled={{valid: false}}
																		type="submit"
																	>
																		{t('pages.login.formSubmit')}
																	</Control.button>
																</Col>
																<Col xs="6" className="text-right">
																	<Link to={"/forgot-password" } className="px-0 btn btn-link">{t('pages.forgotPassword.title')}</Link>
																</Col>
															</Row>
														</LocalForm>
													</CardBody>
												</Card>
												<Card className="text-white bg-primary py-5 d-md-down-none" style={{ width: 44 + '%' }}>
													<CardBody className="text-center">
														<div>
															<h2>Sign up</h2>
															<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut
																labore et dolore magna aliqua.</p>
															<Link to={"/sign-up" }>
																<Button color="primary" className="mt-3" active>Register Now!</Button>
															</Link>
														</div>
													</CardBody>
												</Card>
											</CardGroup>
										</Col>
									</Row>
								</Container>
							</div>
						</div>
					)
				}
			</I18n>
		);
	}
}

const select = (state: AppState) => ({});

export default connect(select)(LoginForm);
