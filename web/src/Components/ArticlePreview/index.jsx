import React, {Component} from 'react';
import {NavLink} from 'react-router-dom';

import './style.scss'

export default class ArticlePreview extends Component{
    render = () => {
        const {header, preview, amountMinutesForRead, id} = this.props.article;
        return(
            <div className="article_preview_container">
                <div className="header">
                    <div className="name">{header}</div>
                    
                </div>
                <div className="preview_image">
                    {/* <h2>Здесь картинка</h2> */}
                    <div className="article_readable">
                    {/* Amount hours for reading over article */}
                        <h2>Amount minutes for reading:</h2>
                        <h2 className="hours_">5 hours</h2>
                    </div>
                </div>
                <div className="preview_box">{preview}</div>
                <div className="article_info">
                    {/* comments */}
                    <div className="comments" />
                    {/* Views amount */}
                    <div className="views" />
                    <div className="likes" />
                    {/* tags */}
                    <div className="tags"/>
                </div>
                <div className="linkOpen">
                    <NavLink to={"/article/" + id}>Read</NavLink>
                </div>
            </div>
        )
    }
}