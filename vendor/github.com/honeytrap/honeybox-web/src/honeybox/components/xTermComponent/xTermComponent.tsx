import * as React from 'react';
import * as XTerm from 'xterm';
import { fit } from 'xterm/lib/addons/fit/fit';

import '../../../../node_modules/xterm/dist/xterm.css';
import { connect } from 'react-redux';
import { StdIn } from '../../honeyboxActions';
import { AppState } from '../../../rootReducer';
import { StdOutInterface } from '../../interfaces/stdOutInterface';
import * as styles from './xTermComponent.css';

const className = require('classnames');

export interface IXtermProps extends React.DOMAttributes<{}> {
    onChange?: (e) => void;
    onInput?: (e) => void;
    onFocusChange?: Function;
    addons?: string[];
    onScroll?: (e) => void;
    onContextMenu?: (e) => void;
    options?: any;
    path?: string;
    value?: string;
    className?: string;
    style?: React.CSSProperties;
    dispatch: any;
    stdOut: StdOutInterface[];
}
export interface IXtermState {
    isFocused: boolean;
}

class XTermComponent extends React.Component<IXtermProps, IXtermState> {
    xterm: XTerm.Terminal;
    container: HTMLDivElement;
    input: string = '';
    writtenIds: string[] = [];
	state = {
		isFocused: false
	};

    componentDidMount() {
        this.xterm = new XTerm.Terminal(this.props.options);
        this.xterm.open(this.container);

		fit(this.xterm);

        this.xterm.on('focus', this.focusChanged.bind(this, true));
        this.xterm.on('blur', this.focusChanged.bind(this, false));

        if (this.props.onContextMenu) {
            this.xterm.element.addEventListener('contextmenu', this.onContextMenu.bind(this));
        }

        if (this.props.onInput) {
            this.xterm.on('data', this.onInput);
        }
        if (this.props.value) {
            this.xterm.write(this.props.value);
        }


		this.runFakeTerminal();
    }

    componentWillUnmount() {
        // is there a lighter-weight way to remove the cm instance?
        if (this.xterm) {
            this.xterm.destroy();
            this.xterm = null;
        }
    }

    shouldComponentUpdate(nextProps, nextState) {
		    // console.log('shouldComponentUpdate', nextProps.hasOwnProperty('value'), nextProps.value != this.props.value);
        if (nextProps.hasOwnProperty('value') && nextProps.value != this.props.value) {
            if (this.xterm) {
				        this.xterm.clear();
				        setTimeout(()=>{
					          this.xterm.write(nextProps.value);
				        },0)
            }
        }
        return false;
    }
    getTerminal() {
        return this.xterm;
    }
    write(data: any) {
        this.xterm && this.xterm.write(data);
    }
    writeln(data: any) {
        this.xterm && this.xterm.writeln(data);
    }
    focus() {
        if (this.xterm) {
            this.xterm.focus();
        }
    }
    focusChanged(focused) {
        this.setState({
            isFocused: focused
        });
        this.props.onFocusChange && this.props.onFocusChange(focused);
    }
    onInput = data => {
        this.props.onInput && this.props.onInput(data);
    };

    resize(cols: number, rows: number) {
        this.xterm && this.xterm.resize(Math.round(cols), Math.round(rows));
    }
    setOption(key: string, value: boolean) {
        this.xterm && this.xterm.setOption(key, value);
    }
    refresh() {
        this.xterm && this.xterm.refresh(0, this.xterm.rows - 1);
    }

    onContextMenu(e) {
        this.props.onContextMenu && this.props.onContextMenu(e);
    }

    sendToServer(string: string) {
		const { dispatch } = this.props;

		dispatch(new StdIn({ data: string }));
	}

    prompt() {
    	this.sendToServer(this.input + '\n');
		this.input = '';
		this.xterm.write('\r\n');
	}

	componentWillReceiveProps(nextProps: IXtermProps) {
    	const { stdOut } = this.props;

    	if (nextProps.stdOut !== stdOut) {
    		const unwritten = nextProps.stdOut.filter(out =>
				this.writtenIds.indexOf(out.id) === -1
			);

    		unwritten.forEach(out => {
				this.xterm.write(out.data);
				this.writtenIds.push(out.id);
			});

			this.xterm.write('$ ');
		}
	}

	runFakeTerminal() {
		this.prompt();

		this.xterm.write('$ ');

		this.xterm.on('key',  (key, ev) => {
			const printable = (
				!ev.altKey && !ev.ctrlKey && !ev.metaKey
			);

			if (ev.keyCode == 13) {
				this.sendToServer('\n');
				this.xterm.write('\r\n');

			} else if (ev.which === 8) {
				this.sendToServer('\x08');
				// this.xterm.write(key);

			} else if (printable) {
				this.sendToServer(key);
				this.xterm.write(key);
			}
		});

		this.xterm.on('paste',  (data, ev) => {
			this.xterm.write(data);
		});
	}

	render() {
        const terminalClassName = className('ReactXTerm', this.state.isFocused ? 'ReactXTerm--focused' : null, this.props.className, styles.container);
        return <div ref={ref => (this.container = ref)} className={terminalClassName} />;
    }
}

const select = (state: AppState) => ({
	stdOut: state.honeyfarm.stdOut
});

export default connect(select)(XTermComponent);