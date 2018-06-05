import * as React from 'react';
import * as styles from './network.css';
import FooterComponent
	from '../../../shared/components/footerComponent/footerComponent';

interface Props {
}

class Network extends React.Component<Props> {
	render() {
		return (
			<FooterComponent title="Network">
				Display all outgoing and incoming network requests.
			</FooterComponent>
		);
	}
}

export default Network;