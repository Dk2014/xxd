import React from 'react';
import {Route, DefaultRoute, NotFoundRoute} from 'react-router';

import App from './app/index.js';
import Index from './pages/index.js';
import Item from './components/item.js';
import Login from './pages/login/index.js';
import NotFound from './pages/misc/notfound.js';

var appRouter = (
	<Route >
		<Route name="login" path="/login" handler={Login} />
	    <Route handler={App}>
		    <DefaultRoute handler={Index}/>
		    <Route name="item" path=":categoryPath/:itemPath" handler={Item} />
		    <NotFoundRoute handler={NotFound} />
	  	</Route>
  	</Route>
);

export default appRouter;
