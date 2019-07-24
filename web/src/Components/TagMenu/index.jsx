import React, {Component} from 'react';

import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import './style.scss';

import * as TagMenuActions from '../../Store/Actions/tags';

class TagMenu extends Component{
    componentDidMount() {
        console.log(this.props);
        const {startFetchingTagsFromBackend} = this.props.actionsForTags;
        const {loaded} = this.props.tags;
        if (!loaded) {
            startFetchingTagsFromBackend();
        }
    }
    render = () => {
        const {loaded, data} = this.props.tags;
        return (
            loaded ?

            <div className="tagmenu_container">
                {
                    data.map((tag, index) => 
                        index <=5 ?
                        <div className="tag_entity" id={tag.id} key={index}>{tag.tagname}</div>   
                        :
                        null
                    )
                }
            </div>

            :
            <div className="no_loaded"></div>
        )
    }
}

const mapStateToProps = (state) => {
    return {
        tags: state.TSI_tagState,
    };
};

const mapDispatchToProps = (dispatch) => {
    return {
        actionsForTags: bindActionCreators(TagMenuActions, dispatch)
    };
};

export default connect(mapStateToProps, mapDispatchToProps)(TagMenu);