import React, {Component} from 'react';

import { connect } from 'react-redux';
import * as ArticlesFetchData  from '../../Store/Actions/articles';
import * as ArticleClear from '../../Store/Actions/article';

import LoadingComponent from '../../Components/Loading';
import ErrorBackend from '../../Components/Errorbackend';
// import Article from '../../Components/Article';

import './style.scss';
import 'react-quill/dist/quill.snow.css';

import TagMenu from '../../Components/TagMenu';
import ArticlePreview from '../../Components/ArticlePreview';
import { bindActionCreators } from 'redux';

class ArticlesPageContainer extends Component{

    componentDidMount(){
        this.props.fetchData.ArticlesFetchData();

    }
    render = () => {
        const {isLoading, items, hasError} = this.props;
        console.log("fetched data: ", items);
        return (
            isLoading ?               
            <LoadingComponent />
            :
            hasError ?
            <ErrorBackend />
            :
            <div className="articles_page_container">
                <TagMenu />
                <div className="articles" >
                    {
                        items.articles.map((article, index) => {
                            return(
                                <ArticlePreview article={article} key={index}/>
                            )
                        })
                    }
                </div>
            </div>
        )
    }
}

const mapStateToProps = (state) => {
    return {
        items: state.AI_stateArticles,
        isLoading: state.AI_itemsIsLoading,
        hasError: state.AI_itemsHasError,
    };
};

const mapDispatchToProps = (dispatch) => {
    return {
        fetchData: bindActionCreators(ArticlesFetchData, dispatch),
        removeArticleData: bindActionCreators(ArticleClear, dispatch),
    };
};

export default connect(mapStateToProps, mapDispatchToProps)(ArticlesPageContainer);