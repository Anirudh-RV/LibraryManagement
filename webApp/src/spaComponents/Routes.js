import React, { Component } from "react";
import { BrowserRouter, Route, Switch, Redirect } from "react-router-dom";
import Cookies from "universal-cookie";
import { Link } from "react-router-dom";
import axios from "axios";
import Home from "./Home";
import Error from "./Error";

// Routes for the webpages in the project

class Routes extends Component {
  render() {
    const cookies = new Cookies();
    return (
      <BrowserRouter>
        <div>
          <Switch />
          <Switch>
            <Route exact path="/" component={Home} />
            <Route exact path="/index.html" component={Home} />
            <Route component={Error} />
          </Switch>
        </div>
      </BrowserRouter>
    );
  }
}
export default Routes;
