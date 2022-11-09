package refresh

import (
	"fmt"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
)

var injectionScriptPath string

type Injection struct {
	Inject string
}

const injectionContent = `
const url = 'ws://localhost:3000/blast/ws';
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
      window.location.reload();
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
`

func CreateRefresher(hot bool, dir string) {
	if !hot {
		return
	}
	base := dir
	injectionScriptPath = fmt.Sprintf("%s/blast-ws.js", base)
	f, err := os.Create(injectionScriptPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte(injectionContent))
}

func InjectScript(router *gin.Engine, indexPath string) {
	tmpl := template.Must(template.ParseFiles(indexPath))
	f, err := os.Create(indexPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	tmpl.Execute(f, Injection{Inject: "<script src=\"/blast-ws.js\" type=\"module\" defer></script>"})
	router.StaticFile("/", indexPath)
}
