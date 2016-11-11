import React from 'react';
import Request from 'superagent';
import GlobalData from '../../data/global.js';
import httpUtil from '../../util/http.js';

let PatchPage = React.createClass({
	getInitialState: function() {
		return {patches: {}, patchUrl:""}
	},

	componentDidMount: function() {
		this.refreshPatch();
		this.refreshPatchUrl();
	},

	handlePatchUrlChange: function(e) {
		this.state.patchUrl = e.target.value;
		this.setState(this.state);
	},

	refreshPatch: function() {
		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let type = GlobalData.App.substring(index+1, GlobalData.App.length);
		Request.get('/api/patches?app=' + app + '&type=' + type)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				console.dir(val);
				let patches = val.data;
				self.setState({patches: patches});
			});
	},

	refreshPatchUrl: function() {
		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		Request.get('/api/getPatchUrl?app=' + app)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				console.dir(val);
				self.setState({patchUrl: val.data});
			});
	},

	handleAddPatch: function(e) {
		let client_version = React.findDOMNode(this.refs.client_version).value
		let server_version = React.findDOMNode(this.refs.server_version).value
		let patch_path = React.findDOMNode(this.refs.patch_path).value

		if (client_version === '' || server_version==='' || patch_path==='') {
			alert('不能为空');
			return;
		}

		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let type = GlobalData.App.substring(index+1, GlobalData.App.length);
		let requestPath = '/api/editPatch?app=' + app + '&type=' + type + 
			'&client_version=' + client_version + '&server_version=' 
			+ server_version + '&path=' + patch_path;
		Request.get(requestPath)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				console.dir(val);
				if (val.code === 0) {
					self.refreshPatch();
				}
			});
	},

	handleEditPatch: function(version) {
		var patch = this.state.patches[version];
		if (!patch) {
			console.log("error, version: " + version);
			return
		}

		React.findDOMNode(this.refs.client_version).value = version.split("_")[0]
		React.findDOMNode(this.refs.server_version).value = version.split("_")[1]
		React.findDOMNode(this.refs.patch_path).value = patch
	},

	handleDelPatch: function(version) {
		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let type = GlobalData.App.substring(index+1, GlobalData.App.length);
		let requestPath = '/api/editPatch?app=' + app + '&type=' + type + 
			'&version=' + version + '&path=';
		Request.get(requestPath)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				console.dir(val);
				if (val.code === 0) {
					self.refreshPatch();
				}
			});
	},

	handleEditPatchUrl: function(e) {
		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let requestPath = '/api/editPatchUrl?app=' + app +
			'&patchurl=' + this.state.patchUrl;
		Request.get(requestPath)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				console.dir(val);
				if (val.code === 0) {
					alert('ok');
				}
			});
	},

	render: function() {
		let self = this;
		return (
			<div>
				<div className="form-inline">
					<label className="label-control col-sm-1">baseurl：</label>
					<input className="form-control col-sm-9" type="text" name="patch_url" value={this.state.patchUrl} onChange={this.handlePatchUrlChange}/>
					<button className="btn-primary form-control" type="button" onClick={this.handleEditPatchUrl}>修改</button>
				</div>
				<br></br>
				<table className="table table-bordered table-responsive">
					<thead>
						<tr className="success">
							<th>版本标识</th>
							<th>patch路径</th>
							<th>操作</th>
						</tr>
					</thead>
					<tbody>
					{Object.keys(this.state.patches).map(function(key) {
						return (
							<tr key={key}>
								<td>{key}</td>
								<td>{self.state.patches[key]}</td>
								<td>
								<button className="btn-primary" type="button" onClick={function(e){self.handleEditPatch(key)}}>编辑</button>
								------
								<button className="btn-primary" type="button" onClick={function(e){self.handleDelPatch(key)}}>删除</button></td>
							</tr>
							);
					})}
					</tbody>
				</table>

				<div className="form-inline">
					<p >配置patch：</p>
					<input className="form-control" type="number" name="client_version" ref="client_version" placeholder="客户端版本号" />
					<span>_</span>
					<input className="form-control" type="number" name="server_version" ref="server_version" placeholder="服务端期待版本号" />
					<span className="label-control">:</span>
					<input className="form-control" type="text" name="patch" ref="patch_path" placeholder="patch文件" />
					<button className="btn-primary" type="button" onClick={this.handleAddPatch}>添加</button>
				</div>
			</div>
			);
	}

});


export default PatchPage;