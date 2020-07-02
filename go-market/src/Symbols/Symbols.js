import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon, Grid, Label, Button } from "semantic-ui-react";

import CardHeader from "./CardHeader";
import CardPages from "./CardPages";
import Rating from './Rating';
import Summary from './Summary';
import Info from './Info';
import History from './History';
import Markers from './Markers';


let endpoint = "http://localhost:8001/api/v1";

class Symbols extends Component {
  constructor(props) {
    super(props);
    this.state = {
      task: "",
      symbols: [],
      activePage: 1,
      totalPages: 1,
      pageSize: 20,
    };
  }

  componentDidMount() {
    this.getSymbols(this.state.activePage);
  }

  componentDidUpdate(prevProps) {
    if (this.props.filter.targetPercent && this.props.filter.targetPercent[0] !== prevProps.filter.targetPercent[0] || this.props.filter.targetPercent[1] !== prevProps.filter.targetPercent[1]) {
      this.getSymbols(this.state.activePage);
    }

    if (this.props.filter.deltaPercent && this.props.filter.deltaPercent[0] !== prevProps.filter.deltaPercent[0] || this.props.filter.deltaPercent[1] !== prevProps.filter.deltaPercent[1]) {
      this.getSymbols(this.state.activePage);
    }

    if (this.props.filter.sync !== prevProps.filter.sync) {
      this.getSymbols(this.state.activePage, this.props.filter.sync);
    }
  }

  onChange = event => {
    this.setState({
      [event.target.name]: event.target.value
    });
  };


  onTagClick = (symbol, state) => {
    axios
      .post(endpoint + "/filters",
        {
          "symbol": symbol,
          "state": state,
        },
        {
          headers: {
            "Content-Type": "application/json"
          }
        })
      .then(res => {
        console.log(res);
        this.getSymbols(this.state.activePage);
      });
  };



  getSymbols = (page = 1, sync = false) => {
    var query = ""
    query += "sortTargetPercent=desc"
    query += "&"
    query += "targetPercents[]=" + this.props.filter.targetPercent[0] + "&" + "targetPercents[]=" + this.props.filter.targetPercent[1]
    query += "&"
    query += "deltaPercents[]=" + this.props.filter.deltaPercent[0] + "&" + "deltaPercents[]=" + this.props.filter.deltaPercent[1]
    query += "&"
    query += "page=" + page
    query += "&"
    query += "pageSize=" + this.state.pageSize

    axios.get(endpoint + "/symbols?" + query).then(res => {
      console.log(res);
      if (res.data) {
        this.setState({
          totalPages: res.data.totalPages,
          symbols: res.data.symbols.map(item => {
            let color = "yellow";
            if (sync && !res.data.filters.includes(item.symbol)) {
              return
            }

            return (
              <Card key={item.id} color={color} fluid>
                <Card.Content>
                  <CardHeader symbol={item.symbol} onTagClick={this.onTagClick} checked={res.data.filters.includes(item.symbol)} />
                  <Card.Meta textAlign="left">
                    <Grid columns={4} divided>
                      <Grid.Row>
                        <Grid.Column width={3}>
                          <Rating ratings={item.ratings} />
                          <div></div>
                          <Summary summaries={item.summaries} />
                        </Grid.Column>
                        <Grid.Column width={3}>
                          <Info summaries={item.summaries} infos={item.infos} />
                        </Grid.Column>
                        <Grid.Column width={5}>
                          <History histories={item.histories} />
                        </Grid.Column>
                      </Grid.Row>
                    </Grid>
                  </Card.Meta>
                </Card.Content>
                <Markers markers={item.markers} />
              </Card>
            );
          })
        });
      } else {
        this.setState({
          symbols: []
        });
      }
    });
  };

  onReload = (e) => {
    this.getSymbols(this.state.activePage);
  }

  onPageChange = (e, data) => {
    this.getSymbols(data.activePage)

    this.setState({
      activePage: data.activePage
    })
  }

  render() {
    return (
      <div>
        <Grid divided='vertically'>
          <Grid.Row>
            <Grid.Column>
              <div className="row">
                <CardPages onPageChange={this.onPageChange} activePage={this.state.activePage} totalPages={this.state.totalPages} />
                <Button onClick={this.onReload} style={{ margin: "5px 0" }} style={{ float: "right" }}>
                  Reload
              </Button>
              </div>
              <div className="row">
                <Card.Group>{this.state.symbols}</Card.Group>
              </div>
            </Grid.Column>
          </Grid.Row>
        </Grid>
      </div>
    );
  }
}

export default Symbols;
