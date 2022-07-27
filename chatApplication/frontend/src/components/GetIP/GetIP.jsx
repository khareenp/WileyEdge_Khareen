import React, { Component } from "react";
import axios from "axios";

class GetIP extends Component {
  constructor() {
    super();

    this.state = {
      ip: "",
      country: "",

      loading: false,
    };
  }

  componentDidMount = () => {
    this.getIpInfo();
  };

  getIpInfo = () => {
    let self = this;
    self.setState({
      loading: true,
    });

    axios
      .get("http://ip-api.com/json")
      .then((response) => {
        let data = response.data || {};
        self.setState({
          ip: data.query || "",
          country: data.country || "",
          loading: false,
        });
      })
      .catch((err) => {
        self.setState({
          loading: false,
        });
      });
  };

  render = () => {
    let self = this;
    if (self.state.loading) {
      return (
        <div className="App">
          <h1>Loading IP Address</h1>
        </div>
      );
    } else {
      return (
        <div className="App">
          <h1>
            IP: {self.state.ip}, {self.state.country}{" "}
          </h1>
        </div>
      );
    }
  };
}

export default GetIP;
