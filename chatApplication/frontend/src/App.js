import React, { Component } from "react";
import "./App.css";
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory/ChatHistory.jsx";
import ChatInput from "./components/ChatInput/ChatInput.jsx";
import { connect, sendMsg } from "./api";

class App extends Component {
  constructor(props) {
    super(props);

    //define our starting chatHistory state
    this.state = {
      chatHistory: [],
    };
  }

  //function will be called automatically as part of our Components life-cycle.
  componentDidMount() {
    connect((msg) => {
      console.log("New Message");
      this.setState((prevState) => ({
        chatHistory: [...this.state.chatHistory, msg],
      }));
      console.log(this.state);
    });
  }

  // By passing in this event, we’ll be able to query if the key pressed was the Enter key,
  // if it is, we’ll be able to send the value of our <input/> field
  // to our WebSocket endpoint and then subsequently clear that <input/>:
  // send(event) {
  //   if (event.keyCode === 13) {
  //     sendMsg(event.target.value);
  //     event.target.value = "";
  //   }
  // }
  //event handler for "enter" key press
  textInputHandler(event) {
    if (event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }
  //event handler for clicking send
  clickHandler(_event) {
    // send message here
    sendMsg(_event.target.value);
    _event.target.value = _event;
  }

  render() {
    return (
      <div className="App flex flex-col w-full">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.textInputHandler} />
        <div className="flex mt-2 mr-3 justify-end md:mr-20">
          <button
            className="w-1/5 mt-1 text-xl border-2 p-2 bg-slate-300 rounded-xl"
            onClick={this.clickHandler}
          >
            Send
          </button>
        </div>
      </div>
    );
  }
}

export default App;
