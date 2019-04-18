import React, {Component} from 'react';
import SearchComponent from '../../Components/Search';
import LogotypeComponent from '../../Components/Logotype';
import FooterComponent from '../../Components/Footer';

import './index.scss';
import ArticleFastView from '../../Components/ArticleFast';
import NavBar from '../../Components/Navbar';

export default class MainPage extends Component{
    constructor(props){
        super(props);
        this.state = {
            positionX: 500
        }
        this.handleMouseMove = this.handleMouseMove.bind(this)
    }

    handleMouseMove(event){
        this.setState({
            positionX: event.screenX
        })
    }
    render = () => 
    <div className="mainpage_container" onMouseMove={this.handleMouseMove}>
        <NavBar />
        <SearchComponent visibility={this.state.positionX > 1800 ? false : true}/>
        <LogotypeComponent visibility={this.state.positionX > 1800 ? false : true}/>
        {
            <ArticleFastView visibility={this.state.positionX > 1800 ? true : false}/>
        }
        <FooterComponent visibility={this.state.positionX > 1800 ? false: true}/>
    </div>
}