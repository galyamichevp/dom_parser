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
        if (this.props.markers && Object.keys(this.props.markers).length > 0) {
            const { delta, delta2, delta3, delta4, delta5 } = this.props.markers
            const { deviation, volatility } = this.props.markers
            const { eps, range52, targetPercent, peRation } = this.props.markers

            return (
                <div>
                    <Grid style={{ margin: "2px" }}>
                        <Grid.Row columns={6}>
                            <Grid.Column>
                                {
                                    delta && delta.bValue ?
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
                                <span>Delta({delta && delta.fValue.toFixed(2)})</span>
                            </Grid.Column>
                            <Grid.Column>
                                {
                                    delta2 && delta2.bValue ?
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
                                <span>Delta2({delta2 && delta2.fValue.toFixed(2)})</span>
                            </Grid.Column>
                            <Grid.Column>
                                {
                                    delta3 && delta3.bValue ?
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
                                <span>Delta3({delta3 && delta3.fValue.toFixed(2)})</span>
                            </Grid.Column>
                            <Grid.Column>
                                {
                                    delta4 && delta4.bValue ?
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
                                <span>Delta4({delta4 && delta4.fValue.toFixed(2)})</span>
                            </Grid.Column>
                            <Grid.Column>
                                {
                                    delta5 && delta5.bValue ?
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
                                <span>Delta5({delta5 && delta5.fValue.toFixed(2)})</span>
                            </Grid.Column>
                            <Grid.Column></Grid.Column>
                        </Grid.Row>
                    </Grid>


                    <Grid style={{ margin: "2px" }}>
                        <Grid.Row columns={6}>
                            <Grid.Column>
                                {
                                    deviation && deviation.bValue ?
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
                                <span>Deviation({deviation && deviation.fValue.toFixed(2)})</span>
                            </Grid.Column>
                            <Grid.Column>
                                {
                                    volatility && volatility.bValue ?
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
                                <span>Volatility({volatility && volatility.fValue.toFixed(2)})</span>
                            </Grid.Column>
                            <Grid.Column></Grid.Column>
                            <Grid.Column></Grid.Column>
                            <Grid.Column></Grid.Column>
                            <Grid.Column></Grid.Column>
                        </Grid.Row>
                    </Grid>

                    <Grid style={{ margin: "2px" }}>
                        <Grid.Row columns={6}>
                            <Grid.Column>
                                {
                                    targetPercent && targetPercent.bValue ?
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
                                <span>Target%({targetPercent && targetPercent.fValue.toFixed(2)})</span>
                            </Grid.Column>
                            <Grid.Column>
                                {
                                    range52 && range52.bValue ?
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
                                <span>Range52({range52 && range52.fValue.toFixed(2)})</span>
                            </Grid.Column>
                            <Grid.Column>
                                {
                                    eps && eps.bValue ?
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
                                <span>EPS({eps && eps.fValue.toFixed(2)})</span>
                            </Grid.Column>
                            <Grid.Column>
                                {
                                    peRation && peRation.bValue ?
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
                                <span>P/E Ration({peRation && peRation.fValue.toFixed(2)})</span>
                            </Grid.Column>
                            <Grid.Column></Grid.Column>
                            <Grid.Column></Grid.Column>
                        </Grid.Row>
                    </Grid>
                </div>
            );
        }
        return (<div>Markers:</div>)
    }
}

export default Marker;
