import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon, Button, Grid, Select, Dropdown, Label } from "semantic-ui-react";
import { RadialBarChart, RadialBar, Legend, PieChart, Pie } from "recharts"

let endpoint = "http://localhost:8001/api/v1";

const targetPrice = 300
const currentPrice = 100
const data01 = [{ name: 'Group A', value: targetPrice, fill: '#8dd1e1' }]

const data02 = [{ name: 'A1', value: currentPrice, fill: '#a4de6c' },
{ name: 'A2', value: targetPrice - currentPrice, fill: '#ffc658' }]


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
      <Grid divided='vertically'>
        <Grid.Row columns={1}>
          <Grid.Column>
            <Header className="header" as="h2">
              GO Market
          </Header>
          </Grid.Column>
        </Grid.Row>
        <Grid.Row columns={1}>
          <Grid.Column>
            <PieChart width={300} height={300}>
              <Pie data={data01} cx={150} cy={150} outerRadius={60} fill="#8884d8"  />
              <Pie data={data02} cx={150} cy={150} innerRadius={70} outerRadius={90} fill="#82ca9d" label />
            </PieChart>
          </Grid.Column>
        </Grid.Row>
      </Grid>

    );
  }
}

export default Summary;
