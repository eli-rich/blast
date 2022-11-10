package refresh

import (
	"fmt"
	"os"
	"text/template"
)

var injectionScriptPath string
var IndexPath string

type Injection struct {
	Inject string
}

const injectionContent = `
const url = 'ws://localhost:3000/ws:blast';
const ws = new WebSocket(url);

ws.onopen = () => {
  console.log('Connected to server');
  ws.send('hello');
};

ws.onmessage = event => {
  const command = event.data;
  switch (command) {
    case 'ping':
      ws.send('pong');
      break;
    case 'close':
      ws.close();
      break;
    case 'reload':
      window.location.reload(true);
      break;
    case 'hello':
      console.log('Hello from server');
      break;
    default:
      console.error('Unknown command', command);
  }
};

ws.onclose = () => {
  console.log('Disconnected from server');
};

window.onbeforeunload = () => {
  ws.close = () => {};
  ws.send('close');
  ws.close();
};


`

func CreateRefresher(hot bool, dir string) {
	IndexPath = dir + "/index.html"
	base := dir
	injectionScriptPath = fmt.Sprintf("%s/blast-ws.js", base)
	f, err := os.Create(injectionScriptPath)
	if err != nil {
		panic(err)
	}
	f.Write([]byte(injectionContent))
	f.Close()
	tmpl := template.Must(template.ParseFiles(IndexPath))
	f, err = os.Create(IndexPath)
	if err != nil {
		panic(err)
	}
	if !hot {
		tmpl.Execute(f, Injection{Inject: "<!-- {{.Inject}} -->"})
	} else {
		tmpl.Execute(f, Injection{Inject: "<script src=\"/blast-ws.js\" type=\"module\" defer></script>"})
	}
	f.Close()
}
