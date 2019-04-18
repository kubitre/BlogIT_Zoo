import React, {Component} from 'react';

import './index.scss';

export default class FooterComponent extends Component{
    render = () => 
        <div className="footer_component" style={{opacity: this.props.visibility ? 1 : 0.3, transition: 'all .5s'}}>
            © 2018 All rights reserved. | IT-Zoo blog. Design by kubitre
        </div>
}