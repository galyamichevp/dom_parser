import React, { Component } from "react";
import { Card, Header, Form, Input, Icon, Grid, Label, Button, Checkbox } from "semantic-ui-react";

class CardHeader extends Component {
  constructor(props) {
    super(props);
    this.state = {
      checked: props.checked
    };
  }

  onClick = () => {
    this.props.onTagClick(this.props.symbol, !this.state.checked)

    this.setState(state => ({
      checked: !state.checked
    }));
  }

  render() {
    return (
      <Card.Header textAlign="left">
        <Label as='a' color={this.state.checked ? 'green' : 'red'} ribbon onClick={this.onClick}>
          {this.props.symbol}
        </Label >
      </Card.Header>
    );
  }
}

export default CardHeader;
