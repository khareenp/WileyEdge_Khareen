import React, { Component } from "react";
import "./App.css";
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory/ChatHistory.jsx";
import ChatInput from "./components/ChatInput/ChatInput.jsx";
import { connect, sendMsg } from "./api";
import { Formik } from "formik";

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
        chatHistory: [...prevState.chatHistory, msg],
      }));
      console.log(this.state);
    });
  }

  render() {
    return (
      <div className="App flex flex-col w-full">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        {/* formik is a library for form handling */}
        <Formik
          initialValues={{ msg: "" }}
          validateOnBlur={false}
          validateOnChange={false}
          onSubmit={(values, actions) => {
            actions.setSubmitting(false);
            console.log(values);
            sendMsg(values.msg);
            // clears form input field after msg sent
            actions.resetForm();
            // onSubmit(values)
          }}
        >
          {({ handleChange, handleSubmit, isSubmitting, values }) => (
            <>
              {/* send message on enter key press */}
              <ChatInput
                send={(event) => {
                  if (event.keyCode === 13) {
                    handleSubmit();
                  }
                }}
                name="msg"
                value={values.msg}
                handleChange={handleChange}
              />
              <div className="flex mt-2 mr-3 justify-end md:mr-20">
                {/* submit message on button press */}
                <button
                  type="submit"
                  className="w-1/5 mt-1 text-xl border-2 p-2 bg-slate-300 rounded-xl"
                  onClick={handleSubmit}
                >
                  Send
                </button>
              </div>
            </>
          )}
        </Formik>
      </div>
    );
  }
}

export default App;
