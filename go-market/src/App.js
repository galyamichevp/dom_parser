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
        targetPercent: 50
      }
    };
  }

  onTargetPercentChange = (value) => {
    console.log(" ==> " + value)
    this.setState(state => ({
      filter: {
        targetPercent: value
      }
    }))
  };

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
          <Filters filter={this.state.filter} onTargetPercentChange={this.onTargetPercentChange} />
          <Divider horizontal>Symbols</Divider>
          <Symbols filter={this.state.filter} />
        </Container>
      </div >
    );
  }
}

export default App;
