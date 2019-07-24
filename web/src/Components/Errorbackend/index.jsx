import React, {Component} from 'react';
import {NavLink} from 'react-router-dom';


import './style.scss';

export default class ErrorBackend extends Component{
    render = () => {
        return(
            <div className="bg-purple">
                <div className="stars">
                    <div className="custom-navbar">
                        <div className="brand-logo">
                            <h1 style={{color: 'white'}}>Kubitre Blog</h1>
                        </div>
                        <div className="navbar-links">
                            <ul>
                            <li><NavLink to="/">Home</NavLink>></li>
                            <li><NavLink to="/articles">Articles</NavLink></li>
                            <li><a href="http://www.kubitre.me/" target="_blank">Autor</a></li>
                            </ul>
                        </div>
                    </div>
                    <div className="central-body">
                        <img className="image-404" src="http://salehriaz.com/404Page/img/404.svg" width="300px"/>
                        <NavLink to="/" className="btn-go-home">GO BACK HOME</NavLink>
                    </div>
                    <div className="objects">
                        <img className="object_rocket" src="http://salehriaz.com/404Page/img/rocket.svg" width="40px"/>
                        <div className="earth-moon">
                            <img className="object_earth" src="http://salehriaz.com/404Page/img/earth.svg" width="100px"/>
                            <img className="object_moon" src="http://salehriaz.com/404Page/img/moon.svg" width="80px"/>
                        </div>
                        <div className="box_astronaut">
                            <img className="object_astronaut" src="http://salehriaz.com/404Page/img/astronaut.svg" width="140px"/>
                        </div>
                    </div>
                    <div className="glowing_stars">
                        <div className="star"></div>
                        <div className="star"></div>
                        <div className="star"></div>
                        <div className="star"></div>
                        <div className="star"></div>

                    </div>

                </div>
            </div>
        )
    }
}