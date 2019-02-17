import React, { Component } from 'react';

export class Message extends Component {

    render () {
        return (
            <li ref={this.props.setRef} className={this.props.className}>
                <pre>{this.props.chat['body']}</pre>
            </li>
        );
    }
}
