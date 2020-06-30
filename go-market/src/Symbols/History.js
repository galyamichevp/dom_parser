import React, { Component } from "react";
import { Card, Header, Form, Input, Icon, Grid, Label, Button, Checkbox, Pagination } from "semantic-ui-react";

class History extends Component {
    constructor(props) {
        super(props);
        this.state = {
            checked: false
        };
    }



    render() {
        if (this.props.histories && Object.keys(this.props.histories).length > 0) {

            const { chart } = this.props.histories['nasdaq']

            return (
                <div>History:
                    {
                        chart && chart.map(h => {
                            return (
                                <p key={h.dateTime} style={{ fontSize: "8pt" }}>{new Date(h.dateTime).toLocaleDateString()}: O:{h.open}/H:{h.high}/L:{h.low}/C:{h.close}/V:{h.volume}</p>)
                        })
                    }
                </div>
            );
        }
        return (<div>History:</div>)
    }
}

export default History;
