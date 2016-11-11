import React from 'react';
import ReactRouter from 'react-router';
import appRouter from './routes.js';
import ReactCookie from 'react-cookie';
import GlobalData from './data/global.js';

var route = ReactCookie.load('requestRoute') || "/";
console.log(route);
const app = {
    start() {
        this.router = ReactRouter.run(appRouter, ReactRouter.HistoryLocation, function(Handler, state) {
            React.render(<Handler />, document.getElementById('mainContainer'));
        });

        this.router.transitionTo(route);
        ReactCookie.remove('requestRoute');
    }
};

export default app;