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
    if (this.props.filter.targetPercent !== prevProps.filter.targetPercent) {
      this.getSymbols(this.state.activePage);
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



  getSymbols = (page = 1) => {
    var query = ""
    query += "sortTargetPercent=desc"
    query += "&"
    query += "targetPercent=" + this.props.filter.targetPercent
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

            // const todayHigh = item.summaries && Object.keys(item.summaries).length > 0 && item.summaries['nasdaq'].todayHigh
            // const todayLow = item.summaries && Object.keys(item.summaries).length > 0 && item.summaries['nasdaq'].todayLow
            // const todayVolatility = item.summaries && Object.keys(item.summaries).length > 0 && item.summaries['nasdaq'].todayVolatility

            // const todayDrop = item.infos && Object.keys(item.infos).length > 0 && item.infos['nasdaq'].percentageChange

            // var lastVolumeDelta = 0

            // if (item.histories && Object.keys(item.histories).length > 0 && item.histories['nasdaq'].chart && item.histories['nasdaq'].chart.length > 0) {
            //   var tmpHistory = [...item.histories['nasdaq'].chart];
            //   tmpHistory.sort(function (a, b) {
            //     return a.volume - b.volume
            //   })
            //   var minVolume = tmpHistory[0].volume
            //   var maxVolume = tmpHistory[tmpHistory.length - 1].volume
            //   var lastVolume = item.histories['nasdaq'].chart[tmpHistory.length - 1].volume

            //   lastVolumeDelta = lastVolume / maxVolume
            // }

            // var deltaLast3Days = 0
            // var deltaLast5Days = 0
            // if (item.infos && Object.keys(item.infos).length > 0 && item.histories && Object.keys(item.histories).length > 0 && item.histories['nasdaq'].chart) {
            //   var closePrice3Days = item.histories['nasdaq'].chart[item.histories['nasdaq'].chart.length - 3].close
            //   var closePrice5Days = item.histories['nasdaq'].chart[item.histories['nasdaq'].chart.length - 5].close

            //   deltaLast3Days = ((item.infos['nasdaq'].lastSalePrice - closePrice3Days) / item.infos['nasdaq'].lastSalePrice * 100) //* (closePrice3Days > item.infos['nasdaq'].lastSalePrice ? 1 : -1)
            //   deltaLast5Days = ((item.infos['nasdaq'].lastSalePrice - closePrice5Days) / item.infos['nasdaq'].lastSalePrice * 100) //* (closePrice5Days > item.infos['nasdaq'].lastSalePrice ? 1 : -1)

            // }

            // var targetLast3Days = 0
            // var targetLast5Days = 0
            // if (item.infos && Object.keys(item.infos).length > 0 && item.histories && Object.keys(item.histories).length > 0 && item.histories['nasdaq'].chart) {

            //   var highPrice3Days = Math.max.apply(Math, item.histories['nasdaq'].chart.slice(2, 5).map(i => i.high))
            //   var highPrice5Days = Math.max.apply(Math, item.histories['nasdaq'].chart.slice(0, 5).map(i => i.high))

            //   targetLast3Days = ((highPrice3Days - item.infos['nasdaq'].lastSalePrice) / item.infos['nasdaq'].lastSalePrice * 100) //* (closePrice3Days > item.infos['nasdaq'].lastSalePrice ? 1 : -1)
            //   targetLast5Days = ((highPrice5Days - item.infos['nasdaq'].lastSalePrice) / item.infos['nasdaq'].lastSalePrice * 100) //* (closePrice5Days > item.infos['nasdaq'].lastSalePrice ? 1 : -1)

            // }


            // var limitLast3Days = 0
            // var limitLast5Days = 0
            // if (item.infos && Object.keys(item.infos).length > 0 && item.histories && Object.keys(item.histories).length > 0 && item.histories['nasdaq'].chart) {
            //   var limitLast3Days = Math.min.apply(Math, item.histories['nasdaq'].chart.slice(2, 5).map(i => i.low))
            //   var limitLast5Days = Math.min.apply(Math, item.histories['nasdaq'].chart.slice(2, 5).map(i => i.low))

            //   limitLast3Days = ((item.infos['nasdaq'].lastSalePrice - limitLast3Days) / item.infos['nasdaq'].lastSalePrice * 100) //* (closePrice3Days > item.infos['nasdaq'].lastSalePrice ? 1 : -1)
            //   limitLast5Days = ((item.infos['nasdaq'].lastSalePrice - limitLast5Days) / item.infos['nasdaq'].lastSalePrice * 100) //* (closePrice5Days > item.infos['nasdaq'].lastSalePrice ? 1 : -1)

            // }

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
                        <Grid.Column>
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
    // this.getSymbols();
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
