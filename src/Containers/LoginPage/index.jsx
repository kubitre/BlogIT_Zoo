import React, {Component} from 'react';
import classNames from 'classnames';
import {NavLink} from 'react-router-dom';
import { connect } from 'react-redux';

import './style.scss';
import { bindActionCreators } from 'redux';

import * as loginActions from '../../Store/Actions/user';

class LoginPage extends Component{
    constructor(props){
        super(props);

        this.state = {
            "clicked": false,
            "username": "",
            "password": "",

            "validation_username": {
                "failed": false,
                "success": false,
            },
            "validation_password": {
                "failed": false,
                "success": false,
            },
        }

        this.checkChangeName = this.checkChangeName.bind(this);
        this.passwordValidation = this.passwordValidation.bind(this);
        this.handleLogin = this.handleLogin.bind(this);
    }

    passwordValidation(event){
        let str = event.target.value;
        this.setState({
            password: str,
        })

        if (str.length >= 6 && str.length <= 28){
            this.setState({
                validation_password: {
                    success: true,
                    failed: false,
                }
            })
        }
        else if (str.length == 0){
            this.setState({
                validation_password: {
                    success: false,
                    failed: false,
                }
            })
        }

        else{
            this.setState({
                validation_password:
                {
                    success: false,
                    failed: true
                }
            })
        }
        
    }

    checkChangeName(event) {
        let str = event.target.value;
        this.setState({
            username: str
        });

        console.log(str);

        if (str.length >= 6 && str.length <= 28){
            this.setState({
                "validation_username": {
                    "failed": false,
                    "success": true,
                }
            })
        }

        else if (str.length == 0){
            this.setState({
                "validation_username": {
                    success: false,
                    failed: false,
                }
            })
        }
        else{
            this.setState({
                "validation_username": {
                    "failed": true,
                    "success": false,
                }
            })
        }
    }

    handleLogin(event){
        if(
            this.state.validation_password.success && this.state.validation_username.success
        ){
            event.preventDefault();

            console.log("Credentials: ", this.state.username, " pass: ", this.state.password);
            const {startFetchingData} = this.props.actionForLogin;

            startFetchingData({
                username: this.state.username,
                password: this.state.password
            })
        }
    }

    render = () => {
        const {loaded, username, token, loadedByLS, userid} = this.props.userStore;
        const {loginByLocalStorage} = this.props.actionForLogin;

        if(loaded){

            console.log("loaded datas");
            let tokenFromStorage = window.localStorage.getItem("token");
            let usernameFromStorage = window.localStorage.getItem("username");
            let useridFromStorage = window.localStorage.getItem("userid");

            console.log("token: ", tokenFromStorage, "; username: ", usernameFromStorage, "userid: ", useridFromStorage);

            if (tokenFromStorage == null && usernameFromStorage == null && useridFromStorage == null){

                console.log("start setting datas to local storage");
                window.localStorage.setItem("username", username);
                window.localStorage.setItem("token", token);
                window.localStorage.setItem("userid", userid)
            }
        }
        else{
            console.log("tururu")
            let tokenFromStorage = window.localStorage.getItem("token");
            let usernameFromStorage = window.localStorage.getItem("username");
            let useridFromStorage = window.localStorage.getItem("userid");
            
            if (tokenFromStorage != null && usernameFromStorage != null && useridFromStorage != null){
                loginByLocalStorage({
                    username: usernameFromStorage,
                    token: tokenFromStorage,
                    userid: useridFromStorage,
                });
            }
        }
        return(
            !loadedByLS ?
            <div className="loginpage_container">
                <div className="box-form">
                    <div className="left">
                        <div className="overlay">
                            <h1>Hello World.</h1>
                            <p>On our page you can login into your account. Please, check your writing data!</p>
                            <span>
                                <p>login with social media(Not implemented)</p>
                                <a href="#"><i className="fa fa-twitter" aria-hidden="true"></i> Login with Twitter</a>
                            </span>
                        </div>
                    </div>
                    
                    <div className="right">
                        <h5>Login</h5>
                        <p>Don't have an account? <NavLink to='/registration'>Creat Your Account</NavLink> it takes less than a minute</p>
                        <div className="inputs">
                            <input type="text" 
                                className={classNames("username_input", {"failed": this.state.validation_username.failed}, 
                                                                        {"success": this.state.validation_username.success})} 
                                placeholder="user name" 
                                onChange={this.checkChangeName}/>
                            
                            <br/>
                            <input  type="password" 
                                    placeholder="password"
                                    className={classNames("password_input", {"success": this.state.validation_password.success},
                                                                            {"failed": this.state.validation_password.failed}
                                    )}
                                    onChange={this.passwordValidation}
                                    />
                        </div>
                            
                        <br></br>
                            
                        <div className="remember-me--forget-password"></div>
                        <label>
                            <input type="checkbox" name="item" checked />
                            <span className="text-checkbox">Remember me</span>
                        </label>
                        <p>forget password?</p>
                            
                        <br></br>
                        <button onClick={this.handleLogin}>Login</button>
                    </div>
                    
                </div>
            </div>

            :

            <div className="loginpage_container">
                <div className="loadedByLS">
                    <div className="status">
                        Вы уже вошли как <div className="username">{username}</div>
                    </div>
                </div>
            </div>
        )
    }
}

const mapStateToProps = (state) => {
    return {
        userStore: state.USI_userState
    };
};

const mapDispatchToProps = (dispatch) => {
    return {
        actionForLogin: bindActionCreators(loginActions, dispatch)
    };
};

export default connect(mapStateToProps, mapDispatchToProps)(LoginPage);