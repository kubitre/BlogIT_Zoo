import React, {Component} from 'react';

import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import * as LoginActions from '../../Store/Actions/user';
import * as CommentsActions from '../../Store/Actions/comments';

import moment from 'moment';

import './style.scss';

class Comments extends Component{
    constructor(props){
        super(props);

        this.state = {
            "value": "",
            "fetching_comments": false,
        };

        this.handleInput = this.handleInput.bind(this);
    }
    componentDidMount(){
        const {fetchDataFromLS} = this.props.actionsForLogin;
        const {loaded} = this.props.userStore;
        const {startFetchCommentsFromBackend} = this.props.commentsActions;
        const {id} = this.props.article;

        if (loaded){
            
        }
        else{
            fetchDataFromLS();
        }
    }

    handleInput = (event) => {
        console.log(event.target.value);
        this.setState({
            value: event.target.value
        });
    }

    handleSend = (event) => {
        event.preventDefault();

        const {userid, token} = this.props.userStore;
        const {id} = this.props.article.data;

        console.log("start sending")
        if (this.state.value != "") {
            this.props.commentsActions.createNewComment({
                id_author: userid,
                body: this.state.value,
                id_article: id
            },token
            )
        }
    }

    render = () => {
        const {loaded} = this.props.userStore;
        const {startFetchCommentsFromBackend} = this.props.commentsActions;
        const {id} = this.props.article.data;
        const loaded_comments = this.props.commentsStore.loaded;

        const loaded_comments_fetching = this.props.commentsStore.loaded;
        const datas = this.props.commentsStore.data;
        
        console.log("loaded user state: ", loaded, " loaded comments state: ", loaded_comments, " id article: ", id);

        if (!this.state.fetching_comments){
            startFetchCommentsFromBackend(id);    
            console.log("start fetch comments" );
            this.setState({
                fetching_comments: true
            });
        }

        return (
            <div className="comments_container">
                <div className="header_comments">Comments</div>
                {loaded ?
                    <div className="comments_body">
                        <div className="values">
                            {
                                loaded_comments_fetching ?
                                
                                datas != null ?
                                datas.map((comment, index) => 
                                    <div className="comment_container" key={index}>
                                        <div className="left">
                                            <div className="avatar"
                                                style={{background: `url(${comment.author.avatar})`,
                                                        backgroundSize: `cover`,
                                                        backgroundRepeat: `no-repeat`,
                                                        borderRadius: `50px`,
                                                        width: `50px`,
                                                        height: `50px`,
                                            }}
                                            />
                                        </div>
                                        <div className="right">
                                            <div className="autor_with_time">
                                                <div className="author">{comment.author.username}</div>                                    
                                                <div className="created">{moment(comment.createdat).format("LLLL")}</div>
                                            </div>
                                            <div className="body">{comment.body}</div>
                                        </div>
                                    </div>
                                )
                                :
                                null

                                :
                                null
                            }
                        </div>
                        <div className="add_new_comment">
                            <div className="header">Добавить новый комментарий:</div>
                            <div className="input_block">
                                <input className="inputField" onChange={this.handleInput}/>
                                <div className="buttonSend" onClick={this.handleSend} style={{background: "#d53"}}>Send</div>
                            </div>

                        </div>

                    </div>
                    :
                    <div className="no_visible_comments">Comments are not available not login users!</div>
                }
            </div>
        )
    }
}

const mapStateToProps = (state) => {
    return {
        userStore: state.USI_userState,
        commentsStore: state.CSI_commentsState,
        article: state.ASI_stateArticle,
    };
};

const mapDispatchToProps = (dispatch) => {
    return {
        actionsForLogin: bindActionCreators(LoginActions, dispatch),
        commentsActions: bindActionCreators(CommentsActions, dispatch)
    };
};

export default connect(mapStateToProps, mapDispatchToProps)(Comments);