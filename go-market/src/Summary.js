import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

let endpoint = "http://localhost:8001/api/v1";

class Summary extends Component {
  constructor(props) {
    super(props);

    this.state = {
      task: "",
      symbols: []
    };
  }

  componentDidMount() {
    this.getSymbols();
  }

  getSymbols = () => {
    return (<div>Summary info ...</div>)
  };

  
  render() {
    return (
      <div>
        <div className="row">
          <Header className="header" as="h2">
            GO Market
          </Header>
        </div>
        <div className="row">
          <Card.Group>{this.state.symbols}</Card.Group>
        </div>
      </div>
    );
  }
}

export default Summary;
