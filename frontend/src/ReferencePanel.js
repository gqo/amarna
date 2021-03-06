import React, { Component } from 'react';
import { 
    Button,
    Card, 
} from 'reactstrap';

export class ReferencePanel extends Component {

    render () {
        return (
            <Card className="lesson">
                <div className="content">
                    <h3>{this.props.lesson['title']}</h3>
                    <h4>{this.props.lesson['section']}</h4>
                    <p>
                        {this.props.lesson['description']}
                    </p>
                </div>
                <Button color="primary" onClick={this.props.onClick}>Request to continue</Button>
            </Card>
        );
    }

}
