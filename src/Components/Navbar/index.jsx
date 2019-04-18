import React, {Component} from 'react';

import {NavLink} from 'react-router-dom';

import './style.scss';

export default class NavBar extends Component{
    render = () => {
        return (
            <nav className="navbar navbar-default navbar-fixed-top">
                <div className="container">
                    <div className="navbar-header">
                    {/* <button type="button" className="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
                        <span className="sr-only">Toggle navigation</span>
                        <span className="icon-bar"></span>
                        <span className="icon-bar"></span>
                        <span className="icon-bar"></span>
                    </button> */}
                    </div>
                    <div id="navbar" className="collapse navbar-collapse">
                    <ul className="nav navbar-nav navbar-right">
                        <li><NavLink className="aa" to="/articles">Articles</NavLink></li>
                        <li><a className="aa" href="https://www.kubitre.me">Author</a></li>
                        <li><NavLink className="aa" to="/login">Sign IN</NavLink></li>
                        <li> | </li>
                        <li><NavLink className="aa" to="/registration">Sign UP</NavLink></li>
                    </ul>
                    </div>
                </div>
            </nav>
        )
    }
}