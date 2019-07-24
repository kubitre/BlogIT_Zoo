import React, {Component} from 'react';

import './style.scss';

import { connect } from 'react-redux';

import LoadingComponent from '../Loading';
import ErrorBackend from '../Errorbackend';

import * as ArticleActions from '../../Store/Actions/article';
import { bindActionCreators } from 'redux';
import TagMenu from '../TagMenu';
import Comments from '../Comments';

class Article extends Component{

    componentDidMount(){
        const {data} = this.props.store;
        const {getArticleByID} = this.props.articleFetch;

        if (data === null) {
            getArticleByID(this.props.match.params.id);   
        }
    }

    render = () => {
        console.log(this.props);        
        const {loading, failed, loaded} = this.props.store;
        return(
            loaded ?
                <div className="article_container_with_panel">

                    <TagMenu />
                    <div className="article_container">
                        <div className="header">{this.props.store.data.header}</div>
                        <div className="body" dangerouslySetInnerHTML={{__html: this.props.store.data.content}}/>
                        <div className="comments_block">
                            <Comments />
                        </div>
                    </div>
                </div>
            :

            loading ? 
                <LoadingComponent />
                :
                failed ?
                    <ErrorBackend />
                    :
                    null
        )
    }
}


const mapStateToProps = (state) => {
    return {
        store: state.ASI_stateArticle,
    };
};

const mapDispatchToProps = (dispatch) => {
    return {
        articleFetch: bindActionCreators(ArticleActions, dispatch)
    };
};

export default connect(mapStateToProps, mapDispatchToProps)(Article);
