import React, { Component } from 'react';
import { 
    Button,
    Container,
    Form,
    FormGroup,
    Label,
    Input
} from 'reactstrap';
import PROPS from './server-properties.json'

export class Profile extends Component {
	constructor(props) {
		super(props);
		this.onFormSubmit = this.onFormSubmit.bind(this);
		this.state = {
			username: "",
			nativeLang: "",
			newLang: "",
		}
	}

	onFormSubmit() {
        let url = PROPS['base'] + PROPS['register']
        url = url.replace('{1}', this.state.username)
        url = url.replace('{2}', this.state.nativeLang)
        url = url.replace('{3}', this.state.newLang)
        fetch(url, {
            method: "GET",
            mode: "cors",
            headers: {
                "Accept": "application/json",
            },
        }).then((response) => {
            return JSON.parse(response)
        }).then((json) => {
            //
        }).catch((error) => {
            console.log(error);
        });
        this.setState({
			username: "",
			nativeLang: "",
			newLang: "",
        });
	}

  	render() {
    	return (
            <Container>
                <Form onSubmit={this.onFormSubmit}>
                <FormGroup>
                  <Label for="username">Username</Label>
                  <Input type="user" 
                         name="user" 
                         id="username" 
                         value={this.state.username}
                         onChange={e => this.setState({ username: e.target.value })}
                         placeholder="ex. johndoe123" />
                </FormGroup>
                <FormGroup>
                  <Label for="nativeLang">Choose the language you are most proficient in: </Label>
                  <Input type="select" 
                         name="nativeLang" 
                         value={this.state.nativeLang}
                         onChange={e => this.setState({ nativeLang: e.target.value })} 
                         id="nativeLang">
                    <option>English</option>
                    <option>Spanish</option>
                    <option>Chinese</option>
                    <option>Japanese</option>
                  </Input>
                </FormGroup>
                <FormGroup>
                  <Label for="newLang">Choose the language you are attempting to learn: </Label>
                  <Input type="select" 
                         name="newLang" 
                         value={this.state.newLang}
                         onChange={e => this.setState({ newLang: e.target.value })} 
                         id="newLang">
                    <option>English</option>
                    <option>Spanish</option>
                    <option>Chinese</option>
                    <option>Japanese</option>
                  </Input>
                </FormGroup>
            <Button>Submit</Button>
          </Form>
        </Container>
		);
	
  	}
}
