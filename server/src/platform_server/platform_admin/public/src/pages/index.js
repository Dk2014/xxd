import React from 'react';
import {Navigation} from 'react-router';
import PageHeader from '../components/pageHeader.js';

const Index = React.createClass({
  	mixins: [Navigation],

    componentDidMount: function() {
  		this.transitionTo('/server/list');
  	},

  render: function () {
    return (
      <div>
      	<PageHeader params={{category:"", item:""}} />
      	<h> welcome xxd platform admin</h>
      </div>
    );
  }
});

export default Index;


