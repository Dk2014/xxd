import React from 'react';
import Router from 'react-router';
import {Route, DefaultRoute, RouteHandler, Link, Navigation, State } from 'react-router';
import GlobalData from '../data/global.js';

let XXDLogoStyle = {
  marginLeft: '55'
};

let CategoryStyle = {
  margin: '15px 20px',
  color: 'white'
};

let sideBarStyle = {
  background: '#22222c',
  margin: '0 20px 20px 0',
  width: '200px',
  cursor: 'pointer'
};


let CategoryNav = React.createClass({
  getInitialState: function () {
    return { isOpen: this.props.defaultIsOpen};
  },

  getDefaultProps: function () {
    return { isOpen: true };
  },

  componentWillReceiveProps: function (newProps) {
    if (!this.state.isOpen)
      this.setState({ isOpen: newProps.defaultIsOpen });
  },

  toggle: function () {
    this.setState({ isOpen: !this.state.isOpen });
  },

  buildToggleClassName: function () {
    let toggleClassName = 'CategoryNav__Toggle';
    if (this.state.isOpen)
      toggleClassName += ' CategoryNav__Toggle--is-open';
    return toggleClassName;
  },

  renderItems: function () {
    let category = this.props.category;
    return this.state.isOpen ? category.items.map(function (item) {
      let params = { name: item.name, category: category.name, item: item.name, categoryPath: category.path, itemPath: item.path};
      return (
        <li key={item.name}>
          <Link to="item" activeStyle={{color: 'red'}} params={params} >{item.name}</Link>
        </li>
      );
    }) : null;
  },

  render: function () {
    let category = this.props.category;
    return (
      <div className="CategoryNav">
        <h5 className={this.buildToggleClassName()} style={CategoryStyle} onClick={this.toggle}>{category.name}</h5>
        <ul>{this.renderItems()}</ul>
      </div>
    );
  }
});

let Sidebar = React.createClass({
  mixins: [Navigation,State],

  getInitialState: function () {
    return {App: GlobalData.App};
  },

  handleAppChange: function(e) {
    let app = e.target.value;
    this.state.App = app;
    GlobalData.App = app;
    this.setState(this.state);

    this.replaceWith("/");
  },

  renderCategory: function (category) {
    return <CategoryNav
      key={category.path}
      defaultIsOpen={category.path === this.props.activeCategory}
      category={category}
    />;
  },

  render: function () {
    return (
      <div className="Sidebar" style={sideBarStyle} >
        <div>
          <div style={XXDLogoStyle}>
            <img src="/dist/logo.jpg" />
          </div>
          <br></br>
          <select className="form-control col-sm-1" value={this.state.App} onChange={this.handleAppChange}>
            <option value="xxd_vn_1">越南IOS</option>
            <option value="xxd_vn_17">越南安卓</option>
            <option value="xxd_tw_1">台湾IOS</option>
            <option value="xxd_tw_17">台湾安卓</option>
          </select>
        </div>
        <br></br>
        <br></br>
        {this.props.categories.map(this.renderCategory)}
      </div>
    );
  }
});

export default Sidebar;
