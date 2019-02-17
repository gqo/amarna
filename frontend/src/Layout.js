import React, { Component } from 'react';
import NavMenu from './NavMenu'

export class Layout extends Component {
    render () {
        return (
            <React.Fragment>
                <NavMenu />
                {this.props.children}
            </React.Fragment>
        );
    }
}
