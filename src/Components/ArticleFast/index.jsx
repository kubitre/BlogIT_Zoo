import React, {Component} from 'react';

import './style.scss';
import ArticleFV from '../ArticleFV';

export default class ArticleFastView extends Component{
    constructor(props){
        super(props);

        this.state = {
            "articles" : [
                {
                    "id": 1,
                    "head": "test1",
                    "prev_text": "tururu rururu\nururur"
                },
                {
                    "id": 2,
                    "head": "test2",
                    "prev_text": "tururu rururu\nururur"
                },
                {
                    "id": 3,
                    "head": "test3",
                    "prev_text": "tururu rururu\nururur"
                }
            ]
        }
    }
    render = () => this.props.visibility ? 
        <div className="fast_articles_view_component">
            <div className="header">
                Последние статьи
            </div>
            <div className="bodies">
                {this.state.articles.map((value, index) => {
                    return(
                        <ArticleFV val={value} key={index}/>
                    )
                })}
            </div>
        </div>
        :
        null
}