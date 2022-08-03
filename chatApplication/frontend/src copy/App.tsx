import { useEffect, useState, useRef } from "react"
import Header from "./components/Header/Header"
import ChatHistory from "./components/ChatHistory/ChatHistory.js"
import ChatInput from "./components/ChatInput/ChatInput.js"
import { Formik } from "formik"
import axios from "axios"
import { atom, useAtom } from "jotai"
import Lobby from "./components/ChatLobies/Lobby"

export const ipAtom = atom<string | undefined>(undefined)
export const useIP = () => {
  const [IP, setIP] = useAtom(ipAtom)

  // get the current user's IP address
  useEffect(() => {
    if (IP === undefined)
      axios
        .get("http://ip-api.com/json")
        .then((response) => {
          console.log("get ip response", response)
          let data = response.data || {}
          setIP(data.query ?? "")
        })
        .catch((err) => {
          console.log("get ip error", err)
          setIP("")
        })
  }, [])

  return IP
}

export type ChatMessage = {
  text: string // message text
  uip: string // user IP address
  timestamp: Date
  chatroom?: string
}

const CHAT_HISTORY: ChatMessage[] = []

// web socket hook
function useChatSocket() {
  const [chatHistory, setChatHistory] = useState<ChatMessage[]>([])
  const ws = useRef<WebSocket | null>(null)

  // handle socket events
  useEffect(() => {
    if (!ws.current) ws.current = new WebSocket("ws://localhost:8080/ws")
    ws.current.onopen = () => console.log("ws opened")
    ws.current.onclose = () => console.log("ws closed")

    ws.current.onmessage = (e) => {
      const message = JSON.parse(e.data)
      //console.log("chat message", JSON.parse(message.body));
      const chatMessage: ChatMessage = JSON.parse(message.body)
      CHAT_HISTORY.push(chatMessage)
      setChatHistory([...CHAT_HISTORY])
    }
  }, [])

  // stringify messae and send through socket
  let sendMsg = (msg: ChatMessage) => {
    if (!ws.current) return
    console.log("sending msg: ", msg)
    ws.current.send(JSON.stringify(msg))
  }

  function getChatroomHistory(chatroom?: string) {
    const roomHistory: ChatMessage[] = []
    chatHistory.forEach((message) => {
      if (message.chatroom === chatroom) roomHistory.push(message)
    })
    return roomHistory
  }

  return { chatHistory, sendMsg, getChatroomHistory }
}

function App() {
  const IP = useIP()
  const [currentChatroom, setCurrentChatroom] = useState<string | undefined>()
  const [currentChatroomHistory, setCurrentChatroomHistory] = useState<
    ChatMessage[]
  >([])
  const { chatHistory, getChatroomHistory, sendMsg } = useChatSocket()

  /// update history when chatroom is changed
  useEffect(() => {
    setCurrentChatroomHistory(getChatroomHistory(currentChatroom))
  }, [currentChatroom])

  useEffect(() => {
    setCurrentChatroomHistory(getChatroomHistory(currentChatroom))
  }, [chatHistory])

  if (IP === undefined) return <div>loading...</div>
  if (IP === "") return <div>Error fetching your IP Address.</div>

  const createNewChatroom = (name: string) => {
    for (let message of chatHistory) {
      if (message.chatroom === "name") {
        alert("There is already a chatroom called " + name)
        return
      }
    }
    sendMsg({
      uip: IP,
      text: `Welcome to ${name}`,
      timestamp: new Date(),
      chatroom: name,
    })
    setCurrentChatroom(name)
  }

  return (
    <div>
      <div className="chat__App flex flex-col w-full h-full">
        <Header />
        <div className="chat__ mainwindow flex h-full w-full">
          <div className="chat__lobby md:w-2/5 h-full bg-slate-200">
            <Lobby
              newChatroom={createNewChatroom}
              chatHistory={chatHistory}
              setChatroom={setCurrentChatroom}
            />
          </div>
          <div className="chat__chathistory w-full">
            <ChatHistory chatHistory={currentChatroomHistory} ip={IP} />
            {/* formik is a library for form handling */}
            <Formik
              initialValues={{ msg: "" }}
              validateOnBlur={false}
              validateOnChange={false}
              onSubmit={(values, actions) => {
                actions.setSubmitting(false)
                //console.log(values);
                sendMsg({
                  text: values.msg,
                  uip: IP,
                  timestamp: new Date(),
                  chatroom: currentChatroom,
                })
                // clears form input field after msg sent
                actions.resetForm()
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
                  <div className="chat__chathistory__submit flex mt-2 mr-3 justify-end md:mr-20">
                    {/* submit message on button press */}
                    <button
                      type="submit"
                      className="w-1/5 mt-1 text-xl border-2 p-2 mr-2 md:mr-7 bg-slate-300 rounded-xl"
                    >
                      Send
                    </button>
                  </div>
                </form>
              )}
            </Formik>
          </div>
        </div>
      </div>
    </div>
  )
}

export default App
