import React from 'react';
import {Switch, Route} from 'react-router-dom';
import MainPage from './Containers/MainPage';
import NotFounPage from './Containers/NotFounPage';
import ArticlesPageContainer from './Containers/ArticlesPage/index';
import LoginPage from './Containers/LoginPage';
import RegistrationPage from './Containers/RegistrationPage';
import Article from './Components/Article';


export default function Routes(props){ 
    return(
        <Switch>
            <Route exact path="/" component={MainPage}/>
            <Route path="/articles" component={ArticlesPageContainer}/>
            <Route path="/article/:id" component={Article}/>
            <Route path="/login" component={LoginPage} />
            <Route path="/registration" component={RegistrationPage}/>
            <Route component={NotFounPage}/>
        </Switch>
    )
}
