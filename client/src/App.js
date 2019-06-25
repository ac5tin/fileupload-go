import React from 'react';
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

import Navbar from './components/nav';
import Landing from './components/landing';


import './static/css/App.css';
import './static/css/simple_latest.min.css';


// react toastify
toast.configure({
    autoClose: 3000,
    hideProgressBar: true
})

const App = () =>{
    return (
        <div className="App">
            <Navbar />
            <div className="container">
                <Landing />
            </div>
        </div>
    );
}

export default App;
