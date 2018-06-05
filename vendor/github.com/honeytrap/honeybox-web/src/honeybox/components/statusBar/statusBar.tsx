import * as React from 'react';
import { connect } from 'react-redux';
import { AppState } from '../../../rootReducer';
import * as styles from './statusBar.css';

interface Props {
	connected: boolean;
	name: string;
}

class StatusBar extends React.Component<Props> {
	render() {
		const { connected, name } = this.props;

		return (
			<div className={styles.statusBar}>
				<h1 className={styles.name}>{name}</h1>
				<div className={styles.connection}>
					<span className={styles.dot + ' ' + (connected ? styles.connected : styles.notConnected)} />
					<p className={styles.message}>{ connected ? 'Connected' : 'Not connected'}</p>
				</div>
			</div>
		)
	}
}

const select = (state: AppState) => ({
	connected: state.connection.connected,
	name: state.honeyfarm.name
});

export default connect(select)(StatusBar);