import React, { Component } from 'react';
import { 
    Button,
    Card,
    Form,
    Input,
} from 'reactstrap';
import { Message } from './Message'

export class ChatPanel extends Component {

    constructor(props) {
        super(props);

        this.last = null;

        this.setRef = element => {
            this.last = element;
        }
    }

    componentDidUpdate() {
        this.last.scrollIntoView({ behavior: "smooth" });
    }

    scroll = () => {
        this.last.scrollIntoView({ behavior: "smooth" });
    }

	createChats(chats) {
        let user = sessionStorage.getItem('username')
		return chats.map((chat, idx) =>
			<Message
                chat={chat}
                key={idx}
                idx={idx}
                setRef={this.setRef}
                className={(chat['from'] === user)? "message user-message" : "message"}
            />
		);
	}

    render () {
        return (
            <Card className="chat-card">
                <Form className="message-form" onSubmit = {this.props.onSubmit}>
                    <Button outline onClick={this.onClick} color="secondary" className="scroll-btn">🡣</Button>
                    <Input type="textarea" value={this.props.message} onChange={this.props.onChange} id="msg" name="message" />
                    <Button type="submit" color="primary" className="submit">⮞</Button>
                </Form>
                <div className="chat-display">
                    <ul className="chats" ref="chats">
                        {this.createChats(this.props.chats)}
                    </ul>
                </div>
            </Card>
        );
    }
}
