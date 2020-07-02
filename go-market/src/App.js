import React, { Component } from "react";
import "./App.css";
import { Container, Divider, Header } from "semantic-ui-react";
import Symbols from "./Symbols/Symbols";
import Filters from "./Filters/Filters";
import Summary from "./Summary/Summary";

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      filter: {
        targetPercent: [50, 100],
        deltaPercent: [0, 100],
        sync: false,
      }
    };
  }

  onDeltaPercentChange = (value) => {
    this.setState(state => ({
      filter: {
        ...this.state.filter,
        deltaPercent: value,
      }
    }))
  };

  onTargetPercentChange = (value) => {
    this.setState(state => ({
      filter: {
        ...this.state.filter,
        targetPercent: value,
      }
    }))
  };

  onSyncChange = (value) => {
    this.setState(state => ({
      filter: {
        ...this.state.filter,
        sync: value,
      }
    }))
  }

  render() {
    return (
      <div>
        <Container>
          <Header className="header" as="h2">
            GO Market
          </Header>
          <Divider horizontal>Symmary</Divider>
          <Summary />
          <Divider horizontal>Filters</Divider>
          <Filters filter={this.state.filter} onDeltaPercentChange={this.onDeltaPercentChange} onTargetPercentChange={this.onTargetPercentChange} onSyncChange={this.onSyncChange} />
          <Divider horizontal>Symbols</Divider>
          <Symbols filter={this.state.filter} />
        </Container>
      </div >
    );
  }
}

export default App;
