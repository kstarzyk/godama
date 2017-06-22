package main

import ui "github.com/gizak/termui"
import "database/sql"
import _ "github.com/lib/pq"
import "projects/dbmanager/wrap"
import "fmt"
import "strings"

const (
  username = "golang-app"
  password = "123456"
  dbname="bomb"
)

func TrimRightPath(path string) string {
  return strings.Split(path,"/")[2]
}
type Mode int
const (
  NORMAL Mode = iota
  QUERY
  INSPECT
  ERROR
)
func (m Mode) String() string {
  switch m {
  case NORMAL: return "NORMAL"
  case QUERY: return "QUERY"
  case INSPECT: return "INSPECT"
  }

  return "ERROR"
}
var db *sql.DB
func InitDB() error {
  var err error
  db, err = sql.Open("postgres", fmt.Sprintf("user=$1 password=$2 dbname=$3 sslmode=disable", username, password, dbname))
  if err != nil {
    return err
  }
  return nil
}

type App struct {
  dataTable *wrap.DataTable
  header *wrap.TextBox
  info *wrap.CList
  tables *wrap.CList
  sequences *wrap.CList
  focused wrap.Selectable
}

func (app *App) HandleKey(e ui.Event) {
  key := strings.Split(e.Path,"/")[3]
  if app.focused != nil {
    err := app.focused.HandleKey(key)
    if err != nil {
      panic(err)
    }
    app.Render()
  } else {
    defaultHandler(key)
  }
}

func defaultHandler(key string) {
  fmt.Println(key)
}

func NewApp() *App {
  app := &App{}
  return app
}

func (app *App) Init() error {
  app.dataTable = wrap.NewDataTable()
  app.dataTable.SetHeader([]string{"header1", "header2", "header3"})
  app.dataTable.SetData([][]string{
    []string{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
    []string{"2016", "11", "11"},
  })
  app.dataTable.SetPosition(0,20)
  app.dataTable.SetSize(62,7)

app.  header = wrap.NewTextBox("Golang Database Manager (" + username +":" + dbname + ")", "", true)
  app.info = wrap.NewCList("Info", []string{
    "ESC - Normal mode",
    "q - Query mode",
    "(T|t)ables",
    "(S|s)equences",
  })
  app.tables = wrap.NewCList("Tables", []string{
    "users",
    "items",
    "comments",
  })
  app.sequences = wrap.NewCList("Sequences", []string{
    "users_id_seq",
  })

  app.focused = app.tables
  return nil
}

func (app *App) Build() {
  ui.Body.AddRows(
    ui.NewRow(
      ui.NewCol(6,0,app.header.UI()),
      ui.NewCol(6,0,app.info.UI())),
    ui.NewRow(
      ui.NewCol(6,0,app.tables.UI()),
      ui.NewCol(6,0,app.sequences.UI())),
    ui.NewRow(
      ui.NewCol(6,0,app.dataTable.Table)))

  ui.Body.Align()

}

func (app *App) Render() {
  ui.Render(ui.Body)
}

func main() {
  err := InitDB()
  if err != nil {
    panic(err)
  }


  var agentid,bomberid int
  err = db.QueryRow("SELECT * FROM agenci").Scan(&agentid,&bomberid)
  fmt.Printf("AGENT:%d, BOMBER:%d,\n",agentid,bomberid)
  err = ui.Init()
  if err != nil {
    panic(err)
  }
  defer ui.Close()


  app := NewApp()
  err = app.Init()
  if err != nil {
    panic(err)
  }

  app.Build()
  app.Render()
  ui.Handle("/sys/kbd/q", func(ui.Event) {
    ui.StopLoop()
  })
  })
  ui.Handle("/sys/kbd", app.HandleKey)
  ui.Handle("/timer/1s", func(e ui.Event) {
    //t := e.Data.(ui.EvtTimer)
    //redraw(int(t.Count))
  })
  ui.Loop()
}
