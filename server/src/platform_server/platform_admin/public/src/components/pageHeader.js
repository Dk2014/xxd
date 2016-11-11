import React from 'react';

var item_header = {
  display: 'flex',
  backgroundColor: 'gray',
  flexFlow: 'row nowrap'
};

var item_head_left = {
  flex: '1 1 auto'
};

var item_head_right = {
  flex: '2 1 auto'
};

var item_head_right_fex = {
  paddingTop: '8',
  float: 'right'
};

let Header = React.createClass({
	contextTypes: {
	    router: React.PropTypes.func
	},

	render: function() {
		let params = this.props.params;
		return (
			<div claseName="item-header" style={item_header}>
	          <div style={item_head_left }>
	           {!!params.category && !!params.item ? <h5>{params.category} / {params.item}</h5> : null}
	          </div>
	          <div style={item_head_right}>
	            <a style={item_head_right_fex} href="/api/logout">退出</a>
	          </div>
	        </div>
		);
	}
});

export default Header;