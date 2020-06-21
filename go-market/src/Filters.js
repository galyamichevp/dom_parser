import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon, Button, Grid, Select, Dropdown, Label } from "semantic-ui-react";

let endpoint = "http://localhost:8001/api/v1";

class Filters extends Component {
  constructor(props) {
    super(props);

    this.state = {
      task: "",
      symbols: [],
      selected: null
    };
  }

  componentDidMount() {
    this.getSymbols();
  }

  getSymbols = () => {
    axios.get(endpoint + "/filters").then(res => {
      console.log(res);
      if (res.data) {
        this.setState({
          symbols: res.data.map(item => {
            return {
              key: item.symbol,
              value: item.symbol,
              text: item.symbol
            }
          })
        });
      } else {
        this.setState({
          symbols: []
        });
      }
    });
  };


  setSymbols = (symbols) => {
    axios
      .post(endpoint + "/filters",
        {
          "symbols": symbols
        },
        {
          headers: {
            "Content-Type": "application/json"
          }
        })
      .then(res => {
        console.log(res);
        this.getSymbols();
      });
  };

  onChange = (e, data) => {
    console.log(data.value);
    this.setState({ selected: data.value });
    this.setSymbols(data.value)
  }

  onReload = (e) => {
    this.getSymbols();
  }


  render() {
    const { selected } = this.state;

    return (
      <div>
        <Grid divided='vertically'>
          <Grid.Row columns={3}>
            <Grid.Column width={1}>
              <Label style={{ margin: "5px" }}>
                Filters
              </Label>
            </Grid.Column>
            <Grid.Column>
              <Dropdown
                placeholder='Symbols'
                fluid
                multiple
                search
                selection
                options={this.state.symbols}
                onChange={this.onChange}
                value={selected}
              />
            </Grid.Column>
            <Grid.Column>
              <Button icon onClick={this.onReload}>
                <Icon name='refresh' />
              </Button>
            </Grid.Column>
          </Grid.Row>
        </Grid>
      </div>
    );
  }
}

export default Filters;
