import * as React from 'react';
import * as styles from './filterMessages.css';
import { AppState } from '../../../rootReducer';
import { Message } from '../../interfaces/message';
import { connect } from 'react-redux';
import { FormEvent } from 'react';
import { SetMessageFilter } from '../../honeyboxActions';

interface Props {
	messages: Message[];
	dispatch: any;
}

interface State {
	input: string;
}

class FilterMessages extends React.Component<Props, State> {
	state: State = {
		input: ''
	};

	onChange(event: FormEvent<HTMLInputElement>) {
		const { dispatch } = this.props;
		const { input } = this.state;

		this.setState({
			input: event.currentTarget.value
		});

		dispatch(new SetMessageFilter({ query: event.currentTarget.value }));
	}

	render() {
		const { input } = this.state;

		return (
			<form className={styles.filterMessages}>
				<input
					onChange={this.onChange.bind(this)}
					value={ input }
					className={styles.input}
					placeholder="Filter messages"
				/>
				<input type="submit" hidden />
			</form>
		);
	}
}

const select = (state: AppState) => ({
	messages: state.honeyfarm.messages
});

export default connect(select)(FilterMessages);