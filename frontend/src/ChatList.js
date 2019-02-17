import React, { Component } from 'react';
import {
    Container,
    Table,
} from 'reactstrap'
import PROPS from './server-properties.json'

export class ChatList extends Component {

    constructor(props) {
        super(props)

        this.state = {
            chats: []
        }
    }

    componentDidMount() {
        let user = sessionStorage.getItem('username');
        let url = PROPS['base'] + PROPS['getChats'];
        url = url.replace('{}', user);
        fetch(url, {
            method: 'GET',
            mode: 'cors',
            headers: {
                'Accept': 'application/json'
            }
        }).then((response) => {
            return response.json();
        }).then((json) => {
            this.setState({
                chats: json['pairings']
            })
        }).catch((error) => {
            console.log(error);
        });
    }

    createRows = () => {
        return this.state.chats.map((user, idx) => {
            return (
                <tr key={idx}>
                    <td><a href={"/chat/"+user}>{user}</a></td>
                </tr>
            );
        });
    }

    render() {
        return(
            <Container>
                <Table hover responsive>
                    <thead>
                        <tr>
                            <th>Username</th>
                        </tr>
                    </thead>
                    <tbody>
                        {this.createRows()}
                    </tbody>
                </Table>
            </Container>
        );
    }
}
