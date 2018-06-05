import * as React from 'react';
import * as styles from './syscalls.css';
import FooterComponent
	from '../../../shared/components/footerComponent/footerComponent';

interface Props {
}

class Syscalls extends React.Component<Props> {
	render() {
		return (
			<FooterComponent title="Syscalls">
				Cool stuff and more!
			</FooterComponent>
		);
	}
}

export default Syscalls;