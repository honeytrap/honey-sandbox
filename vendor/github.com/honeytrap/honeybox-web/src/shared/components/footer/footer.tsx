import * as React from 'react';

class DefaultFooter extends React.Component<any, any> {
    render() {

        // eslint-disable-next-line
        const { children, ...attributes } = this.props;

        return (
            <React.Fragment>
                <span><a href="https://dutchsec.com">Honeyfarm</a> &copy; 2018 DutchSec.</span>
                <span className="ml-auto">Powered by <a href="https://dutchsec.com">DutchSec</a></span>
            </React.Fragment>
        );
    }
}

export default DefaultFooter;
