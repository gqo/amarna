import React from 'react';
import {
    Route,
    Redirect,
} from 'react-router-dom'
import { Layout } from './Layout'
import { Home } from './Home'
import { ChatList } from './ChatList'
import { Chatroom } from './Chatroom'
import Login from './Login'
import { Profile } from './Profile'
import './App.css';

const App = () => (
            <Layout>
                <Route exact path='/' component={Home} />
                <Route path='/login' component={Login} />
                <Route path='/logout' component={Login} />
                <PrivateRoute path="/chats" component={ChatList} />
                <PrivateRoute path="/chat/:user" component={Chatroom} />
                <Route path="/profile" component={Profile} />
            </Layout>
);

const PrivateRoute = ({ component: Component, ...rest }) => (
    <Route
        {...rest}
        render={props =>
            (sessionStorage.getItem('username') !== null)? (
                <Component {...props} />
            ) : (
                <Redirect
                to={{
                    pathname: "/login",
                    state: { from: props.location }
                }}
                />
            )
        }
    />
);

export default App;
