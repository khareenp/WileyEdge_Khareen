import React, { Component } from "react";
import axios from "axios";

//This component is used to  get ip address and render to browser
class GetIP extends Component {
  constructor() {
    super();
    var today = new Date(),
      time =
        today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();

    this.state = {
      ip: "",
      // country: "",
      currentTime: time,
      // currentDateTime: Date().toLocaleString(),
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
          // country: data.country || "",
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
        <div className="App text-sm w-full">
          <h1>Loading User IP Address</h1>
        </div>
      );
    } else {
      return (
        <div className="App text-sm">
          {/* <p>{this.state.currentDateTime}</p> */}
          <p>Time stamp:{this.state.currentTime}</p>
          <h1>
            User: {self.state.ip}, {self.state.country}{" "}
          </h1>
        </div>
      );
    }
  };
}

export default GetIP;
