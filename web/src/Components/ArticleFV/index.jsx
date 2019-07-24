import React, {Component} from 'react';

export default class ArticleFV extends Component{
    render = () => {
        const {head, prev_text} = this.props.val;
        return(
            <div className="article_fast_view">
                <div className="header">
                    {head}
                </div>
                <div className="img_">
                </div>
                <div className="body_prev">
                    {prev_text}
                </div>
            </div>
        )
    }
}