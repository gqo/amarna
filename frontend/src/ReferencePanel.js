import React, { Component } from 'react';
import { 
    Button,
    Card, 
} from 'reactstrap';

export class ReferencePanel extends Component {

    render () {
        return (
            <Card className="lesson">
                <pre>
                    {this.props.lesson}
                </pre>
                <Button color="primary" onClick={this.props.onClick}>Request to continue</Button>
            </Card>
        );
    }

}
