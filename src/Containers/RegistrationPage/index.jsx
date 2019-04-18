import React, {Component} from 'react';

import {NavLink} from 'react-router-dom';

import './style.scss';

export default class RegistrationPage extends Component{
    constructor(props){
        super(props);

        this.state = {
            "selectedPhoto": null,
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
        };
    }
    onEnterImageData(event){
        this.setState({
            selectedPhoto: event.target.files[0],
        })
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

    handleRegister(event){
        if(
            this.state.validation_password.success && this.state.validation_username.success
        ){
            event.preventDefault();

            console.log("Start fetching: ");
            console.log(this.state)
            // const {startFetchingData} = this.props.actionForLogin;

            // startFetchingData({
            //     username: this.state.username,
            //     password: this.state.password
            // })
        }
    }

    emailValidation(event){
        this.setState({
            email: event.target.value
        });
    }

    render = () => {
        return(
            <div className="registrationpage_container">
                <div class="box-form">
                    <div class="left">
                        <div class="overlay">
                            <h1>Hello World.</h1>
                            <p>please, have a white data</p>
                        </div>
                    </div>
                    
                    <div class="right">
                        <h5>Registration</h5>
                        <p>Have an account? <NavLink to='/login'>Login</NavLink></p>
                        <div class="inputs">
                            <input type="file" onChange={this.onEnterImageData.bind(this)}/>
                            <br/>
                            <input type="text" placeholder="user name" onChange={this.checkChangeName.bind(this)}/>
                            <br/>
                            <input type="text" placeholder="email" onChange={this.emailValidation.bind(this)}/>
                            <br/>
                            <input type="password" placeholder="password" onChange={this.passwordValidation.bind(this)}/>
                            
                        </div>
                            
                        <br></br>
                            
                        <div class="remember-me--forget-password"></div>
                        <label>
                            <input type="checkbox" name="item" checked />
                            <span class="text-checkbox">Remember me(Not implemented)</span>
                        </label>
                        <p>forget password?(Not implemented)</p>
                            
                        <br></br>
                        <button onClick={this.handleRegister.bind(this)}>Registration</button>
                    </div>
                    
                </div>
            </div>
        )
    }
}