import { useEffect, useState, useRef } from "react";
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory/ChatHistory.js";
import ChatInput from "./components/ChatInput/ChatInput.js";
import { Formik } from "formik";
import axios from "axios";
import { atom, useAtom } from "jotai";

export const ipAtom = atom<string | undefined>(undefined);

export const useIP = () => {
  const [IP, setIP] = useAtom(ipAtom);

  useEffect(() => {
    if (IP === undefined)
      axios
        .get("http://ip-api.com/json")
        .then((response) => {
          console.log("get ip response", response);
          let data = response.data || {};
          setIP(data.query ?? "");
        })
        .catch((err) => {
          console.log("get ip error", err);
          setIP("");
        });
  }, []);

  return IP;
};

const CHAT_HISTORY: string[] = [];

function useChatSocket() {
  const [chatHistory, setChatHistory] = useState<string[]>([]);
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    if (!ws.current) ws.current = new WebSocket("ws://localhost:8080/ws");
    ws.current.onopen = () => console.log("ws opened");
    ws.current.onclose = () => console.log("ws closed");

    ws.current.onmessage = (e) => {
      const message = JSON.parse(e.data);
      console.log(message);
      CHAT_HISTORY.push(message.body);
      setChatHistory([...CHAT_HISTORY]);
    };
  }, []);

  let sendMsg = (msg: string) => {
    if (!ws.current) return;
    console.log("sending msg: ", msg);
    ws.current.send(msg);
  };

  return { chatHistory, sendMsg };
}

function App() {
  const IP = useIP();
  const { chatHistory, sendMsg } = useChatSocket();
  if (IP === undefined) return <div>loading...</div>;
  if (IP === "") return <div>Error fetching your IP Address.</div>;
  return (
    <div className="App flex flex-col w-full h-full">
      <Header />
      <ChatHistory chatHistory={chatHistory} ip={IP} />
      {/* formik is a library for form handling */}
      <Formik
        initialValues={{ msg: "" }}
        validateOnBlur={false}
        validateOnChange={false}
        onSubmit={(values, actions) => {
          actions.setSubmitting(false);
          //console.log(values);
          sendMsg(values.msg);
          // clears form input field after msg sent
          actions.resetForm();
          // onSubmit(values)
        }}
      >
        {({ handleChange, handleSubmit, isSubmitting, values }) => (
          <form onSubmit={handleSubmit}>
            {/* send message on enter key press */}
            <ChatInput
              handleSubmit={handleSubmit}
              name="msg"
              value={values.msg}
              handleChange={handleChange}
            />
            <div className="flex mt-2 mr-3 justify-end md:mr-20">
              {/* submit message on button press */}
              <button
                type="submit"
                className="w-1/5 mt-1 text-xl border-2 p-2 bg-slate-300 rounded-xl"
              >
                Send
              </button>
            </div>
          </form>
        )}
      </Formik>
    </div>
  );
}

export default App;
