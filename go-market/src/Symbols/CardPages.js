import React, { Component } from "react";
import { Card, Header, Form, Input, Icon, Grid, Label, Button, Checkbox, Pagination } from "semantic-ui-react";

class CardPages extends Component {
    constructor(props) {
        super(props);
        this.state = {
            checked: false
        };
    }

    onPageChange = (e, data) => {
        this.props.onPageChange(e, data)
    }

    render() {
        return (
            <Pagination
                defaultActivePage={1}
                firstItem={null}
                lastItem={null}
                pointing
                secondary
                totalPages={this.props.totalPages}
                onPageChange={this.onPageChange}
            />
        );
    }
}

export default CardPages;
