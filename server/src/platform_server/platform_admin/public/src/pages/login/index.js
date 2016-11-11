import React from 'react';
import Radium from 'radium';
import request from 'superagent';
import App from '../../app.jsx';
import MD5 from './md5.js';


let flexItemStyle = {
	flex: '0 1 auto',
	alignSelf: 'center'
};

let flexContainerStyle = {
	display: 'flex',
	flexDirection: 'column',
	alignContent: 'center'
};

//for test radium
let redStyle = {
	color: 'red'
};


var page = React.createClass({
	getInitialState: function() {
		return {user: '', password: ''}
	},

	onUserChange: function(e) {
		this.setState({user:e.target.value})
	},

	onPasswordChange: function(e) {
		this.setState({password:e.target.value})
	},

	handleLoginClick: function(e) {
		let md5_psw = MD5(this.state.password);
		
		request
			.post('/api/login')
			.send(JSON.stringify({user: this.state.user,password: md5_psw}))
			.set('Accept', 'application/json')
			.end(function(err, data) {
	        	if (!err && data.body.code === 0) {
	        		App.router.transitionTo('/');
	        	} else {
		        	alert('failed');
	        	}
	    	});
	},

    render: function() {
    	return (
			<div className='form-login' style={flexContainerStyle}>
				<div style={flexItemStyle}>
					<label>用户名：</label>
					<input type='text' className='form-control' value={this.state.user}  onChange={this.onUserChange}/>
				</div>
				<br />
				<div style={flexItemStyle}>
					<label>密码：</label>
					<input type='password' className='form-control' value={this.state.password} onChange={this.onPasswordChange}/>
				</div>
				<br />
				<div style={flexItemStyle}>
					<button className='btn btn-primary' type='button' onClick={this.handleLoginClick}>登录</button>
				</div>
			</div>
		);
    }
});

//page = Radium.Enhancer(page);

export default page;