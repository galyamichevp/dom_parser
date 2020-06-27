import React, { Component } from "react";
import { Card, Header, Form, Input, Icon, Grid, Label, Button, Checkbox, Pagination } from "semantic-ui-react";

class Summary extends Component {
    constructor(props) {
        super(props);
        this.state = {
            checked: false
        };
    }



    render() {
        if (this.props.summaries && Object.keys(this.props.summaries).length > 0) {

            const { sector, industry } = this.props.summaries['nasdaq']

            return (
                <div>Summary:
                    <p style={{ fontSize: "8pt" }}>Sector: {sector}</p>
                    <p style={{ fontSize: "8pt" }}>Industry: {industry}</p>
                    {/* 
                <div>Summary:
                            <p style={{ fontSize: "8pt" }}>Sector: {item.summaries && Object.keys(item.summaries).length > 0 && item.summaries['nasdaq'].sector}</p>
                    <p style={{ fontSize: "8pt" }}>Industry: {item.summaries && Object.keys(item.summaries).length > 0 && item.summaries['nasdaq'].industry}</p>
                    <p style={{ fontSize: "8pt" }}>TodayHighLow: {todayHigh}/{todayLow}</p>
                    <p style={{ fontSize: "8pt" }}>52 wk H/L: {item.summaries && Object.keys(item.summaries).length > 0 && item.summaries['nasdaq'].fiftTwoWeekHighLow}</p>
                    <p style={{ fontSize: "8pt" }}>EPS: {item.summaries && Object.keys(item.summaries).length > 0 && item.summaries['nasdaq'].earningsPerShare}</p>
                </div> */}
                </div>
            );
        }
        return (<div>Summary:</div>)
    }
}

export default Summary;
