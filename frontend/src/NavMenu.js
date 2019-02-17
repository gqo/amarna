import React, { Component }  from 'react';
import {
    Collapse,
    Navbar,
    NavbarBrand,
    NavbarToggler,
    Nav,
    NavItem,
    NavLink,
} from 'reactstrap'

export default class NavMenu extends Component {

    constructor(props) {
        super(props)

        this.toggle = this.toggle.bind(this);
        this.state = {
            isOpen: false
        }
    }

    toggle () {
        this.setState({
            isOpen: !this.state.isOpen
        });
    }

    render () {
        return (
            <div className="nav-header">
                <Navbar color="dark" dark expand="md">
                    <NavbarBrand href="/">Amarna</NavbarBrand>
                    <NavbarToggler onClick={this.toggle} />
                    <Collapse isOpen={this.state.isOpen} navbar>
                        <Nav className="ml-auto" navbar>
                            {(sessionStorage.getItem('username') === null)? <React.Fragment><NavItem><NavLink href="/profile">Register</NavLink> </NavItem><NavItem><NavLink href="/login">Login</NavLink></NavItem></React.Fragment> : <React.Fragment><NavItem><NavLink href="/chats">Chats</NavLink></NavItem><NavItem><NavLink href="/logout">Logout</NavLink></NavItem></React.Fragment>}
                        </Nav>
                    </Collapse>
                </Navbar>
            </div>
        );
    }
}
