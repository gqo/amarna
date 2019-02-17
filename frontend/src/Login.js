import React, { Component } from 'react';
import {Form, FormGroup, Label, Input} from 'reactstrap';
import { 
    Button,
    Container,
} from 'reactstrap'
import { Redirect} from 'react-router'
import PROPS from './server-properties.json'

class Login extends Component {

    constructor (props) {
        super(props)

        this.state = {
            username: "",
            redirect: false
        }
    }

    login = () => {

       let url = PROPS['base'] + PROPS['validate']
       url = url.replace('{}', this.state.username)
        fetch(url, {
            method: 'GET',
            mode: 'cors',
            headers: {
                'Accept': 'application/json'
            }
        }).then((response) => {
            return response.json();
        }).then((json) => {
            if (json['validate']) {
                sessionStorage.setItem('username', this.state.username);
            }
            this.setState({ redirect: true });
        })
    }

    componentDidMount () {
        if (sessionStorage.getItem('username') !== null) {
            sessionStorage.removeItem('username')
            this.setState({ redirect: true });
        }
    }

    render () {
        const { from } = this.props.location.state || { from : { pathname: "/" } }
        const { redirect } = this.state;

        if (redirect) {
            return <Redirect to={from} />;
        }

        return (
            <Container>
                <h1>Log In</h1>
                <Form onSubmit={this.onFormSubmit}>
                <FormGroup>
                <Label for="username"></Label>
                <Input type="username" 
                        name="username" 
                        id="username" 
                        value={this.state.username} 
                        onChange={e => this.setState({ username: e.target.value })}
                        placeholder="ex. janedone123" />
                </FormGroup>
                </Form>
                <Button outline color="secondary" id="username" onClick={this.login}>Log In</Button>
            </Container>
        );
    }
}

export default Login;
