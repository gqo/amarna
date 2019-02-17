import React, { Component } from 'react';
import { ChatPanel } from './ChatPanel'
import { ReferencePanel } from './ReferencePanel'
import { Container } from 'reactstrap'
import './Chatroom.css'
import PROPS from './server-properties.json'

export class Chatroom extends Component {

	constructor(props) {
		super(props)

		this.state = {
			chats: [],
			message: ""
		}
	}

	componentDidMount() {
        let url = PROPS['base'] + PROPS['getMessages'];
        url = url.replace('{1}', sessionStorage.getItem('username'));
        url = url.replace('{2}', this.props.match.params['user'])
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
                chats: json['letters']
            })
        }).then(() => {
            this.updateLesson(this.props.match.params['user']);
        }).catch((error) => {
                console.log(error);
        });
	}

    updateLesson = (user) => {
        let url = PROPS['base'] + PROPS['getLesson'];
        url = url.replace('{1}', sessionStorage.getItem('username'));
        url = url.replace('{2}', user)
        fetch(url, {
            method: 'GET',
            mode: 'cors',
            headers: {
                'Accept': 'application/json'
            }
        }).then((response) => {
            return response.json();
        }).then((json) => {
            let newState = this.state
            newState['lesson'] = json['lesson']
            this.setState(newState)
        }).catch((error) => {
            console.log(error);
        });
    }

    requestNext = () => {
        let url = PROPS['base'] + PROPS['requestNext'];
        url = url.replace('{1}', sessionStorage.getItem('username'));
        url = url.replace('{2}', this.props.match.params['user'])
        fetch(url, {
            method: 'GET',
            mode: 'cors',
            headers: {
                'Accept': 'application/json'
            }
        }).then((response) => {
            return response.json();
        }).then((json) => {
            //
        }).then(() => {
            window.location.reload();
        }).catch((error) => {
            console.log(error);
        });
    }

	handleChange = (event) => {
		let newState = this.state
		newState['message'] = event.target.value
		this.setState(newState)
	}

	handleSubmit = (event) => {
		event.preventDefault();
        let body = JSON.stringify({ 
            "letter": this.state['message'],
            "leftUsername": sessionStorage.getItem('username'),
            "rightUsername": this.props.match.params['user']
        })
        let url = PROPS['base'] + PROPS['submit']
        fetch(url, {
            method: "POST",
            mode: "cors",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json"
            },
            body: body
        }).then((response) => {
            return response.json();
        }).then((json) => {
            //
        }).catch((error) => {
            console.log(error);
        });
		let newState = this.state
		newState['message'] = ""
		this.setState(newState)
	}

    render() {
        return (
        	<Container className="chatroom-container">
	        	<ChatPanel
	        		chats={this.state.chats}
                    onChange={this.handleChange}
                    onSubmit={this.handleSubmit}
                    message={this.state.message}
	    		/>
                <ReferencePanel 
                    user={this.props.match.params['user']}
                    lesson={this.state.lesson}
                    onClick={this.requestNext}
                />
			</Container>
        );
    }

}
