import React, { Component } from "react";
import { Card, Header, Form, Input, Icon, Grid, Label, Button, Checkbox, Pagination } from "semantic-ui-react";

class Marker extends Component {
    constructor(props) {
        super(props);
        this.state = {
            checked: false
        };
    }



    render() {
        if (this.props.ratings && Object.keys(this.props.ratings).length > 0) {
            const { targetPercent, targetPrice } = this.props.ratings['marketbeat']

            return (
                <div>Rating:
                    <p style={{ fontSize: "8pt" }}>{targetPrice} ({targetPercent}%)</p>
                </div>
            );
        }
        return (<div>Summary:</div>)
    }
}

export default Marker;
