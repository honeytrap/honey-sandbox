import * as React from 'react';
import { connect, Dispatch } from 'react-redux';
import { withRouter } from 'react-router-dom';
import { replace, push } from 'react-router-redux'
import { Link } from 'react-router-dom';
import { LocalForm, Errors, Control, track, actions } from 'react-redux-form';
import classNames from 'classnames';
import { I18n, translate } from 'react-i18next';
import ReactPasswordStrength from 'react-password-strength';
import {
	Button,
	Card,
	CardBody,
	CardGroup,
	Col,
	Container,
	Input,
	InputGroup,
	InputGroupAddon,
	InputGroupText,
	FormGroup,
	Row
} from 'reactstrap';
import { SmallScreen } from '../../../shared/components/smallScreen/smallScreen';
import { t } from 'i18next';
import { SignUp } from '../../authenticationActions';
import { AppState } from '../../../rootReducer';

interface Props {
	dispatch: Dispatch<any>;
	history: any;
	i18n: any;
	t: any;
	signUpErrors: any;
}

interface State {
	loading: boolean;
}

class SignupForm extends React.Component<Props, State> {
	passwordInput;
	formDispatch;

	constructor(props) {
		super(props);

		this.state = {
			loading: false,
		};

		this.passwordInput = (props) => <ReactPasswordStrength
			className=""
			minLength={8}
			minScore={2}
			tooShortWord="te zwak"
			scoreWords={['zwak', 'matig', 'goed', 'sterk', 'sterker']}
			changeCallback={(props) => {
				this.formDispatch(actions.change('user.password', props.password));
				this.formDispatch(actions.setValidity('user.password', props.isValid));
			}}
			inputProps={{
				...props,
				className: "text_field1",
			}}
		/>;
	}

	componentDidMount() {
		const {i18n} = this.props;
	}

	handleSubmit(values) {
		const { dispatch } = this.props;
		const { name, company, email, password } = values;

		dispatch(new SignUp({
			name,
			company,
			email,
			password
		}));
	}

	componentWillUpdate(nextProps, nextState) {
		const {dispatch} = nextProps;
	}

	renderFormGroup(model: string, type: 'text' | 'password', placeholderKey: string, icon: any) {
		const errorInput = ({valid, ...otherProps}) =>
			<input className={`text_field1 ${valid ? 'no-error' : 'has-error'}`} { ...otherProps } />;

		return (
			<FormGroup>
				<InputGroup className="mb-3">
					<InputGroupAddon addonType="prepend">
						<InputGroupText>
							{icon}
						</InputGroupText>
					</InputGroupAddon>
					<Control.text
						model={model}
						placeholder={t(placeholderKey)}
						component={errorInput}
						className="form-control"
						type={type}
						mapProps={{
							valid: ({fieldValue}) => fieldValue.valid,
						}}
						validators={{
							required: (val) => val.length,
						}}
					/>
				</InputGroup>

				<Errors
					className="error-msg text-danger"
					model={model}
					show="touched"
					messages={{
						required: t('required'),
					}}
				/>
			</FormGroup>
		)
	}

	componentWillReceiveProps(nextProps: Props) {
		if (this.formDispatch && nextProps.signUpErrors) {
			this.formDispatch(actions.setFieldsErrors('user', nextProps.signUpErrors));
		}
	}

	render() {
		return (
			<SmallScreen>
				<h1>Sign up</h1>
				<p>Create your account</p>

				<LocalForm
					model="derp"
					initialState={{
						name: '',
						errors: {
							name: 'vdenj'
						}
					}}>
				</LocalForm>

				<LocalForm
					className="cnt_form"
					model="user"
					initialState={{
						name: '',
						company: '',
						email: '',
						password: '',
						password2: ''
					}}
					validators={{
						'': {
							passwordsMatch: (vals) => vals.password === vals.password2,
						},
					}}
					onSubmit={(values) => this.handleSubmit(values)}
					getDispatch={dispatch => this.formDispatch = dispatch}>

					{this.renderFormGroup(
						'user.name',
						'text',
						'fullName',
						<i className="icon-user" />
					)}

					{this.renderFormGroup(
						'user.company',
						'text',
						'company',
						<i className="icon-user" />
					)}

					{this.renderFormGroup(
						'user.email',
						'text',
						'email',
						'@'
					)}

					{this.renderFormGroup(
						'user.password',
						'password',
						'password',
						<i className="icon-lock" />
					)}

					{this.renderFormGroup(
						'user.password2',
						'password',
						'confirmPassword',
						<i className="icon-lock" />
					)}

					<Control.button
						model="user"
						className={ classNames({ btn: true, 'btn-success': true, 'btn-block': true }) }
						disabled={{ valid: false }}
						>
						{t('pages.signup.formSubmit')}
					</Control.button>
				</LocalForm>
			</SmallScreen>
		);
	}
}

const select = (state: AppState) => ({
	signUpErrors: state.authentication.signUpErrors
});

export default connect(select)(SignupForm);
