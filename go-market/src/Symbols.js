import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon, Grid, Label, Button } from "semantic-ui-react";

let endpoint = "http://localhost:8001/api/v1";


class Symbols extends Component {
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

  componentDidUpdate(prevProps) {
    // // Популярный пример (не забудьте сравнить пропсы):
    // if (this.props.userID !== prevProps.userID) {
    //   this.fetchData(this.props.userID);
    // }
  }

  onChange = event => {
    this.setState({
      [event.target.name]: event.target.value
    });
  };

  onSubmit = () => {
    let { task } = this.state;
    // console.log("pRINTING task", this.state.task);
    if (task) {
      axios
        .post(
          endpoint + "/api/task",
          {
            task
          },
          {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded"
            }
          }
        )
        .then(res => {
          this.getSymbols();
          this.setState({
            task: ""
          });
          console.log(res);
        });
    }
  };

  getSymbols = () => {
    axios.get(endpoint + "/symbols?sort_percent=desc&percent_limit=60").then(res => {
      console.log(res);
      if (res.data) {
        this.setState({
          symbols: res.data.map(item => {
            let color = "yellow";

            // if (item.status) {
            //   color = "green";
            // }

            return (
              <Card key={item.id} color={color} fluid>
                <Card.Content>
                  <Card.Header textAlign="left">
                    <div style={{ wordWrap: "break-word" }}>{item.symbol}</div>
                  </Card.Header>
                  <Card.Meta textAlign="left">

                    <Grid columns={4} divided>
                      <Grid.Row>
                        <Grid.Column>
                          <div>Target:
                            <p style={{ fontSize: "8pt" }}>Percent: {item.ratings && Object.keys(item.ratings).length > 0 && item.ratings['marketbeat'].targetPercent}</p>
                            <p style={{ fontSize: "8pt" }}>Price: {item.ratings && Object.keys(item.ratings).length > 0 && item.ratings['marketbeat'].targetPrice}</p>
                          </div>
                        </Grid.Column>
                        <Grid.Column>
                          <div>Info:
                            <p style={{ fontSize: "8pt" }}>PreviousClose: {item.infos && Object.keys(item.infos).length > 0 && item.infos['nasdaq'].previousClose}</p>
                            <p style={{ fontSize: "8pt" }}>LastSalePrice: {item.infos && Object.keys(item.infos).length > 0 && item.infos['nasdaq'].lastSalePrice}</p>
                            <p style={{ fontSize: "8pt" }}>PercentageChange: {item.infos && Object.keys(item.infos).length > 0 && item.infos['nasdaq'].percentageChange}</p>
                          </div>
                        </Grid.Column>
                        <Grid.Column>
                          <div>Summary:
                            <p style={{ fontSize: "8pt" }}>Sector: {item.summaries && Object.keys(item.summaries).length > 0 && item.summaries['nasdaq'].sector}</p>
                            <p style={{ fontSize: "8pt" }}>Industry: {item.summaries && Object.keys(item.summaries).length > 0 && item.summaries['nasdaq'].industry}</p>
                            <p style={{ fontSize: "8pt" }}>TodayHighLow: {item.summaries && Object.keys(item.summaries).length > 0 && item.summaries['nasdaq'].todayHighLow}</p>
                            <p style={{ fontSize: "8pt" }}>52 wk H/L: {item.summaries && Object.keys(item.summaries).length > 0 && item.summaries['nasdaq'].fiftTwoWeekHighLow}</p>
                            <p style={{ fontSize: "8pt" }}>EPS: {item.summaries && Object.keys(item.summaries).length > 0 && item.summaries['nasdaq'].earningsPerShare}</p>
                          </div>
                        </Grid.Column>
                      </Grid.Row>
                    </Grid>
                  </Card.Meta>
                </Card.Content>
                <Grid style={{ margin: "2px" }}>
                  <Grid.Row columns={8}>
                    <Grid.Column>
                      <Icon
                        name="arrow circle down"
                        color="red"
                        onClick={() => this.updateTask(item.id)}
                      />
                      <span>Text</span>
                    </Grid.Column>
                    <Grid.Column>
                      {
                        item.summaries && Object.keys(item.summaries).length > 0 && item.summaries['nasdaq'].earningsPerShare > -10 ?
                          <Icon
                            name="thumbs up outline"
                            color="green"
                          />
                          :
                          <Icon
                            name="thumbs down outline"
                            color="red"
                          />
                      }
                      <span>EPS({item.summaries && Object.keys(item.summaries).length > 0 && item.summaries['nasdaq'].earningsPerShare})</span>
                    </Grid.Column>
                    <Grid.Column>
                      {
                        item.infos && Object.keys(item.infos).length > 0 && item.infos['nasdaq'].deltaIndicator === "up" ?
                          (<Icon
                            name="arrow circle up"
                            color="green"
                          />)
                          : item.infos && Object.keys(item.infos).length > 0 && item.infos['nasdaq'].deltaIndicator === "down" ?
                            (<Icon
                              name="arrow circle down"
                              color="red"
                            />)
                            :
                            (<Icon
                              name="pause circle"
                              color="yellow"
                            />)
                      }
                      <span>{item.infos && Object.keys(item.infos).length > 0 && item.infos['nasdaq'].percentageChange}</span>
                    </Grid.Column>
                  </Grid.Row>
                </Grid>

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

  updateTask = id => {
    axios
      .put(endpoint + "/api/task/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log(res);
        this.getSymbols();
      });
  };

  undoTask = id => {
    axios
      .put(endpoint + "/api/undoTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log(res);
        this.getSymbols();
      });
  };

  deleteTask = id => {
    axios
      .delete(endpoint + "/api/deleteTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log(res);
        this.getSymbols();
      });
  };

  onReload = (e) => {
    this.getSymbols();
  }

  render() {
    return (
      <div>
        <Grid divided='vertically'>
          <Grid.Row>
            <Grid.Column>
              <Button onClick={this.onReload} style={{ margin: "5px 0" }}>
                Reload
              </Button>
              <div className="row">
                <Card.Group>{this.state.symbols}</Card.Group>
              </div>
            </Grid.Column>
          </Grid.Row>
        </Grid>

        {/* <div className="row">
          <Header className="header" as="h2">
            GO Market
          </Header>
        </div> 
        <div className="row">
          <Form onSubmit={this.onSubmit}>
            <Input
              type="text"
              name="task"
              onChange={this.onChange}
              value={this.state.task}
              fluid
              placeholder="Create Task"
            />
          </Form>
        </div> */}

      </div>
    );
  }
}

export default Symbols;
