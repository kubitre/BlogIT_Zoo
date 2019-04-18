import React, {Component} from 'react';

import './index.scss';

export default class LogotypeComponent extends Component {
    render = () => {
        return <div className="logotype_component" style={{opacity: this.props.visibility ? 1 : 0.3, transition: 'all .4s'}}>
                    <div className="text_wrapper">
                        <span>В\&lt;IT-ZOO&gt;</span>
                    </div>
                    {this.props.smart_activate ? 
                        <div className="smart_functions"></div>
                        :
                        null
                    }
                </div>
    }
}