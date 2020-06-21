import React, { Component } from "react";
import "./App.css";
import { Container } from "semantic-ui-react";
import Symbols from "./Symbols";
import Filters from "./Filters";
import Summary from "./Summary";

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      filter: {
        targetPercentLimit: 100
      }
    };
  }

  render() {
    return (
      <div>
        <Container>
          <Summary />
          <Filters />
          <Symbols filter={this.state.filter} />
        </Container>
      </div >
    );
  }
}

export default App;
