import React, { Component } from 'react';
import {
    Route,
} from 'react-router-dom'
import { Layout } from './Layout'
import { Home } from './Home'
import './App.css';

class App extends Component {
    static displayName = App.name;

    render() {
        return (
            <Layout>
                <Route exact path='/' component={Home} />
            </Layout>
        );
    }
}

export default App;
