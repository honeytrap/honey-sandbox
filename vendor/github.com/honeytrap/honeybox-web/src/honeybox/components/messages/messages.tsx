import * as React from 'react';
import * as styles from './messages.css';
import FooterComponent
	from '../../../shared/components/footerComponent/footerComponent';
import { AppState } from '../../../rootReducer';
import { Message } from '../../interfaces/message';
import { connect } from 'react-redux';
import FilterMessages from '../filterMessages/filterMessages';
import { selectFilteredMessages } from '../../honeyboxSelectors';

interface Props {
	messages: Message[];
}

class Messages extends React.Component<Props> {
	render() {
		const { messages } = this.props;

		return (
			<FooterComponent title="Messages" headerContent={<FilterMessages/>}>
				<ul className={styles.messageList}>
				{messages.map((message, i) => (
					<li key={i} className={styles.message}>{message.text}</li>
				))}
				</ul>
			</FooterComponent>
		);
	}
}

const select = (state: AppState) => ({
	messages: selectFilteredMessages(state)
});

export default connect(select)(Messages);