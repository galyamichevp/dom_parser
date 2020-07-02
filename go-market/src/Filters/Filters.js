import React, { Component } from "react";
import axios from "axios";
import { Slider } from "react-semantic-ui-range";
import { Card, Header, Form, Input, Icon, Button, Grid, Select, Dropdown, Label, Checkbox } from "semantic-ui-react";

let endpoint = "http://localhost:8001/api/v1";

class Filters extends Component {
  constructor(props) {
    super(props);

    this.state = {
      task: "",
      symbols: [],
      selected: [],//null
      targetPercent: [...this.props.filter.targetPercent],
      deltaPercent: [...this.props.filter.deltaPercent],
      fromSliderSettings: {
        start: this.props.filter.targetPercent[0],
        min: 0,
        max: 600,
        step: 10,
        onChange: value => {
          var range = [...this.state.targetPercent]
          range[0] = value
          if (range[0] > range[1]) {
            range[1] = range[0]
          }

          this.onReload(range)
        }
      },
      toSliderSettings: {
        start: this.props.filter.targetPercent[1],
        min: 0,
        max: 600,
        step: 10,
        onChange: value => {
          var range = [...this.state.targetPercent]
          range[1] = value
          if (range[0] > range[1]) {
            range[0] = range[1]
          }

          this.onReload(range)
        }
      },
      fromDeltaSliderSettings: {
        start: this.props.filter.deltaPercent[0],
        min: -200,
        max: 200,
        step: 1,
        onChange: value => {
          var range = [...this.state.deltaPercent]
          range[0] = value
          if (range[0] > range[1]) {
            range[1] = range[0]
          }

          this.onDeltaReload(range)
        }
      },
      toDeltaSliderSettings: {
        start: this.props.filter.deltaPercent[1],
        min: -200,
        max: 200,
        step: 1,
        onChange: value => {
          var range = [...this.state.deltaPercent]
          range[1] = value
          if (range[0] > range[1]) {
            range[0] = range[1]
          }

          this.onDeltaReload(range)
        }
      },
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


  onChange = (e, data) => {
    this.setState({ selected: data.value });
    this.setSymbols(data.value)
  }

  onReload = (range) => {
    this.props.onTargetPercentChange([...range])

    this.setState({
      targetPercent: [...range]
    });
  }

  onDeltaReload = (range) => {
    this.props.onDeltaPercentChange([...range])

    this.setState({
      deltaPercent: [...range]
    });
  }

  onSync = (e, { checked }) => {
    this.props.onSyncChange(checked)
  }

  render() {
    const { selected, fromSliderSettings, toSliderSettings, targetPercent } = this.state;
    const { fromDeltaSliderSettings, toDeltaSliderSettings, deltaPercent } = this.state;

    return (
      <div>
        <Grid divided='vertically'>
          <Grid.Row columns={5}>
            <Grid.Column width={1}>
              <Label color="red">{targetPercent[0]}</Label>
            </Grid.Column>
            <Grid.Column width={4}>
              <Slider value={targetPercent[0]} color="red" settings={fromSliderSettings} />
            </Grid.Column>
            <Grid.Column width={1}>
              <Label color="red">{targetPercent[1]}</Label>
            </Grid.Column>
            <Grid.Column width={4}>
              <Slider value={targetPercent[1]} color="red" settings={toSliderSettings} />
            </Grid.Column>
            <Grid.Column>
              <Checkbox label='Sync only' onChange={this.onSync} />
            </Grid.Column>
          </Grid.Row>
        </Grid>

        <Grid divided='vertically'>
          <Grid.Row columns={5}>
            <Grid.Column width={1}>
              <Label color="red">{deltaPercent[0]}</Label>
            </Grid.Column>
            <Grid.Column width={4}>
              <Slider value={deltaPercent[0]} color="red" settings={fromDeltaSliderSettings} />
            </Grid.Column>
            <Grid.Column width={1}>
              <Label color="red">{deltaPercent[1]}</Label>
            </Grid.Column>
            <Grid.Column width={4}>
              <Slider value={deltaPercent[1]} color="red" settings={toDeltaSliderSettings} />
            </Grid.Column>
            <Grid.Column>
              {/* <Checkbox label='Sync only' onChange={this.onSync} /> */}
            </Grid.Column>
          </Grid.Row>
        </Grid>
      </div>
    );
  }
}

export default Filters;
