import React, { Component } from "react";
import { Card, Header, Form, Input, Icon, Grid, Label, Button, Checkbox, Pagination } from "semantic-ui-react";

class Info extends Component {
    constructor(props) {
        super(props);
        this.state = {
            checked: false
        };
    }



    render() {
        if (this.props.summaries && Object.keys(this.props.summaries).length > 0 &&
            this.props.infos && Object.keys(this.props.infos).length > 0) {

            const { todayHigh, todayLow, fiftTwoWeekHighLow, earningsPerShare } = this.props.summaries['nasdaq']
            const { previousClose, lastSalePrice, percentageChange } = this.props.infos['nasdaq']

            return (
                <div>Info:
                    <p style={{ fontSize: "8pt" }}>H/L: {todayHigh} / {todayLow}</p>
                    <p style={{ fontSize: "8pt" }}>52wk: {fiftTwoWeekHighLow}</p>
                    <p style={{ fontSize: "8pt" }}>EPS: {earningsPerShare}</p>
                    <p style={{ fontSize: "8pt" }}>Previious Close: {previousClose}</p>
                    <p style={{ fontSize: "8pt" }}>Sale Price: {lastSalePrice}</p>
                    <p style={{ fontSize: "8pt" }}>Percentage Change: {percentageChange}</p>
                </div>
            );
        }
        return (
            <div>Info:</div>)
    }
}

export default Info;
