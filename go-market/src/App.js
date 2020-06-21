import React from "react";
import "./App.css";
import { Container } from "semantic-ui-react";
import Symbols from "./Symbols";
import Filters from "./Filters";
import Summary from "./Summary";

function App() {
  return (
    <div>
      <Container>
        <Summary />
        <Filters />
        <Symbols />
      </Container>
    </div>
  );
}

export default App;
