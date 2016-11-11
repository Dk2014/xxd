import React from 'react';
import {RouteHandler} from 'react-router';
import Sidebar from '../components/sidebar';
import ItemData from '../data/item.js';

var flexContainerStyle = {
  display: 'flex',
  flexFlow: 'row nowrap',
  height: "1000px"
};

var flexSideStyle = {
  flex: '1 1 auto'
};

var flexContentStyle = {
  flex: '2 1 auto'
};

var App = React.createClass({
  contextTypes: {
    router: React.PropTypes.func
  },

  render: function () {
    let activeCategory = this.context.router.getCurrentParams().categoryPath;
    let activeItem = this.context.router.getCurrentParams().itemPath;

    return (
      <div style={flexContainerStyle}>
        <Sidebar style={flexSideStyle} activeCategory={activeCategory} activeItem={activeItem} categories={ItemData.getAll()}/>
        <div style={flexContentStyle}>
          <RouteHandler/>
        </div>
      </div>
    );
  }
});
export default App;
