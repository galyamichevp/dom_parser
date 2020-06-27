import React, { Component } from "react";

class SummarySyncSlider extends Component {
  render() {
    return (
      <div class="ui slider checkbox" style={{ padding: "5px 0" }}>
        <input type="checkbox" name={this.props.name} onChange={this.props.onChange} />
        <label style={{ fontSize: "8pt" }}>{this.props.title}</label>
      </div>
    );
  }
}

export default SummarySyncSlider;
